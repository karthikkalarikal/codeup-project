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

// InsertProblem

type adminServiceImpl struct {
	cfg         config.Config
	problemPool *sync.Pool // a different way to initialize rpc connections
}

func NewAdminService(cfg *config.Config) interfaces.AdminRPCService {
	return &adminServiceImpl{
		cfg: *cfg,
		problemPool: &sync.Pool{
			New: func() interface{} {
				client, err := rpc.Dial("tcp", cfg.ProblemServiceUrl)
				if err != nil {
					panic(err)
				}
				return client
			},
		},
	}
}

func (u *adminServiceImpl) InsertProblem(ctx echo.Context, in request.InsertProblem) (response.Problem, error) {
	fmt.Println("here in rpc", in)
	client := u.problemPool.Get().(*rpc.Client)

	defer u.problemPool.Put(client)

	out := new(response.Problem)
	err := client.Call("AdminClientImpl.InsertProblem", in, out)
	if err != nil {
		fmt.Println("err in the end", err)
		// app.ErrorJson(c, err)
		return response.Problem{}, err
	}

	fmt.Println("out", out)
	return *out, nil
}
