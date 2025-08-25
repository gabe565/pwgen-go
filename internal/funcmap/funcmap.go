package funcmap

import (
	"text/template"

	"gabe565.com/pwgen/internal/rand"
	"gabe565.com/pwgen/internal/wordlist"
	"github.com/Masterminds/sprig/v3"
)

func New(list wordlist.Wordlist) template.FuncMap {
	funcs := sprig.TxtFuncMap()
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
	funcs["randBytes"] = rand.BytesN
	funcs["randBinary"] = rand.BinaryN
	funcs["binary"] = rand.BinaryN
	funcs["shuffle"] = rand.ShuffleSlice[any]
	funcs["randFromStr"] = rand.FromString
	return funcs
}
