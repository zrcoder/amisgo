package config

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

type Option func(*Config)

func WithTheme(theme Theme) Option {
	return func(c *Config) {
		c.Theme = theme
	}
}

func WithLang(lang Lang) Option {
	return func(c *Config) {
		c.Lang = lang
	}
}

func WithTitle(title string) Option {
	return func(c *Config) {
		c.Title = title
	}
}

func WithIcon(icon string) Option {
	return func(c *Config) {
		c.Icon = icon
	}
}

func WithCustomCSS(customCSS string) Option {
	return func(c *Config) {
		c.CustomCSS = customCSS
	}
}

func WithCustomJS(customJS string) Option {
	return func(c *Config) {
		c.CustomJS = customJS
	}
}

func WithHost(host string) Option {
	return func(c *Config) {
		c.Host = host
	}
}

func Apply(options ...Option) *Config {
	cfg := GetDefaultConfig()
	for _, option := range options {
		option(cfg)
	}
	return cfg
}
