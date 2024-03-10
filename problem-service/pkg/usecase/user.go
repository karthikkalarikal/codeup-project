package usecase

import (
	"context"
	"problem-service/pkg/domain"
	"problem-service/pkg/repository/interfaces"
	user "problem-service/pkg/usecase/interfaces"
	"problem-service/pkg/utils/request"
)

type userUseCase struct {
	repo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) user.UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

//view all

func (u *userUseCase) ViewAllProblems(ctx context.Context) ([]domain.Problem, error) {
	body, err := u.repo.ViewAllProblems(ctx)
	if err != nil {
		return []domain.Problem{}, err
	}
	return body, nil
}

// get one

func (u *userUseCase) GetProblemById(ctx context.Context, id request.ProblemById) (domain.Problem, error) {
	body, err := u.repo.GetProblemById(ctx, id)
	if err != nil {
		return domain.Problem{}, err
	}
	return body, nil
}

func (u *userUseCase) SubmitCodeById(ctx context.Context, req request.SubmitCodeIdRequest) ([]byte, error) {
	return nil, nil
}
