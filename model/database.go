package model

// Database is DB interface
type Database interface {
	Health() (status bool, err error)
	C2F(cels float64) (fahr float64, err error)
	F2C(fahr float64) (cels float64, err error)
}
