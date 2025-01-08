package comp

import "github.com/zrcoder/amisgo/model"

// Spinner represents a spinner configuration.
type spinner model.Schema

// Spinner creates a new Spinner instance.
func Spinner() spinner {
	return spinner{}.set("type", "spinner")
}

func (s spinner) set(key string, value any) spinner {
	s[key] = value
	return s
}

// Body sets the content as a container.
func (s spinner) Body(value ...any) spinner {
	return s.set("body", value)
}

// ClassName sets the custom spinner class.
func (s spinner) ClassName(value string) spinner {
	return s.set("className", value)
}

// Delay sets the delay for display.
func (s spinner) Delay(value string) spinner {
	return s.set("delay", value)
}

// Disabled sets whether the spinner is disabled.
func (s spinner) Disabled(value bool) spinner {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to disable the spinner.
func (s spinner) DisabledOn(value string) spinner {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (s spinner) EditorSetting(value string) spinner {
	return s.set("editorSetting", value)
}

// Hidden sets whether the spinner is hidden.
func (s spinner) Hidden(value bool) spinner {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to hide the spinner.
func (s spinner) HiddenOn(value string) spinner {
	return s.set("hiddenOn", value)
}

// Icon sets the custom icon.
func (s spinner) Icon(value string) spinner {
	return s.set("icon", value)
}

// ID sets the unique component ID.
func (s spinner) ID(value string) spinner {
	return s.set("id", value)
}

// LoadingConfig sets the loading configuration.
func (s spinner) LoadingConfig(value string) spinner {
	return s.set("loadingConfig", value)
}

// Mode sets the mode.
func (s spinner) Mode(value string) spinner {
	return s.set("mode", value)
}

// OnEvent sets the event action configuration.
func (s spinner) OnEvent(value any) spinner {
	return s.set("onEvent", value)
}

// Overlay sets whether to show the overlay layer.
func (s spinner) Overlay(value bool) spinner {
	return s.set("overlay", value)
}

// Show controls the display of the Spinner.
func (s spinner) Show(value bool) spinner {
	return s.set("show", value)
}

// Size sets the spinner icon size.
func (s spinner) Size(value string) spinner {
	return s.set("size", value)
}

// SpinnerClassName sets the custom class for the spinner icon wrapper.
func (s spinner) SpinnerClassName(value string) spinner {
	return s.set("spinnerClassName", value)
}

// SpinnerWrapClassName sets the class for the outermost element when used as a container.
func (s spinner) SpinnerWrapClassName(value string) spinner {
	return s.set("spinnerWrapClassName", value)
}

// Static sets whether to display statically.
func (s spinner) Static(value bool) spinner {
	return s.set("static", value)
}

// StaticClassName sets the class for static display.
func (s spinner) StaticClassName(value string) spinner {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the class for static value display.
func (s spinner) StaticInputClassName(value string) spinner {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class for static label display.
func (s spinner) StaticLabelClassName(value string) spinner {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display.
func (s spinner) StaticOn(value string) spinner {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for empty values in static display.
func (s spinner) StaticPlaceholder(value string) spinner {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display.
func (s spinner) StaticSchema(value string) spinner {
	return s.set("staticSchema", value)
}

// Style sets the component style.
func (s spinner) Style(value any) spinner {
	return s.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (s spinner) TestIdBuilder(value string) spinner {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (s spinner) Testid(value string) spinner {
	return s.set("testid", value)
}

// Tip sets the spinner text.
func (s spinner) Tip(value string) spinner {
	return s.set("tip", value)
}

// TipPlacement sets the placement of the spinner text.
func (s spinner) TipPlacement(value string) spinner {
	return s.set("tipPlacement", value)
}

// UseMobileUI sets whether to use mobile UI styles.
func (s spinner) UseMobileUI(value bool) spinner {
	return s.set("useMobileUI", value)
}

// Visible sets whether the spinner is visible.
func (s spinner) Visible(value bool) spinner {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine visibility.
func (s spinner) VisibleOn(value string) spinner {
	return s.set("visibleOn", value)
}
