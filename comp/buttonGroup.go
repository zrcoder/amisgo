package comp

// ButtonGroup represents a button group renderer
type buttonGroup Schema

// ButtonGroup creates a new ButtonGroup instance
func ButtonGroup() buttonGroup {
	return make(buttonGroup).set("type", "button-group")
}

func (br buttonGroup) set(key string, value any) buttonGroup {
	br[key] = value
	return br
}

// BtnActiveClassName sets the CSS class name for button active state
func (bg buttonGroup) BtnActiveClassName(value string) buttonGroup {
	return bg.set("btnActiveClassName", value)
}

// BtnActiveLevel sets the style level for selected buttons
// Supported values: 'link' | 'primary' | 'secondary' | 'info' | 'success' | 'warning' | 'danger' | 'light' | 'dark' | 'default'
func (bg buttonGroup) BtnActiveLevel(value string) buttonGroup {
	return bg.set("btnActiveLevel", value)
}

// BtnClassName sets the CSS class name for buttons
func (bg buttonGroup) BtnClassName(value string) buttonGroup {
	return bg.set("btnClassName", value)
}

// BtnLevel sets the style level for buttons
// Supported values: 'link' | 'primary' | 'secondary' | 'info' | 'success' | 'warning' | 'danger' | 'light' | 'dark' | 'default'
func (bg buttonGroup) BtnLevel(value string) buttonGroup {
	return bg.set("btnLevel", value)
}

// Buttons configures the collection of buttons
func (bg buttonGroup) Buttons(value ...any) buttonGroup {
	return bg.set("buttons", value)
}

// ClassName sets the CSS class name for the container
func (bg buttonGroup) ClassName(value string) buttonGroup {
	return bg.set("className", value)
}

// Disabled enables or disables the button group
func (bg buttonGroup) Disabled(value bool) buttonGroup {
	return bg.set("disabled", value)
}

// DisabledOn configures the disabled state using a JavaScript expression
func (bg buttonGroup) DisabledOn(value string) buttonGroup {
	return bg.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings
func (bg buttonGroup) EditorSetting(value string) buttonGroup {
	return bg.set("editorSetting", value)
}

// Hidden controls the visibility of the button group
func (bg buttonGroup) Hidden(value bool) buttonGroup {
	return bg.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the button group
func (bg buttonGroup) HiddenOn(value string) buttonGroup {
	return bg.set("hiddenOn", value)
}

// ID sets a unique identifier for the component
func (bg buttonGroup) ID(value string) buttonGroup {
	return bg.set("id", value)
}

// OnEvent configures event-driven actions
func (bg buttonGroup) OnEvent(value any) buttonGroup {
	return bg.set("onEvent", value)
}

// Size sets the button size
func (bg buttonGroup) Size(value string) buttonGroup {
	return bg.set("size", value)
}

// Static determines if the button group is statically displayed
func (bg buttonGroup) Static(value bool) buttonGroup {
	return bg.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (bg buttonGroup) StaticClassName(value string) buttonGroup {
	return bg.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (bg buttonGroup) StaticInputClassName(value string) buttonGroup {
	return bg.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (bg buttonGroup) StaticLabelClassName(value string) buttonGroup {
	return bg.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (bg buttonGroup) StaticOn(value string) buttonGroup {
	return bg.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (bg buttonGroup) StaticPlaceholder(value string) buttonGroup {
	return bg.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (bg buttonGroup) StaticSchema(value string) buttonGroup {
	return bg.set("staticSchema", value)
}

// Style sets custom inline styles
func (bg buttonGroup) Style(value any) buttonGroup {
	return bg.set("style", value)
}

// TestIdBuilder configures test ID generation
func (bg buttonGroup) TestIdBuilder(value string) buttonGroup {
	return bg.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (bg buttonGroup) Testid(value string) buttonGroup {
	return bg.set("testid", value)
}

// Tiled enables or disables tiled display
func (bg buttonGroup) Tiled(value bool) buttonGroup {
	return bg.set("tiled", value)
}

// UseMobileUI enables or disables mobile UI styling
func (bg buttonGroup) UseMobileUI(value bool) buttonGroup {
	return bg.set("useMobileUI", value)
}

// Vertical controls vertical or horizontal layout
func (bg buttonGroup) Vertical(value bool) buttonGroup {
	return bg.set("vertical", value)
}

// Visible controls the overall visibility of the button group
func (bg buttonGroup) Visible(value bool) buttonGroup {
	return bg.set("visible", value)
}

// VisibleOn sets a conditional expression for button group visibility
func (bg buttonGroup) VisibleOn(value string) buttonGroup {
	return bg.set("visibleOn", value)
}
