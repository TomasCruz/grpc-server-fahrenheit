// +build integration

package main

import (
	"context"
	"testing"
	"time"

	"github.com/TomasCruz/grpc-server-fahrenheit/api"

	"gotest.tools/assert"
)

func TestC2F(t *testing.T) {
	conn, err := testSetup()
	assert.NilError(t, err)
	defer conn.Close()

	c := api.NewConvertorClient(conn)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(100)*time.Millisecond)
	defer cancel()

	var response *api.ConversionMsg

	response, err = c.C2F(ctx, &api.ConversionMsg{Number: -40})
	assert.NilError(t, err)
	assert.Assert(t, response.Number == -40.0)

	response, err = c.C2F(ctx, &api.ConversionMsg{Number: 0})
	assert.NilError(t, err)
	assert.Assert(t, response.Number == 32)

	response, err = c.C2F(ctx, &api.ConversionMsg{Number: 100})
	assert.NilError(t, err)
	assert.Assert(t, response.Number == 212)
}

func TestF2C(t *testing.T) {
	conn, err := testSetup()
	assert.NilError(t, err)
	defer conn.Close()

	c := api.NewConvertorClient(conn)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(100)*time.Millisecond)
	defer cancel()

	var response *api.ConversionMsg

	response, err = c.F2C(ctx, &api.ConversionMsg{Number: -40})
	assert.NilError(t, err)
	assert.Assert(t, response.Number == -40.0)

	response, err = c.F2C(ctx, &api.ConversionMsg{Number: 32})
	assert.NilError(t, err)
	assert.Assert(t, response.Number == 0)

	response, err = c.F2C(ctx, &api.ConversionMsg{Number: 212})
	assert.NilError(t, err)
	assert.Assert(t, response.Number == 100)
}
