package interfaces

import (
	"authentication/pkg/domain"
	"authentication/pkg/utils/request"
	"context"
)

type UserUseCase interface {
	UserSignUp(context.Context, request.UserSignUpRequest) (domain.User, error)
}
