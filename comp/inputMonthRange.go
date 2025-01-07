package comp

// inputMonthRange represents a month range control
type inputMonthRange Schema

// InputMonthRange creates a new MonthRangeControl instance
func InputMonthRange() inputMonthRange {
	return inputMonthRange{}.set("type", "input-month-range")
}

// Animation enables or disables cursor animation
func (m inputMonthRange) Animation(value bool) inputMonthRange {
	return m.set("animation", value)
}

// AutoFill sets auto-fill value
func (m inputMonthRange) AutoFill(value string) inputMonthRange {
	return m.set("autoFill", value)
}

// BorderMode sets border mode: full, half, or none
func (m inputMonthRange) BorderMode(value string) inputMonthRange {
	return m.set("borderMode", value)
}

// ClassName sets container CSS class name
func (m inputMonthRange) ClassName(value string) inputMonthRange {
	return m.set("className", value)
}

// ClearValueOnHidden clears value when hidden
func (m inputMonthRange) ClearValueOnHidden(value bool) inputMonthRange {
	return m.set("clearValueOnHidden", value)
}

// Delimiter sets the delimiter for the range values
func (m inputMonthRange) Delimiter(value string) inputMonthRange {
	return m.set("delimiter", value)
}

// Desc sets description text
func (m inputMonthRange) Desc(value string) inputMonthRange {
	return m.set("desc", value)
}

// Description sets HTML description text
func (m inputMonthRange) Description(value string) inputMonthRange {
	return m.set("description", value)
}

// DescriptionClassName sets CSS class name for description
func (m inputMonthRange) DescriptionClassName(value string) inputMonthRange {
	return m.set("descriptionClassName", value)
}

// Disabled enables or disables the control
func (m inputMonthRange) Disabled(value bool) inputMonthRange {
	return m.set("disabled", value)
}

// DisabledOn sets expression to disable the control
func (m inputMonthRange) DisabledOn(value string) inputMonthRange {
	return m.set("disabledOn", value)
}

// DisplayFormat sets display date format
func (m inputMonthRange) DisplayFormat(value string) inputMonthRange {
	return m.set("displayFormat", value)
}

// EditorSetting sets editor configuration
func (m inputMonthRange) EditorSetting(value string) inputMonthRange {
	return m.set("editorSetting", value)
}

// Embed enables or disables inline mode
func (m inputMonthRange) Embed(value bool) inputMonthRange {
	return m.set("embed", value)
}

// EndPlaceholder sets end date placeholder
func (m inputMonthRange) EndPlaceholder(value string) inputMonthRange {
	return m.set("endPlaceholder", value)
}

// ExtraName sets extra field name for range component
func (m inputMonthRange) ExtraName(value string) inputMonthRange {
	return m.set("extraName", value)
}

// Format sets submission date format
func (m inputMonthRange) Format(value string) inputMonthRange {
	return m.set("format", value)
}

// Hidden hides or shows the control
func (m inputMonthRange) Hidden(value bool) inputMonthRange {
	return m.set("hidden", value)
}

// HiddenOn sets expression to hide the control
func (m inputMonthRange) HiddenOn(value string) inputMonthRange {
	return m.set("hiddenOn", value)
}

// Hint sets input hint text
func (m inputMonthRange) Hint(value string) inputMonthRange {
	return m.set("hint", value)
}

// Horizontal sets horizontal layout configuration
func (m inputMonthRange) Horizontal(value string) inputMonthRange {
	return m.set("horizontal", value)
}

// ID sets unique component ID
func (m inputMonthRange) ID(value string) inputMonthRange {
	return m.set("id", value)
}

// InitAutoFill sets initial auto-fill value
func (m inputMonthRange) InitAutoFill(value string) inputMonthRange {
	return m.set("initAutoFill", value)
}

// Inline enables or disables inline mode
func (m inputMonthRange) Inline(value bool) inputMonthRange {
	return m.set("inline", value)
}

// InputClassName sets input CSS class name
func (m inputMonthRange) InputClassName(value string) inputMonthRange {
	return m.set("inputClassName", value)
}

// InputFormat sets input date format
func (m inputMonthRange) InputFormat(value string) inputMonthRange {
	return m.set("inputFormat", value)
}

// JoinValues joins selected values with delimiter
func (m inputMonthRange) JoinValues(value bool) inputMonthRange {
	return m.set("joinValues", value)
}

// Label sets label text
func (m inputMonthRange) Label(value string) inputMonthRange {
	return m.set("label", value)
}

// LabelAlign sets label alignment: right, left, top, or inherit
func (m inputMonthRange) LabelAlign(value string) inputMonthRange {
	return m.set("labelAlign", value)
}

// LabelClassName sets label CSS class name
func (m inputMonthRange) LabelClassName(value string) inputMonthRange {
	return m.set("labelClassName", value)
}

// LabelRemark sets label remark text
func (m inputMonthRange) LabelRemark(value string) inputMonthRange {
	return m.set("labelRemark", value)
}

// LabelWidth sets label width
func (m inputMonthRange) LabelWidth(value string) inputMonthRange {
	return m.set("labelWidth", value)
}

// MaxDate sets maximum date limit
func (m inputMonthRange) MaxDate(value string) inputMonthRange {
	return m.set("maxDate", value)
}

// MaxDuration sets maximum duration
func (m inputMonthRange) MaxDuration(value string) inputMonthRange {
	return m.set("maxDuration", value)
}

// MinDate sets minimum date limit
func (m inputMonthRange) MinDate(value string) inputMonthRange {
	return m.set("minDate", value)
}

