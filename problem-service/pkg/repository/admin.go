package repository

import (
	"context"
	"fmt"
	"log"
	"problem-service/pkg/domain"
	"problem-service/pkg/repository/interfaces"
	"problem-service/pkg/utils/request"
	"time"

	"github.com/jinzhu/copier"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type problemDatabaseAdmin struct {
	DB *mongo.Client
}

func NewAdmimRepository(DB *mongo.Client) interfaces.AdminRepository {
	return &problemDatabaseAdmin{
		DB: DB,
	}
}

// insert
func (p *problemDatabaseAdmin) InsertProblem(ctx context.Context, entry request.Problem) (primitive.ObjectID, error) {
	fmt.Println("here inside repository", entry)
	collection := p.DB.Database("problems").Collection("problems")
	var problem domain.Problem
	err := copier.CopyWithOption(&problem, &entry, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return primitive.ObjectID{}, err
	}
	problem.CreatedAt = time.Now()
	fmt.Println("problem", problem)
	body, err := collection.InsertOne(context.TODO(), problem)
	fmt.Println("body", body)
	if err != nil {
		log.Println("error inerting into problems: ", err)
		return primitive.ObjectID{}, err
	}
	fmt.Println("id", body.InsertedID)

	return body.InsertedID.(primitive.ObjectID), err
}

// get one problem by id
func (p *problemDatabaseAdmin) GetProblemById(ctx context.Context, id primitive.ObjectID) (domain.Problem, error) {

	collection := p.DB.Database("problems").Collection("problems")

	var entry domain.Problem

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&entry)

	if err != nil {
		return domain.Problem{}, err
	}

	return entry, nil
}

// drop the collection
// func (l *Models) DropCollection() error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()

// 	collection := l.Client.Database("problems").Collection("problems")

// 	if err := collection.Drop(ctx); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (l *Models) GetOne(id string) (*Problem, error) {
//
// }
