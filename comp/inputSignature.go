package comp

import "github.com/zrcoder/amisgo/model"

// inputSignature

type inputSignature model.Schema

// InputSignature creates a new InputSignature instance with default type
func InputSignature() inputSignature {
	return inputSignature{"type": "input-signature"}
}

// set sets a field value and returns the instance
func (i inputSignature) set(key string, value any) inputSignature {
	i[key] = value
	return i
}

// AutoFill sets autoFill value
func (i inputSignature) AutoFill(value string) inputSignature {
	return i.set("autoFill", value)
}

// BgColor sets background color
func (i inputSignature) BgColor(value string) inputSignature {
	return i.set("bgColor", value)
}

// ClassName sets CSS class name
func (i inputSignature) ClassName(value string) inputSignature {
	return i.set("className", value)
}

// ClearBtnIcon sets clear button icon
func (i inputSignature) ClearBtnIcon(value string) inputSignature {
	return i.set("clearBtnIcon", value)
}

// ClearBtnLabel sets clear button label
func (i inputSignature) ClearBtnLabel(value string) inputSignature {
	return i.set("clearBtnLabel", value)
}

// ClearValueOnHidden sets whether to clear value when hidden
func (i inputSignature) ClearValueOnHidden(value bool) inputSignature {
	return i.set("clearValueOnHidden", value)
}

// Color sets component color
func (i inputSignature) Color(value string) inputSignature {
	return i.set("color", value)
}

// ConfirmBtnIcon sets confirm button icon
func (i inputSignature) ConfirmBtnIcon(value string) inputSignature {
	return i.set("confirmBtnIcon", value)
}

// ConfirmBtnLabel sets confirm button label
func (i inputSignature) ConfirmBtnLabel(value string) inputSignature {
	return i.set("confirmBtnLabel", value)
}

// Desc sets description
func (i inputSignature) Desc(value string) inputSignature {
	return i.set("desc", value)
}

// Description sets HTML description
func (i inputSignature) Description(value string) inputSignature {
	return i.set("description", value)
}

// DescriptionClassName sets description CSS class name
func (i inputSignature) DescriptionClassName(value string) inputSignature {
	return i.set("descriptionClassName", value)
}

// Disabled sets whether the component is disabled
func (i inputSignature) Disabled(value bool) inputSignature {
	return i.set("disabled", value)
}

// DisabledOn sets disabled expression
func (i inputSignature) DisabledOn(value string) inputSignature {
	return i.set("disabledOn", value)
}

// EbmedCancelIcon sets embedded cancel button icon
func (i inputSignature) EbmedCancelIcon(value string) inputSignature {
	return i.set("ebmedCancelIcon", value)
}

// EbmedCancelLabel sets embedded cancel button label
func (i inputSignature) EbmedCancelLabel(value string) inputSignature {
	return i.set("ebmedCancelLabel", value)
}

// EditorSetting sets editor settings
func (i inputSignature) EditorSetting(value string) inputSignature {
	return i.set("editorSetting", value)
}

// Embed sets whether the component is embedded
func (i inputSignature) Embed(value bool) inputSignature {
	return i.set("embed", value)
}

// EmbedBtnIcon sets embedded button icon
func (i inputSignature) EmbedBtnIcon(value string) inputSignature {
	return i.set("embedBtnIcon", value)
}

// EmbedBtnLabel sets embedded button label
func (i inputSignature) EmbedBtnLabel(value string) inputSignature {
	return i.set("embedBtnLabel", value)
}

// EmbedConfirmIcon sets embedded confirm button icon
func (i inputSignature) EmbedConfirmIcon(value string) inputSignature {
	return i.set("embedConfirmIcon", value)
}

// EmbedConfirmLabel sets embedded confirm button label
func (i inputSignature) EmbedConfirmLabel(value string) inputSignature {
	return i.set("embedConfirmLabel", value)
}

// ExtraName sets extra field name
func (i inputSignature) ExtraName(value string) inputSignature {
	return i.set("extraName", value)
}

// Height sets component height
func (i inputSignature) Height(value string) inputSignature {
	return i.set("height", value)
}

// Hidden sets whether the component is hidden
func (i inputSignature) Hidden(value bool) inputSignature {
	return i.set("hidden", value)
}

// HiddenOn sets hidden expression
func (i inputSignature) HiddenOn(value string) inputSignature {
	return i.set("hiddenOn", value)
}

// Hint sets input hint
func (i inputSignature) Hint(value string) inputSignature {
	return i.set("hint", value)
}

// Horizontal sets horizontal layout configuration
func (i inputSignature) Horizontal(value string) inputSignature {
	return i.set("horizontal", value)
}

// ID sets component unique ID
func (i inputSignature) ID(value string) inputSignature {
	return i.set("id", value)
}

// InitAutoFill sets initial autoFill value
func (i inputSignature) InitAutoFill(value string) inputSignature {
	return i.set("initAutoFill", value)
}

// Inline sets whether the control is inline
func (i inputSignature) Inline(value bool) inputSignature {
	return i.set("inline", value)
}

