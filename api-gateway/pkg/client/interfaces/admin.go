package interfaces

import (
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/response"
	"github.com/labstack/echo/v4"
)

type AdminClient interface {
	InsertProblem(e echo.Context, req request.InsertProblem) (response.Problem, error)
	InsertFirstHalfProblem(e echo.Context, in request.FirstHalfCode) (response.InsertProblem, error)
	InsertSecondHalfProblem(e echo.Context, in request.SecondHalfCode) (response.InsertProblem, error)
	ViewUsers(e echo.Context) ([]response.User, error)
	SearchUser(e echo.Context, req request.Search) ([]response.User, error)
}
