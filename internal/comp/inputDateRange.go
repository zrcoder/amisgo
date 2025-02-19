package comp

import "github.com/zrcoder/amisgo/schema"

// InputDateRange represents a date range input control
type InputDateRange schema.Schema

func NewInputDateRange() InputDateRange {
	return InputDateRange{"type": "input-date-range"}
}

// set sets a key-value pair and returns the current instance
func (d InputDateRange) set(key string, value any) InputDateRange {
	d[key] = value
	return d
}

// Animation sets the animation effect
func (d InputDateRange) Animation(value bool) InputDateRange {
	return d.set("animation", value)
}

// AutoFill sets the auto-fill value
func (d InputDateRange) AutoFill(value string) InputDateRange {
	return d.set("autoFill", value)
}

// BorderMode sets the border mode
func (d InputDateRange) BorderMode(value string) InputDateRange {
	return d.set("borderMode", value)
}

// ClassName sets the class name
func (d InputDateRange) ClassName(value string) InputDateRange {
	return d.set("className", value)
}

// ClearValueOnHidden sets whether to clear value when hidden
func (d InputDateRange) ClearValueOnHidden(value bool) InputDateRange {
	return d.set("clearValueOnHidden", value)
}

// Delimiter sets the delimiter
func (d InputDateRange) Delimiter(value string) InputDateRange {
	return d.set("delimiter", value)
}

// Desc sets the description
func (d InputDateRange) Desc(value string) InputDateRange {
	return d.set("desc", value)
}

// Description sets the description
func (d InputDateRange) Description(value string) InputDateRange {
	return d.set("description", value)
}

// DescriptionClassName sets the description class name
func (d InputDateRange) DescriptionClassName(value string) InputDateRange {
	return d.set("descriptionClassName", value)
}

// Disabled sets whether the control is disabled
func (d InputDateRange) Disabled(value bool) InputDateRange {
	return d.set("disabled", value)
}

// DisabledOn sets the condition for disabling the control
func (d InputDateRange) DisabledOn(value string) InputDateRange {
	return d.set("disabledOn", value)
}

// DisplayFormat sets the display format
func (d InputDateRange) DisplayFormat(value string) InputDateRange {
	return d.set("displayFormat", value)
}

// EditorSetting sets the editor configuration
func (d InputDateRange) EditorSetting(value string) InputDateRange {
	return d.set("editorSetting", value)
}

// Embed sets whether to use inline mode
func (d InputDateRange) Embed(value bool) InputDateRange {
	return d.set("embed", value)
}

// EndPlaceholder sets the end placeholder
func (d InputDateRange) EndPlaceholder(value string) InputDateRange {
	return d.set("endPlaceholder", value)
}

// ExtraName sets the extra field name
func (d InputDateRange) ExtraName(value string) InputDateRange {
	return d.set("extraName", value)
}

// Format sets the submission format
func (d InputDateRange) Format(value string) InputDateRange {
	return d.set("format", value)
}

// Hidden sets whether the control is hidden
func (d InputDateRange) Hidden(value bool) InputDateRange {
	return d.set("hidden", value)
}

// HiddenOn sets the condition for hiding the control
func (d InputDateRange) HiddenOn(value string) InputDateRange {
	return d.set("hiddenOn", value)
}

// Hint sets the input hint
func (d InputDateRange) Hint(value string) InputDateRange {
	return d.set("hint", value)
}

// Horizontal sets the horizontal layout
func (d InputDateRange) Horizontal(value string) InputDateRange {
	return d.set("horizontal", value)
}

