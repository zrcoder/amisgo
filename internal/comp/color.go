package comp

import "github.com/zrcoder/amisgo/model"

// Color represents a Color selection and display component
type Color model.Schema

func NewColor() Color {
	return Color{"type": "color"}
}

func (c Color) set(key string, value any) Color {
	c[key] = value
	return c
}

// ClassName sets the CSS class name for the container
func (c Color) ClassName(value string) Color {
	return c.set("className", value)
}

// DefaultColor sets the initial default color value
func (c Color) DefaultColor(value string) Color {
	return c.set("defaultColor", value)
}

// Disabled enables or disables the color picker component
func (c Color) Disabled(value bool) Color {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the component
func (c Color) DisabledOn(value string) Color {
	return c.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings for the component
func (c Color) EditorSetting(value string) Color {
	return c.set("editorSetting", value)
}

// Hidden controls the overall visibility of the color picker
func (c Color) Hidden(value bool) Color {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the component
func (c Color) HiddenOn(value string) Color {
	return c.set("hiddenOn", value)
}

// ID sets a unique identifier for the component
func (c Color) ID(value string) Color {
	return c.set("id", value)
}

// OnEvent configures event-driven actions for the color picker
func (c Color) OnEvent(value any) Color {
	return c.set("onEvent", value)
}

// ShowValue enables or disables displaying the color value as text
func (c Color) ShowValue(value bool) Color {
	return c.set("showValue", value)
}

// Static enables or disables static display mode
func (c Color) Static(value bool) Color {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display
func (c Color) StaticClassName(value string) Color {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display
func (c Color) StaticInputClassName(value string) Color {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c Color) StaticLabelClassName(value string) Color {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c Color) StaticOn(value string) Color {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c Color) StaticPlaceholder(value string) Color {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (c Color) StaticSchema(value string) Color {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c Color) Style(value any) Color {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c Color) TestIdBuilder(value string) Color {
	return c.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (c Color) Testid(value string) Color {
	return c.set("testid", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c Color) UseMobileUI(value bool) Color {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the component
func (c Color) Visible(value bool) Color {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c Color) VisibleOn(value string) Color {
	return c.set("visibleOn", value)
}
