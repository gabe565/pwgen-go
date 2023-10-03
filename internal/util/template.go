package util

import (
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func TemplateFuncMap() template.FuncMap {
	funcs := sprig.FuncMap()
	funcs["randWord"] = RandWord
	funcs["randWords"] = RandWords
	return funcs
}
