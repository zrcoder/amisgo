package template

import (
	_ "embed"
	"html/template"

	"github.com/zrcoder/amisgo/internal/conf"
)

const (
	amisVersion = "6.12.0"
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
	Themes      []conf.Theme
	themeIndex  int
	Locales     []conf.Local
	localeIndex int
}

func New() *Templ {
	res := &Templ{
		AmisVersion: amisVersion,
		AmisBaseURL: amisBaseURL,
		Themes:      []conf.Theme{{Value: conf.ThemeDefault, Label: conf.ThemeDefault}},
		Locales:     []conf.Local{{Value: conf.LocaleDefault, Label: conf.LocaleDefault}},
	}
	funcs := template.FuncMap{
		"CurrentTheme":  res.CurrentTheme,
		"CurrentLocale": res.CurrentLocale,
		"CurrentI18n":   res.CurrentI18n,
	}
	tmpl := template.New("").Funcs(funcs)
	tmpl = template.Must(tmpl.Parse(amisTempl))
	res.AmisTempl = tmpl
	return res
}

func (t *Templ) UpdateTheme(name string) {
	for i, theme := range t.Themes {
		if theme.Value == name {
			t.themeIndex = i
			return
		}
	}
}

func (t *Templ) CurrentTheme() string {
	return t.Themes[t.themeIndex].Value
}

func (t *Templ) SetLocales(ls []conf.Local) {
	t.Locales = ls
}

func (t *Templ) GetLocales() []conf.Local {
	return t.Locales
}

func (t *Templ) UpdateLocale(name string) {
	for i, locale := range t.Locales {
		if locale.Value == name {
			t.localeIndex = i
			return
		}
	}
}

func (t *Templ) CurrentLocale() string {
	return t.Locales[t.localeIndex].Value
}

func (t *Templ) CurrentI18n() any {
	return t.Locales[t.localeIndex].Dict
}
