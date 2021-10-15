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
		fmt.Printf(getTextFromIndex(alphabet, index, 2, 4))
	}

}

//Question 1
func hash_MD5(text string) [16]byte {
	data := []byte(text)
	return md5.Sum(data)
}

// Question 1
func hash_SHA1(text string) [20]byte {
	data := []byte(text)
	return sha1.Sum(data)
}

// Question 2
func nb_posibilities_Q2(alphabet string, size_min int, size_max int) int {
	var n float64 = 0
	for i := size_min; i <= size_max; i++ {
		n += math.Pow(float64(len(alphabet)), float64(i))
	}
	return int(n)
}

//for i2C
func nb_posibilities(alphabet string, size_min int, size_max int) []int {
	var len_alphabet float64 = float64(len(alphabet))

	var tab_res = make([]int, (size_max-size_min)+1)

	var iteration int = 0
	for i := size_min; i <= size_max; i++ {
		// nb_total += math.Pow(len_alphabet, float64(i))
		tab_res[iteration] = int(math.Pow(len_alphabet, float64(i)))
		iteration++
	}

	return tab_res
}

//Question 3
func getTextFromIndex_monoSize(alphabet string, index int, mono_size int) string {
	var len_alphabet = len(alphabet)
	var alphabet_product = len_alphabet
	var indextemp int = index
	var result string = ""
	for indextemp >= alphabet_product {
		indextemp -= alphabet_product
		alphabet_product *= alphabet_product
	}
	for i := 0; i <= mono_size; i++ {
		result = result + string(alphabet[indextemp%len(alphabet)])
		indextemp /= len(alphabet)
	}
	return Reverse(result)
}

//Question 3
func getTextFromIndex(alphabet string, index int, size_min int, size_max int) string {
	//on compte le nombre de possibilités de size_min à size_max
	var tab_possibilities []int = nb_posibilities(alphabet, size_min, size_max)

	//la bonne size est la size minimum qui est supérieur à index
	var trouve bool = false
	var good_size int = 0

	var i int = 0
	for !trouve {
		if tab_possibilities[i] >= index {
			good_size = i + 1
			trouve = true
		}
		i++
	}
	//on fait getTextFromIndex_monoSize avec la size trouvé
	return getTextFromIndex_monoSize(alphabet, index, good_size)
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
