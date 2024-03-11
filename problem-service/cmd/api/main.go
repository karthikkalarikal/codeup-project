package main

import (
	"log"
	"problem-service/pkg/di"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	// cfg, err := config.LoadConfig()

	// if err != nil {
	// 	log.Fatalf("failed to load config: %s", err)
	// }

	service, err := di.InitializeServices()
	if err != nil {
		log.Fatalf("failed to initialize service : %s", err)
	}

	service.Start()
}
