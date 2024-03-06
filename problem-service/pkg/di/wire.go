//go:build wireinject
// +build wireinject

package di

import (
	"problem-service/pkg/client"
	"problem-service/pkg/config"
	"problem-service/pkg/db"
	"problem-service/pkg/repository"
	"problem-service/pkg/server"
	"problem-service/pkg/usecase"

	"github.com/google/wire"
)

func InitializeServices(cfg *config.Config) (*server.RpcServer, error) {
	wire.Build(
		db.ConnectToMongo,
		client.NewUserClient,
		repository.NewUserRepository,
		usecase.NewUserUseCase,

		server.NewRPCServer,
	)
	return &server.RpcServer{}, nil
}
