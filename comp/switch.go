package comp

// switchControl documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/switch

type switchControl Schema

// Switch creates a new SwitchControl instance
func Switch() switchControl {
	return switchControl{}.set("type", "switch")
}

func (s switchControl) set(key string, value any) switchControl {
	s[key] = value
	return s
}

// AutoFill sets autoFill when an option is selected.
func (s switchControl) AutoFill(value string) switchControl {
	return s.set("autoFill", value)
}

// ClassName sets the container CSS class name.
func (s switchControl) ClassName(value string) switchControl {
	return s.set("className", value)
}

// ClearValueOnHidden removes the form item value when hidden.
func (s switchControl) ClearValueOnHidden(value bool) switchControl {
	return s.set("clearValueOnHidden", value)
}

// Desc sets the description.
func (s switchControl) Desc(value string) switchControl {
	return s.set("desc", value)
}

// Description sets the description content, supports HTML.
func (s switchControl) Description(value string) switchControl {
	return s.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (s switchControl) DescriptionClassName(value string) switchControl {
	return s.set("descriptionClassName", value)
}

// Disabled sets whether the switch is disabled.
func (s switchControl) Disabled(value bool) switchControl {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to disable the switch.
func (s switchControl) DisabledOn(value string) switchControl {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (s switchControl) EditorSetting(value string) switchControl {
	return s.set("editorSetting", value)
}

// ExtraName sets an extra field name for range components.
func (s switchControl) ExtraName(value string) switchControl {
	return s.set("extraName", value)
}

// FalseValue sets the value when unchecked.
func (s switchControl) FalseValue(value string) switchControl {
	return s.set("falseValue", value)
}

// Hidden sets whether the switch is hidden.
func (s switchControl) Hidden(value bool) switchControl {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to hide the switch.
func (s switchControl) HiddenOn(value string) switchControl {
	return s.set("hiddenOn", value)
}

// Hint sets the input hint shown on focus.
func (s switchControl) Hint(value string) switchControl {
	return s.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (s switchControl) Horizontal(value string) switchControl {
	return s.set("horizontal", value)
}

// ID sets the unique component ID.
func (s switchControl) ID(value string) switchControl {
	return s.set("id", value)
}

// InitAutoFill sets the initial autoFill value.
func (s switchControl) InitAutoFill(value string) switchControl {
	return s.set("initAutoFill", value)
}

// Inline sets whether the form control is inline.
func (s switchControl) Inline(value bool) switchControl {
	return s.set("inline", value)
}

// InputClassName sets the input CSS class name.
func (s switchControl) InputClassName(value string) switchControl {
	return s.set("inputClassName", value)
}

// Label sets the label text.
func (s switchControl) Label(value string) switchControl {
	return s.set("label", value)
}

// LabelAlign sets the label alignment.
func (s switchControl) LabelAlign(value string) switchControl {
	return s.set("labelAlign", value)
}

// LabelClassName sets the label CSS class name.
func (s switchControl) LabelClassName(value string) switchControl {
	return s.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (s switchControl) LabelRemark(value string) switchControl {
	return s.set("labelRemark", value)
}

// LabelWidth sets the custom label width.
func (s switchControl) LabelWidth(value string) switchControl {
	return s.set("labelWidth", value)
}

// Loading sets whether the switch is in loading state.
func (s switchControl) Loading(value bool) switchControl {
	return s.set("loading", value)
}

// Mode sets the display mode of the form item.
func (s switchControl) Mode(value string) switchControl {
	return s.set("mode", value)
}

// Name sets the field name for form submission.
func (s switchControl) Name(value string) switchControl {
	return s.set("name", value)
}

// OffText sets the content displayed when off.
func (s switchControl) OffText(value ...any) switchControl {
	return s.set("offText", value)
}

// OnEvent sets the event action configuration.
func (s switchControl) OnEvent(value any) switchControl {
	return s.set("onEvent", value)
}

// OnText sets the content displayed when on.
func (s switchControl) OnText(value ...any) switchControl {
	return s.set("onText", value)
}

// Option sets the option description.
func (s switchControl) Option(value string) switchControl {
	return s.set("option", value)
}

// Placeholder sets the placeholder text.
func (s switchControl) Placeholder(value string) switchControl {
	return s.set("placeholder", value)
}

// ReadOnly sets whether the switch is read-only.
func (s switchControl) ReadOnly(value bool) switchControl {
	return s.set("readOnly", value)
}

// ReadOnlyOn sets the read-only condition.
func (s switchControl) ReadOnlyOn(value string) switchControl {
	return s.set("readOnlyOn", value)
}

// Remark sets the remark.
func (s switchControl) Remark(value string) switchControl {
	return s.set("remark", value)
}

// Required sets whether the switch is required.
func (s switchControl) Required(value bool) switchControl {
	return s.set("required", value)
}

// Row sets the row value.
func (s switchControl) Row(value string) switchControl {
	return s.set("row", value)
}

// SaveImmediately sets whether to save immediately in TableColumn.
func (s switchControl) SaveImmediately(value bool) switchControl {
	return s.set("saveImmediately", value)
}

// Size sets the switch size.
func (s switchControl) Size(value string) switchControl {
	return s.set("size", value)
}

// Static sets whether the switch is displayed statically.
func (s switchControl) Static(value bool) switchControl {
	return s.set("static", value)
}

// StaticClassName sets the static display form item class name.
func (s switchControl) StaticClassName(value string) switchControl {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name.
func (s switchControl) StaticInputClassName(value string) switchControl {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name.
func (s switchControl) StaticLabelClassName(value string) switchControl {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression to display statically.
func (s switchControl) StaticOn(value string) switchControl {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder.
func (s switchControl) StaticPlaceholder(value string) switchControl {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (s switchControl) StaticSchema(value string) switchControl {
	return s.set("staticSchema", value)
}

// Style sets the component style.
func (s switchControl) Style(value any) switchControl {
	return s.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (s switchControl) SubmitOnChange(value bool) switchControl {
	return s.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (s switchControl) TestIdBuilder(value string) switchControl {
	return s.set("testIdBuilder", value)
}

// TrueValue sets the value when checked.
func (s switchControl) TrueValue(value string) switchControl {
	return s.set("trueValue", value)
}

// UseMobileUI sets whether to disable mobile UI styles.
func (s switchControl) UseMobileUI(value bool) switchControl {
	return s.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API.
func (s switchControl) ValidateApi(value string) switchControl {
	return s.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (s switchControl) ValidateOnChange(value bool) switchControl {
	return s.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (s switchControl) ValidationErrors(value string) switchControl {
	return s.set("validationErrors", value)
}

// Validations sets the validations.
func (s switchControl) Validations(value string) switchControl {
	return s.set("validations", value)
}

// Value sets the default value.
func (s switchControl) Value(value any) switchControl {
	return s.set("value", value)
}

// Visible sets whether the switch is visible.
func (s switchControl) Visible(value bool) switchControl {
	return s.set("visible", value)
}

// VisibleOn sets the expression to display the switch.
func (s switchControl) VisibleOn(value string) switchControl {
	return s.set("visibleOn", value)
}

// Width sets the width in Table.
func (s switchControl) Width(value string) switchControl {
	return s.set("width", value)
}
