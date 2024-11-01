package amisgo

import (
	"encoding/json"
	"io"
	"net/http"
	"text/template"
)

// Global template for HTML rendering
var tmpl = template.Must(template.New("").Parse(htmlTemplate))

// Engine represents the core web application structure
type Engine struct {
	Config *Config
}

// New creates and initializes a new Engine instance
func New() *Engine {
	return &Engine{
		Config: GetDefaultConfig(),
	}
}

// Register associates a component with a URL path
func (e *Engine) Register(path string, component any) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		e.serveComponent(w, component)
	})
}

// Redirect sets up a URL redirection from source to destination path
func (e *Engine) Redirect(src, dst string) {
	http.HandleFunc(src, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, dst, http.StatusTemporaryRedirect)
	})
}

// Run starts the HTTP server with the specified address
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, nil)
}

// serveComponent handles the rendering of registered components
func (e *Engine) serveComponent(w io.Writer, component any) {
	amisJson, _ := json.Marshal(component)
	data := struct {
		*Config
		AmisJson string
	}{
		Config:   e.Config,
		AmisJson: string(amisJson),
	}
	if err := tmpl.Execute(w, data); err != nil {
		panic(err)
	}
}

// For backward compatibility
var defaultEngine = New()

// Deprecated: Use Engine.Register instead
func Serve(path string, component any) {
	defaultEngine.Register(path, component)
}

// Deprecated: Use Engine.Redirect instead
func Redirect(src, dst string) {
	defaultEngine.Redirect(src, dst)
}

// Deprecated: Use Engine.Run instead
func ListenAndServe(addr string, cfg ...*Config) error {
	if len(cfg) > 0 {
		defaultEngine.Config = cfg[0]
	}
	return defaultEngine.Run(addr)
}
