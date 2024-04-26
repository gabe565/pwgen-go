package template

import (
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/gabe565/pwgen-go/internal/config"
	"github.com/gabe565/pwgen-go/internal/rand"
	"github.com/gabe565/pwgen-go/internal/wordlist"
)

func FuncMap(conf *config.Config) template.FuncMap {
	funcs := sprig.FuncMap()

	var list wordlist.Wordlist
	switch conf.Wordlist {
	case config.WordlistLong:
		list = wordlist.EFF_Long
	case config.WordlistShort1:
		list = wordlist.EFF_Short1
	case config.WordlistShort2:
		list = wordlist.EFF_Short2
	}

	funcs["randWord"] = list.RandWord
	funcs["word"] = list.RandWord

	funcs["randWords"] = list.RandWords
	funcs["words"] = list.RandWords
	funcs["wordsWithNumber"] = list.RandWordsWithNumber
	funcs["wordsWithNum"] = list.RandWordsWithNumber

	funcs["number"] = funcs["randNumeric"]
	funcs["num"] = funcs["randNumeric"]
	funcs["numeric"] = funcs["randNumeric"]
	funcs["alpha"] = funcs["randAlpha"]
	funcs["alphaNum"] = funcs["randAlphaNum"]
	funcs["ascii"] = funcs["randAscii"]

	funcs["shuffle"] = rand.ShuffleSlice[any]

	return funcs
}
