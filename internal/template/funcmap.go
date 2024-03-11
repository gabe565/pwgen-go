package template

import (
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/gabe565/pwgen-go/internal/wordlists"
)

func FuncMap() template.FuncMap {
	funcs := sprig.FuncMap()

	funcs["randWord"] = wordlists.EFF_Long.RandWord
	funcs["word"] = wordlists.EFF_Long.RandWord

	funcs["randWords"] = wordlists.EFF_Long.RandWords
	funcs["words"] = wordlists.EFF_Long.RandWords

	funcs["number"] = funcs["randNumeric"]
	funcs["num"] = funcs["randNumeric"]

	return funcs
}
