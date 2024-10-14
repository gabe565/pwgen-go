package wordlist

import (
	"strconv"

	"gabe565.com/pwgen/internal/rand"
)

//go:generate go run ./generate

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
