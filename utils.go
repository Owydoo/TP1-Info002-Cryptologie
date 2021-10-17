package main

import (
	"fmt"
	"math"
)

func pow(i1 int, i2 int) int {
	return int(math.Pow(float64(i1), float64(i2)))
}

func logInt(i int) {
	fmt.Printf("%d\n", i)
}

func reverse(s string) (result string) {
	for _, v := range s {
		result= string(v) + result
	}
	return
}