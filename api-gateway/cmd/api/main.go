package main

import (
	"fmt"
	"log"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("start service on port %s", webPort)

	err := app.Routes().Start(fmt.Sprintf(":%s", webPort))
	if err != nil {
		fmt.Println("here", err.Error())
		log.Panic(err)
	}
}

// func hello(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World")
// }
