package usecase

import (
	"authentication/pkg/domain"
	repo "authentication/pkg/repository/interfaces"
	"authentication/pkg/usecase/interfaces"
	"context"
	"fmt"
	"time"
)

type adminUsecaseImpl struct {
	repo repo.AdminRepository
}

func NewAdminUsecase(repo repo.AdminRepository) interfaces.AdminUsecase {
	return &adminUsecaseImpl{
		repo: repo,
	}
}

func (a *adminUsecaseImpl) BlockUser(ctx context.Context, id int) (domain.User, error) {
	fmt.Println("in usecase block user")
	ctxDeadline, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	body := domain.User{}
	var err error
	err = a.repo.Transactions(func(repo repo.AdminRepository) error {
		body, err = a.repo.BlockUser(ctxDeadline, id)
		if err != nil {
			return err
		}
		return nil
	})
	fmt.Println("body in usecase ", body)
	if err != nil {
		return domain.User{}, err
	}
	// body, err = a.repo.BlockUser(ctxDeadline, id)
	// if err != nil {
	// 	return domain.User{}, err
	// }
	return body, nil
}
