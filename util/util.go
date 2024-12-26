package util

import "net/http"

// Redirect is a utility function to perform an HTTP redirect.
func Redirect(w http.ResponseWriter, r *http.Request, url string, code int) {
	http.Redirect(w, r, url, code)
}

// WithMiddlewares applies a chain of middleware functions to an HTTP handler.
func WithMiddlewares(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
