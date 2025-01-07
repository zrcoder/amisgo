package comp

// MOption represents an option schema
type MOption Schema

// Option creates a new MOption instance
func Option() MOption {
	return MOption{}
}

// Set sets a field value
func (o MOption) Set(key string, value any) MOption {
	o[key] = value
	return o
}

// Children supports nested options
func (o MOption) Children(value ...MOption) MOption {
	return o.Set("children", value)
}

// Defer marks data for deferred loading
func (o MOption) Defer(value bool) MOption {
	return o.Set("defer", value)
}

// DeferApi sets a higher priority API for deferred loading
func (o MOption) DeferApi(value string) MOption {
	return o.Set("deferApi", value)
}

// Description sets a description
func (o MOption) Description(value string) MOption {
	return o.Set("description", value)
}

// Disabled sets the disabled state
func (o MOption) Disabled(value bool) MOption {
	return o.Set("disabled", value)
}

// Hidden sets the hidden state (use visible instead)
func (o MOption) Hidden(value bool) MOption {
	return o.Set("hidden", value)
}

// Label sets the display text
func (o MOption) Label(value string) MOption {
	return o.Set("label", value)
}

// Loaded indicates if deferred loading is complete (internal use only)
func (o MOption) Loaded(value bool) MOption {
	return o.Set("loaded", value)
}

// Loading marks as loading (internal use only)
func (o MOption) Loading(value bool) MOption {
	return o.Set("loading", value)
}

// ScopeLabel sets a scope label for clearer data display
func (o MOption) ScopeLabel(value string) MOption {
	return o.Set("scopeLabel", value)
}

// Value sets a unique value for the option
func (o MOption) Value(value any) MOption {
	return o.Set("value", value)
}

// Visible sets the visibility state
func (o MOption) Visible(value bool) MOption {
	return o.Set("visible", value)
}

// Image sets the image path
func (o MOption) Image(value string) MOption {
	return o.Set("image", value)
}

// Body sets the body content
func (o MOption) Body(value any) MOption {
	return o.Set("body", value)
}
