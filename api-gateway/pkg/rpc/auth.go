package rpc

import (
	"fmt"
	"net/rpc"

	"github.com/karthikkalarikal/api-gateway/pkg/config"
	"github.com/karthikkalarikal/api-gateway/pkg/rpc/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/response"
	"github.com/labstack/echo/v4"
)

type authServiceImpl struct {
	cfg config.Config
}

func NewAuthService(cfg *config.Config) interfaces.AuthService {
	return &authServiceImpl{cfg: *cfg}
}

func (c *authServiceImpl) UserSignUp(ctx echo.Context, in request.UserSignUpRequest) (*response.UserSignUpResponse, error) {
	fmt.Println("tcp", c.cfg.AuthServiceUrl)
	client, err := rpc.Dial("tcp", c.cfg.AuthServiceUrl)
	fmt.Println("err", err)
	if err != nil {

		return nil, err
	}

	out := new(response.UserSignUpResponse)
	err = client.Call("AuthUserService.SignUp", in, out)
	if err != nil {
		fmt.Println("err in the end", err)
		// app.ErrorJson(c, err)
		return nil, err
	}

	fmt.Println("out", out)
	return out, nil
}
