// +build integration

package main

import (
	"fmt"

	"google.golang.org/grpc"
)

func testSetup() (*grpc.ClientConn, error) {
	port := readAndCheckEnvVar("GRPC_PORT")
	serviceConfig := grpc.WithDefaultServiceConfig(`{
		"healthCheckConfig": {
		  "serviceName": ""
		}
	  }`)

	return grpc.Dial(fmt.Sprintf("localhost:%s", port), grpc.WithInsecure(), serviceConfig)
}
