package interfaces

import (
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/response"
	"github.com/labstack/echo/v4"
)

type UserClient interface {
	ViewAllProblems(in request.AllProbles) ([]response.Problem, error)
	GetProblemById(ctx echo.Context, in request.GetOneProblemById) (response.Problem, error)
	ExecuteGoCodyById(ctx echo.Context, in request.SubmitCodeIdRequest) (code []byte, err error)
	ForgetPassword(echo.Context, request.ForgotPassword) (response.User, error)
	GetProblemBy(e echo.Context, req request.SearchBy) ([]response.Problem, error)
}
