package comp

// inputQuarterRange represents a quarter range control.
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-quarter-range
type inputQuarterRange Schema

// InputQuarterRange creates a new QuarterRangeControl instance
func InputQuarterRange() inputQuarterRange {
	return inputQuarterRange{}.set("type", "input-quarter-range")
}

// set sets a field value
func (qrc inputQuarterRange) set(key string, value any) inputQuarterRange {
	qrc[key] = value
	return qrc
}

// Animation enables or disables cursor animation, default is enabled
func (qrc inputQuarterRange) Animation(value bool) inputQuarterRange {
	return qrc.set("animation", value)
}

// AutoFill auto-fills other values in the form when an option is selected
func (qrc inputQuarterRange) AutoFill(value string) inputQuarterRange {
	return qrc.set("autoFill", value)
}

// BorderMode sets the border mode: full, half, or none
func (qrc inputQuarterRange) BorderMode(value string) inputQuarterRange {
	return qrc.set("borderMode", value)
}

// ClassName sets the container CSS class name
func (qrc inputQuarterRange) ClassName(value string) inputQuarterRange {
	return qrc.set("className", value)
}

// ClearValueOnHidden removes the form item value when hidden
func (qrc inputQuarterRange) ClearValueOnHidden(value bool) inputQuarterRange {
	return qrc.set("clearValueOnHidden", value)
}

// Delimiter sets the delimiter for the range, default is a comma
func (qrc inputQuarterRange) Delimiter(value string) inputQuarterRange {
	return qrc.set("delimiter", value)
}

// Desc sets the description
func (qrc inputQuarterRange) Desc(value string) inputQuarterRange {
	return qrc.set("desc", value)
}

