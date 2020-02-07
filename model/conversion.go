package model

// C2F returns appropriate number of fahrenheits, storing and caching the pair
func C2F(cels float64) (fahr float64, err error) {
	var notFound bool
	fahr, notFound, err = database.GetByC(cels)
	if notFound {
		err = nil
	} else {
		return
	}

	fahr = cels*9/5 + 32
	err = database.SetPair(cels, fahr)
	return
}

// F2C returns appropriate number of celsius, storing and caching the pair
func F2C(fahr float64) (cels float64, err error) {
	var notFound bool
	cels, notFound, err = database.GetByF(fahr)
	if notFound {
		err = nil
	} else {
		return
	}

	cels = (fahr - 32) * 5 / 9
	err = database.SetPair(cels, fahr)
	return
}
