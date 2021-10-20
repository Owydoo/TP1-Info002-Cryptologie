package main

import (
	"encoding/binary"
)

//Q6
func i2i(t uint64, index uint64) uint64 {
	return h2i(hash(i2c(index)), t)
}

//Q5
func h2i(hash []byte, t uint64) uint64 {
	first8o := hash[:8]
	data := binary.LittleEndian.Uint64(first8o)
	return (data + t) % N
}

//Q6
func newString(index uint64, width uint64) uint64 {
	var newIndex uint64 = index
	for i := 1; i < int(width); i++ {
		newIndex = i2i(newIndex, uint64(i))
	}
	return newIndex
}

func i2c(index uint64) string {
	var nbPossibilitiesMinSize = nbPossibilities(lenAlphabet, 1, sizeMin-1)
	var nbLetters uint64 = 0
	var indexTemp = index + nbPossibilitiesMinSize
	var powStock = lenAlphabet
	for indexTemp >= powStock {
		indexTemp -= powStock
		powStock *= lenAlphabet
		nbLetters++
	}

	var result = ""
	for i := 0; i <= int(nbLetters); i++ {
		result = result + string(alphabet[indexTemp%lenAlphabet])
		indexTemp /= lenAlphabet
	}
	return reverse(result)
}
