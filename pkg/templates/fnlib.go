package templates

import (
	"strings"
	"text/template"
)

var (
	fnMap = template.FuncMap{
		"kebap": func(s string) string {
			return strings.ToLower(strings.ReplaceAll(s, " ", "-"))
		},
	}
)
