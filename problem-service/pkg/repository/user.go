package repository

import (
	"context"
	"log"
	"problem-service/pkg/domain"
	"problem-service/pkg/repository/interfaces"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type problemDatabase struct {
	DB *mongo.Client
}

func NewUserRepository(DB *mongo.Client) interfaces.UserRepository {
	return &problemDatabase{
		DB: DB,
	}
}

func (p *problemDatabase) ViewAllProblems(ctx context.Context) ([]domain.Problem, error) {

	ctxTO, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	collection := p.DB.Database("problems").Collection("problems")

	opts := options.Find()

	opts.SetSort(bson.D{{"created_at", -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("finding all docs error:", err)
		return nil, err
	}

	defer cursor.Close(ctxTO)

	var problems []domain.Problem

	for cursor.Next(ctx) {
		var item domain.Problem

		err := cursor.Decode(&item)
		if err != nil {
			log.Println("error decoding problems into slices", err)
			return nil, err
		} else {
			problems = append(problems, item)
		}

	}
	return problems, nil
}
