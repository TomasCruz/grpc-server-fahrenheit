package model

var database Database

// SetDatabaseInterface sets DB interface
func SetDatabaseInterface(dBase Database) {
	database = dBase
}
