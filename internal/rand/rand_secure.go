//go:build !rand_insecure_for_docs

package rand

import (
	"bufio"
	cryptoRand "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"log/slog"
	"math/rand/v2"
	"os"
)

//nolint:gochecknoglobals,gosec
var (
	source = &Source{
		r: bufio.NewReaderSize(cryptoRand.Reader, 1024),
	}
	globalRand = rand.New(source)
)

func BinaryN(n int) (string, error) {
	b := make([]byte, n)
	_, err := source.Read(b)
	return string(b), err
}

func BytesN(n int) (string, error) {
	b := make([]byte, n)
	if _, err := source.Read(b); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

type Source struct {
	r *bufio.Reader
}

func (s *Source) Uint64() uint64 {
	b := make([]byte, 8)
	if _, err := s.Read(b); err != nil {
		fmt.Println() //nolint:forbidigo
		slog.Error("Crypto read failed", "error", err)
		os.Exit(1)
	}
	return binary.BigEndian.Uint64(b)
}

func (s *Source) Read(p []byte) (int, error) {
	return io.ReadFull(s.r, p)
}
