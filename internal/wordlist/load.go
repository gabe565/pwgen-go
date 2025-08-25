package wordlist

import (
	"bufio"
	"os"
)

func Load(name string) (Wordlist, error) {
	l, err := MetaString(name)
	if err == nil {
		return l.List(), nil
	}

	if name == "" {
		return Meta(0).List(), nil
	}

	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	scan := bufio.NewScanner(f)
	words := make([]string, 0, 1024)
	for scan.Scan() {
		if len(scan.Bytes()) != 0 {
			words = append(words, scan.Text())
		}
	}
	return words, nil
}
