package comp

import "github.com/zrcoder/amisgo/model"

// Steps represents a schema for Steps component
type Steps model.Schema

func NewSteps() Steps {
	return Steps{"type": "steps"}
}

func (s Steps) set(key string, value any) Steps {
	s[key] = value
	return s
}

// ClassName sets the CSS class name for the container
func (s Steps) ClassName(value string) Steps {
	return s.set("className", value)
}

// Disabled sets whether the component is disabled
func (s Steps) Disabled(value bool) Steps {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled
func (s Steps) DisabledOn(value string) Steps {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (s Steps) EditorSetting(value string) Steps {
	return s.set("editorSetting", value)
}

// Hidden sets whether the component is hidden
func (s Steps) Hidden(value bool) Steps {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden
func (s Steps) HiddenOn(value string) Steps {
	return s.set("hiddenOn", value)
}

// ID sets the unique ID for the component, mainly for logging
func (s Steps) ID(value string) Steps {
	return s.set("id", value)
}

// LabelPlacement sets the label placement, either horizontal or vertical
func (s Steps) LabelPlacement(value string) Steps {
	return s.set("labelPlacement", value)
}

// Mode sets the display mode, either horizontal or vertical
func (s Steps) Mode(value string) Steps {
	return s.set("mode", value)
}

// Name sets the variable mapping
func (s Steps) Name(value string) Steps {
	return s.set("name", value)
}

// OnEvent sets the event action configuration
func (s Steps) OnEvent(value any) Steps {
	return s.set("onEvent", value)
}

// ProgressDot sets whether to use dot style for steps
func (s Steps) ProgressDot(value bool) Steps {
	return s.set("progressDot", value)
}

// Source sets the API or data mapping
func (s Steps) Source(value string) Steps {
	return s.set("source", value)
}

// Static sets whether the component is displayed statically
func (s Steps) Static(value bool) Steps {
	return s.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (s Steps) StaticClassName(value string) Steps {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (s Steps) StaticInputClassName(value string) Steps {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (s Steps) StaticLabelClassName(value string) Steps {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is displayed statically
func (s Steps) StaticOn(value string) Steps {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (s Steps) StaticPlaceholder(value string) Steps {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (s Steps) StaticSchema(value string) Steps {
	return s.set("staticSchema", value)
}

// Status sets the status of the steps
func (s Steps) Status(value string) Steps {
	return s.set("status", value)
}

// Steps sets the steps
func (s Steps) Steps(value string) Steps {
	return s.set("steps", value)
}

// Style sets the style for the component
func (s Steps) Style(value any) Steps {
	return s.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (s Steps) TestIdBuilder(value string) Steps {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID
func (s Steps) Testid(value string) Steps {
	return s.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (s Steps) UseMobileUI(value bool) Steps {
	return s.set("useMobileUI", value)
}

// Value sets the current step
func (s Steps) Value(value string) Steps {
	return s.set("value", value)
}

// Visible sets whether the component is visible
func (s Steps) Visible(value bool) Steps {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible
func (s Steps) VisibleOn(value string) Steps {
	return s.set("visibleOn", value)
}
