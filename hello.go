package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	H := os.Args[1]

	switch H {
	case "FP_MD5":
		fmt.Printf("%x", hashMd5(os.Args[2]))
	case "FP_SHA1":
		fmt.Printf("%x", hashSha1(os.Args[2]))
	case "N": // Question 2
		alphabet := os.Args[2]
		sizeMin, _ := strconv.Atoi(os.Args[3])
		sizeMax, _ := strconv.Atoi(os.Args[4])
		fmt.Printf("%d", nbPossibilities(alphabet, sizeMin, sizeMax))
	case "INDEX":
		alphabet := os.Args[2]
		lenAlphabet := len(alphabet)
		index, _ := strconv.Atoi(os.Args[3])
		sizeMin, _ := strconv.Atoi(os.Args[4])
		_, _ = strconv.Atoi(os.Args[5])
		fmt.Printf(i2c(alphabet,lenAlphabet, index, sizeMin))
	}

}

func hashMd5(text string) [16]byte {
	data := []byte(text)
	return md5.Sum(data)
}

func hashSha1(text string) [20]byte {
	data := []byte(text)
	return sha1.Sum(data)
}

func nbPossibilities(alphabet string, sizeMin int, sizeMax int) int {
	var n float64 = 0
	for i := sizeMin; i <= sizeMax; i++ {
		n += math.Pow(float64(len(alphabet)), float64(i))
	}
	return int(n)
}

func i2c(alphabet string, lenAlphabet int, index int, minSize int) string {
	var nbPossibilitiesMinSize = nbPossibilities(alphabet, 1, minSize - 1)
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

