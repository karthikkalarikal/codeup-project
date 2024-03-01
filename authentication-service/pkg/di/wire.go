//go:build wireinject
// +build wireinject

package di

import (
	"authentication/pkg/api"
	"authentication/pkg/client"
	"authentication/pkg/config"
	"authentication/pkg/db"
	"authentication/pkg/repository"
	"authentication/pkg/usecase"

	"github.com/google/wire"
)

func InitializeServices(cfg *config.Config) (*api.RpcServer, error) {
	wire.Build(

		db.ConnectDatabase,
		client.NewUserService,
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		api.NewRPCServer,
	)

	return &api.RpcServer{}, nil
}
