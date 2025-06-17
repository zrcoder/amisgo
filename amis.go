// Package amisgo provides a web framework for building amis applications
package amisgo

import (
	"context"
	"net/http"
	"strings"

	"github.com/zrcoder/amisgo/conf"
	"github.com/zrcoder/amisgo/util"
)

// App represents the web application
type App struct {
	Conf   *conf.Config
	mux    *http.ServeMux
	server *http.Server
}

// New creates an Engine instance with options
func New(opts ...conf.Option) *App {
	cfg := conf.Default()
	cfg.Apply(opts...)

	app := &App{
		Conf: cfg,
		mux:  http.NewServeMux(),
	}

	if cfg.LocalSdkFS != nil {
		app.StaticFS("/__amisgo__sdk/", cfg.LocalSdkFS)
	}

	return app
}

// Mount registers an amis component at the given path
func (a *App) Mount(path string, component any, middlewares ...func(http.Handler) http.Handler) {
	var h http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.renderComponent(w, component)
	})
	h = util.WithMiddlewares(h, middlewares...)
	a.mux.Handle(path, h)
}

// Redirect sets up a URL redirection
func (a *App) Redirect(src, dst string, code int) {
	a.mux.HandleFunc(src, func(w http.ResponseWriter, r *http.Request) {
		util.Redirect(w, r, dst, code)
	})
}

// Handle registers a handler at the given path
func (a *App) Handle(path string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) {
	a.mux.Handle(path, util.WithMiddlewares(handler, middlewares...))
}

// HandleFunc registers a handler function at the given path
func (a *App) HandleFunc(path string, handler http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) {
	a.Handle(path, handler, middlewares...)
}

// StaticFS registers a file server for serving static files
func (a *App) StaticFS(prefix string, fs http.FileSystem) {
	// regular prefix to: "/xxx/"" (eg. "abc", "/abc", "abc/", "//abc/" all will be changed to "/abc/")
	// if prefix is "" or "/", regular it to "/"
	prefix = "/" + strings.TrimLeft(prefix, "/")
	prefix = strings.TrimRight(prefix, "/") + "/"
	a.mux.Handle(prefix, http.StripPrefix(prefix, http.FileServer(fs)))
}

// StaticFiles registers a file server for serving static files
func (a *App) StaticFiles(prefix string, root string) {
	a.StaticFS(prefix, http.Dir(root))
}

// Run starts the HTTP server
func (a *App) Run(addr ...string) error {
	address := ":80"
	if len(addr) > 0 && addr[0] != "" {
		address = addr[0]
	}
	a.server = &http.Server{Addr: address, Handler: a}

	return a.server.ListenAndServe()
}

// Shutdown gracefully shuts down the server without interrupting any
// active connections.
func (a *App) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}

// ServeHTTP implements http.Handler interface
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

// renderComponent renders an amis component
func (a *App) renderComponent(w http.ResponseWriter, component any) {
	data := struct {
		*conf.Config
		AmisJson any
	}{
		Config:   a.Conf,
		AmisJson: component,
	}
	if err := a.Conf.AmisTempl.Execute(w, data); err != nil {
		panic(err)
	}
}
