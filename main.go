package main

import (
	"fmt"
	"os"
	"strconv"
)

var alphabet string
var sizeMin uint64
var sizeMax uint64
var lenAlphabet uint64
var N uint64
var tabN []uint64
var hashMethod string

func main() {

	// Init global values
	// go run . FUNC H ALPHABET SIZE_MIN SIZE_MAX
	function := os.Args[1]
	hashMethod = os.Args[2]
	alphabet = os.Args[3]
	lenAlphabet = uint64(len(alphabet))
	sizeMinTemp, _ := strconv.Atoi(os.Args[4])
	sizeMin = uint64(sizeMinTemp)

	sizeMaxTemp, _ := strconv.Atoi(os.Args[5])
	sizeMax = uint64(sizeMaxTemp)
	N = nbPossibilities(lenAlphabet, sizeMin, sizeMax)
	tabN = nbPosibilitiesTab(lenAlphabet, sizeMin, sizeMax)

	fmt.Printf("%d\n", h2i(hash("oups"),1))

	switch function {
	case "FP_MD5":
		fmt.Printf("%x", hashMd5(os.Args[2]))
	case "FP_SHA1":
		fmt.Printf("%x", hashSha1(os.Args[2]))
	case "N": // Question 2
		fmt.Printf("%d", nbPossibilities(uint64(len(alphabet)), uint64(sizeMin), uint64(sizeMax)))
	case "INDEX":
		index, _ := strconv.Atoi(os.Args[6])
		//tabN := nbPosibilitiesTab(lenAlphabet, sizeMin, sizeMax)
		// Display Tab N
		/* for _, value := range tabN {
			fmt.Printf("%d\n", value)
		} */

		fmt.Printf(i2c(uint64(index)))

	}

}
