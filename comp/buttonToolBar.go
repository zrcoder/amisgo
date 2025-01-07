package comp

// ButtonToolbar represents a button toolbar control renderer
type buttonToolbar Schema

// ButtonToolbar creates a new ButtonToolbar instance
func ButtonToolbar() buttonToolbar {
	return make(buttonToolbar).set("type", "button-toolbar")
}

// set overrides the BaseRenderer's Set method, returning ButtonToolbar
func (bt buttonToolbar) set(key string, value any) buttonToolbar {
	bt[key] = value
	return bt
}

// Buttons configures the collection of buttons
func (bt buttonToolbar) Buttons(value ...action) buttonToolbar {
	return bt.set("buttons", value)
}

// Label sets the label for the button toolbar
func (bt buttonToolbar) Label(value string) buttonToolbar {
	return bt.set("label", value)
}

// ClassName sets the CSS class name for the container
func (bt buttonToolbar) ClassName(value string) buttonToolbar {
	return bt.set("className", value)
}

// Disabled enables or disables the button toolbar
func (bt buttonToolbar) Disabled(value bool) buttonToolbar {
	return bt.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the button toolbar
func (bt buttonToolbar) DisabledOn(value string) buttonToolbar {
	return bt.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings
func (bt buttonToolbar) EditorSetting(value string) buttonToolbar {
	return bt.set("editorSetting", value)
}

// Hidden controls the visibility of the button toolbar
func (bt buttonToolbar) Hidden(value bool) buttonToolbar {
	return bt.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the button toolbar
func (bt buttonToolbar) HiddenOn(value string) buttonToolbar {
	return bt.set("hiddenOn", value)
}

// ID sets a unique identifier for the component
func (bt buttonToolbar) ID(value string) buttonToolbar {
	return bt.set("id", value)
}

// OnEvent configures event-driven actions
func (bt buttonToolbar) OnEvent(value any) buttonToolbar {
	return bt.set("onEvent", value)
}

// Static determines if the button toolbar is statically displayed
func (bt buttonToolbar) Static(value bool) buttonToolbar {
	return bt.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (bt buttonToolbar) StaticClassName(value string) buttonToolbar {
	return bt.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (bt buttonToolbar) StaticInputClassName(value string) buttonToolbar {
	return bt.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (bt buttonToolbar) StaticLabelClassName(value string) buttonToolbar {
	return bt.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (bt buttonToolbar) StaticOn(value string) buttonToolbar {
	return bt.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (bt buttonToolbar) StaticPlaceholder(value string) buttonToolbar {
	return bt.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (bt buttonToolbar) StaticSchema(value string) buttonToolbar {
	return bt.set("staticSchema", value)
}

// Style sets custom inline styles
func (bt buttonToolbar) Style(value any) buttonToolbar {
	return bt.set("style", value)
}

// TestIdBuilder configures test ID generation
func (bt buttonToolbar) TestIdBuilder(value string) buttonToolbar {
	return bt.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (bt buttonToolbar) Testid(value string) buttonToolbar {
	return bt.set("testid", value)
}

// UseMobileUI enables or disables mobile UI styling
func (bt buttonToolbar) UseMobileUI(value bool) buttonToolbar {
	return bt.set("useMobileUI", value)
}

// Visible controls the overall visibility of the button toolbar
func (bt buttonToolbar) Visible(value bool) buttonToolbar {
	return bt.set("visible", value)
}

// VisibleOn sets a conditional expression for button toolbar visibility
func (bt buttonToolbar) VisibleOn(value string) buttonToolbar {
	return bt.set("visibleOn", value)
}
