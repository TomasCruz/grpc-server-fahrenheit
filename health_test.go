// +build integration

package main

import (
	"context"
	"testing"
	"time"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"gotest.tools/assert"
)

func TestHealthOK(t *testing.T) {
	conn, err := testSetup()
	assert.NilError(t, err)
	defer conn.Close()

	healthClient := healthpb.NewHealthClient(conn)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(100)*time.Millisecond)
	defer cancel()

	response, err := healthClient.Check(ctx, &healthpb.HealthCheckRequest{})

	assert.NilError(t, err)
	assert.Assert(t, response.Status == healthpb.HealthCheckResponse_SERVING)
}
