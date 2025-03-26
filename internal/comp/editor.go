package comp

import "github.com/zrcoder/amisgo/schema"

// Editor represents a code Editor
type Editor schema.Schema

func NewEditor() Editor {
	return Editor{"type": "editor"}
}

func NewJsonEditor() Editor {
	return NewEditor().Language("json")
}

func (ec Editor) set(key string, value any) Editor {
	ec[key] = value
	return ec
}

// AllowFullscreen sets whether fullscreen mode is allowed
func (ec Editor) AllowFullscreen(value bool) Editor {
	return ec.set("allowFullscreen", value)
}

// AutoFill sets auto-fill value
func (ec Editor) AutoFill(value string) Editor {
	return ec.set("autoFill", value)
}

// ClassName sets the container CSS class name
func (ec Editor) ClassName(value string) Editor {
	return ec.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when the form item is hidden
func (ec Editor) ClearValueOnHidden(value bool) Editor {
	return ec.set("clearValueOnHidden", value)
}

// Desc sets the description
func (ec Editor) Desc(value string) Editor {
	return ec.set("desc", value)
}

// Description sets the description, supports HTML
func (ec Editor) Description(value string) Editor {
	return ec.set("description", value)
}

// DescriptionClassName sets the class name for the description
func (ec Editor) DescriptionClassName(value string) Editor {
	return ec.set("descriptionClassName", value)
}

// Disabled sets whether the editor is disabled/read-only
func (ec Editor) Disabled(value bool) Editor {
	return ec.set("disabled", value)
}

// DisabledOn sets the expression to determine if the editor is disabled
func (ec Editor) DisabledOn(value string) Editor {
	return ec.set("disabledOn", value)
}

// EditorDidMount sets the callback for when the editor is mounted
func (ec Editor) EditorDidMount(value string) Editor {
	return ec.set("editorDidMount", value)
}

// EditorSetting sets the editor configuration
func (ec Editor) EditorSetting(value string) Editor {
	return ec.set("editorSetting", value)
}

// ExtraName sets an additional field name
func (ec Editor) ExtraName(value string) Editor {
	return ec.set("extraName", value)
}

// Hidden sets whether the editor is hidden
func (ec Editor) Hidden(value bool) Editor {
	return ec.set("hidden", value)
}

// HiddenOn sets the expression to determine if the editor is hidden
func (ec Editor) HiddenOn(value string) Editor {
	return ec.set("hiddenOn", value)
}

// Hint sets the input hint
func (ec Editor) Hint(value string) Editor {
	return ec.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (ec Editor) Horizontal(value string) Editor {
	return ec.set("horizontal", value)
}

// ID sets the unique component ID
func (ec Editor) ID(value string) Editor {
	return ec.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (ec Editor) InitAutoFill(value string) Editor {
	return ec.set("initAutoFill", value)
}

// Inline sets whether the form control is in inline mode
func (ec Editor) Inline(value bool) Editor {
	return ec.set("inline", value)
}

// InputClassName sets the input class name
func (ec Editor) InputClassName(value string) Editor {
	return ec.set("inputClassName", value)
}

// Label sets the label
func (ec Editor) Label(value string) Editor {
	return ec.set("label", value)
}

// LabelAlign sets the label alignment
func (ec Editor) LabelAlign(value string) Editor {
	return ec.set("labelAlign", value)
}

// LabelClassName sets the label class name
func (ec Editor) LabelClassName(value string) Editor {
	return ec.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (ec Editor) LabelRemark(value string) Editor {
	return ec.set("labelRemark", value)
}

// LabelWidth sets the label width
func (ec Editor) LabelWidth(value string) Editor {
	return ec.set("labelWidth", value)
}

// Language sets the language type
func (ec Editor) Language(value string) Editor {
	return ec.set("language", value)
}

// Mode sets the display mode of the form item
func (ec Editor) Mode(value string) Editor {
	return ec.set("mode", value)
}

// Name sets the field name for form submission
func (ec Editor) Name(value string) Editor {
	return ec.set("name", value)
}

// OnEvent sets the event action configuration
func (ec Editor) OnEvent(value Event) Editor {
	return ec.set("onEvent", value)
}

// Placeholder sets the placeholder
func (ec Editor) Placeholder(value string) Editor {
	return ec.set("placeholder", value)
}

// ReadOnly sets whether the editor is read-only
func (ec Editor) ReadOnly(value bool) Editor {
	return ec.set("readOnly", value)
}

// ReadOnlyOn sets the expression to determine if the editor is read-only
func (ec Editor) ReadOnlyOn(value string) Editor {
	return ec.set("readOnlyOn", value)
}

// Remark sets the remark
func (ec Editor) Remark(value string) Editor {
	return ec.set("remark", value)
}

// Required sets whether the field is required
func (ec Editor) Required(value bool) Editor {
	return ec.set("required", value)
}

// Row sets the number of rows
func (ec Editor) Row(value string) Editor {
	return ec.set("row", value)
}

// SaveImmediately sets whether to save immediately
func (ec Editor) SaveImmediately(value bool) Editor {
	return ec.set("saveImmediately", value)
}

// Size sets the editor size.  can be: md、lg、xl、xxl, default is md
func (ec Editor) Size(value string) Editor {
	return ec.set("size", value)
}

// Static sets whether the editor is in static display mode
func (ec Editor) Static(value bool) Editor {
	return ec.set("static", value)
}

// StaticClassName sets the class name for static display
func (ec Editor) StaticClassName(value string) Editor {
	return ec.set("staticClassName", value)
}

// StaticInputClassName sets the class name for static input display
func (ec Editor) StaticInputClassName(value string) Editor {
	return ec.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class name for static label display
func (ec Editor) StaticLabelClassName(value string) Editor {
	return ec.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the editor is in static display mode
func (ec Editor) StaticOn(value string) Editor {
	return ec.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (ec Editor) StaticPlaceholder(value string) Editor {
	return ec.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (ec Editor) StaticSchema(value string) Editor {
	return ec.set("staticSchema", value)
}

// Style sets the component style
func (ec Editor) Style(value any) Editor {
	return ec.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change
func (ec Editor) SubmitOnChange(value bool) Editor {
	return ec.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (ec Editor) TestIdBuilder(value string) Editor {
	return ec.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI
func (ec Editor) UseMobileUI(value bool) Editor {
	return ec.set("useMobileUI", value)
}

// ValidateApi sets the API for remote validation
func (ec Editor) ValidateApi(value string) Editor {
	return ec.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (ec Editor) ValidateOnChange(value bool) Editor {
	return ec.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (ec Editor) ValidationErrors(value string) Editor {
	return ec.set("validationErrors", value)
}

// Validations sets the validation rules
func (ec Editor) Validations(value string) Editor {
	return ec.set("validations", value)
}

// Value sets the default value
func (ec Editor) Value(value string) Editor {
	return ec.set("value", value)
}

// Visible sets whether the editor is visible
func (ec Editor) Visible(value bool) Editor {
	return ec.set("visible", value)
}

// VisibleOn sets the expression to determine if the editor is visible
func (ec Editor) VisibleOn(value string) Editor {
	return ec.set("visibleOn", value)
}

// Width sets the width in a table
func (ec Editor) Width(value string) Editor {
	return ec.set("width", value)
}

// Options sets additional options for the Monaco editor
func (e Editor) Options(value schema.Schema) Editor {
	return e.set("options", value)
}
