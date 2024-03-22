// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/karthikkalarikal/api-gateway/pkg/api"
	"github.com/karthikkalarikal/api-gateway/pkg/api/handlers"
	"github.com/karthikkalarikal/api-gateway/pkg/client"
	"github.com/karthikkalarikal/api-gateway/pkg/config"
	"github.com/karthikkalarikal/api-gateway/pkg/rpc"
	"github.com/karthikkalarikal/api-gateway/pkg/utils"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*api.Server, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	authService := rpc.NewAuthService(configConfig)
	authClient := client.NewAuthClient(authService)
	utilsUtils := utils.NewUtils()
	authHandler := handlers.NewAuthHandler(authClient, utilsUtils)
	userRPCService := rpc.NewUserService(configConfig)
	userClient := client.NewUserClient(userRPCService, authService)
	goCodeExecRPC := rpc.NewGoExexRPC(configConfig)
	goCodeExecClient := client.NewGoExecClient(goCodeExecRPC)
	userHandler := handlers.NewUserHandler(userClient, utilsUtils, goCodeExecClient)
	adminRPCService := rpc.NewAdminService(configConfig)
	adminClient := client.NewAdminClient(adminRPCService, authService)
	adminHandler := handlers.NewAdminHandler(adminClient, utilsUtils)
	server := api.NewServerHTTP(configConfig, authHandler, userHandler, adminHandler)
	return server, nil
}
