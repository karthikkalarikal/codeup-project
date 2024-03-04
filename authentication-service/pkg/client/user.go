package client

import (
	"authentication/pkg/domain"
	user "authentication/pkg/usecase/interfaces"
	"authentication/pkg/utils/request"
	"authentication/pkg/utils/response"
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

func (u *AuthUserService) UserSignIn(req request.UserSignInRequest, reply *response.UserSignInResponse) (err error) {
	ctx := context.Background()
	body, err := u.useCase.UserSignIn(ctx, req)

	if err != nil {
		return
	}

	*reply = body
	return nil
}
