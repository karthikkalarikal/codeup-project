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
)

func InitializeAPI(cfg config.Config) (*api.Server, error) {
	wire.Build(
		client.NewAuthClient,

		handlers.NewAuthHandler,

		rpc.NewAuthService,
		api.NewServerHTTP,
	)
	return &api.Server{}, nil
}
