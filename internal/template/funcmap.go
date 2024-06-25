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
	funcs["randNumeric"] = rand.Numeric
	funcs["number"] = rand.Numeric
	funcs["num"] = rand.Numeric
	funcs["numeric"] = rand.Numeric
	funcs["randAlpha"] = rand.Alpha
	funcs["alpha"] = rand.Alpha
	funcs["randAlphaNum"] = rand.AlphaNum
	funcs["alphaNum"] = rand.AlphaNum
	funcs["randAscii"] = rand.ASCII
	funcs["ascii"] = rand.ASCII
	funcs["randBinary"] = rand.BinaryN
	funcs["binary"] = rand.BinaryN
	funcs["shuffle"] = rand.ShuffleSlice[any]
	return funcs
}
