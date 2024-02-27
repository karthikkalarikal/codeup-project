package main

import (
	"context"
	"fmt"
	"log"
	"problem-service/data"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RPCServer struct {
}

// type TestCase struct {
// 	Input  string `bson:"input" json:"input"`
// 	Output string `bson:"output" json:"output"`
// }

type RPCProblme struct {
	// ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// Title       string             `bson:"title" json:"title"`
	// Description string             `bson:"description" json:"description"`
	// Difficulty  string             `bson:"difficulty" json:"difficulty"`
	// TestCases   []TestCase         `bson:"test_cases" json:"test_cases"`
	// TimeLimit   int                `bson:"time_limit" json:"time_limit"`
	// MemoryLimit int                `bson:"memory_limit" json:"memory_limit"`
	// Tags        []string           `bson:"tags" json:"tags"`
	// CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

func (r *RPCServer) All(payload RPCProblme, resp *[]data.Problem) error {
	fmt.Println("here in handler")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("problems").Collection("problems")

	opts := options.Find()

	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("finding all docs error:", err)
		return err
	}

	defer cursor.Close(ctx)

	var problems []data.Problem

	for cursor.Next(ctx) {
		var item data.Problem

		err := cursor.Decode(&item)
		if err != nil {
			log.Println("error decoding problems into slices", err)
			return err
		} else {
			problems = append(problems, item)
		}

	}
	fmt.Println("problems", problems)
	*resp = problems
	return nil
}
