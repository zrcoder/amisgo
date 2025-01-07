package comp

// textarea represents a multi-line text input field. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/textarea

type textarea Schema

// Textarea creates a new textarea component.
func Textarea() textarea {
	return textarea{}.set("type", "textarea")
}

// set sets a key-value pair for the textarea component.
func (t textarea) set(key string, value any) textarea {
	t[key] = value
	return t
}

// AutoFill sets the autoFill property.
func (t textarea) AutoFill(value string) textarea {
	return t.set("autoFill", value)
}

// BorderMode sets the border mode. Options: full | half | none
func (t textarea) BorderMode(value string) textarea {
	return t.set("borderMode", value)
}

// ClassName sets the CSS class name for the container.
func (t textarea) ClassName(value string) textarea {
	return t.set("className", value)
}

// ClearValueOnHidden removes the value when the field is hidden.
func (t textarea) ClearValueOnHidden(value bool) textarea {
	return t.set("clearValueOnHidden", value)
}

// Clearable sets whether the input content can be cleared.
func (t textarea) Clearable(value bool) textarea {
	return t.set("clearable", value)
}

// Desc sets the description.
func (t textarea) Desc(value string) textarea {
	return t.set("desc", value)
}

// Description sets the description content, supports HTML.
func (t textarea) Description(value string) textarea {
	return t.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (t textarea) DescriptionClassName(value string) textarea {
	return t.set("descriptionClassName", value)
}

// Disabled sets whether the component is disabled.
func (t textarea) Disabled(value bool) textarea {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to disable the component.
func (t textarea) DisabledOn(value string) textarea {
	return t.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (t textarea) EditorSetting(value string) textarea {
	return t.set("editorSetting", value)
}

// ExtraName sets an extra field name for range components.
func (t textarea) ExtraName(value string) textarea {
	return t.set("extraName", value)
}

// Hidden sets whether the component is hidden.
func (t textarea) Hidden(value bool) textarea {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to hide the component.
func (t textarea) HiddenOn(value string) textarea {
	return t.set("hiddenOn", value)
}

// Hint sets the input hint displayed on focus.
func (t textarea) Hint(value string) textarea {
	return t.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (t textarea) Horizontal(value string) textarea {
	return t.set("horizontal", value)
}

// ID sets the unique ID for the component.
func (t textarea) ID(value string) textarea {
	return t.set("id", value)
}

// InitAutoFill sets the initial autoFill value.
func (t textarea) InitAutoFill(value string) textarea {
	return t.set("initAutoFill", value)
}

// Inline sets whether the form control is in inline mode.
func (t textarea) Inline(value bool) textarea {
	return t.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (t textarea) InputClassName(value string) textarea {
	return t.set("inputClassName", value)
}

// Label sets the label text.
func (t textarea) Label(value string) textarea {
	return t.set("label", value)
}

// LabelAlign sets the label alignment. Options: right | left | top | inherit
func (t textarea) LabelAlign(value string) textarea {
	return t.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (t textarea) LabelClassName(value string) textarea {
	return t.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (t textarea) LabelRemark(value string) textarea {
	return t.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label.
func (t textarea) LabelWidth(value string) textarea {
	return t.set("labelWidth", value)
}

// MaxLength sets the maximum number of characters.
func (t textarea) MaxLength(value string) textarea {
	return t.set("maxLength", value)
}

// MaxRows sets the maximum number of rows.
func (t textarea) MaxRows(value string) textarea {
	return t.set("maxRows", value)
}

// MinRows sets the minimum number of rows.
func (t textarea) MinRows(value string) textarea {
	return t.set("minRows", value)
}

// Mode sets the display mode. Options: normal | inline | horizontal
func (t textarea) Mode(value string) textarea {
	return t.set("mode", value)
}

// Name sets the field name for form submission.
func (t textarea) Name(value string) textarea {
	return t.set("name", value)
}

// OnEvent sets the event action configuration.
func (t textarea) OnEvent(value any) textarea {
	return t.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (t textarea) Placeholder(value string) textarea {
	return t.set("placeholder", value)
}

// ReadOnly sets whether the component is read-only.
func (t textarea) ReadOnly(value bool) textarea {
	return t.set("readOnly", value)
}

// ReadOnlyOn sets the expression to make the component read-only.
func (t textarea) ReadOnlyOn(value string) textarea {
	return t.set("readOnlyOn", value)
}

// Remark sets the remark.
func (t textarea) Remark(value string) textarea {
	return t.set("remark", value)
}

// Required sets whether the field is required.
func (t textarea) Required(value bool) textarea {
	return t.set("required", value)
}

// ResetValue sets the reset value.
func (t textarea) ResetValue(value string) textarea {
	return t.set("resetValue", value)
}

// Row sets the row value.
func (t textarea) Row(value string) textarea {
	return t.set("row", value)
}

// SaveImmediately sets whether to save immediately (used in TableColumn).
func (t textarea) SaveImmediately(value bool) textarea {
	return t.set("saveImmediately", value)
}

// ShowCounter sets whether to show the counter.
func (t textarea) ShowCounter(value bool) textarea {
	return t.set("showCounter", value)
}

// Size sets the size of the form item. Options: xs | sm | md | lg | full
func (t textarea) Size(value string) textarea {
	return t.set("size", value)
}

// Static sets whether the component is displayed statically.
func (t textarea) Static(value bool) textarea {
	return t.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (t textarea) StaticClassName(value string) textarea {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display.
func (t textarea) StaticInputClassName(value string) textarea {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display.
func (t textarea) StaticLabelClassName(value string) textarea {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display.
func (t textarea) StaticOn(value string) textarea {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (t textarea) StaticPlaceholder(value string) textarea {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (t textarea) StaticSchema(value string) textarea {
	return t.set("staticSchema", value)
}

// Style sets the component style.
func (t textarea) Style(value any) textarea {
	return t.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (t textarea) SubmitOnChange(value bool) textarea {
	return t.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (t textarea) TestIdBuilder(value string) textarea {
	return t.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI styles.
func (t textarea) UseMobileUI(value bool) textarea {
	return t.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API.
func (t textarea) ValidateApi(value string) textarea {
	return t.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (t textarea) ValidateOnChange(value bool) textarea {
	return t.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (t textarea) ValidationErrors(value string) textarea {
	return t.set("validationErrors", value)
}

// Validations sets the validation rules.
func (t textarea) Validations(value string) textarea {
	return t.set("validations", value)
}

// Value sets the default value.
func (t textarea) Value(value string) textarea {
	return t.set("value", value)
}

// Visible sets whether the component is visible.
func (t textarea) Visible(value bool) textarea {
	return t.set("visible", value)
}

// VisibleOn sets the expression to make the component visible.
func (t textarea) VisibleOn(value string) textarea {
	return t.set("visibleOn", value)
}

// Width sets the width in a table.
func (t textarea) Width(value string) textarea {
	return t.set("width", value)
}
