package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sandbox/pkg/config"
	"sandbox/pkg/executer"
	"time"
)

func main() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":80", nil)
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Panic(err)
	}
	exec := executer.Executer{}
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", cfg.GoSandboxServer))
	err = rpc.Register(exec)
	if err != nil {
		log.Println(err.Error())
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			retries, maxRetries := 0, 3
			for retries < maxRetries {
				log.Println("error in accepting connection:", err)
				time.Sleep(1 * time.Second)
				conn, err = listen.Accept()
				if err == nil {
					break // success
				}

				retries++
			}
			if retries >= maxRetries {
				log.Println(fmt.Errorf("max retries exceeded"))
			}

		}
		go rpc.ServeConn(conn)
	}
}
