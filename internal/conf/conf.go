package conf

import "encoding/json"

const (
	// Available themes
	ThemeDefault = ThemeCxd
	ThemeCxd     = "cxd"
	ThemeAntd    = "antd"
	ThemeAng     = "ang"
	ThemeDark    = "dark"

	// Supported languages
	LocaleDefault = LocaleEnUS
	LocalZhCN     = "zh-CN"
	LocaleEnUS    = "en-US"
)

type Theme struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Icon  string `json:"icon"`
}

type Local struct {
	Value string          `json:"value"`
	Label string          `json:"label"`
	Dict  json.RawMessage `json:"-"`
}

func RegularThemes(themes []Theme) {
	for i := range themes {
		themes[i] = themes[i].regular()
	}
}

func RegularLocales(langs []Local) {
	for i := range langs {
		langs[i] = langs[i].regular()
	}
}

func (t Theme) regular() Theme {
	if t.Value == "" {
		t.Value = ThemeDefault
	}
	return t
}

func (l Local) regular() Local {
	if l.Value == "" {
		l.Value = LocaleDefault
	}
	if l.Label == "" {
		l.Label = l.Value
	}
	return l
}
