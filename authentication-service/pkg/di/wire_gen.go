// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"authentication/pkg/api"
	"authentication/pkg/client"
	"authentication/pkg/config"
	"authentication/pkg/db"
	"authentication/pkg/repository"
	"authentication/pkg/usecase"
)

// Injectors from wire.go:

func InitializeServices(cfg *config.Config) (*api.RpcServer, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	authUserService := client.NewUserService(userUseCase)
	rpcServer := api.NewRPCServer(cfg, authUserService)
	return rpcServer, nil
}
