package wordlist

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gabe565/pwgen-go/internal/config"
	"github.com/gabe565/pwgen-go/internal/rand"
)

//go:generate go run ./generate

var ErrUnknownName = errors.New("unknown wordlist name")

func New(name string) (Wordlist, error) {
	switch name {
	case config.WordlistLong:
		return EFF_Long(), nil
	case config.WordlistShort1:
		return EFF_Short1(), nil
	case config.WordlistShort2:
		return EFF_Short2(), nil
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnknownName, name)
	}
}

type Wordlist []string

func (w Wordlist) RandWord() string {
	lineNum := rand.IntN(len(w))
	return w[lineNum]
}

func (w Wordlist) RandWords(n int) []string {
	result := make([]string, 0, n)
	for range n {
		result = append(result, w.RandWord())
	}
	return result
}

func (w Wordlist) RandWordsWithNumber(n int) []string {
	words := w.RandWords(n)
	if len(words) != 0 {
		words[len(words)-1] += strconv.Itoa(rand.IntN(10))
		rand.ShuffleSlice(words)
	}
	return words
}
