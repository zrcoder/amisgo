// Package amisgo provides a lightweight web framework for building web applications using Amis.
package amisgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/template"

	"github.com/zrcoder/amisgo/config"
)

// Global template for HTML rendering.
// This template will be used to render all Amis components.
var tmpl = template.Must(template.New("").Parse(htmlTemplate))

// Engine represents the core web application structure.
// It handles routing, static file serving, and component rendering.
type Engine struct {
	Config *config.Config
}

// New creates and initializes a new Engine instance with the given options.
func New(opts ...config.Option) *Engine {
	cfg := config.GetDefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	return &Engine{
		Config: cfg,
	}
}

// Mount associates an Amis component with a URL path.
// The component will be rendered as JSON and embedded in the HTML template.
func (e *Engine) Mount(path string, component any) *Engine {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		e.renderComponent(w, component)
	})
	return e
}

// HandleFunc registers a new API endpoint with the given URL path and handler function.
// This method is primarily used for handling API requests that require custom response handling.
func (e *Engine) HandleFunc(path string, handler func(http.ResponseWriter, *http.Request)) *Engine {
	http.HandleFunc(path, handler)
	return e
}

// Handle registers a new API endpoint with the given URL path and http.Handler.
// This method is primarily used for handling API requests with pre-defined http.Handler implementations.
func (e *Engine) Handle(path string, handler http.Handler) *Engine {
	http.Handle(path, handler)
	return e
}

// Redirect sets up a URL redirection from source to destination path
// using HTTP 307 Temporary Redirect.
func (e *Engine) Redirect(src, dst string) *Engine {
	http.HandleFunc(src, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, dst, http.StatusTemporaryRedirect)
	})
	return e
}

// Run starts the HTTP server with the specified address.
func (e *Engine) Run(addr string) error {
	if addr == "" {
		addr = ":80" // default port
	}
	if e.Config.StaticFS != nil {
		http.Handle(e.Config.StaticPrefix,
			http.StripPrefix(e.Config.StaticPrefix, http.FileServer(e.Config.StaticFS)))
	}
	return http.ListenAndServe(addr, http.DefaultServeMux)
}

// renderComponent handles the rendering of registered components.
// It converts the component to JSON and embeds it in the HTML template.
func (e *Engine) renderComponent(w io.Writer, component any) {
	amisJson, _ := json.Marshal(component)
	data := struct {
		*config.Config
		AmisJson string
	}{
		Config:   e.Config,
		AmisJson: string(amisJson),
	}
	if err := tmpl.Execute(w, data); err != nil {
		panic(fmt.Sprintf("failed to execute template: %v", err))
	}
}
