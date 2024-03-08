package client

import (
	"context"
	"fmt"
	"log"
	"problem-service/pkg/domain"
	usecase "problem-service/pkg/usecase/interfaces"
	"problem-service/pkg/utils/request"
)

type ProblemUserClient struct {
	user usecase.UserUseCase
}

func NewUserClient(user usecase.UserUseCase) *ProblemUserClient {
	return &ProblemUserClient{
		user: user,
	}
}

// view
func (p *ProblemUserClient) ViewAllProblems(request struct{}, reply *[]domain.Problem) error {
	log.Println("User wants to view all problems")
	ctx := context.Background()
	body, err := p.user.ViewAllProblems(ctx)
	if err != nil {
		return err
	}
	*reply = body
	// copy(*reply, body)
	return nil
}

// get one problem

func (p *ProblemUserClient) GetProblemById(request request.ProblemById, reply *domain.Problem) error {
	fmt.Println("get one problem")
	ctx := context.Background()
	body, err := p.user.GetProblemById(ctx, request)
	if err != nil {
		return err
	}

	*reply = body
	return nil
}
