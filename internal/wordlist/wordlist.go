package wordlist

import (
	"github.com/gabe565/pwgen-go/internal/util"
)

//go:generate go run ./generate

type Wordlist []string

func (w Wordlist) RandWord() (string, error) {
	lineNum, err := util.CryptoRandn(len(w))
	if err != nil {
		return "", err
	}

	return w[lineNum], nil
}

func (w Wordlist) RandWords(n int) ([]string, error) {
	result := make([]string, 0, n)
	for i := 0; i < n; i += 1 {
		word, err := w.RandWord()
		if err != nil {
			return result, err
		}
		result = append(result, word)
	}
	return result, nil
}
