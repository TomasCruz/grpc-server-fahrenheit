package model

var (
	database Database
)

// SetDatabase sets DB interface
func SetDatabase(dBase Database) {
	database = dBase
}
