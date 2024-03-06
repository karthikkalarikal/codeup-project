package client

import (
	"context"
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

func (p *ProblemUserClient) ViewAllProblems(request struct{}, reply *[]domain.Problem) error {
	ctx := context.Background()
	body, err := p.user.ViewAllProblems(ctx)
	if err != nil {
		return err
	}
	*reply = body
	// copy(*reply, body)
	return nil
}
