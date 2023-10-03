package util

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"math/rand"

	"github.com/gabe565/pwgen-go/internal/wordlists"
)

func RandWord() (string, error) {
	var line int
	for line < 32 {
		line = rand.Intn(wordlists.EnLines)
	}

	return GetLine(bytes.NewReader(wordlists.En), line)
}

func RandWords(n uint) ([]string, error) {
	result := make([]string, 0, n)
	for i := uint(0); i < n; i += 1 {
		word, err := RandWord()
		if err != nil {
			return result, err
		}
		result = append(result, word)
	}
	return result, nil
}

var ErrNotEnoughLines = errors.New("not enough lines")

func GetLine(r io.Reader, line int) (string, error) {
	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan(); i += 1 {
		if i == line {
			return scanner.Text(), nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", ErrNotEnoughLines
}
