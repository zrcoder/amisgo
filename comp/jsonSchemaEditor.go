package comp

import "github.com/zrcoder/amisgo/model"

// jsonSchemaEditor JSON model.Schema Editor component
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/json-schema-editor

type jsonSchemaEditor model.Schema

func JsonSchemaEditor() jsonSchemaEditor {
	return jsonSchemaEditor{"type": "json-schema-editor"}
}

// set sets a field value and returns the editor
func (j jsonSchemaEditor) set(key string, value any) jsonSchemaEditor {
	j[key] = value
	return j
}

// AdvancedSettings sets custom detail panel configuration
func (j jsonSchemaEditor) AdvancedSettings(value string) jsonSchemaEditor {
	return j.set("advancedSettings", value)
}

// AutoFill sets auto-fill options
func (j jsonSchemaEditor) AutoFill(value string) jsonSchemaEditor {
	return j.set("autoFill", value)
}

// ClassName sets the container CSS class name
func (j jsonSchemaEditor) ClassName(value string) jsonSchemaEditor {
	return j.set("className", value)
}

// ClearValueOnHidden removes the form item value when hidden
func (j jsonSchemaEditor) ClearValueOnHidden(value bool) jsonSchemaEditor {
	return j.set("clearValueOnHidden", value)
}

// Definitions sets type templates for complex types
func (j jsonSchemaEditor) Definitions(value string) jsonSchemaEditor {
	return j.set("definitions", value)
}

// Desc sets the description
func (j jsonSchemaEditor) Desc(value string) jsonSchemaEditor {
	return j.set("desc", value)
}

// Description sets the HTML description
func (j jsonSchemaEditor) Description(value string) jsonSchemaEditor {
	return j.set("description", value)
}

// DescriptionClassName sets the description CSS class name
func (j jsonSchemaEditor) DescriptionClassName(value string) jsonSchemaEditor {
	return j.set("descriptionClassName", value)
}

