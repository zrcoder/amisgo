package comp

// checkbox
type checkbox schema

// Checkbox 创建一个新的 Checkbox 实例
func Checkbox() checkbox {
	return make(checkbox).set("type", "checkbox")
}

func (c checkbox) set(key string, value interface{}) checkbox {
	c[key] = value
	return c
}

// AutoFill 设置 autoFill 属性
func (c checkbox) AutoFill(value string) checkbox {
	return c.set("autoFill", value)
}

// Badge 设置 badge 属性
func (c checkbox) Badge(value string) checkbox {
	return c.set("badge", value)
}

// Checked 设置 checked 属性
func (c checkbox) Checked(value bool) checkbox {
	return c.set("checked", value)
}

// ClassName 设置 className 属性
func (c checkbox) ClassName(value string) checkbox {
	return c.set("className", value)
}

// ClearValueOnHidden 设置 clearValueOnHidden 属性
func (c checkbox) ClearValueOnHidden(value bool) checkbox {
	return c.set("clearValueOnHidden", value)
}

// Desc 设置 desc 属性
func (c checkbox) Desc(value string) checkbox {
	return c.set("desc", value)
}

// Description 设置 description 属性
func (c checkbox) Description(value string) checkbox {
	return c.set("description", value)
}

// DescriptionClassName 设置 descriptionClassName 属性
func (c checkbox) DescriptionClassName(value string) checkbox {
	return c.set("descriptionClassName", value)
}

// Disabled 设置 disabled 属性
func (c checkbox) Disabled(value bool) checkbox {
	return c.set("disabled", value)
}

// DisabledOn 设置 disabledOn 属性
func (c checkbox) DisabledOn(value string) checkbox {
	return c.set("disabledOn", value)
}

// EditorSetting 设置 editorSetting 属性
func (c checkbox) EditorSetting(value string) checkbox {
	return c.set("editorSetting", value)
}

// ExtraName 设置 extraName 属性
func (c checkbox) ExtraName(value string) checkbox {
	return c.set("extraName", value)
}

// FalseValue 设置 falseValue 属性
func (c checkbox) FalseValue(value string) checkbox {
	return c.set("falseValue", value)
}

// Hidden 设置 hidden 属性
func (c checkbox) Hidden(value bool) checkbox {
	return c.set("hidden", value)
}

// HiddenOn 设置 hiddenOn 属性
func (c checkbox) HiddenOn(value string) checkbox {
	return c.set("hiddenOn", value)
}

// Hint 设置 hint 属性
func (c checkbox) Hint(value string) checkbox {
	return c.set("hint", value)
}

// Horizontal 设置 horizontal 属性
func (c checkbox) Horizontal(value string) checkbox {
	return c.set("horizontal", value)
}

// ID 设置 id 属性
func (c checkbox) ID(value string) checkbox {
	return c.set("id", value)
}

// InitAutoFill 设置 initAutoFill 属性
func (c checkbox) InitAutoFill(value string) checkbox {
	return c.set("initAutoFill", value)
}

// Inline 设置 inline 属性
func (c checkbox) Inline(value bool) checkbox {
	return c.set("inline", value)
}

// InputClassName 设置 inputClassName 属性
func (c checkbox) InputClassName(value string) checkbox {
	return c.set("inputClassName", value)
}

// Label 设置 label 属性
func (c checkbox) Label(value string) checkbox {
	return c.set("label", value)
}

// LabelAlign 设置 labelAlign 属性
func (c checkbox) LabelAlign(value string) checkbox {
	return c.set("labelAlign", value)
}

// LabelClassName 设置 labelClassName 属性
func (c checkbox) LabelClassName(value string) checkbox {
	return c.set("labelClassName", value)
}

// LabelRemark 设置 labelRemark 属性
func (c checkbox) LabelRemark(value string) checkbox {
	return c.set("labelRemark", value)
}

// LabelWidth 设置 labelWidth 属性
func (c checkbox) LabelWidth(value string) checkbox {
	return c.set("labelWidth", value)
}

// Mode 设置 mode 属性
func (c checkbox) Mode(value string) checkbox {
	return c.set("mode", value)
}

// Name 设置 name 属性
func (c checkbox) Name(value string) checkbox {
	return c.set("name", value)
}

