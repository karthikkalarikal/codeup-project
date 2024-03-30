package usecase

import (
	"context"
	"errors"
	"fmt"
	"problem-service/pkg/domain"
	"problem-service/pkg/repository/interfaces"
	rpc "problem-service/pkg/rpc/interfaces"
	user "problem-service/pkg/usecase/interfaces"
	"problem-service/pkg/utils/request"
)

type userUseCase struct {
	repo interfaces.UserRepository
	rpc  rpc.UserRPCProblem
}

func NewUserUseCase(repo interfaces.UserRepository, rpc rpc.UserRPCProblem) user.UserUseCase {
	return &userUseCase{
		repo: repo,
		rpc:  rpc,
	}
}

//view all

func (u *userUseCase) ViewAllProblems(ctx context.Context) ([]domain.Problem, error) {
	body, err := u.repo.ViewAllProblems(ctx)
	if err != nil {
		return []domain.Problem{}, err
	}
	return body, nil
}

// get one

func (u *userUseCase) GetProblemById(ctx context.Context, id request.ProblemById) (domain.Problem, error) {
	body, err := u.repo.GetProblemById(ctx, id)
	if err != nil {
		return domain.Problem{}, err
	}
	return body, nil
}

func (u *userUseCase) SubmitCodeById(ctx context.Context, req request.SubmitCodeIdRequest) ([]byte, error) {

	id := &request.ProblemById{
		ID: req.ID,
	}
	body, err := u.repo.GetProblemById(ctx, *id)
	if err != nil {
		return nil, err
	}
	fmt.Println("body", body)

	finalCode := string(body.FirstHalfCode) + string(req.Code) + string(body.SecondHalfCode)
	code, err := u.rpc.ExecuteGoCode(ctx, []byte(finalCode))
	if err != nil {
		return nil, err
	}

	return code, nil
}

func (u *userUseCase) GetProblemBy(ctx context.Context, in request.SearchBy) ([]domain.Problem, error) {

	if in.Field == "tag" {
		body, err := u.repo.GetProblemByTags(ctx, in.Search)
		if err != nil {
			return nil, err
		}
		return body, nil
	} else if in.Field == "difficulty" {
		body, err := u.repo.GetProblemByDifficulty(ctx, in.Search)
		if err != nil {
			return nil, err
		}
		return body, nil
	} else {
		return nil, errors.New("invalid input")
	}
}
