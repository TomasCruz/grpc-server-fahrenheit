package database

func (pDb postgresDb) C2F(cels float64) (fahr float64, err error) {
	fahr = 32 + cels*9/5
	return
}

func (pDb postgresDb) F2C(fahr float64) (cels float64, err error) {
	cels = (fahr - 32) * 5 / 9
	return
}
