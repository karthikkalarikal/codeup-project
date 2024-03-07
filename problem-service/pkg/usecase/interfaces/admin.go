package interfaces

import (
	"context"
	"problem-service/pkg/domain"
	"problem-service/pkg/utils/request"
)

type AdminUseCase interface {
	InsertProblem(ctx context.Context, req request.Problem) (domain.Problem, error)
}
