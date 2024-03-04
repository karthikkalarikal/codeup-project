package interfaces

import (
	"authentication/pkg/domain"
	"authentication/pkg/utils/request"
	"context"
)

type UserRepository interface {
	UserSignUp(context.Context, request.UserSignUpRequest) (domain.User, error)
	FindUserByEmail(context.Context, string) (domain.User, error)
}
