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

func NewUserService(user user.UserUseCase) *AuthUserService {
	return &AuthUserService{
		useCase: user,
	}
}

func (u *AuthUserService) SignUp(req request.UserSignUpRequest, reply *domain.User) error {
	ctx := context.Background()
	body, err := u.useCase.UserSignUp(ctx, req)

	if err != nil {
		return err
	}

	*reply = body
	return nil
}