// Description sets the description content, supports HTML
func (qrc inputQuarterRange) Description(value string) inputQuarterRange {
	return qrc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (qrc inputQuarterRange) DescriptionClassName(value string) inputQuarterRange {
	return qrc.set("descriptionClassName", value)
}

// Disabled enables or disables the control
func (qrc inputQuarterRange) Disabled(value bool) inputQuarterRange {
	return qrc.set("disabled", value)
}

// DisabledOn sets the expression to disable the control
func (qrc inputQuarterRange) DisabledOn(value string) inputQuarterRange {
	return qrc.set("disabledOn", value)
}

// DisplayFormat sets the display format for the time
func (qrc inputQuarterRange) DisplayFormat(value string) inputQuarterRange {
	return qrc.set("displayFormat", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (qrc inputQuarterRange) EditorSetting(value string) inputQuarterRange {
	return qrc.set("editorSetting", value)
}

// Embed enables inline mode
func (qrc inputQuarterRange) Embed(value bool) inputQuarterRange {
	return qrc.set("embed", value)
}

// EndPlaceholder sets the placeholder for the end date
func (qrc inputQuarterRange) EndPlaceholder(value string) inputQuarterRange {
	return qrc.set("endPlaceholder", value)
}

// ExtraName sets an extra field name for range components
func (qrc inputQuarterRange) ExtraName(value string) inputQuarterRange {
	return qrc.set("extraName", value)
}

// Format sets the format for the submitted time, default is timestamp
func (qrc inputQuarterRange) Format(value string) inputQuarterRange {
	return qrc.set("format", value)
}

// Hidden hides the control
func (qrc inputQuarterRange) Hidden(value bool) inputQuarterRange {
	return qrc.set("hidden", value)
}

// HiddenOn sets the expression to hide the control
func (qrc inputQuarterRange) HiddenOn(value string) inputQuarterRange {
	return qrc.set("hiddenOn", value)
}

// Hint sets the input hint, shown on focus
func (qrc inputQuarterRange) Hint(value string) inputQuarterRange {
	return qrc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (qrc inputQuarterRange) Horizontal(value string) inputQuarterRange {
	return qrc.set("horizontal", value)
}

// ID sets the unique component ID, mainly for logging
func (qrc inputQuarterRange) ID(value string) inputQuarterRange {
	return qrc.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (qrc inputQuarterRange) InitAutoFill(value string) inputQuarterRange {
	return qrc.set("initAutoFill", value)
}

// Inline sets the control to inline mode
func (qrc inputQuarterRange) Inline(value bool) inputQuarterRange {
	return qrc.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (qrc inputQuarterRange) InputClassName(value string) inputQuarterRange {
	return qrc.set("inputClassName", value)
}

// InputFormat sets the display format for the input time, default is YYYY-MM-DD
func (qrc inputQuarterRange) InputFormat(value string) inputQuarterRange {
	return qrc.set("inputFormat", value)
}

// JoinValues joins selected option values with a delimiter
func (qrc inputQuarterRange) JoinValues(value bool) inputQuarterRange {
	return qrc.set("joinValues", value)
}

// Label sets the label text
func (qrc inputQuarterRange) Label(value string) inputQuarterRange {
	return qrc.set("label", value)
}

// LabelAlign sets the label alignment: right, left, top, or inherit
func (qrc inputQuarterRange) LabelAlign(value string) inputQuarterRange {
	return qrc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (qrc inputQuarterRange) LabelClassName(value string) inputQuarterRange {
	return qrc.set("labelClassName", value)
}

// LabelRemark sets the label remark, shown on hover
func (qrc inputQuarterRange) LabelRemark(value string) inputQuarterRange {
	return qrc.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label, default is px
func (qrc inputQuarterRange) LabelWidth(value string) inputQuarterRange {
	return qrc.set("labelWidth", value)
}

// MaxDate sets the maximum date limit, supports variables and relative values
func (qrc inputQuarterRange) MaxDate(value string) inputQuarterRange {
	return qrc.set("maxDate", value)
}

// MaxDuration sets the maximum duration, e.g., 2 days
func (qrc inputQuarterRange) MaxDuration(value string) inputQuarterRange {
	return qrc.set("maxDuration", value)
}

// MinDate sets the minimum date limit, supports variables and relative values
func (qrc inputQuarterRange) MinDate(value string) inputQuarterRange {
	return qrc.set("minDate", value)
}

// MinDuration sets the minimum duration, e.g., 2 days
func (qrc inputQuarterRange) MinDuration(value string) inputQuarterRange {
	return qrc.set("minDuration", value)
}

// Mode sets the display mode: normal, inline, or horizontal
func (qrc inputQuarterRange) Mode(value string) inputQuarterRange {
	return qrc.set("mode", value)
}

// Name sets the field name for form submission, supports nested keys
func (qrc inputQuarterRange) Name(value string) inputQuarterRange {
	return qrc.set("name", value)
}

// OnEvent sets the event action configuration
func (qrc inputQuarterRange) OnEvent(value any) inputQuarterRange {
	return qrc.set("onEvent", value)
}

// Placeholder sets the placeholder text
func (qrc inputQuarterRange) Placeholder(value string) inputQuarterRange {
	return qrc.set("placeholder", value)
}

// PopOverContainerSelector sets the popover container selector
func (qrc inputQuarterRange) PopOverContainerSelector(value string) inputQuarterRange {
	return qrc.set("popOverContainerSelector", value)
}

// Ranges sets the date range shortcuts
func (qrc inputQuarterRange) Ranges(value string) inputQuarterRange {
	return qrc.set("ranges", value)
}

// ReadOnly sets the control to read-only mode
func (qrc inputQuarterRange) ReadOnly(value bool) inputQuarterRange {
	return qrc.set("readOnly", value)
}

// ReadOnlyOn sets the expression to make the control read-only
func (qrc inputQuarterRange) ReadOnlyOn(value string) inputQuarterRange {
	return qrc.set("readOnlyOn", value)
}

// Remark sets the remark, shown on hover
func (qrc inputQuarterRange) Remark(value string) inputQuarterRange {
	return qrc.set("remark", value)
}

// Required sets the control as required
func (qrc inputQuarterRange) Required(value bool) inputQuarterRange {
	return qrc.set("required", value)
}

// Row sets the row value
func (qrc inputQuarterRange) Row(value string) inputQuarterRange {
	return qrc.set("row", value)
}

// SaveImmediately sets whether to save immediately in TableColumn
func (qrc inputQuarterRange) SaveImmediately(value bool) inputQuarterRange {
	return qrc.set("saveImmediately", value)
}

// Shortcuts sets the date range shortcuts
func (qrc inputQuarterRange) Shortcuts(value string) inputQuarterRange {
	return qrc.set("shortcuts", value)
}

// Size sets the control size: xs, sm, md, lg, or full
func (qrc inputQuarterRange) Size(value string) inputQuarterRange {
	return qrc.set("size", value)
}

// StartPlaceholder sets the placeholder for the start date
func (qrc inputQuarterRange) StartPlaceholder(value string) inputQuarterRange {
	return qrc.set("startPlaceholder", value)
}

// Static sets the control to static display mode
func (qrc inputQuarterRange) Static(value bool) inputQuarterRange {
	return qrc.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (qrc inputQuarterRange) StaticClassName(value string) inputQuarterRange {
	return qrc.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (qrc inputQuarterRange) StaticInputClassName(value string) inputQuarterRange {
	return qrc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (qrc inputQuarterRange) StaticLabelClassName(value string) inputQuarterRange {
	return qrc.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (qrc inputQuarterRange) StaticOn(value string) inputQuarterRange {
	return qrc.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (qrc inputQuarterRange) StaticPlaceholder(value string) inputQuarterRange {
	return qrc.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (qrc inputQuarterRange) StaticSchema(value string) inputQuarterRange {
	return qrc.set("staticSchema", value)
}

// Style sets the component style
func (qrc inputQuarterRange) Style(value any) inputQuarterRange {
	return qrc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change
func (qrc inputQuarterRange) SubmitOnChange(value bool) inputQuarterRange {
	return qrc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (qrc inputQuarterRange) TestIdBuilder(value string) inputQuarterRange {
	return qrc.set("testIdBuilder", value)
}

// Transform sets the date transformation function
func (qrc inputQuarterRange) Transform(value string) inputQuarterRange {
	return qrc.set("transform", value)
}

// UseMobileUI enables or disables mobile UI styles
func (qrc inputQuarterRange) UseMobileUI(value bool) inputQuarterRange {
	return qrc.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API
func (qrc inputQuarterRange) ValidateApi(value string) inputQuarterRange {
	return qrc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (qrc inputQuarterRange) ValidateOnChange(value bool) inputQuarterRange {
	return qrc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (qrc inputQuarterRange) ValidationErrors(value string) inputQuarterRange {
	return qrc.set("validationErrors", value)
}

// Validations sets the validation rules
func (qrc inputQuarterRange) Validations(value string) inputQuarterRange {
	return qrc.set("validations", value)
}

// Value sets the value, supports relative values
func (qrc inputQuarterRange) Value(value string) inputQuarterRange {
	return qrc.set("value", value)
}

// ValueFormat sets the format for the submitted time
func (qrc inputQuarterRange) ValueFormat(value string) inputQuarterRange {
	return qrc.set("valueFormat", value)
}

// Visible sets the control visibility
func (qrc inputQuarterRange) Visible(value bool) inputQuarterRange {
	return qrc.set("visible", value)
}

// VisibleOn sets the expression for control visibility
func (qrc inputQuarterRange) VisibleOn(value string) inputQuarterRange {
	return qrc.set("visibleOn", value)
}

// Width sets the width in Table
func (qrc inputQuarterRange) Width(value string) inputQuarterRange {
	return qrc.set("width", value)
}
