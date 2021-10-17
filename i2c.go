package main

func i2c(alphabet string, lenAlphabet int, index int, minSize int) string {
	var nbPossibilitiesMinSize = nbPossibilities(lenAlphabet, 1, minSize - 1)
	var nbLetters = 0
	var indexTemp = index  + nbPossibilitiesMinSize
	var powStock = lenAlphabet
	for indexTemp >= powStock {
		indexTemp -= powStock
		powStock *= lenAlphabet
		nbLetters++
	}

	var result = ""
	for i := 0; i <= nbLetters; i++ {
		result = result + string(alphabet[indexTemp % lenAlphabet])
		indexTemp /= len(alphabet)
	}
	return reverse(result)
}
