package comp

import "github.com/zrcoder/amisgo/model"

// inputDate represents a date picker component
type inputDate model.Schema

func InputDate() inputDate {
	return inputDate{"type": "input-date"}
}

func (d inputDate) set(key string, value any) inputDate {
	d[key] = value
	return d
}

// AutoFill sets autoFill value
func (d inputDate) AutoFill(value string) inputDate {
	return d.set("autoFill", value)
}

// BorderMode sets border mode
func (d inputDate) BorderMode(value string) inputDate {
	return d.set("borderMode", value)
}

// ClassName sets container CSS class name
func (d inputDate) ClassName(value string) inputDate {
	return d.set("className", value)
}

// ClearValueOnHidden clears value when hidden
func (d inputDate) ClearValueOnHidden(value bool) inputDate {
	return d.set("clearValueOnHidden", value)
}

// Clearable shows clear button
func (d inputDate) Clearable(value bool) inputDate {
	return d.set("clearable", value)
}

// CloseOnSelect closes popup on date select
func (d inputDate) CloseOnSelect(value bool) inputDate {
	return d.set("closeOnSelect", value)
}

// Desc sets description
func (d inputDate) Desc(value string) inputDate {
	return d.set("desc", value)
}

// Description sets HTML description
func (d inputDate) Description(value string) inputDate {
	return d.set("description", value)
}

// DescriptionClassName sets description CSS class name
func (d inputDate) DescriptionClassName(value string) inputDate {
	return d.set("descriptionClassName", value)
}

// Disabled disables the component
func (d inputDate) Disabled(value bool) inputDate {
	return d.set("disabled", value)
}

// DisabledDate sets disabled date function
func (d inputDate) DisabledDate(value string) inputDate {
	return d.set("disabledDate", value)
}

// DisabledOn sets disabled expression
func (d inputDate) DisabledOn(value string) inputDate {
	return d.set("disabledOn", value)
}

// DisplayFormat sets display format
func (d inputDate) DisplayFormat(value string) inputDate {
	return d.set("displayFormat", value)
}

// EditorSetting sets editor configuration
func (d inputDate) EditorSetting(value string) inputDate {
	return d.set("editorSetting", value)
}

// Embed sets inline mode
func (d inputDate) Embed(value bool) inputDate {
	return d.set("embed", value)
}

// ExtraName sets extra field name
func (d inputDate) ExtraName(value string) inputDate {
	return d.set("extraName", value)
}

// Format sets storage format
func (d inputDate) Format(value string) inputDate {
	return d.set("format", value)
}

// Hidden hides the component
func (d inputDate) Hidden(value bool) inputDate {
	return d.set("hidden", value)
}

// HiddenOn sets hidden expression
func (d inputDate) HiddenOn(value string) inputDate {
	return d.set("hiddenOn", value)
}

// Hint sets input hint
func (d inputDate) Hint(value string) inputDate {
	return d.set("hint", value)
}

// Horizontal sets horizontal layout configuration
func (d inputDate) Horizontal(value string) inputDate {
	return d.set("horizontal", value)
}

// ID sets unique component ID
func (d inputDate) ID(value string) inputDate {
	return d.set("id", value)
}

// InitAutoFill sets initial autoFill value
func (d inputDate) InitAutoFill(value string) inputDate {
	return d.set("initAutoFill", value)
}

// Inline sets inline mode
func (d inputDate) Inline(value bool) inputDate {
	return d.set("inline", value)
}

// InputClassName sets input CSS class name
func (d inputDate) InputClassName(value string) inputDate {
	return d.set("inputClassName", value)
}

// InputFormat sets input format
func (d inputDate) InputFormat(value string) inputDate {
	return d.set("inputFormat", value)
}

// Label sets label
func (d inputDate) Label(value string) inputDate {
	return d.set("label", value)
}

// LabelAlign sets label alignment
func (d inputDate) LabelAlign(value string) inputDate {
	return d.set("labelAlign", value)
}

// LabelClassName sets label CSS class name
func (d inputDate) LabelClassName(value string) inputDate {
	return d.set("labelClassName", value)
}

// LabelRemark sets label remark
func (d inputDate) LabelRemark(value string) inputDate {
	return d.set("labelRemark", value)
}

// LabelWidth sets label width
func (d inputDate) LabelWidth(value string) inputDate {
	return d.set("labelWidth", value)
}

// MaxDate sets maximum date
func (d inputDate) MaxDate(value string) inputDate {
	return d.set("maxDate", value)
}

