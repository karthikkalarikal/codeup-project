package client

import (
	"context"
	"fmt"
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

// first half
func (a *AdminClientImpl) InsertFirstHalfProblem(req request.FirstHalfCode, reply *domain.Problem) error {
	log.Println("Admin wants to insert first half of code")
	fmt.Println("code ", string(req.FirstHalfCode))

	ctx := context.Background()
	body, err := a.admin.InsertFirstHalfProblem(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println("body: ", string(reply.FirstHalfCode))
	*reply = body
	return nil
}

// second half
func (a *AdminClientImpl) InsertSecondHalfProblem(req request.SecondHalfCode, reply *domain.Problem) error {
	log.Println("Admin want's to insert second half of code")

	ctx := context.Background()

	// body, err := a.admin.InsertSecondHalfProblem(ctx, req)

	body, err := a.admin.InsertSecondHalfProblem(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println("body: ", string(reply.FirstHalfCode))
	*reply = body
	return nil
}
