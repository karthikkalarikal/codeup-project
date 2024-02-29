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

func (c *authServiceImpl) UserSignUp(ctx echo.Context, in request.UserSignUp) (*response.UserSignUp, error) {

	client, err := rpc.Dial("tcp", c.cfg.AuthServiceUrl)
	if err != nil {

		return nil, err
	}

	out := new(response.UserSignUp)
	err = client.Call("User.SignUp", in, out)
	if err != nil {
		fmt.Println("err in the end", err)
		// app.ErrorJson(c, err)
		return nil, err
	}

	fmt.Println("out", out)
	return out, nil
}
