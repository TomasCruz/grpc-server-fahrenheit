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

func readEnvVar(varName string) string {
	return os.Getenv(varName)
}

func readAndCheckEnvVar(varName string) (varVal string) {
	if varVal = readEnvVar(varName); varVal == "" {
		err := fmt.Errorf("%s environment variable not set properly", varName)
		log.Fatal(err)
	}

	return
}

func readAndCheckIntEnvVar(varName string) (varVal string) {
	varVal = readAndCheckEnvVar(varName)
	if _, err := strconv.Atoi(varVal); err != nil {
		err := fmt.Errorf("Value of %s environment variable has to be an integer", varName)
		log.Fatal(err)
	}

	return
}

func setupFromEnvVars() (config configuration.Config) {
	config.Port = readAndCheckIntEnvVar("GRPC_PORT")
	config.DbPort = readAndCheckIntEnvVar("GRPC_DB_PORT")
	config.DbReqPswd = readEnvVar("GRPC_DB_REQ_PSWD")
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

func composeRedisDbURL(c configuration.Config) (url string) {
	//url = fmt.Sprintf("redis://guest:guest@%s:%s/0", c.DbHost, c.DbPort)
	url = fmt.Sprintf("redis://%s:%s/0", c.DbHost, c.DbPort)
	if c.DbReqPswd != "" {
		url = fmt.Sprintf("%s?password=%s", url, c.DbReqPswd)
	}

	return
}

func main() {
	// populate configuration
	config := setupFromEnvVars()

	// set DB interface to service
	databaseInterface := database.InitializeDatabase(composeRedisDbURL(config))
	model.SetDatabase(databaseInterface)

	// fire the gRPC server in a goroutine
	go func() {
		if err := startGRPCServer(config.Port); err != nil {
			log.Fatalf("failed to start gRPC server: %s", err)
		}
	}()

	select {}
}
