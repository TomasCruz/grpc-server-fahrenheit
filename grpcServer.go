package main

import (
	"fmt"
	"log"
	"net"

	"github.com/TomasCruz/grpc-server-fahrenheit/api"
	"github.com/TomasCruz/grpc-server-fahrenheit/presenter"
	"google.golang.org/grpc"
	ghealth "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func registerGRPCServer(port string) (grpcServer *grpc.Server, listener net.Listener, err error) {
	address := fmt.Sprintf("localhost:%s", port)

	// create a listener on TCP port
	if listener, err = net.Listen("tcp", address); err != nil {
		return
	}

	// create and register server instance
	grpcServer = grpc.NewServer(withServerUnaryInterceptor())
	api.RegisterConvertorServer(grpcServer, &presenter.Server{})

	// create and register health server
	healthServer := ghealth.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	return
}

func withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(
		presenter.UnaryInterceptorChainer(
			presenter.ErrRepackagingInterceptor,
			presenter.DummyInterceptor))
}

func startGRPCServer(grpcServer *grpc.Server, listener net.Listener, port string) (err error) {
	// start the server
	log.Printf("starting gRPC server on :%s", port)
	err = grpcServer.Serve(listener)
	return
}
