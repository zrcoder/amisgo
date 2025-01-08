package comp

import "github.com/zrcoder/amisgo/model"

// Date represents a date component.
type date model.Schema

// Date creates a new date component and sets the default type.
func Date() date {
	return make(date).set("type", "date")
}

func (d date) set(key string, value any) date {
	d[key] = value
	return d
}

// ClassName sets the class name of the container.
func (d date) ClassName(value string) date {
	return d.set("className", value)
}

// Disabled sets whether the component is disabled.
func (d date) Disabled(value bool) date {
	return d.set("disabled", value)
}

// DisabledOn sets the expression for disabling the component.
func (d date) DisabledOn(value string) date {
	return d.set("disabledOn", value)
}

// DisplayFormat sets the format of the displayed date.
func (d date) DisplayFormat(value string) date {
	return d.set("displayFormat", value)
}

// DisplayTimeZone sets the time zone for the displayed date.
func (d date) DisplayTimeZone(value string) date {
	return d.set("displayTimeZone", value)
}

// EditorSetting sets the configuration for the editor.
func (d date) EditorSetting(value string) date {
	return d.set("editorSetting", value)
}

// Format sets the format of the displayed date.
func (d date) Format(value string) date {
	return d.set("format", value)
}

// FromNow shows the date relative to the current time.
func (d date) FromNow(value bool) date {
	return d.set("fromNow", value)
}

// Hidden sets whether the component is hidden.
func (d date) Hidden(value bool) date {
	return d.set("hidden", value)
}

// HiddenOn sets the expression for hiding the component.
func (d date) HiddenOn(value string) date {
	return d.set("hiddenOn", value)
}

// ID sets the unique id of the component.
func (d date) ID(value string) date {
	return d.set("id", value)
}

// OnEvent sets the event handler configuration.
func (d date) OnEvent(value any) date {
	return d.set("onEvent", value)
}

// Placeholder sets the placeholder of the component.
func (d date) Placeholder(value string) date {
	return d.set("placeholder", value)
}

// Static sets whether the component is statically displayed.
func (d date) Static(value bool) date {
	return d.set("static", value)
}

// StaticClassName sets the class name for statically displayed form items.
func (d date) StaticClassName(value string) date {
	return d.set("staticClassName", value)
}

// StaticInputClassName sets the class name for statically displayed form item values.
func (d date) StaticInputClassName(value string) date {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class name for statically displayed form item labels.
func (d date) StaticLabelClassName(value string) date {
	return d.set("staticLabelClassName", value)
}

// StaticOn sets the expression for statically displaying the component.
func (d date) StaticOn(value string) date {
	return d.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for statically displayed empty values.
func (d date) StaticPlaceholder(value string) date {
	return d.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for statically displayed form items.
func (d date) StaticSchema(value string) date {
	return d.set("staticSchema", value)
}

// Style sets the style of the component.
func (d date) Style(value any) date {
	return d.set("style", value)
}

// TestIDBuilder sets the id builder for testing.
func (d date) TestIDBuilder(value string) date {
	return d.set("testIdBuilder", value)
}

// TestID sets the id for testing.
func (d date) TestID(value string) date {
	return d.set("testid", value)
}

// UpdateFrequency sets the frequency of updating the component.
func (d date) UpdateFrequency(value string) date {
	return d.set("updateFrequency", value)
}

// UseMobileUI sets whether to use the mobile UI.
func (d date) UseMobileUI(value bool) date {
	return d.set("useMobileUI", value)
}

// ValueFormat sets the format of the value.
func (d date) ValueFormat(value string) date {
	return d.set("valueFormat", value)
}

// Visible sets whether the component is visible.
func (d date) Visible(value bool) date {
	return d.set("visible", value)
}

// VisibleOn sets the expression for showing the component.
func (d date) VisibleOn(value string) date {
	return d.set("visibleOn", value)
}
