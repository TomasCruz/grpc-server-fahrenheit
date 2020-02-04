// +build integration

package main

import (
	"fmt"

	"google.golang.org/grpc"
)

func testSetup() (*grpc.ClientConn, error) {
	port := readAndCheckEnvVar("GRPC_SERVER_PORT")
	return grpc.Dial(fmt.Sprintf("localhost:%s", port), grpc.WithInsecure())
}
