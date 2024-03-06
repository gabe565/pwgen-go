package util

import (
	"crypto/rand"
	"math/big"
)

func CryptoRandn(n int) (int, error) {
	val, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		return 0, err
	}
	return int(val.Int64()), nil
}
