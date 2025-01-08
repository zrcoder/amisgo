package comp

type step Schema

func Step() step {
	return step{}
}

func (s step) set(key string, value any) step {
	s[key] = value
	return s
}

// ClassName sets the container CSS class name
// Can be a string or an object for conditional classes
// Example: {"red": "data.progress > 80", "blue": "data.progress > 60"}
func (s step) ClassName(value string) step {
	return s.set("className", value)
}

// Description sets the description text
func (s step) Description(value string) step {
	return s.set("description", value)
}

// Disabled sets whether the step is disabled
func (s step) Disabled(value bool) step {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to determine if the step should be disabled
func (s step) DisabledOn(value string) step {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, can be ignored at runtime
func (s step) EditorSetting(value string) step {
	return s.set("editorSetting", value)
}

// Hidden sets whether the step is hidden
func (s step) Hidden(value bool) step {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to determine if the step should be hidden
// Expression syntax: data.xxx > 5
func (s step) HiddenOn(value string) step {
	return s.set("hiddenOn", value)
}

// Icon sets the step icon
func (s step) Icon(value string) step {
	return s.set("icon", value)
}

// ID sets the unique component ID, mainly used for log collection
func (s step) ID(value string) step {
	return s.set("id", value)
}

// OnEvent sets the event action configuration
func (s step) OnEvent(value any) step {
	return s.set("onEvent", value)
}

// Static sets whether to display in static mode
func (s step) Static(value bool) step {
	return s.set("static", value)
}

// StaticClassName sets the CSS class for static display
// Can be a string or an object for conditional classes
// Example: {"red": "data.progress > 80", "blue": "data.progress > 60"}
func (s step) StaticClassName(value string) step {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class for static value display
// Can be a string or an object for conditional classes
// Example: {"red": "data.progress > 80", "blue": "data.progress > 60"}
func (s step) StaticInputClassName(value string) step {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class for static label display
// Can be a string or an object for conditional classes
// Example: {"red": "data.progress > 80", "blue": "data.progress > 60"}
func (s step) StaticLabelClassName(value string) step {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the step should be displayed in static mode
// Expression syntax: data.xxx > 5
func (s step) StaticOn(value string) step {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder text for empty values in static display
func (s step) StaticPlaceholder(value string) step {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (s step) StaticSchema(value string) step {
	return s.set("staticSchema", value)
}

// Style sets the component inline styles
func (s step) Style(value any) step {
	return s.set("style", value)
}

// SubTitle sets the step subtitle
func (s step) SubTitle(value any) step {
	return s.set("subTitle", value)
}

// TestIdBuilder sets the test ID builder
func (s step) TestIdBuilder(value string) step {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID
func (s step) Testid(value string) step {
	return s.set("testid", value)
}

// Title sets the step title
func (s step) Title(value any) step {
	return s.set("title", value)
}

// UseMobileUI sets whether to disable mobile UI styles at component level
func (s step) UseMobileUI(value bool) step {
	return s.set("useMobileUI", value)
}

// Value sets the step value
func (s step) Value(value string) step {
	return s.set("value", value)
}

// Visible sets whether the step is visible
func (s step) Visible(value bool) step {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine if the step should be visible
// Expression syntax: data.xxx > 5
func (s step) VisibleOn(value string) step {
	return s.set("visibleOn", value)
}
