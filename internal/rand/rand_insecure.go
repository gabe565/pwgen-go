//go:build rand_insecure_for_docs

package rand

import (
	"math/rand/v2"

	"github.com/fatih/color"
)

func init() {
	color.New(color.FgRed, color.Bold).Println(`
WARNING: DO NOT USE THESE PASSWORDS!
This binary was built with an insecure flag that generates predictable passwords specifically for documentation purposes.
`)
}

//nolint:gochecknoglobals,gosec
var Rand = rand.New(rand.NewPCG(1, 2))

func BinaryN(n int) (string, error) {
	v := make([]byte, 0, n)
	for range n {
		v = append(v, byte(Rand.Uint64()))
	}
	return string(v), nil
}
