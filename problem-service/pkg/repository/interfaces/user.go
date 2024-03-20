package interfaces

import (
	"context"
	"problem-service/pkg/domain"
	"problem-service/pkg/utils/request"
)

type UserRepository interface {
	ViewAllProblems(context.Context) ([]domain.Problem, error)
	GetProblemById(ctx context.Context, id request.ProblemById) (domain.Problem, error)
	GetProblemByDifficulty(ctx context.Context, difficulty string) ([]domain.Problem, error)
}
