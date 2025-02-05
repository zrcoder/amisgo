package comp

import "github.com/zrcoder/amisgo/model"

// Textarea represents a multi-line text input field. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/Textarea
type Textarea model.Schema

func NewTextarea() Textarea {
	return Textarea{"type": "textarea"}
}

// set sets a key-value pair for the textarea component.
func (t Textarea) set(key string, value any) Textarea {
	t[key] = value
	return t
}

// AutoFill sets the autoFill property.
func (t Textarea) AutoFill(value string) Textarea {
	return t.set("autoFill", value)
}

// BorderMode sets the border mode. Options: full | half | none
func (t Textarea) BorderMode(value string) Textarea {
	return t.set("borderMode", value)
}

// ClassName sets the CSS class name for the container.
func (t Textarea) ClassName(value string) Textarea {
	return t.set("className", value)
}

// ClearValueOnHidden removes the value when the field is hidden.
func (t Textarea) ClearValueOnHidden(value bool) Textarea {
	return t.set("clearValueOnHidden", value)
}

// Clearable sets whether the input content can be cleared.
func (t Textarea) Clearable(value bool) Textarea {
	return t.set("clearable", value)
}

// Desc sets the description.
func (t Textarea) Desc(value string) Textarea {
	return t.set("desc", value)
}

// Description sets the description content, supports HTML.
func (t Textarea) Description(value string) Textarea {
	return t.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (t Textarea) DescriptionClassName(value string) Textarea {
	return t.set("descriptionClassName", value)
}

// Disabled sets whether the component is disabled.
func (t Textarea) Disabled(value bool) Textarea {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to disable the component.
func (t Textarea) DisabledOn(value string) Textarea {
	return t.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (t Textarea) EditorSetting(value string) Textarea {
	return t.set("editorSetting", value)
}

// ExtraName sets an extra field name for range components.
func (t Textarea) ExtraName(value string) Textarea {
	return t.set("extraName", value)
}

// Hidden sets whether the component is hidden.
func (t Textarea) Hidden(value bool) Textarea {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to hide the component.
func (t Textarea) HiddenOn(value string) Textarea {
	return t.set("hiddenOn", value)
}

// Hint sets the input hint displayed on focus.
func (t Textarea) Hint(value string) Textarea {
	return t.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (t Textarea) Horizontal(value string) Textarea {
	return t.set("horizontal", value)
}

// ID sets the unique ID for the component.
func (t Textarea) ID(value string) Textarea {
	return t.set("id", value)
}

// InitAutoFill sets the initial autoFill value.
func (t Textarea) InitAutoFill(value string) Textarea {
	return t.set("initAutoFill", value)
}

// Inline sets whether the form control is in inline mode.
func (t Textarea) Inline(value bool) Textarea {
	return t.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (t Textarea) InputClassName(value string) Textarea {
	return t.set("inputClassName", value)
}

// Label sets the label text.
func (t Textarea) Label(value string) Textarea {
	return t.set("label", value)
}

// LabelAlign sets the label alignment. Options: right | left | top | inherit
func (t Textarea) LabelAlign(value string) Textarea {
	return t.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (t Textarea) LabelClassName(value string) Textarea {
	return t.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (t Textarea) LabelRemark(value string) Textarea {
	return t.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label.
func (t Textarea) LabelWidth(value string) Textarea {
	return t.set("labelWidth", value)
}

// MaxLength sets the maximum number of characters.
func (t Textarea) MaxLength(value string) Textarea {
	return t.set("maxLength", value)
}

// MaxRows sets the maximum number of rows.
func (t Textarea) MaxRows(value string) Textarea {
	return t.set("maxRows", value)
}

// MinRows sets the minimum number of rows.
func (t Textarea) MinRows(value string) Textarea {
	return t.set("minRows", value)
}

// Mode sets the display mode. Options: normal | inline | horizontal
func (t Textarea) Mode(value string) Textarea {
	return t.set("mode", value)
}

// Name sets the field name for form submission.
func (t Textarea) Name(value string) Textarea {
	return t.set("name", value)
}

// OnEvent sets the event action configuration.
func (t Textarea) OnEvent(value any) Textarea {
	return t.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (t Textarea) Placeholder(value string) Textarea {
	return t.set("placeholder", value)
}

// ReadOnly sets whether the component is read-only.
func (t Textarea) ReadOnly(value bool) Textarea {
	return t.set("readOnly", value)
}

// ReadOnlyOn sets the expression to make the component read-only.
func (t Textarea) ReadOnlyOn(value string) Textarea {
	return t.set("readOnlyOn", value)
}

// Remark sets the remark.
func (t Textarea) Remark(value string) Textarea {
	return t.set("remark", value)
}

// Required sets whether the field is required.
func (t Textarea) Required(value bool) Textarea {
	return t.set("required", value)
}

// ResetValue sets the reset value.
func (t Textarea) ResetValue(value string) Textarea {
	return t.set("resetValue", value)
}

// Row sets the row value.
func (t Textarea) Row(value string) Textarea {
	return t.set("row", value)
}

// SaveImmediately sets whether to save immediately (used in TableColumn).
func (t Textarea) SaveImmediately(value bool) Textarea {
	return t.set("saveImmediately", value)
}

// ShowCounter sets whether to show the counter.
func (t Textarea) ShowCounter(value bool) Textarea {
	return t.set("showCounter", value)
}

// Size sets the size of the form item. Options: xs | sm | md | lg | full
func (t Textarea) Size(value string) Textarea {
	return t.set("size", value)
}

// Static sets whether the component is displayed statically.
func (t Textarea) Static(value bool) Textarea {
	return t.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (t Textarea) StaticClassName(value string) Textarea {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display.
func (t Textarea) StaticInputClassName(value string) Textarea {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display.
func (t Textarea) StaticLabelClassName(value string) Textarea {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display.
func (t Textarea) StaticOn(value string) Textarea {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (t Textarea) StaticPlaceholder(value string) Textarea {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (t Textarea) StaticSchema(value string) Textarea {
	return t.set("staticSchema", value)
}

// Style sets the component style.
func (t Textarea) Style(value any) Textarea {
	return t.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (t Textarea) SubmitOnChange(value bool) Textarea {
	return t.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (t Textarea) TestIdBuilder(value string) Textarea {
	return t.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI styles.
func (t Textarea) UseMobileUI(value bool) Textarea {
	return t.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API.
func (t Textarea) ValidateApi(value string) Textarea {
	return t.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (t Textarea) ValidateOnChange(value bool) Textarea {
	return t.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (t Textarea) ValidationErrors(value string) Textarea {
	return t.set("validationErrors", value)
}

// Validations sets the validation rules.
func (t Textarea) Validations(value string) Textarea {
	return t.set("validations", value)
}

// Value sets the default value.
func (t Textarea) Value(value string) Textarea {
	return t.set("value", value)
}

// Visible sets whether the component is visible.
func (t Textarea) Visible(value bool) Textarea {
	return t.set("visible", value)
}

// VisibleOn sets the expression to make the component visible.
func (t Textarea) VisibleOn(value string) Textarea {
	return t.set("visibleOn", value)
}

// Width sets the width in a table.
func (t Textarea) Width(value string) Textarea {
	return t.set("width", value)
}
