package main

import (
	"fmt"
	"math"
)

func pow(i1 uint64, i2 uint64) uint64 {
	return uint64(math.Pow(float64(i1), float64(i2)))
}

func logInt(i uint64) {
	fmt.Printf("%d\n", i)
}

func reverse(s string) (result string) {
	for _, v := range s {
		result= string(v) + result
	}
	return
}