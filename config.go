package amisgo

type Theme string

const (
	ThemeDefault Theme = ""
	ThemeCxd     Theme = "cxd"
	ThemeAntd    Theme = "antd"
	ThemeAng     Theme = "ang"
	ThemeDark    Theme = "dark"
)

type Lang string

const (
	LangDefault Lang = LangZh
	LangZh      Lang = "zh"
	LangEn      Lang = "en"
)

type Config struct {
	Theme Theme
	Lang  Lang
	Title string
}

func GetDefaultConfig() *Config {
	return &Config{
		Title: "Amisgo",
		Lang:  LangDefault,
		Theme: ThemeDefault,
	}
}

func getConfig(cfg []*Config) *Config {
	if len(cfg) > 0 {
		return cfg[0]
	}
	return GetDefaultConfig()
}
