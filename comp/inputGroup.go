package comp

import "github.com/zrcoder/amisgo/model"

// inputGroup represents the schema for an input group component.
type inputGroup model.Schema

// InputGroup initializes an input group component.
func InputGroup() inputGroup {
	return inputGroup{"type": "input-group"}
}

// set sets a key-value pair and returns the input group.
func (i inputGroup) set(key string, value any) inputGroup {
	i[key] = value
	return i
}

// AutoFill sets the auto-fill value.
func (i inputGroup) AutoFill(value string) inputGroup {
	return i.set("autoFill", value)
}

// Body sets the form items.
func (i inputGroup) Body(value ...any) inputGroup {
	return i.set("body", value)
}

// ClassName sets the CSS class name.
func (i inputGroup) ClassName(value string) inputGroup {
	return i.set("className", value)
}

// ClearValueOnHidden sets whether to clear the value when hidden.
func (i inputGroup) ClearValueOnHidden(value bool) inputGroup {
	return i.set("clearValueOnHidden", value)
}

// Desc sets the description.
func (i inputGroup) Desc(value string) inputGroup {
	return i.set("desc", value)
}

// Description sets the HTML description.
func (i inputGroup) Description(value string) inputGroup {
	return i.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (i inputGroup) DescriptionClassName(value string) inputGroup {
	return i.set("descriptionClassName", value)
}

// Disabled sets whether the input group is disabled.
func (i inputGroup) Disabled(value bool) inputGroup {
	return i.set("disabled", value)
}

// DisabledOn sets the expression to disable the input group.
func (i inputGroup) DisabledOn(value string) inputGroup {
	return i.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (i inputGroup) EditorSetting(value string) inputGroup {
	return i.set("editorSetting", value)
}

// ExtraName sets an extra field name for range components.
func (i inputGroup) ExtraName(value string) inputGroup {
	return i.set("extraName", value)
}

// Hidden sets whether the input group is hidden.
func (i inputGroup) Hidden(value bool) inputGroup {
	return i.set("hidden", value)
}

// HiddenOn sets the expression to hide the input group.
func (i inputGroup) HiddenOn(value string) inputGroup {
	return i.set("hiddenOn", value)
}

// Hint sets the input hint.
func (i inputGroup) Hint(value string) inputGroup {
	return i.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (i inputGroup) Horizontal(value string) inputGroup {
	return i.set("horizontal", value)
}

// ID sets the unique ID for the component.
func (i inputGroup) ID(value string) inputGroup {
	return i.set("id", value)
}

// InitAutoFill sets the initial auto-fill value.
func (i inputGroup) InitAutoFill(value string) inputGroup {
	return i.set("initAutoFill", value)
}

// Inline sets whether the form control is inline.
func (i inputGroup) Inline(value bool) inputGroup {
	return i.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (i inputGroup) InputClassName(value string) inputGroup {
	return i.set("inputClassName", value)
}

// Label sets the label text.
func (i inputGroup) Label(value string) inputGroup {
	return i.set("label", value)
}

// LabelAlign sets the label alignment.
func (i inputGroup) LabelAlign(value string) inputGroup {
	return i.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (i inputGroup) LabelClassName(value string) inputGroup {
	return i.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (i inputGroup) LabelRemark(value string) inputGroup {
	return i.set("labelRemark", value)
}

// LabelWidth sets the label width.
func (i inputGroup) LabelWidth(value string) inputGroup {
	return i.set("labelWidth", value)
}

// Mode sets the display mode.
func (i inputGroup) Mode(value string) inputGroup {
	return i.set("mode", value)
}

// Name sets the field name.
func (i inputGroup) Name(value string) inputGroup {
	return i.set("name", value)
}

// OnEvent sets the event configuration.
func (i inputGroup) OnEvent(value any) inputGroup {
	return i.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (i inputGroup) Placeholder(value string) inputGroup {
	return i.set("placeholder", value)
}

// ReadOnly sets whether the input group is read-only.
func (i inputGroup) ReadOnly(value bool) inputGroup {
	return i.set("readOnly", value)
}

// ReadOnlyOn sets the expression to make the input group read-only.
func (i inputGroup) ReadOnlyOn(value string) inputGroup {
	return i.set("readOnlyOn", value)
}

// Remark sets the remark text.
func (i inputGroup) Remark(value string) inputGroup {
	return i.set("remark", value)
}

// Required sets whether the input group is required.
func (i inputGroup) Required(value bool) inputGroup {
	return i.set("required", value)
}

// Row sets the row configuration.
func (i inputGroup) Row(value string) inputGroup {
	return i.set("row", value)
}

// SaveImmediately sets whether to save immediately.
func (i inputGroup) SaveImmediately(value bool) inputGroup {
	return i.set("saveImmediately", value)
}

// Size sets the size of the input group.
func (i inputGroup) Size(value string) inputGroup {
	return i.set("size", value)
}

// Static sets whether the input group is static.
func (i inputGroup) Static(value bool) inputGroup {
	return i.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (i inputGroup) StaticClassName(value string) inputGroup {
	return i.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value.
func (i inputGroup) StaticInputClassName(value string) inputGroup {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label.
func (i inputGroup) StaticLabelClassName(value string) inputGroup {
	return i.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display.
func (i inputGroup) StaticOn(value string) inputGroup {
	return i.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (i inputGroup) StaticPlaceholder(value string) inputGroup {
	return i.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display.
func (i inputGroup) StaticSchema(value string) inputGroup {
	return i.set("staticSchema", value)
}

// Style sets the component style.
func (i inputGroup) Style(value any) inputGroup {
	return i.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (i inputGroup) SubmitOnChange(value bool) inputGroup {
	return i.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (i inputGroup) TestIdBuilder(value string) inputGroup {
	return i.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI.
func (i inputGroup) UseMobileUI(value bool) inputGroup {
	return i.set("useMobileUI", value)
}

// ValidateApi sets the API for remote validation.
func (i inputGroup) ValidateApi(value string) inputGroup {
	return i.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (i inputGroup) ValidateOnChange(value bool) inputGroup {
	return i.set("validateOnChange", value)
}

// ValidationConfig sets the validation configuration.
func (i inputGroup) ValidationConfig(value string) inputGroup {
	return i.set("validationConfig", value)
}

// ValidationErrors sets the validation error messages.
func (i inputGroup) ValidationErrors(value string) inputGroup {
	return i.set("validationErrors", value)
}

// Validations sets the validation rules.
func (i inputGroup) Validations(value string) inputGroup {
	return i.set("validations", value)
}

// Value sets the default value.
func (i inputGroup) Value(value string) inputGroup {
	return i.set("value", value)
}

// Visible sets whether the input group is visible.
func (i inputGroup) Visible(value bool) inputGroup {
	return i.set("visible", value)
}

// VisibleOn sets the expression to make the input group visible.
func (i inputGroup) VisibleOn(value string) inputGroup {
	return i.set("visibleOn", value)
}

// Width sets the width in a table.
func (i inputGroup) Width(value string) inputGroup {
	return i.set("width", value)
}
