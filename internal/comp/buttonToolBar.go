package comp

import "github.com/zrcoder/amisgo/schema"

// ButtonToolbar represents a button toolbar control renderer
type ButtonToolbar schema.Schema

func NewButtonToolbar() ButtonToolbar {
	return ButtonToolbar{"type": "button-toolbar"}
}

// set overrides the BaseRenderer's Set method, returning ButtonToolbar
func (bt ButtonToolbar) set(key string, value any) ButtonToolbar {
	bt[key] = value
	return bt
}

// Buttons configures the collection of buttons
func (bt ButtonToolbar) Buttons(value ...Action) ButtonToolbar {
	return bt.set("buttons", value)
}

// Label sets the label for the button toolbar
func (bt ButtonToolbar) Label(value string) ButtonToolbar {
	return bt.set("label", value)
}

// ClassName sets the CSS class name for the container
func (bt ButtonToolbar) ClassName(value string) ButtonToolbar {
	return bt.set("className", value)
}

// Disabled enables or disables the button toolbar
func (bt ButtonToolbar) Disabled(value bool) ButtonToolbar {
	return bt.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the button toolbar
func (bt ButtonToolbar) DisabledOn(value string) ButtonToolbar {
	return bt.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings
func (bt ButtonToolbar) EditorSetting(value string) ButtonToolbar {
	return bt.set("editorSetting", value)
}

// Hidden controls the visibility of the button toolbar
func (bt ButtonToolbar) Hidden(value bool) ButtonToolbar {
	return bt.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the button toolbar
func (bt ButtonToolbar) HiddenOn(value string) ButtonToolbar {
	return bt.set("hiddenOn", value)
}

// ID sets a unique identifier for the component
func (bt ButtonToolbar) ID(value string) ButtonToolbar {
	return bt.set("id", value)
}

// OnEvent configures event-driven actions
func (bt ButtonToolbar) OnEvent(value Event) ButtonToolbar {
	return bt.set("onEvent", value)
}

// Static determines if the button toolbar is statically displayed
func (bt ButtonToolbar) Static(value bool) ButtonToolbar {
	return bt.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (bt ButtonToolbar) StaticClassName(value string) ButtonToolbar {
	return bt.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (bt ButtonToolbar) StaticInputClassName(value string) ButtonToolbar {
	return bt.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (bt ButtonToolbar) StaticLabelClassName(value string) ButtonToolbar {
	return bt.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (bt ButtonToolbar) StaticOn(value string) ButtonToolbar {
	return bt.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (bt ButtonToolbar) StaticPlaceholder(value string) ButtonToolbar {
	return bt.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (bt ButtonToolbar) StaticSchema(value string) ButtonToolbar {
	return bt.set("staticSchema", value)
}

// Style sets custom inline styles
func (bt ButtonToolbar) Style(value any) ButtonToolbar {
	return bt.set("style", value)
}

// TestIdBuilder configures test ID generation
func (bt ButtonToolbar) TestIdBuilder(value string) ButtonToolbar {
	return bt.set("testIdBuilder", value)
}

// Testid sets a specific test identifier
func (bt ButtonToolbar) Testid(value string) ButtonToolbar {
	return bt.set("testid", value)
}

// UseMobileUI enables or disables mobile UI styling
func (bt ButtonToolbar) UseMobileUI(value bool) ButtonToolbar {
	return bt.set("useMobileUI", value)
}

// Visible controls the overall visibility of the button toolbar
func (bt ButtonToolbar) Visible(value bool) ButtonToolbar {
	return bt.set("visible", value)
}

// VisibleOn sets a conditional expression for button toolbar visibility
func (bt ButtonToolbar) VisibleOn(value string) ButtonToolbar {
	return bt.set("visibleOn", value)
}
