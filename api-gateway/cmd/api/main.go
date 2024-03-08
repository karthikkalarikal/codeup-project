package main

import (
	"fmt"
	"log"

	_ "github.com/karthikkalarikal/api-gateway/cmd/api/docs" // docs
	"github.com/karthikkalarikal/api-gateway/pkg/config"
	"github.com/karthikkalarikal/api-gateway/pkg/di"
)

// const webPort = "80"

//	@title			Code Up Project API Documentation
//	@version		1.0
//	@description	This is a sample code execution platform.
//	@termsOfService	http://swagger.io/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
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

}