// MinDate sets minimum date
func (d inputDate) MinDate(value string) inputDate {
	return d.set("minDate", value)
}

// Mode sets display mode
func (d inputDate) Mode(value string) inputDate {
	return d.set("mode", value)
}

// Name sets field name
func (d inputDate) Name(value string) inputDate {
	return d.set("name", value)
}

// OnEvent sets event configuration
func (d inputDate) OnEvent(value any) inputDate {
	return d.set("onEvent", value)
}

// Placeholder sets placeholder
func (d inputDate) Placeholder(value string) inputDate {
	return d.set("placeholder", value)
}

// PopOverContainerSelector sets popup container selector
func (d inputDate) PopOverContainerSelector(value string) inputDate {
	return d.set("popOverContainerSelector", value)
}

// ReadOnly sets read-only mode
func (d inputDate) ReadOnly(value bool) inputDate {
	return d.set("readOnly", value)
}

// ReadOnlyOn sets read-only expression
func (d inputDate) ReadOnlyOn(value string) inputDate {
	return d.set("readOnlyOn", value)
}

// Remark sets remark
func (d inputDate) Remark(value string) inputDate {
	return d.set("remark", value)
}

// Required sets required field
func (d inputDate) Required(value bool) inputDate {
	return d.set("required", value)
}

// Row sets row configuration
func (d inputDate) Row(value string) inputDate {
	return d.set("row", value)
}

// SaveImmediately sets immediate save
func (d inputDate) SaveImmediately(value bool) inputDate {
	return d.set("saveImmediately", value)
}

// Shortcuts sets date shortcuts
func (d inputDate) Shortcuts(value string) inputDate {
	return d.set("shortcuts", value)
}

// Size sets component size
func (d inputDate) Size(value string) inputDate {
	return d.set("size", value)
}

// Static sets static display
func (d inputDate) Static(value bool) inputDate {
	return d.set("static", value)
}

// StaticClassName sets static CSS class name
func (d inputDate) StaticClassName(value string) inputDate {
	return d.set("staticClassName", value)
}

// StaticInputClassName sets static input CSS class name
func (d inputDate) StaticInputClassName(value string) inputDate {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label CSS class name
func (d inputDate) StaticLabelClassName(value string) inputDate {
	return d.set("staticLabelClassName", value)
}

// StaticOn sets static display expression
func (d inputDate) StaticOn(value string) inputDate {
	return d.set("staticOn", value)
}

// StaticPlaceholder sets static placeholder
func (d inputDate) StaticPlaceholder(value string) inputDate {
	return d.set("staticPlaceholder", value)
}

// StaticSchema sets static schema
func (d inputDate) StaticSchema(value string) inputDate {
	return d.set("staticSchema", value)
}

// Style sets component style
func (d inputDate) Style(value any) inputDate {
	return d.set("style", value)
}

// SubmitOnChange submits form on change
func (d inputDate) SubmitOnChange(value bool) inputDate {
	return d.set("submitOnChange", value)
}

// TestIdBuilder sets test ID builder
func (d inputDate) TestIdBuilder(value string) inputDate {
	return d.set("testIdBuilder", value)
}

// UseMobileUI sets mobile UI mode
func (d inputDate) UseMobileUI(value bool) inputDate {
	return d.set("useMobileUI", value)
}

// UTC sets UTC storage
func (d inputDate) UTC(value bool) inputDate {
	return d.set("utc", value)
}

// ValidateApi sets remote validation API
func (d inputDate) ValidateApi(value string) inputDate {
	return d.set("validateApi", value)
}

// ValidateOnChange validates on change
func (d inputDate) ValidateOnChange(value bool) inputDate {
	return d.set("validateOnChange", value)
}

// ValidationErrors sets validation error messages
func (d inputDate) ValidationErrors(value string) inputDate {
	return d.set("validationErrors", value)
}

// Validations sets validation rules
func (d inputDate) Validations(value string) inputDate {
	return d.set("validations", value)
}

// Value sets default value
func (d inputDate) Value(value string) inputDate {
	return d.set("value", value)
}

// ValueFormat sets value format
func (d inputDate) ValueFormat(value string) inputDate {
	return d.set("valueFormat", value)
}

// Visible sets visibility
func (d inputDate) Visible(value bool) inputDate {
	return d.set("visible", value)
}

// VisibleOn sets visibility expression
func (d inputDate) VisibleOn(value string) inputDate {
	return d.set("visibleOn", value)
}

// Width sets width in table
func (d inputDate) Width(value string) inputDate {
	return d.set("width", value)
}