// ID sets the component ID
func (d InputDateRange) ID(value string) InputDateRange {
	return d.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (d InputDateRange) InitAutoFill(value string) InputDateRange {
	return d.set("initAutoFill", value)
}

// Inline sets whether to use inline mode
func (d InputDateRange) Inline(value bool) InputDateRange {
	return d.set("inline", value)
}

// Value sets the value
func (d InputDateRange) Value(value string) InputDateRange {
	return d.set("value", value)
}

// CloseOnSelect sets whether to close on select
func (d InputDateRange) CloseOnSelect(value bool) InputDateRange {
	return d.set("closeOnSelect", value)
}

// Clearable sets whether the control is clearable
func (d InputDateRange) Clearable(value bool) InputDateRange {
	return d.set("clearable", value)
}

// InputClassName sets the input class name
func (d InputDateRange) InputClassName(value string) InputDateRange {
	return d.set("inputClassName", value)
}

// InputFormat sets the input format
func (d InputDateRange) InputFormat(value string) InputDateRange {
	return d.set("inputFormat", value)
}

// JoinValues sets whether to join values into a single string
func (d InputDateRange) JoinValues(value bool) InputDateRange {
	return d.set("joinValues", value)
}

// Label sets the control label
func (d InputDateRange) Label(value string) InputDateRange {
	return d.set("label", value)
}

// LabelAlign sets the label alignment
func (d InputDateRange) LabelAlign(value string) InputDateRange {
	return d.set("labelAlign", value)
}

// LabelClassName sets the label class name
func (d InputDateRange) LabelClassName(value string) InputDateRange {
	return d.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (d InputDateRange) LabelRemark(value string) InputDateRange {
	return d.set("labelRemark", value)
}

// LabelWidth sets the label width
func (d InputDateRange) LabelWidth(value string) InputDateRange {
	return d.set("labelWidth", value)
}

// MaxDate sets the maximum date limit
func (d InputDateRange) MaxDate(value string) InputDateRange {
	return d.set("maxDate", value)
}

// MaxDuration sets the maximum duration
func (d InputDateRange) MaxDuration(value string) InputDateRange {
	return d.set("maxDuration", value)
}

// MinDate sets the minimum date limit
func (d InputDateRange) MinDate(value string) InputDateRange {
	return d.set("minDate", value)
}

// MinDuration sets the minimum duration
func (d InputDateRange) MinDuration(value string) InputDateRange {
	return d.set("minDuration", value)
}

// Mode sets the display mode
func (d InputDateRange) Mode(value string) InputDateRange {
	return d.set("mode", value)
}

// Name sets the form submission field name
func (d InputDateRange) Name(value string) InputDateRange {
	return d.set("name", value)
}

// OnEvent sets the event configuration
func (d InputDateRange) OnEvent(value any) InputDateRange {
	return d.set("onEvent", value)
}

// Placeholder sets the placeholder text
func (d InputDateRange) Placeholder(value string) InputDateRange {
	return d.set("placeholder", value)
}

// PopOverContainerSelector sets the popover container selector
func (d InputDateRange) PopOverContainerSelector(value string) InputDateRange {
	return d.set("popOverContainerSelector", value)
}

// Ranges sets the date range shortcuts
func (d InputDateRange) Ranges(value string) InputDateRange {
	return d.set("ranges", value)
}

// ReadOnly sets whether the control is read-only
func (d InputDateRange) ReadOnly(value bool) InputDateRange {
	return d.set("readOnly", value)
}

// ReadOnlyOn sets the condition for read-only
func (d InputDateRange) ReadOnlyOn(value string) InputDateRange {
	return d.set("readOnlyOn", value)
}

// Remark sets the control remark
func (d InputDateRange) Remark(value string) InputDateRange {
	return d.set("remark", value)
}

// Required sets whether the control is required
func (d InputDateRange) Required(value bool) InputDateRange {
	return d.set("required", value)
}

// Row sets the number of rows
func (d InputDateRange) Row(value string) InputDateRange {
	return d.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (d InputDateRange) SaveImmediately(value bool) InputDateRange {
	return d.set("saveImmediately", value)
}

// Shortcuts sets the date range shortcuts
func (d InputDateRange) Shortcuts(value string) InputDateRange {
	return d.set("shortcuts", value)
}

// StartPlaceholder sets the start placeholder
func (d InputDateRange) StartPlaceholder(value string) InputDateRange {
	return d.set("startPlaceholder", value)
}

// TimeFormat sets the time format
func (d InputDateRange) TimeFormat(value string) InputDateRange {
	return d.set("timeFormat", value)
}

// UseMobileUI sets whether to use mobile UI
func (d InputDateRange) UseMobileUI(value bool) InputDateRange {
	return d.set("useMobileUI", value)
}
