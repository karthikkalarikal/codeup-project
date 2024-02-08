package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) Api(c echo.Context) error {

	payload := jsonResponse{
		Error:   false,
		Message: "Hit the api",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusAccepted)

	return c.String(http.StatusAccepted, string(out))

}
