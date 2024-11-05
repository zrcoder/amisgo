package amisgo

import (
	"encoding/json"
	"io"
	"net/http"
	"text/template"

	"github.com/zrcoder/amisgo/config"
)

// Global template for HTML rendering
var tmpl = template.Must(template.New("").Parse(htmlTemplate))

// Engine represents the core web application structure
type Engine struct {
	Config *config.Config
}

// New creates and initializes a new Engine instance
func New(opts ...config.Option) *Engine {
	cfg := config.GetDefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	return &Engine{
		Config: cfg,
	}
}

// Register associates a component with a URL path
func (e *Engine) Register(path string, component any) *Engine {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		e.serveComponent(w, component)
	})
	return e
}

// Redirect sets up a URL redirection from source to destination path
func (e *Engine) Redirect(src, dst string) *Engine {
	http.HandleFunc(src, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, dst, http.StatusTemporaryRedirect)
	})
	return e
}

// Run starts the HTTP server with the specified address
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, nil)
}

// serveComponent handles the rendering of registered components
func (e *Engine) serveComponent(w io.Writer, component any) {
	amisJson, _ := json.Marshal(component)
	data := struct {
		*config.Config
		AmisJson string
	}{
		Config:   e.Config,
		AmisJson: string(amisJson),
	}
	if err := tmpl.Execute(w, data); err != nil {
		panic(err)
	}
}
