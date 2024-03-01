package client

import (
	"authentication/pkg/domain"
	user "authentication/pkg/usecase/interfaces"
	"authentication/pkg/utils/request"
	"context"
)

type AuthUserService struct {
	useCase user.UserUseCase
}

func NewUserService(user user.UserUseCase) AuthUserService {
	return AuthUserService{
		useCase: user,
	}
}

func (u *AuthUserService) SignUp(ctx context.Context, req request.UserSignUpRequest) (domain.User, error) {
	body, err := u.useCase.UserSignUp(ctx, req)

	if err != nil {
		return domain.User{}, err
	}
	return body, nil
}
