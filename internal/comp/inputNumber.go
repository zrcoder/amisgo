package comp

import "github.com/zrcoder/amisgo/model"

// InputNumber represents a number input field
type InputNumber model.Schema

func NewInputNumber() InputNumber {
	return InputNumber{"type": "input-number"}
}

// set sets a field value
func (nc InputNumber) set(key string, value any) InputNumber {
	nc[key] = value
	return nc
}

// AutoFill sets autoFill value
func (nc InputNumber) AutoFill(value string) InputNumber {
	return nc.set("autoFill", value)
}

// Big sets big value
func (nc InputNumber) Big(value bool) InputNumber {
	return nc.set("big", value)
}

// BorderMode sets border mode
func (nc InputNumber) BorderMode(value string) InputNumber {
	return nc.set("borderMode", value)
}

// ClassName sets container CSS class name
func (nc InputNumber) ClassName(value string) InputNumber {
	return nc.set("className", value)
}

// ClearValueOnHidden sets clearValueOnHidden value
func (nc InputNumber) ClearValueOnHidden(value bool) InputNumber {
	return nc.set("clearValueOnHidden", value)
}

// Desc sets description
func (nc InputNumber) Desc(value string) InputNumber {
	return nc.set("desc", value)
}

// Description sets HTML description
func (nc InputNumber) Description(value string) InputNumber {
	return nc.set("description", value)
}

// DescriptionClassName sets description class name
func (nc InputNumber) DescriptionClassName(value string) InputNumber {
	return nc.set("descriptionClassName", value)
}

// Disabled sets disabled value
func (nc InputNumber) Disabled(value bool) InputNumber {
	return nc.set("disabled", value)
}

// DisabledOn sets disabled expression
func (nc InputNumber) DisabledOn(value string) InputNumber {
	return nc.set("disabledOn", value)
}

// DisplayMode sets display mode
func (nc InputNumber) DisplayMode(value string) InputNumber {
	return nc.set("displayMode", value)
}

// EditorSetting sets editor settings
func (nc InputNumber) EditorSetting(value string) InputNumber {
	return nc.set("editorSetting", value)
}

// ExtraName sets extra field name
func (nc InputNumber) ExtraName(value string) InputNumber {
	return nc.set("extraName", value)
}

// Hidden sets hidden value
func (nc InputNumber) Hidden(value bool) InputNumber {
	return nc.set("hidden", value)
}

// HiddenOn sets hidden expression
func (nc InputNumber) HiddenOn(value string) InputNumber {
	return nc.set("hiddenOn", value)
}

// Hint sets input hint
func (nc InputNumber) Hint(value string) InputNumber {
	return nc.set("hint", value)
}

// Horizontal sets horizontal layout
func (nc InputNumber) Horizontal(value string) InputNumber {
	return nc.set("horizontal", value)
}

// ID sets unique component ID
func (nc InputNumber) ID(value string) InputNumber {
	return nc.set("id", value)
}

// InitAutoFill sets initial autoFill value
func (nc InputNumber) InitAutoFill(value string) InputNumber {
	return nc.set("initAutoFill", value)
}

// Inline sets inline mode
func (nc InputNumber) Inline(value bool) InputNumber {
	return nc.set("inline", value)
}

// InputClassName sets input class name
func (nc InputNumber) InputClassName(value string) InputNumber {
	return nc.set("inputClassName", value)
}

// Keyboard enables keyboard behavior
func (nc InputNumber) Keyboard(value bool) InputNumber {
	return nc.set("keyboard", value)
}

// KilobitSeparator sets kilobit separator
func (nc InputNumber) KilobitSeparator(value bool) InputNumber {
	return nc.set("kilobitSeparator", value)
}

// Label sets label
func (nc InputNumber) Label(value string) InputNumber {
	return nc.set("label", value)
}

// LabelAlign sets label alignment
func (nc InputNumber) LabelAlign(value string) InputNumber {
	return nc.set("labelAlign", value)
}

// LabelClassName sets label class name
func (nc InputNumber) LabelClassName(value string) InputNumber {
	return nc.set("labelClassName", value)
}

// LabelRemark sets label remark
func (nc InputNumber) LabelRemark(value string) InputNumber {
	return nc.set("labelRemark", value)
}

// LabelWidth sets label width
func (nc InputNumber) LabelWidth(value string) InputNumber {
	return nc.set("labelWidth", value)
}

// Max sets maximum value
func (nc InputNumber) Max(value int) InputNumber {
	return nc.set("max", value)
}

// Min sets minimum value
func (nc InputNumber) Min(value int) InputNumber {
	return nc.set("min", value)
}

// Mode sets display mode
func (nc InputNumber) Mode(value string) InputNumber {
	return nc.set("mode", value)
}

// Name sets field name
func (nc InputNumber) Name(value string) InputNumber {
	return nc.set("name", value)
}

