package main

func nbPossibilities(lenAlphabet int, sizeMin int, sizeMax int) int {
	var n  = 0
	for i := sizeMin; i <= sizeMax; i++ {
		n += pow(lenAlphabet, i)
	}
	return n
}

//for i2C
func nbPosibilitiesTab(lenAlphabet int, sizeMin int, sizeMax int) []int {
	var tabRes = make([]int, (sizeMax-sizeMin)+1)

	var iteration= 0
	for i := sizeMin; i <= sizeMax; i++ {
		tabRes[iteration] = pow(lenAlphabet, i)
		iteration++
	}

	return tabRes
}
