package interfaces

import (
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/response"
)

type UserClient interface {
	ViewAllProblems(in request.AllProbles) ([]response.Problem, error)
}
