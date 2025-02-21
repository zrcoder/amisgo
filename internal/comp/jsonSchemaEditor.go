package comp

import "github.com/zrcoder/amisgo/schema"

// JsonSchemaEditor JSON schema.Schema Editor component
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/json-schema-editor
type JsonSchemaEditor schema.Schema

func NewJsonSchemaEditor() JsonSchemaEditor {
	return JsonSchemaEditor{"type": "json-schema-editor"}
}

// set sets a field value and returns the editor
func (j JsonSchemaEditor) set(key string, value any) JsonSchemaEditor {
	j[key] = value
	return j
}

// AdvancedSettings sets custom detail panel configuration
func (j JsonSchemaEditor) AdvancedSettings(value string) JsonSchemaEditor {
	return j.set("advancedSettings", value)
}

// AutoFill sets auto-fill options
func (j JsonSchemaEditor) AutoFill(value string) JsonSchemaEditor {
	return j.set("autoFill", value)
}

// ClassName sets the container CSS class name
func (j JsonSchemaEditor) ClassName(value string) JsonSchemaEditor {
	return j.set("className", value)
}

// ClearValueOnHidden removes the form item value when hidden
func (j JsonSchemaEditor) ClearValueOnHidden(value bool) JsonSchemaEditor {
	return j.set("clearValueOnHidden", value)
}

// Definitions sets type templates for complex types
func (j JsonSchemaEditor) Definitions(value string) JsonSchemaEditor {
	return j.set("definitions", value)
}

// Desc sets the description
func (j JsonSchemaEditor) Desc(value string) JsonSchemaEditor {
	return j.set("desc", value)
}

// Description sets the HTML description
func (j JsonSchemaEditor) Description(value string) JsonSchemaEditor {
	return j.set("description", value)
}

// DescriptionClassName sets the description CSS class name
func (j JsonSchemaEditor) DescriptionClassName(value string) JsonSchemaEditor {
	return j.set("descriptionClassName", value)
}

