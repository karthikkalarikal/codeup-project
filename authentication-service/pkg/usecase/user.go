package usecase

import (
	"authentication/pkg/domain"
	repo "authentication/pkg/repository/interfaces"
	"authentication/pkg/usecase/interfaces"
	"authentication/pkg/utils/request"
	"authentication/pkg/utils/response"
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/copier"

	"gorm.io/gorm"
)

type userUseCase struct {
	repo repo.UserRepository
}

func NewUserUseCase(repo repo.UserRepository) interfaces.UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (u *userUseCase) UserSignUp(ctx context.Context, user request.UserSignUpRequest) (domain.User, error) {
	body, err := u.repo.UserSignUp(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return body, nil
}

func (u *userUseCase) UserSignIn(ctx context.Context, user request.UserSignInRequest) (out response.UserSignInResponse, err error) {

	ctxDeadline, cancel := context.
		WithTimeout(ctx, 5*time.Second)

	defer cancel()

	body, err := u.repo.FindUserByEmail(ctxDeadline, user.Email)

	fmt.Println("err ", err, "gorm error:", gorm.ErrRecordNotFound)

	if err != nil {
		return
	}
	err = copier.CopyWithOption(&out, &body, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return
	}
	fmt.Println("out: ", out)
	return out, nil
}
