package interfaces

import (
	"authentication/pkg/domain"
	"authentication/pkg/utils/request"
	"context"
)

type UserRepository interface {
	Transactions(func(UserRepository) error) error

	UserSignUp(context.Context, request.UserSignUpRequest) (domain.User, error)
	FindUserByEmail(context.Context, string) (domain.User, error)
	GetAllUsers(context.Context) ([]domain.User, error)
	SearchUserByEmail(ctx context.Context, email string) ([]domain.User, error)
	SearchUserByUsername(ctx context.Context, email string) ([]domain.User, error)
	GetUserById(ctx context.Context, id int) (domain.User, error)
	ForgetPassword(ctx context.Context, obj request.ForgotPassword) error
}
