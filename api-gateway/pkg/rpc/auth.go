package rpc

import (
	"fmt"
	"net/rpc"
	"sync"

	"github.com/karthikkalarikal/api-gateway/pkg/config"
	"github.com/karthikkalarikal/api-gateway/pkg/rpc/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/response"
	"github.com/labstack/echo/v4"
)

type authServiceImpl struct {
	cfg      config.Config
	authPool *sync.Pool // a different way to initialize rpc connections
}

func NewAuthService(cfg *config.Config) interfaces.AuthService {
	return &authServiceImpl{
		cfg: *cfg,
		authPool: &sync.Pool{
			New: func() interface{} {
				client, err := rpc.Dial("tcp", cfg.AuthServiceUrl)
				if err != nil {
					panic(err)
				}
				return client
			},
		},
	}
}

// sign up
func (c *authServiceImpl) UserSignUp(ctx echo.Context, in request.UserSignUpRequest) (*response.UserSignUpResponse, error) {
	// fmt.Println("tcp", c.cfg.AuthServiceUrl)
	// client, err := rpc.Dial("tcp", c.cfg.AuthServiceUrl)
	// fmt.Println("err", err)
	// if err != nil {

	// 	return nil, err
	// }
	client := c.authPool.Get().(*rpc.Client)
	defer c.authPool.Put(client)

	out := new(response.UserSignUpResponse)
	err := client.Call("AuthUserService.SignUp", in, out)
	if err != nil {
		fmt.Println("err in the end", err)
		// app.ErrorJson(c, err)
		return nil, err
	}

	fmt.Println("out", out)
	return out, nil
}

// sign in
func (c *authServiceImpl) UserSignIn(ctx echo.Context, in request.UserSignInRequest) (*response.UserSignInResponse, error) {

	client := c.authPool.Get().(*rpc.Client)
	defer c.authPool.Put(client)

	out := new(response.UserSignInResponse)
	err := client.Call("AuthUserService.UserSignIn", in, out)
	if err != nil {
		fmt.Println("err ", err)
		return nil, err
	}
	fmt.Println("out", out)
	return out, nil
}

// view all users
func (a *authServiceImpl) ViewUsers(e echo.Context) ([]response.User, error) {
	client := a.authPool.Get().(*rpc.Client)
	defer a.authPool.Put(client)

	out := new([]response.User)

	err := client.Call("AuthUserService.GetAllUsers", struct{}{}, out)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	fmt.Println("out ", out)
	return *out, nil
}

// search users by email, username
func (a *authServiceImpl) SearchUser(e echo.Context, req request.Search) ([]response.User, error) {
	fmt.Println("req usecase ", req)
	client := a.authPool.Get().(*rpc.Client)
	defer a.authPool.Put(client)

	out := new([]response.User)

	err := client.Call("AuthUserService.SearchUsers", req, out)
	if err != nil {
		fmt.Println("err ", err)
		return nil, err
	}

	fmt.Println("out ", out)
	return *out, nil
}

// block user
func (a *authServiceImpl) BlockUser(e echo.Context, in int) (response.BlockedStatus, error) {
	fmt.Println("block user rpc", in)

	client := a.authPool.Get().(*rpc.Client)

	defer a.authPool.Put(client)

	out := new(response.User)
	err := client.Call("AuthUserService.BlockUser", in, out)
	if err != nil {
		fmt.Println("err ", err)
		// log.Panic(err)
		return response.BlockedStatus{}, err
	}
	fmt.Println("out ", out)
	return response.BlockedStatus{
		ID:      out.ID,
		Blocked: out.Blocked,
	}, nil
}

// forget password
func (a *authServiceImpl) ForgetPassword(e echo.Context, in request.ForgotPassword) (response.User, error) {
	fmt.Println("forget password")

	client := a.authPool.Get().(*rpc.Client)
	defer a.authPool.Put(client)

	out := new(response.User)
	err := client.Call("AuthUserService.ForgetPassword", in, out)
	if err != nil {
		fmt.Println("err ", err)
		return response.User{}, err
	}
	fmt.Println("out ", out)
	return *out, nil
}
