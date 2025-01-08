package comp

import "github.com/zrcoder/amisgo/model"

// inputRichText Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-rich-text

type inputRichText model.Schema

func InputRichText() inputRichText {
	return inputRichText{}.set("type", "input-rich-text")
}

func (rc inputRichText) set(key string, value any) inputRichText {
	rc[key] = value
	return rc
}

// AutoFill sets autoFill when an option is selected.
func (rc inputRichText) AutoFill(value string) inputRichText {
	return rc.set("autoFill", value)
}

// BorderMode sets border mode: full, half, or none.
func (rc inputRichText) BorderMode(value string) inputRichText {
	return rc.set("borderMode", value)
}

// ClassName sets the container CSS class name.
func (rc inputRichText) ClassName(value string) inputRichText {
	return rc.set("className", value)
}

// ClearValueOnHidden removes the value when the form item is hidden.
func (rc inputRichText) ClearValueOnHidden(value bool) inputRichText {
	return rc.set("clearValueOnHidden", value)
}

// Desc sets the description.
func (rc inputRichText) Desc(value string) inputRichText {
	return rc.set("desc", value)
}

// Description sets the description content, supports HTML.
func (rc inputRichText) Description(value string) inputRichText {
	return rc.set("description", value)
}

// DescriptionClassName sets the class name for the description.
func (rc inputRichText) DescriptionClassName(value string) inputRichText {
	return rc.set("descriptionClassName", value)
}

// Disabled sets whether the form item is disabled.
func (rc inputRichText) Disabled(value bool) inputRichText {
	return rc.set("disabled", value)
}

