package conf

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

type Theme = option

type Local = option

type option struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Dict  any    `json:"-"`
}

func RegularThemes(themes []Theme) {
	for i := range themes {
		themes[i] = themes[i].regular(ThemeDefault)
	}
}

func RegularLocales(langs []Local) {
	for i := range langs {
		langs[i] = langs[i].regular(LocaleDefault)
	}
}

func (o option) regular(placeholder string) option {
	if o.Value == "" {
		o.Value = placeholder
	}
	if o.Label == "" {
		o.Label = o.Value
	}
	return o
}
