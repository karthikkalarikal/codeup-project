//go:build wireinject
// +build wireinject

package di

import (
	"problem-service/pkg/client"
	"problem-service/pkg/config"
	"problem-service/pkg/db"
	"problem-service/pkg/repository"
	"problem-service/pkg/rpc"
	"problem-service/pkg/server"
	"problem-service/pkg/usecase"

	"github.com/google/wire"
)

func InitializeServices() (*server.RpcServer, error) {
	wire.Build(
		provideConfig,
		db.ConnectToMongo,

		client.NewUserClient,
		client.NewAdminClient,

		repository.NewAdmimRepository,
		repository.NewUserRepository,

		usecase.NewAdminUseCase,
		usecase.NewUserUseCase,
		rpc.NewUserProblemRPC,
		// config.NewConfig,

		server.NewRPCServer,
	)
	return &server.RpcServer{}, nil
}
func provideConfig() (*config.Config, error) {

	return config.LoadConfig()
}
