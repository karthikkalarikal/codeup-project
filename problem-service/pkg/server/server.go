package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"problem-service/pkg/client"
	"problem-service/pkg/config"
	"time"
)

type RpcServer struct {
	lis   net.Listener
	port  string
	user  *client.ProblemUserClient
	admin *client.AdminClientImpl
}

func NewRPCServer(cfg *config.Config, user *client.ProblemUserClient, admin *client.AdminClientImpl) *RpcServer {
	// log.Println("starting rpc server on port ", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", cfg.ProblemServicePort))
	if err != nil {
		log.Panic(err)
	}

	err = rpc.Register(user) // need to register the functions
	if err != nil {
		log.Println(err.Error())
	}
	err = rpc.Register(admin)
	if err != nil {
		log.Println(err.Error())
	}

	return &RpcServer{
		lis:   listen,
		port:  cfg.ProblemServicePort,
		user:  user,
		admin: admin,
	}
}

func (l *RpcServer) Start() error {
	log.Println("problem service listening on port ", l.port)
	log.Println("method called: ")
	for {
		conn, err := l.lis.Accept()
		if err != nil {
			retries, maxRetries := 0, 3
			for retries < maxRetries {
				log.Println("error in accepting connection:", err)
				time.Sleep(1 * time.Second)
				conn, err = l.lis.Accept()
				if err == nil {
					break // success
				}

				retries++
			}
			if retries >= maxRetries {
				return fmt.Errorf("max retries exceeded")
			}

		}
		go rpc.ServeConn(conn)
	}
}
