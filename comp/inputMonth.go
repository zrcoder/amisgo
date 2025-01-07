package comp

// inputMonth represents a month selection component
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/Month

type inputMonth Schema

func InputMonth() inputMonth {
	return inputMonth{}.set("type", "input-month")
}

func (mc inputMonth) set(key string, value any) inputMonth {
	mc[key] = value
	return mc
}

// AutoFill sets auto-fill value
func (mc inputMonth) AutoFill(value string) inputMonth {
	return mc.set("autoFill", value)
}

// BorderMode sets border mode: full | half | none
func (mc inputMonth) BorderMode(value string) inputMonth {
	return mc.set("borderMode", value)
}

// ClassName sets container CSS class name
func (mc inputMonth) ClassName(value string) inputMonth {
	return mc.set("className", value)
}

// ClearValueOnHidden removes value when hidden
func (mc inputMonth) ClearValueOnHidden(value bool) inputMonth {
	return mc.set("clearValueOnHidden", value)
}

// Clearable shows clear button
func (mc inputMonth) Clearable(value bool) inputMonth {
	return mc.set("clearable", value)
}

// Desc sets description
func (mc inputMonth) Desc(value string) inputMonth {
	return mc.set("desc", value)
}

// Description sets HTML description
func (mc inputMonth) Description(value string) inputMonth {
	return mc.set("description", value)
}

// DescriptionClassName sets description CSS class name
func (mc inputMonth) DescriptionClassName(value string) inputMonth {
	return mc.set("descriptionClassName", value)
}

// Disabled disables the component
func (mc inputMonth) Disabled(value bool) inputMonth {
	return mc.set("disabled", value)
}

// DisabledDate sets disabled date function
func (mc inputMonth) DisabledDate(value string) inputMonth {
	return mc.set("disabledDate", value)
}

// DisabledOn sets disabled expression
func (mc inputMonth) DisabledOn(value string) inputMonth {
	return mc.set("disabledOn", value)
}

// DisplayFormat sets display format
func (mc inputMonth) DisplayFormat(value string) inputMonth {
	return mc.set("displayFormat", value)
}

// EditorSetting sets editor configuration
func (mc inputMonth) EditorSetting(value string) inputMonth {
	return mc.set("editorSetting", value)
}

// Emebed sets inline mode
func (mc inputMonth) Emebed(value bool) inputMonth {
	return mc.set("emebed", value)
}

// ExtraName sets extra field name
func (mc inputMonth) ExtraName(value string) inputMonth {
	return mc.set("extraName", value)
}

// Format sets month storage format
func (mc inputMonth) Format(value string) inputMonth {
	return mc.set("format", value)
}

// Hidden hides the component
func (mc inputMonth) Hidden(value bool) inputMonth {
	return mc.set("hidden", value)
}

// HiddenOn sets hidden expression
func (mc inputMonth) HiddenOn(value string) inputMonth {
	return mc.set("hiddenOn", value)
}

// Hint sets input hint
func (mc inputMonth) Hint(value string) inputMonth {
	return mc.set("hint", value)
}

// Horizontal sets horizontal layout configuration
func (mc inputMonth) Horizontal(value string) inputMonth {
	return mc.set("horizontal", value)
}

// ID sets unique component ID
func (mc inputMonth) ID(value string) inputMonth {
	return mc.set("id", value)
}

// InitAutoFill sets initial auto-fill value
func (mc inputMonth) InitAutoFill(value string) inputMonth {
	return mc.set("initAutoFill", value)
}

// Inline sets inline mode
func (mc inputMonth) Inline(value bool) inputMonth {
	return mc.set("inline", value)
}

// InputClassName sets input CSS class name
func (mc inputMonth) InputClassName(value string) inputMonth {
	return mc.set("inputClassName", value)
}

// InputFormat sets input format
func (mc inputMonth) InputFormat(value string) inputMonth {
	return mc.set("inputFormat", value)
}

// Label sets label text
func (mc inputMonth) Label(value string) inputMonth {
	return mc.set("label", value)
}

// LabelAlign sets label alignment: right | left | top | inherit
func (mc inputMonth) LabelAlign(value string) inputMonth {
	return mc.set("labelAlign", value)
}

// LabelClassName sets label CSS class name
func (mc inputMonth) LabelClassName(value string) inputMonth {
	return mc.set("labelClassName", value)
}

// LabelRemark sets label remark
func (mc inputMonth) LabelRemark(value string) inputMonth {
	return mc.set("labelRemark", value)
}

// LabelWidth sets label width
func (mc inputMonth) LabelWidth(value string) inputMonth {
	return mc.set("labelWidth", value)
}

