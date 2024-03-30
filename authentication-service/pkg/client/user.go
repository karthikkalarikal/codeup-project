package client

import (
	"authentication/pkg/domain"
	"time"

	// "authentication/pkg/usecase/interfaces"
	user "authentication/pkg/usecase/interfaces"
	"authentication/pkg/utils/request"
	"authentication/pkg/utils/response"
	"context"
	"fmt"
	"strings"
)

type AuthUserService struct {
	useCase user.UserUseCase
	admin   user.AdminUsecase
	payment user.PaymentUsecase
}

func NewUserService(user user.UserUseCase, admin user.AdminUsecase, payment user.PaymentUsecase) *AuthUserService {
	return &AuthUserService{
		useCase: user,
		admin:   admin,
		payment: payment,
	}
}

// sign up
func (u *AuthUserService) SignUp(req request.UserSignUpRequest, reply *domain.User) error {
	ctx := context.Background()
	body, err := u.useCase.UserSignUp(ctx, req)

	if err != nil {
		return err
	}

	*reply = body
	return nil
}

// user sign in
func (u *AuthUserService) UserSignIn(req request.UserSignInRequest, reply *response.UserSignInResponse) (err error) {
	ctx := context.Background()
	body, err := u.useCase.UserSignIn(ctx, req)

	if err != nil {
		return
	}

	*reply = body
	return nil
}

// get all users
func (u *AuthUserService) GetAllUsers(req struct{}, reply *[]domain.User) error {
	ctx := context.Background()
	body, err := u.useCase.GetAllUsers(ctx)
	if err != nil {
		return err
	}
	*reply = body
	return nil
}

// serach users by email and username
func (u *AuthUserService) SearchUsers(req request.Search, reply *[]domain.User) error {
	fmt.Println("keyword ", req.Keyword)
	key := strings.TrimSpace(req.Keyword)
	fmt.Println("keyword length", len(req.Keyword))
	ctx := context.Background()
	if key == "" {
		body, err := u.useCase.GetAllUsers(ctx)
		if err != nil {
			return err
		}
		fmt.Println("here :")
		*reply = body
		return nil
	} else {
		body, err := u.useCase.SearchTheUser(ctx, req)
		if err != nil {
			return err
		}
		*reply = body
		return nil
	}

}

// block user
func (u *AuthUserService) BlockUser(req int, reply *domain.User) error {
	fmt.Println("here in block user auth service ", req)
	ctx := context.Background()

	out, err := u.admin.BlockUser(ctx, req)
	fmt.Println("err ", err)
	if err != nil {
		return err
	}
	fmt.Println("out ", out)
	*reply = out
	return nil
}

// forget password
func (u *AuthUserService) ForgetPassword(req request.ForgotPassword, reply *domain.User) error {
	fmt.Println("here in forget password", req)
	ctx := context.Background()

	out, err := u.useCase.ForgotPassword(ctx, req)
	if err != nil {
		return err
	}
	*reply = out
	return nil
}

// charge
func (u *AuthUserService) Charge(req request.Payment, reply *[]byte) error {
	fmt.Println("here in charge")
	ctx := context.Background()
	out, err := u.payment.GetPaymentIntent(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println(out)

	*reply = out
	return nil
}

// make prime
func (u *AuthUserService) MakePrime(req string, reply *string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	out, err := u.useCase.MakePrime(ctx, req)
	if err != nil {
		return err
	}
	*reply = out
	return nil
}

// unsubscribe
func (u *AuthUserService) UnSubscribe(req int, reply *domain.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	out, err := u.useCase.UnSubscribe(ctx, req)
	if err != nil {
		return err
	}
	*reply = out
	return nil
}
