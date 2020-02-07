package configuration

// Config holds all the globals and environment variable values, populated on startup
type Config struct {
	Port      string
	DbHost    string
	DbPort    string
	DbReqPswd string
}
