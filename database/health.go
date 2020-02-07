package database

// Health verifies the schema
func (rDb radisDb) Health() (status bool, err error) {
	conn := rDb.pool.Get()
	defer conn.Close()

	// Test the connection
	r, err := conn.Do("PING")
	if err != nil {
		return
	}

	status = r.(string) == "PONG"
	return
}
