package comp

import "github.com/zrcoder/amisgo/schema"

// InputSignature
type InputSignature schema.Schema

func NewInputSignature() InputSignature {
	return InputSignature{"type": "input-signature"}
}

// set sets a field value and returns the instance
func (i InputSignature) set(key string, value any) InputSignature {
	i[key] = value
	return i
}

// AutoFill sets autoFill value
func (i InputSignature) AutoFill(value string) InputSignature {
	return i.set("autoFill", value)
}

// BgColor sets background color
func (i InputSignature) BgColor(value string) InputSignature {
	return i.set("bgColor", value)
}

// ClassName sets CSS class name
func (i InputSignature) ClassName(value string) InputSignature {
	return i.set("className", value)
}

// ClearBtnIcon sets clear button icon
func (i InputSignature) ClearBtnIcon(value string) InputSignature {
	return i.set("clearBtnIcon", value)
}

// ClearBtnLabel sets clear button label
func (i InputSignature) ClearBtnLabel(value string) InputSignature {
	return i.set("clearBtnLabel", value)
}

// ClearValueOnHidden sets whether to clear value when hidden
func (i InputSignature) ClearValueOnHidden(value bool) InputSignature {
	return i.set("clearValueOnHidden", value)
}

// Color sets component color
func (i InputSignature) Color(value string) InputSignature {
	return i.set("color", value)
}

// ConfirmBtnIcon sets confirm button icon
func (i InputSignature) ConfirmBtnIcon(value string) InputSignature {
	return i.set("confirmBtnIcon", value)
}

// ConfirmBtnLabel sets confirm button label
func (i InputSignature) ConfirmBtnLabel(value string) InputSignature {
	return i.set("confirmBtnLabel", value)
}

// Desc sets description
func (i InputSignature) Desc(value string) InputSignature {
	return i.set("desc", value)
}

// Description sets HTML description
func (i InputSignature) Description(value string) InputSignature {
	return i.set("description", value)
}

// DescriptionClassName sets description CSS class name
func (i InputSignature) DescriptionClassName(value string) InputSignature {
	return i.set("descriptionClassName", value)
}

// Disabled sets whether the component is disabled
func (i InputSignature) Disabled(value bool) InputSignature {
	return i.set("disabled", value)
}

// DisabledOn sets disabled expression
func (i InputSignature) DisabledOn(value string) InputSignature {
	return i.set("disabledOn", value)
}

// EbmedCancelIcon sets embedded cancel button icon
func (i InputSignature) EbmedCancelIcon(value string) InputSignature {
	return i.set("ebmedCancelIcon", value)
}

// EbmedCancelLabel sets embedded cancel button label
func (i InputSignature) EbmedCancelLabel(value string) InputSignature {
	return i.set("ebmedCancelLabel", value)
}

// EditorSetting sets editor settings
func (i InputSignature) EditorSetting(value string) InputSignature {
	return i.set("editorSetting", value)
}

// Embed sets whether the component is embedded
func (i InputSignature) Embed(value bool) InputSignature {
	return i.set("embed", value)
}

// EmbedBtnIcon sets embedded button icon
func (i InputSignature) EmbedBtnIcon(value string) InputSignature {
	return i.set("embedBtnIcon", value)
}

// EmbedBtnLabel sets embedded button label
func (i InputSignature) EmbedBtnLabel(value string) InputSignature {
	return i.set("embedBtnLabel", value)
}

// EmbedConfirmIcon sets embedded confirm button icon
func (i InputSignature) EmbedConfirmIcon(value string) InputSignature {
	return i.set("embedConfirmIcon", value)
}

// EmbedConfirmLabel sets embedded confirm button label
func (i InputSignature) EmbedConfirmLabel(value string) InputSignature {
	return i.set("embedConfirmLabel", value)
}

// ExtraName sets extra field name
func (i InputSignature) ExtraName(value string) InputSignature {
	return i.set("extraName", value)
}

// Height sets component height
func (i InputSignature) Height(value string) InputSignature {
	return i.set("height", value)
}

// Hidden sets whether the component is hidden
func (i InputSignature) Hidden(value bool) InputSignature {
	return i.set("hidden", value)
}

// HiddenOn sets hidden expression
func (i InputSignature) HiddenOn(value string) InputSignature {
	return i.set("hiddenOn", value)
}

// Hint sets input hint
func (i InputSignature) Hint(value string) InputSignature {
	return i.set("hint", value)
}

// Horizontal sets horizontal layout configuration
func (i InputSignature) Horizontal(value string) InputSignature {
	return i.set("horizontal", value)
}

// ID sets component unique ID
func (i InputSignature) ID(value string) InputSignature {
	return i.set("id", value)
}

// InitAutoFill sets initial autoFill value
func (i InputSignature) InitAutoFill(value string) InputSignature {
	return i.set("initAutoFill", value)
}

// Inline sets whether the control is inline
func (i InputSignature) Inline(value bool) InputSignature {
	return i.set("inline", value)
}

