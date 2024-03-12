package interfaces

import (
	"context"
	"problem-service/pkg/domain"
	"problem-service/pkg/utils/request"
	req "problem-service/pkg/utils/request"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminRepository interface {
	InsertProblem(ctx context.Context, entry req.Problem) (primitive.ObjectID, error)
	GetProblemById(ctx context.Context, entry primitive.ObjectID) (domain.Problem, error)
	InsertFirstHalfProblem(ctx context.Context, entry request.FirstHalfCode) (domain.Problem, error)
	InsertSecondHalfProblem(ctx context.Context, entry request.SecondHalfCode) (domain.Problem, error)
}
