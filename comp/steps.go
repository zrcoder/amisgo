package comp

import "github.com/zrcoder/amisgo/model"

// steps represents a schema for steps component

type steps model.Schema

// Steps creates a new Steps instance
func Steps() steps {
	return steps{}.set("type", "steps")
}

func (s steps) set(key string, value any) steps {
	s[key] = value
	return s
}

// ClassName sets the CSS class name for the container
func (s steps) ClassName(value string) steps {
	return s.set("className", value)
}

// Disabled sets whether the component is disabled
func (s steps) Disabled(value bool) steps {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled
func (s steps) DisabledOn(value string) steps {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (s steps) EditorSetting(value string) steps {
	return s.set("editorSetting", value)
}

// Hidden sets whether the component is hidden
func (s steps) Hidden(value bool) steps {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden
func (s steps) HiddenOn(value string) steps {
	return s.set("hiddenOn", value)
}

// ID sets the unique ID for the component, mainly for logging
func (s steps) ID(value string) steps {
	return s.set("id", value)
}

// LabelPlacement sets the label placement, either horizontal or vertical
func (s steps) LabelPlacement(value string) steps {
	return s.set("labelPlacement", value)
}

// Mode sets the display mode, either horizontal or vertical
func (s steps) Mode(value string) steps {
	return s.set("mode", value)
}

// Name sets the variable mapping
func (s steps) Name(value string) steps {
	return s.set("name", value)
}

// OnEvent sets the event action configuration
func (s steps) OnEvent(value any) steps {
	return s.set("onEvent", value)
}

// ProgressDot sets whether to use dot style for steps
func (s steps) ProgressDot(value bool) steps {
	return s.set("progressDot", value)
}

// Source sets the API or data mapping
func (s steps) Source(value string) steps {
	return s.set("source", value)
}

// Static sets whether the component is displayed statically
func (s steps) Static(value bool) steps {
	return s.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (s steps) StaticClassName(value string) steps {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (s steps) StaticInputClassName(value string) steps {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (s steps) StaticLabelClassName(value string) steps {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is displayed statically
func (s steps) StaticOn(value string) steps {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (s steps) StaticPlaceholder(value string) steps {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (s steps) StaticSchema(value string) steps {
	return s.set("staticSchema", value)
}

// Status sets the status of the steps
func (s steps) Status(value string) steps {
	return s.set("status", value)
}

// Steps sets the steps
func (s steps) Steps(value string) steps {
	return s.set("steps", value)
}

// Style sets the style for the component
func (s steps) Style(value any) steps {
	return s.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (s steps) TestIdBuilder(value string) steps {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID
func (s steps) Testid(value string) steps {
	return s.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (s steps) UseMobileUI(value bool) steps {
	return s.set("useMobileUI", value)
}

// Value sets the current step
func (s steps) Value(value string) steps {
	return s.set("value", value)
}

// Visible sets whether the component is visible
func (s steps) Visible(value bool) steps {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible
func (s steps) VisibleOn(value string) steps {
	return s.set("visibleOn", value)
}
