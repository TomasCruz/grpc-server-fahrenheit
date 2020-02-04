package database

// Health verifies the schema
func (pDb postgresDb) Health() (status bool, err error) {
	var cnt int64
	if err = pDb.db.QueryRow("select count(*) from pg_catalog.pg_tables where tablename='degrees';").
		Scan(&cnt); err != nil {

		return
	}

	status = cnt == 1
	return
}
