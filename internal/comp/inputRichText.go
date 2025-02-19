package comp

import "github.com/zrcoder/amisgo/schema"

// InputRichText Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-rich-text
type InputRichText schema.Schema

func NewInputRichText() InputRichText {
	return InputRichText{"type": "input-rich-text"}
}

func (rc InputRichText) set(key string, value any) InputRichText {
	rc[key] = value
	return rc
}

// AutoFill sets autoFill when an option is selected.
func (rc InputRichText) AutoFill(value string) InputRichText {
	return rc.set("autoFill", value)
}

// BorderMode sets border mode: full, half, or none.
func (rc InputRichText) BorderMode(value string) InputRichText {
	return rc.set("borderMode", value)
}

// ClassName sets the container CSS class name.
func (rc InputRichText) ClassName(value string) InputRichText {
	return rc.set("className", value)
}

// ClearValueOnHidden removes the value when the form item is hidden.
func (rc InputRichText) ClearValueOnHidden(value bool) InputRichText {
	return rc.set("clearValueOnHidden", value)
}

// Desc sets the description.
func (rc InputRichText) Desc(value string) InputRichText {
	return rc.set("desc", value)
}

// Description sets the description content, supports HTML.
func (rc InputRichText) Description(value string) InputRichText {
	return rc.set("description", value)
}

// DescriptionClassName sets the class name for the description.
func (rc InputRichText) DescriptionClassName(value string) InputRichText {
	return rc.set("descriptionClassName", value)
}

// Disabled sets whether the form item is disabled.
func (rc InputRichText) Disabled(value bool) InputRichText {
	return rc.set("disabled", value)
}

