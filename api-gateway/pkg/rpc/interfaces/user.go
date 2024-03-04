package interfaces

import (
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/response"
)

type UserRPCService interface {
	ViewAllProblems(request.AllProbles) ([]response.Problem, error)
}
