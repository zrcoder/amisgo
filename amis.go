// Package amisgo provides a web framework for building Amis applications
package amisgo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/zrcoder/amisgo/config"
	"github.com/zrcoder/amisgo/internal/servermux"
)

// Engine represents the web application
type Engine struct {
	Config      *config.Config
	mux         *http.ServeMux
	middlewares []func(http.Handler) http.Handler
	handler     http.Handler
	handlerOnce sync.Once
}

// New creates an Engine instance with options
func New(opts ...config.Option) *Engine {
	cfg := config.Default()
	cfg.Apply(opts...)

	e := &Engine{
		Config: cfg,
		mux:    servermux.Mux(),
	}

	if cfg.StaticFS != nil {
		e.mux.Handle(cfg.StaticPrefix,
			http.StripPrefix(cfg.StaticPrefix, http.FileServer(cfg.StaticFS)))
	}

	return e
}

// Mount registers an Amis component at the given path
func (e *Engine) Mount(path string, component any) *Engine {
	e.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		e.renderComponent(w, component)
	})
	return e
}

// Redirect sets up a URL redirection
func (e *Engine) Redirect(src, dst string) *Engine {
	e.mux.HandleFunc(src, func(w http.ResponseWriter, r *http.Request) {
		Redirect(w, r, dst, http.StatusPermanentRedirect)
	})
	return e
}

// Handle registers a handler at the given path
func (e *Engine) Handle(path string, handler http.Handler) *Engine {
	e.mux.Handle(path, handler)
	return e
}

// HandleFunc registers a handler function at the given path
func (e *Engine) HandleFunc(path string, handler func(http.ResponseWriter, *http.Request)) *Engine {
	e.mux.HandleFunc(path, handler)
	return e
}

// Use adds a middleware function to the engine for processing requests
func (e *Engine) Use(middleware func(http.Handler) http.Handler) *Engine {
	e.middlewares = append(e.middlewares, middleware)
	return e
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
	e.handlerOnce.Do(func() {
		e.handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			e.mux.ServeHTTP(w, r)
		})
		for _, m := range e.middlewares {
			e.handler = m(e.handler)
		}
	})

	e.handler.ServeHTTP(w, r)
}

// renderComponent renders an Amis component
func (e *Engine) renderComponent(w io.Writer, component any) {
	amisJson, err := json.Marshal(component)
	if err != nil {
		http.Error(w.(http.ResponseWriter), "failed to marshal component:"+err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		*config.Config
		AmisJson string
	}{
		Config:   e.Config,
		AmisJson: string(amisJson),
	}
	if err = e.Config.AmisTemplate.Execute(w, data); err != nil {
		panic(fmt.Sprintf("failed to execute template: %v", err))
	}
}
