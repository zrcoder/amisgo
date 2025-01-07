package comp

// inputNumber represents a number input field
type inputNumber Schema

// InputNumber creates a new NumberControl instance
func InputNumber() inputNumber {
	return inputNumber{}.set("type", "input-number")
}

// set sets a field value
func (nc inputNumber) set(key string, value any) inputNumber {
	nc[key] = value
	return nc
}

// AutoFill sets autoFill value
func (nc inputNumber) AutoFill(value string) inputNumber {
	return nc.set("autoFill", value)
}

// Big sets big value
func (nc inputNumber) Big(value bool) inputNumber {
	return nc.set("big", value)
}

// BorderMode sets border mode
func (nc inputNumber) BorderMode(value string) inputNumber {
	return nc.set("borderMode", value)
}

// ClassName sets container CSS class name
func (nc inputNumber) ClassName(value string) inputNumber {
	return nc.set("className", value)
}

// ClearValueOnHidden sets clearValueOnHidden value
func (nc inputNumber) ClearValueOnHidden(value bool) inputNumber {
	return nc.set("clearValueOnHidden", value)
}

// Desc sets description
func (nc inputNumber) Desc(value string) inputNumber {
	return nc.set("desc", value)
}

// Description sets HTML description
func (nc inputNumber) Description(value string) inputNumber {
	return nc.set("description", value)
}

// DescriptionClassName sets description class name
func (nc inputNumber) DescriptionClassName(value string) inputNumber {
	return nc.set("descriptionClassName", value)
}

// Disabled sets disabled value
func (nc inputNumber) Disabled(value bool) inputNumber {
	return nc.set("disabled", value)
}

// DisabledOn sets disabled expression
func (nc inputNumber) DisabledOn(value string) inputNumber {
	return nc.set("disabledOn", value)
}

// DisplayMode sets display mode
func (nc inputNumber) DisplayMode(value string) inputNumber {
	return nc.set("displayMode", value)
}

// EditorSetting sets editor settings
func (nc inputNumber) EditorSetting(value string) inputNumber {
	return nc.set("editorSetting", value)
}

// ExtraName sets extra field name
func (nc inputNumber) ExtraName(value string) inputNumber {
	return nc.set("extraName", value)
}

// Hidden sets hidden value
func (nc inputNumber) Hidden(value bool) inputNumber {
	return nc.set("hidden", value)
}

// HiddenOn sets hidden expression
func (nc inputNumber) HiddenOn(value string) inputNumber {
	return nc.set("hiddenOn", value)
}

// Hint sets input hint
func (nc inputNumber) Hint(value string) inputNumber {
	return nc.set("hint", value)
}

// Horizontal sets horizontal layout
func (nc inputNumber) Horizontal(value string) inputNumber {
	return nc.set("horizontal", value)
}

// ID sets unique component ID
func (nc inputNumber) ID(value string) inputNumber {
	return nc.set("id", value)
}

// InitAutoFill sets initial autoFill value
func (nc inputNumber) InitAutoFill(value string) inputNumber {
	return nc.set("initAutoFill", value)
}

// Inline sets inline mode
func (nc inputNumber) Inline(value bool) inputNumber {
	return nc.set("inline", value)
}

// InputClassName sets input class name
func (nc inputNumber) InputClassName(value string) inputNumber {
	return nc.set("inputClassName", value)
}

// Keyboard enables keyboard behavior
func (nc inputNumber) Keyboard(value bool) inputNumber {
	return nc.set("keyboard", value)
}

// KilobitSeparator sets kilobit separator
func (nc inputNumber) KilobitSeparator(value bool) inputNumber {
	return nc.set("kilobitSeparator", value)
}

// Label sets label
func (nc inputNumber) Label(value string) inputNumber {
	return nc.set("label", value)
}

// LabelAlign sets label alignment
func (nc inputNumber) LabelAlign(value string) inputNumber {
	return nc.set("labelAlign", value)
}

// LabelClassName sets label class name
func (nc inputNumber) LabelClassName(value string) inputNumber {
	return nc.set("labelClassName", value)
}

// LabelRemark sets label remark
func (nc inputNumber) LabelRemark(value string) inputNumber {
	return nc.set("labelRemark", value)
}

// LabelWidth sets label width
func (nc inputNumber) LabelWidth(value string) inputNumber {
	return nc.set("labelWidth", value)
}

// Max sets maximum value
func (nc inputNumber) Max(value int) inputNumber {
	return nc.set("max", value)
}

// Min sets minimum value
func (nc inputNumber) Min(value int) inputNumber {
	return nc.set("min", value)
}

// Mode sets display mode
func (nc inputNumber) Mode(value string) inputNumber {
	return nc.set("mode", value)
}

// Name sets field name
func (nc inputNumber) Name(value string) inputNumber {
	return nc.set("name", value)
}

