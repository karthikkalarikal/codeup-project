//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/karthikkalarikal/api-gateway/pkg/api"
	"github.com/karthikkalarikal/api-gateway/pkg/api/handlers"
	"github.com/karthikkalarikal/api-gateway/pkg/client"
	"github.com/karthikkalarikal/api-gateway/pkg/config"
	"github.com/karthikkalarikal/api-gateway/pkg/rpc"
	"github.com/karthikkalarikal/api-gateway/pkg/utils"
)

func InitializeAPI(cfg config.Config) (*api.Server, error) {
	wire.Build(
		client.NewAuthClient,
		client.NewUserClient,
		client.NewAdminClient,

		handlers.NewAuthHandler,
		handlers.NewUserHandler,
		handlers.NewAdminHandler,

		rpc.NewAuthService,
		rpc.NewUserService,
		rpc.NewAdminService,

		utils.NewUtils,

		config.NewConfig,

		api.NewServerHTTP,
	)
	return &api.Server{}, nil
}
