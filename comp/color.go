package comp

import "github.com/zrcoder/amisgo/model"

// Color represents a color selection and display component
type color model.Schema

// Color creates a new Color instance
func Color() color {
	return color{"type": "color"}
}

func (c color) set(key string, value any) color {
	c[key] = value
	return c
}

// ClassName sets the CSS class name for the container
func (c color) ClassName(value string) color {
	return c.set("className", value)
}

// DefaultColor sets the initial default color value
func (c color) DefaultColor(value string) color {
	return c.set("defaultColor", value)
}

// Disabled enables or disables the color picker component
func (c color) Disabled(value bool) color {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the component
func (c color) DisabledOn(value string) color {
	return c.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings for the component
func (c color) EditorSetting(value string) color {
	return c.set("editorSetting", value)
}

// Hidden controls the overall visibility of the color picker
func (c color) Hidden(value bool) color {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the component
func (c color) HiddenOn(value string) color {
	return c.set("hiddenOn", value)
}

// ID sets a unique identifier for the component
func (c color) ID(value string) color {
	return c.set("id", value)
}

// OnEvent configures event-driven actions for the color picker
func (c color) OnEvent(value any) color {
	return c.set("onEvent", value)
}

// ShowValue enables or disables displaying the color value as text
func (c color) ShowValue(value bool) color {
	return c.set("showValue", value)
}

// Static enables or disables static display mode
func (c color) Static(value bool) color {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display
func (c color) StaticClassName(value string) color {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display
func (c color) StaticInputClassName(value string) color {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c color) StaticLabelClassName(value string) color {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c color) StaticOn(value string) color {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c color) StaticPlaceholder(value string) color {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (c color) StaticSchema(value string) color {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c color) Style(value any) color {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c color) TestIdBuilder(value string) color {
	return c.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (c color) Testid(value string) color {
	return c.set("testid", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c color) UseMobileUI(value bool) color {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the component
func (c color) Visible(value bool) color {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c color) VisibleOn(value string) color {
	return c.set("visibleOn", value)
}
