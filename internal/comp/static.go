package comp

import "github.com/zrcoder/amisgo/model"

// Static Document: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/Static
type Static model.Schema

func NewStatic() Static {
	return Static{"type": "static"}
}

func (s Static) set(key string, value any) Static {
	s[key] = value
	return s
}

// AutoFill sets the auto-fill function
func (s Static) AutoFill(value string) Static {
	return s.set("autoFill", value)
}

// BorderMode sets the border mode, default is none
func (s Static) BorderMode(value string) Static {
	return s.set("borderMode", value)
}

// ClassName sets the container CSS class name
func (s Static) ClassName(value string) Static {
	return s.set("className", value)
}

// ClearValueOnHidden sets whether to clear the form item value when it is hidden
func (s Static) ClearValueOnHidden(value bool) Static {
	return s.set("clearValueOnHidden", value)
}

// Copyable sets the copy function
func (s Static) Copyable(value any) Static {
	return s.set("copyable", value)
}

// Desc sets the description
func (s Static) Desc(value string) Static {
	return s.set("desc", value)
}

// Description sets the description content, supports HTML fragments
func (s Static) Description(value string) Static {
	return s.set("description", value)
}

// DescriptionClassName sets the description class name
func (s Static) DescriptionClassName(value string) Static {
	return s.set("descriptionClassName", value)
}

// Disabled sets whether the form item is disabled
func (s Static) Disabled(value bool) Static {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to determine if the form item should be disabled
func (s Static) DisabledOn(value string) Static {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, can be ignored at runtime
func (s Static) EditorSetting(value string) Static {
	return s.set("editorSetting", value)
}

// ExtraName sets the extra field name, used when it is a range component
func (s Static) ExtraName(value string) Static {
	return s.set("extraName", value)
}

// Hidden sets whether the form item is hidden
func (s Static) Hidden(value bool) Static {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to determine if the form item should be hidden
func (s Static) HiddenOn(value string) Static {
	return s.set("hiddenOn", value)
}

// Hint sets the input hint
func (s Static) Hint(value string) Static {
	return s.set("hint", value)
}

// Horizontal sets the specific left-right allocation when in horizontal layout
func (s Static) Horizontal(value string) Static {
	return s.set("horizontal", value)
}

// Id sets the unique component ID, mainly used for log collection
func (s Static) ID(value string) Static {
	return s.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (s Static) InitAutoFill(value string) Static {
	return s.set("initAutoFill", value)
}

// Inline sets whether the form control is in inline mode
func (s Static) Inline(value bool) Static {
	return s.set("inline", value)
}

// InputClassName sets the input class name
func (s Static) InputClassName(value string) Static {
	return s.set("inputClassName", value)
}

// Label sets the description title
func (s Static) Label(value string) Static {
	return s.set("label", value)
}

// LabelAlign sets the description title alignment
func (s Static) LabelAlign(value string) Static {
	return s.set("labelAlign", value)
}

// LabelClassName sets the label class name
func (s Static) LabelClassName(value string) Static {
	return s.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (s Static) LabelRemark(value string) Static {
	return s.set("labelRemark", value)
}

// LabelWidth sets the label custom width, default unit is px
func (s Static) LabelWidth(value string) Static {
	return s.set("labelWidth", value)
}

// Mode sets the current form item display mode
func (s Static) Mode(value string) Static {
	return s.set("mode", value)
}

// Name sets the field name, key for form submission, supports multi-level, connected by .
func (s Static) Name(value string) Static {
	return s.set("name", value)
}

// OnEvent sets the event action configuration
func (s Static) OnEvent(value any) Static {
	return s.set("onEvent", value)
}

// Placeholder sets the placeholder
func (s Static) Placeholder(value string) Static {
	return s.set("placeholder", value)
}

// PopOver sets the view details function
func (s Static) PopOver(value string) Static {
	return s.set("popOver", value)
}

// QuickEdit sets the quick edit function
func (s Static) QuickEdit(value string) Static {
	return s.set("quickEdit", value)
}

// ReadOnly sets whether the form item is read-only
func (s Static) ReadOnly(value bool) Static {
	return s.set("readOnly", value)
}

// ReadOnlyOn sets the condition for read-only
func (s Static) ReadOnlyOn(value string) Static {
	return s.set("readOnlyOn", value)
}

// Remark sets the remark
func (s Static) Remark(value string) Static {
	return s.set("remark", value)
}

// Required sets whether the form item is required
func (s Static) Required(value bool) Static {
	return s.set("required", value)
}

// Row sets the row
func (s Static) Row(value string) Static {
	return s.set("row", value)
}

// SaveImmediately sets whether to save immediately (used in TableColumn)
func (s Static) SaveImmediately(value bool) Static {
	return s.set("saveImmediately", value)
}

// Size sets the form item size
func (s Static) Size(value string) Static {
	return s.set("size", value)
}

// Static sets whether to display in static mode
func (s Static) Static(value bool) Static {
	return s.set("static", value)
}

// StaticClassName sets the static display form item class name
func (s Static) StaticClassName(value string) Static {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (s Static) StaticInputClassName(value string) Static {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (s Static) StaticLabelClassName(value string) Static {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the form item should be displayed in static mode
func (s Static) StaticOn(value string) Static {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for empty values in static display
func (s Static) StaticPlaceholder(value string) Static {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (s Static) StaticSchema(value string) Static {
	return s.set("staticSchema", value)
}

// Style sets the component style
func (s Static) Style(value any) Static {
	return s.set("style", value)
}

// SubmitOnChange sets whether to submit the form immediately after modification
func (s Static) SubmitOnChange(value bool) Static {
	return s.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (s Static) TestIdBuilder(value string) Static {
	return s.set("testIdBuilder", value)
}

// Text sets the content template, does not support HTML
// More document: https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template
func (s Static) Text(value string) Static {
	return s.set("text", value)
}

// Tpl sets the content template, supports HTML
// More document: https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template
func (s Static) Tpl(value string) Static {
	return s.set("tpl", value)
}

// UseMobileUI sets whether to disable mobile UI styles at component level
func (s Static) UseMobileUI(value bool) Static {
	return s.set("useMobileUI", value)
}

// ValidateApi sets the remote validation form item interface
func (s Static) ValidateApi(value string) Static {
	return s.set("validateApi", value)
}

// ValidateOnChange sets whether to validate the form item immediately after modification
func (s Static) ValidateOnChange(value bool) Static {
	return s.set("validateOnChange", value)
}

// ValidationErrors sets the validation failure prompt
func (s Static) ValidationErrors(value string) Static {
	return s.set("validationErrors", value)
}

// Validations sets the validation rules
func (s Static) Validations(value string) Static {
	return s.set("validations", value)
}

// Value sets the default value, only static values are supported, does not support variables, data association is achieved by setting the name property
func (s Static) Value(value string) Static {
	return s.set("value", value)
}

// Visible sets whether the form item is visible
func (s Static) Visible(value bool) Static {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine if the form item should be visible
func (s Static) VisibleOn(value string) Static {
	return s.set("visibleOn", value)
}

// Width sets the width in Table
func (s Static) Width(value string) Static {
	return s.set("width", value)
}
