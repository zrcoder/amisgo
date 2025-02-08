package comp

import "github.com/zrcoder/amisgo/model"

// InputGroup represents the schema for an input group component.
type InputGroup model.Schema

func NewInputGroup() InputGroup {
	return InputGroup{"type": "input-group"}
}

// set sets a key-value pair and returns the input group.
func (i InputGroup) set(key string, value any) InputGroup {
	i[key] = value
	return i
}

// AutoFill sets the auto-fill value.
func (i InputGroup) AutoFill(value string) InputGroup {
	return i.set("autoFill", value)
}

// Body sets the form items.
func (i InputGroup) Body(value ...any) InputGroup {
	return i.set("body", value)
}

// ClassName sets the CSS class name.
func (i InputGroup) ClassName(value string) InputGroup {
	return i.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden.
func (i InputGroup) ClearValueOnHidden(value bool) InputGroup {
	return i.set("clearValueOnHidden", value)
}

// Desc sets the description.
func (i InputGroup) Desc(value string) InputGroup {
	return i.set("desc", value)
}

// Description sets the HTML description.
func (i InputGroup) Description(value string) InputGroup {
	return i.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (i InputGroup) DescriptionClassName(value string) InputGroup {
	return i.set("descriptionClassName", value)
}

// Disabled sets whether the input group is disabled.
func (i InputGroup) Disabled(value bool) InputGroup {
	return i.set("disabled", value)
}

// DisabledOn sets the expression to disable the input group.
func (i InputGroup) DisabledOn(value string) InputGroup {
	return i.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (i InputGroup) EditorSetting(value string) InputGroup {
	return i.set("editorSetting", value)
}

// ExtraName sets an extra field name for range components.
func (i InputGroup) ExtraName(value string) InputGroup {
	return i.set("extraName", value)
}

// Hidden sets whether the input group is hidden.
func (i InputGroup) Hidden(value bool) InputGroup {
	return i.set("hidden", value)
}

// HiddenOn sets the expression to hide the input group.
func (i InputGroup) HiddenOn(value string) InputGroup {
	return i.set("hiddenOn", value)
}

// Hint sets the input hint.
func (i InputGroup) Hint(value string) InputGroup {
	return i.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (i InputGroup) Horizontal(value string) InputGroup {
	return i.set("horizontal", value)
}

// ID sets the unique ID for the component.
func (i InputGroup) ID(value string) InputGroup {
	return i.set("id", value)
}

// InitAutoFill sets the initial auto-fill value.
func (i InputGroup) InitAutoFill(value string) InputGroup {
	return i.set("initAutoFill", value)
}

// Inline sets whether the form control is inline.
func (i InputGroup) Inline(value bool) InputGroup {
	return i.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (i InputGroup) InputClassName(value string) InputGroup {
	return i.set("inputClassName", value)
}

// Label sets the label text.
func (i InputGroup) Label(value string) InputGroup {
	return i.set("label", value)
}

// LabelAlign sets the label alignment.
func (i InputGroup) LabelAlign(value string) InputGroup {
	return i.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (i InputGroup) LabelClassName(value string) InputGroup {
	return i.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (i InputGroup) LabelRemark(value string) InputGroup {
	return i.set("labelRemark", value)
}

// LabelWidth sets the label width.
func (i InputGroup) LabelWidth(value string) InputGroup {
	return i.set("labelWidth", value)
}

// Mode sets the display mode.
func (i InputGroup) Mode(value string) InputGroup {
	return i.set("mode", value)
}

// Name sets the field name.
func (i InputGroup) Name(value string) InputGroup {
	return i.set("name", value)
}

// OnEvent sets the event configuration.
func (i InputGroup) OnEvent(value any) InputGroup {
	return i.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (i InputGroup) Placeholder(value string) InputGroup {
	return i.set("placeholder", value)
}

// ReadOnly sets whether the input group is read-only.
func (i InputGroup) ReadOnly(value bool) InputGroup {
	return i.set("readOnly", value)
}

// ReadOnlyOn sets the expression to make the input group read-only.
func (i InputGroup) ReadOnlyOn(value string) InputGroup {
	return i.set("readOnlyOn", value)
}

// Remark sets the remark text.
func (i InputGroup) Remark(value string) InputGroup {
	return i.set("remark", value)
}

// Required sets whether the input group is required.
func (i InputGroup) Required(value bool) InputGroup {
	return i.set("required", value)
}

// Row sets the row configuration.
func (i InputGroup) Row(value string) InputGroup {
	return i.set("row", value)
}

// SaveImmediately sets whether to save immediately.
func (i InputGroup) SaveImmediately(value bool) InputGroup {
	return i.set("saveImmediately", value)
}

// Size sets the size of the input group.
func (i InputGroup) Size(value string) InputGroup {
	return i.set("size", value)
}

// Static sets whether the input group is static.
func (i InputGroup) Static(value bool) InputGroup {
	return i.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (i InputGroup) StaticClassName(value string) InputGroup {
	return i.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value.
func (i InputGroup) StaticInputClassName(value string) InputGroup {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label.
func (i InputGroup) StaticLabelClassName(value string) InputGroup {
	return i.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display.
func (i InputGroup) StaticOn(value string) InputGroup {
	return i.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (i InputGroup) StaticPlaceholder(value string) InputGroup {
	return i.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display.
func (i InputGroup) StaticSchema(value string) InputGroup {
	return i.set("staticSchema", value)
}

// Style sets the component style.
func (i InputGroup) Style(value any) InputGroup {
	return i.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (i InputGroup) SubmitOnChange(value bool) InputGroup {
	return i.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (i InputGroup) TestIdBuilder(value string) InputGroup {
	return i.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI.
func (i InputGroup) UseMobileUI(value bool) InputGroup {
	return i.set("useMobileUI", value)
}

// ValidateApi sets the API for remote validation.
func (i InputGroup) ValidateApi(value string) InputGroup {
	return i.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (i InputGroup) ValidateOnChange(value bool) InputGroup {
	return i.set("validateOnChange", value)
}

// ValidationConfig sets the validation configuration.
func (i InputGroup) ValidationConfig(value string) InputGroup {
	return i.set("validationConfig", value)
}

// ValidationErrors sets the validation error messages.
func (i InputGroup) ValidationErrors(value string) InputGroup {
	return i.set("validationErrors", value)
}

// Validations sets the validation rules.
func (i InputGroup) Validations(value string) InputGroup {
	return i.set("validations", value)
}

// Value sets the default value.
func (i InputGroup) Value(value string) InputGroup {
	return i.set("value", value)
}

// Visible sets whether the input group is visible.
func (i InputGroup) Visible(value bool) InputGroup {
	return i.set("visible", value)
}

// VisibleOn sets the expression to make the input group visible.
func (i InputGroup) VisibleOn(value string) InputGroup {
	return i.set("visibleOn", value)
}

// Width sets the width in a table.
func (i InputGroup) Width(value string) InputGroup {
	return i.set("width", value)
}
