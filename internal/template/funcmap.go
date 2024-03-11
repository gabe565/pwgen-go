package template

import (
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/gabe565/pwgen-go/internal/wordlist"
)

func FuncMap() template.FuncMap {
	funcs := sprig.FuncMap()

	funcs["randWord"] = wordlist.EFF_Long.RandWord
	funcs["word"] = wordlist.EFF_Long.RandWord

	funcs["randWords"] = wordlist.EFF_Long.RandWords
	funcs["words"] = wordlist.EFF_Long.RandWords

	funcs["number"] = funcs["randNumeric"]
	funcs["num"] = funcs["randNumeric"]

	return funcs
}
