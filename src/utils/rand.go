package utils

import (
	"crypto/rand"
	"math/big"
)

var letters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetRandomBytes(n int) []byte {
	resultBytes := make([]byte, n)
	for i := range resultBytes {
		resultBytes[i] = letters[GetRandomNumber(int64(len(letters)))]
	}
	return resultBytes
}

func GetRandomNumber(max int64) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	return nBig.Int64()
}
