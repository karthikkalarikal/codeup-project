package rpc

import (
	"fmt"
	"net/rpc"
	"sync"

	"github.com/karthikkalarikal/api-gateway/pkg/config"
	"github.com/karthikkalarikal/api-gateway/pkg/rpc/interfaces"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/request"
	"github.com/karthikkalarikal/api-gateway/pkg/utils/response"
)

type userServiceImpl struct {
	cfg         config.Config
	problemPool *sync.Pool // a different way to initialize rpc connections
}

func NewUserService(cfg *config.Config) interfaces.UserRPCService {
	return &userServiceImpl{
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

func (u *userServiceImpl) ViewAllProblems(in request.AllProbles) ([]response.Problem, error) {
	client := u.problemPool.Get().(*rpc.Client)

	defer u.problemPool.Put(client)

	out := new([]response.Problem)
	err := client.Call("RPCServer.All", in, out)
	if err != nil {
		fmt.Println("err in the end", err)
		// app.ErrorJson(c, err)
		return nil, err
	}

	fmt.Println("out", out)
	return *out, nil
}