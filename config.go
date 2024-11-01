package amisgo

// Theme represents the UI theme of the application
type Theme string

// Available themes for the application
const (
	ThemeDefault Theme = ""
	ThemeCxd     Theme = "cxd"
	ThemeAntd    Theme = "antd"
	ThemeAng     Theme = "ang"
	ThemeDark    Theme = "dark"
)

// Lang represents the language setting of the application
type Lang string

// Supported languages
const (
	LangDefault Lang = LangZh
	LangZh      Lang = "zh"
	LangEn      Lang = "en"
)

// Config holds all configuration options for the application
type Config struct {
	Theme     Theme
	Lang      Lang
	Title     string
	Icon      string
	CustomCSS string
	CustomJS  string
	Host      string
}

// GetDefaultConfig returns a new Config instance with default settings
func GetDefaultConfig() *Config {
	return &Config{
		Title: "Amisgo",
		Lang:  LangDefault,
		Theme: ThemeDefault,
		Host:  "http://localhost",
	}
}
