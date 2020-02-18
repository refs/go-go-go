package templates

import "text/template"

const text = `
## Go-Go-Go Is a collection of sources I find

Over the time I found starring repos is very sub-optimal when it comes to discovery, so I decided to create this small project for easy discoverability and categorization.

{{range .}}
	- [{{ .Owner }}/{{ .Name }}]({{ .URL }})
{{end}}
`

// Readme returns a readme template ready to be compiled
func Readme() *template.Template {
	return template.Must(template.New("readme").Parse(text))
}
