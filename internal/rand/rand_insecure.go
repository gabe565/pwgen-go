//go:build rand_insecure_for_docs

package rand

import (
	"math/rand/v2"
	"strings"

	"github.com/fatih/color"
)

func init() {
	color.New(color.FgRed, color.Bold).Println(`
WARNING: DO NOT USE THESE PASSWORDS!
This binary was built with an insecure flag that generates predictable passwords specifically for documentation purposes.
`)
}

//nolint:gochecknoglobals,gosec
var globalRand = rand.New(rand.NewPCG(1, 2))

func BinaryN(n int) (string, error) {
	var s strings.Builder
	s.Grow(n)
	for range n {
		s.WriteByte(byte(globalRand.Uint64()))
	}
	return s.String(), nil
}
