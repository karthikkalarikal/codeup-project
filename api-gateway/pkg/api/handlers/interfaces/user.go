package interfaces

import (
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	ViewAllProblems(e echo.Context) error
	GetOneProblemById(e echo.Context) error
	WriteCode(e echo.Context) error
	ExecuteGoCodyById(echo.Context) error
	ForgetPassword(echo.Context) error
	GetProblemBy(e echo.Context) error
	Payment(e echo.Context) error
	GetPaymentIntent(e echo.Context) error
	PaymentSuccess(e echo.Context) error
	UnSubscrbe(e echo.Context) error
}
