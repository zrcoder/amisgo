package comp

// CheckboxControl 继承自 BaseRenderer
type CheckboxControl BaseRenderer

// NewCheckboxControl 创建一个新的 CheckboxControl 实例
func NewCheckboxControl() CheckboxControl {
	c := CheckboxControl(make(BaseRenderer))
	return c.set("type", "checkbox")
}

func (c CheckboxControl) set(key string, value interface{}) CheckboxControl {
	c[key] = value
	return c
}

// AutoFill 设置 autoFill 属性
func (c CheckboxControl) AutoFill(value string) CheckboxControl {
	return c.set("autoFill", value)
}

// Badge 设置 badge 属性
func (c CheckboxControl) Badge(value string) CheckboxControl {
	return c.set("badge", value)
}

// Checked 设置 checked 属性
func (c CheckboxControl) Checked(value bool) CheckboxControl {
	return c.set("checked", value)
}

// ClassName 设置 className 属性
func (c CheckboxControl) ClassName(value string) CheckboxControl {
	return c.set("className", value)
}

// ClearValueOnHidden 设置 clearValueOnHidden 属性
func (c CheckboxControl) ClearValueOnHidden(value bool) CheckboxControl {
	return c.set("clearValueOnHidden", value)
}

// Desc 设置 desc 属性
func (c CheckboxControl) Desc(value string) CheckboxControl {
	return c.set("desc", value)
}

// Description 设置 description 属性
func (c CheckboxControl) Description(value string) CheckboxControl {
	return c.set("description", value)
}

// DescriptionClassName 设置 descriptionClassName 属性
func (c CheckboxControl) DescriptionClassName(value string) CheckboxControl {
	return c.set("descriptionClassName", value)
}

// Disabled 设置 disabled 属性
func (c CheckboxControl) Disabled(value bool) CheckboxControl {
	return c.set("disabled", value)
}

// DisabledOn 设置 disabledOn 属性
func (c CheckboxControl) DisabledOn(value string) CheckboxControl {
	return c.set("disabledOn", value)
}

// EditorSetting 设置 editorSetting 属性
func (c CheckboxControl) EditorSetting(value string) CheckboxControl {
	return c.set("editorSetting", value)
}

// ExtraName 设置 extraName 属性
func (c CheckboxControl) ExtraName(value string) CheckboxControl {
	return c.set("extraName", value)
}

// FalseValue 设置 falseValue 属性
func (c CheckboxControl) FalseValue(value string) CheckboxControl {
	return c.set("falseValue", value)
}

// Hidden 设置 hidden 属性
func (c CheckboxControl) Hidden(value bool) CheckboxControl {
	return c.set("hidden", value)
}

// HiddenOn 设置 hiddenOn 属性
func (c CheckboxControl) HiddenOn(value string) CheckboxControl {
	return c.set("hiddenOn", value)
}

// Hint 设置 hint 属性
func (c CheckboxControl) Hint(value string) CheckboxControl {
	return c.set("hint", value)
}

// Horizontal 设置 horizontal 属性
func (c CheckboxControl) Horizontal(value string) CheckboxControl {
	return c.set("horizontal", value)
}

// ID 设置 id 属性
func (c CheckboxControl) ID(value string) CheckboxControl {
	return c.set("id", value)
}

// InitAutoFill 设置 initAutoFill 属性
func (c CheckboxControl) InitAutoFill(value string) CheckboxControl {
	return c.set("initAutoFill", value)
}

// Inline 设置 inline 属性
func (c CheckboxControl) Inline(value bool) CheckboxControl {
	return c.set("inline", value)
}

// InputClassName 设置 inputClassName 属性
func (c CheckboxControl) InputClassName(value string) CheckboxControl {
	return c.set("inputClassName", value)
}

// Label 设置 label 属性
func (c CheckboxControl) Label(value string) CheckboxControl {
	return c.set("label", value)
}

// LabelAlign 设置 labelAlign 属性
func (c CheckboxControl) LabelAlign(value string) CheckboxControl {
	return c.set("labelAlign", value)
}

// LabelClassName 设置 labelClassName 属性
func (c CheckboxControl) LabelClassName(value string) CheckboxControl {
	return c.set("labelClassName", value)
}

// LabelRemark 设置 labelRemark 属性
func (c CheckboxControl) LabelRemark(value string) CheckboxControl {
	return c.set("labelRemark", value)
}

// LabelWidth 设置 labelWidth 属性
func (c CheckboxControl) LabelWidth(value string) CheckboxControl {
	return c.set("labelWidth", value)
}

// Mode 设置 mode 属性
func (c CheckboxControl) Mode(value string) CheckboxControl {
	return c.set("mode", value)
}

// Name 设置 name 属性
func (c CheckboxControl) Name(value string) CheckboxControl {
	return c.set("name", value)
}

