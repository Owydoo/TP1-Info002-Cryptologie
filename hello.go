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
		fmt.Printf("%x", hash_MD5(os.Args[2]))
	case "FP_SHA1":
		fmt.Printf("%x", hash_SHA1(os.Args[2]))
	case "CALCUL_N": // Question 2
		alphabet := os.Args[2]
		size_min, _ := strconv.Atoi(os.Args[3])
		size_max, _ := strconv.Atoi(os.Args[4])
		fmt.Printf("%d", nb_posibilities(alphabet, size_min, size_max))
	case "INDEX":
		alphabet := os.Args[2]
		index, _ := strconv.Atoi(os.Args[3])
		fmt.Printf(getTextFromIndex(alphabet, index))
	}

}

func nb_posibilities(alphabet string, size_min int, size_max int) int {
	var n float64 = 0
	for i := size_min; i <= size_max; i++ {
		n += math.Pow(float64(len(alphabet)), float64(i))
	}
	return int(n)
}

func hash_MD5(text string) [16]byte {
	data := []byte(text)
	return md5.Sum(data)
}

func hash_SHA1(text string) [20]byte {
	data := []byte(text)
	return sha1.Sum(data)
}

func getTextFromIndex(alphabet string, index int) string {
	var len_alphabet = len(alphabet)
	var alphabet_product = len_alphabet
	var nbLetters int = 0
	var indextemp int = index
	var result string = ""
	for indextemp >= alphabet_product {
		indextemp -= alphabet_product
		alphabet_product *= alphabet_product
		nbLetters++
	}
	for i := 0; i <= nbLetters; i++ {
		result = result + string(alphabet[indextemp%len(alphabet)])
		indextemp /= len(alphabet)
	}
	return Reverse(result)

}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
