package amisgo

import (
	"github.com/zrcoder/amisgo/conf"
	"github.com/zrcoder/amisgo/internal/servermux"
)

// New creates an Engine instance with options
func New(opts ...conf.Option) *Engine {
	cfg := conf.Default()
	cfg.Apply(opts...)

	e := &Engine{
		Config: cfg,
		mux:    servermux.Mux(),
	}

	if cfg.LocalSdkFS != nil {
		e.StaticFS("/__amisgo__sdk/", cfg.LocalSdkFS)
	}

	return e
}
