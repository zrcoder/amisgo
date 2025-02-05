package comp

import "github.com/zrcoder/amisgo/model"

// IconPicker represents an icon picker component
type IconPicker model.Schema

func NewIconPicker() IconPicker {
	return IconPicker{"type": "icon-picker"}
}

func (i IconPicker) set(key string, value any) IconPicker {
	i[key] = value
	return i
}

// AutoFill sets autoFill value
func (i IconPicker) AutoFill(value string) IconPicker {
	return i.set("autoFill", value)
}

// ClassName sets container CSS class name
func (i IconPicker) ClassName(value string) IconPicker {
	return i.set("className", value)
}

// ClearValueOnHidden clears value when hidden
func (i IconPicker) ClearValueOnHidden(value bool) IconPicker {
	return i.set("clearValueOnHidden", value)
}

// Desc sets description
func (i IconPicker) Desc(value string) IconPicker {
	return i.set("desc", value)
}

// Description sets HTML description
func (i IconPicker) Description(value string) IconPicker {
	return i.set("description", value)
}

// DescriptionClassName sets description CSS class name
func (i IconPicker) DescriptionClassName(value string) IconPicker {
	return i.set("descriptionClassName", value)
}

// Disabled sets disabled state
func (i IconPicker) Disabled(value bool) IconPicker {
	return i.set("disabled", value)
}

// DisabledOn sets disabled expression
func (i IconPicker) DisabledOn(value string) IconPicker {
	return i.set("disabledOn", value)
}

// EditorSetting sets editor configuration
func (i IconPicker) EditorSetting(value string) IconPicker {
	return i.set("editorSetting", value)
}

// ExtraName sets extra field name
func (i IconPicker) ExtraName(value string) IconPicker {
	return i.set("extraName", value)
}

// Hidden sets hidden state
func (i IconPicker) Hidden(value bool) IconPicker {
	return i.set("hidden", value)
}

// HiddenOn sets hidden expression
func (i IconPicker) HiddenOn(value string) IconPicker {
	return i.set("hiddenOn", value)
}

// Hint sets input hint
func (i IconPicker) Hint(value string) IconPicker {
	return i.set("hint", value)
}

// Horizontal sets horizontal layout configuration
func (i IconPicker) Horizontal(value string) IconPicker {
	return i.set("horizontal", value)
}

// ID sets unique component ID
func (i IconPicker) ID(value string) IconPicker {
	return i.set("id", value)
}

// InitAutoFill sets initial autoFill value
func (i IconPicker) InitAutoFill(value string) IconPicker {
	return i.set("initAutoFill", value)
}

// Inline sets inline mode for form control
func (i IconPicker) Inline(value bool) IconPicker {
	return i.set("inline", value)
}

// InputClassName sets input CSS class name
func (i IconPicker) InputClassName(value string) IconPicker {
	return i.set("inputClassName", value)
}

// Label sets label text
func (i IconPicker) Label(value string) IconPicker {
	return i.set("label", value)
}

// LabelAlign sets label alignment
func (i IconPicker) LabelAlign(value string) IconPicker {
	return i.set("labelAlign", value)
}

// LabelClassName sets label CSS class name
func (i IconPicker) LabelClassName(value string) IconPicker {
	return i.set("labelClassName", value)
}

// LabelRemark sets label remark
func (i IconPicker) LabelRemark(value string) IconPicker {
	return i.set("labelRemark", value)
}

// LabelWidth sets custom label width
func (i IconPicker) LabelWidth(value string) IconPicker {
	return i.set("labelWidth", value)
}

// Mode sets display mode
func (i IconPicker) Mode(value string) IconPicker {
	return i.set("mode", value)
}

// Name sets field name
func (i IconPicker) Name(value string) IconPicker {
	return i.set("name", value)
}

// OnEvent sets event configuration
func (i IconPicker) OnEvent(value any) IconPicker {
	return i.set("onEvent", value)
}

// Placeholder sets placeholder text
func (i IconPicker) Placeholder(value string) IconPicker {
	return i.set("placeholder", value)
}

// ReadOnly sets read-only state
func (i IconPicker) ReadOnly(value bool) IconPicker {
	return i.set("readOnly", value)
}

// ReadOnlyOn sets read-only expression
func (i IconPicker) ReadOnlyOn(value string) IconPicker {
	return i.set("readOnlyOn", value)
}

// Remark sets remark text
func (i IconPicker) Remark(value string) IconPicker {
	return i.set("remark", value)
}

// Required sets required state
func (i IconPicker) Required(value bool) IconPicker {
	return i.set("required", value)
}

// Row sets row value
func (i IconPicker) Row(value string) IconPicker {
	return i.set("row", value)
}

// SaveImmediately sets save immediately state
func (i IconPicker) SaveImmediately(value bool) IconPicker {
	return i.set("saveImmediately", value)
}

// Size sets component size
func (i IconPicker) Size(value string) IconPicker {
	return i.set("size", value)
}

// Static sets static display state
func (i IconPicker) Static(value bool) IconPicker {
	return i.set("static", value)
}

// StaticClassName sets static display CSS class name
func (i IconPicker) StaticClassName(value string) IconPicker {
	return i.set("staticClassName", value)
}

// StaticInputClassName sets static input CSS class name
func (i IconPicker) StaticInputClassName(value string) IconPicker {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label CSS class name
func (i IconPicker) StaticLabelClassName(value string) IconPicker {
	return i.set("staticLabelClassName", value)
}

// StaticOn sets static display expression
func (i IconPicker) StaticOn(value string) IconPicker {
	return i.set("staticOn", value)
}

// StaticPlaceholder sets static placeholder text
func (i IconPicker) StaticPlaceholder(value string) IconPicker {
	return i.set("staticPlaceholder", value)
}

// StaticSchema sets static schema
func (i IconPicker) StaticSchema(value string) IconPicker {
	return i.set("staticSchema", value)
}

// Style sets component style
func (i IconPicker) Style(value any) IconPicker {
	return i.set("style", value)
}

// SubmitOnChange sets submit on change state
func (i IconPicker) SubmitOnChange(value bool) IconPicker {
	return i.set("submitOnChange", value)
}

// TestIdBuilder sets test ID builder
func (i IconPicker) TestIdBuilder(value string) IconPicker {
	return i.set("testIdBuilder", value)
}

// UseMobileUI sets mobile UI state
func (i IconPicker) UseMobileUI(value bool) IconPicker {
	return i.set("useMobileUI", value)
}

// ValidateApi sets remote validation API
func (i IconPicker) ValidateApi(value string) IconPicker {
	return i.set("validateApi", value)
}

// ValidateOnChange sets validate on change state
func (i IconPicker) ValidateOnChange(value bool) IconPicker {
	return i.set("validateOnChange", value)
}

// ValidationErrors sets validation error messages
func (i IconPicker) ValidationErrors(value string) IconPicker {
	return i.set("validationErrors", value)
}

// Validations sets validation rules
func (i IconPicker) Validations(value string) IconPicker {
	return i.set("validations", value)
}

// Value sets default value
func (i IconPicker) Value(value string) IconPicker {
	return i.set("value", value)
}

// Visible sets visibility state
func (i IconPicker) Visible(value bool) IconPicker {
	return i.set("visible", value)
}

// VisibleOn sets visibility expression
func (i IconPicker) VisibleOn(value string) IconPicker {
	return i.set("visibleOn", value)
}

// Width sets width in table
func (i IconPicker) Width(value string) IconPicker {
	return i.set("width", value)
}
