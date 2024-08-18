package comp

// IconPickerControl 图标选择器
type IconPickerControl Schema

// NewIconPickerControl 创建一个新的 IconPickerControl 实例
func NewIconPickerControl() IconPickerControl {
	return make(IconPickerControl).set("type", "icon-picker")
}

func (i IconPickerControl) set(key string, value interface{}) IconPickerControl {
	i[key] = value
	return i
}

// AutoFill 自动填充
func (i IconPickerControl) AutoFill(value string) IconPickerControl {
	return i.set("autoFill", value)
}

// ClassName 容器 css 类名
func (i IconPickerControl) ClassName(value string) IconPickerControl {
	return i.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时是否清除值
func (i IconPickerControl) ClearValueOnHidden(value bool) IconPickerControl {
	return i.set("clearValueOnHidden", value)
}

// Desc 描述
func (i IconPickerControl) Desc(value string) IconPickerControl {
	return i.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (i IconPickerControl) Description(value string) IconPickerControl {
	return i.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (i IconPickerControl) DescriptionClassName(value string) IconPickerControl {
	return i.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (i IconPickerControl) Disabled(value bool) IconPickerControl {
	return i.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (i IconPickerControl) DisabledOn(value string) IconPickerControl {
	return i.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (i IconPickerControl) EditorSetting(value string) IconPickerControl {
	return i.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (i IconPickerControl) ExtraName(value string) IconPickerControl {
	return i.set("extraName", value)
}

// Hidden 是否隐藏
func (i IconPickerControl) Hidden(value bool) IconPickerControl {
	return i.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (i IconPickerControl) HiddenOn(value string) IconPickerControl {
	return i.set("hiddenOn", value)
}

// Hint 输入提示
func (i IconPickerControl) Hint(value string) IconPickerControl {
	return i.set("hint", value)
}

// Horizontal 配置水平布局的左右分配
func (i IconPickerControl) Horizontal(value string) IconPickerControl {
	return i.set("horizontal", value)
}

// ID 组件唯一 id
func (i IconPickerControl) ID(value string) IconPickerControl {
	return i.set("id", value)
}

// InitAutoFill 初始自动填充
func (i IconPickerControl) InitAutoFill(value string) IconPickerControl {
	return i.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式
func (i IconPickerControl) Inline(value bool) IconPickerControl {
	return i.set("inline", value)
}

// InputClassName 配置 input className
func (i IconPickerControl) InputClassName(value string) IconPickerControl {
	return i.set("inputClassName", value)
}

// Label 描述标题
func (i IconPickerControl) Label(value string) IconPickerControl {
	return i.set("label", value)
}

// LabelAlign 描述标题对齐
func (i IconPickerControl) LabelAlign(value string) IconPickerControl {
	return i.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (i IconPickerControl) LabelClassName(value string) IconPickerControl {
	return i.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (i IconPickerControl) LabelRemark(value string) IconPickerControl {
	return i.set("labelRemark", value)
}

// LabelWidth label 自定义宽度
func (i IconPickerControl) LabelWidth(value string) IconPickerControl {
	return i.set("labelWidth", value)
}

// Mode 当前表单项展示模式
func (i IconPickerControl) Mode(value string) IconPickerControl {
	return i.set("mode", value)
}

// Name 字段名
func (i IconPickerControl) Name(value string) IconPickerControl {
	return i.set("name", value)
}

// OnEvent 事件动作配置
func (i IconPickerControl) OnEvent(value string) IconPickerControl {
	return i.set("onEvent", value)
}

// Placeholder 占位符
func (i IconPickerControl) Placeholder(value string) IconPickerControl {
	return i.set("placeholder", value)
}

// ReadOnly 是否只读
func (i IconPickerControl) ReadOnly(value bool) IconPickerControl {
	return i.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (i IconPickerControl) ReadOnlyOn(value string) IconPickerControl {
	return i.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (i IconPickerControl) Remark(value string) IconPickerControl {
	return i.set("remark", value)
}

// Required 是否为必填
func (i IconPickerControl) Required(value bool) IconPickerControl {
	return i.set("required", value)
}

// Row 行
func (i IconPickerControl) Row(value string) IconPickerControl {
	return i.set("row", value)
}

// SaveImmediately 是否立即保存
func (i IconPickerControl) SaveImmediately(value bool) IconPickerControl {
	return i.set("saveImmediately", value)
}

// Size 表单项大小
func (i IconPickerControl) Size(value string) IconPickerControl {
	return i.set("size", value)
}

// Static 是否静态展示
func (i IconPickerControl) Static(value bool) IconPickerControl {
	return i.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (i IconPickerControl) StaticClassName(value string) IconPickerControl {
	return i.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (i IconPickerControl) StaticInputClassName(value string) IconPickerControl {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (i IconPickerControl) StaticLabelClassName(value string) IconPickerControl {
	return i.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (i IconPickerControl) StaticOn(value string) IconPickerControl {
	return i.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (i IconPickerControl) StaticPlaceholder(value string) IconPickerControl {
	return i.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (i IconPickerControl) StaticSchema(value string) IconPickerControl {
	return i.set("staticSchema", value)
}

// Style 组件样式
func (i IconPickerControl) Style(value string) IconPickerControl {
	return i.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (i IconPickerControl) SubmitOnChange(value bool) IconPickerControl {
	return i.set("submitOnChange", value)
}

// TestIdBuilder 测试 id builder
func (i IconPickerControl) TestIdBuilder(value string) IconPickerControl {
	return i.set("testIdBuilder", value)
}

// UseMobileUI 关闭移动端样式
func (i IconPickerControl) UseMobileUI(value bool) IconPickerControl {
	return i.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (i IconPickerControl) ValidateApi(value string) IconPickerControl {
	return i.set("validateApi", value)
}

// ValidateOnChange 是否每次修改都触发验证
func (i IconPickerControl) ValidateOnChange(value bool) IconPickerControl {
	return i.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (i IconPickerControl) ValidationErrors(value string) IconPickerControl {
	return i.set("validationErrors", value)
}

// Validations 验证规则
func (i IconPickerControl) Validations(value string) IconPickerControl {
	return i.set("validations", value)
}

// Value 默认值
func (i IconPickerControl) Value(value string) IconPickerControl {
	return i.set("value", value)
}

// Visible 是否显示
func (i IconPickerControl) Visible(value bool) IconPickerControl {
	return i.set("visible", value)
}

// VisibleOn 是否显示表达式
func (i IconPickerControl) VisibleOn(value string) IconPickerControl {
	return i.set("visibleOn", value)
}

// Width 在 Table 中调整宽度
func (i IconPickerControl) Width(value string) IconPickerControl {
	return i.set("width", value)
}
