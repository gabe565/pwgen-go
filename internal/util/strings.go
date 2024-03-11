package util

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"

	"github.com/gabe565/pwgen-go/internal/wordlists"
)

func RandWord() (string, error) {
	var line int
	for line < 32 {
		var err error
		if line, err = CryptoRandn(wordlists.EffLines); err != nil {
			return "", err
		}
	}

	return GetLine(bytes.NewReader(wordlists.Eff), line)
}

func RandWords(n int) ([]string, error) {
	result := make([]string, 0, n)
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

func GetLine(r io.Reader, line int) (string, error) {
	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan(); i += 1 {
		if i == line {
			split := strings.Split(scanner.Text(), "\t")
			return split[1], nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", ErrNotEnoughLines
}
