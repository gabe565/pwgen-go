package wordlist

import (
	"strconv"

	"github.com/gabe565/pwgen-go/internal/rand"
)

//go:generate go run ./generate

type Wordlist []string

func (w Wordlist) RandWord() string {
	lineNum := rand.Rand.IntN(len(w))
	return w[lineNum]
}

func (w Wordlist) RandWords(n int) []string {
	result := make([]string, 0, n)
	for i := 0; i < n; i += 1 {
		result = append(result, w.RandWord())
	}
	return result
}

func (w Wordlist) RandWordsWithNumber(n int) []string {
	words := w.RandWords(n)
	words[len(words)-1] += strconv.Itoa(rand.Rand.IntN(10))
	rand.ShuffleSlice(words)
	return words
}