// InputClassName sets input CSS class name
func (i InputSignature) InputClassName(value string) InputSignature {
	return i.set("inputClassName", value)
}

// Label sets label text
func (i InputSignature) Label(value string) InputSignature {
	return i.set("label", value)
}

// LabelAlign sets label alignment
func (i InputSignature) LabelAlign(value string) InputSignature {
	return i.set("labelAlign", value)
}

// LabelClassName sets label CSS class name
func (i InputSignature) LabelClassName(value string) InputSignature {
	return i.set("labelClassName", value)
}

// LabelRemark sets label remark
func (i InputSignature) LabelRemark(value string) InputSignature {
	return i.set("labelRemark", value)
}

// LabelWidth sets label width
func (i InputSignature) LabelWidth(value string) InputSignature {
	return i.set("labelWidth", value)
}

// Mode sets display mode
func (i InputSignature) Mode(value string) InputSignature {
	return i.set("mode", value)
}

// Name sets field name
func (i InputSignature) Name(value string) InputSignature {
	return i.set("name", value)
}

// OnEvent sets event actions
func (i InputSignature) OnEvent(value Event) InputSignature {
	return i.set("onEvent", value)
}

// Placeholder sets placeholder text
func (i InputSignature) Placeholder(value string) InputSignature {
	return i.set("placeholder", value)
}

// ReadOnly sets whether the component is read-only
func (i InputSignature) ReadOnly(value bool) InputSignature {
	return i.set("readOnly", value)
}

// ReadOnlyOn sets read-only expression
func (i InputSignature) ReadOnlyOn(value string) InputSignature {
	return i.set("readOnlyOn", value)
}

// Remark sets remark
func (i InputSignature) Remark(value string) InputSignature {
	return i.set("remark", value)
}

// Required sets whether the field is required
func (i InputSignature) Required(value bool) InputSignature {
	return i.set("required", value)
}

// Row sets row value
func (i InputSignature) Row(value string) InputSignature {
	return i.set("row", value)
}

// Size sets component size
func (i InputSignature) Size(value string) InputSignature {
	return i.set("size", value)
}

// Static sets whether the component is static
func (i InputSignature) Static(value bool) InputSignature {
	return i.set("static", value)
}

// StaticClassName sets static CSS class name
func (i InputSignature) StaticClassName(value string) InputSignature {
	return i.set("staticClassName", value)
}

// StaticInputClassName sets static input CSS class name
func (i InputSignature) StaticInputClassName(value string) InputSignature {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label CSS class name
func (i InputSignature) StaticLabelClassName(value string) InputSignature {
	return i.set("staticLabelClassName", value)
}

// StaticOn sets static expression
func (i InputSignature) StaticOn(value string) InputSignature {
	return i.set("staticOn", value)
}

// StaticPlaceholder sets static placeholder
func (i InputSignature) StaticPlaceholder(value string) InputSignature {
	return i.set("staticPlaceholder", value)
}

// StaticSchema sets static schema.Schema
func (i InputSignature) StaticSchema(value string) InputSignature {
	return i.set("staticSchema", value)
}

// Style sets component style
func (i InputSignature) Style(value any) InputSignature {
	return i.set("style", value)
}

// SubmitOnChange sets whether to submit on change
func (i InputSignature) SubmitOnChange(value bool) InputSignature {
	return i.set("submitOnChange", value)
}

// TestIdBuilder sets test ID builder
func (i InputSignature) TestIdBuilder(value string) InputSignature {
	return i.set("testIdBuilder", value)
}

// UndoBtnIcon sets undo button icon
func (i InputSignature) UndoBtnIcon(value string) InputSignature {
	return i.set("undoBtnIcon", value)
}

// UndoBtnLabel sets undo button label
func (i InputSignature) UndoBtnLabel(value string) InputSignature {
	return i.set("undoBtnLabel", value)
}

// UseMobileUI sets whether to use mobile UI
func (i InputSignature) UseMobileUI(value bool) InputSignature {
	return i.set("useMobileUI", value)
}

// ValidateApi sets validation API
func (i InputSignature) ValidateApi(value string) InputSignature {
	return i.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (i InputSignature) ValidateOnChange(value bool) InputSignature {
	return i.set("validateOnChange", value)
}

// ValidationErrors sets validation error messages
func (i InputSignature) ValidationErrors(value string) InputSignature {
	return i.set("validationErrors", value)
}

// Validations sets validation rules
func (i InputSignature) Validations(value string) InputSignature {
	return i.set("validations", value)
}

// Value sets default value
func (i InputSignature) Value(value string) InputSignature {
	return i.set("value", value)
}

// Visible sets whether the component is visible
func (i InputSignature) Visible(value bool) InputSignature {
	return i.set("visible", value)
}

// VisibleOn sets visibility expression
func (i InputSignature) VisibleOn(value string) InputSignature {
	return i.set("visibleOn", value)
}

// Width sets component width
func (i InputSignature) Width(value string) InputSignature {
	return i.set("width", value)
}
