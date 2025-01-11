package comp

import "github.com/zrcoder/amisgo/model"

// custom represents a custom component.
type custom model.Schema

// Custom creates a new instance of custom and sets the default type.
func Custom() custom {
	return custom{"type": "custom"}
}

func (c custom) set(key string, value any) custom {
	c[key] = value
	return c
}

// ClassName sets the class of the node.
func (c custom) ClassName(value string) custom {
	return c.set("className", value)
}

// HTML sets the initial HTML of the node.
func (c custom) HTML(value string) custom {
	return c.set("html", value)
}

// ID sets the ID of the node.
func (c custom) ID(value string) custom {
	return c.set("id", value)
}

// Inline sets whether to use a span tag instead of a div tag.
func (c custom) Inline(value bool) custom {
	return c.set("inline", value)
}

// Name sets the name of the node.
func (c custom) Name(value string) custom {
	return c.set("name", value)
}

// OnMount sets the function to be called after the node is mounted.
func (c custom) OnMount(value string) custom {
	return c.set("onMount", value)
}

// OnUnmount sets the function to be called after the node is unmounted.
func (c custom) OnUnmount(value string) custom {
	return c.set("onUnmount", value)
}

// OnUpdate sets the function to be called when the data is updated.
func (c custom) OnUpdate(value string) custom {
	return c.set("onUpdate", value)
}
