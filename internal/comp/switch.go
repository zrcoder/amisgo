package comp

import "github.com/zrcoder/amisgo/schema"

// SwitchControl documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/switch
type SwitchControl schema.Schema

func NewSwitch() SwitchControl {
	return SwitchControl{"type": "switch"}
}

func (s SwitchControl) set(key string, value any) SwitchControl {
	s[key] = value
	return s
}

// AutoFill sets autoFill when an option is selected.
func (s SwitchControl) AutoFill(value string) SwitchControl {
	return s.set("autoFill", value)
}

// ClassName sets the container CSS class name.
func (s SwitchControl) ClassName(value string) SwitchControl {
	return s.set("className", value)
}

// ClearValueOnHidden removes the form item value when hidden.
func (s SwitchControl) ClearValueOnHidden(value bool) SwitchControl {
	return s.set("clearValueOnHidden", value)
}

// Desc sets the description.
func (s SwitchControl) Desc(value string) SwitchControl {
	return s.set("desc", value)
}

// Description sets the description content, supports HTML.
func (s SwitchControl) Description(value string) SwitchControl {
	return s.set("description", value)
}

// DescriptionClassName sets the CSS class name for the description.
func (s SwitchControl) DescriptionClassName(value string) SwitchControl {
	return s.set("descriptionClassName", value)
}

