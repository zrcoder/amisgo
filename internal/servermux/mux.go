package servermux

import (
	"net/http"
	"sync"
)

var (
	mux  *http.ServeMux
	once sync.Once
)

func Mux() *http.ServeMux {
	once.Do(func() {
		mux = http.NewServeMux()
	})
	return mux
}
