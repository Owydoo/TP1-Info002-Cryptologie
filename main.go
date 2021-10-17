package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
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
		fmt.Printf("%d", nbPossibilities(len(alphabet), sizeMin, sizeMax))
	case "INDEX":
		alphabet := os.Args[2]
		sizeMin, _ := strconv.Atoi(os.Args[3])
		sizeMax, _ := strconv.Atoi(os.Args[4])
		index, _ := strconv.Atoi(os.Args[5])

		//Calcul useful things
		lenAlphabet := len(alphabet)
		tabN := nbPosibilitiesTab(lenAlphabet, sizeMin, sizeMax)
		// Display Tab N
		for _, value := range tabN {
			fmt.Printf("%d\n", value)
		}

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