// Disabled sets the disabled state
func (j jsonSchemaEditor) Disabled(value bool) jsonSchemaEditor {
	return j.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (j jsonSchemaEditor) DisabledOn(value string) jsonSchemaEditor {
	return j.set("disabledOn", value)
}

// DisabledTypes sets the disabled types
func (j jsonSchemaEditor) DisabledTypes(value string) jsonSchemaEditor {
	return j.set("disabledTypes", value)
}

// EditorSetting sets the editor configuration
func (j jsonSchemaEditor) EditorSetting(value string) jsonSchemaEditor {
	return j.set("editorSetting", value)
}

// EnableAdvancedSetting enables advanced settings
func (j jsonSchemaEditor) EnableAdvancedSetting(value bool) jsonSchemaEditor {
	return j.set("enableAdvancedSetting", value)
}

// ExtraName sets an extra field name
func (j jsonSchemaEditor) ExtraName(value string) jsonSchemaEditor {
	return j.set("extraName", value)
}

// Hidden sets the hidden state
func (j jsonSchemaEditor) Hidden(value bool) jsonSchemaEditor {
	return j.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (j jsonSchemaEditor) HiddenOn(value string) jsonSchemaEditor {
	return j.set("hiddenOn", value)
}

// Hint sets the input hint
func (j jsonSchemaEditor) Hint(value string) jsonSchemaEditor {
	return j.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (j jsonSchemaEditor) Horizontal(value string) jsonSchemaEditor {
	return j.set("horizontal", value)
}

// ID sets the component unique ID
func (j jsonSchemaEditor) ID(value string) jsonSchemaEditor {
	return j.set("id", value)
}

// InitAutoFill sets the initial auto-fill
func (j jsonSchemaEditor) InitAutoFill(value string) jsonSchemaEditor {
	return j.set("initAutoFill", value)
}

// Inline sets the inline mode
func (j jsonSchemaEditor) Inline(value bool) jsonSchemaEditor {
	return j.set("inline", value)
}

// InputClassName sets the input CSS class name
func (j jsonSchemaEditor) InputClassName(value string) jsonSchemaEditor {
	return j.set("inputClassName", value)
}

// Label sets the label
func (j jsonSchemaEditor) Label(value string) jsonSchemaEditor {
	return j.set("label", value)
}

// LabelAlign sets the label alignment
func (j jsonSchemaEditor) LabelAlign(value string) jsonSchemaEditor {
	return j.set("labelAlign", value)
}

// LabelClassName sets the label CSS class name
func (j jsonSchemaEditor) LabelClassName(value string) jsonSchemaEditor {
	return j.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (j jsonSchemaEditor) LabelRemark(value string) jsonSchemaEditor {
	return j.set("labelRemark", value)
}

// LabelWidth sets the label width
func (j jsonSchemaEditor) LabelWidth(value string) jsonSchemaEditor {
	return j.set("labelWidth", value)
}

// Mini sets the mini mode
func (j jsonSchemaEditor) Mini(value bool) jsonSchemaEditor {
	return j.set("mini", value)
}

// Mode sets the display mode
func (j jsonSchemaEditor) Mode(value string) jsonSchemaEditor {
	return j.set("mode", value)
}

// Name sets the field name
func (j jsonSchemaEditor) Name(value string) jsonSchemaEditor {
	return j.set("name", value)
}

// OnEvent sets the event actions
func (j jsonSchemaEditor) OnEvent(value any) jsonSchemaEditor {
	return j.set("onEvent", value)
}

// Placeholder sets the placeholder text
func (j jsonSchemaEditor) Placeholder(value string) jsonSchemaEditor {
	return j.set("placeholder", value)
}

// ReadOnly sets the read-only state
func (j jsonSchemaEditor) ReadOnly(value bool) jsonSchemaEditor {
	return j.set("readOnly", value)
}

// ReadOnlyOn sets the read-only expression
func (j jsonSchemaEditor) ReadOnlyOn(value string) jsonSchemaEditor {
	return j.set("readOnlyOn", value)
}

// Remark sets the remark
func (j jsonSchemaEditor) Remark(value string) jsonSchemaEditor {
	return j.set("remark", value)
}

// Required sets the required state
func (j jsonSchemaEditor) Required(value bool) jsonSchemaEditor {
	return j.set("required", value)
}

// RootTypeMutable sets whether the root type is mutable
func (j jsonSchemaEditor) RootTypeMutable(value bool) jsonSchemaEditor {
	return j.set("rootTypeMutable", value)
}

// Row sets the row configuration
func (j jsonSchemaEditor) Row(value string) jsonSchemaEditor {
	return j.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (j jsonSchemaEditor) SaveImmediately(value bool) jsonSchemaEditor {
	return j.set("saveImmediately", value)
}

// ShowRootInfo sets whether to show root type info
func (j jsonSchemaEditor) ShowRootInfo(value bool) jsonSchemaEditor {
	return j.set("showRootInfo", value)
}

// Size sets the form item size
func (j jsonSchemaEditor) Size(value string) jsonSchemaEditor {
	return j.set("size", value)
}

// Static sets the static display state
func (j jsonSchemaEditor) Static(value bool) jsonSchemaEditor {
	return j.set("static", value)
}

// StaticClassName sets the static display CSS class name
func (j jsonSchemaEditor) StaticClassName(value string) jsonSchemaEditor {
	return j.set("staticClassName", value)
}

// StaticInputClassName sets the static input CSS class name
func (j jsonSchemaEditor) StaticInputClassName(value string) jsonSchemaEditor {
	return j.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label CSS class name
func (j jsonSchemaEditor) StaticLabelClassName(value string) jsonSchemaEditor {
	return j.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (j jsonSchemaEditor) StaticOn(value string) jsonSchemaEditor {
	return j.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder
func (j jsonSchemaEditor) StaticPlaceholder(value string) jsonSchemaEditor {
	return j.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema
func (j jsonSchemaEditor) StaticSchema(value string) jsonSchemaEditor {
	return j.set("staticSchema", value)
}

// Style sets the component style
func (j jsonSchemaEditor) Style(value any) jsonSchemaEditor {
	return j.set("style", value)
}

// SubmitOnChange sets whether to submit on change
func (j jsonSchemaEditor) SubmitOnChange(value bool) jsonSchemaEditor {
	return j.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (j jsonSchemaEditor) TestIdBuilder(value string) jsonSchemaEditor {
	return j.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI
func (j jsonSchemaEditor) UseMobileUI(value bool) jsonSchemaEditor {
	return j.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API
func (j jsonSchemaEditor) ValidateApi(value string) jsonSchemaEditor {
	return j.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (j jsonSchemaEditor) ValidateOnChange(value bool) jsonSchemaEditor {
	return j.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (j jsonSchemaEditor) ValidationErrors(value string) jsonSchemaEditor {
	return j.set("validationErrors", value)
}

// Validations sets the validation rules
func (j jsonSchemaEditor) Validations(value string) jsonSchemaEditor {
	return j.set("validations", value)
}

// Value sets the default value
func (j jsonSchemaEditor) Value(value string) jsonSchemaEditor {
	return j.set("value", value)
}

// Visible sets the visibility
func (j jsonSchemaEditor) Visible(value bool) jsonSchemaEditor {
	return j.set("visible", value)
}

// VisibleOn sets the visibility expression
func (j jsonSchemaEditor) VisibleOn(value string) jsonSchemaEditor {
	return j.set("visibleOn", value)
}

// Width sets the width in a table
func (j jsonSchemaEditor) Width(value string) jsonSchemaEditor {
	return j.set("width", value)
}
