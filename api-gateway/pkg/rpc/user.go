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

// get all problems
func (u *userServiceImpl) ViewAllProblems(in request.AllProbles) ([]response.Problem, error) {
	client := u.problemPool.Get().(*rpc.Client)

	defer u.problemPool.Put(client)

	out := new([]response.Problem)
	err := client.Call("ProblemUserClient.ViewAllProblems", in, out)
	if err != nil {
		fmt.Println("err in the end", err)
		// app.ErrorJson(c, err)
		return nil, err
	}

	fmt.Println("out", out)
	return *out, nil
}

// get problem by id
func (u *userServiceImpl) GetProblemById(ctx echo.Context, in request.GetOneProblemById) (response.Problem, error) {
	fmt.Println("in rpc ")
	client := u.problemPool.Get().(*rpc.Client)

	defer u.problemPool.Put(client)

	out := new(response.Problem)
	err := client.Call("ProblemUserClient.GetProblemById", in, out)
	if err != nil {
		fmt.Println("error at rpc connection :", err)
		return response.Problem{}, err
	}
	fmt.Println("out", out)
	return *out, nil
}

// execute the problem
func (u *userServiceImpl) ExecuteGoCodyById(ctx echo.Context, in request.SubmitCodeIdRequest) (code []byte, err error) {
	fmt.Println("api gateway rpc")

	client := u.problemPool.Get().(*rpc.Client)

	defer u.problemPool.Put(client)

	out := new([]byte)
	err = client.Call("ProblemUserClient.SubmitCodeById", in, out)
	if err != nil {
		fmt.Println("error at rpc connection :", err)
		return
	}
	fmt.Println("out", out)
	return *out, nil
}

func (u *userServiceImpl) SortProblemBy(ctx echo.Context, in request.SearchBy) ([]response.Problem, error) {
	fmt.Println("api gateway rpc")
	fmt.Println("get problem ", in)
	client := u.problemPool.Get().(*rpc.Client)
	defer u.problemPool.Put(client)

	out := new([]response.Problem)

	err := client.Call("ProblemUserClient.GetProblemBy", in, out)
	if err != nil {
		return nil, err
	}

	fmt.Println("out ", out)
	return *out, nil
}
