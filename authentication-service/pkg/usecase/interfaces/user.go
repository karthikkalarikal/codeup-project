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
	SearchTheUser(ctx context.Context, s request.Search) ([]domain.User, error)
}
