package comp

import "github.com/zrcoder/amisgo/model"

// ButtonGroup represents a button group renderer
type ButtonGroup model.Schema

func NewButtonGroup() ButtonGroup {
	return ButtonGroup{"type": "button-group"}
}

func (br ButtonGroup) set(key string, value any) ButtonGroup {
	br[key] = value
	return br
}

// BtnActiveClassName sets the CSS class name for button active state
func (bg ButtonGroup) BtnActiveClassName(value string) ButtonGroup {
	return bg.set("btnActiveClassName", value)
}

// BtnActiveLevel sets the style level for selected buttons
// Supported values: 'link' | 'primary' | 'secondary' | 'info' | 'success' | 'warning' | 'danger' | 'light' | 'dark' | 'default'
func (bg ButtonGroup) BtnActiveLevel(value string) ButtonGroup {
	return bg.set("btnActiveLevel", value)
}

// BtnClassName sets the CSS class name for buttons
func (bg ButtonGroup) BtnClassName(value string) ButtonGroup {
	return bg.set("btnClassName", value)
}

// BtnLevel sets the style level for buttons
// Supported values: 'link' | 'primary' | 'secondary' | 'info' | 'success' | 'warning' | 'danger' | 'light' | 'dark' | 'default'
func (bg ButtonGroup) BtnLevel(value string) ButtonGroup {
	return bg.set("btnLevel", value)
}

// Buttons configures the collection of buttons
func (bg ButtonGroup) Buttons(value ...any) ButtonGroup {
	return bg.set("buttons", value)
}

// ClassName sets the CSS class name for the container
func (bg ButtonGroup) ClassName(value string) ButtonGroup {
	return bg.set("className", value)
}

// Disabled enables or disables the button group
func (bg ButtonGroup) Disabled(value bool) ButtonGroup {
	return bg.set("disabled", value)
}

// DisabledOn configures the disabled state using a JavaScript expression
func (bg ButtonGroup) DisabledOn(value string) ButtonGroup {
	return bg.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings
func (bg ButtonGroup) EditorSetting(value string) ButtonGroup {
	return bg.set("editorSetting", value)
}

// Hidden controls the visibility of the button group
func (bg ButtonGroup) Hidden(value bool) ButtonGroup {
	return bg.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the button group
func (bg ButtonGroup) HiddenOn(value string) ButtonGroup {
	return bg.set("hiddenOn", value)
}

// ID sets a unique identifier for the component
func (bg ButtonGroup) ID(value string) ButtonGroup {
	return bg.set("id", value)
}

// OnEvent configures event-driven actions
func (bg ButtonGroup) OnEvent(value any) ButtonGroup {
	return bg.set("onEvent", value)
}

// Size sets the button size
func (bg ButtonGroup) Size(value string) ButtonGroup {
	return bg.set("size", value)
}

// Static determines if the button group is statically displayed
func (bg ButtonGroup) Static(value bool) ButtonGroup {
	return bg.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (bg ButtonGroup) StaticClassName(value string) ButtonGroup {
	return bg.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (bg ButtonGroup) StaticInputClassName(value string) ButtonGroup {
	return bg.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (bg ButtonGroup) StaticLabelClassName(value string) ButtonGroup {
	return bg.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (bg ButtonGroup) StaticOn(value string) ButtonGroup {
	return bg.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (bg ButtonGroup) StaticPlaceholder(value string) ButtonGroup {
	return bg.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (bg ButtonGroup) StaticSchema(value string) ButtonGroup {
	return bg.set("staticSchema", value)
}

// Style sets custom inline styles
func (bg ButtonGroup) Style(value any) ButtonGroup {
	return bg.set("style", value)
}

// TestIdBuilder configures test ID generation
func (bg ButtonGroup) TestIdBuilder(value string) ButtonGroup {
	return bg.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (bg ButtonGroup) Testid(value string) ButtonGroup {
	return bg.set("testid", value)
}

// Tiled enables or disables tiled display
func (bg ButtonGroup) Tiled(value bool) ButtonGroup {
	return bg.set("tiled", value)
}

// UseMobileUI enables or disables mobile UI styling
func (bg ButtonGroup) UseMobileUI(value bool) ButtonGroup {
	return bg.set("useMobileUI", value)
}

// Vertical controls vertical or horizontal layout
func (bg ButtonGroup) Vertical(value bool) ButtonGroup {
	return bg.set("vertical", value)
}

// Visible controls the overall visibility of the button group
func (bg ButtonGroup) Visible(value bool) ButtonGroup {
	return bg.set("visible", value)
}

// VisibleOn sets a conditional expression for button group visibility
func (bg ButtonGroup) VisibleOn(value string) ButtonGroup {
	return bg.set("visibleOn", value)
}
