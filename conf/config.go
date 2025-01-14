package conf

import (
	"net/http"

	"github.com/zrcoder/amisgo/internal/template"
)

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
	LangDefault Lang = LangEn
	LangZh      Lang = "zh"
	LangEn      Lang = "en"
)

// Config holds all configuration options for the application
type Config struct {
	Lang       Lang
	Title      string
	Icon       string
	CustomCSS  string
	CustomJS   string
	LocalSdkFS http.FileSystem
	template.Template
}

func (c *Config) UseLocalSDK() bool {
	return c.LocalSdkFS != nil
}

// Default returns a new Config instance with default settings
func Default() *Config {
	return &Config{
		Title:    "amisgo",
		Lang:     LangDefault,
		Template: template.GetTemplate(),
	}
}

// Apply applies the given options to the config.
func (c *Config) Apply(options ...Option) {
	for _, option := range options {
		option(c)
	}
}

// Option represents a function that modifies the Config.
type Option func(*Config)

// WithTheme sets the UI theme.
func WithTheme(theme Theme) Option {
	return func(c *Config) {
		if len(template.GetThemes()) > 0 {
			return
		}
		template.SetTheme(string(theme))
	}
}

// WithThemes sets multiple UI themes, overriding an set theme via WithTheme.
// Once themes are configured, you can use the comp.ThemeSelect or comp.ThemeButtonGroupSelect component in your pages to enable users to switch between the available themes.
func WithThemes(themes ...Theme) Option {
	return func(c *Config) {
		if len(themes) < 2 {
			panic("WithThemes: at least 2 themes are required")
		}
		ts := make([]string, len(themes))
		for i, v := range themes {
			ts[i] = string(v)
		}
		template.SetThemes(ts)
	}
}

// WithLang sets the interface language.
func WithLang(lang Lang) Option {
	return func(c *Config) {
		c.Lang = lang
	}
}

// WithTitle sets the page title.
func WithTitle(title string) Option {
	return func(c *Config) {
		c.Title = title
	}
}

// WithIcon sets the favicon URL.
func WithIcon(icon string) Option {
	return func(c *Config) {
		c.Icon = icon
	}
}

// WithCustomCSS sets the custom CSS URL.
func WithCustomCSS(customCSS string) Option {
	return func(c *Config) {
		c.CustomCSS = customCSS
	}
}

// WithCustomJS sets the custom JavaScript URL.
func WithCustomJS(customJS string) Option {
	return func(c *Config) {
		c.CustomJS = customJS
	}
}

// WithLocalSdk sets the local SDK file system
func WithLocalSdk(fs http.FileSystem) Option {
	return func(c *Config) {
		c.LocalSdkFS = fs
	}
}
