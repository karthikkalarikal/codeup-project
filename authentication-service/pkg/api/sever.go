package api

import (
	"authentication/pkg/client"
	"authentication/pkg/config"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type RpcServer struct {
	lis    net.Listener
	port   string
	auth   *client.AuthUserService
	stripe *config.StripeConfig
}

func NewRPCServer(stripe *config.StripeConfig, cfg *config.Config, auth *client.AuthUserService) *RpcServer {
	// log.Println("starting rpc server on port ", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", cfg.ServicePort))
	if err != nil {
		log.Panic(err)
	}

	err = rpc.Register(auth)
	if err != nil {
		log.Println(err.Error())
	}

	return &RpcServer{
		lis:    listen,
		port:   cfg.ServicePort,
		auth:   auth,
		stripe: stripe,
	}
}

func (l *RpcServer) Start() error {
	fmt.Println("auth service listening on port ", l.port)
	for {
		conn, err := l.lis.Accept()
		if err != nil {
			log.Println("error in accepting connection:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
