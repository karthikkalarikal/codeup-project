package rpc

import (
	"context"
	"fmt"
	"net/rpc"
	"problem-service/pkg/config"
	"problem-service/pkg/rpc/interfaces"
	"sync"
)

type userRPCProblemImpl struct {
	cfg      *config.Config
	execPool *sync.Pool
}

func NewUserProblemRPC(cfg *config.Config) interfaces.UserRPCProblem {
	return &userRPCProblemImpl{
		cfg: cfg,
		execPool: &sync.Pool{
			New: func() interface{} {
				client, err := rpc.Dial("tcp", cfg.GoSandboxUrl)
				if err != nil {
					panic(err)
				}
				return client
			},
		},
	}
}

func (u *userRPCProblemImpl) ExecuteGoCode(ctx context.Context, code []byte) (body []byte, err error) {
	fmt.Println("here in exec go code")
	client := u.execPool.Get().(*rpc.Client)
	defer u.execPool.Put(client)
	out := new([]byte)
	err = client.Call("Executer.GoCodeExec", code, out)
	if err != nil {
		fmt.Println("err in the end", err)
		// app.ErrorJson(c, err)
		return
	}

	fmt.Println("out", out)
	return *out, nil
}
