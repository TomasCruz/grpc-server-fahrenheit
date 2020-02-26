package model

// Database is DB interface
type Database interface {
	GetByC(cels float64) (fahr float64, notFound bool, err error)
	GetByF(fahr float64) (cels float64, notFound bool, err error)
	SetPair(cels, fahr float64) (err error)
}