// MinDuration sets minimum duration
func (m inputMonthRange) MinDuration(value string) inputMonthRange {
	return m.set("minDuration", value)
}

// Mode sets display mode: normal, inline, or horizontal
func (m inputMonthRange) Mode(value string) inputMonthRange {
	return m.set("mode", value)
}

// Name sets field name for form submission
func (m inputMonthRange) Name(value string) inputMonthRange {
	return m.set("name", value)
}

// OnEvent sets event action configuration
func (m inputMonthRange) OnEvent(value any) inputMonthRange {
	return m.set("onEvent", value)
}

// Placeholder sets placeholder text
func (m inputMonthRange) Placeholder(value string) inputMonthRange {
	return m.set("placeholder", value)
}

// PopOverContainerSelector sets popover container selector
func (m inputMonthRange) PopOverContainerSelector(value string) inputMonthRange {
	return m.set("popOverContainerSelector", value)
}

// Ranges sets date range shortcuts
func (m inputMonthRange) Ranges(value string) inputMonthRange {
	return m.set("ranges", value)
}

// ReadOnly sets read-only mode
func (m inputMonthRange) ReadOnly(value bool) inputMonthRange {
	return m.set("readOnly", value)
}

// ReadOnlyOn sets expression for read-only mode
func (m inputMonthRange) ReadOnlyOn(value string) inputMonthRange {
	return m.set("readOnlyOn", value)
}

// Remark sets remark text
func (m inputMonthRange) Remark(value string) inputMonthRange {
	return m.set("remark", value)
}

// Required sets required field
func (m inputMonthRange) Required(value bool) inputMonthRange {
	return m.set("required", value)
}

// Row sets row configuration
func (m inputMonthRange) Row(value string) inputMonthRange {
	return m.set("row", value)
}

// SaveImmediately enables or disables immediate save
func (m inputMonthRange) SaveImmediately(value bool) inputMonthRange {
	return m.set("saveImmediately", value)
}

// Shortcuts sets date range shortcuts
func (m inputMonthRange) Shortcuts(value string) inputMonthRange {
	return m.set("shortcuts", value)
}

// Size sets form item size: xs, sm, md, lg, or full
func (m inputMonthRange) Size(value string) inputMonthRange {
	return m.set("size", value)
}

// StartPlaceholder sets start date placeholder
func (m inputMonthRange) StartPlaceholder(value string) inputMonthRange {
	return m.set("startPlaceholder", value)
}

// Static enables or disables static display
func (m inputMonthRange) Static(value bool) inputMonthRange {
	return m.set("static", value)
}

// StaticClassName sets static display CSS class name
func (m inputMonthRange) StaticClassName(value string) inputMonthRange {
	return m.set("staticClassName", value)
}

// StaticInputClassName sets static input CSS class name
func (m inputMonthRange) StaticInputClassName(value string) inputMonthRange {
	return m.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label CSS class name
func (m inputMonthRange) StaticLabelClassName(value string) inputMonthRange {
	return m.set("staticLabelClassName", value)
}

// StaticOn sets expression for static display
func (m inputMonthRange) StaticOn(value string) inputMonthRange {
	return m.set("staticOn", value)
}

// StaticPlaceholder sets static display placeholder
func (m inputMonthRange) StaticPlaceholder(value string) inputMonthRange {
	return m.set("staticPlaceholder", value)
}

// StaticSchema sets static display schema
func (m inputMonthRange) StaticSchema(value string) inputMonthRange {
	return m.set("staticSchema", value)
}

// Style sets component style
func (m inputMonthRange) Style(value any) inputMonthRange {
	return m.set("style", value)
}

// SubmitOnChange enables or disables form submission on change
func (m inputMonthRange) SubmitOnChange(value bool) inputMonthRange {
	return m.set("submitOnChange", value)
}

// TestIdBuilder sets test ID builder
func (m inputMonthRange) TestIdBuilder(value string) inputMonthRange {
	return m.set("testIdBuilder", value)
}

// Transform sets date transformation function
func (m inputMonthRange) Transform(value string) inputMonthRange {
	return m.set("transform", value)
}

// UseMobileUI enables or disables mobile UI
func (m inputMonthRange) UseMobileUI(value bool) inputMonthRange {
	return m.set("useMobileUI", value)
}

// ValidateApi sets remote validation API
func (m inputMonthRange) ValidateApi(value string) inputMonthRange {
	return m.set("validateApi", value)
}

// ValidateOnChange enables or disables validation on change
func (m inputMonthRange) ValidateOnChange(value bool) inputMonthRange {
	return m.set("validateOnChange", value)
}

// ValidationErrors sets validation error messages
func (m inputMonthRange) ValidationErrors(value string) inputMonthRange {
	return m.set("validationErrors", value)
}

// Validations sets validation rules
func (m inputMonthRange) Validations(value string) inputMonthRange {
	return m.set("validations", value)
}

// Value sets the value, supporting relative values
func (m inputMonthRange) Value(value string) inputMonthRange {
	return m.set("value", value)
}

// ValueFormat sets submission date format
func (m inputMonthRange) ValueFormat(value string) inputMonthRange {
	return m.set("valueFormat", value)
}

// Visible shows or hides the control
func (m inputMonthRange) Visible(value bool) inputMonthRange {
	return m.set("visible", value)
}

// VisibleOn sets expression to show the control
func (m inputMonthRange) VisibleOn(value string) inputMonthRange {
	return m.set("visibleOn", value)
}

// Width sets width in table
func (m inputMonthRange) Width(value string) inputMonthRange {
	return m.set("width", value)
}

func (m inputMonthRange) set(key string, value any) inputMonthRange {
	m[key] = value
	return m
}
