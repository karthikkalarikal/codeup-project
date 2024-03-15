package interfaces

import (
	"github.com/labstack/echo/v4"
)

type AdminHandler interface {
	CreateProblem(echo.Context) error
	InsertFirstHalfProblem(echo.Context) error
	InsertSecondHalfProblem(echo.Context) error
	ViewUsers(echo.Context) error
	SearchUser(echo.Context) error
}
