package usecase

import (
	"context"
	"problem-service/pkg/domain"
	"problem-service/pkg/repository/interfaces"
	user "problem-service/pkg/usecase/interfaces"
)

type userUseCase struct {
	repo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) user.UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (u *userUseCase) ViewAllProblems(ctx context.Context) ([]domain.Problem, error) {
	body, err := u.repo.ViewAllProblems(ctx)
	if err != nil {
		return []domain.Problem{}, err
	}
	return body, nil
}
