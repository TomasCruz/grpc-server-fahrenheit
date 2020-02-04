// +build integration

package main

import (
	"context"
	"testing"
	"time"

	"github.com/TomasCruz/grpc-server-fahrenheit/presenter"

	"gotest.tools/assert"
)

func TestHealthOK(t *testing.T) {
	conn, err := testSetup()
	assert.NilError(t, err)
	defer conn.Close()

	c := presenter.NewConvertorClient(conn)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(100)*time.Millisecond)
	defer cancel()

	response, err := c.Health(ctx, &presenter.NoParamsMsg{})

	assert.NilError(t, err)
	assert.Assert(t, response.Health == true)
}
