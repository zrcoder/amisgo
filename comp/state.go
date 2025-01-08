package comp

import "github.com/zrcoder/amisgo/model"

// state represents a state configuration.

type state model.Schema

// State creates a new State instance.
func State() state {
	return state{}
}

func (s state) set(key string, value any) state {
	s[key] = value
	return s
}

// Body sets the content.
func (s state) Body(value ...any) state {
	return s.set("body", value)
}

// ClassName sets the container CSS class name.
func (s state) ClassName(value string) state {
	return s.set("className", value)
}

// Disabled sets whether the component is disabled.
func (s state) Disabled(value bool) state {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component should be disabled.
func (s state) DisabledOn(value string) state {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (s state) EditorSetting(value string) state {
	return s.set("editorSetting", value)
}

// Hidden sets whether the component is hidden.
func (s state) Hidden(value bool) state {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component should be hidden.
func (s state) HiddenOn(value string) state {
	return s.set("hiddenOn", value)
}

// ID sets the unique component ID.
func (s state) ID(value string) state {
	return s.set("id", value)
}

// OnEvent sets the event action configuration.
func (s state) OnEvent(value any) state {
	return s.set("onEvent", value)
}

// Static sets whether to display in static mode.
func (s state) Static(value bool) state {
	return s.set("static", value)
}

// StaticClassName sets the CSS class for static display.
func (s state) StaticClassName(value string) state {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class for static value display.
func (s state) StaticInputClassName(value string) state {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class for static label display.
func (s state) StaticLabelClassName(value string) state {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component should be displayed in static mode.
func (s state) StaticOn(value string) state {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder text for empty values in static display.
func (s state) StaticPlaceholder(value string) state {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display.
func (s state) StaticSchema(value string) state {
	return s.set("staticSchema", value)
}

// Style sets the component inline styles.
func (s state) Style(value any) state {
	return s.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (s state) TestIdBuilder(value string) state {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (s state) Testid(value string) state {
	return s.set("testid", value)
}

// Title sets the state title.
func (s state) Title(value any) state {
	return s.set("title", value)
}

// UseMobileUI sets whether to disable mobile UI styles at component level.
func (s state) UseMobileUI(value bool) state {
	return s.set("useMobileUI", value)
}

// Visible sets whether the component is visible.
func (s state) Visible(value bool) state {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine if the component should be visible.
func (s state) VisibleOn(value string) state {
	return s.set("visibleOn", value)
}
