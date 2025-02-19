package comp

import "github.com/zrcoder/amisgo/schema"

// DiffEditor represents the editor schema.Schema
type DiffEditor schema.Schema

func NewDiffEditor() DiffEditor {
	return DiffEditor{"type": "diff-editor"}
}

func (d DiffEditor) set(key string, value any) DiffEditor {
	d[key] = value
	return d
}

// AutoFill sets the autoFill value
func (d DiffEditor) AutoFill(value string) DiffEditor {
	return d.set("autoFill", value)
}

// ClassName sets the container CSS class name
func (d DiffEditor) ClassName(value string) DiffEditor {
	return d.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when the form item is hidden
func (d DiffEditor) ClearValueOnHidden(value bool) DiffEditor {
	return d.set("clearValueOnHidden", value)
}

// Desc sets the description
func (d DiffEditor) Desc(value string) DiffEditor {
	return d.set("desc", value)
}

// Description sets the description content
func (d DiffEditor) Description(value string) DiffEditor {
	return d.set("description", value)
}

// DescriptionClassName sets the class name for the description
func (d DiffEditor) DescriptionClassName(value string) DiffEditor {
	return d.set("descriptionClassName", value)
}

// DiffValue sets the value for the left panel
func (d DiffEditor) DiffValue(value string) DiffEditor {
	return d.set("diffValue", value)
}

// Disabled sets whether the editor is disabled
func (d DiffEditor) Disabled(value bool) DiffEditor {
	return d.set("disabled", value)
}

// DisabledOn sets the expression to determine if the editor is disabled
func (d DiffEditor) DisabledOn(value string) DiffEditor {
	return d.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (d DiffEditor) EditorSetting(value string) DiffEditor {
	return d.set("editorSetting", value)
}

// ExtraName sets the extra field name
func (d DiffEditor) ExtraName(value string) DiffEditor {
	return d.set("extraName", value)
}

// Hidden sets whether the editor is hidden
func (d DiffEditor) Hidden(value bool) DiffEditor {
	return d.set("hidden", value)
}

// HiddenOn sets the expression to determine if the editor is hidden
func (d DiffEditor) HiddenOn(value string) DiffEditor {
	return d.set("hiddenOn", value)
}

// Hint sets the input hint
func (d DiffEditor) Hint(value string) DiffEditor {
	return d.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (d DiffEditor) Horizontal(value string) DiffEditor {
	return d.set("horizontal", value)
}

// ID sets the unique component ID
func (d DiffEditor) ID(value string) DiffEditor {
	return d.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (d DiffEditor) InitAutoFill(value string) DiffEditor {
	return d.set("initAutoFill", value)
}

// Inline sets whether the form control is in inline mode
func (d DiffEditor) Inline(value bool) DiffEditor {
	return d.set("inline", value)
}

// InputClassName sets the input class name
func (d DiffEditor) InputClassName(value string) DiffEditor {
	return d.set("inputClassName", value)
}

// Label sets the label
func (d DiffEditor) Label(value string) DiffEditor {
	return d.set("label", value)
}

// LabelAlign sets the label alignment
func (d DiffEditor) LabelAlign(value string) DiffEditor {
	return d.set("labelAlign", value)
}

// LabelClassName sets the label class name
func (d DiffEditor) LabelClassName(value string) DiffEditor {
	return d.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (d DiffEditor) LabelRemark(value string) DiffEditor {
	return d.set("labelRemark", value)
}

// LabelWidth sets the custom label width
func (d DiffEditor) LabelWidth(value string) DiffEditor {
	return d.set("labelWidth", value)
}

// Language sets the language, refer to monaco-editor
func (d DiffEditor) Language(value string) DiffEditor {
	return d.set("language", value)
}

// Mode sets the display mode of the current form item
func (d DiffEditor) Mode(value string) DiffEditor {
	return d.set("mode", value)
}

// Name sets the field name
func (d DiffEditor) Name(value string) DiffEditor {
	return d.set("name", value)
}

// OnEvent sets the event action configuration
func (d DiffEditor) OnEvent(value any) DiffEditor {
	return d.set("onEvent", value)
}

// Options sets the editor options
func (d DiffEditor) Options(value schema.Schema) DiffEditor {
	return d.set("options", value)
}

// Placeholder sets the placeholder
func (d DiffEditor) Placeholder(value string) DiffEditor {
	return d.set("placeholder", value)
}

// ReadOnly sets whether the editor is read-only
func (d DiffEditor) ReadOnly(value bool) DiffEditor {
	return d.set("readOnly", value)
}

// ReadOnlyOn sets the expression to determine if the editor is read-only
func (d DiffEditor) ReadOnlyOn(value string) DiffEditor {
	return d.set("readOnlyOn", value)
}

// Remark sets the remark
func (d DiffEditor) Remark(value string) DiffEditor {
	return d.set("remark", value)
}

// Required sets whether the field is required
func (d DiffEditor) Required(value bool) DiffEditor {
	return d.set("required", value)
}

// Row sets the row configuration
func (d DiffEditor) Row(value string) DiffEditor {
	return d.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (d DiffEditor) SaveImmediately(value bool) DiffEditor {
	return d.set("saveImmediately", value)
}

// Size sets the form item size
func (d DiffEditor) Size(value string) DiffEditor {
	return d.set("size", value)
}

// Static sets whether to display statically
func (d DiffEditor) Static(value bool) DiffEditor {
	return d.set("static", value)
}

// StaticClassName sets the class name for static display form items
func (d DiffEditor) StaticClassName(value string) DiffEditor {
	return d.set("staticClassName", value)
}

// StaticInputClassName sets the class name for static display form item values
func (d DiffEditor) StaticInputClassName(value string) DiffEditor {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class name for static display form item labels
func (d DiffEditor) StaticLabelClassName(value string) DiffEditor {
	return d.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the editor is displayed statically
func (d DiffEditor) StaticOn(value string) DiffEditor {
	return d.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (d DiffEditor) StaticPlaceholder(value string) DiffEditor {
	return d.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema.Schema
func (d DiffEditor) StaticSchema(value string) DiffEditor {
	return d.set("staticSchema", value)
}

// Style sets the component style
func (d DiffEditor) Style(value any) DiffEditor {
	return d.set("style", value)
}

// SubmitOnChange sets whether to submit the form when changes are made
func (d DiffEditor) SubmitOnChange(value bool) DiffEditor {
	return d.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (d DiffEditor) TestIdBuilder(value string) DiffEditor {
	return d.set("testIdBuilder", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (d DiffEditor) UseMobileUI(value bool) DiffEditor {
	return d.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API for the form item
func (d DiffEditor) ValidateApi(value string) DiffEditor {
	return d.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (d DiffEditor) ValidateOnChange(value bool) DiffEditor {
	return d.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (d DiffEditor) ValidationErrors(value string) DiffEditor {
	return d.set("validationErrors", value)
}

// Validations sets the validation configuration
func (d DiffEditor) Validations(value string) DiffEditor {
	return d.set("validations", value)
}

// Value sets the default value
func (d DiffEditor) Value(value string) DiffEditor {
	return d.set("value", value)
}

// Visible sets whether the editor is visible
func (d DiffEditor) Visible(value bool) DiffEditor {
	return d.set("visible", value)
}

// VisibleOn sets the expression to determine if the editor is visible
func (d DiffEditor) VisibleOn(value string) DiffEditor {
	return d.set("visibleOn", value)
}

// Width sets the width in the table
func (d DiffEditor) Width(value string) DiffEditor {
	return d.set("width", value)
}
