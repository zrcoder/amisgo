package comp

import "github.com/zrcoder/amisgo/model"

// InputQuarterRange represents a quarter range control.
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-quarter-range
type InputQuarterRange model.Schema

func NewInputQuarterRange() InputQuarterRange {
	return InputQuarterRange{"type": "input-quarter-range"}
}

// set sets a field value
func (qrc InputQuarterRange) set(key string, value any) InputQuarterRange {
	qrc[key] = value
	return qrc
}

// Animation enables or disables cursor animation, default is enabled
func (qrc InputQuarterRange) Animation(value bool) InputQuarterRange {
	return qrc.set("animation", value)
}

// AutoFill auto-fills other values in the form when an option is selected
func (qrc InputQuarterRange) AutoFill(value string) InputQuarterRange {
	return qrc.set("autoFill", value)
}

// BorderMode sets the border mode: full, half, or none
func (qrc InputQuarterRange) BorderMode(value string) InputQuarterRange {
	return qrc.set("borderMode", value)
}

// ClassName sets the container CSS class name
func (qrc InputQuarterRange) ClassName(value string) InputQuarterRange {
	return qrc.set("className", value)
}

// ClearValueOnHidden removes the form item value when hidden
func (qrc InputQuarterRange) ClearValueOnHidden(value bool) InputQuarterRange {
	return qrc.set("clearValueOnHidden", value)
}

// Delimiter sets the delimiter for the range, default is a comma
func (qrc InputQuarterRange) Delimiter(value string) InputQuarterRange {
	return qrc.set("delimiter", value)
}

// Desc sets the description
func (qrc InputQuarterRange) Desc(value string) InputQuarterRange {
	return qrc.set("desc", value)
}

