package comp

// inputDateRange represents a date range input control
type inputDateRange Schema

// InputDateRange creates a new InputDateRange instance with default type
func InputDateRange() inputDateRange {
	return make(inputDateRange).set("type", "input-date-range")
}

// set sets a key-value pair and returns the current instance
func (d inputDateRange) set(key string, value any) inputDateRange {
	d[key] = value
	return d
}

// Animation sets the animation effect
func (d inputDateRange) Animation(value bool) inputDateRange {
	return d.set("animation", value)
}

// AutoFill sets the auto-fill value
func (d inputDateRange) AutoFill(value string) inputDateRange {
	return d.set("autoFill", value)
}

// BorderMode sets the border mode
func (d inputDateRange) BorderMode(value string) inputDateRange {
	return d.set("borderMode", value)
}

// ClassName sets the class name
func (d inputDateRange) ClassName(value string) inputDateRange {
	return d.set("className", value)
}

// ClearValueOnHidden sets whether to clear value when hidden
func (d inputDateRange) ClearValueOnHidden(value bool) inputDateRange {
	return d.set("clearValueOnHidden", value)
}

// Delimiter sets the delimiter
func (d inputDateRange) Delimiter(value string) inputDateRange {
	return d.set("delimiter", value)
}

// Desc sets the description
func (d inputDateRange) Desc(value string) inputDateRange {
	return d.set("desc", value)
}

// Description sets the description
func (d inputDateRange) Description(value string) inputDateRange {
	return d.set("description", value)
}

// DescriptionClassName sets the description class name
func (d inputDateRange) DescriptionClassName(value string) inputDateRange {
	return d.set("descriptionClassName", value)
}

// Disabled sets whether the control is disabled
func (d inputDateRange) Disabled(value bool) inputDateRange {
	return d.set("disabled", value)
}

// DisabledOn sets the condition for disabling the control
func (d inputDateRange) DisabledOn(value string) inputDateRange {
	return d.set("disabledOn", value)
}

// DisplayFormat sets the display format
func (d inputDateRange) DisplayFormat(value string) inputDateRange {
	return d.set("displayFormat", value)
}

// EditorSetting sets the editor configuration
func (d inputDateRange) EditorSetting(value string) inputDateRange {
	return d.set("editorSetting", value)
}

// Embed sets whether to use inline mode
func (d inputDateRange) Embed(value bool) inputDateRange {
	return d.set("embed", value)
}

// EndPlaceholder sets the end placeholder
func (d inputDateRange) EndPlaceholder(value string) inputDateRange {
	return d.set("endPlaceholder", value)
}

// ExtraName sets the extra field name
func (d inputDateRange) ExtraName(value string) inputDateRange {
	return d.set("extraName", value)
}

// Format sets the submission format
func (d inputDateRange) Format(value string) inputDateRange {
	return d.set("format", value)
}

// Hidden sets whether the control is hidden
func (d inputDateRange) Hidden(value bool) inputDateRange {
	return d.set("hidden", value)
}

// HiddenOn sets the condition for hiding the control
func (d inputDateRange) HiddenOn(value string) inputDateRange {
	return d.set("hiddenOn", value)
}

// Hint sets the input hint
func (d inputDateRange) Hint(value string) inputDateRange {
	return d.set("hint", value)
}

// Horizontal sets the horizontal layout
func (d inputDateRange) Horizontal(value string) inputDateRange {
	return d.set("horizontal", value)
}

