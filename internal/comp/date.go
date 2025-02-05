package comp

import "github.com/zrcoder/amisgo/model"

// Date represents a Date component.
type Date model.Schema

func NewDate() Date {
	return Date{"type": "date"}
}

func (d Date) set(key string, value any) Date {
	d[key] = value
	return d
}

// ClassName sets the class name of the container.
func (d Date) ClassName(value string) Date {
	return d.set("className", value)
}

// Disabled sets whether the component is disabled.
func (d Date) Disabled(value bool) Date {
	return d.set("disabled", value)
}

// DisabledOn sets the expression for disabling the component.
func (d Date) DisabledOn(value string) Date {
	return d.set("disabledOn", value)
}

// DisplayFormat sets the format of the displayed date.
func (d Date) DisplayFormat(value string) Date {
	return d.set("displayFormat", value)
}

// DisplayTimeZone sets the time zone for the displayed date.
func (d Date) DisplayTimeZone(value string) Date {
	return d.set("displayTimeZone", value)
}

// EditorSetting sets the configuration for the editor.
func (d Date) EditorSetting(value string) Date {
	return d.set("editorSetting", value)
}

// Format sets the format of the displayed date.
func (d Date) Format(value string) Date {
	return d.set("format", value)
}

// FromNow shows the date relative to the current time.
func (d Date) FromNow(value bool) Date {
	return d.set("fromNow", value)
}

// Hidden sets whether the component is hidden.
func (d Date) Hidden(value bool) Date {
	return d.set("hidden", value)
}

// HiddenOn sets the expression for hiding the component.
func (d Date) HiddenOn(value string) Date {
	return d.set("hiddenOn", value)
}

// ID sets the unique id of the component.
func (d Date) ID(value string) Date {
	return d.set("id", value)
}

// OnEvent sets the event handler configuration.
func (d Date) OnEvent(value any) Date {
	return d.set("onEvent", value)
}

// Placeholder sets the placeholder of the component.
func (d Date) Placeholder(value string) Date {
	return d.set("placeholder", value)
}

// Static sets whether the component is statically displayed.
func (d Date) Static(value bool) Date {
	return d.set("static", value)
}

// StaticClassName sets the class name for statically displayed form items.
func (d Date) StaticClassName(value string) Date {
	return d.set("staticClassName", value)
}

// StaticInputClassName sets the class name for statically displayed form item values.
func (d Date) StaticInputClassName(value string) Date {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class name for statically displayed form item labels.
func (d Date) StaticLabelClassName(value string) Date {
	return d.set("staticLabelClassName", value)
}

// StaticOn sets the expression for statically displaying the component.
func (d Date) StaticOn(value string) Date {
	return d.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for statically displayed empty values.
func (d Date) StaticPlaceholder(value string) Date {
	return d.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for statically displayed form items.
func (d Date) StaticSchema(value string) Date {
	return d.set("staticSchema", value)
}

// Style sets the style of the component.
func (d Date) Style(value any) Date {
	return d.set("style", value)
}

// TestIDBuilder sets the id builder for testing.
func (d Date) TestIDBuilder(value string) Date {
	return d.set("testIdBuilder", value)
}

// TestID sets the id for testing.
func (d Date) TestID(value string) Date {
	return d.set("testid", value)
}

// UpdateFrequency sets the frequency of updating the component.
func (d Date) UpdateFrequency(value string) Date {
	return d.set("updateFrequency", value)
}

// UseMobileUI sets whether to use the mobile UI.
func (d Date) UseMobileUI(value bool) Date {
	return d.set("useMobileUI", value)
}

// ValueFormat sets the format of the value.
func (d Date) ValueFormat(value string) Date {
	return d.set("valueFormat", value)
}

// Visible sets whether the component is visible.
func (d Date) Visible(value bool) Date {
	return d.set("visible", value)
}

// VisibleOn sets the expression for showing the component.
func (d Date) VisibleOn(value string) Date {
	return d.set("visibleOn", value)
}
