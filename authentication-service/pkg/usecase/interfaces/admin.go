package interfaces

import (
	"authentication/pkg/domain"
	"context"
)

type AdminUsecase interface {
	BlockUser(ctx context.Context, id int) (domain.User, error)
}
