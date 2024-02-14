package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Authenticate(e echo.Context) error {
	fmt.Println("here---")
	var payload RequestPayload
	// e.Request().Body
	err := e.Bind(&payload)
	// err := app.readJSON(e, &payload)
	// fmt.Println("err in auth --:", err)
	if err != nil {
		app.errorJSON(e, err, http.StatusBadRequest)
		return err
	}
	fmt.Println("payload", payload)
	// validate the user against the database

	user, err := app.Models.User.GetByEmail(payload.Auth.Email)
	fmt.Println("user:", user)
	if err != nil {
		app.errorJSON(e, err, http.StatusBadRequest)
		return err
	}
	app.writeJSON(e, http.StatusAccepted, user)
	return nil
}
