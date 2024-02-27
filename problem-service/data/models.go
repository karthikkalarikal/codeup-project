package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Models struct {
	Problem Problem
	Client  *mongo.Client
}

func New(mongo *mongo.Client) Models {
	// client = mongo
	if mongo == nil {
		log.Panic("Cannot create data models with a nil MongoDB client")
	}
	return Models{
		Problem: Problem{},
		Client:  mongo,
	}
}

type TestCase struct {
	Input  string `bson:"input" json:"input"`
	Output string `bson:"output" json:"output"`
}

type Problem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Difficulty  string             `bson:"difficulty" json:"difficulty"`
	TestCases   []TestCase         `bson:"test_cases" json:"test_cases"`
	TimeLimit   int                `bson:"time_limit" json:"time_limit"`
	MemoryLimit int                `bson:"memory_limit" json:"memory_limit"`
	Tags        []string           `bson:"tags" json:"tags"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

// insert
func (l *Models) Insert(entry Problem) error {
	collection := l.Client.Database("problems").Collection("problems")

	if entry.CreatedAt.IsZero() {
		entry.CreatedAt = time.Now()
	}
	_, err := collection.InsertOne(context.TODO(), entry)
	if err != nil {
		log.Println("error inerting into problems: ", err)
		return err
	}

	return nil
}

// view all
func (l *Models) All() ([]*Problem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := l.Client.Database("problems").Collection("problems")

	opts := options.Find()

	opts.SetSort(bson.D{{"created_at", -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("finding all docs error:", err)
		return nil, err
	}

	defer cursor.Close(ctx)

	var problems []*Problem

	for cursor.Next(ctx) {
		var item Problem

		err := cursor.Decode(&item)
		if err != nil {
			log.Println("error decoding problems into slices", err)
			return nil, err
		} else {
			problems = append(problems, &item)
		}

	}
	return problems, nil
}

// view one by id
func (l *Models) GetOne(id string) (*Problem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := l.Client.Database("problems").Collection("problems")

	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var entry Problem

	err = collection.FindOne(ctx, bson.M{"_id": docId}).Decode(&entry)

	if err != nil {
		return nil, err
	}

	return &entry, nil
}

// drop the collection
func (l *Models) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := l.Client.Database("problems").Collection("problems")

	if err := collection.Drop(ctx); err != nil {
		return err
	}
	return nil
}

// func (l *Problem) Update() (*mongo.UpdateResult, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()

// 	collection := client.Database("problems").Collection("problems")

// 	docId, err := primitive.ObjectIDFromHex(l.ID.Hex())
// 	if err != nil {
// 		return nil, err
// 	}

// 	result, err := collection.UpdateOne(
// 		ctx,
// 		bson.M{"_id": docId},
// 		bson.D{
// 			{"$set", bson.D{
// 				{"title": l.Title},
// 			}},
// 		},
// 	)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }
