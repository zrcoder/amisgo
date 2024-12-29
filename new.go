//go:build !amisgo_use_local_sdk

package amisgo

import (
	"github.com/zrcoder/amisgo/conf"
	"github.com/zrcoder/amisgo/internal/servermux"
)

// New creates an Engine instance with options
func New(opts ...conf.Option) *Engine {
	cfg := conf.Default()
	cfg.Apply(opts...)

	return &Engine{
		Config: cfg,
		mux:    servermux.Mux(),
	}
}
