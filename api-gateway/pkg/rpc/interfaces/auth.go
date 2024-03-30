package interfaces

import (
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/response"
	"github.com/labstack/echo/v4"
)

type AuthService interface {
	UserSignUp(echo.Context, request.UserSignUpRequest) (*response.UserSignUpResponse, error)
	UserSignIn(echo.Context, request.UserSignInRequest) (*response.UserSignInResponse, error)

	ViewUsers(e echo.Context) ([]response.User, error)
	SearchUser(e echo.Context, req request.Search) ([]response.User, error)
	BlockUser(e echo.Context, in int) (response.BlockedStatus, error)
	ForgetPassword(e echo.Context, in request.ForgotPassword) (response.User, error)

	Payment(e echo.Context, in request.Stripe) ([]byte, error)

	MakePrime(echo.Context, string) error
	UnSubscribe(e echo.Context, id int) (response.User, error)
}
