package comp

import "github.com/zrcoder/amisgo/schema"

// Radio control for single selection.
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/radios
type Radio schema.Schema

func NewRadio() Radio {
	return Radio{"type": "radio"}
}

// set is a private method for setting field values
func (rc Radio) set(key string, value any) Radio {
	rc[key] = value
	return rc
}

// AutoFill enables automatic filling
func (rc Radio) AutoFill(value string) Radio {
	return rc.set("autoFill", value)
}

// Badge sets the badge configuration
func (rc Radio) Badge(value string) Radio {
	return rc.set("badge", value)
}

// ClassName sets the container CSS class name
func (rc Radio) ClassName(value string) Radio {
	return rc.set("className", value)
}

// ClearValueOnHidden determines whether the form item value should be deleted when hidden
func (rc Radio) ClearValueOnHidden(value bool) Radio {
	return rc.set("clearValueOnHidden", value)
}

// Desc sets the description
func (rc Radio) Desc(value string) Radio {
	return rc.set("desc", value)
}

// Description sets the description content
func (rc Radio) Description(value string) Radio {
	return rc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (rc Radio) DescriptionClassName(value string) Radio {
	return rc.set("descriptionClassName", value)
}

// Disabled disables the component
func (rc Radio) Disabled(value bool) Radio {
	return rc.set("disabled", value)
}

// DisabledOn sets the disabling expression
func (rc Radio) DisabledOn(value string) Radio {
	return rc.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (rc Radio) EditorSetting(value string) Radio {
	return rc.set("editorSetting", value)
}

// ExtraName sets the extra field name
func (rc Radio) ExtraName(value string) Radio {
	return rc.set("extraName", value)
}

// FalseValue sets the unselected value
func (rc Radio) FalseValue(value string) Radio {
	return rc.set("falseValue", value)
}

// Hidden hides the component
func (rc Radio) Hidden(value bool) Radio {
	return rc.set("hidden", value)
}

// HiddenOn sets the hiding expression
func (rc Radio) HiddenOn(value string) Radio {
	return rc.set("hiddenOn", value)
}

// Hint sets the input hint
func (rc Radio) Hint(value string) Radio {
	return rc.set("hint", value)
}

// Horizontal sets the horizontal layout
func (rc Radio) Horizontal(value string) Radio {
	return rc.set("horizontal", value)
}

// ID sets the component unique ID
func (rc Radio) ID(value string) Radio {
	return rc.set("id", value)
}

// InitAutoFill
func (rc Radio) InitAutoFill(value string) Radio {
	return rc.set("initAutoFill", value)
}

// Inline sets the inline mode
func (rc Radio) Inline(value bool) Radio {
	return rc.set("inline", value)
}

// InputClassName sets the input CSS class name
func (rc Radio) InputClassName(value string) Radio {
	return rc.set("inputClassName", value)
}

// Label sets the label
func (rc Radio) Label(value string) Radio {
	return rc.set("label", value)
}

// LabelAlign sets the label alignment
func (rc Radio) LabelAlign(value string) Radio {
	return rc.set("labelAlign", value)
}

// LabelClassName sets the label CSS class name
func (rc Radio) LabelClassName(value string) Radio {
	return rc.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (rc Radio) LabelRemark(value string) Radio {
	return rc.set("labelRemark", value)
}

// LabelWidth sets the label width
func (rc Radio) LabelWidth(value string) Radio {
	return rc.set("labelWidth", value)
}

// Mode sets the component mode
func (rc Radio) Mode(value string) Radio {
	return rc.set("mode", value)
}

// Name sets the field name
func (rc Radio) Name(value string) Radio {
	return rc.set("name", value)
}

// OnEvent sets the event action configuration
func (rc Radio) OnEvent(value any) Radio {
	return rc.set("onEvent", value)
}

// Option sets the option explanation
func (rc Radio) Option(value string) Radio {
	return rc.set("option", value)
}

// OptionType sets the option type
func (rc Radio) OptionType(value string) Radio {
	return rc.set("optionType", value)
}

// Partial
func (rc Radio) Partial(value bool) Radio {
	return rc.set("partial", value)
}

// Placeholder sets the placeholder
func (rc Radio) Placeholder(value string) Radio {
	return rc.set("placeholder", value)
}

// ReadOnly sets the read-only mode
func (rc Radio) ReadOnly(value bool) Radio {
	return rc.set("readOnly", value)
}

// ReadOnlyOn sets the read-only expression
func (rc Radio) ReadOnlyOn(value string) Radio {
	return rc.set("readOnlyOn", value)
}

// Remark sets the remark
func (rc Radio) Remark(value string) Radio {
	return rc.set("remark", value)
}

// Required sets the required mode
func (rc Radio) Required(value bool) Radio {
	return rc.set("required", value)
}

// Row sets the row configuration
func (rc Radio) Row(value string) Radio {
	return rc.set("row", value)
}

// SaveImmediately sets the save immediately mode
func (rc Radio) SaveImmediately(value bool) Radio {
	return rc.set("saveImmediately", value)
}

// Size sets the component size
func (rc Radio) Size(value string) Radio {
	return rc.set("size", value)
}

// Static sets the static mode
func (rc Radio) Static(value bool) Radio {
	return rc.set("static", value)
}

// StaticClassName sets the static CSS class name
func (rc Radio) StaticClassName(value string) Radio {
	return rc.set("staticClassName", value)
}

// StaticInputClassName sets the static input CSS class name
func (rc Radio) StaticInputClassName(value string) Radio {
	return rc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label CSS class name
func (rc Radio) StaticLabelClassName(value string) Radio {
	return rc.set("staticLabelClassName", value)
}

// StaticOn sets the static expression
func (rc Radio) StaticOn(value string) Radio {
	return rc.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder
func (rc Radio) StaticPlaceholder(value string) Radio {
	return rc.set("staticPlaceholder", value)
}

// StaticSchema
func (rc Radio) StaticSchema(value string) Radio {
	return rc.set("staticSchema", value)
}

// Style sets the component style
func (rc Radio) Style(value any) Radio {
	return rc.set("style", value)
}

// SubmitOnChange sets the submit on change mode
func (rc Radio) SubmitOnChange(value bool) Radio {
	return rc.set("submitOnChange", value)
}

// TestIdBuilder
func (rc Radio) TestIdBuilder(value string) Radio {
	return rc.set("testIdBuilder", value)
}

// TrueValue sets the selected value
func (rc Radio) TrueValue(value string) Radio {
	return rc.set("trueValue", value)
}

// UseMobileUI sets the mobile UI mode
func (rc Radio) UseMobileUI(value bool) Radio {
	return rc.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API
func (rc Radio) ValidateApi(value string) Radio {
	return rc.set("validateApi", value)
}

// ValidateOnChange sets the validation on change mode
func (rc Radio) ValidateOnChange(value bool) Radio {
	return rc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (rc Radio) ValidationErrors(value string) Radio {
	return rc.set("validationErrors", value)
}

// Validations sets the validation rules
func (rc Radio) Validations(value string) Radio {
	return rc.set("validations", value)
}

// Value sets the default value
func (rc Radio) Value(value string) Radio {
	return rc.set("value", value)
}

// Visible sets the visible mode
func (rc Radio) Visible(value bool) Radio {
	return rc.set("visible", value)
}

// VisibleOn sets the visible expression
func (rc Radio) VisibleOn(value string) Radio {
	return rc.set("visibleOn", value)
}

// Width sets the component width
func (rc Radio) Width(value string) Radio {
	return rc.set("width", value)
}
