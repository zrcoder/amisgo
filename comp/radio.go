package comp

import "github.com/zrcoder/amisgo/model"

// Radio control for single selection.
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/radios

type radio model.Schema

// Radio creates a new RadioControl instance
func Radio() radio {
	return radio{"type": "radio"}
}

// set is a private method for setting field values
func (rc radio) set(key string, value any) radio {
	rc[key] = value
	return rc
}

// AutoFill enables automatic filling
func (rc radio) AutoFill(value string) radio {
	return rc.set("autoFill", value)
}

// Badge sets the badge configuration
func (rc radio) Badge(value string) radio {
	return rc.set("badge", value)
}

// ClassName sets the container CSS class name
func (rc radio) ClassName(value string) radio {
	return rc.set("className", value)
}

// ClearValueOnHidden determines whether the form item value should be deleted when hidden
func (rc radio) ClearValueOnHidden(value bool) radio {
	return rc.set("clearValueOnHidden", value)
}

// Desc sets the description
func (rc radio) Desc(value string) radio {
	return rc.set("desc", value)
}

// Description sets the description content
func (rc radio) Description(value string) radio {
	return rc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (rc radio) DescriptionClassName(value string) radio {
	return rc.set("descriptionClassName", value)
}

// Disabled disables the component
func (rc radio) Disabled(value bool) radio {
	return rc.set("disabled", value)
}

// DisabledOn sets the disabling expression
func (rc radio) DisabledOn(value string) radio {
	return rc.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (rc radio) EditorSetting(value string) radio {
	return rc.set("editorSetting", value)
}

// ExtraName sets the extra field name
func (rc radio) ExtraName(value string) radio {
	return rc.set("extraName", value)
}

// FalseValue sets the unselected value
func (rc radio) FalseValue(value string) radio {
	return rc.set("falseValue", value)
}

// Hidden hides the component
func (rc radio) Hidden(value bool) radio {
	return rc.set("hidden", value)
}

// HiddenOn sets the hiding expression
func (rc radio) HiddenOn(value string) radio {
	return rc.set("hiddenOn", value)
}

// Hint sets the input hint
func (rc radio) Hint(value string) radio {
	return rc.set("hint", value)
}

// Horizontal sets the horizontal layout
func (rc radio) Horizontal(value string) radio {
	return rc.set("horizontal", value)
}

// ID sets the component unique ID
func (rc radio) ID(value string) radio {
	return rc.set("id", value)
}

// InitAutoFill
func (rc radio) InitAutoFill(value string) radio {
	return rc.set("initAutoFill", value)
}

// Inline sets the inline mode
func (rc radio) Inline(value bool) radio {
	return rc.set("inline", value)
}

// InputClassName sets the input CSS class name
func (rc radio) InputClassName(value string) radio {
	return rc.set("inputClassName", value)
}

// Label sets the label
func (rc radio) Label(value string) radio {
	return rc.set("label", value)
}

// LabelAlign sets the label alignment
func (rc radio) LabelAlign(value string) radio {
	return rc.set("labelAlign", value)
}

// LabelClassName sets the label CSS class name
func (rc radio) LabelClassName(value string) radio {
	return rc.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (rc radio) LabelRemark(value string) radio {
	return rc.set("labelRemark", value)
}

// LabelWidth sets the label width
func (rc radio) LabelWidth(value string) radio {
	return rc.set("labelWidth", value)
}

// Mode sets the component mode
func (rc radio) Mode(value string) radio {
	return rc.set("mode", value)
}

// Name sets the field name
func (rc radio) Name(value string) radio {
	return rc.set("name", value)
}

// OnEvent sets the event action configuration
func (rc radio) OnEvent(value any) radio {
	return rc.set("onEvent", value)
}

// Option sets the option explanation
func (rc radio) Option(value string) radio {
	return rc.set("option", value)
}

// OptionType sets the option type
func (rc radio) OptionType(value string) radio {
	return rc.set("optionType", value)
}

// Partial
func (rc radio) Partial(value bool) radio {
	return rc.set("partial", value)
}

// Placeholder sets the placeholder
func (rc radio) Placeholder(value string) radio {
	return rc.set("placeholder", value)
}

// ReadOnly sets the read-only mode
func (rc radio) ReadOnly(value bool) radio {
	return rc.set("readOnly", value)
}

// ReadOnlyOn sets the read-only expression
func (rc radio) ReadOnlyOn(value string) radio {
	return rc.set("readOnlyOn", value)
}

// Remark sets the remark
func (rc radio) Remark(value string) radio {
	return rc.set("remark", value)
}

// Required sets the required mode
func (rc radio) Required(value bool) radio {
	return rc.set("required", value)
}

// Row sets the row configuration
func (rc radio) Row(value string) radio {
	return rc.set("row", value)
}

// SaveImmediately sets the save immediately mode
func (rc radio) SaveImmediately(value bool) radio {
	return rc.set("saveImmediately", value)
}

// Size sets the component size
func (rc radio) Size(value string) radio {
	return rc.set("size", value)
}

// Static sets the static mode
func (rc radio) Static(value bool) radio {
	return rc.set("static", value)
}

// StaticClassName sets the static CSS class name
func (rc radio) StaticClassName(value string) radio {
	return rc.set("staticClassName", value)
}

// StaticInputClassName sets the static input CSS class name
func (rc radio) StaticInputClassName(value string) radio {
	return rc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label CSS class name
func (rc radio) StaticLabelClassName(value string) radio {
	return rc.set("staticLabelClassName", value)
}

// StaticOn sets the static expression
func (rc radio) StaticOn(value string) radio {
	return rc.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder
func (rc radio) StaticPlaceholder(value string) radio {
	return rc.set("staticPlaceholder", value)
}

// StaticSchema
func (rc radio) StaticSchema(value string) radio {
	return rc.set("staticSchema", value)
}

// Style sets the component style
func (rc radio) Style(value any) radio {
	return rc.set("style", value)
}

// SubmitOnChange sets the submit on change mode
func (rc radio) SubmitOnChange(value bool) radio {
	return rc.set("submitOnChange", value)
}

// TestIdBuilder
func (rc radio) TestIdBuilder(value string) radio {
	return rc.set("testIdBuilder", value)
}

// TrueValue sets the selected value
func (rc radio) TrueValue(value string) radio {
	return rc.set("trueValue", value)
}

// UseMobileUI sets the mobile UI mode
func (rc radio) UseMobileUI(value bool) radio {
	return rc.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API
func (rc radio) ValidateApi(value string) radio {
	return rc.set("validateApi", value)
}

// ValidateOnChange sets the validation on change mode
func (rc radio) ValidateOnChange(value bool) radio {
	return rc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (rc radio) ValidationErrors(value string) radio {
	return rc.set("validationErrors", value)
}

// Validations sets the validation rules
func (rc radio) Validations(value string) radio {
	return rc.set("validations", value)
}

// Value sets the default value
func (rc radio) Value(value string) radio {
	return rc.set("value", value)
}

// Visible sets the visible mode
func (rc radio) Visible(value bool) radio {
	return rc.set("visible", value)
}

// VisibleOn sets the visible expression
func (rc radio) VisibleOn(value string) radio {
	return rc.set("visibleOn", value)
}

// Width sets the component width
func (rc radio) Width(value string) radio {
	return rc.set("width", value)
}
