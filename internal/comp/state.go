package comp

import "github.com/zrcoder/amisgo/schema"

// State represents a State configuration.
type State schema.Schema

func NewState() State {
	return State{}
}

func (s State) set(key string, value any) State {
	s[key] = value
	return s
}

// Body sets the content.
func (s State) Body(value ...any) State {
	return s.set("body", value)
}

// ClassName sets the container CSS class name.
func (s State) ClassName(value string) State {
	return s.set("className", value)
}

// Disabled sets whether the component is disabled.
func (s State) Disabled(value bool) State {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component should be disabled.
func (s State) DisabledOn(value string) State {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (s State) EditorSetting(value string) State {
	return s.set("editorSetting", value)
}

// Hidden sets whether the component is hidden.
func (s State) Hidden(value bool) State {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component should be hidden.
func (s State) HiddenOn(value string) State {
	return s.set("hiddenOn", value)
}

// ID sets the unique component ID.
func (s State) ID(value string) State {
	return s.set("id", value)
}

// OnEvent sets the event action configuration.
func (s State) OnEvent(value any) State {
	return s.set("onEvent", value)
}

// Static sets whether to display in static mode.
func (s State) Static(value bool) State {
	return s.set("static", value)
}

// StaticClassName sets the CSS class for static display.
func (s State) StaticClassName(value string) State {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class for static value display.
func (s State) StaticInputClassName(value string) State {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class for static label display.
func (s State) StaticLabelClassName(value string) State {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component should be displayed in static mode.
func (s State) StaticOn(value string) State {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder text for empty values in static display.
func (s State) StaticPlaceholder(value string) State {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display.
func (s State) StaticSchema(value string) State {
	return s.set("staticSchema", value)
}

// Style sets the component inline styles.
func (s State) Style(value any) State {
	return s.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (s State) TestIdBuilder(value string) State {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (s State) Testid(value string) State {
	return s.set("testid", value)
}

// Title sets the state title.
func (s State) Title(value any) State {
	return s.set("title", value)
}

// UseMobileUI sets whether to disable mobile UI styles at component level.
func (s State) UseMobileUI(value bool) State {
	return s.set("useMobileUI", value)
}

// Visible sets whether the component is visible.
func (s State) Visible(value bool) State {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine if the component should be visible.
func (s State) VisibleOn(value string) State {
	return s.set("visibleOn", value)
}
