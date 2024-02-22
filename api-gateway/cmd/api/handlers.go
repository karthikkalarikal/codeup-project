package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type requestPayload struct {
	Action  string         `json:"action"`
	Auth    AuthPayload    `json:"auth"`
	Problem problemPayload `json:"problem,omitempty"` //in case there is a problem with auth comment this line
}

// auth
type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// test
type TestCase struct {
	Input  string `bson:"input" json:"input"`
	Output string `bson:"output" json:"output"`
}

// test
type problemPayload struct {
	Title       string     `bson:"title" json:"title"`
	Description string     `bson:"description" json:"description"`
	Difficulty  string     `bson:"difficulty" json:"difficulty"`
	TestCases   []TestCase `bson:"test_cases" json:"test_cases"`
	TimeLimit   int        `bson:"time_limit" json:"time_limit"`
	MemoryLimit int        `bson:"memory_limit" json:"memory_limit"`
	Tags        []string   `bson:"tags" json:"tags"`
}

// sign up
type requestSignUpPayload struct {
	Action string      `json:"action"`
	SignUp SignUpModel `json:"signup"`
}

// sign up
type SignUpModel struct {
	Username        sql.NullString `json:"username"`
	Email           sql.NullString `json:"email"`
	Password        string         `json:"password"`
	ConfirmPassword string         `json:"confirmpassword"`
	Name            sql.NullString `json:"name"`
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
		fmt.Println("case auth:")
		fmt.Println("payload", requestPayload)
		app.authenticate(c, requestPayload)
	case "problem":
		fmt.Println("problem")
		app.problemItem(c, requestPayload.Problem)

	default:
		app.errorJSON(c, errors.New("unknown error"))
	}
	return nil
}

func (app *Config) authenticate(c echo.Context, a requestPayload) error {
	// create some json we'll send to auth microservice
	fmt.Println("a: ", a)
	jsonData, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		log.Println("error in auth api gateway", err)

	}
	fmt.Println("jsonData ", string(jsonData))

	// call service
	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
	fmt.Println("request---", request)
	// fmt.Println("error:  ", err)
	if err != nil {
		app.errorJSON(c, err)
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	//logging
	fmt.Printf("Request Details:\n%s %s\n", request.Method, request.URL.String())
	fmt.Printf("Request Headers: %v\n", request.Header)
	fmt.Printf("Request Body: %s\n", jsonData)

	client := &http.Client{}
	fmt.Println("request", request)
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(c, err)
		return err
	}
	fmt.Println("response --- ", response, "err ::", err)
	// {
	// 	fmt.Println("here 1")
	// 	body, err := io.ReadAll(response.Body)

	// 	log.Printf("Response Body: %s\n err: %s", body, err)
	// }
	fmt.Println("here 2")
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
	var user User

	// decode teh json from auth service

	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		fmt.Println("here error in decoding")
		app.errorJSON(c, err)
		return err
	}
	fmt.Println("body ", user)

	if jsonFromService.Error {
		app.errorJSON(c, err, http.StatusUnauthorized)
		return err
	}

	var payload jsonResponse

	payload.Error = false
	payload.Message = "Authenticated"
	payload.Data = user

	app.writeJSON(c, http.StatusAccepted, payload)
	return nil
}

func (app *Config) signUp(c echo.Context, a requestSignUpPayload) error {
	jsonData, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		log.Println("error in auth api gateway", err)

	}
	fmt.Println("jsonData ", string(jsonData))

	// call service
	request, err := http.NewRequest("POST", "http://authentication-service/signup", bytes.NewBuffer(jsonData))
	fmt.Println("request---", request)
	// fmt.Println("error:  ", err)
	if err != nil {
		app.errorJSON(c, err)
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	//logging
	fmt.Printf("Request Details:\n%s %s\n", request.Method, request.URL.String())
	fmt.Printf("Request Headers: %v\n", request.Header)
	fmt.Printf("Request Body: %s\n", jsonData)

	client := &http.Client{}
	fmt.Println("request", request)
	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(c, err)
		return err
	}
	fmt.Println("response --- ", response, "err ::", err)
	// {
	// 	fmt.Println("here 1")
	// 	body, err := io.ReadAll(response.Body)

	// 	log.Printf("Response Body: %s\n err: %s", body, err)
	// }
	fmt.Println("here 2")
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
	var user User

	// decode teh json from auth service

	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		fmt.Println("here error in decoding")
		app.errorJSON(c, err)
		return err
	}
	fmt.Println("body ", user)

	if jsonFromService.Error {
		app.errorJSON(c, err, http.StatusUnauthorized)
		return err
	}

	var payload jsonResponse

	payload.Error = false
	payload.Message = "Authenticated"
	payload.Data = user

	app.writeJSON(c, http.StatusAccepted, payload)
	return nil
}

// problem service

func (app *Config) problemItem(c echo.Context, entry problemPayload) error {
	jsonData, _ := json.MarshalIndent(entry, "", "\t") // user marshal in production,

	problemServiceURL := "http://problem-service/write"

	request, err := http.NewRequest("POST", problemServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(c, err)
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.errorJSON(c, err)
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		app.errorJSON(c, err)
		return err
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "inserted"

	app.writeJSON(c, http.StatusAccepted, payload)

	return nil
}
