//go:build !rand_insecure_for_docs

package rand

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand/v2"
	"os"
)

//nolint:gochecknoglobals,gosec
var globalRand = rand.New(cryptoSource{})

func BinaryN(n int) (string, error) {
	v := make([]byte, n)
	if err := binary.Read(cryptoRand.Reader, binary.BigEndian, &v); err != nil {
		return "", err
	}
	return string(v), nil
}

type cryptoSource struct{}

func (s cryptoSource) Uint64() uint64 {
	var v uint64
	if err := binary.Read(cryptoRand.Reader, binary.BigEndian, &v); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return v
}
