package comp

import "github.com/zrcoder/amisgo/schema"

// Option represents an option schema.Schema
type Option schema.Schema

func NewOption() Option {
	return Option{}
}

// Set sets a field value
func (o Option) Set(key string, value any) Option {
	o[key] = value
	return o
}

// Children supports nested options
func (o Option) Children(value ...Option) Option {
	return o.Set("children", value)
}

// Defer marks data for deferred loading
func (o Option) Defer(value bool) Option {
	return o.Set("defer", value)
}

// DeferApi sets a higher priority API for deferred loading
func (o Option) DeferApi(value string) Option {
	return o.Set("deferApi", value)
}

// Description sets a description
func (o Option) Description(value string) Option {
	return o.Set("description", value)
}

// Disabled sets the disabled state
func (o Option) Disabled(value bool) Option {
	return o.Set("disabled", value)
}

// Hidden sets the hidden state (use visible instead)
func (o Option) Hidden(value bool) Option {
	return o.Set("hidden", value)
}

// Label sets the display text
func (o Option) Label(value string) Option {
	return o.Set("label", value)
}

// Loaded indicates if deferred loading is complete (internal use only)
func (o Option) Loaded(value bool) Option {
	return o.Set("loaded", value)
}

// Loading marks as loading (internal use only)
func (o Option) Loading(value bool) Option {
	return o.Set("loading", value)
}

// ScopeLabel sets a scope label for clearer data display
func (o Option) ScopeLabel(value string) Option {
	return o.Set("scopeLabel", value)
}

// Value sets a unique value for the option
func (o Option) Value(value any) Option {
	return o.Set("value", value)
}

// Visible sets the visibility state
func (o Option) Visible(value bool) Option {
	return o.Set("visible", value)
}

// Image sets the image path
func (o Option) Image(value string) Option {
	return o.Set("image", value)
}

// Body sets the body content
func (o Option) Body(value any) Option {
	return o.Set("body", value)
}
