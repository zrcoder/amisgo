package comp

import "github.com/zrcoder/amisgo/schema"

// InputRepeat documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/repeat
type InputRepeat schema.Schema

func NewInputRepeat() InputRepeat {
	return InputRepeat{"type": "input-repeat"}
}

func (rc InputRepeat) set(key string, value any) InputRepeat {
	rc[key] = value
	return rc
}

// AutoFill sets autoFill when an option is selected
func (rc InputRepeat) AutoFill(value string) InputRepeat {
	return rc.set("autoFill", value)
}

// ClassName sets the container CSS class name
func (rc InputRepeat) ClassName(value string) InputRepeat {
	return rc.set("className", value)
}

// ClearValueOnHidden removes the form item value when hidden
func (rc InputRepeat) ClearValueOnHidden(value bool) InputRepeat {
	return rc.set("clearValueOnHidden", value)
}

// Desc sets the description
func (rc InputRepeat) Desc(value string) InputRepeat {
	return rc.set("desc", value)
}

// Description sets the description content, supports HTML
func (rc InputRepeat) Description(value string) InputRepeat {
	return rc.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description
func (rc InputRepeat) DescriptionClassName(value string) InputRepeat {
	return rc.set("descriptionClassName", value)
}

// Disabled sets whether the input is disabled
func (rc InputRepeat) Disabled(value bool) InputRepeat {
	return rc.set("disabled", value)
}

// DisabledOn sets the condition to disable the input
func (rc InputRepeat) DisabledOn(value string) InputRepeat {
	return rc.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (rc InputRepeat) EditorSetting(value string) InputRepeat {
	return rc.set("editorSetting", value)
}

// ExtraName sets an extra field name for range components
func (rc InputRepeat) ExtraName(value string) InputRepeat {
	return rc.set("extraName", value)
}

// Hidden sets whether the input is hidden
func (rc InputRepeat) Hidden(value bool) InputRepeat {
	return rc.set("hidden", value)
}

// HiddenOn sets the condition to hide the input
func (rc InputRepeat) HiddenOn(value string) InputRepeat {
	return rc.set("hiddenOn", value)
}

// Hint sets the input hint displayed on focus
func (rc InputRepeat) Hint(value string) InputRepeat {
	return rc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration
func (rc InputRepeat) Horizontal(value string) InputRepeat {
	return rc.set("horizontal", value)
}

// ID sets the unique component ID
func (rc InputRepeat) ID(value string) InputRepeat {
	return rc.set("id", value)
}

// InitAutoFill sets the initial autoFill value
func (rc InputRepeat) InitAutoFill(value string) InputRepeat {
	return rc.set("initAutoFill", value)
}

// Inline sets whether the form control is inline
func (rc InputRepeat) Inline(value bool) InputRepeat {
	return rc.set("inline", value)
}

// InputClassName sets the CSS class name for the input
func (rc InputRepeat) InputClassName(value string) InputRepeat {
	return rc.set("inputClassName", value)
}

// Label sets the label text
func (rc InputRepeat) Label(value string) InputRepeat {
	return rc.set("label", value)
}

// LabelAlign sets the label alignment
func (rc InputRepeat) LabelAlign(value string) InputRepeat {
	return rc.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label
func (rc InputRepeat) LabelClassName(value string) InputRepeat {
	return rc.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (rc InputRepeat) LabelRemark(value string) InputRepeat {
	return rc.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label
func (rc InputRepeat) LabelWidth(value string) InputRepeat {
	return rc.set("labelWidth", value)
}

// Mode sets the display mode of the form item
func (rc InputRepeat) Mode(value string) InputRepeat {
	return rc.set("mode", value)
}

// Name sets the field name for form submission
func (rc InputRepeat) Name(value string) InputRepeat {
	return rc.set("name", value)
}

// OnEvent sets the event action configuration
func (rc InputRepeat) OnEvent(value any) InputRepeat {
	return rc.set("onEvent", value)
}

// Options sets the options
func (rc InputRepeat) Options(value ...any) InputRepeat {
	return rc.set("options", value)
}

// Placeholder sets the placeholder text
func (rc InputRepeat) Placeholder(value string) InputRepeat {
	return rc.set("placeholder", value)
}

// ReadOnly sets whether the input is read-only
func (rc InputRepeat) ReadOnly(value bool) InputRepeat {
	return rc.set("readOnly", value)
}

// ReadOnlyOn sets the condition for read-only
func (rc InputRepeat) ReadOnlyOn(value string) InputRepeat {
	return rc.set("readOnlyOn", value)
}

// Remark sets the remark
func (rc InputRepeat) Remark(value string) InputRepeat {
	return rc.set("remark", value)
}

// Required sets whether the input is required
func (rc InputRepeat) Required(value bool) InputRepeat {
	return rc.set("required", value)
}

// Row sets the row value
func (rc InputRepeat) Row(value string) InputRepeat {
	return rc.set("row", value)
}

// SaveImmediately sets whether to save immediately in TableColumn
func (rc InputRepeat) SaveImmediately(value bool) InputRepeat {
	return rc.set("saveImmediately", value)
}

// Size sets the size of the form item
func (rc InputRepeat) Size(value string) InputRepeat {
	return rc.set("size", value)
}

// Static sets whether the input is displayed statically
func (rc InputRepeat) Static(value bool) InputRepeat {
	return rc.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (rc InputRepeat) StaticClassName(value string) InputRepeat {
	return rc.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (rc InputRepeat) StaticInputClassName(value string) InputRepeat {
	return rc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (rc InputRepeat) StaticLabelClassName(value string) InputRepeat {
	return rc.set("staticLabelClassName", value)
}

// StaticOn sets the condition for static display
func (rc InputRepeat) StaticOn(value string) InputRepeat {
	return rc.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (rc InputRepeat) StaticPlaceholder(value string) InputRepeat {
	return rc.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema
func (rc InputRepeat) StaticSchema(value string) InputRepeat {
	return rc.set("staticSchema", value)
}

// Style sets the component style
func (rc InputRepeat) Style(value any) InputRepeat {
	return rc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change
func (rc InputRepeat) SubmitOnChange(value bool) InputRepeat {
	return rc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (rc InputRepeat) TestIdBuilder(value string) InputRepeat {
	return rc.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI
func (rc InputRepeat) UseMobileUI(value bool) InputRepeat {
	return rc.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API
func (rc InputRepeat) ValidateApi(value string) InputRepeat {
	return rc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change
func (rc InputRepeat) ValidateOnChange(value bool) InputRepeat {
	return rc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages
func (rc InputRepeat) ValidationErrors(value string) InputRepeat {
	return rc.set("validationErrors", value)
}

// Validations sets the validations
func (rc InputRepeat) Validations(value string) InputRepeat {
	return rc.set("validations", value)
}

// Value sets the default value
func (rc InputRepeat) Value(value string) InputRepeat {
	return rc.set("value", value)
}

// Visible sets whether the input is visible
func (rc InputRepeat) Visible(value bool) InputRepeat {
	return rc.set("visible", value)
}

// VisibleOn sets the condition for visibility
func (rc InputRepeat) VisibleOn(value string) InputRepeat {
	return rc.set("visibleOn", value)
}

// Width sets the width in Table
func (rc InputRepeat) Width(value string) InputRepeat {
	return rc.set("width", value)
}
