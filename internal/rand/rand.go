package rand

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand/v2"
	"os"
)

//nolint:gochecknoglobals
var Rand = rand.New(cryptoSource{})

func ShuffleSlice[T any](s []T) []T {
	Rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	return s
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
