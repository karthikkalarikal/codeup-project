package db

import (
	"context"
	"log"
	"problem-service/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo(cfg *config.Config) (*mongo.Client, error) {
	// create connection options

	clientOptions := options.Client().ApplyURI(cfg.MongoURL)

	clientOptions.SetAuth(options.Credential{
		Username:      cfg.Username,
		Password:      cfg.Password,
		AuthSource:    cfg.AuthSource,
		AuthMechanism: cfg.AuthMechanism,
	}) // remembet to take these  values from env when fine tuning the code.

	// connect

	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("error connection ", err)
		return nil, err
	}

	log.Println("connected to mongo")

	return c, nil
}
