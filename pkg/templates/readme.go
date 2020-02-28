package templates

import (
	"text/template"
)

const text = `# Go-Go-Go Is a collection of sources I find, as well as a way of keeping it visible and categorized

Over the time I found starring repos is very sub-optimal when it comes to discovery, and so this project was born. A single file with the list of repositories to be tracked is used to generate the output. This small tool can be added as part of a Drone pipeline and generate your readme.

## Index 🔎
{{range $key, $value := . }}
- [{{$key}}](#{{kebap $key}})
{{end}}
---
{{range $key, $value := . }}
## {{ $key }}
{{ range $value }}
- [{{ .Owner }}/{{ .Name }}]({{ .URL }}) - {{ .Description }}
  - {{ .Stargazers }}⭐
{{end}}{{end}}`

// Readme returns a readme template ready to be executed
func Readme() *template.Template {
	return template.Must(template.New("readme").Funcs(fnMap).Parse(text))
}
