package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) readJSON(c echo.Context, data any) error {
	maxBytes := 1048576 //one megabyte

	c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, int64(maxBytes))

	dec := json.NewDecoder(c.Request().Body)
	fmt.Println("dec ", dec)
	err := dec.Decode(data)
	fmt.Println("data", data)

	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single JSON value")

	}
	return nil
}

func (app *Config) writeJSON(c echo.Context, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			c.Response().Header()[key] = value
		}
	}

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(status)
	_, err = c.Response().Write(out)

	if err != nil {
		return err
	}

	return nil

}

func (app *Config) ErrorJson(c echo.Context, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.writeJSON(c, statusCode, payload)
}
