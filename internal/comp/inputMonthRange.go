package comp

import "github.com/zrcoder/amisgo/model"

// InputMonthRange represents a month range control
type InputMonthRange model.Schema

func NewInputMonthRange() InputMonthRange {
	return InputMonthRange{"type": "input-month-range"}
}

// Animation enables or disables cursor animation
func (m InputMonthRange) Animation(value bool) InputMonthRange {
	return m.set("animation", value)
}

// AutoFill sets auto-fill value
func (m InputMonthRange) AutoFill(value string) InputMonthRange {
	return m.set("autoFill", value)
}

// BorderMode sets border mode: full, half, or none
func (m InputMonthRange) BorderMode(value string) InputMonthRange {
	return m.set("borderMode", value)
}

// ClassName sets container CSS class name
func (m InputMonthRange) ClassName(value string) InputMonthRange {
	return m.set("className", value)
}

// ClearValueOnHidden clears value when hidden
func (m InputMonthRange) ClearValueOnHidden(value bool) InputMonthRange {
	return m.set("clearValueOnHidden", value)
}

// Delimiter sets the delimiter for the range values
func (m InputMonthRange) Delimiter(value string) InputMonthRange {
	return m.set("delimiter", value)
}

// Desc sets description text
func (m InputMonthRange) Desc(value string) InputMonthRange {
	return m.set("desc", value)
}

// Description sets HTML description text
func (m InputMonthRange) Description(value string) InputMonthRange {
	return m.set("description", value)
}

// DescriptionClassName sets CSS class name for description
func (m InputMonthRange) DescriptionClassName(value string) InputMonthRange {
	return m.set("descriptionClassName", value)
}

// Disabled enables or disables the control
func (m InputMonthRange) Disabled(value bool) InputMonthRange {
	return m.set("disabled", value)
}

// DisabledOn sets expression to disable the control
func (m InputMonthRange) DisabledOn(value string) InputMonthRange {
	return m.set("disabledOn", value)
}

// DisplayFormat sets display date format
func (m InputMonthRange) DisplayFormat(value string) InputMonthRange {
	return m.set("displayFormat", value)
}

// EditorSetting sets editor configuration
func (m InputMonthRange) EditorSetting(value string) InputMonthRange {
	return m.set("editorSetting", value)
}

// Embed enables or disables inline mode
func (m InputMonthRange) Embed(value bool) InputMonthRange {
	return m.set("embed", value)
}

// EndPlaceholder sets end date placeholder
func (m InputMonthRange) EndPlaceholder(value string) InputMonthRange {
	return m.set("endPlaceholder", value)
}

// ExtraName sets extra field name for range component
func (m InputMonthRange) ExtraName(value string) InputMonthRange {
	return m.set("extraName", value)
}

// Format sets submission date format
func (m InputMonthRange) Format(value string) InputMonthRange {
	return m.set("format", value)
}

// Hidden hides or shows the control
func (m InputMonthRange) Hidden(value bool) InputMonthRange {
	return m.set("hidden", value)
}

// HiddenOn sets expression to hide the control
func (m InputMonthRange) HiddenOn(value string) InputMonthRange {
	return m.set("hiddenOn", value)
}

// Hint sets input hint text
func (m InputMonthRange) Hint(value string) InputMonthRange {
	return m.set("hint", value)
}

// Horizontal sets horizontal layout configuration
func (m InputMonthRange) Horizontal(value string) InputMonthRange {
	return m.set("horizontal", value)
}

// ID sets unique component ID
func (m InputMonthRange) ID(value string) InputMonthRange {
	return m.set("id", value)
}

// InitAutoFill sets initial auto-fill value
func (m InputMonthRange) InitAutoFill(value string) InputMonthRange {
	return m.set("initAutoFill", value)
}

// Inline enables or disables inline mode
func (m InputMonthRange) Inline(value bool) InputMonthRange {
	return m.set("inline", value)
}

// InputClassName sets input CSS class name
func (m InputMonthRange) InputClassName(value string) InputMonthRange {
	return m.set("inputClassName", value)
}

// InputFormat sets input date format
func (m InputMonthRange) InputFormat(value string) InputMonthRange {
	return m.set("inputFormat", value)
}

// JoinValues joins selected values with delimiter
func (m InputMonthRange) JoinValues(value bool) InputMonthRange {
	return m.set("joinValues", value)
}

// Label sets label text
func (m InputMonthRange) Label(value string) InputMonthRange {
	return m.set("label", value)
}

// LabelAlign sets label alignment: right, left, top, or inherit
func (m InputMonthRange) LabelAlign(value string) InputMonthRange {
	return m.set("labelAlign", value)
}

// LabelClassName sets label CSS class name
func (m InputMonthRange) LabelClassName(value string) InputMonthRange {
	return m.set("labelClassName", value)
}

// LabelRemark sets label remark text
func (m InputMonthRange) LabelRemark(value string) InputMonthRange {
	return m.set("labelRemark", value)
}

// LabelWidth sets label width
func (m InputMonthRange) LabelWidth(value string) InputMonthRange {
	return m.set("labelWidth", value)
}

// MaxDate sets maximum date limit
func (m InputMonthRange) MaxDate(value string) InputMonthRange {
	return m.set("maxDate", value)
}

// MaxDuration sets maximum duration
func (m InputMonthRange) MaxDuration(value string) InputMonthRange {
	return m.set("maxDuration", value)
}

// MinDate sets minimum date limit
func (m InputMonthRange) MinDate(value string) InputMonthRange {
	return m.set("minDate", value)
}

