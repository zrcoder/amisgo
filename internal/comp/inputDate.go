package comp

import "github.com/zrcoder/amisgo/schema"

// InputDate represents a date picker component
type InputDate schema.Schema

func NewInputDate() InputDate {
	return InputDate{"type": "input-date"}
}

func (d InputDate) set(key string, value any) InputDate {
	d[key] = value
	return d
}

// AutoFill sets autoFill value
func (d InputDate) AutoFill(value string) InputDate {
	return d.set("autoFill", value)
}

// BorderMode sets border mode
func (d InputDate) BorderMode(value string) InputDate {
	return d.set("borderMode", value)
}

// ClassName sets container CSS class name
func (d InputDate) ClassName(value string) InputDate {
	return d.set("className", value)
}

// ClearValueOnHidden clears value when hidden
func (d InputDate) ClearValueOnHidden(value bool) InputDate {
	return d.set("clearValueOnHidden", value)
}

// Clearable shows clear button
func (d InputDate) Clearable(value bool) InputDate {
	return d.set("clearable", value)
}

// CloseOnSelect closes popup on date select
func (d InputDate) CloseOnSelect(value bool) InputDate {
	return d.set("closeOnSelect", value)
}

// Desc sets description
func (d InputDate) Desc(value string) InputDate {
	return d.set("desc", value)
}

// Description sets HTML description
func (d InputDate) Description(value string) InputDate {
	return d.set("description", value)
}

// DescriptionClassName sets description CSS class name
func (d InputDate) DescriptionClassName(value string) InputDate {
	return d.set("descriptionClassName", value)
}

// Disabled disables the component
func (d InputDate) Disabled(value bool) InputDate {
	return d.set("disabled", value)
}

// DisabledDate sets disabled date function
func (d InputDate) DisabledDate(value string) InputDate {
	return d.set("disabledDate", value)
}

// DisabledOn sets disabled expression
func (d InputDate) DisabledOn(value string) InputDate {
	return d.set("disabledOn", value)
}

// DisplayFormat sets display format
func (d InputDate) DisplayFormat(value string) InputDate {
	return d.set("displayFormat", value)
}

// EditorSetting sets editor configuration
func (d InputDate) EditorSetting(value string) InputDate {
	return d.set("editorSetting", value)
}

// Embed sets inline mode
func (d InputDate) Embed(value bool) InputDate {
	return d.set("embed", value)
}

// ExtraName sets extra field name
func (d InputDate) ExtraName(value string) InputDate {
	return d.set("extraName", value)
}

// Format sets storage format
func (d InputDate) Format(value string) InputDate {
	return d.set("format", value)
}

// Hidden hides the component
func (d InputDate) Hidden(value bool) InputDate {
	return d.set("hidden", value)
}

// HiddenOn sets hidden expression
func (d InputDate) HiddenOn(value string) InputDate {
	return d.set("hiddenOn", value)
}

// Hint sets input hint
func (d InputDate) Hint(value string) InputDate {
	return d.set("hint", value)
}

// Horizontal sets horizontal layout configuration
func (d InputDate) Horizontal(value string) InputDate {
	return d.set("horizontal", value)
}

// ID sets unique component ID
func (d InputDate) ID(value string) InputDate {
	return d.set("id", value)
}

// InitAutoFill sets initial autoFill value
func (d InputDate) InitAutoFill(value string) InputDate {
	return d.set("initAutoFill", value)
}

// Inline sets inline mode
func (d InputDate) Inline(value bool) InputDate {
	return d.set("inline", value)
}

// InputClassName sets input CSS class name
func (d InputDate) InputClassName(value string) InputDate {
	return d.set("inputClassName", value)
}

// InputFormat sets input format
func (d InputDate) InputFormat(value string) InputDate {
	return d.set("inputFormat", value)
}

// Label sets label
func (d InputDate) Label(value string) InputDate {
	return d.set("label", value)
}

// LabelAlign sets label alignment
func (d InputDate) LabelAlign(value string) InputDate {
	return d.set("labelAlign", value)
}

// LabelClassName sets label CSS class name
func (d InputDate) LabelClassName(value string) InputDate {
	return d.set("labelClassName", value)
}

// LabelRemark sets label remark
func (d InputDate) LabelRemark(value string) InputDate {
	return d.set("labelRemark", value)
}

// LabelWidth sets label width
func (d InputDate) LabelWidth(value string) InputDate {
	return d.set("labelWidth", value)
}

// MaxDate sets maximum date
func (d InputDate) MaxDate(value string) InputDate {
	return d.set("maxDate", value)
}

