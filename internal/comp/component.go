package comp

import "github.com/zrcoder/amisgo/schema"

// Component represents a flexible Component with configurable attributes
type Component schema.Schema

func NewComponent() Component {
	return make(Component)
}

// Set allows setting additional properties for the component
func (c Component) Set(key string, value any) Component {
	c[key] = value
	return c
}
