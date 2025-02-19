package comp

import "github.com/zrcoder/amisgo/schema"

// InputMonth represents a month selection component
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/Month
type InputMonth schema.Schema

func NewInputMonth() InputMonth {
	return InputMonth{"type": "input-month"}
}

func (mc InputMonth) set(key string, value any) InputMonth {
	mc[key] = value
	return mc
}

// AutoFill sets auto-fill value
func (mc InputMonth) AutoFill(value string) InputMonth {
	return mc.set("autoFill", value)
}

// BorderMode sets border mode: full | half | none
func (mc InputMonth) BorderMode(value string) InputMonth {
	return mc.set("borderMode", value)
}

// ClassName sets container CSS class name
func (mc InputMonth) ClassName(value string) InputMonth {
	return mc.set("className", value)
}

// ClearValueOnHidden removes value when hidden
func (mc InputMonth) ClearValueOnHidden(value bool) InputMonth {
	return mc.set("clearValueOnHidden", value)
}

// Clearable shows clear button
func (mc InputMonth) Clearable(value bool) InputMonth {
	return mc.set("clearable", value)
}

// Desc sets description
func (mc InputMonth) Desc(value string) InputMonth {
	return mc.set("desc", value)
}

// Description sets HTML description
func (mc InputMonth) Description(value string) InputMonth {
	return mc.set("description", value)
}

// DescriptionClassName sets description CSS class name
func (mc InputMonth) DescriptionClassName(value string) InputMonth {
	return mc.set("descriptionClassName", value)
}

// Disabled disables the component
func (mc InputMonth) Disabled(value bool) InputMonth {
	return mc.set("disabled", value)
}

// DisabledDate sets disabled date function
func (mc InputMonth) DisabledDate(value string) InputMonth {
	return mc.set("disabledDate", value)
}

// DisabledOn sets disabled expression
func (mc InputMonth) DisabledOn(value string) InputMonth {
	return mc.set("disabledOn", value)
}

// DisplayFormat sets display format
func (mc InputMonth) DisplayFormat(value string) InputMonth {
	return mc.set("displayFormat", value)
}

// EditorSetting sets editor configuration
func (mc InputMonth) EditorSetting(value string) InputMonth {
	return mc.set("editorSetting", value)
}

// Emebed sets inline mode
func (mc InputMonth) Emebed(value bool) InputMonth {
	return mc.set("emebed", value)
}

// ExtraName sets extra field name
func (mc InputMonth) ExtraName(value string) InputMonth {
	return mc.set("extraName", value)
}

// Format sets month storage format
func (mc InputMonth) Format(value string) InputMonth {
	return mc.set("format", value)
}

// Hidden hides the component
func (mc InputMonth) Hidden(value bool) InputMonth {
	return mc.set("hidden", value)
}

// HiddenOn sets hidden expression
func (mc InputMonth) HiddenOn(value string) InputMonth {
	return mc.set("hiddenOn", value)
}

// Hint sets input hint
func (mc InputMonth) Hint(value string) InputMonth {
	return mc.set("hint", value)
}

// Horizontal sets horizontal layout configuration
func (mc InputMonth) Horizontal(value string) InputMonth {
	return mc.set("horizontal", value)
}

// ID sets unique component ID
func (mc InputMonth) ID(value string) InputMonth {
	return mc.set("id", value)
}

// InitAutoFill sets initial auto-fill value
func (mc InputMonth) InitAutoFill(value string) InputMonth {
	return mc.set("initAutoFill", value)
}

// Inline sets inline mode
func (mc InputMonth) Inline(value bool) InputMonth {
	return mc.set("inline", value)
}

// InputClassName sets input CSS class name
func (mc InputMonth) InputClassName(value string) InputMonth {
	return mc.set("inputClassName", value)
}

// InputFormat sets input format
func (mc InputMonth) InputFormat(value string) InputMonth {
	return mc.set("inputFormat", value)
}

// Label sets label text
func (mc InputMonth) Label(value string) InputMonth {
	return mc.set("label", value)
}

// LabelAlign sets label alignment: right | left | top | inherit
func (mc InputMonth) LabelAlign(value string) InputMonth {
	return mc.set("labelAlign", value)
}

// LabelClassName sets label CSS class name
func (mc InputMonth) LabelClassName(value string) InputMonth {
	return mc.set("labelClassName", value)
}

// LabelRemark sets label remark
func (mc InputMonth) LabelRemark(value string) InputMonth {
	return mc.set("labelRemark", value)
}

// LabelWidth sets label width
func (mc InputMonth) LabelWidth(value string) InputMonth {
	return mc.set("labelWidth", value)
}

