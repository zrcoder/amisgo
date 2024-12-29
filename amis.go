// Package amisgo provides a web framework for building Amis applications
package amisgo

import (
	"net/http"
	"strings"

	"github.com/zrcoder/amisgo/conf"
	"github.com/zrcoder/amisgo/util"
)

// Engine represents the web application
type Engine struct {
	Config *conf.Config
	mux    *http.ServeMux
}

// Mount registers an Amis component at the given path
func (e *Engine) Mount(path string, component any, middlewares ...func(http.Handler) http.Handler) *Engine {
	var h http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		e.renderComponent(w, component)
	})
	h = util.WithMiddlewares(h, middlewares...)
	e.mux.Handle(path, h)
	return e
}

// Redirect sets up a URL redirection
func (e *Engine) Redirect(src, dst string, code int) *Engine {
	e.mux.HandleFunc(src, func(w http.ResponseWriter, r *http.Request) {
		util.Redirect(w, r, dst, code)
	})
	return e
}

// Handle registers a handler at the given path
func (e *Engine) Handle(path string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) *Engine {
	e.mux.Handle(path, util.WithMiddlewares(handler, middlewares...))
	return e
}

// HandleFunc registers a handler function at the given path
func (e *Engine) HandleFunc(path string, handler http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) *Engine {
	return e.Handle(path, handler, middlewares...)
}

// StaticFS registers a file server for serving static files
func (e *Engine) StaticFS(prefix string, fs http.FileSystem) *Engine {
	if prefix == "" || prefix == "/" {
		prefix = "/"
	}
	prefix = "/" + strings.TrimLeft(prefix, "/")
	prefix = strings.TrimRight(prefix, "/") + "/"
	e.mux.Handle(prefix, http.StripPrefix(prefix, http.FileServer(fs)))
	return e
}

// StaticFiles registers a file server for serving static files
func (e *Engine) StaticFiles(prefix string, root string) *Engine {
	return e.StaticFS(prefix, http.Dir(root))
}

// Run starts the HTTP server
func (e *Engine) Run(addr ...string) error {
	address := ":80"
	if len(addr) > 0 && addr[0] != "" {
		address = addr[0]
	}

	return http.ListenAndServe(address, e)
}

// ServeHTTP implements http.Handler interface
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.mux.ServeHTTP(w, r)
}

// renderComponent renders an Amis component
func (e *Engine) renderComponent(w http.ResponseWriter, component any) {
	data := struct {
		*conf.Config
		AmisJson any
	}{
		Config:   e.Config,
		AmisJson: component,
	}
	if err := e.Config.AmisTemplate.Execute(w, data); err != nil {
		panic(err)
	}
}
