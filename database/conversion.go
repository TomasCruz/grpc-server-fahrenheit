package database

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func (rDb radisDb) GetByC(cels float64) (fahr float64, notFound bool, err error) {
	conn := rDb.pool.Get()
	defer conn.Close()

	if fahr, err = rDb.get(conn, "C", cels); err == redis.ErrNil {
		notFound = true
		err = nil
	}

	return
}

func (rDb radisDb) GetByF(fahr float64) (cels float64, notFound bool, err error) {
	conn := rDb.pool.Get()
	defer conn.Close()

	if cels, err = rDb.get(conn, "F", fahr); err == redis.ErrNil {
		notFound = true
		err = nil
	}

	return
}

func (rDb radisDb) SetPair(cels, fahr float64) (err error) {
	conn := rDb.pool.Get()
	defer conn.Close()

	if err = conn.Send("MULTI"); err != nil {
		return
	}

	if err = rDb.set(conn, "C", cels, fahr); err != nil {
		return
	}

	if err = rDb.set(conn, "F", fahr, cels); err != nil {
		return
	}

	if _, err = conn.Do("EXEC"); err != nil {
		return
	}

	return
}

func (rDb radisDb) get(conn redis.Conn, id string, deg float64) (out float64, err error) {
	key := fmt.Sprintf("%s%e", id, deg)
	out, err = redis.Float64(conn.Do("GET", key))
	return
}

func (rDb radisDb) set(conn redis.Conn, id string, deg1 float64, deg2 float64) (err error) {
	key := fmt.Sprintf("%s%e", id, deg1)
	_, err = conn.Do("SET", key, deg2)
	return
}