// Mode sets display mode: normal | inline | horizontal
func (mc InputMonth) Mode(value string) InputMonth {
	return mc.set("mode", value)
}

// Name sets field name
func (mc InputMonth) Name(value string) InputMonth {
	return mc.set("name", value)
}

// OnEvent sets event actions
func (mc InputMonth) OnEvent(value any) InputMonth {
	return mc.set("onEvent", value)
}

// Placeholder sets placeholder text
func (mc InputMonth) Placeholder(value string) InputMonth {
	return mc.set("placeholder", value)
}

// ReadOnly sets read-only mode
func (mc InputMonth) ReadOnly(value bool) InputMonth {
	return mc.set("readOnly", value)
}

// ReadOnlyOn sets read-only expression
func (mc InputMonth) ReadOnlyOn(value string) InputMonth {
	return mc.set("readOnlyOn", value)
}

// Remark sets remark
func (mc InputMonth) Remark(value string) InputMonth {
	return mc.set("remark", value)
}

// Required sets required field
func (mc InputMonth) Required(value bool) InputMonth {
	return mc.set("required", value)
}

// Row sets row value
func (mc InputMonth) Row(value string) InputMonth {
	return mc.set("row", value)
}

// SaveImmediately sets immediate save
func (mc InputMonth) SaveImmediately(value bool) InputMonth {
	return mc.set("saveImmediately", value)
}

// Shortcuts sets date shortcuts
func (mc InputMonth) Shortcuts(value string) InputMonth {
	return mc.set("shortcuts", value)
}

// Size sets component size
func (mc InputMonth) Size(value string) InputMonth {
	return mc.set("size", value)
}

// Static sets static display
func (mc InputMonth) Static(value bool) InputMonth {
	return mc.set("static", value)
}

// StaticClassName sets static display CSS class name
func (mc InputMonth) StaticClassName(value string) InputMonth {
	return mc.set("staticClassName", value)
}

// StaticInputClassName sets static input CSS class name
func (mc InputMonth) StaticInputClassName(value string) InputMonth {
	return mc.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label CSS class name
func (mc InputMonth) StaticLabelClassName(value string) InputMonth {
	return mc.set("staticLabelClassName", value)
}

// StaticOn sets static display expression
func (mc InputMonth) StaticOn(value string) InputMonth {
	return mc.set("staticOn", value)
}

// StaticPlaceholder sets static placeholder
func (mc InputMonth) StaticPlaceholder(value string) InputMonth {
	return mc.set("staticPlaceholder", value)
}

// StaticSchema sets static schema.Schema
func (mc InputMonth) StaticSchema(value string) InputMonth {
	return mc.set("staticSchema", value)
}

// Style sets component style
func (mc InputMonth) Style(value any) InputMonth {
	return mc.set("style", value)
}

// SubmitOnChange submits form on change
func (mc InputMonth) SubmitOnChange(value bool) InputMonth {
	return mc.set("submitOnChange", value)
}

// TestIdBuilder sets test ID builder
func (mc InputMonth) TestIdBuilder(value string) InputMonth {
	return mc.set("testIdBuilder", value)
}

// UseMobileUI sets mobile UI usage
func (mc InputMonth) UseMobileUI(value bool) InputMonth {
	return mc.set("useMobileUI", value)
}

// Utc sets UTC time storage
func (mc InputMonth) Utc(value bool) InputMonth {
	return mc.set("utc", value)
}

// ValidateApi sets remote validation API
func (mc InputMonth) ValidateApi(value string) InputMonth {
	return mc.set("validateApi", value)
}

// ValidateOnChange sets validation on change
func (mc InputMonth) ValidateOnChange(value bool) InputMonth {
	return mc.set("validateOnChange", value)
}

// ValidationErrors sets validation error messages
func (mc InputMonth) ValidationErrors(value string) InputMonth {
	return mc.set("validationErrors", value)
}

// Validations sets validation rules
func (mc InputMonth) Validations(value string) InputMonth {
	return mc.set("validations", value)
}

// Value sets default value
func (mc InputMonth) Value(value string) InputMonth {
	return mc.set("value", value)
}

// ValueFormat sets value format
func (mc InputMonth) ValueFormat(value string) InputMonth {
	return mc.set("valueFormat", value)
}

// Visible sets visibility
func (mc InputMonth) Visible(value bool) InputMonth {
	return mc.set("visible", value)
}

// VisibleOn sets visibility expression
func (mc InputMonth) VisibleOn(value string) InputMonth {
	return mc.set("visibleOn", value)
}

// Width sets width in table
func (mc InputMonth) Width(value string) InputMonth {
	return mc.set("width", value)
}
