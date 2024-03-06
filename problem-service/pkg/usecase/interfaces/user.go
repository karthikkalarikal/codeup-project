package interfaces

import (
	"context"
	"problem-service/pkg/domain"
)

type UserUseCase interface {
	ViewAllProblems(context.Context) ([]domain.Problem, error)
}
