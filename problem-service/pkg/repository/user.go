package repository

import (
	"context"
	"fmt"
	"log"
	"problem-service/pkg/domain"
	"problem-service/pkg/repository/interfaces"
	"problem-service/pkg/utils/request"
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

// // insert
// func (p *problemDatabase) InsertProblem(ctx context.Context, entry request.Problem) (int, error) {
// 	collection := p.DB.Database("problems").Collection("problems")
// 	var problem domain.Problem
// 	copier.Copy(problem, entry)
// 	problem.CreatedAt = time.Now()

// 	body, err := collection.InsertOne(context.TODO(), problem)

// 	if err != nil {
// 		log.Println("error inerting into problems: ", err)
// 		return body.InsertedID.(int), err
// 	}

// 	return body.InsertedID.(int), err
// }

// view all
func (p *problemDatabase) ViewAllProblems(ctx context.Context) ([]domain.Problem, error) {

	ctxTO, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	collection := p.DB.Database("problems").Collection("problems")

	opts := options.Find()

	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})

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
			log.Println("appended item: ", item.Title)
		}
		if err := cursor.Err(); err != nil {
			log.Println("cursor error: ", err)
			return nil, err
		}

	}
	log.Println("total problems found :", len(problems))
	return problems, nil
}

// execute problem
func (p *problemDatabase) GetProblemById(ctx context.Context, id request.ProblemById) (domain.Problem, error) {
	collection := p.DB.Database("problems").Collection("problems")

	var entry domain.Problem
	fmt.Println("id", id)
	err := collection.FindOne(ctx, bson.M{"_id": id.ID}).Decode(&entry)

	if err != nil {
		return domain.Problem{}, err
	}

	return entry, nil

}

func (p *problemDatabase) GetProblemByDifficulty(ctx context.Context, difficulty string) ([]domain.Problem, error) {
	fmt.Println("difficulty ", difficulty)
	collection := p.DB.Database("problems").Collection("problems")
	// var entry []domain.Problem

	filter := bson.D{{Key: "difficulty", Value: difficulty}}

	opts := options.Find()

	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
	// opts.
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Println("finding all docs error:", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	var problems []domain.Problem

	for cursor.Next(ctx) {
		var item domain.Problem

		err := cursor.Decode(&item)
		if err != nil {
			log.Println("error decoding problems into slices", err)
			return nil, err
		} else {
			problems = append(problems, item)
			log.Println("appended item: ", item.Title)
		}
		if err := cursor.Err(); err != nil {
			log.Println("cursor error: ", err)
			return nil, err
		}

	}
	log.Println("total problems found :", len(problems))
	return problems, nil
}

func (p *problemDatabase) GetProblemByTags(ctx context.Context, tag string) ([]domain.Problem, error) {
	fmt.Println("tag ", tag)
	collection := p.DB.Database("problems").Collection("problems")
	// var entry []domain.Problem

	filter := bson.D{{Key: "tags", Value: tag}}

	opts := options.Find()

	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
	// opts.
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Println("finding all docs error:", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	var problems []domain.Problem

	for cursor.Next(ctx) {
		var item domain.Problem

		err := cursor.Decode(&item)
		if err != nil {
			log.Println("error decoding problems into slices", err)
			return nil, err
		} else {
			problems = append(problems, item)
			log.Println("appended item: ", item.Title)
		}
		if err := cursor.Err(); err != nil {
			log.Println("cursor error: ", err)
			return nil, err
		}

	}
	log.Println("total problems found :", len(problems))
	return problems, nil
}
