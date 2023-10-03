package wordlists

import (
	"bytes"
	_ "embed"
)

//go:embed en.txt
var En []byte

var EnLines int

func init() {
	EnLines = bytes.Count(En, []byte{'\n'})
}
