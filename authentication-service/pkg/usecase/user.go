package usecase

import (
	"authentication/pkg/domain"
	repo "authentication/pkg/repository/interfaces"
	"authentication/pkg/usecase/interfaces"
	"authentication/pkg/utils/request"
	"context"
)

type userUseCase struct {
	repo repo.UserRepository
}

func NewUserUseCase(repo repo.UserRepository) interfaces.UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (u *userUseCase) UserSignUp(ctx context.Context, user request.UserSignUpRequest) (domain.User, error) {
	body, err := u.repo.UserSignUp(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return body, nil
}
