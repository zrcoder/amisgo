package config

import (
	"net/http"
	"strings"

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
	Theme        Theme
	Lang         Lang
	Title        string
	Icon         string
	CustomCSS    string
	CustomJS     string
	StaticPrefix string
	StaticFS     http.FileSystem
	template.Template
}

// Default returns a new Config instance with default settings
func Default() *Config {
	return &Config{
		Title:    "Amisgo",
		Lang:     LangDefault,
		Theme:    ThemeDefault,
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
		c.Theme = theme
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

// WithStaticFS sets up static file system serving with the given URL prefix.
// If urlPrefix is empty or "/", it serves files at the root path.
// The urlPrefix will be normalized to have a leading and trailing slash.
func WithStaticFS(urlPrefix string, staticFS http.FileSystem) Option {
	if staticFS == nil {
		panic("staticFS cannot be nil")
	}
	if urlPrefix == "" || urlPrefix == "/" {
		return func(c *Config) {
			c.StaticPrefix = "/"
			c.StaticFS = staticFS
		}
	}
	urlPrefix = "/" + strings.TrimLeft(urlPrefix, "/")
	urlPrefix = strings.TrimRight(urlPrefix, "/") + "/"
	return func(c *Config) {
		c.StaticPrefix = urlPrefix
		c.StaticFS = staticFS
	}
}

// WithStatic is a convenience function that wraps WithStaticFS using a local directory.
// It panics if the path is empty.
func WithStatic(urlPrefix, path string) Option {
	if path == "" {
		panic("path cannot be empty")
	}
	return WithStaticFS(urlPrefix, http.Dir(path))
}
