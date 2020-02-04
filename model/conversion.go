package model

// C2F returns appropriate number of fahrenheits, storing and caching the pair
func C2F(cels float64) (fahr float64, err error) {
	fahr = 32 + cels*9/5
	return
}

// F2C returns appropriate number of celsius, storing and caching the pair
func F2C(fahr float64) (cels float64, err error) {
	cels = (fahr - 32) * 5 / 9
	return
}
