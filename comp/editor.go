package comp

import "github.com/zrcoder/amisgo/model"

// editor represents a code editor
type editor model.Schema

func Editor() editor {
	return editor{"type": "editor"}
}

func JsonEditor() editor {
	return Editor().Language("json")
}

func (ec editor) set(key string, value any) editor {
	ec[key] = value
	return ec
}

// AllowFullscreen sets whether fullscreen mode is allowed
func (ec editor) AllowFullscreen(value bool) editor {
	return ec.set("allowFullscreen", value)
}

// AutoFill sets auto-fill value
func (ec editor) AutoFill(value string) editor {
	return ec.set("autoFill", value)
}

// ClassName sets the container CSS class name
func (ec editor) ClassName(value string) editor {
	return ec.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when the form item is hidden
func (ec editor) ClearValueOnHidden(value bool) editor {
	return ec.set("clearValueOnHidden", value)
}

// Desc sets the description
func (ec editor) Desc(value string) editor {
	return ec.set("desc", value)
}

// Description sets the description, supports HTML
func (ec editor) Description(value string) editor {
	return ec.set("description", value)
}

// DescriptionClassName sets the class name for the description
func (ec editor) DescriptionClassName(value string) editor {
	return ec.set("descriptionClassName", value)
}

// Disabled sets whether the editor is disabled/read-only
func (ec editor) Disabled(value bool) editor {
	return ec.set("disabled", value)
}

// DisabledOn sets the expression to determine if the editor is disabled
func (ec editor) DisabledOn(value string) editor {
	return ec.set("disabledOn", value)
}

// EditorDidMount sets the callback for when the editor is mounted
func (ec editor) EditorDidMount(value string) editor {
	return ec.set("editorDidMount", value)
}

// EditorSetting sets the editor configuration
func (ec editor) EditorSetting(value string) editor {
	return ec.set("editorSetting", value)
}

// ExtraName sets an additional field name
func (ec editor) ExtraName(value string) editor {
	return ec.set("extraName", value)
}

// Hidden sets whether the editor is hidden
func (ec editor) Hidden(value bool) editor {
	return ec.set("hidden", value)
}

// HiddenOn sets the expression to determine if the editor is hidden
func (ec editor) HiddenOn(value string) editor {
	return ec.set("hiddenOn", value)
}

// Hint sets the input hint
func (ec editor) Hint(value string) editor {
	return ec.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (ec editor) Horizontal(value string) editor {
	return ec.set("horizontal", value)
}

// ID sets the unique component ID
func (ec editor) ID(value string) editor {
	return ec.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (ec editor) InitAutoFill(value string) editor {
	return ec.set("initAutoFill", value)
}

// Inline sets whether the form control is in inline mode
func (ec editor) Inline(value bool) editor {
	return ec.set("inline", value)
}

// InputClassName sets the input class name
func (ec editor) InputClassName(value string) editor {
	return ec.set("inputClassName", value)
}

// Label sets the label
func (ec editor) Label(value string) editor {
	return ec.set("label", value)
}

// LabelAlign sets the label alignment
func (ec editor) LabelAlign(value string) editor {
	return ec.set("labelAlign", value)
}

// LabelClassName sets the label class name
func (ec editor) LabelClassName(value string) editor {
	return ec.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (ec editor) LabelRemark(value string) editor {
	return ec.set("labelRemark", value)
}

// LabelWidth sets the label width
func (ec editor) LabelWidth(value string) editor {
	return ec.set("labelWidth", value)
}

// Language sets the language type
func (ec editor) Language(value string) editor {
	return ec.set("language", value)
}

// Mode sets the display mode of the form item
func (ec editor) Mode(value string) editor {
	return ec.set("mode", value)
}

// Name sets the field name for form submission
func (ec editor) Name(value string) editor {
	return ec.set("name", value)
}

// OnEvent sets the event action configuration
func (ec editor) OnEvent(value any) editor {
	return ec.set("onEvent", value)
}

// Placeholder sets the placeholder
func (ec editor) Placeholder(value string) editor {
	return ec.set("placeholder", value)
}

// ReadOnly sets whether the editor is read-only
func (ec editor) ReadOnly(value bool) editor {
	return ec.set("readOnly", value)
}

// ReadOnlyOn sets the expression to determine if the editor is read-only
func (ec editor) ReadOnlyOn(value string) editor {
	return ec.set("readOnlyOn", value)
}

// Remark sets the remark
func (ec editor) Remark(value string) editor {
	return ec.set("remark", value)
}

// Required sets whether the field is required
func (ec editor) Required(value bool) editor {
	return ec.set("required", value)
}

// Row sets the number of rows
func (ec editor) Row(value string) editor {
	return ec.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (ec editor) SaveImmediately(value bool) editor {
	return ec.set("saveImmediately", value)
}

// Size sets the editor size
func (ec editor) Size(value string) editor {
	return ec.set("size", value)
}

// Static sets whether the editor is in static display mode
func (ec editor) Static(value bool) editor {
	return ec.set("static", value)
}

// StaticClassName sets the class name for static display
func (ec editor) StaticClassName(value string) editor {
	return ec.set("staticClassName", value)
}

// StaticInputClassName sets the class name for static input display
func (ec editor) StaticInputClassName(value string) editor {
	return ec.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class name for static label display
func (ec editor) StaticLabelClassName(value string) editor {
	return ec.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the editor is in static display mode
func (ec editor) StaticOn(value string) editor {
	return ec.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (ec editor) StaticPlaceholder(value string) editor {
	return ec.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (ec editor) StaticSchema(value string) editor {
	return ec.set("staticSchema", value)
}

// Style sets the component style
func (ec editor) Style(value any) editor {
	return ec.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change
func (ec editor) SubmitOnChange(value bool) editor {
	return ec.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (ec editor) TestIdBuilder(value string) editor {
	return ec.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI
func (ec editor) UseMobileUI(value bool) editor {
	return ec.set("useMobileUI", value)
}

// ValidateApi sets the API for remote validation
func (ec editor) ValidateApi(value string) editor {
	return ec.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (ec editor) ValidateOnChange(value bool) editor {
	return ec.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (ec editor) ValidationErrors(value string) editor {
	return ec.set("validationErrors", value)
}

// Validations sets the validation rules
func (ec editor) Validations(value string) editor {
	return ec.set("validations", value)
}

// Value sets the default value
func (ec editor) Value(value string) editor {
	return ec.set("value", value)
}

// Visible sets whether the editor is visible
func (ec editor) Visible(value bool) editor {
	return ec.set("visible", value)
}

// VisibleOn sets the expression to determine if the editor is visible
func (ec editor) VisibleOn(value string) editor {
	return ec.set("visibleOn", value)
}

// Width sets the width in a table
func (ec editor) Width(value string) editor {
	return ec.set("width", value)
}

// Options sets additional options for the Monaco editor
func (e editor) Options(value model.Schema) editor {
	return e.set("options", value)
}
