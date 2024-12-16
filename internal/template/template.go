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

// Global template for HTML rendering.
// This template will be used to render all Amis components.
var tmpl = template.Must(template.New("").Parse(amisTemplate))

type Template struct {
	AmisTemplate *template.Template
	AmisVersion  string
	AmisBaseURL  string
}

func GetTemplate() Template {
	return Template{
		AmisTemplate: tmpl,
		AmisVersion:  amisVersion,
		AmisBaseURL:  amisBaseURL,
	}
}
