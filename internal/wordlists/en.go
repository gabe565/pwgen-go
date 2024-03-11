package wordlists

import (
	"bytes"
	_ "embed"
)

//go:embed en.txt
var En []byte

var EnLines = bytes.Count(En, []byte{'\n'})
