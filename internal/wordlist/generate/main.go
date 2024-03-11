package main

import (
	"bufio"
	"net/http"
	"strings"

	. "github.com/dave/jennifer/jen"
)

func main() {
	type generatorConfig struct {
		url         string
		name        string
		description string
	}

	configs := []generatorConfig{
		{
			"https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt",
			"EFF_Long",
			"is the default wordlist used for passphrases.\nIt contains 7776 words and is designed for memorability and password strength.",
		},
		{
			"https://eff.org/files/2016/09/08/eff_short_wordlist_1.txt",
			"EFF_Short1",
			"features only short words.\nIt contains 1296 words and may be more efficient to type, but suggests using more words to maintain security.",
		},
		{
			"https://eff.org/files/2016/09/08/eff_short_wordlist_2_0.txt",
			"EFF_Short2",
			"features longer words and may be more memorable.\nIt contains 1296 words and may be more efficient to type, but suggests using more words to maintain security.",
		},
	}

	for _, config := range configs {
		resp, err := http.Get(config.url)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != http.StatusOK {
			panic("invalid response: " + resp.Status)
		}
		defer resp.Body.Close()

		f := NewFile("wordlist")
		f.HeaderComment("// Code generated by internal/wordlist/generate; DO NOT EDIT.")

		f.Comment(config.name + " " + config.description + "\n\n" + "Source: https://www.eff.org/dice")
		f.Var().Id(config.name).Op("=").Id("Wordlist").BlockFunc(func(group *Group) {
			scanner := bufio.NewScanner(resp.Body)
			for scanner.Scan() {
				if scanner.Bytes()[0] != '#' {
					split := strings.Split(scanner.Text(), "\t")
					group.Lit(split[1]).Op(",")
				}
			}
		})

		if err := f.Save(strings.ToLower(config.name) + ".go"); err != nil {
			panic(err)
		}
	}
}
