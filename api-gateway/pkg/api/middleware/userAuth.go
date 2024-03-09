package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	customerrors "github.com/karthikkalarikal/api-gateway/pkg/utils/customErrors"
	"github.com/labstack/echo/v4"
)

// type _getHeader struct {
// 	Token string `header: "Authorization"`
// }

func UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// a := new(_getHeader)
		// b := &echo.DefaultBinder{}
		// b.BindHeaders(c, a)
		// fmt.Println("header token", a.Token)
		// fmt.Println("auth", c.Request().Header.Get("Authorization"))
		tokenString := c.Request().Header.Get("Authorization") // the default value is user
		// if !ok {
		// 	c.JSON(http.StatusUnauthorized, echo.Map{"error": customerrors.JwtTokenMissingError})
		// 	return errors.New(customerrors.JwtTokenMissingError)
		// }

		fmt.Println("token ", tokenString)

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, echo.Map{"error": customerrors.JwtTokenMissingError})
			return errors.New(customerrors.JwtTokenMissingError)
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, echo.Map{"error": customerrors.JwtTokenMissingError})
			return err

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
		fmt.Println("id", id)
		return next(c)
	}
}

// admin auth

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// a := new(_getHeader)
		// b := &echo.DefaultBinder{}
		// b.BindHeaders(c, a)
		// fmt.Println("header token", a.Token)
		// fmt.Println("auth", c.Request().Header.Get("Authorization"))
		tokenString := c.Request().Header.Get("Authorization") // the default value is user
		// if !ok {
		// 	c.JSON(http.StatusUnauthorized, echo.Map{"error": customerrors.JwtTokenMissingError})
		// 	return errors.New(customerrors.JwtTokenMissingError)
		// }

		fmt.Println("token ", tokenString)

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, echo.Map{"error": customerrors.JwtTokenMissingError})
			return errors.New(customerrors.JwtTokenMissingError)
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, echo.Map{"error": customerrors.JwtTokenMissingError})
			return err

		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return errors.New("failed to cast claims as jwt.MapClaims")
		}

		id, ok := claims["id"].(float64)
		if !ok || id == 0 {

			return errors.New("error error in retrieving id")

		}
		admin, ok := claims["admin"].(bool)
		if !ok || !admin {
			return errors.New("authorization error")
		}

		// c.Set("id", int(id))

		return next(c)
	}
}
