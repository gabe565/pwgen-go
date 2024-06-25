package rand

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
)

//nolint:gochecknoglobals,gosec
var Rand = rand.New(cryptoSource{})

func ShuffleSlice[T any](s []T) []T {
	Rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	return s
}

func BinaryN(n int) (string, error) {
	v := make([]byte, n)
	if err := binary.Read(cryptoRand.Reader, binary.BigEndian, &v); err != nil {
		return "", err
	}
	return string(v), nil
}

func Numeric(n int) string {
	var result string
	for range n {
		result += strconv.Itoa(Rand.IntN(10))
	}
	return result
}

func Alpha(n int) string {
	const bytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	result := make([]byte, 0, n)
	for range n {
		result = append(result, bytes[Rand.IntN(len(bytes))])
	}
	return string(result)
}

func AlphaNum(n int) string {
	const bytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	result := make([]byte, 0, n)
	for range n {
		result = append(result, bytes[Rand.IntN(len(bytes))])
	}
	return string(result)
}

func ASCII(n int) string {
	result := make([]byte, 0, n)
	for range n {
		// Generate random bytes between 32 (space) and 126 (~)
		result = append(result, byte(Rand.IntN(127-32)+32))
	}
	return string(result)
}

type cryptoSource struct{}

func (s cryptoSource) Seed(_ int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() uint64 {
	var v uint64
	if err := binary.Read(cryptoRand.Reader, binary.BigEndian, &v); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return v
}
