package main

func nbPossibilities(lenAlphabet uint64, sizeMin uint64, sizeMax uint64) uint64 {
	var n uint64 = 0
	for i := sizeMin; i <= sizeMax; i++ {
		n += pow(lenAlphabet, i)
	}
	return n
}

//for i2C
func nbPosibilitiesTab(lenAlphabet uint64, sizeMin uint64, sizeMax uint64) []uint64 {
	var tabRes = make([]uint64, (sizeMax-sizeMin)+1)

	var iteration= 0
	for i := sizeMin; i <= sizeMax; i++ {
		tabRes[iteration] = pow(lenAlphabet, i)
		iteration++
	}

	return tabRes
}
