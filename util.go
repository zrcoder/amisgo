package amisgo

import "net/http"

// Redirect is a utility function to perform an HTTP redirect.
func Redirect(w http.ResponseWriter, r *http.Request, url string, code int) {
	http.Redirect(w, r, url, code)
}
