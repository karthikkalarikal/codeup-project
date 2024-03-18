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
	user interfaces.AdminRPCService
	auth interfaces.AuthService
}

func NewAdminClient(user interfaces.AdminRPCService, auth interfaces.AuthService) client.AdminClient {
	return &adminClientImpl{
		user: user,
		auth: auth,
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

// insert first half
func (u *adminClientImpl) InsertFirstHalfProblem(e echo.Context, in request.FirstHalfCode) (response.InsertProblem, error) {
	body, err := u.user.InsertFirstHalfProblem(e, in)
	if err != nil {
		return response.InsertProblem{}, err
	}
	return body, nil
}

// insert second half
func (u *adminClientImpl) InsertSecondHalfProblem(e echo.Context, in request.SecondHalfCode) (response.InsertProblem, error) {
	body, err := u.user.InsertSecondHalfProblem(e, in)
	if err != nil {
		return response.InsertProblem{}, err
	}
	return body, nil
}

// get all the users
func (u *adminClientImpl) ViewUsers(e echo.Context) ([]response.User, error) {
	body, err := u.auth.ViewUsers(e)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// search users
func (a *adminClientImpl) SearchUser(e echo.Context, req request.Search) ([]response.User, error) {
	res, err := a.auth.SearchUser(e, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// block user
func (a *adminClientImpl) BlockUser(e echo.Context, id int) (response.BlockedStatus, error) {
	res, err := a.auth.BlockUser(e, id)
	if err != nil {
		return response.BlockedStatus{}, err
	}
	return res, nil
}
