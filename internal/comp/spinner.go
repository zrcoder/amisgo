package comp

import "github.com/zrcoder/amisgo/schema"

// Spinner represents a Spinner configuration.
type Spinner schema.Schema

func NewSpinner() Spinner {
	return Spinner{"type": "spinner"}
}

func (s Spinner) set(key string, value any) Spinner {
	s[key] = value
	return s
}

// Body sets the content as a container.
func (s Spinner) Body(value ...any) Spinner {
	return s.set("body", value)
}

// ClassName sets the custom spinner class.
func (s Spinner) ClassName(value string) Spinner {
	return s.set("className", value)
}

// Delay sets the delay for display.
func (s Spinner) Delay(value string) Spinner {
	return s.set("delay", value)
}

// Disabled sets whether the spinner is disabled.
func (s Spinner) Disabled(value bool) Spinner {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to disable the spinner.
func (s Spinner) DisabledOn(value string) Spinner {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (s Spinner) EditorSetting(value string) Spinner {
	return s.set("editorSetting", value)
}

// Hidden sets whether the spinner is hidden.
func (s Spinner) Hidden(value bool) Spinner {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to hide the spinner.
func (s Spinner) HiddenOn(value string) Spinner {
	return s.set("hiddenOn", value)
}

// Icon sets the custom icon.
func (s Spinner) Icon(value string) Spinner {
	return s.set("icon", value)
}

// ID sets the unique component ID.
func (s Spinner) ID(value string) Spinner {
	return s.set("id", value)
}

// LoadingConfig sets the loading configuration.
func (s Spinner) LoadingConfig(value string) Spinner {
	return s.set("loadingConfig", value)
}

// Mode sets the mode.
func (s Spinner) Mode(value string) Spinner {
	return s.set("mode", value)
}

// OnEvent sets the event action configuration.
func (s Spinner) OnEvent(value Event) Spinner {
	return s.set("onEvent", value)
}

// Overlay sets whether to show the overlay layer.
func (s Spinner) Overlay(value bool) Spinner {
	return s.set("overlay", value)
}

// Show controls the display of the Spinner.
func (s Spinner) Show(value bool) Spinner {
	return s.set("show", value)
}

// Size sets the spinner icon size.
func (s Spinner) Size(value string) Spinner {
	return s.set("size", value)
}

// SpinnerClassName sets the custom class for the spinner icon wrapper.
func (s Spinner) SpinnerClassName(value string) Spinner {
	return s.set("spinnerClassName", value)
}

// SpinnerWrapClassName sets the class for the outermost element when used as a container.
func (s Spinner) SpinnerWrapClassName(value string) Spinner {
	return s.set("spinnerWrapClassName", value)
}

// Static sets whether to display statically.
func (s Spinner) Static(value bool) Spinner {
	return s.set("static", value)
}

// StaticClassName sets the class for static display.
func (s Spinner) StaticClassName(value string) Spinner {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the class for static value display.
func (s Spinner) StaticInputClassName(value string) Spinner {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class for static label display.
func (s Spinner) StaticLabelClassName(value string) Spinner {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display.
func (s Spinner) StaticOn(value string) Spinner {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for empty values in static display.
func (s Spinner) StaticPlaceholder(value string) Spinner {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display.
func (s Spinner) StaticSchema(value string) Spinner {
	return s.set("staticSchema", value)
}

// Style sets the component style.
func (s Spinner) Style(value any) Spinner {
	return s.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (s Spinner) TestIdBuilder(value string) Spinner {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (s Spinner) Testid(value string) Spinner {
	return s.set("testid", value)
}

// Tip sets the spinner text.
func (s Spinner) Tip(value string) Spinner {
	return s.set("tip", value)
}

// TipPlacement sets the placement of the spinner text.
func (s Spinner) TipPlacement(value string) Spinner {
	return s.set("tipPlacement", value)
}

// UseMobileUI sets whether to use mobile UI styles.
func (s Spinner) UseMobileUI(value bool) Spinner {
	return s.set("useMobileUI", value)
}

// Visible sets whether the spinner is visible.
func (s Spinner) Visible(value bool) Spinner {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine visibility.
func (s Spinner) VisibleOn(value string) Spinner {
	return s.set("visibleOn", value)
}
