package database

// Health verifies the schema
func (rDb radisDb) Health() (status bool, err error) {
	conn := rDb.pool.Get()
	defer conn.Close()

	// Test the connection
	status, err = pingPong(conn)
	return
}
