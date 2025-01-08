package comp

// static Document: https://aisuda.bce.baidu.com/amis/zh-CN/components/form/static

type static Schema

func Static() static {
	return static{}.set("type", "static")
}

func (s static) set(key string, value any) static {
	s[key] = value
	return s
}

// AutoFill sets the auto-fill function
func (s static) AutoFill(value string) static {
	return s.set("autoFill", value)
}

// BorderMode sets the border mode, default is none
func (s static) BorderMode(value string) static {
	return s.set("borderMode", value)
}

// ClassName sets the container CSS class name
func (s static) ClassName(value string) static {
	return s.set("className", value)
}

// ClearValueOnHidden sets whether to clear the form item value when it is hidden
func (s static) ClearValueOnHidden(value bool) static {
	return s.set("clearValueOnHidden", value)
}

// Copyable sets the copy function
func (s static) Copyable(value any) static {
	return s.set("copyable", value)
}

// Desc sets the description
func (s static) Desc(value string) static {
	return s.set("desc", value)
}

// Description sets the description content, supports HTML fragments
func (s static) Description(value string) static {
	return s.set("description", value)
}

// DescriptionClassName sets the description class name
func (s static) DescriptionClassName(value string) static {
	return s.set("descriptionClassName", value)
}

// Disabled sets whether the form item is disabled
func (s static) Disabled(value bool) static {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to determine if the form item should be disabled
func (s static) DisabledOn(value string) static {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, can be ignored at runtime
func (s static) EditorSetting(value string) static {
	return s.set("editorSetting", value)
}

// ExtraName sets the extra field name, used when it is a range component
func (s static) ExtraName(value string) static {
	return s.set("extraName", value)
}

// Hidden sets whether the form item is hidden
func (s static) Hidden(value bool) static {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to determine if the form item should be hidden
func (s static) HiddenOn(value string) static {
	return s.set("hiddenOn", value)
}

// Hint sets the input hint
func (s static) Hint(value string) static {
	return s.set("hint", value)
}

// Horizontal sets the specific left-right allocation when in horizontal layout
func (s static) Horizontal(value string) static {
	return s.set("horizontal", value)
}

// Id sets the unique component ID, mainly used for log collection
func (s static) ID(value string) static {
	return s.set("id", value)
}

// InitAutoFill sets the initial auto-fill value
func (s static) InitAutoFill(value string) static {
	return s.set("initAutoFill", value)
}

// Inline sets whether the form control is in inline mode
func (s static) Inline(value bool) static {
	return s.set("inline", value)
}

// InputClassName sets the input class name
func (s static) InputClassName(value string) static {
	return s.set("inputClassName", value)
}

// Label sets the description title
func (s static) Label(value string) static {
	return s.set("label", value)
}

// LabelAlign sets the description title alignment
func (s static) LabelAlign(value string) static {
	return s.set("labelAlign", value)
}

// LabelClassName sets the label class name
func (s static) LabelClassName(value string) static {
	return s.set("labelClassName", value)
}

// LabelRemark sets the label remark
func (s static) LabelRemark(value string) static {
	return s.set("labelRemark", value)
}

// LabelWidth sets the label custom width, default unit is px
func (s static) LabelWidth(value string) static {
	return s.set("labelWidth", value)
}

// Mode sets the current form item display mode
func (s static) Mode(value string) static {
	return s.set("mode", value)
}

// Name sets the field name, key for form submission, supports multi-level, connected by .
func (s static) Name(value string) static {
	return s.set("name", value)
}

// OnEvent sets the event action configuration
func (s static) OnEvent(value any) static {
	return s.set("onEvent", value)
}

// Placeholder sets the placeholder
func (s static) Placeholder(value string) static {
	return s.set("placeholder", value)
}

// PopOver sets the view details function
func (s static) PopOver(value string) static {
	return s.set("popOver", value)
}

// QuickEdit sets the quick edit function
func (s static) QuickEdit(value string) static {
	return s.set("quickEdit", value)
}

// ReadOnly sets whether the form item is read-only
func (s static) ReadOnly(value bool) static {
	return s.set("readOnly", value)
}

// ReadOnlyOn sets the condition for read-only
func (s static) ReadOnlyOn(value string) static {
	return s.set("readOnlyOn", value)
}

// Remark sets the remark
func (s static) Remark(value string) static {
	return s.set("remark", value)
}

// Required sets whether the form item is required
func (s static) Required(value bool) static {
	return s.set("required", value)
}

// Row sets the row
func (s static) Row(value string) static {
	return s.set("row", value)
}

// SaveImmediately sets whether to save immediately (used in TableColumn)
func (s static) SaveImmediately(value bool) static {
	return s.set("saveImmediately", value)
}

// Size sets the form item size
func (s static) Size(value string) static {
	return s.set("size", value)
}

// Static sets whether to display in static mode
func (s static) Static(value bool) static {
	return s.set("static", value)
}

// StaticClassName sets the static display form item class name
func (s static) StaticClassName(value string) static {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (s static) StaticInputClassName(value string) static {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (s static) StaticLabelClassName(value string) static {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the form item should be displayed in static mode
func (s static) StaticOn(value string) static {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for empty values in static display
func (s static) StaticPlaceholder(value string) static {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (s static) StaticSchema(value string) static {
	return s.set("staticSchema", value)
}

// Style sets the component style
func (s static) Style(value any) static {
	return s.set("style", value)
}

// SubmitOnChange sets whether to submit the form immediately after modification
func (s static) SubmitOnChange(value bool) static {
	return s.set("submitOnChange", value)
}

// TestIdBuilder sets the test ID builder
func (s static) TestIdBuilder(value string) static {
	return s.set("testIdBuilder", value)
}

// Text sets the content template, does not support HTML
// More document: https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template
func (s static) Text(value string) static {
	return s.set("text", value)
}

// Tpl sets the content template, supports HTML
// More document: https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template
func (s static) Tpl(value string) static {
	return s.set("tpl", value)
}

// UseMobileUI sets whether to disable mobile UI styles at component level
func (s static) UseMobileUI(value bool) static {
	return s.set("useMobileUI", value)
}

// ValidateApi sets the remote validation form item interface
func (s static) ValidateApi(value string) static {
	return s.set("validateApi", value)
}

// ValidateOnChange sets whether to validate the form item immediately after modification
func (s static) ValidateOnChange(value bool) static {
	return s.set("validateOnChange", value)
}

// ValidationErrors sets the validation failure prompt
func (s static) ValidationErrors(value string) static {
	return s.set("validationErrors", value)
}

// Validations sets the validation rules
func (s static) Validations(value string) static {
	return s.set("validations", value)
}

// Value sets the default value, only static values are supported, does not support variables, data association is achieved by setting the name property
func (s static) Value(value string) static {
	return s.set("value", value)
}

// Visible sets whether the form item is visible
func (s static) Visible(value bool) static {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine if the form item should be visible
func (s static) VisibleOn(value string) static {
	return s.set("visibleOn", value)
}

// Width sets the width in Table
func (s static) Width(value string) static {
	return s.set("width", value)
}
