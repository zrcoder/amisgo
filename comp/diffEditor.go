package comp

import "github.com/zrcoder/amisgo/model"

// diffEditor represents the editor schema
type diffEditor model.Schema

// DiffEditor creates a new diffEditor instance
func DiffEditor() diffEditor {
	return make(diffEditor).set("type", "diff-editor")
}

func (d diffEditor) set(key string, value any) diffEditor {
	d[key] = value
	return d
}

// AutoFill sets the autoFill value
func (d diffEditor) AutoFill(value string) diffEditor {
	return d.set("autoFill", value)
}

// ClassName sets the container CSS class name
func (d diffEditor) ClassName(value string) diffEditor {
	return d.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when the form item is hidden
func (d diffEditor) ClearValueOnHidden(value bool) diffEditor {
	return d.set("clearValueOnHidden", value)
}

// Desc sets the description
func (d diffEditor) Desc(value string) diffEditor {
	return d.set("desc", value)
}

// Description sets the description content
func (d diffEditor) Description(value string) diffEditor {
	return d.set("description", value)
}

// DescriptionClassName sets the class name for the description
func (d diffEditor) DescriptionClassName(value string) diffEditor {
	return d.set("descriptionClassName", value)
}

// DiffValue sets the value for the left panel
func (d diffEditor) DiffValue(value string) diffEditor {
	return d.set("diffValue", value)
}

// Disabled sets whether the editor is disabled
func (d diffEditor) Disabled(value bool) diffEditor {
	return d.set("disabled", value)
}

// DisabledOn sets the expression to determine if the editor is disabled
func (d diffEditor) DisabledOn(value string) diffEditor {
	return d.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (d diffEditor) EditorSetting(value string) diffEditor {
	return d.set("editorSetting", value)
}

// ExtraName sets the extra field name
func (d diffEditor) ExtraName(value string) diffEditor {
	return d.set("extraName", value)
}

// Hidden sets whether the editor is hidden
func (d diffEditor) Hidden(value bool) diffEditor {
	return d.set("hidden", value)
}

// HiddenOn sets the expression to determine if the editor is hidden
func (d diffEditor) HiddenOn(value string) diffEditor {
	return d.set("hiddenOn", value)
}

// Hint sets the input hint
func (d diffEditor) Hint(value string) diffEditor {
	return d.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (d diffEditor) Horizontal(value string) diffEditor {
	return d.set("horizontal", value)
}

// ID sets the unique component ID
func (d diffEditor) ID(value string) diffEditor {
	return d.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (d diffEditor) InitAutoFill(value string) diffEditor {
	return d.set("initAutoFill", value)
}

// Inline sets whether the form control is in inline mode
func (d diffEditor) Inline(value bool) diffEditor {
	return d.set("inline", value)
}

// InputClassName sets the input class name
func (d diffEditor) InputClassName(value string) diffEditor {
	return d.set("inputClassName", value)
}

// Label sets the label
func (d diffEditor) Label(value string) diffEditor {
	return d.set("label", value)
}

// LabelAlign sets the label alignment
func (d diffEditor) LabelAlign(value string) diffEditor {
	return d.set("labelAlign", value)
}

// LabelClassName sets the label class name
func (d diffEditor) LabelClassName(value string) diffEditor {
	return d.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (d diffEditor) LabelRemark(value string) diffEditor {
	return d.set("labelRemark", value)
}

// LabelWidth sets the custom label width
func (d diffEditor) LabelWidth(value string) diffEditor {
	return d.set("labelWidth", value)
}

// Language sets the language, refer to monaco-editor
func (d diffEditor) Language(value string) diffEditor {
	return d.set("language", value)
}

// Mode sets the display mode of the current form item
func (d diffEditor) Mode(value string) diffEditor {
	return d.set("mode", value)
}

// Name sets the field name
func (d diffEditor) Name(value string) diffEditor {
	return d.set("name", value)
}

// OnEvent sets the event action configuration
func (d diffEditor) OnEvent(value any) diffEditor {
	return d.set("onEvent", value)
}

// Options sets the editor options
func (d diffEditor) Options(value model.Schema) diffEditor {
	return d.set("options", value)
}

// Placeholder sets the placeholder
func (d diffEditor) Placeholder(value string) diffEditor {
	return d.set("placeholder", value)
}

// ReadOnly sets whether the editor is read-only
func (d diffEditor) ReadOnly(value bool) diffEditor {
	return d.set("readOnly", value)
}

// ReadOnlyOn sets the expression to determine if the editor is read-only
func (d diffEditor) ReadOnlyOn(value string) diffEditor {
	return d.set("readOnlyOn", value)
}

// Remark sets the remark
func (d diffEditor) Remark(value string) diffEditor {
	return d.set("remark", value)
}

// Required sets whether the field is required
func (d diffEditor) Required(value bool) diffEditor {
	return d.set("required", value)
}

// Row sets the row configuration
func (d diffEditor) Row(value string) diffEditor {
	return d.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (d diffEditor) SaveImmediately(value bool) diffEditor {
	return d.set("saveImmediately", value)
}

// Size sets the form item size
func (d diffEditor) Size(value string) diffEditor {
	return d.set("size", value)
}

// Static sets whether to display statically
func (d diffEditor) Static(value bool) diffEditor {
	return d.set("static", value)
}

// StaticClassName sets the class name for static display form items
func (d diffEditor) StaticClassName(value string) diffEditor {
	return d.set("staticClassName", value)
}

// StaticInputClassName sets the class name for static display form item values
func (d diffEditor) StaticInputClassName(value string) diffEditor {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class name for static display form item labels
func (d diffEditor) StaticLabelClassName(value string) diffEditor {
	return d.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the editor is displayed statically
func (d diffEditor) StaticOn(value string) diffEditor {
	return d.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (d diffEditor) StaticPlaceholder(value string) diffEditor {
	return d.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema
func (d diffEditor) StaticSchema(value string) diffEditor {
	return d.set("staticSchema", value)
}

// Style sets the component style
func (d diffEditor) Style(value any) diffEditor {
	return d.set("style", value)
}

// SubmitOnChange sets whether to submit the form when changes are made
func (d diffEditor) SubmitOnChange(value bool) diffEditor {
	return d.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (d diffEditor) TestIdBuilder(value string) diffEditor {
	return d.set("testIdBuilder", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (d diffEditor) UseMobileUI(value bool) diffEditor {
	return d.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API for the form item
func (d diffEditor) ValidateApi(value string) diffEditor {
	return d.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (d diffEditor) ValidateOnChange(value bool) diffEditor {
	return d.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (d diffEditor) ValidationErrors(value string) diffEditor {
	return d.set("validationErrors", value)
}

// Validations sets the validation configuration
func (d diffEditor) Validations(value string) diffEditor {
	return d.set("validations", value)
}

// Value sets the default value
func (d diffEditor) Value(value string) diffEditor {
	return d.set("value", value)
}

// Visible sets whether the editor is visible
func (d diffEditor) Visible(value bool) diffEditor {
	return d.set("visible", value)
}

// VisibleOn sets the expression to determine if the editor is visible
func (d diffEditor) VisibleOn(value string) diffEditor {
	return d.set("visibleOn", value)
}

// Width sets the width in the table
func (d diffEditor) Width(value string) diffEditor {
	return d.set("width", value)
}
