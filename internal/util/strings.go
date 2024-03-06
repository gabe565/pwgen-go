package util

import (
	"bufio"
	"bytes"
	"errors"
	"io"

	"github.com/gabe565/pwgen-go/internal/wordlists"
)

func RandWord() ([]byte, error) {
	var line int
	for line < 32 {
		var err error
		if line, err = CryptoRandn(wordlists.EnLines); err != nil {
			return nil, err
		}
	}

	return GetLine(bytes.NewReader(wordlists.En), line)
}

func RandWords(n int) ([][]byte, error) {
	result := make([][]byte, 0, n)
	for i := 0; i < n; i += 1 {
		word, err := RandWord()
		if err != nil {
			return result, err
		}
		result = append(result, word)
	}
	return result, nil
}

var ErrNotEnoughLines = errors.New("not enough lines")

func GetLine(r io.Reader, line int) ([]byte, error) {
	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan(); i += 1 {
		if i == line {
			return scanner.Bytes(), nil
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return nil, ErrNotEnoughLines
}
