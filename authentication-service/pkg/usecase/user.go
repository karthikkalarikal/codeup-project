package usecase

import (
	"authentication/pkg/domain"
	repo "authentication/pkg/repository/interfaces"
	"authentication/pkg/usecase/interfaces"
	"authentication/pkg/utils/request"
	"authentication/pkg/utils/response"
	"context"
	"errors"
	"fmt"
	"time"

	emailverifier "github.com/AfterShip/email-verifier"
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

func (u *userUseCase) UserSignUp(ctx context.Context, user request.UserSignUpRequest) (body domain.User, err error) {
	ctxDeadline, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	// ch := make(chan error)
	err = u.repo.Transactions(func(ur repo.UserRepository) error {
		if err = EmailVerify(ctxDeadline, body); err != nil {
			return err
		}

		body, err = u.repo.UserSignUp(ctxDeadline, user)
		if err != nil {
			return err
		}
		return nil
		// select {
		// case emailErr := <-ch:
		// 	if emailErr != nil {
		// 		return emailErr
		// 	}
		// case <-ctxDeadline.Done():
		// 	return ctxDeadline.Err()
		// default:
		// 	return errors.New("time out")
		// }
		// return <-ch

	})
	if err != nil {
		return
	}
	return body, nil
	// body, err := u.repo.UserSignUp(ctx, user)
	// if err != nil {
	// 	return domain.User{}, err
	// }

	// return body, nil
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

// get all users
func (u *userUseCase) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	ctxDeadline, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	body, err := u.repo.GetAllUsers(ctxDeadline)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// search by email , username
func (u *userUseCase) SearchTheUser(ctx context.Context, s request.Search) ([]domain.User, error) {
	ctxDeadlin, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	switch s.SearchBy {
	case "email":
		body, err := u.repo.SearchUserByEmail(ctxDeadlin, s.Keyword)
		if err != nil {
			return nil, err
		}
		if body == nil {
			s.SearchBy = "username"
			body, err = u.SearchTheUser(ctxDeadlin, s)
			if err != nil {
				return nil, err
			}
			return body, nil
		}
		return body, err
	case "username":
		body, err := u.repo.SearchUserByUsername(ctxDeadlin, s.Keyword)
		if err != nil {
			return nil, err
		}
		return body, err
	default:
		err := fmt.Errorf("unsupported search type: %s", s.Keyword)
		return nil, err
	}

}

// forget password
func (a *userUseCase) ForgotPassword(ctx context.Context, req request.ForgotPassword) (out domain.User, err error) {
	fmt.Println("in usecase block user")
	ctxDeadline, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	// Transactions(func(UserRepository) error) error

	err = a.repo.Transactions(func(repo repo.UserRepository) error {
		_, err = repo.GetUserById(ctxDeadline, req.Id)
		if err != nil {
			return err
		}
		err = repo.ForgetPassword(ctxDeadline, req)
		if err != nil {
			return err
		}
		out, err = repo.GetUserById(ctxDeadline, req.Id)
		if err != nil {
			return err
		}
		return nil

	})
	if err != nil {
		return domain.User{}, err
	}
	return out, err
}

func EmailVerify(ctx context.Context, body domain.User) (err error) {
	verifier := emailverifier.NewVerifier()
	ret, err := verifier.Verify(body.Email)
	if err != nil {
		err = errors.New("verify email address failed, error is: " + err.Error())
		// ch <- err
		return
	}
	if !ret.Syntax.Valid {
		err = errors.New("email address syntax is invalid")
		// ch <- err
		return
	}
	// _, err = verifier.CheckSMTP("", "")
	// if err != nil {
	// 	err = errors.New("check smtp failed: " + err.Error())
	// 	// ch <- err
	// 	return
	// }
	// smtp.CatchAll
	// ch <- nil
	return nil

}

// verify email by generating otp
// func (u *userUseCase) EmailVerify(ctx context.Context, id int) (string, error) {
// 	err := u.repo.Transactions(func(ur repo.UserRepository) error {
// 		body, err := ur.GetUserById(ctx, id)
// 		if err != nil {
// 			return err
// 		}

// 	})
// }
