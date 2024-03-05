package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"problem-service/pkg/client"
	"problem-service/pkg/config"
)

type RpcServer struct {
	lis     net.Listener
	port    string
	problem client.ProblemUserClient
}

func NewRPCServer(cfg *config.Config, problem client.ProblemUserClient) *RpcServer {
	// log.Println("starting rpc server on port ", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", cfg.ProblemServicePort))
	if err != nil {
		log.Panic(err)
	}

	err = rpc.Register(problem) // need to register the functions
	if err != nil {
		log.Println(err.Error())
	}

	return &RpcServer{
		lis:     listen,
		port:    cfg.ProblemServicePort,
		problem: problem,
	}
}

func (l *RpcServer) Start() error {
	fmt.Println("problem service listening on port ", l.port)
	for {
		conn, err := l.lis.Accept()
		if err != nil {
			log.Println("error in accepting connection:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
