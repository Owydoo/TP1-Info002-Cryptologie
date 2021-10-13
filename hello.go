package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"os"
)

func main() {
	arg1 := os.Args[1]
	arg2 := os.Args[2]

	switch arg1 {
	case "MD5":
		fmt.Printf("%x", hash_MD5(arg2))
	case "SHA1":
		fmt.Printf("%x", hash_SHA1(arg2))
	}
}

func hash_MD5(text string) [16]byte {
	data := []byte(text)
	return md5.Sum(data)
}

func hash_SHA1(text string) [20]byte {
	data := []byte(text)
	return sha1.Sum(data)
}
