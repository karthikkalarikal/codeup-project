package middleware

import (
	"errors"

	"github.com/golang-jwt/jwt"
	customerrors "github.com/karthikkalarikal/api-gateway/pkg/utils/customErrors"
	"github.com/labstack/echo/v4"
)

func UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("Authorization").(*jwt.Token) // the default value is user
		if !ok {
			return errors.New(customerrors.JwtTokenMissingError)
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return errors.New("failed to cast claims as jwt.MapClaims")
		}

		id, ok := claims["id"].(float64)
		if !ok || id == 0 {

			return errors.New("error error in retrieving id")

		}

		c.Set("id", int(id))
		return nil
	}
}
