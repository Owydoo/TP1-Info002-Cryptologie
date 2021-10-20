package main

import (
	"crypto/md5"
	"crypto/sha1"
)

func hashMd5(text string) [16]byte {
	data := []byte(text)
	return md5.Sum(data)
}

func hashSha1(text string) [20]byte {
	data := []byte(text)
	return sha1.Sum(data)
}

func hash(text string) []byte {
	var temp []byte
	if hashMethod == "SHA" {
		hash := hashSha1(text)
		temp = hash[:]
	} else {
		hash := hashMd5(text)
		temp = hash[:]
	}

	return temp
}
