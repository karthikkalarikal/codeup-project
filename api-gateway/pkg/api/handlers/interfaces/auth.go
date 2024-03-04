package interfaces

import "github.com/labstack/echo/v4"

type AuthHandler interface {
	UserSignUp(e echo.Context) error
	UserSignIn(e echo.Context) error
}
