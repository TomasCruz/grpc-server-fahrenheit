// +build integration

package main

import (
	"context"
	"testing"
	"time"

	"github.com/TomasCruz/grpc-server-fahrenheit/api"

	"gotest.tools/assert"
)

func TestHealthOK(t *testing.T) {
	conn, err := testSetup()
	assert.NilError(t, err)
	defer conn.Close()

	c := api.NewConvertorClient(conn)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(100)*time.Millisecond)
	defer cancel()

	response, err := c.Health(ctx, &api.NoParamsMsg{})

	assert.NilError(t, err)
	assert.Assert(t, response.Health == true)
}
