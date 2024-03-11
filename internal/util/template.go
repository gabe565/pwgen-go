package util

import (
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func TemplateFuncMap() template.FuncMap {
	funcs := sprig.FuncMap()

	funcs["randWord"] = RandWord
	funcs["word"] = RandWord

	funcs["randWords"] = RandWords
	funcs["words"] = RandWords

	funcs["number"] = funcs["randNumeric"]
	funcs["num"] = funcs["randNumeric"]

	return funcs
}
