package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type requestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Api(c echo.Context) error {

	payload := jsonResponse{
		Error:   false,
		Message: "Hit the api",
	}

	err := app.writeJSON(c, http.StatusOK, payload)

	return err
	// out, _ := json.MarshalIndent(payload, "", "\t")
	// c.Response().Header().Set("Content-Type", "application/json")
	// c.Response().WriteHeader(http.StatusAccepted)

	// return c.String(http.StatusAccepted, string(out))

}

func (app *Config) HandleSubmission(c echo.Context) error {
	var requestPayload requestPayload

	err := app.readJSON(c, &requestPayload)

	if err != nil {
		app.errorJSON(c, err)
		return err
	}

	switch requestPayload.Action {
	case "auth":
		app.authenticate(c, requestPayload.Auth)

	default:
		app.errorJSON(c, errors.New("unknown error"))
	}
	return nil
}

func (app *Config) authenticate(c echo.Context, a AuthPayload) error {
	// create some json we'll send to auth microservice

	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// call service
	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
	fmt.Println("request---", request)
	fmt.Println("error:  ", err)
	if err != nil {
		app.errorJSON(c, err)
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	fmt.Println("request", request)
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(c, err)
		return err
	}
	fmt.Println("response --- ", response, "err ::", err)
	defer response.Body.Close()
	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(c, errors.New("invalid credentials"))
		return err
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(c, errors.New("error calling auth service"))

		return err
	}

	// create a variabel that read response.Body into

	var jsonFromService jsonResponse

	// decode teh json from auth service

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.errorJSON(c, err)
		return err
	}

	if jsonFromService.Error {
		app.errorJSON(c, err, http.StatusUnauthorized)
		return err
	}

	var payload jsonResponse

	payload.Error = false
	payload.Message = "Authenticated"
	payload.Data = jsonFromService.Data

	app.writeJSON(c, http.StatusAccepted, payload)
	return nil
}
