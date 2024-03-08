package interfaces

import (
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/response"
	"github.com/labstack/echo/v4"
)

type UserRPCService interface {
	ViewAllProblems(request.AllProbles) ([]response.Problem, error)
	GetProblemById(ctx echo.Context, in request.GetOneProblemById) (response.Problem, error)
}
