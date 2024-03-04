package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

type Utils struct {
}

func NewUtils() *Utils {
	return &Utils{}
}

type JwtCustomClaims struct {
	Id int `json:"id"`
	// Admin bool   `json:"admin"`

	jwt.RegisteredClaims
}

func (u *Utils) GetTokenString(userId int) (string, error) {

	claims := &JwtCustomClaims{
		userId,

		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Utils) ReadJSON(c echo.Context, data any) error {
	maxBytes := 1048576 //one megabyte

	c.Request().Body = http.MaxBytesReader(c.Response(), c.Request().Body, int64(maxBytes))

	dec := json.NewDecoder(c.Request().Body)
	err := dec.Decode(data)

	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single JSON value")

	}
	return nil
}

func (app *Utils) WriteJSON(c echo.Context, status int, data any, headers ...http.Header) error {
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

func (app *Utils) ErrorJson(c echo.Context, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.WriteJSON(c, statusCode, payload)
}
