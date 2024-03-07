package client

import (
	"context"
	"log"
	"problem-service/pkg/domain"
	"problem-service/pkg/usecase/interfaces"
	"problem-service/pkg/utils/request"
)

type AdminClientImpl struct {
	admin interfaces.AdminUseCase
}

func NewAdminClient(admin interfaces.AdminUseCase) *AdminClientImpl {
	return &AdminClientImpl{
		admin: admin,
	}
}

func (a *AdminClientImpl) InsertProblem(req request.Problem, reply *domain.Problem) error {
	log.Println("Admin wants to insert a problems", req)
	ctx := context.Background()
	body, err := a.admin.InsertProblem(ctx, req)
	if err != nil {
		return err
	}
	*reply = body
	return nil
}
