package client

import (
	"github.com/karthikkalarikal/api-gateway/pkg/client/interfaces"
	service "github.com/karthikkalarikal/api-gateway/pkg/rpc/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/response"
	"github.com/labstack/echo/v4"
)

type authClientImpl struct {
	client interfaces.AuthClient
}

func NewAuthClient(service service.AuthService) interfaces.AuthClient {
	return &authClientImpl{
		client: service,
	}
}

func (auth *authClientImpl) UserSignUp(e echo.Context, in request.UserSignUpRequest) (*response.UserSignUpResponse, error) {
	res, err := auth.client.UserSignUp(e, in)
	if err != nil {
		return &response.UserSignUpResponse{}, err
	}
	return res, err

}

func (auth *authClientImpl) UserSignIn(ctx echo.Context, in request.UserSignInRequest) (*response.UserSignInResponse, error) {
	res, err := auth.client.UserSignIn(ctx, in)
	if err != nil {
		return &response.UserSignInResponse{}, err
	}
	return res, err
}
