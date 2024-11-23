package mock

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

const (
	intNumber int64 = 99999
)

func RandomInt() int {
	maxNumber := big.NewInt(intNumber)
	n, _ := rand.Int(rand.Reader, maxNumber)
	return int(n.Int64())
}

func RandomString() string {
	return strconv.Itoa(RandomInt())
}
