package interfaces

import (
	"authentication/pkg/domain"
	"context"
)

type AdminRepository interface {
	Transactions(tx func(repo AdminRepository) error) error

	GetUserById(ctx context.Context, id int) (domain.User, error)
	BlockUser(ctx context.Context, id int) (domain.User, error)
}