// Disabled sets whether the switch is disabled.
func (s SwitchControl) Disabled(value bool) SwitchControl {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to disable the switch.
func (s SwitchControl) DisabledOn(value string) SwitchControl {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (s SwitchControl) EditorSetting(value string) SwitchControl {
	return s.set("editorSetting", value)
}

// ExtraName sets an extra field name for range components.
func (s SwitchControl) ExtraName(value string) SwitchControl {
	return s.set("extraName", value)
}

// FalseValue sets the value when unchecked.
func (s SwitchControl) FalseValue(value string) SwitchControl {
	return s.set("falseValue", value)
}

// Hidden sets whether the switch is hidden.
func (s SwitchControl) Hidden(value bool) SwitchControl {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to hide the switch.
func (s SwitchControl) HiddenOn(value string) SwitchControl {
	return s.set("hiddenOn", value)
}

// Hint sets the input hint shown on focus.
func (s SwitchControl) Hint(value string) SwitchControl {
	return s.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (s SwitchControl) Horizontal(value string) SwitchControl {
	return s.set("horizontal", value)
}

// ID sets the unique component ID.
func (s SwitchControl) ID(value string) SwitchControl {
	return s.set("id", value)
}

// InitAutoFill sets the initial autoFill value.
func (s SwitchControl) InitAutoFill(value string) SwitchControl {
	return s.set("initAutoFill", value)
}

// Inline sets whether the form control is inline.
func (s SwitchControl) Inline(value bool) SwitchControl {
	return s.set("inline", value)
}

// InputClassName sets the input CSS class name.
func (s SwitchControl) InputClassName(value string) SwitchControl {
	return s.set("inputClassName", value)
}

// Label sets the label text.
func (s SwitchControl) Label(value string) SwitchControl {
	return s.set("label", value)
}

// LabelAlign sets the label alignment.
func (s SwitchControl) LabelAlign(value string) SwitchControl {
	return s.set("labelAlign", value)
}

// LabelClassName sets the label CSS class name.
func (s SwitchControl) LabelClassName(value string) SwitchControl {
	return s.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (s SwitchControl) LabelRemark(value string) SwitchControl {
	return s.set("labelRemark", value)
}

// LabelWidth sets the custom label width.
func (s SwitchControl) LabelWidth(value string) SwitchControl {
	return s.set("labelWidth", value)
}

// Loading sets whether the switch is in loading state.
func (s SwitchControl) Loading(value bool) SwitchControl {
	return s.set("loading", value)
}

// Mode sets the display mode of the form item.
func (s SwitchControl) Mode(value string) SwitchControl {
	return s.set("mode", value)
}

// Name sets the field name for form submission.
func (s SwitchControl) Name(value string) SwitchControl {
	return s.set("name", value)
}

// OffText sets the content displayed when off.
func (s SwitchControl) OffText(value ...any) SwitchControl {
	return s.set("offText", value)
}

// OnEvent sets the event action configuration.
func (s SwitchControl) OnEvent(value Event) SwitchControl {
	return s.set("onEvent", value)
}

// OnText sets the content displayed when on.
func (s SwitchControl) OnText(value ...any) SwitchControl {
	return s.set("onText", value)
}

// Option sets the option description.
func (s SwitchControl) Option(value string) SwitchControl {
	return s.set("option", value)
}

// Placeholder sets the placeholder text.
func (s SwitchControl) Placeholder(value string) SwitchControl {
	return s.set("placeholder", value)
}

// ReadOnly sets whether the switch is read-only.
func (s SwitchControl) ReadOnly(value bool) SwitchControl {
	return s.set("readOnly", value)
}

// ReadOnlyOn sets the read-only condition.
func (s SwitchControl) ReadOnlyOn(value string) SwitchControl {
	return s.set("readOnlyOn", value)
}

// Remark sets the remark.
func (s SwitchControl) Remark(value string) SwitchControl {
	return s.set("remark", value)
}

// Required sets whether the switch is required.
func (s SwitchControl) Required(value bool) SwitchControl {
	return s.set("required", value)
}

// Row sets the row value.
func (s SwitchControl) Row(value string) SwitchControl {
	return s.set("row", value)
}

// SaveImmediately sets whether to save immediately in TableColumn.
func (s SwitchControl) SaveImmediately(value bool) SwitchControl {
	return s.set("saveImmediately", value)
}

// Size sets the switch size.
func (s SwitchControl) Size(value string) SwitchControl {
	return s.set("size", value)
}

// Static sets whether the switch is displayed statically.
func (s SwitchControl) Static(value bool) SwitchControl {
	return s.set("static", value)
}

// StaticClassName sets the static display form item class name.
func (s SwitchControl) StaticClassName(value string) SwitchControl {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name.
func (s SwitchControl) StaticInputClassName(value string) SwitchControl {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name.
func (s SwitchControl) StaticLabelClassName(value string) SwitchControl {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression to display statically.
func (s SwitchControl) StaticOn(value string) SwitchControl {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder.
func (s SwitchControl) StaticPlaceholder(value string) SwitchControl {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema.
func (s SwitchControl) StaticSchema(value string) SwitchControl {
	return s.set("staticSchema", value)
}

// Style sets the component style.
func (s SwitchControl) Style(value any) SwitchControl {
	return s.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (s SwitchControl) SubmitOnChange(value bool) SwitchControl {
	return s.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (s SwitchControl) TestIdBuilder(value string) SwitchControl {
	return s.set("testIdBuilder", value)
}

// TrueValue sets the value when checked.
func (s SwitchControl) TrueValue(value string) SwitchControl {
	return s.set("trueValue", value)
}

// UseMobileUI sets whether to disable mobile UI styles.
func (s SwitchControl) UseMobileUI(value bool) SwitchControl {
	return s.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API.
func (s SwitchControl) ValidateApi(value string) SwitchControl {
	return s.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (s SwitchControl) ValidateOnChange(value bool) SwitchControl {
	return s.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (s SwitchControl) ValidationErrors(value string) SwitchControl {
	return s.set("validationErrors", value)
}

// Validations sets the validations.
func (s SwitchControl) Validations(value string) SwitchControl {
	return s.set("validations", value)
}

// Value sets the default value.
func (s SwitchControl) Value(value any) SwitchControl {
	return s.set("value", value)
}

// Visible sets whether the switch is visible.
func (s SwitchControl) Visible(value bool) SwitchControl {
	return s.set("visible", value)
}

// VisibleOn sets the expression to display the switch.
func (s SwitchControl) VisibleOn(value string) SwitchControl {
	return s.set("visibleOn", value)
}

// Width sets the width in Table.
func (s SwitchControl) Width(value string) SwitchControl {
	return s.set("width", value)
}
