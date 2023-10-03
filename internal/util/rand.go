package util

import (
	"crypto/rand"
	"math/big"

	"github.com/gabe565/pwgen-go/internal/wordlists"
)

func CryptoRandn(n int) (int, error) {
	val, err := rand.Int(rand.Reader, big.NewInt(int64(wordlists.EnLines)))
	if err != nil {
		return 0, err
	}
	return int(val.Int64()), nil
}
