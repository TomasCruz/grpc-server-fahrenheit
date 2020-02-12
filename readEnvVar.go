package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/grpc-server-fahrenheit/callstack"
	"github.com/pkg/errors"
)

func readEnvVar(varName string) string {
	return os.Getenv(varName)
}

func readAndCheckEnvVar(varName string) (varVal string) {
	if varVal = readEnvVar(varName); varVal == "" {
		err := errors.Errorf("%s environment variable not set properly", varName)
		callstack.LogErrStack(err)
		log.Fatal(err)
	}

	return
}

func readAndCheckIntEnvVar(varName string) (varVal string) {
	varVal = readAndCheckEnvVar(varName)
	if _, err := strconv.Atoi(varVal); err != nil {
		err = errors.Wrapf(err, "Value of %s environment variable has to be an integer", varName)
		callstack.LogErrStack(err)
		log.Fatal(err)
	}

	return
}