// DisabledOn sets the expression to disable the form item.
func (rc inputRichText) DisabledOn(value string) inputRichText {
	return rc.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (rc inputRichText) EditorSetting(value string) inputRichText {
	return rc.set("editorSetting", value)
}

// ExtraName sets an extra field name for range components.
func (rc inputRichText) ExtraName(value string) inputRichText {
	return rc.set("extraName", value)
}

// FileField sets the receiver's field name.
func (rc inputRichText) FileField(value string) inputRichText {
	return rc.set("fileField", value)
}

// Hidden sets whether the form item is hidden.
func (rc inputRichText) Hidden(value bool) inputRichText {
	return rc.set("hidden", value)
}

// HiddenOn sets the expression to hide the form item.
func (rc inputRichText) HiddenOn(value string) inputRichText {
	return rc.set("hiddenOn", value)
}

// Hint sets the input hint displayed on focus.
func (rc inputRichText) Hint(value string) inputRichText {
	return rc.set("hint", value)
}

// Horizontal sets the horizontal layout configuration.
func (rc inputRichText) Horizontal(value string) inputRichText {
	return rc.set("horizontal", value)
}

// ID sets the unique component ID.
func (rc inputRichText) ID(value string) inputRichText {
	return rc.set("id", value)
}

// InitAutoFill sets the initial autoFill value.
func (rc inputRichText) InitAutoFill(value string) inputRichText {
	return rc.set("initAutoFill", value)
}

// Inline sets whether the form control is inline.
func (rc inputRichText) Inline(value bool) inputRichText {
	return rc.set("inline", value)
}

// InputClassName sets the input class name.
func (rc inputRichText) InputClassName(value string) inputRichText {
	return rc.set("inputClassName", value)
}

// Label sets the label text.
func (rc inputRichText) Label(value string) inputRichText {
	return rc.set("label", value)
}

// LabelAlign sets the label alignment: right, left, top, or inherit.
func (rc inputRichText) LabelAlign(value string) inputRichText {
	return rc.set("labelAlign", value)
}

// LabelClassName sets the label class name.
func (rc inputRichText) LabelClassName(value string) inputRichText {
	return rc.set("labelClassName", value)
}

// LabelRemark sets the label remark.
func (rc inputRichText) LabelRemark(value string) inputRichText {
	return rc.set("labelRemark", value)
}

// LabelWidth sets the custom label width, default unit is px.
func (rc inputRichText) LabelWidth(value string) inputRichText {
	return rc.set("labelWidth", value)
}

// Mode sets the display mode: normal, inline, or horizontal.
func (rc inputRichText) Mode(value string) inputRichText {
	return rc.set("mode", value)
}

// Name sets the field name for form submission.
func (rc inputRichText) Name(value string) inputRichText {
	return rc.set("name", value)
}

// OnEvent sets the event action configuration.
func (rc inputRichText) OnEvent(value any) inputRichText {
	return rc.set("onEvent", value)
}

// Options sets the tinymce or froala configuration.
func (rc inputRichText) Options(value model.Schema) inputRichText {
	return rc.set("options", value)
}

// Placeholder sets the placeholder text.
func (rc inputRichText) Placeholder(value string) inputRichText {
	return rc.set("placeholder", value)
}

// ReadOnly sets whether the form item is read-only.
func (rc inputRichText) ReadOnly(value bool) inputRichText {
	return rc.set("readOnly", value)
}

// ReadOnlyOn sets the expression for read-only condition.
func (rc inputRichText) ReadOnlyOn(value string) inputRichText {
	return rc.set("readOnlyOn", value)
}

// Receiver sets the image save API.
func (rc inputRichText) Receiver(value string) inputRichText {
	return rc.set("receiver", value)
}

// Remark sets the remark with a small icon.
func (rc inputRichText) Remark(value string) inputRichText {
	return rc.set("remark", value)
}

// Required sets whether the form item is required.
func (rc inputRichText) Required(value bool) inputRichText {
	return rc.set("required", value)
}

// Row sets the row value.
func (rc inputRichText) Row(value string) inputRichText {
	return rc.set("row", value)
}

// SaveImmediately sets whether to save immediately in TableColumn.
func (rc inputRichText) SaveImmediately(value bool) inputRichText {
	return rc.set("saveImmediately", value)
}

// Size sets the form item size: md or lg.
func (rc inputRichText) Size(value string) inputRichText {
	return rc.set("size", value)
}

// Static sets whether to display statically.
func (rc inputRichText) Static(value bool) inputRichText {
	return rc.set("static", value)
}

// StaticClassName sets the static display class name.
func (rc inputRichText) StaticClassName(value string) inputRichText {
	return rc.set("staticClassName", value)
}

// StaticInputClassName sets the static input class name.
func (rc inputRichText) StaticInputClassName(value string) inputRichText {
	return rc.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label class name.
func (rc inputRichText) StaticLabelClassName(value string) inputRichText {
	return rc.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display.
func (rc inputRichText) StaticOn(value string) inputRichText {
	return rc.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder.
func (rc inputRichText) StaticPlaceholder(value string) inputRichText {
	return rc.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.
func (rc inputRichText) StaticSchema(value string) inputRichText {
	return rc.set("staticSchema", value)
}

// Style sets the component style.
func (rc inputRichText) Style(value any) inputRichText {
	return rc.set("style", value)
}

// SubmitOnChange sets whether to submit the form on change.
func (rc inputRichText) SubmitOnChange(value bool) inputRichText {
	return rc.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder.
func (rc inputRichText) TestIdBuilder(value string) inputRichText {
	return rc.set("testIdBuilder", value)
}

// UseMobileUI sets whether to disable mobile UI styles.
func (rc inputRichText) UseMobileUI(value bool) inputRichText {
	return rc.set("useMobileUI", value)
}

// ValidateApi sets the remote validation API.
func (rc inputRichText) ValidateApi(value string) inputRichText {
	return rc.set("validateApi", value)
}

// ValidateOnChange sets whether to validate on change.
func (rc inputRichText) ValidateOnChange(value bool) inputRichText {
	return rc.set("validateOnChange", value)
}

// ValidationErrors sets the validation error messages.
func (rc inputRichText) ValidationErrors(value string) inputRichText {
	return rc.set("validationErrors", value)
}

// Validations sets the validation rules.
func (rc inputRichText) Validations(value string) inputRichText {
	return rc.set("validations", value)
}

// Value sets the default value.
func (rc inputRichText) Value(value string) inputRichText {
	return rc.set("value", value)
}

// Vendor sets the editor type: froala or tinymce, default is tinymce.
func (rc inputRichText) Vendor(value string) inputRichText {
	return rc.set("vendor", value)
}

// VideoReceiver sets the video save API.
func (rc inputRichText) VideoReceiver(value string) inputRichText {
	return rc.set("videoReceiver", value)
}

// Visible sets whether the form item is visible.
func (rc inputRichText) Visible(value bool) inputRichText {
	return rc.set("visible", value)
}

// VisibleOn sets the expression for visibility.
func (rc inputRichText) VisibleOn(value string) inputRichText {
	return rc.set("visibleOn", value)
}

// Width sets the width in Table.
func (rc inputRichText) Width(value string) inputRichText {
	return rc.set("width", value)
}