// MinDate sets minimum date
func (d InputDate) MinDate(value string) InputDate {
	return d.set("minDate", value)
}

// Mode sets display mode
func (d InputDate) Mode(value string) InputDate {
	return d.set("mode", value)
}

// Name sets field name
func (d InputDate) Name(value string) InputDate {
	return d.set("name", value)
}

// OnEvent sets event configuration
func (d InputDate) OnEvent(value any) InputDate {
	return d.set("onEvent", value)
}

// Placeholder sets placeholder
func (d InputDate) Placeholder(value string) InputDate {
	return d.set("placeholder", value)
}

// PopOverContainerSelector sets popup container selector
func (d InputDate) PopOverContainerSelector(value string) InputDate {
	return d.set("popOverContainerSelector", value)
}

// ReadOnly sets read-only mode
func (d InputDate) ReadOnly(value bool) InputDate {
	return d.set("readOnly", value)
}

// ReadOnlyOn sets read-only expression
func (d InputDate) ReadOnlyOn(value string) InputDate {
	return d.set("readOnlyOn", value)
}

// Remark sets remark
func (d InputDate) Remark(value string) InputDate {
	return d.set("remark", value)
}

// Required sets required field
func (d InputDate) Required(value bool) InputDate {
	return d.set("required", value)
}

// Row sets row configuration
func (d InputDate) Row(value string) InputDate {
	return d.set("row", value)
}

// SaveImmediately sets immediate save
func (d InputDate) SaveImmediately(value bool) InputDate {
	return d.set("saveImmediately", value)
}

// Shortcuts sets date shortcuts
func (d InputDate) Shortcuts(value string) InputDate {
	return d.set("shortcuts", value)
}

// Size sets component size
func (d InputDate) Size(value string) InputDate {
	return d.set("size", value)
}

// Static sets static display
func (d InputDate) Static(value bool) InputDate {
	return d.set("static", value)
}

// StaticClassName sets static CSS class name
func (d InputDate) StaticClassName(value string) InputDate {
	return d.set("staticClassName", value)
}

// StaticInputClassName sets static input CSS class name
func (d InputDate) StaticInputClassName(value string) InputDate {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label CSS class name
func (d InputDate) StaticLabelClassName(value string) InputDate {
	return d.set("staticLabelClassName", value)
}

// StaticOn sets static display expression
func (d InputDate) StaticOn(value string) InputDate {
	return d.set("staticOn", value)
}

// StaticPlaceholder sets static placeholder
func (d InputDate) StaticPlaceholder(value string) InputDate {
	return d.set("staticPlaceholder", value)
}

// StaticSchema sets static schema.Schema
func (d InputDate) StaticSchema(value string) InputDate {
	return d.set("staticSchema", value)
}

// Style sets component style
func (d InputDate) Style(value any) InputDate {
	return d.set("style", value)
}

// SubmitOnChange submits form on change
func (d InputDate) SubmitOnChange(value bool) InputDate {
	return d.set("submitOnChange", value)
}

// TestIdBuilder sets test ID builder
func (d InputDate) TestIdBuilder(value string) InputDate {
	return d.set("testIdBuilder", value)
}

// UseMobileUI sets mobile UI mode
func (d InputDate) UseMobileUI(value bool) InputDate {
	return d.set("useMobileUI", value)
}

// UTC sets UTC storage
func (d InputDate) UTC(value bool) InputDate {
	return d.set("utc", value)
}

// ValidateApi sets remote validation API
func (d InputDate) ValidateApi(value string) InputDate {
	return d.set("validateApi", value)
}

// ValidateOnChange validates on change
func (d InputDate) ValidateOnChange(value bool) InputDate {
	return d.set("validateOnChange", value)
}

// ValidationErrors sets validation error messages
func (d InputDate) ValidationErrors(value string) InputDate {
	return d.set("validationErrors", value)
}

// Validations sets validation rules
func (d InputDate) Validations(value string) InputDate {
	return d.set("validations", value)
}

// Value sets default value
func (d InputDate) Value(value string) InputDate {
	return d.set("value", value)
}

// ValueFormat sets value format
func (d InputDate) ValueFormat(value string) InputDate {
	return d.set("valueFormat", value)
}

// Visible sets visibility
func (d InputDate) Visible(value bool) InputDate {
	return d.set("visible", value)
}

// VisibleOn sets visibility expression
func (d InputDate) VisibleOn(value string) InputDate {
	return d.set("visibleOn", value)
}

// Width sets width in table
func (d InputDate) Width(value string) InputDate {
	return d.set("width", value)
}