// InputClassName sets input CSS class name
func (i inputSignature) InputClassName(value string) inputSignature {
	return i.set("inputClassName", value)
}

// Label sets label text
func (i inputSignature) Label(value string) inputSignature {
	return i.set("label", value)
}

// LabelAlign sets label alignment
func (i inputSignature) LabelAlign(value string) inputSignature {
	return i.set("labelAlign", value)
}

// LabelClassName sets label CSS class name
func (i inputSignature) LabelClassName(value string) inputSignature {
	return i.set("labelClassName", value)
}

// LabelRemark sets label remark
func (i inputSignature) LabelRemark(value string) inputSignature {
	return i.set("labelRemark", value)
}

// LabelWidth sets label width
func (i inputSignature) LabelWidth(value string) inputSignature {
	return i.set("labelWidth", value)
}

// Mode sets display mode
func (i inputSignature) Mode(value string) inputSignature {
	return i.set("mode", value)
}

// Name sets field name
func (i inputSignature) Name(value string) inputSignature {
	return i.set("name", value)
}

// OnEvent sets event actions
func (i inputSignature) OnEvent(value any) inputSignature {
	return i.set("onEvent", value)
}

// Placeholder sets placeholder text
func (i inputSignature) Placeholder(value string) inputSignature {
	return i.set("placeholder", value)
}

// ReadOnly sets whether the component is read-only
func (i inputSignature) ReadOnly(value bool) inputSignature {
	return i.set("readOnly", value)
}

// ReadOnlyOn sets read-only expression
func (i inputSignature) ReadOnlyOn(value string) inputSignature {
	return i.set("readOnlyOn", value)
}

// Remark sets remark
func (i inputSignature) Remark(value string) inputSignature {
	return i.set("remark", value)
}

// Required sets whether the field is required
func (i inputSignature) Required(value bool) inputSignature {
	return i.set("required", value)
}

// Row sets row value
func (i inputSignature) Row(value string) inputSignature {
	return i.set("row", value)
}

// Size sets component size
func (i inputSignature) Size(value string) inputSignature {
	return i.set("size", value)
}

// Static sets whether the component is static
func (i inputSignature) Static(value bool) inputSignature {
	return i.set("static", value)
}

// StaticClassName sets static CSS class name
func (i inputSignature) StaticClassName(value string) inputSignature {
	return i.set("staticClassName", value)
}

// StaticInputClassName sets static input CSS class name
func (i inputSignature) StaticInputClassName(value string) inputSignature {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName sets static label CSS class name
func (i inputSignature) StaticLabelClassName(value string) inputSignature {
	return i.set("staticLabelClassName", value)
}

// StaticOn sets static expression
func (i inputSignature) StaticOn(value string) inputSignature {
	return i.set("staticOn", value)
}

// StaticPlaceholder sets static placeholder
func (i inputSignature) StaticPlaceholder(value string) inputSignature {
	return i.set("staticPlaceholder", value)
}

// StaticSchema sets static schema
func (i inputSignature) StaticSchema(value string) inputSignature {
	return i.set("staticSchema", value)
}

// Style sets component style
func (i inputSignature) Style(value any) inputSignature {
	return i.set("style", value)
}

// SubmitOnChange sets whether to submit on change
func (i inputSignature) SubmitOnChange(value bool) inputSignature {
	return i.set("submitOnChange", value)
}

// TestIdBuilder sets test ID builder
func (i inputSignature) TestIdBuilder(value string) inputSignature {
	return i.set("testIdBuilder", value)
}

// UndoBtnIcon sets undo button icon
func (i inputSignature) UndoBtnIcon(value string) inputSignature {
	return i.set("undoBtnIcon", value)
}

// UndoBtnLabel sets undo button label
func (i inputSignature) UndoBtnLabel(value string) inputSignature {
	return i.set("undoBtnLabel", value)
}

// UseMobileUI sets whether to use mobile UI
func (i inputSignature) UseMobileUI(value bool) inputSignature {
	return i.set("useMobileUI", value)
}

// ValidateApi sets validation API
func (i inputSignature) ValidateApi(value string) inputSignature {
	return i.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (i inputSignature) ValidateOnChange(value bool) inputSignature {
	return i.set("validateOnChange", value)
}

// ValidationErrors sets validation error messages
func (i inputSignature) ValidationErrors(value string) inputSignature {
	return i.set("validationErrors", value)
}

// Validations sets validation rules
func (i inputSignature) Validations(value string) inputSignature {
	return i.set("validations", value)
}

// Value sets default value
func (i inputSignature) Value(value string) inputSignature {
	return i.set("value", value)
}

// Visible sets whether the component is visible
func (i inputSignature) Visible(value bool) inputSignature {
	return i.set("visible", value)
}

// VisibleOn sets visibility expression
func (i inputSignature) VisibleOn(value string) inputSignature {
	return i.set("visibleOn", value)
}

// Width sets component width
func (i inputSignature) Width(value string) inputSignature {
	return i.set("width", value)
}
