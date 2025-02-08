package conf

import (
	"net/http"

	"github.com/zrcoder/amisgo/internal/template"
	"github.com/zrcoder/amisgo/theme"
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
	*template.Templ
}

func (c *Config) UseLocalSDK() bool {
	return c.LocalSdkFS != nil
}

// Default returns a new Config instance with default settings
func Default() *Config {
	return &Config{
		Title: "amisgo",
		Lang:  LangDefault,
		Templ: template.New(),
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
func WithTheme(name string) Option {
	return func(c *Config) {
		if len(c.Templ.GetThemes()) > 0 {
			return
		}
		c.Templ.Theme = theme.Theme{Value: name, Label: name}
	}
}

// WithThemes sets multiple UI themes, overriding an set theme via WithTheme.
// Once themes are configured, you can use the comp.ThemeSelect or comp.ThemeButtonGroupSelect component in your pages to enable users to switch between the available themes.
func WithThemes(themes ...theme.Theme) Option {
	return func(c *Config) {
		if len(themes) < 2 {
			panic("WithThemes: at least 2 themes are required")
		}
		regularThemes(themes)
		c.Templ.SetThemes(themes)
	}
}

func regularThemes(themes []theme.Theme) {
	for i := range themes {
		if themes[i].Value == "" {
			themes[i].Value = theme.Default
		}
		if themes[i].Label == "" {
			themes[i].Label = themes[i].Value
		}
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