// MinDuration sets minimum duration
func (m InputMonthRange) MinDuration(value string) InputMonthRange {
	return m.set("minDuration", value)
}

// Mode sets display mode: normal, inline, or horizontal
func (m InputMonthRange) Mode(value string) InputMonthRange {
	return m.set("mode", value)
}

// Name sets field name for form submission
func (m InputMonthRange) Name(value string) InputMonthRange {
	return m.set("name", value)
}

// OnEvent sets event action configuration
func (m InputMonthRange) OnEvent(value any) InputMonthRange {
	return m.set("onEvent", value)
}

// Placeholder sets placeholder text
func (m InputMonthRange) Placeholder(value string) InputMonthRange {
	return m.set("placeholder", value)
}

// PopOverContainerSelector sets popover container selector
func (m InputMonthRange) PopOverContainerSelector(value string) InputMonthRange {
	return m.set("popOverContainerSelector", value)
}

// Ranges sets date range shortcuts
func (m InputMonthRange) Ranges(value string) InputMonthRange {
	return m.set("ranges", value)
}

// ReadOnly sets read-only mode
func (m InputMonthRange) ReadOnly(value bool) InputMonthRange {
	return m.set("readOnly", value)
}

// ReadOnlyOn sets expression for read-only mode
func (m InputMonthRange) ReadOnlyOn(value string) InputMonthRange {
	return m.set("readOnlyOn", value)
}

// Remark sets remark text
func (m InputMonthRange) Remark(value string) InputMonthRange {
	return m.set("remark", value)
}

// Required sets required field
func (m InputMonthRange) Required(value bool) InputMonthRange {
	return m.set("required", value)
}

// Row sets row configuration
func (m InputMonthRange) Row(value string) InputMonthRange {
	return m.set("row", value)
}

// SaveImmediately enables or disables immediate save
func (m InputMonthRange) SaveImmediately(value bool) InputMonthRange {
	return m.set("saveImmediately", value)
}

// Shortcuts sets date range shortcuts
func (m InputMonthRange) Shortcuts(value string) InputMonthRange {
	return m.set("shortcuts", value)
}

// Size sets form item size: xs, sm, md, lg, or full
func (m InputMonthRange) Size(value string) InputMonthRange {
	return m.set("size", value)
}

// StartPlaceholder sets start date placeholder
func (m InputMonthRange) StartPlaceholder(value string) InputMonthRange {
	return m.set("startPlaceholder", value)
}

// Static enables or disables static display
func (m InputMonthRange) Static(value bool) InputMonthRange {
	return m.set("static", value)
}

// StaticClassName sets static display CSS class name
func (m InputMonthRange) StaticClassName(value string) InputMonthRange {
	return m.set("staticClassName", value)
}

// StaticInputClassName sets static input CSS class name
func (m InputMonthRange) StaticInputClassName(value string) InputMonthRange {
	return m.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label CSS class name
func (m InputMonthRange) StaticLabelClassName(value string) InputMonthRange {
	return m.set("staticLabelClassName", value)
}

// StaticOn sets expression for static display
func (m InputMonthRange) StaticOn(value string) InputMonthRange {
	return m.set("staticOn", value)
}

// StaticPlaceholder sets static display placeholder
func (m InputMonthRange) StaticPlaceholder(value string) InputMonthRange {
	return m.set("staticPlaceholder", value)
}

// StaticSchema sets static display schema
func (m InputMonthRange) StaticSchema(value string) InputMonthRange {
	return m.set("staticSchema", value)
}

// Style sets component style
func (m InputMonthRange) Style(value any) InputMonthRange {
	return m.set("style", value)
}

// SubmitOnChange enables or disables form submission on change
func (m InputMonthRange) SubmitOnChange(value bool) InputMonthRange {
	return m.set("submitOnChange", value)
}

// TestIdBuilder sets test ID builder
func (m InputMonthRange) TestIdBuilder(value string) InputMonthRange {
	return m.set("testIdBuilder", value)
}

// Transform sets date transformation function
func (m InputMonthRange) Transform(value string) InputMonthRange {
	return m.set("transform", value)
}

// UseMobileUI enables or disables mobile UI
func (m InputMonthRange) UseMobileUI(value bool) InputMonthRange {
	return m.set("useMobileUI", value)
}

// ValidateApi sets remote validation API
func (m InputMonthRange) ValidateApi(value string) InputMonthRange {
	return m.set("validateApi", value)
}

// ValidateOnChange enables or disables validation on change
func (m InputMonthRange) ValidateOnChange(value bool) InputMonthRange {
	return m.set("validateOnChange", value)
}

// ValidationErrors sets validation error messages
func (m InputMonthRange) ValidationErrors(value string) InputMonthRange {
	return m.set("validationErrors", value)
}

// Validations sets validation rules
func (m InputMonthRange) Validations(value string) InputMonthRange {
	return m.set("validations", value)
}

// Value sets the value, supporting relative values
func (m InputMonthRange) Value(value string) InputMonthRange {
	return m.set("value", value)
}

// ValueFormat sets submission date format
func (m InputMonthRange) ValueFormat(value string) InputMonthRange {
	return m.set("valueFormat", value)
}

// Visible shows or hides the control
func (m InputMonthRange) Visible(value bool) InputMonthRange {
	return m.set("visible", value)
}

// VisibleOn sets expression to show the control
func (m InputMonthRange) VisibleOn(value string) InputMonthRange {
	return m.set("visibleOn", value)
}

// Width sets width in table
func (m InputMonthRange) Width(value string) InputMonthRange {
	return m.set("width", value)
}

func (m InputMonthRange) set(key string, value any) InputMonthRange {
	m[key] = value
	return m
}