// OnEvent sets event actions
func (nc InputNumber) OnEvent(value any) InputNumber {
	return nc.set("onEvent", value)
}

// Placeholder sets placeholder text
func (nc InputNumber) Placeholder(value string) InputNumber {
	return nc.set("placeholder", value)
}

// Precision sets precision
func (nc InputNumber) Precision(value string) InputNumber {
	return nc.set("precision", value)
}

// Prefix sets prefix
func (nc InputNumber) Prefix(value string) InputNumber {
	return nc.set("prefix", value)
}

// ReadOnly sets read-only mode
func (nc InputNumber) ReadOnly(value bool) InputNumber {
	return nc.set("readOnly", value)
}

// ReadOnlyOn sets read-only expression
func (nc InputNumber) ReadOnlyOn(value string) InputNumber {
	return nc.set("readOnlyOn", value)
}

// Remark sets remark
func (nc InputNumber) Remark(value string) InputNumber {
	return nc.set("remark", value)
}

// Required sets required field
func (nc InputNumber) Required(value bool) InputNumber {
	return nc.set("required", value)
}

// Row sets row value
func (nc InputNumber) Row(value string) InputNumber {
	return nc.set("row", value)
}

// SaveImmediately sets save immediately flag
func (nc InputNumber) SaveImmediately(value bool) InputNumber {
	return nc.set("saveImmediately", value)
}

// ShowAsPercent sets percent display
func (nc InputNumber) ShowAsPercent(value bool) InputNumber {
	return nc.set("showAsPercent", value)
}

// ShowSteps sets show steps flag
func (nc InputNumber) ShowSteps(value bool) InputNumber {
	return nc.set("showSteps", value)
}

// Size sets size
func (nc InputNumber) Size(value string) InputNumber {
	return nc.set("size", value)
}

// Static sets static display
func (nc InputNumber) Static(value bool) InputNumber {
	return nc.set("static", value)
}

// StaticClassName sets static class name
func (nc InputNumber) StaticClassName(value string) InputNumber {
	return nc.set("staticClassName", value)
}

// StaticInputClassName sets static input class name
func (nc InputNumber) StaticInputClassName(value string) InputNumber {
	return nc.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label class name
func (nc InputNumber) StaticLabelClassName(value string) InputNumber {
	return nc.set("staticLabelClassName", value)
}

// StaticOn sets static expression
func (nc InputNumber) StaticOn(value string) InputNumber {
	return nc.set("staticOn", value)
}

// StaticPlaceholder sets static placeholder
func (nc InputNumber) StaticPlaceholder(value string) InputNumber {
	return nc.set("staticPlaceholder", value)
}

// StaticSchema sets static schema
func (nc InputNumber) StaticSchema(value string) InputNumber {
	return nc.set("staticSchema", value)
}

// Step sets step value
func (nc InputNumber) Step(value string) InputNumber {
	return nc.set("step", value)
}

// Style sets component style
func (nc InputNumber) Style(value any) InputNumber {
	return nc.set("style", value)
}

// SubmitOnChange sets submit on change flag
func (nc InputNumber) SubmitOnChange(value bool) InputNumber {
	return nc.set("submitOnChange", value)
}

// Suffix sets suffix
func (nc InputNumber) Suffix(value string) InputNumber {
	return nc.set("suffix", value)
}

// TestIdBuilder sets test ID builder
func (nc InputNumber) TestIdBuilder(value string) InputNumber {
	return nc.set("testIdBuilder", value)
}

// UnitOptions sets unit options
func (nc InputNumber) UnitOptions(value ...any) InputNumber {
	return nc.set("unitOptions", value)
}

// UseMobileUI sets mobile UI flag
func (nc InputNumber) UseMobileUI(value bool) InputNumber {
	return nc.set("useMobileUI", value)
}

// ValidateApi sets validation API
func (nc InputNumber) ValidateApi(value string) InputNumber {
	return nc.set("validateApi", value)
}

// ValidateOnChange sets validate on change flag
func (nc InputNumber) ValidateOnChange(value bool) InputNumber {
	return nc.set("validateOnChange", value)
}

// ValidationErrors sets validation errors
func (nc InputNumber) ValidationErrors(value string) InputNumber {
	return nc.set("validationErrors", value)
}

// Validations sets validations
func (nc InputNumber) Validations(value string) InputNumber {
	return nc.set("validations", value)
}

// Value sets default value
func (nc InputNumber) Value(value int) InputNumber {
	return nc.set("value", value)
}

// Visible sets visible flag
func (nc InputNumber) Visible(value bool) InputNumber {
	return nc.set("visible", value)
}

// VisibleOn sets visible expression
func (nc InputNumber) VisibleOn(value string) InputNumber {
	return nc.set("visibleOn", value)
}

// Width sets width in table
func (nc InputNumber) Width(value string) InputNumber {
	return nc.set("width", value)
}
