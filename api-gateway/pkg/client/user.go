package client

import (
	client "github.com/karthikkalarikal/api-gateway/pkg/client/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/rpc/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/response"
	"github.com/labstack/echo/v4"
)

type userClientImpl struct {
	user client.UserClient
}

func NewUserClient(user interfaces.UserRPCService) client.UserClient {
	return &userClientImpl{
		user: user,
	}
}

func (u *userClientImpl) ViewAllProblems(in request.AllProbles) ([]response.Problem, error) {
	body, err := u.user.ViewAllProblems(in)
	if err != nil {
		return []response.Problem{}, err
	}

	return body, nil
}

func (u *userClientImpl) GetProblemById(ctx echo.Context, in request.GetOneProblemById) (response.Problem, error) {
	body, err := u.user.GetProblemById(ctx, in)
	if err != nil {
		return response.Problem{}, err
	}
	return body, nil
}
