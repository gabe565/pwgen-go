package template

import (
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/gabe565/pwgen-go/internal/rand"
	"github.com/gabe565/pwgen-go/internal/wordlist"
)

func FuncMap() template.FuncMap {
	funcs := sprig.FuncMap()

	funcs["randWord"] = wordlist.EFF_Long.RandWord
	funcs["word"] = wordlist.EFF_Long.RandWord

	funcs["randWords"] = wordlist.EFF_Long.RandWords
	funcs["words"] = wordlist.EFF_Long.RandWords
	funcs["wordsWithNumber"] = wordlist.EFF_Long.RandWordsWithNumber
	funcs["wordsWithNum"] = wordlist.EFF_Long.RandWordsWithNumber

	funcs["number"] = funcs["randNumeric"]
	funcs["num"] = funcs["randNumeric"]

	funcs["shuffle"] = rand.ShuffleSlice[any]

	return funcs
}
