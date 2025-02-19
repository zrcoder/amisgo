package comp

import "github.com/zrcoder/amisgo/schema"

// Custom represents a Custom component.
type Custom schema.Schema

func NewCustom() Custom {
	return Custom{"type": "custom"}
}

func (c Custom) set(key string, value any) Custom {
	c[key] = value
	return c
}

// ClassName sets the class of the node.
func (c Custom) ClassName(value string) Custom {
	return c.set("className", value)
}

// HTML sets the initial HTML of the node.
func (c Custom) HTML(value string) Custom {
	return c.set("html", value)
}

// ID sets the ID of the node.
func (c Custom) ID(value string) Custom {
	return c.set("id", value)
}

// Inline sets whether to use a span tag instead of a div tag.
func (c Custom) Inline(value bool) Custom {
	return c.set("inline", value)
}

// Name sets the name of the node.
func (c Custom) Name(value string) Custom {
	return c.set("name", value)
}

// OnMount sets the function to be called after the node is mounted.
func (c Custom) OnMount(value string) Custom {
	return c.set("onMount", value)
}

// OnUnmount sets the function to be called after the node is unmounted.
func (c Custom) OnUnmount(value string) Custom {
	return c.set("onUnmount", value)
}

// OnUpdate sets the function to be called when the data is updated.
func (c Custom) OnUpdate(value string) Custom {
	return c.set("onUpdate", value)
}
