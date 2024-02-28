package main

// import (
// 	"encoding/json"
// 	"errors"
// 	"io"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// type Config struct{}
// type jsonResponse struct {
// 	Error   bool   `json:"error"`
// 	Message string `json:"message"`
// 	Data    any    `json:"data,omitempty"`
// }

// type User struct {
// 	ID        int64      `json:"id"`
// 	Username  NullString `json:"username"`
// 	Email     NullString `json:"email"`
// 	Name      NullString `json:"name"`
// 	CreatedAt Timestamp  `json:"created_at"`
// 	UpdatedAt Timestamp  `json:"updated_at"`
// }

// type NullString struct {
// 	String string `json:"String"`
// 	Valid  bool   `json:"Valid"`
// }

// type Timestamp struct {
// 	Time  string `json:"Time"`
// 	Valid bool   `json:"Valid"`
// }

// func (app *Config) readJSON(c echo.Context, data any) error {
// 	maxBytes := 1048576 //one megabyte

// 	c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, int64(maxBytes))

// 	dec := json.NewDecoder(c.Request().Body)
// 	err := dec.Decode(data)

// 	if err != nil {
// 		return err
// 	}

// 	err = dec.Decode(&struct{}{})
// 	if err != io.EOF {
// 		return errors.New("body must have only a single JSON value")

// 	}
// 	return nil
// }

// func (app *Config) writeJSON(c echo.Context, status int, data any, headers ...http.Header) error {
// 	out, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}

// 	if len(headers) > 0 {
// 		for key, value := range headers[0] {
// 			c.Response().Header()[key] = value
// 		}
// 	}

// 	c.Response().Header().Set("Content-Type", "application/json")
// 	c.Response().WriteHeader(status)
// 	_, err = c.Response().Write(out)

// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }

// func (app *Config) ErrorJson(c echo.Context, err error, status ...int) error {
// 	statusCode := http.StatusBadRequest

// 	if len(status) > 0 {
// 		statusCode = status[0]
// 	}

// 	var payload jsonResponse
// 	payload.Error = true
// 	payload.Message = err.Error()

// 	return app.writeJSON(c, statusCode, payload)
// }