// Disabled sets the disabled state
func (j JsonSchemaEditor) Disabled(value bool) JsonSchemaEditor {
	return j.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (j JsonSchemaEditor) DisabledOn(value string) JsonSchemaEditor {
	return j.set("disabledOn", value)
}

// DisabledTypes sets the disabled types
func (j JsonSchemaEditor) DisabledTypes(value string) JsonSchemaEditor {
	return j.set("disabledTypes", value)
}

// EditorSetting sets the editor configuration
func (j JsonSchemaEditor) EditorSetting(value string) JsonSchemaEditor {
	return j.set("editorSetting", value)
}

// EnableAdvancedSetting enables advanced settings
func (j JsonSchemaEditor) EnableAdvancedSetting(value bool) JsonSchemaEditor {
	return j.set("enableAdvancedSetting", value)
}

// ExtraName sets an extra field name
func (j JsonSchemaEditor) ExtraName(value string) JsonSchemaEditor {
	return j.set("extraName", value)
}

// Hidden sets the hidden state
func (j JsonSchemaEditor) Hidden(value bool) JsonSchemaEditor {
	return j.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (j JsonSchemaEditor) HiddenOn(value string) JsonSchemaEditor {
	return j.set("hiddenOn", value)
}

// Hint sets the input hint
func (j JsonSchemaEditor) Hint(value string) JsonSchemaEditor {
	return j.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (j JsonSchemaEditor) Horizontal(value string) JsonSchemaEditor {
	return j.set("horizontal", value)
}

// ID sets the component unique ID
func (j JsonSchemaEditor) ID(value string) JsonSchemaEditor {
	return j.set("id", value)
}

// InitAutoFill sets the initial auto-fill
func (j JsonSchemaEditor) InitAutoFill(value string) JsonSchemaEditor {
	return j.set("initAutoFill", value)
}

// Inline sets the inline mode
func (j JsonSchemaEditor) Inline(value bool) JsonSchemaEditor {
	return j.set("inline", value)
}

// InputClassName sets the input CSS class name
func (j JsonSchemaEditor) InputClassName(value string) JsonSchemaEditor {
	return j.set("inputClassName", value)
}

// Label sets the label
func (j JsonSchemaEditor) Label(value string) JsonSchemaEditor {
	return j.set("label", value)
}

// LabelAlign sets the label alignment
func (j JsonSchemaEditor) LabelAlign(value string) JsonSchemaEditor {
	return j.set("labelAlign", value)
}

// LabelClassName sets the label CSS class name
func (j JsonSchemaEditor) LabelClassName(value string) JsonSchemaEditor {
	return j.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (j JsonSchemaEditor) LabelRemark(value string) JsonSchemaEditor {
	return j.set("labelRemark", value)
}

// LabelWidth sets the label width
func (j JsonSchemaEditor) LabelWidth(value string) JsonSchemaEditor {
	return j.set("labelWidth", value)
}

// Mini sets the mini mode
func (j JsonSchemaEditor) Mini(value bool) JsonSchemaEditor {
	return j.set("mini", value)
}

// Mode sets the display mode
func (j JsonSchemaEditor) Mode(value string) JsonSchemaEditor {
	return j.set("mode", value)
}

// Name sets the field name
func (j JsonSchemaEditor) Name(value string) JsonSchemaEditor {
	return j.set("name", value)
}

// OnEvent sets the event actions
func (j JsonSchemaEditor) OnEvent(value Event) JsonSchemaEditor {
	return j.set("onEvent", value)
}

// Placeholder sets the placeholder text
func (j JsonSchemaEditor) Placeholder(value string) JsonSchemaEditor {
	return j.set("placeholder", value)
}

// ReadOnly sets the read-only state
func (j JsonSchemaEditor) ReadOnly(value bool) JsonSchemaEditor {
	return j.set("readOnly", value)
}

// ReadOnlyOn sets the read-only expression
func (j JsonSchemaEditor) ReadOnlyOn(value string) JsonSchemaEditor {
	return j.set("readOnlyOn", value)
}

// Remark sets the remark
func (j JsonSchemaEditor) Remark(value string) JsonSchemaEditor {
	return j.set("remark", value)
}

// Required sets the required state
func (j JsonSchemaEditor) Required(value bool) JsonSchemaEditor {
	return j.set("required", value)
}

// RootTypeMutable sets whether the root type is mutable
func (j JsonSchemaEditor) RootTypeMutable(value bool) JsonSchemaEditor {
	return j.set("rootTypeMutable", value)
}

// Row sets the row configuration
func (j JsonSchemaEditor) Row(value string) JsonSchemaEditor {
	return j.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (j JsonSchemaEditor) SaveImmediately(value bool) JsonSchemaEditor {
	return j.set("saveImmediately", value)
}

// ShowRootInfo sets whether to show root type info
func (j JsonSchemaEditor) ShowRootInfo(value bool) JsonSchemaEditor {
	return j.set("showRootInfo", value)
}

// Size sets the form item size
func (j JsonSchemaEditor) Size(value string) JsonSchemaEditor {
	return j.set("size", value)
}

// Static sets the static display state
func (j JsonSchemaEditor) Static(value bool) JsonSchemaEditor {
	return j.set("static", value)
}

// StaticClassName sets the static display CSS class name
func (j JsonSchemaEditor) StaticClassName(value string) JsonSchemaEditor {
	return j.set("staticClassName", value)
}

// StaticInputClassName sets the static input CSS class name
func (j JsonSchemaEditor) StaticInputClassName(value string) JsonSchemaEditor {
	return j.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label CSS class name
func (j JsonSchemaEditor) StaticLabelClassName(value string) JsonSchemaEditor {
	return j.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (j JsonSchemaEditor) StaticOn(value string) JsonSchemaEditor {
	return j.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder
func (j JsonSchemaEditor) StaticPlaceholder(value string) JsonSchemaEditor {
	return j.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema.Schema
func (j JsonSchemaEditor) StaticSchema(value string) JsonSchemaEditor {
	return j.set("staticSchema", value)
}

// Style sets the component style
func (j JsonSchemaEditor) Style(value any) JsonSchemaEditor {
	return j.set("style", value)
}

// SubmitOnChange sets whether to submit on change
func (j JsonSchemaEditor) SubmitOnChange(value bool) JsonSchemaEditor {
	return j.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (j JsonSchemaEditor) TestIdBuilder(value string) JsonSchemaEditor {
	return j.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI
func (j JsonSchemaEditor) UseMobileUI(value bool) JsonSchemaEditor {
	return j.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API
func (j JsonSchemaEditor) ValidateApi(value string) JsonSchemaEditor {
	return j.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (j JsonSchemaEditor) ValidateOnChange(value bool) JsonSchemaEditor {
	return j.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (j JsonSchemaEditor) ValidationErrors(value string) JsonSchemaEditor {
	return j.set("validationErrors", value)
}

// Validations sets the validation rules
func (j JsonSchemaEditor) Validations(value string) JsonSchemaEditor {
	return j.set("validations", value)
}

// Value sets the default value
func (j JsonSchemaEditor) Value(value string) JsonSchemaEditor {
	return j.set("value", value)
}

// Visible sets the visibility
func (j JsonSchemaEditor) Visible(value bool) JsonSchemaEditor {
	return j.set("visible", value)
}

// VisibleOn sets the visibility expression
func (j JsonSchemaEditor) VisibleOn(value string) JsonSchemaEditor {
	return j.set("visibleOn", value)
}

// Width sets the width in a table
func (j JsonSchemaEditor) Width(value string) JsonSchemaEditor {
	return j.set("width", value)
}