// OnEvent sets event actions
func (nc inputNumber) OnEvent(value any) inputNumber {
	return nc.set("onEvent", value)
}

// Placeholder sets placeholder text
func (nc inputNumber) Placeholder(value string) inputNumber {
	return nc.set("placeholder", value)
}

// Precision sets precision
func (nc inputNumber) Precision(value string) inputNumber {
	return nc.set("precision", value)
}

// Prefix sets prefix
func (nc inputNumber) Prefix(value string) inputNumber {
	return nc.set("prefix", value)
}

// ReadOnly sets read-only mode
func (nc inputNumber) ReadOnly(value bool) inputNumber {
	return nc.set("readOnly", value)
}

// ReadOnlyOn sets read-only expression
func (nc inputNumber) ReadOnlyOn(value string) inputNumber {
	return nc.set("readOnlyOn", value)
}

// Remark sets remark
func (nc inputNumber) Remark(value string) inputNumber {
	return nc.set("remark", value)
}

// Required sets required field
func (nc inputNumber) Required(value bool) inputNumber {
	return nc.set("required", value)
}

// Row sets row value
func (nc inputNumber) Row(value string) inputNumber {
	return nc.set("row", value)
}

// SaveImmediately sets save immediately flag
func (nc inputNumber) SaveImmediately(value bool) inputNumber {
	return nc.set("saveImmediately", value)
}

// ShowAsPercent sets percent display
func (nc inputNumber) ShowAsPercent(value bool) inputNumber {
	return nc.set("showAsPercent", value)
}

// ShowSteps sets show steps flag
func (nc inputNumber) ShowSteps(value bool) inputNumber {
	return nc.set("showSteps", value)
}

// Size sets size
func (nc inputNumber) Size(value string) inputNumber {
	return nc.set("size", value)
}

// Static sets static display
func (nc inputNumber) Static(value bool) inputNumber {
	return nc.set("static", value)
}

// StaticClassName sets static class name
func (nc inputNumber) StaticClassName(value string) inputNumber {
	return nc.set("staticClassName", value)
}

// StaticInputClassName sets static input class name
func (nc inputNumber) StaticInputClassName(value string) inputNumber {
	return nc.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label class name
func (nc inputNumber) StaticLabelClassName(value string) inputNumber {
	return nc.set("staticLabelClassName", value)
}

// StaticOn sets static expression
func (nc inputNumber) StaticOn(value string) inputNumber {
	return nc.set("staticOn", value)
}

// StaticPlaceholder sets static placeholder
func (nc inputNumber) StaticPlaceholder(value string) inputNumber {
	return nc.set("staticPlaceholder", value)
}

// StaticSchema sets static schema
func (nc inputNumber) StaticSchema(value string) inputNumber {
	return nc.set("staticSchema", value)
}

// Step sets step value
func (nc inputNumber) Step(value string) inputNumber {
	return nc.set("step", value)
}

// Style sets component style
func (nc inputNumber) Style(value any) inputNumber {
	return nc.set("style", value)
}

// SubmitOnChange sets submit on change flag
func (nc inputNumber) SubmitOnChange(value bool) inputNumber {
	return nc.set("submitOnChange", value)
}

// Suffix sets suffix
func (nc inputNumber) Suffix(value string) inputNumber {
	return nc.set("suffix", value)
}

// TestIdBuilder sets test ID builder
func (nc inputNumber) TestIdBuilder(value string) inputNumber {
	return nc.set("testIdBuilder", value)
}

// UnitOptions sets unit options
func (nc inputNumber) UnitOptions(value ...any) inputNumber {
	return nc.set("unitOptions", value)
}

// UseMobileUI sets mobile UI flag
func (nc inputNumber) UseMobileUI(value bool) inputNumber {
	return nc.set("useMobileUI", value)
}

// ValidateApi sets validation API
func (nc inputNumber) ValidateApi(value string) inputNumber {
	return nc.set("validateApi", value)
}

// ValidateOnChange sets validate on change flag
func (nc inputNumber) ValidateOnChange(value bool) inputNumber {
	return nc.set("validateOnChange", value)
}

// ValidationErrors sets validation errors
func (nc inputNumber) ValidationErrors(value string) inputNumber {
	return nc.set("validationErrors", value)
}

// Validations sets validations
func (nc inputNumber) Validations(value string) inputNumber {
	return nc.set("validations", value)
}

// Value sets default value
func (nc inputNumber) Value(value int) inputNumber {
	return nc.set("value", value)
}

// Visible sets visible flag
func (nc inputNumber) Visible(value bool) inputNumber {
	return nc.set("visible", value)
}

// VisibleOn sets visible expression
func (nc inputNumber) VisibleOn(value string) inputNumber {
	return nc.set("visibleOn", value)
}

// Width sets width in table
func (nc inputNumber) Width(value string) inputNumber {
	return nc.set("width", value)
}
