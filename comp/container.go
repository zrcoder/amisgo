package comp

// Container represents a container renderer with multiple configuration options
type container Schema

// Container creates a new Container instance
func Container() container {
	return make(container).set("type", "container")
}

func (c container) set(key string, value any) container {
	c[key] = value
	return c
}

// Body sets the content for the container
func (c container) Body(value ...any) container {
	return c.set("body", value)
}

// BodyClassName sets the CSS class name for the body
func (c container) BodyClassName(value string) container {
	return c.set("bodyClassName", value)
}

// ClassName sets the CSS class name for the container
func (c container) ClassName(value string) container {
	return c.set("className", value)
}

// Disabled enables or disables the entire component
func (c container) Disabled(value bool) container {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the component
func (c container) DisabledOn(value string) container {
	return c.set("disabledOn", value)
}

// Draggable enables or disables container drag functionality
func (c container) Draggable(value bool) container {
	return c.set("draggable", value)
}

// DraggableConfig sets the configuration for container dragging
func (c container) DraggableConfig(value string) container {
	return c.set("draggableConfig", value)
}

// EditorSetting configures editor-specific settings
func (c container) EditorSetting(value string) container {
	return c.set("editorSetting", value)
}

// Hidden controls the overall visibility of the component
func (c container) Hidden(value bool) container {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the component
func (c container) HiddenOn(value string) container {
	return c.set("hiddenOn", value)
}

// ID sets a unique identifier for the component
func (c container) ID(value string) container {
	return c.set("id", value)
}

// OnEvent configures event-driven actions for the component
func (c container) OnEvent(value any) container {
	return c.set("onEvent", value)
}

// Static enables or disables static display mode
func (c container) Static(value bool) container {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display
func (c container) StaticClassName(value string) container {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display
func (c container) StaticInputClassName(value string) container {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c container) StaticLabelClassName(value string) container {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c container) StaticOn(value string) container {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c container) StaticPlaceholder(value string) container {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (c container) StaticSchema(value string) container {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c container) Style(value any) container {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c container) TestIdBuilder(value string) container {
	return c.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (c container) Testid(value string) container {
	return c.set("testid", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c container) UseMobileUI(value bool) container {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the component
func (c container) Visible(value bool) container {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component
func (c container) VisibleOn(value string) container {
	return c.set("visibleOn", value)
}

// WrapperBody determines whether to wrap the body content in an additional div
func (c container) WrapperBody(value bool) container {
	return c.set("wrapperBody", value)
}

// WrapperComponent sets the HTML tag used for wrapping the container
func (c container) WrapperComponent(value string) container {
	return c.set("wrapperComponent", value)
}
