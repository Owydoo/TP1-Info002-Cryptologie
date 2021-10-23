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

var gWidth uint64
var gHeight uint64

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

	switch function {
	case "FP_MD5": //Q1
		fmt.Printf("%x", hashMd5(os.Args[2]))
	case "FP_SHA1": //Q1
		fmt.Printf("%x", hashSha1(os.Args[2]))
	case "N": // Question 2
		fmt.Printf("%d", nbPossibilities(uint64(len(alphabet)), uint64(sizeMin), uint64(sizeMax)))
	case "INDEX": //Q3
		index, _ := strconv.Atoi(os.Args[6])
		//tabN := nbPosibilitiesTab(lenAlphabet, sizeMin, sizeMax)
		// Display Tab N
		/* for _, value := range tabN {
			fmt.Printf("%d\n", value)
		} */
		fmt.Printf(i2c(uint64(index)))
	case "Q5":
		fmt.Printf("%d\n", h2i(hash("oups"), 1)) //Q5
	case "Q6":
		width, _ := strconv.Atoi(os.Args[6])
		fmt.Printf("%d\n", newString(1, uint64(width))) //Q6
	case "Q8":
		width, _ := strconv.Atoi(os.Args[6])
		height, _ := strconv.Atoi(os.Args[7])
		fmt.Printf("%v\n", creerTable(uint64(width), uint64(height)))
		// sauveTable(creerTable(uint64(width), uint64(height)), uint64(width), uint64(height), "test.txt")
	case "Q9":
		width, _ := strconv.Atoi(os.Args[6])
		height, _ := strconv.Atoi(os.Args[7])
		var tab [][2]uint64 = creerTable(uint64(width), uint64(height))
		sauveTable(tab, uint64(width), uint64(height), "test.txt")
		poolSize, _ := strconv.Atoi(os.Args[8])
		afficheTable(ouvreTable("test.txt"), uint64(poolSize))

	case "RECHERCHE":
		//test de la recherche question 10
		tab := ouvreTable("test.txt")

		a, b := recherche(tab, gHeight, 10589183)
		fmt.Printf("%d %d", a, b)
	case "Q10":
		//hacher un mot
		mot_hashe := hash("eabd")
		var width uint64 = 1000
		var height uint64 = 100000

		//créer la rainbow table
		var tab = creerTable(width, height)
		err := sauveTable(tab, width, height, "test.txt")
		if err != nil {
			return 
		}
		//inverser
		clair, err := inverse(tab, height, width, mot_hashe)
		if err != nil {
			fmt.Println(err)
		}

		// fmt.Printf("%s", clair)
		fmt.Println("clair : ", clair)

	}
}
