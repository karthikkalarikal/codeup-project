package client

import (
	"authentication/pkg/domain"
	user "authentication/pkg/usecase/interfaces"
	"authentication/pkg/utils/request"
	"authentication/pkg/utils/response"
	"context"
	"fmt"
	"strings"
)

type AuthUserService struct {
	useCase user.UserUseCase
}

func NewUserService(user user.UserUseCase) *AuthUserService {
	return &AuthUserService{
		useCase: user,
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
// angia tech \\ //