// ID sets the component ID
func (d inputDateRange) ID(value string) inputDateRange {
	return d.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (d inputDateRange) InitAutoFill(value string) inputDateRange {
	return d.set("initAutoFill", value)
}

// Inline sets whether to use inline mode
func (d inputDateRange) Inline(value bool) inputDateRange {
	return d.set("inline", value)
}

// Value sets the value
func (d inputDateRange) Value(value string) inputDateRange {
	return d.set("value", value)
}

// CloseOnSelect sets whether to close on select
func (d inputDateRange) CloseOnSelect(value bool) inputDateRange {
	return d.set("closeOnSelect", value)
}

// Clearable sets whether the control is clearable
func (d inputDateRange) Clearable(value bool) inputDateRange {
	return d.set("clearable", value)
}

// InputClassName sets the input class name
func (d inputDateRange) InputClassName(value string) inputDateRange {
	return d.set("inputClassName", value)
}

// InputFormat sets the input format
func (d inputDateRange) InputFormat(value string) inputDateRange {
	return d.set("inputFormat", value)
}

// JoinValues sets whether to join values into a single string
func (d inputDateRange) JoinValues(value bool) inputDateRange {
	return d.set("joinValues", value)
}

// Label sets the control label
func (d inputDateRange) Label(value string) inputDateRange {
	return d.set("label", value)
}

// LabelAlign sets the label alignment
func (d inputDateRange) LabelAlign(value string) inputDateRange {
	return d.set("labelAlign", value)
}

// LabelClassName sets the label class name
func (d inputDateRange) LabelClassName(value string) inputDateRange {
	return d.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (d inputDateRange) LabelRemark(value string) inputDateRange {
	return d.set("labelRemark", value)
}

// LabelWidth sets the label width
func (d inputDateRange) LabelWidth(value string) inputDateRange {
	return d.set("labelWidth", value)
}

// MaxDate sets the maximum date limit
func (d inputDateRange) MaxDate(value string) inputDateRange {
	return d.set("maxDate", value)
}

// MaxDuration sets the maximum duration
func (d inputDateRange) MaxDuration(value string) inputDateRange {
	return d.set("maxDuration", value)
}

// MinDate sets the minimum date limit
func (d inputDateRange) MinDate(value string) inputDateRange {
	return d.set("minDate", value)
}

// MinDuration sets the minimum duration
func (d inputDateRange) MinDuration(value string) inputDateRange {
	return d.set("minDuration", value)
}

// Mode sets the display mode
func (d inputDateRange) Mode(value string) inputDateRange {
	return d.set("mode", value)
}

// Name sets the form submission field name
func (d inputDateRange) Name(value string) inputDateRange {
	return d.set("name", value)
}

// OnEvent sets the event configuration
func (d inputDateRange) OnEvent(value any) inputDateRange {
	return d.set("onEvent", value)
}

// Placeholder sets the placeholder text
func (d inputDateRange) Placeholder(value string) inputDateRange {
	return d.set("placeholder", value)
}

// PopOverContainerSelector sets the popover container selector
func (d inputDateRange) PopOverContainerSelector(value string) inputDateRange {
	return d.set("popOverContainerSelector", value)
}

// Ranges sets the date range shortcuts
func (d inputDateRange) Ranges(value string) inputDateRange {
	return d.set("ranges", value)
}

// ReadOnly sets whether the control is read-only
func (d inputDateRange) ReadOnly(value bool) inputDateRange {
	return d.set("readOnly", value)
}

// ReadOnlyOn sets the condition for read-only
func (d inputDateRange) ReadOnlyOn(value string) inputDateRange {
	return d.set("readOnlyOn", value)
}

// Remark sets the control remark
func (d inputDateRange) Remark(value string) inputDateRange {
	return d.set("remark", value)
}

// Required sets whether the control is required
func (d inputDateRange) Required(value bool) inputDateRange {
	return d.set("required", value)
}

// Row sets the number of rows
func (d inputDateRange) Row(value string) inputDateRange {
	return d.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (d inputDateRange) SaveImmediately(value bool) inputDateRange {
	return d.set("saveImmediately", value)
}

// Shortcuts sets the date range shortcuts
func (d inputDateRange) Shortcuts(value string) inputDateRange {
	return d.set("shortcuts", value)
}

// StartPlaceholder sets the start placeholder
func (d inputDateRange) StartPlaceholder(value string) inputDateRange {
	return d.set("startPlaceholder", value)
}

// TimeFormat sets the time format
func (d inputDateRange) TimeFormat(value string) inputDateRange {
	return d.set("timeFormat", value)
}

// UseMobileUI sets whether to use mobile UI
func (d inputDateRange) UseMobileUI(value bool) inputDateRange {
	return d.set("useMobileUI", value)
}
