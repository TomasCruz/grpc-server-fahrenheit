package database

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

// Health verifies the schema
func (rDb radisDb) Health() (status bool, err error) {
	conn := rDb.pool.Get()
	defer conn.Close()

	// Test the connection
	status, err = pingPong(conn)
	return
}

func pingPong(conn redis.Conn) (status bool, err error) {
	var r interface{}
	if r, err = conn.Do("PING"); err != nil {
		err = errors.WithStack(err)
		return
	}

	status = r.(string) == "PONG"
	return
}
