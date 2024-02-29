package main

import (
	"fmt"
	"log"

	"github.com/karthikkalarikal/api-gateway/pkg/config"
	"github.com/karthikkalarikal/api-gateway/pkg/di"
)

// const webPort = "80"

func main() {

	cfg, err := config.LoadConfig()
	fmt.Println("here in main 1")
	if err != nil {
		log.Fatalf("failed to load config error: %s", err.Error())

	}

	service, err := di.InitializeAPI(cfg)
	fmt.Println("here in main 2")
	if err != nil {
		log.Fatalf("failed initialize api error: %s", err.Error())
	}
	fmt.Println("here 3")

	service.Start()
	fmt.Println("here 4")
	// app := Config{}

	// log.Printf("start service on port %s", webPort)

	// err := app.Routes().Start(fmt.Sprintf(":%s", webPort))
	// if err != nil {
	// 	fmt.Println("here", err.Error())
	// 	log.Panic(err)
	// }
}

// func hello(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World")
// }
