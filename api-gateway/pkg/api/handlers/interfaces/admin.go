package interfaces

import (
	"github.com/labstack/echo/v4"
)

type AdminHandler interface {
	CreateProblem(echo.Context) error
}