// DisabledOn sets the expression to disable the form item.
func (rc InputRichText) DisabledOn(value string) InputRichText {
	return rc.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (rc InputRichText) EditorSetting(value string) InputRichText {
	return rc.set("editorSetting", value)
}

// ExtraName sets an extra field name for range components.
func (rc InputRichText) ExtraName(value string) InputRichText {
	return rc.set("extraName", value)
}

// FileField sets the receiver's field name.
func (rc InputRichText) FileField(value string) InputRichText {
	return rc.set("fileField", value)
}

// Hidden sets whether the form item is hidden.
func (rc InputRichText) Hidden(value bool) InputRichText {
	return rc.set("hidden", value)
}

// HiddenOn sets the expression to hide the form item.
func (rc InputRichText) HiddenOn(value string) InputRichText {
	return rc.set("hiddenOn", value)
}

// Hint sets the input hint displayed on focus.
func (rc InputRichText) Hint(value string) InputRichText {
	return rc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (rc InputRichText) Horizontal(value string) InputRichText {
	return rc.set("horizontal", value)
}

// ID sets the unique component ID.
func (rc InputRichText) ID(value string) InputRichText {
	return rc.set("id", value)
}

// InitAutoFill sets the initial autoFill value.
func (rc InputRichText) InitAutoFill(value string) InputRichText {
	return rc.set("initAutoFill", value)
}

// Inline sets whether the form control is inline.
func (rc InputRichText) Inline(value bool) InputRichText {
	return rc.set("inline", value)
}

// InputClassName sets the input class name.
func (rc InputRichText) InputClassName(value string) InputRichText {
	return rc.set("inputClassName", value)
}

// Label sets the label text.
func (rc InputRichText) Label(value string) InputRichText {
	return rc.set("label", value)
}

// LabelAlign sets the label alignment: right, left, top, or inherit.
func (rc InputRichText) LabelAlign(value string) InputRichText {
	return rc.set("labelAlign", value)
}

// LabelClassName sets the label class name.
func (rc InputRichText) LabelClassName(value string) InputRichText {
	return rc.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (rc InputRichText) LabelRemark(value string) InputRichText {
	return rc.set("labelRemark", value)
}

// LabelWidth sets the custom label width, default unit is px.
func (rc InputRichText) LabelWidth(value string) InputRichText {
	return rc.set("labelWidth", value)
}

// Mode sets the display mode: normal, inline, or horizontal.
func (rc InputRichText) Mode(value string) InputRichText {
	return rc.set("mode", value)
}

// Name sets the field name for form submission.
func (rc InputRichText) Name(value string) InputRichText {
	return rc.set("name", value)
}

// OnEvent sets the event action configuration.
func (rc InputRichText) OnEvent(value any) InputRichText {
	return rc.set("onEvent", value)
}

// Options sets the tinymce or froala configuration.
func (rc InputRichText) Options(value schema.Schema) InputRichText {
	return rc.set("options", value)
}

// Placeholder sets the placeholder text.
func (rc InputRichText) Placeholder(value string) InputRichText {
	return rc.set("placeholder", value)
}

// ReadOnly sets whether the form item is read-only.
func (rc InputRichText) ReadOnly(value bool) InputRichText {
	return rc.set("readOnly", value)
}

// ReadOnlyOn sets the expression for read-only condition.
func (rc InputRichText) ReadOnlyOn(value string) InputRichText {
	return rc.set("readOnlyOn", value)
}

// Receiver sets the image save API.
func (rc InputRichText) Receiver(value string) InputRichText {
	return rc.set("receiver", value)
}

// Remark sets the remark with a small icon.
func (rc InputRichText) Remark(value string) InputRichText {
	return rc.set("remark", value)
}

// Required sets whether the form item is required.
func (rc InputRichText) Required(value bool) InputRichText {
	return rc.set("required", value)
}

// Row sets the row value.
func (rc InputRichText) Row(value string) InputRichText {
	return rc.set("row", value)
}

// SaveImmediately sets whether to save immediately in TableColumn.
func (rc InputRichText) SaveImmediately(value bool) InputRichText {
	return rc.set("saveImmediately", value)
}

// Size sets the form item size: md or lg.
func (rc InputRichText) Size(value string) InputRichText {
	return rc.set("size", value)
}

// Static sets whether to display statically.
func (rc InputRichText) Static(value bool) InputRichText {
	return rc.set("static", value)
}

// StaticClassName sets the static display class name.
func (rc InputRichText) StaticClassName(value string) InputRichText {
	return rc.set("staticClassName", value)
}

// StaticInputClassName sets the static input class name.
func (rc InputRichText) StaticInputClassName(value string) InputRichText {
	return rc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label class name.
func (rc InputRichText) StaticLabelClassName(value string) InputRichText {
	return rc.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display.
func (rc InputRichText) StaticOn(value string) InputRichText {
	return rc.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder.
func (rc InputRichText) StaticPlaceholder(value string) InputRichText {
	return rc.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema.
func (rc InputRichText) StaticSchema(value string) InputRichText {
	return rc.set("staticSchema", value)
}

// Style sets the component style.
func (rc InputRichText) Style(value any) InputRichText {
	return rc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (rc InputRichText) SubmitOnChange(value bool) InputRichText {
	return rc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (rc InputRichText) TestIdBuilder(value string) InputRichText {
	return rc.set("testIdBuilder", value)
}

// UseMobileUI sets whether to disable mobile UI styles.
func (rc InputRichText) UseMobileUI(value bool) InputRichText {
	return rc.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API.
func (rc InputRichText) ValidateApi(value string) InputRichText {
	return rc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (rc InputRichText) ValidateOnChange(value bool) InputRichText {
	return rc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (rc InputRichText) ValidationErrors(value string) InputRichText {
	return rc.set("validationErrors", value)
}

// Validations sets the validation rules.
func (rc InputRichText) Validations(value string) InputRichText {
	return rc.set("validations", value)
}

// Value sets the default value.
func (rc InputRichText) Value(value string) InputRichText {
	return rc.set("value", value)
}

// Vendor sets the editor type: froala or tinymce, default is tinymce.
func (rc InputRichText) Vendor(value string) InputRichText {
	return rc.set("vendor", value)
}

// VideoReceiver sets the video save API.
func (rc InputRichText) VideoReceiver(value string) InputRichText {
	return rc.set("videoReceiver", value)
}

// Visible sets whether the form item is visible.
func (rc InputRichText) Visible(value bool) InputRichText {
	return rc.set("visible", value)
}

// VisibleOn sets the expression for visibility.
func (rc InputRichText) VisibleOn(value string) InputRichText {
	return rc.set("visibleOn", value)
}

// Width sets the width in Table.
func (rc InputRichText) Width(value string) InputRichText {
	return rc.set("width", value)
}
