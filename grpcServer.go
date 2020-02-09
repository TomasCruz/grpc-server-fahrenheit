package main

import (
	"fmt"
	"log"
	"net"

	"github.com/TomasCruz/grpc-server-fahrenheit/presenter"
	"google.golang.org/grpc"
)

func registerGRPCServer(port string) (grpcServer *grpc.Server, listener net.Listener, err error) {
	address := fmt.Sprintf("localhost:%s", port)

	// create a listener on TCP port
	if listener, err = net.Listen("tcp", address); err != nil {
		return
	}

	// create and register server instance
	grpcServer = grpc.NewServer()
	presenter.RegisterConvertorServer(grpcServer, &presenter.Server{})
	return
}

func startGRPCServer(grpcServer *grpc.Server, listener net.Listener, port string) (err error) {
	// start the server
	log.Printf("starting gRPC server on localhost:%s", port)
	err = grpcServer.Serve(listener)
	return
}
