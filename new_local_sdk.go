//go:build amisgo_local_sdk

package amisgo

import (
	"net/http"

	"github.com/zrcoder/amisgo/conf"
	"github.com/zrcoder/amisgo/internal/servermux"
	"github.com/zrcoder/amisgo/jssdk"
)

// New creates an Engine instance with options
func New(opts ...conf.Option) *Engine {
	cfg := conf.Default()
	cfg.Apply(opts...)
	cfg.UseLocalSDK = true

	e := &Engine{
		Config: cfg,
		mux:    servermux.Mux(),
	}

	e.StaticFS("/__amisgo__sdk/", http.FS(jssdk.FS))

	return e
}
