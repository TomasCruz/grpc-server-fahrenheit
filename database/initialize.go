package database

import (
	"strings"
	"time"

	"github.com/TomasCruz/grpc-server-fahrenheit/model"
	"github.com/gomodule/redigo/redis"
)

type radisDb struct {
	pool             *redis.Pool
	requiredPassword string
}

var rDb radisDb

// InitializeDatabase verifies DB accessibility
func InitializeDatabase(dbURL string) model.Database {
	rDb = radisDb{}
	pswdStaticString := "password="
	pswdStringIndex := strings.Index(dbURL, pswdStaticString)
	if pswdStringIndex != -1 {
		rDb.requiredPassword = dbURL[pswdStringIndex+len(pswdStaticString):]
	}

	rDb.pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (conn redis.Conn, err error) {
			if conn, err = redis.DialURL(dbURL); err != nil {
				return
			}

			if rDb.requiredPassword != "" {
				if _, err = conn.Do("AUTH", rDb.requiredPassword); err != nil {
					return
				}
			}

			return
		},
	}

	return rDb
}
