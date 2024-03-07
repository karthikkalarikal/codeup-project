package client

import (
	"fmt"

	"github.com/karthikkalarikal/api-gateway/pkg/rpc/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/response"
	"github.com/labstack/echo/v4"

	client "github.com/karthikkalarikal/api-gateway/pkg/client/interfaces"
)

type adminClientImpl struct {
	user client.AdminClient
}

func NewAdminClient(user interfaces.AdminRPCService) client.AdminClient {
	return &adminClientImpl{
		user: user,
	}
}

func (u *adminClientImpl) InsertProblem(e echo.Context, req request.InsertProblem) (response.Problem, error) {

	body, err := u.user.InsertProblem(e, req)
	fmt.Println("here in inser problem client", body)
	if err != nil {
		return response.Problem{}, err
	}
	return body, nil
}
