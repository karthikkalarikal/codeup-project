package interfaces

import (
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/response"
	"github.com/labstack/echo/v4"
)

type AdminRPCService interface {
	InsertProblem(echo.Context, request.InsertProblem) (response.Problem, error)
}
