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

type RPCProblme struct{}

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
