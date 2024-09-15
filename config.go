package amisgo

import "io/fs"

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
	Theme     Theme
	Lang      Lang
	Title     string
	Icon      string
	StaticDir string
	// if you want to build you app as a single binary, pass a embed.FS to StaticFS is helpful
	StaticFS  fs.FS
	CustomCSS string
	CustomJS  string

	Host string
}

func GetDefaultConfig() *Config {
	return &Config{
		Title: "Amisgo",
		Lang:  LangDefault,
		Theme: ThemeDefault,
		Host:  "http://localhost",
	}
}

func getConfig(cfg []*Config) *Config {
	if len(cfg) > 0 {
		return cfg[0]
	}
	return GetDefaultConfig()
}
