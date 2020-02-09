package configuration

// Config holds environment variable values, it's populated on startup
type Config struct {
	Port      string
	DbHost    string
	DbPort    string
	DbReqPswd string
}
