package wordlist

import (
	"errors"
)

//go:generate go tool enumer -type Meta -transform lower -text -output meta_string.go

type Meta uint8

const (
	Long Meta = iota
	Short1
	Short2
)

func (m Meta) Description() string {
	switch m {
	case Long:
		return "is the default wordlist used for passphrases.\nIt contains 7776 words and is designed for memorability and password strength."
	case Short1:
		return "features only short words.\nIt contains 1296 words and may be more efficient to type, but suggests using more words to maintain security."
	case Short2:
		return "features longer words and may be more memorable.\nIt contains 1296 words and may be more efficient to type, but suggests using more words to maintain security."
	default:
		panic("unknown meta ID")
	}
}

func (m Meta) URL() string {
	switch m {
	case Long:
		return "https://eff.org/files/2016/07/18/eff_large_wordlist.txt"
	case Short1:
		return "https://eff.org/files/2016/09/08/eff_short_wordlist_1.txt"
	case Short2:
		return "https://eff.org/files/2016/09/08/eff_short_wordlist_2_0.txt"
	default:
		panic("unknown meta ID")
	}
}

func (m Meta) Var() string {
	switch m {
	case Long:
		return "EFF_Long"
	case Short1:
		return "EFF_Short1"
	case Short2:
		return "EFF_Short2"
	default:
		panic("unknown meta ID")
	}
}

var ErrUnknownName = errors.New("unknown wordlist name")

func (m Meta) List() (Wordlist, error) {
	switch m {
	case Long:
		return EFF_Long(), nil
	case Short1:
		return EFF_Short1(), nil
	case Short2:
		return EFF_Short2(), nil
	default:
		return nil, ErrUnknownName
	}
}
