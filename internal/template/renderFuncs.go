package template

import (
	tpl "html/template"

	"github.com/Masterminds/sprig/v3"
)

func padLeft(s string, fill string, size int) string {
	for len(s) < size {
		s = fill + s
	}
	return s
}

func funcMap() tpl.FuncMap {
	funcs := sprig.FuncMap()
	funcs["padLeft"] = padLeft
	return funcs
}
