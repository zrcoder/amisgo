package comp

// iconPicker 图标选择器
type iconPicker schema

// IconPicker 创建一个新的 IconPicker 实例
func IconPicker() iconPicker {
	return make(iconPicker).set("type", "icon-picker")
}

func (i iconPicker) set(key string, value interface{}) iconPicker {
	i[key] = value
	return i
}

// AutoFill 自动填充
func (i iconPicker) AutoFill(value string) iconPicker {
	return i.set("autoFill", value)
}

// ClassName 容器 css 类名
func (i iconPicker) ClassName(value string) iconPicker {
	return i.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时是否清除值
func (i iconPicker) ClearValueOnHidden(value bool) iconPicker {
	return i.set("clearValueOnHidden", value)
}

// Desc 描述
func (i iconPicker) Desc(value string) iconPicker {
	return i.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (i iconPicker) Description(value string) iconPicker {
	return i.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (i iconPicker) DescriptionClassName(value string) iconPicker {
	return i.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (i iconPicker) Disabled(value bool) iconPicker {
	return i.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (i iconPicker) DisabledOn(value string) iconPicker {
	return i.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (i iconPicker) EditorSetting(value string) iconPicker {
	return i.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (i iconPicker) ExtraName(value string) iconPicker {
	return i.set("extraName", value)
}

// Hidden 是否隐藏
func (i iconPicker) Hidden(value bool) iconPicker {
	return i.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (i iconPicker) HiddenOn(value string) iconPicker {
	return i.set("hiddenOn", value)
}

// Hint 输入提示
func (i iconPicker) Hint(value string) iconPicker {
	return i.set("hint", value)
}

// Horizontal 配置水平布局的左右分配
func (i iconPicker) Horizontal(value string) iconPicker {
	return i.set("horizontal", value)
}

// ID 组件唯一 id
func (i iconPicker) ID(value string) iconPicker {
	return i.set("id", value)
}

// InitAutoFill 初始自动填充
func (i iconPicker) InitAutoFill(value string) iconPicker {
	return i.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式
func (i iconPicker) Inline(value bool) iconPicker {
	return i.set("inline", value)
}

// InputClassName 配置 input className
func (i iconPicker) InputClassName(value string) iconPicker {
	return i.set("inputClassName", value)
}

// Label 描述标题
func (i iconPicker) Label(value string) iconPicker {
	return i.set("label", value)
}

// LabelAlign 描述标题对齐
func (i iconPicker) LabelAlign(value string) iconPicker {
	return i.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (i iconPicker) LabelClassName(value string) iconPicker {
	return i.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (i iconPicker) LabelRemark(value string) iconPicker {
	return i.set("labelRemark", value)
}

// LabelWidth label 自定义宽度
func (i iconPicker) LabelWidth(value string) iconPicker {
	return i.set("labelWidth", value)
}

// Mode 当前表单项展示模式
func (i iconPicker) Mode(value string) iconPicker {
	return i.set("mode", value)
}

// Name 字段名
func (i iconPicker) Name(value string) iconPicker {
	return i.set("name", value)
}

// OnEvent 事件动作配置
func (i iconPicker) OnEvent(value string) iconPicker {
	return i.set("onEvent", value)
}

// Placeholder 占位符
func (i iconPicker) Placeholder(value string) iconPicker {
	return i.set("placeholder", value)
}

// ReadOnly 是否只读
func (i iconPicker) ReadOnly(value bool) iconPicker {
	return i.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (i iconPicker) ReadOnlyOn(value string) iconPicker {
	return i.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (i iconPicker) Remark(value string) iconPicker {
	return i.set("remark", value)
}

// Required 是否为必填
func (i iconPicker) Required(value bool) iconPicker {
	return i.set("required", value)
}

// Row 行
func (i iconPicker) Row(value string) iconPicker {
	return i.set("row", value)
}

// SaveImmediately 是否立即保存
func (i iconPicker) SaveImmediately(value bool) iconPicker {
	return i.set("saveImmediately", value)
}

// Size 表单项大小
func (i iconPicker) Size(value string) iconPicker {
	return i.set("size", value)
}

// Static 是否静态展示
func (i iconPicker) Static(value bool) iconPicker {
	return i.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (i iconPicker) StaticClassName(value string) iconPicker {
	return i.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (i iconPicker) StaticInputClassName(value string) iconPicker {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (i iconPicker) StaticLabelClassName(value string) iconPicker {
	return i.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (i iconPicker) StaticOn(value string) iconPicker {
	return i.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (i iconPicker) StaticPlaceholder(value string) iconPicker {
	return i.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (i iconPicker) StaticSchema(value string) iconPicker {
	return i.set("staticSchema", value)
}

// Style 组件样式
func (i iconPicker) Style(value string) iconPicker {
	return i.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (i iconPicker) SubmitOnChange(value bool) iconPicker {
	return i.set("submitOnChange", value)
}

// TestIdBuilder 测试 id builder
func (i iconPicker) TestIdBuilder(value string) iconPicker {
	return i.set("testIdBuilder", value)
}

// UseMobileUI 关闭移动端样式
func (i iconPicker) UseMobileUI(value bool) iconPicker {
	return i.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (i iconPicker) ValidateApi(value string) iconPicker {
	return i.set("validateApi", value)
}

// ValidateOnChange 是否每次修改都触发验证
func (i iconPicker) ValidateOnChange(value bool) iconPicker {
	return i.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (i iconPicker) ValidationErrors(value string) iconPicker {
	return i.set("validationErrors", value)
}

// Validations 验证规则
func (i iconPicker) Validations(value string) iconPicker {
	return i.set("validations", value)
}

// Value 默认值
func (i iconPicker) Value(value string) iconPicker {
	return i.set("value", value)
}

// Visible 是否显示
func (i iconPicker) Visible(value bool) iconPicker {
	return i.set("visible", value)
}

// VisibleOn 是否显示表达式
func (i iconPicker) VisibleOn(value string) iconPicker {
	return i.set("visibleOn", value)
}

// Width 在 Table 中调整宽度
func (i iconPicker) Width(value string) iconPicker {
	return i.set("width", value)
}
