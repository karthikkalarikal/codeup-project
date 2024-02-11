package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

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
