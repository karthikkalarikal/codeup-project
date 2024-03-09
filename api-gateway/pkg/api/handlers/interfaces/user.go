package interfaces

import "github.com/labstack/echo/v4"

type UserHandler interface {
	ViewAllProblems(e echo.Context) error
	GetOneProblemById(e echo.Context) error
	GetOneProblemById(e echo.Context) error
	WriteCode(e echo.Context) error
}
