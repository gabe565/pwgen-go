package template

import (
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/gabe565/pwgen-go/internal/rand"
	"github.com/gabe565/pwgen-go/internal/wordlist"
)

func FuncMap(list wordlist.Wordlist) template.FuncMap {
	funcs := sprig.FuncMap()

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

	funcs["randBinary"] = rand.BinaryN
	funcs["binary"] = rand.BinaryN

	funcs["shuffle"] = rand.ShuffleSlice[any]

	return funcs
}
