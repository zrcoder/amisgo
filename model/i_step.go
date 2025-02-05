package model

type Step Schema

func NewStep() Step {
	return Step{}
}

func (s Step) set(key string, value any) Step {
	s[key] = value
	return s
}

// ClassName sets the container CSS class name
// Can be a string or an object for conditional classes
// Example: {"red": "data.progress > 80", "blue": "data.progress > 60"}
func (s Step) ClassName(value string) Step {
	return s.set("className", value)
}

// Description sets the description text
func (s Step) Description(value string) Step {
	return s.set("description", value)
}

// Disabled sets whether the step is disabled
func (s Step) Disabled(value bool) Step {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to determine if the step should be disabled
func (s Step) DisabledOn(value string) Step {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, can be ignored at runtime
func (s Step) EditorSetting(value string) Step {
	return s.set("editorSetting", value)
}

// Hidden sets whether the step is hidden
func (s Step) Hidden(value bool) Step {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to determine if the step should be hidden
// Expression syntax: data.xxx > 5
func (s Step) HiddenOn(value string) Step {
	return s.set("hiddenOn", value)
}

// Icon sets the step icon
func (s Step) Icon(value string) Step {
	return s.set("icon", value)
}

// ID sets the unique component ID, mainly used for log collection
func (s Step) ID(value string) Step {
	return s.set("id", value)
}

// OnEvent sets the event action configuration
func (s Step) OnEvent(value any) Step {
	return s.set("onEvent", value)
}

// Static sets whether to display in static mode
func (s Step) Static(value bool) Step {
	return s.set("static", value)
}

// StaticClassName sets the CSS class for static display
// Can be a string or an object for conditional classes
// Example: {"red": "data.progress > 80", "blue": "data.progress > 60"}
func (s Step) StaticClassName(value string) Step {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class for static value display
// Can be a string or an object for conditional classes
// Example: {"red": "data.progress > 80", "blue": "data.progress > 60"}
func (s Step) StaticInputClassName(value string) Step {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class for static label display
// Can be a string or an object for conditional classes
// Example: {"red": "data.progress > 80", "blue": "data.progress > 60"}
func (s Step) StaticLabelClassName(value string) Step {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the step should be displayed in static mode
// Expression syntax: data.xxx > 5
func (s Step) StaticOn(value string) Step {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder text for empty values in static display
func (s Step) StaticPlaceholder(value string) Step {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (s Step) StaticSchema(value string) Step {
	return s.set("staticSchema", value)
}

// Style sets the component inline styles
func (s Step) Style(value any) Step {
	return s.set("style", value)
}

// SubTitle sets the step subtitle
func (s Step) SubTitle(value any) Step {
	return s.set("subTitle", value)
}

// TestIdBuilder sets the test ID builder
func (s Step) TestIdBuilder(value string) Step {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID
func (s Step) Testid(value string) Step {
	return s.set("testid", value)
}

// Title sets the step title
func (s Step) Title(value any) Step {
	return s.set("title", value)
}

// UseMobileUI sets whether to disable mobile UI styles at component level
func (s Step) UseMobileUI(value bool) Step {
	return s.set("useMobileUI", value)
}

// Value sets the step value
func (s Step) Value(value string) Step {
	return s.set("value", value)
}

// Visible sets whether the step is visible
func (s Step) Visible(value bool) Step {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine if the step should be visible
// Expression syntax: data.xxx > 5
func (s Step) VisibleOn(value string) Step {
	return s.set("visibleOn", value)
}
