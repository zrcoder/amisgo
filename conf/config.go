package conf

import (
	stdtpl "html/template"
	"net/http"

	"github.com/zrcoder/amisgo/internal/conf"
	"github.com/zrcoder/amisgo/internal/template"
)

// Available confs
const (
	ThemeDefault = conf.ThemeDefault
	ThemeCxd     = conf.ThemeCxd
	ThemeAntd    = conf.ThemeAntd
	ThemeAng     = conf.ThemeAng
	ThemeDark    = conf.ThemeDark

	LocaleDefault = conf.LocaleDefault
	LocaleZhCN    = conf.LocalZhCN
	LocaleEnUS    = conf.LocaleEnUS
)

type (
	Theme  = conf.Theme
	Locale = conf.Local
)

// Config holds all configuration options for the application
type Config struct {
	Title      string
	Icon       string
	CustomCSS  stdtpl.CSS
	CustomJS   stdtpl.JS
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
	return WithThemes(Theme{Value: name, Label: name})
}

// WithThemes sets multiple UI themes, overriding an set theme via WithTheme.
func WithThemes(themes ...Theme) Option {
	return func(c *Config) {
		if len(themes) == 0 {
			panic("WithThemes: at least 1 theme is required")
		}
		conf.RegularThemes(themes)
		c.Templ.Themes = themes
	}
}

// WithLocale sets the locale.
func WithLocale(locale string) Option {
	return WithLocales(Locale{Value: locale, Label: locale})
}

// WithLocales sets multiple locales, overriding an set locale via WithLocale.
func WithLocales(locales ...Locale) Option {
	return func(c *Config) {
		if len(locales) == 0 {
			panic("WithLocales: at least 1 lang is required")
		}
		conf.RegularLocales(locales)
		c.Templ.Locales = locales
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
		c.CustomCSS = stdtpl.CSS(customCSS)
	}
}

// WithCustomJS sets the custom JavaScript URL.
func WithCustomJS(customJS string) Option {
	return func(c *Config) {
		c.CustomJS = stdtpl.JS(customJS)
	}
}

// WithLocalSdk sets the local SDK file system
func WithLocalSdk(fs http.FileSystem) Option {
	return func(c *Config) {
		c.LocalSdkFS = fs
	}
}
