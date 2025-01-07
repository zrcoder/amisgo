package comp

// hidden represents a hidden field component.
type hidden Schema

func Hidden() hidden {
	return make(hidden).set("type", "hidden")
}

func (h hidden) set(key string, value any) hidden {
	h[key] = value
	return h
}

// AutoFill sets the autoFill property.
func (h hidden) AutoFill(value string) hidden {
	return h.set("autoFill", value)
}

// ClassName sets the container CSS class name.
func (h hidden) ClassName(value string) hidden {
	return h.set("className", value)
}

// ClearValueOnHidden sets whether to remove the field value when hidden.
func (h hidden) ClearValueOnHidden(value bool) hidden {
	return h.set("clearValueOnHidden", value)
}

// Desc sets the description.
func (h hidden) Desc(value string) hidden {
	return h.set("desc", value)
}

// Description sets the description content, supports HTML.
func (h hidden) Description(value string) hidden {
	return h.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (h hidden) DescriptionClassName(value string) hidden {
	return h.set("descriptionClassName", value)
}

// Disabled sets whether the field is disabled.
func (h hidden) Disabled(value bool) hidden {
	return h.set("disabled", value)
}

// DisabledOn sets the expression to determine if the field is disabled.
func (h hidden) DisabledOn(value string) hidden {
	return h.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (h hidden) EditorSetting(value string) hidden {
	return h.set("editorSetting", value)
}

// ExtraName sets an additional field name for range components.
func (h hidden) ExtraName(value string) hidden {
	return h.set("extraName", value)
}

// Hidden sets whether the field is hidden.
func (h hidden) Hidden(value bool) hidden {
	return h.set("hidden", value)
}

// HiddenOn sets the expression to determine if the field is hidden.
func (h hidden) HiddenOn(value string) hidden {
	return h.set("hiddenOn", value)
}

// Hint sets the input hint displayed on focus.
func (h hidden) Hint(value string) hidden {
	return h.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (h hidden) Horizontal(value string) hidden {
	return h.set("horizontal", value)
}

// ID sets the unique component ID.
func (h hidden) ID(value string) hidden {
	return h.set("id", value)
}

// InitAutoFill sets the initial autoFill value.
func (h hidden) InitAutoFill(value string) hidden {
	return h.set("initAutoFill", value)
}

// Inline sets whether the control is in inline mode.
func (h hidden) Inline(value bool) hidden {
	return h.set("inline", value)
}

// InputClassName sets the CSS class name for the input.
func (h hidden) InputClassName(value string) hidden {
	return h.set("inputClassName", value)
}

// Label sets the label text.
func (h hidden) Label(value string) hidden {
	return h.set("label", value)
}

// LabelAlign sets the label alignment.
func (h hidden) LabelAlign(value string) hidden {
	return h.set("labelAlign", value)
}

// LabelClassName sets the CSS class name for the label.
func (h hidden) LabelClassName(value string) hidden {
	return h.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (h hidden) LabelRemark(value string) hidden {
	return h.set("labelRemark", value)
}

// LabelWidth sets the custom width for the label.
func (h hidden) LabelWidth(value string) hidden {
	return h.set("labelWidth", value)
}

// Mode sets the display mode of the form item.
func (h hidden) Mode(value string) hidden {
	return h.set("mode", value)
}

// Name sets the field name.
func (h hidden) Name(value string) hidden {
	return h.set("name", value)
}

// OnEvent sets the event action configuration.
func (h hidden) OnEvent(value any) hidden {
	return h.set("onEvent", value)
}

// Placeholder sets the placeholder text.
func (h hidden) Placeholder(value string) hidden {
	return h.set("placeholder", value)
}

// ReadOnly sets whether the field is read-only.
func (h hidden) ReadOnly(value bool) hidden {
	return h.set("readOnly", value)
}

// ReadOnlyOn sets the expression to determine if the field is read-only.
func (h hidden) ReadOnlyOn(value string) hidden {
	return h.set("readOnlyOn", value)
}

// Remark sets the remark.
func (h hidden) Remark(value string) hidden {
	return h.set("remark", value)
}

// Required sets whether the field is required.
func (h hidden) Required(value bool) hidden {
	return h.set("required", value)
}

// Row sets the row value.
func (h hidden) Row(value string) hidden {
	return h.set("row", value)
}

// SaveImmediately sets whether to save immediately in TableColumn.
func (h hidden) SaveImmediately(value bool) hidden {
	return h.set("saveImmediately", value)
}

// Size sets the size of the form item.
func (h hidden) Size(value string) hidden {
	return h.set("size", value)
}

// Static sets whether the field is displayed statically.
func (h hidden) Static(value bool) hidden {
	return h.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (h hidden) StaticClassName(value string) hidden {
	return h.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display.
func (h hidden) StaticInputClassName(value string) hidden {
	return h.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display.
func (h hidden) StaticLabelClassName(value string) hidden {
	return h.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the field is displayed statically.
func (h hidden) StaticOn(value string) hidden {
	return h.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (h hidden) StaticPlaceholder(value string) hidden {
	return h.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (h hidden) StaticSchema(value string) hidden {
	return h.set("staticSchema", value)
}

// Style sets the component style.
func (h hidden) Style(value any) hidden {
	return h.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (h hidden) SubmitOnChange(value bool) hidden {
	return h.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (h hidden) TestIdBuilder(value string) hidden {
	return h.set("testIdBuilder", value)
}

// UseMobileUI sets whether to use mobile UI styles.
func (h hidden) UseMobileUI(value bool) hidden {
	return h.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API.
func (h hidden) ValidateApi(value string) hidden {
	return h.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (h hidden) ValidateOnChange(value bool) hidden {
	return h.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (h hidden) ValidationErrors(value string) hidden {
	return h.set("validationErrors", value)
}

// Validations sets the validation rules.
func (h hidden) Validations(value string) hidden {
	return h.set("validations", value)
}

// Value sets the default value.
func (h hidden) Value(value string) hidden {
	return h.set("value", value)
}

// Visible sets whether the field is visible.
func (h hidden) Visible(value bool) hidden {
	return h.set("visible", value)
}

// VisibleOn sets the expression to determine if the field is visible.
func (h hidden) VisibleOn(value string) hidden {
	return h.set("visibleOn", value)
}

// Width sets the width in a table.
func (h hidden) Width(value string) hidden {
	return h.set("width", value)
}
