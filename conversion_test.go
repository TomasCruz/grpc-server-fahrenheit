// +build integration

package main

import (
	"context"
	"testing"
	"time"

	"github.com/TomasCruz/grpc-server-fahrenheit/presenter"

	"gotest.tools/assert"
)

func TestC2F(t *testing.T) {
	conn, err := testSetup()
	assert.NilError(t, err)
	defer conn.Close()

	c := presenter.NewConvertorClient(conn)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(100)*time.Millisecond)
	defer cancel()

	var response *presenter.ConversionMsg

	response, err = c.C2F(ctx, &presenter.ConversionMsg{Input: -40})
	assert.NilError(t, err)
	assert.Assert(t, response.Input == -40.0)

	response, err = c.C2F(ctx, &presenter.ConversionMsg{Input: 0})
	assert.NilError(t, err)
	assert.Assert(t, response.Input == 32)

	response, err = c.C2F(ctx, &presenter.ConversionMsg{Input: 100})
	assert.NilError(t, err)
	assert.Assert(t, response.Input == 212)
}

func TestF2C(t *testing.T) {
	conn, err := testSetup()
	assert.NilError(t, err)
	defer conn.Close()

	c := presenter.NewConvertorClient(conn)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(100)*time.Millisecond)
	defer cancel()

	var response *presenter.ConversionMsg

	response, err = c.F2C(ctx, &presenter.ConversionMsg{Input: -40})
	assert.NilError(t, err)
	assert.Assert(t, response.Input == -40.0)

	response, err = c.F2C(ctx, &presenter.ConversionMsg{Input: 32})
	assert.NilError(t, err)
	assert.Assert(t, response.Input == 0)

	response, err = c.F2C(ctx, &presenter.ConversionMsg{Input: 212})
	assert.NilError(t, err)
	assert.Assert(t, response.Input == 100)
}