// Description sets the description content, supports HTML
func (qrc InputQuarterRange) Description(value string) InputQuarterRange {
	return qrc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (qrc InputQuarterRange) DescriptionClassName(value string) InputQuarterRange {
	return qrc.set("descriptionClassName", value)
}

// Disabled enables or disables the control
func (qrc InputQuarterRange) Disabled(value bool) InputQuarterRange {
	return qrc.set("disabled", value)
}

// DisabledOn sets the expression to disable the control
func (qrc InputQuarterRange) DisabledOn(value string) InputQuarterRange {
	return qrc.set("disabledOn", value)
}

// DisplayFormat sets the display format for the time
func (qrc InputQuarterRange) DisplayFormat(value string) InputQuarterRange {
	return qrc.set("displayFormat", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (qrc InputQuarterRange) EditorSetting(value string) InputQuarterRange {
	return qrc.set("editorSetting", value)
}

// Embed enables inline mode
func (qrc InputQuarterRange) Embed(value bool) InputQuarterRange {
	return qrc.set("embed", value)
}

// EndPlaceholder sets the placeholder for the end date
func (qrc InputQuarterRange) EndPlaceholder(value string) InputQuarterRange {
	return qrc.set("endPlaceholder", value)
}

// ExtraName sets an extra field name for range components
func (qrc InputQuarterRange) ExtraName(value string) InputQuarterRange {
	return qrc.set("extraName", value)
}

// Format sets the format for the submitted time, default is timestamp
func (qrc InputQuarterRange) Format(value string) InputQuarterRange {
	return qrc.set("format", value)
}

// Hidden hides the control
func (qrc InputQuarterRange) Hidden(value bool) InputQuarterRange {
	return qrc.set("hidden", value)
}

// HiddenOn sets the expression to hide the control
func (qrc InputQuarterRange) HiddenOn(value string) InputQuarterRange {
	return qrc.set("hiddenOn", value)
}

// Hint sets the input hint, shown on focus
func (qrc InputQuarterRange) Hint(value string) InputQuarterRange {
	return qrc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (qrc InputQuarterRange) Horizontal(value string) InputQuarterRange {
	return qrc.set("horizontal", value)
}

// ID sets the unique component ID, mainly for logging
func (qrc InputQuarterRange) ID(value string) InputQuarterRange {
	return qrc.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (qrc InputQuarterRange) InitAutoFill(value string) InputQuarterRange {
	return qrc.set("initAutoFill", value)
}

// Inline sets the control to inline mode
func (qrc InputQuarterRange) Inline(value bool) InputQuarterRange {
	return qrc.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (qrc InputQuarterRange) InputClassName(value string) InputQuarterRange {
	return qrc.set("inputClassName", value)
}

// InputFormat sets the display format for the input time, default is YYYY-MM-DD
func (qrc InputQuarterRange) InputFormat(value string) InputQuarterRange {
	return qrc.set("inputFormat", value)
}

// JoinValues joins selected option values with a delimiter
func (qrc InputQuarterRange) JoinValues(value bool) InputQuarterRange {
	return qrc.set("joinValues", value)
}

// Label sets the label text
func (qrc InputQuarterRange) Label(value string) InputQuarterRange {
	return qrc.set("label", value)
}

// LabelAlign sets the label alignment: right, left, top, or inherit
func (qrc InputQuarterRange) LabelAlign(value string) InputQuarterRange {
	return qrc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (qrc InputQuarterRange) LabelClassName(value string) InputQuarterRange {
	return qrc.set("labelClassName", value)
}

// LabelRemark sets the label remark, shown on hover
func (qrc InputQuarterRange) LabelRemark(value string) InputQuarterRange {
	return qrc.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label, default is px
func (qrc InputQuarterRange) LabelWidth(value string) InputQuarterRange {
	return qrc.set("labelWidth", value)
}

// MaxDate sets the maximum date limit, supports variables and relative values
func (qrc InputQuarterRange) MaxDate(value string) InputQuarterRange {
	return qrc.set("maxDate", value)
}

// MaxDuration sets the maximum duration, e.g., 2 days
func (qrc InputQuarterRange) MaxDuration(value string) InputQuarterRange {
	return qrc.set("maxDuration", value)
}

// MinDate sets the minimum date limit, supports variables and relative values
func (qrc InputQuarterRange) MinDate(value string) InputQuarterRange {
	return qrc.set("minDate", value)
}

// MinDuration sets the minimum duration, e.g., 2 days
func (qrc InputQuarterRange) MinDuration(value string) InputQuarterRange {
	return qrc.set("minDuration", value)
}

// Mode sets the display mode: normal, inline, or horizontal
func (qrc InputQuarterRange) Mode(value string) InputQuarterRange {
	return qrc.set("mode", value)
}

// Name sets the field name for form submission, supports nested keys
func (qrc InputQuarterRange) Name(value string) InputQuarterRange {
	return qrc.set("name", value)
}

// OnEvent sets the event action configuration
func (qrc InputQuarterRange) OnEvent(value any) InputQuarterRange {
	return qrc.set("onEvent", value)
}

// Placeholder sets the placeholder text
func (qrc InputQuarterRange) Placeholder(value string) InputQuarterRange {
	return qrc.set("placeholder", value)
}

// PopOverContainerSelector sets the popover container selector
func (qrc InputQuarterRange) PopOverContainerSelector(value string) InputQuarterRange {
	return qrc.set("popOverContainerSelector", value)
}

// Ranges sets the date range shortcuts
func (qrc InputQuarterRange) Ranges(value string) InputQuarterRange {
	return qrc.set("ranges", value)
}

// ReadOnly sets the control to read-only mode
func (qrc InputQuarterRange) ReadOnly(value bool) InputQuarterRange {
	return qrc.set("readOnly", value)
}

// ReadOnlyOn sets the expression to make the control read-only
func (qrc InputQuarterRange) ReadOnlyOn(value string) InputQuarterRange {
	return qrc.set("readOnlyOn", value)
}

// Remark sets the remark, shown on hover
func (qrc InputQuarterRange) Remark(value string) InputQuarterRange {
	return qrc.set("remark", value)
}

// Required sets the control as required
func (qrc InputQuarterRange) Required(value bool) InputQuarterRange {
	return qrc.set("required", value)
}

// Row sets the row value
func (qrc InputQuarterRange) Row(value string) InputQuarterRange {
	return qrc.set("row", value)
}

// SaveImmediately sets whether to save immediately in TableColumn
func (qrc InputQuarterRange) SaveImmediately(value bool) InputQuarterRange {
	return qrc.set("saveImmediately", value)
}

// Shortcuts sets the date range shortcuts
func (qrc InputQuarterRange) Shortcuts(value string) InputQuarterRange {
	return qrc.set("shortcuts", value)
}

// Size sets the control size: xs, sm, md, lg, or full
func (qrc InputQuarterRange) Size(value string) InputQuarterRange {
	return qrc.set("size", value)
}

// StartPlaceholder sets the placeholder for the start date
func (qrc InputQuarterRange) StartPlaceholder(value string) InputQuarterRange {
	return qrc.set("startPlaceholder", value)
}

// Static sets the control to static display mode
func (qrc InputQuarterRange) Static(value bool) InputQuarterRange {
	return qrc.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (qrc InputQuarterRange) StaticClassName(value string) InputQuarterRange {
	return qrc.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (qrc InputQuarterRange) StaticInputClassName(value string) InputQuarterRange {
	return qrc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (qrc InputQuarterRange) StaticLabelClassName(value string) InputQuarterRange {
	return qrc.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (qrc InputQuarterRange) StaticOn(value string) InputQuarterRange {
	return qrc.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (qrc InputQuarterRange) StaticPlaceholder(value string) InputQuarterRange {
	return qrc.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (qrc InputQuarterRange) StaticSchema(value string) InputQuarterRange {
	return qrc.set("staticSchema", value)
}

// Style sets the component style
func (qrc InputQuarterRange) Style(value any) InputQuarterRange {
	return qrc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change
func (qrc InputQuarterRange) SubmitOnChange(value bool) InputQuarterRange {
	return qrc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (qrc InputQuarterRange) TestIdBuilder(value string) InputQuarterRange {
	return qrc.set("testIdBuilder", value)
}

// Transform sets the date transformation function
func (qrc InputQuarterRange) Transform(value string) InputQuarterRange {
	return qrc.set("transform", value)
}

// UseMobileUI enables or disables mobile UI styles
func (qrc InputQuarterRange) UseMobileUI(value bool) InputQuarterRange {
	return qrc.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API
func (qrc InputQuarterRange) ValidateApi(value string) InputQuarterRange {
	return qrc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (qrc InputQuarterRange) ValidateOnChange(value bool) InputQuarterRange {
	return qrc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (qrc InputQuarterRange) ValidationErrors(value string) InputQuarterRange {
	return qrc.set("validationErrors", value)
}

// Validations sets the validation rules
func (qrc InputQuarterRange) Validations(value string) InputQuarterRange {
	return qrc.set("validations", value)
}

// Value sets the value, supports relative values
func (qrc InputQuarterRange) Value(value string) InputQuarterRange {
	return qrc.set("value", value)
}

// ValueFormat sets the format for the submitted time
func (qrc InputQuarterRange) ValueFormat(value string) InputQuarterRange {
	return qrc.set("valueFormat", value)
}

// Visible sets the control visibility
func (qrc InputQuarterRange) Visible(value bool) InputQuarterRange {
	return qrc.set("visible", value)
}

// VisibleOn sets the expression for control visibility
func (qrc InputQuarterRange) VisibleOn(value string) InputQuarterRange {
	return qrc.set("visibleOn", value)
}

// Width sets the width in Table
func (qrc InputQuarterRange) Width(value string) InputQuarterRange {
	return qrc.set("width", value)
}
