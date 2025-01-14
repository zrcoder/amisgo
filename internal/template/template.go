package template

import (
	_ "embed"
	"html/template"
)

const (
	amisVersion = "6.10.0"
	amisBaseURL = "https://cdn.jsdelivr.net/npm/amis"
)

//go:embed template.gohtml
var amisTemplate string

var (
	theme  = ""
	themes []string
)

func SetTheme(t string) {
	theme = t
}

func GetTheme() string {
	return theme
}

func SetThemes(ts []string) {
	themes = ts
	if len(ts) > 0 {
		theme = ts[0]
	}
}

func GetThemes() []string {
	return themes
}

// Global template for HTML rendering.
// This template will be used to render all Amis components.
// var tmpl = template.Must(template.New("").Parse(amisTemplate))

type Template struct {
	AmisTemplate *template.Template
	AmisVersion  string
	AmisBaseURL  string
}

func GetTemplate() Template {
	funcs := template.FuncMap{
		"GetTheme": func() string {
			return theme
		},
	}
	tmpl := template.New("").Funcs(funcs)
	tmpl = template.Must(tmpl.Parse(amisTemplate))
	return Template{
		AmisTemplate: tmpl,
		AmisVersion:  amisVersion,
		AmisBaseURL:  amisBaseURL,
	}
}
