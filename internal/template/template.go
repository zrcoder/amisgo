package template

import (
	_ "embed"
	"html/template"

	"github.com/zrcoder/amisgo/theme"
)

const (
	amisVersion = "6.11.0"
	amisBaseURL = "https://cdn.jsdelivr.net/npm/amis"
)

//go:embed template.gohtml
var amisTempl string

type Templ struct {
	// Global template for HTML rendering.
	// This template will be used to render all Amis components.
	AmisTempl   *template.Template
	AmisVersion string
	AmisBaseURL string
	Theme       theme.Theme
	themes      []theme.Theme
}

func New() *Templ {
	res := &Templ{
		AmisVersion: amisVersion,
		AmisBaseURL: amisBaseURL,
	}
	funcs := template.FuncMap{
		"GetTheme": func() string {
			return res.Theme.Value
		},
	}
	tmpl := template.New("").Funcs(funcs)
	tmpl = template.Must(tmpl.Parse(amisTempl))
	res.AmisTempl = tmpl
	return res
}

func (t *Templ) SetThemes(ts []theme.Theme) {
	t.themes = ts
	if len(ts) > 0 {
		t.Theme = ts[0]
	}
}

func (t *Templ) GetThemes() []theme.Theme {
	return t.themes
}

func (t *Templ) GetTheme(name string) theme.Theme {
	for _, theme := range t.themes {
		if theme.Value == name {
			return theme
		}
	}
	return theme.Theme{}
}
