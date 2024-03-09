package rpc

import (
	"fmt"
	"net/rpc"
	"sync"

	"github.com/karthikkalarikal/api-gateway/pkg/config"
	"github.com/karthikkalarikal/api-gateway/pkg/rpc/interfaces"
	"github.com/labstack/echo/v4"
)

type goCodeExecRPC struct {
	cfg        config.Config
	goexecPool *sync.Pool
}

func NewGoExexRPC(cfg *config.Config) interfaces.GoCodeExecRPC {
	return &goCodeExecRPC{
		cfg: *cfg,
		goexecPool: &sync.Pool{
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

func (u *goCodeExecRPC) WriteGoCode(e echo.Context, in []byte) ([]byte, error) {
	client := u.goexecPool.Get().(*rpc.Client)
	defer u.goexecPool.Put(client)

	out := new([]byte)
	err := client.Call("Executer.GoCodeExec", in, out)
	if err != nil {
		fmt.Println("err in the end", err)
		return []byte{}, err
	}
	fmt.Println("out", out)
	return *out, nil
}
