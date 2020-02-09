package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/grpc-server-fahrenheit/configuration"
)

func setupFromEnvVars() (config configuration.Config) {
	config.Port = readAndCheckIntEnvVar("GRPC_PORT")
	config.DbPort = readAndCheckIntEnvVar("GRPC_DB_PORT")
	config.DbReqPswd = readEnvVar("GRPC_DB_REQ_PSWD")
	return
}

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
