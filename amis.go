package amisgo

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"text/template"
)

// Global template for HTML rendering
var tmpl = template.Must(template.New("").Parse(htmlTemplate))

// Engine represents the core web application structure
type Engine struct {
	routes map[string]any
	Config *Config
}

// New creates and initializes a new Engine instance
func New() *Engine {
	return &Engine{
		routes: make(map[string]any),
		Config: GetDefaultConfig(),
	}
}

// Register associates a component with a URL path
func (e *Engine) Register(path string, component any) {
	e.routes[path] = component
}

// Redirect sets up a URL redirection from source to destination path
func (e *Engine) Redirect(src, dst string) {
	http.HandleFunc(src, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, dst, http.StatusTemporaryRedirect)
	})
}

// Run starts the HTTP server with the specified address
func (e *Engine) Run(addr string) error {
	// Setup static file server if configured
	if e.Config.StaticDir != "" {
		dir := strings.Trim(e.Config.StaticDir, "/")
		e.Config.StaticDir = "/" + dir + "/"
		if e.Config.StaticFS == nil {
			http.Handle(e.Config.StaticDir, http.StripPrefix(e.Config.StaticDir, http.FileServer(http.Dir(dir))))
		} else {
			http.Handle(e.Config.StaticDir, http.FileServer(http.FS(e.Config.StaticFS)))
		}
	}

	// Register all routes to DefaultServeMux
	for path, component := range e.routes {
		component := component // Create new variable for closure
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			e.serveComponent(w, component)
		})
	}

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