// Mode sets display mode: normal | inline | horizontal
func (mc inputMonth) Mode(value string) inputMonth {
	return mc.set("mode", value)
}

// Name sets field name
func (mc inputMonth) Name(value string) inputMonth {
	return mc.set("name", value)
}

// OnEvent sets event actions
func (mc inputMonth) OnEvent(value any) inputMonth {
	return mc.set("onEvent", value)
}

// Placeholder sets placeholder text
func (mc inputMonth) Placeholder(value string) inputMonth {
	return mc.set("placeholder", value)
}

// ReadOnly sets read-only mode
func (mc inputMonth) ReadOnly(value bool) inputMonth {
	return mc.set("readOnly", value)
}

// ReadOnlyOn sets read-only expression
func (mc inputMonth) ReadOnlyOn(value string) inputMonth {
	return mc.set("readOnlyOn", value)
}

// Remark sets remark
func (mc inputMonth) Remark(value string) inputMonth {
	return mc.set("remark", value)
}

// Required sets required field
func (mc inputMonth) Required(value bool) inputMonth {
	return mc.set("required", value)
}

// Row sets row value
func (mc inputMonth) Row(value string) inputMonth {
	return mc.set("row", value)
}

// SaveImmediately sets immediate save
func (mc inputMonth) SaveImmediately(value bool) inputMonth {
	return mc.set("saveImmediately", value)
}

// Shortcuts sets date shortcuts
func (mc inputMonth) Shortcuts(value string) inputMonth {
	return mc.set("shortcuts", value)
}

// Size sets component size
func (mc inputMonth) Size(value string) inputMonth {
	return mc.set("size", value)
}

// Static sets static display
func (mc inputMonth) Static(value bool) inputMonth {
	return mc.set("static", value)
}

// StaticClassName sets static display CSS class name
func (mc inputMonth) StaticClassName(value string) inputMonth {
	return mc.set("staticClassName", value)
}

// StaticInputClassName sets static input CSS class name
func (mc inputMonth) StaticInputClassName(value string) inputMonth {
	return mc.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label CSS class name
func (mc inputMonth) StaticLabelClassName(value string) inputMonth {
	return mc.set("staticLabelClassName", value)
}

// StaticOn sets static display expression
func (mc inputMonth) StaticOn(value string) inputMonth {
	return mc.set("staticOn", value)
}

// StaticPlaceholder sets static placeholder
func (mc inputMonth) StaticPlaceholder(value string) inputMonth {
	return mc.set("staticPlaceholder", value)
}

// StaticSchema sets static schema
func (mc inputMonth) StaticSchema(value string) inputMonth {
	return mc.set("staticSchema", value)
}

// Style sets component style
func (mc inputMonth) Style(value any) inputMonth {
	return mc.set("style", value)
}

// SubmitOnChange submits form on change
func (mc inputMonth) SubmitOnChange(value bool) inputMonth {
	return mc.set("submitOnChange", value)
}

// TestIdBuilder sets test ID builder
func (mc inputMonth) TestIdBuilder(value string) inputMonth {
	return mc.set("testIdBuilder", value)
}

// UseMobileUI sets mobile UI usage
func (mc inputMonth) UseMobileUI(value bool) inputMonth {
	return mc.set("useMobileUI", value)
}

// Utc sets UTC time storage
func (mc inputMonth) Utc(value bool) inputMonth {
	return mc.set("utc", value)
}

// ValidateApi sets remote validation API
func (mc inputMonth) ValidateApi(value string) inputMonth {
	return mc.set("validateApi", value)
}

// ValidateOnChange sets validation on change
func (mc inputMonth) ValidateOnChange(value bool) inputMonth {
	return mc.set("validateOnChange", value)
}

// ValidationErrors sets validation error messages
func (mc inputMonth) ValidationErrors(value string) inputMonth {
	return mc.set("validationErrors", value)
}

// Validations sets validation rules
func (mc inputMonth) Validations(value string) inputMonth {
	return mc.set("validations", value)
}

// Value sets default value
func (mc inputMonth) Value(value string) inputMonth {
	return mc.set("value", value)
}

// ValueFormat sets value format
func (mc inputMonth) ValueFormat(value string) inputMonth {
	return mc.set("valueFormat", value)
}

// Visible sets visibility
func (mc inputMonth) Visible(value bool) inputMonth {
	return mc.set("visible", value)
}

// VisibleOn sets visibility expression
func (mc inputMonth) VisibleOn(value string) inputMonth {
	return mc.set("visibleOn", value)
}

// Width sets width in table
func (mc inputMonth) Width(value string) inputMonth {
	return mc.set("width", value)
}
