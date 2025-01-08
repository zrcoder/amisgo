package comp

import "github.com/zrcoder/amisgo/model"

// Component represents a flexible component with configurable attributes
type component model.Schema

// Component creates a new Component instance
func Component() component {
	return make(component)
}

// Set allows setting additional properties for the component
func (c component) Set(key string, value any) component {
	c[key] = value
	return c
}
