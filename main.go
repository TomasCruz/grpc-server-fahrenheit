package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/TomasCruz/grpc-server-fahrenheit/configuration"
	"github.com/TomasCruz/grpc-server-fahrenheit/database"
	"github.com/TomasCruz/grpc-server-fahrenheit/model"
	"google.golang.org/grpc"
)

func main() {
	// populate configuration
	config := setupFromEnvVars()

	// set DB interface to service
	databaseInterface := database.InitializeDatabase(composeRedisDbURL(config))
	model.SetDatabase(databaseInterface)

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)

	// register gRPC server
	var (
		grpcServer *grpc.Server
		listener   net.Listener
		err        error
	)
	if grpcServer, listener, err = registerGRPCServer(config.Port); err != nil {
		log.Fatalf("failed to register gRPC server: %s", err)
	}

	// fire up the gRPC server
	go func() {
		if err = startGRPCServer(grpcServer, listener, config.Port); err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	<-stop
	gracefulShutdown(grpcServer)
}

func setupFromEnvVars() (config configuration.Config) {
	config.Port = readAndCheckIntEnvVar("GRPC_PORT")
	config.DbPort = readAndCheckIntEnvVar("GRPC_DB_PORT")
	config.DbReqPswd = readEnvVar("GRPC_DB_REQ_PSWD")
	return
}

func composeRedisDbURL(c configuration.Config) (url string) {
	url = fmt.Sprintf("redis://%s:%s/0", c.DbHost, c.DbPort)
	if c.DbReqPswd != "" {
		url = fmt.Sprintf("%s?password=%s", url, c.DbReqPswd)
	}

	return
}

func gracefulShutdown(grpcServer *grpc.Server) {
	grpcServer.GracefulStop()
	//fmt.Println("Graceful shutdown")
}
