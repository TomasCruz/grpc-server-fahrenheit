package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/TomasCruz/grpc-server-fahrenheit/configuration"
	"github.com/TomasCruz/grpc-server-fahrenheit/database"
	"github.com/TomasCruz/grpc-server-fahrenheit/model"
	"github.com/TomasCruz/grpc-server-fahrenheit/presenter"
	"google.golang.org/grpc"
)

func readAndCheckEnvVar(varName string) (varVal string) {
	if varVal = os.Getenv(varName); varVal == "" {
		err := fmt.Errorf("%s environment variable not set properly", varName)
		log.Fatal(err)
	}

	return
}

func setupFromEnvVars() (config configuration.Config) {
	config.Port = readAndCheckEnvVar("GRPC_SERVER_PORT")
	config.DB = readAndCheckEnvVar("GRPC_SERVER_DB")
	return
}

func startGRPCServer(port string) error {
	address := fmt.Sprintf("localhost:%s", port)

	// create a listener on TCP port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	// create a server instance
	s := presenter.Server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Convertor service to the server
	presenter.RegisterConvertorServer(grpcServer, &s)

	// start the server
	log.Printf("starting gRPC server on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}

	return nil
}

func main() {
	// populate configuration
	config := setupFromEnvVars()
	if _, err := strconv.Atoi(config.Port); err != nil {
		log.Fatalf("Port environment variable has to be an integer")
	}

	// set DB interface to service
	databaseInterface := database.InitializeDatabase(config.DB)
	model.SetDatabaseInterface(databaseInterface)

	// fire the gRPC server in a goroutine
	go func() {
		if err := startGRPCServer(config.Port); err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	select {}
}
