package usecase

import (
	"context"
	"fmt"
	"problem-service/pkg/domain"
	"problem-service/pkg/repository/interfaces"
	admin "problem-service/pkg/usecase/interfaces"
	"problem-service/pkg/utils/request"
	"time"
)

type adminUseCaseImpl struct {
	admin interfaces.AdminRepository
}

func NewAdminUseCase(admin interfaces.AdminRepository) admin.AdminUseCase {
	return &adminUseCaseImpl{
		admin: admin,
	}
}

// insert
func (u *adminUseCaseImpl) InsertProblem(ctx context.Context, req request.Problem) (domain.Problem, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	fmt.Println("inside usecase", req)
	defer cancel()
	id, err := u.admin.InsertProblem(ctxWithTimeout, req)
	fmt.Println("id inside usecase", id)
	if err != nil {
		return domain.Problem{}, err
	}
	body, err := u.admin.GetProblemById(ctxWithTimeout, id)
	fmt.Println("body inisde usecase", body)
	if err != nil {
		return domain.Problem{}, err
	}
	return body, nil
}

// insert first half
func (u *adminUseCaseImpl) InsertFirstHalfProblem(ctx context.Context, entry request.FirstHalfCode) (domain.Problem, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	fmt.Println("inside usecase", entry)
	defer cancel()

	out, err := u.admin.InsertFirstHalfProblem(ctxWithTimeout, entry)
	fmt.Println("id inside usecase", out)
	if err != nil {
		return domain.Problem{}, err
	}

	return out, nil
}

// insert second half
func (u *adminUseCaseImpl) InsertSecondHalfProblem(ctx context.Context, entry request.SecondHalfCode) (domain.Problem, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	fmt.Println("inside usecase", entry)
	defer cancel()

	out, err := u.admin.InsertSecondHalfProblem(ctxWithTimeout, entry)
	fmt.Println("id inside usecase", out)
	if err != nil {
		return domain.Problem{}, err
	}

	return out, nil
}
