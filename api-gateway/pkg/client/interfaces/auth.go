package interfaces

import (
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/response"
	"github.com/labstack/echo/v4"
)

type AuthClient interface {
	UserSignUp(echo.Context, request.UserSignUpRequest) (*response.UserSignUpResponse, error)
	UserSignIn(echo.Context, request.UserSignInRequest) (*response.UserSignInResponse, error)
	// SearchUser(e echo.Context, req request.Search) ([]response.User, error)
}
