package interfaces

import (
	"context"
	"problem-service/pkg/domain"
)

type UserRepository interface {
	ViewAllProblems(context.Context) ([]*domain.Problem, error)
}
