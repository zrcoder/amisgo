package comp

import "github.com/zrcoder/amisgo/schema"

// Container represents a Container renderer with multiple configuration options
type Container schema.Schema

func NewContainer() Container {
	return Container{"type": "container"}
}

func (c Container) set(key string, value any) Container {
	c[key] = value
	return c
}

// Body sets the content for the container
func (c Container) Body(value ...any) Container {
	return c.set("body", value)
}

// BodyClassName sets the CSS class name for the body
func (c Container) BodyClassName(value string) Container {
	return c.set("bodyClassName", value)
}

// ClassName sets the CSS class name for the container
func (c Container) ClassName(value string) Container {
	return c.set("className", value)
}

// Disabled enables or disables the entire component
func (c Container) Disabled(value bool) Container {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the component
func (c Container) DisabledOn(value string) Container {
	return c.set("disabledOn", value)
}

// Draggable enables or disables container drag functionality
func (c Container) Draggable(value bool) Container {
	return c.set("draggable", value)
}

// DraggableConfig sets the configuration for container dragging
func (c Container) DraggableConfig(value string) Container {
	return c.set("draggableConfig", value)
}

// EditorSetting configures editor-specific settings
func (c Container) EditorSetting(value string) Container {
	return c.set("editorSetting", value)
}

// Hidden controls the overall visibility of the component
func (c Container) Hidden(value bool) Container {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the component
func (c Container) HiddenOn(value string) Container {
	return c.set("hiddenOn", value)
}

// ID sets a unique identifier for the component
func (c Container) ID(value string) Container {
	return c.set("id", value)
}

// OnEvent configures event-driven actions for the component
func (c Container) OnEvent(value any) Container {
	return c.set("onEvent", value)
}

// Static enables or disables static display mode
func (c Container) Static(value bool) Container {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display
func (c Container) StaticClassName(value string) Container {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display
func (c Container) StaticInputClassName(value string) Container {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c Container) StaticLabelClassName(value string) Container {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c Container) StaticOn(value string) Container {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c Container) StaticPlaceholder(value string) Container {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (c Container) StaticSchema(value string) Container {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c Container) Style(value any) Container {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c Container) TestIdBuilder(value string) Container {
	return c.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (c Container) Testid(value string) Container {
	return c.set("testid", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c Container) UseMobileUI(value bool) Container {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the component
func (c Container) Visible(value bool) Container {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component
func (c Container) VisibleOn(value string) Container {
	return c.set("visibleOn", value)
}

// WrapperBody determines whether to wrap the body content in an additional div
func (c Container) WrapperBody(value bool) Container {
	return c.set("wrapperBody", value)
}

// WrapperComponent sets the HTML tag used for wrapping the container
func (c Container) WrapperComponent(value string) Container {
	return c.set("wrapperComponent", value)
}
