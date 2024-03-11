package wordlists

import (
	"bytes"
	_ "embed"
)

//go:embed eff_large_wordlist.txt
var Eff []byte

var EffLines = bytes.Count(Eff, []byte{'\n'})