// OnEvent 设置 onEvent 属性
func (c CheckboxControl) OnEvent(value string) CheckboxControl {
	return c.set("onEvent", value)
}

// Option 设置 option 属性
func (c CheckboxControl) Option(value string) CheckboxControl {
	return c.set("option", value)
}

// OptionType 设置 optionType 属性
func (c CheckboxControl) OptionType(value string) CheckboxControl {
	return c.set("optionType", value)
}

// Partial 设置 partial 属性
func (c CheckboxControl) Partial(value bool) CheckboxControl {
	return c.set("partial", value)
}

// Placeholder 设置 placeholder 属性
func (c CheckboxControl) Placeholder(value string) CheckboxControl {
	return c.set("placeholder", value)
}

// ReadOnly 设置 readOnly 属性
func (c CheckboxControl) ReadOnly(value bool) CheckboxControl {
	return c.set("readOnly", value)
}

// ReadOnlyOn 设置 readOnlyOn 属性
func (c CheckboxControl) ReadOnlyOn(value string) CheckboxControl {
	return c.set("readOnlyOn", value)
}

// Remark 设置 remark 属性
func (c CheckboxControl) Remark(value string) CheckboxControl {
	return c.set("remark", value)
}

// Required 设置 required 属性
func (c CheckboxControl) Required(value bool) CheckboxControl {
	return c.set("required", value)
}

// Row 设置 row 属性
func (c CheckboxControl) Row(value string) CheckboxControl {
	return c.set("row", value)
}

// SaveImmediately 设置 saveImmediately 属性
func (c CheckboxControl) SaveImmediately(value bool) CheckboxControl {
	return c.set("saveImmediately", value)
}

// Size 设置 size 属性
func (c CheckboxControl) Size(value string) CheckboxControl {
	return c.set("size", value)
}

// Static 设置 static 属性
func (c CheckboxControl) Static(value bool) CheckboxControl {
	return c.set("static", value)
}

// StaticClassName 设置 staticClassName 属性
func (c CheckboxControl) StaticClassName(value string) CheckboxControl {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置 staticInputClassName 属性
func (c CheckboxControl) StaticInputClassName(value string) CheckboxControl {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置 staticLabelClassName 属性
func (c CheckboxControl) StaticLabelClassName(value string) CheckboxControl {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置 staticOn 属性
func (c CheckboxControl) StaticOn(value string) CheckboxControl {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置 staticPlaceholder 属性
func (c CheckboxControl) StaticPlaceholder(value string) CheckboxControl {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置 staticSchema 属性
func (c CheckboxControl) StaticSchema(value string) CheckboxControl {
	return c.set("staticSchema", value)
}

// Style 设置 style 属性
func (c CheckboxControl) Style(value string) CheckboxControl {
	return c.set("style", value)
}

// SubmitOnChange 设置 submitOnChange 属性
func (c CheckboxControl) SubmitOnChange(value bool) CheckboxControl {
	return c.set("submitOnChange", value)
}

// TestIdBuilder 设置 testIdBuilder 属性
func (c CheckboxControl) TestIdBuilder(value string) CheckboxControl {
	return c.set("testIdBuilder", value)
}

// TrueValue 设置 trueValue 属性
func (c CheckboxControl) TrueValue(value string) CheckboxControl {
	return c.set("trueValue", value)
}

// UseMobileUI 设置 useMobileUI 属性
func (c CheckboxControl) UseMobileUI(value bool) CheckboxControl {
	return c.set("useMobileUI", value)
}

// ValidateApi 设置 validateApi 属性
func (c CheckboxControl) ValidateApi(value string) CheckboxControl {
	return c.set("validateApi", value)
}

// ValidateOnChange 设置 validateOnChange 属性
func (c CheckboxControl) ValidateOnChange(value bool) CheckboxControl {
	return c.set("validateOnChange", value)
}

// ValidationErrors 设置 validationErrors 属性
func (c CheckboxControl) ValidationErrors(value string) CheckboxControl {
	return c.set("validationErrors", value)
}

// Validations 设置 validations 属性
func (c CheckboxControl) Validations(value string) CheckboxControl {
	return c.set("validations", value)
}

// Value 设置 value 属性
func (c CheckboxControl) Value(value string) CheckboxControl {
	return c.set("value", value)
}

// Visible 设置 visible 属性
func (c CheckboxControl) Visible(value bool) CheckboxControl {
	return c.set("visible", value)
}

// VisibleOn 设置 visibleOn 属性
func (c CheckboxControl) VisibleOn(value string) CheckboxControl {
	return c.set("visibleOn", value)
}

// Width 设置 width 属性
func (c CheckboxControl) Width(value string) CheckboxControl {
	return c.set("width", value)
}
