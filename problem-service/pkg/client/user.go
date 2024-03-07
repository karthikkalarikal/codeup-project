package client

import (
	"context"
	"log"
	"problem-service/pkg/domain"
	usecase "problem-service/pkg/usecase/interfaces"
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