// OnEvent 设置 onEvent 属性
func (c checkbox) OnEvent(value string) checkbox {
	return c.set("onEvent", value)
}

// Option 设置 option 属性
func (c checkbox) Option(value string) checkbox {
	return c.set("option", value)
}

// OptionType 设置 optionType 属性
func (c checkbox) OptionType(value string) checkbox {
	return c.set("optionType", value)
}

// Partial 设置 partial 属性
func (c checkbox) Partial(value bool) checkbox {
	return c.set("partial", value)
}

// Placeholder 设置 placeholder 属性
func (c checkbox) Placeholder(value string) checkbox {
	return c.set("placeholder", value)
}

// ReadOnly 设置 readOnly 属性
func (c checkbox) ReadOnly(value bool) checkbox {
	return c.set("readOnly", value)
}

// ReadOnlyOn 设置 readOnlyOn 属性
func (c checkbox) ReadOnlyOn(value string) checkbox {
	return c.set("readOnlyOn", value)
}

// Remark 设置 remark 属性
func (c checkbox) Remark(value string) checkbox {
	return c.set("remark", value)
}

// Required 设置 required 属性
func (c checkbox) Required(value bool) checkbox {
	return c.set("required", value)
}

// Row 设置 row 属性
func (c checkbox) Row(value string) checkbox {
	return c.set("row", value)
}

// SaveImmediately 设置 saveImmediately 属性
func (c checkbox) SaveImmediately(value bool) checkbox {
	return c.set("saveImmediately", value)
}

// Size 设置 size 属性
func (c checkbox) Size(value string) checkbox {
	return c.set("size", value)
}

// Static 设置 static 属性
func (c checkbox) Static(value bool) checkbox {
	return c.set("static", value)
}

// StaticClassName 设置 staticClassName 属性
func (c checkbox) StaticClassName(value string) checkbox {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置 staticInputClassName 属性
func (c checkbox) StaticInputClassName(value string) checkbox {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置 staticLabelClassName 属性
func (c checkbox) StaticLabelClassName(value string) checkbox {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置 staticOn 属性
func (c checkbox) StaticOn(value string) checkbox {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置 staticPlaceholder 属性
func (c checkbox) StaticPlaceholder(value string) checkbox {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置 staticSchema 属性
func (c checkbox) StaticSchema(value string) checkbox {
	return c.set("staticSchema", value)
}

// Style 设置 style 属性
func (c checkbox) Style(value string) checkbox {
	return c.set("style", value)
}

// SubmitOnChange 设置 submitOnChange 属性
func (c checkbox) SubmitOnChange(value bool) checkbox {
	return c.set("submitOnChange", value)
}

// TestIdBuilder 设置 testIdBuilder 属性
func (c checkbox) TestIdBuilder(value string) checkbox {
	return c.set("testIdBuilder", value)
}

// TrueValue 设置 trueValue 属性
func (c checkbox) TrueValue(value string) checkbox {
	return c.set("trueValue", value)
}

// UseMobileUI 设置 useMobileUI 属性
func (c checkbox) UseMobileUI(value bool) checkbox {
	return c.set("useMobileUI", value)
}

// ValidateApi 设置 validateApi 属性
func (c checkbox) ValidateApi(value string) checkbox {
	return c.set("validateApi", value)
}

// ValidateOnChange 设置 validateOnChange 属性
func (c checkbox) ValidateOnChange(value bool) checkbox {
	return c.set("validateOnChange", value)
}

// ValidationErrors 设置 validationErrors 属性
func (c checkbox) ValidationErrors(value string) checkbox {
	return c.set("validationErrors", value)
}

// Validations 设置 validations 属性
func (c checkbox) Validations(value string) checkbox {
	return c.set("validations", value)
}

// Value 设置 value 属性
func (c checkbox) Value(value string) checkbox {
	return c.set("value", value)
}

// Visible 设置 visible 属性
func (c checkbox) Visible(value bool) checkbox {
	return c.set("visible", value)
}

// VisibleOn 设置 visibleOn 属性
func (c checkbox) VisibleOn(value string) checkbox {
	return c.set("visibleOn", value)
}

// Width 设置 width 属性
func (c checkbox) Width(value string) checkbox {
	return c.set("width", value)
}
