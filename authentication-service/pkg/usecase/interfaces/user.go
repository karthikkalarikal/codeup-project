package interfaces

import (
	"authentication/pkg/domain"
	"authentication/pkg/utils/request"
	"authentication/pkg/utils/response"
	"context"
)

type UserUseCase interface {
	UserSignUp(context.Context, request.UserSignUpRequest) (domain.User, error)
	UserSignIn(context.Context, request.UserSignInRequest) (response.UserSignInResponse, error)
	GetAllUsers(context.Context) ([]domain.User, error)
	SearchTheUser(context.Context, request.Search) ([]domain.User, error)
	ForgotPassword(context.Context, request.ForgotPassword) (domain.User, error)
	// EmailVerify(context.Context, int) (string, error)
}
