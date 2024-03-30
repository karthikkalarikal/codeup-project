package client

import (
	"fmt"

	client "github.com/karthikkalarikal/api-gateway/pkg/client/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/rpc/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/response"
	"github.com/labstack/echo/v4"
)

type userClientImpl struct {
	user interfaces.UserRPCService
	auth interfaces.AuthService
}

func NewUserClient(user interfaces.UserRPCService, auth interfaces.AuthService) client.UserClient {
	return &userClientImpl{
		user: user,
		auth: auth,
	}
}

// view all problems
func (u *userClientImpl) ViewAllProblems(in request.AllProbles) ([]response.Problem, error) {
	body, err := u.user.ViewAllProblems(in)

	if err != nil {
		return []response.Problem{}, err
	}

	return body, nil
}

// get problem by id
func (u *userClientImpl) GetProblemById(ctx echo.Context, in request.GetOneProblemById) (response.Problem, error) {
	fmt.Println("in usecase")
	body, err := u.user.GetProblemById(ctx, in)
	if err != nil {
		return response.Problem{}, err
	}
	return body, nil
}

// execute code by id
func (u *userClientImpl) ExecuteGoCodyById(ctx echo.Context, in request.SubmitCodeIdRequest) (code []byte, err error) {
	fmt.Println("in usecase exec code by id rpc")

	body, err := u.user.ExecuteGoCodyById(ctx, in)
	if err != nil {
		return
	}
	return body, nil
}

// forget password
func (u *userClientImpl) ForgetPassword(e echo.Context, req request.ForgotPassword) (response.User, error) {
	body, err := u.auth.ForgetPassword(e, req)
	if err != nil {
		return response.User{}, err
	}
	return body, nil
}

func (u *userClientImpl) GetProblemBy(e echo.Context, req request.SearchBy) ([]response.Problem, error) {
	fmt.Println("get problem ", req)
	body, err := u.user.SortProblemBy(e, req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (u *userClientImpl) MakePrime(e echo.Context, email string) error {
	fmt.Println("make prime")

	err := u.auth.MakePrime(e, email)
	if err != nil {
		return err
	}
	return nil
}

func (u *userClientImpl) UnSubscrbe(e echo.Context, id int) (response.User, error) {
	fmt.Println("unsubscrbe")
	body, err := u.auth.UnSubscribe(e, id)
	if err != nil {
		return response.User{}, err
	}
	return body, nil
}
