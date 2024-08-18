package comp

// diffEditor 编辑器
type diffEditor schema

func DiffEditor() diffEditor {
	return make(diffEditor).set("type", "diff-editor")
}

func (d diffEditor) set(key string, value interface{}) diffEditor {
	d[key] = value
	return d
}

// AutoFill 自动填充
func (d diffEditor) AutoFill(value string) diffEditor {
	return d.set("autoFill", value)
}

// ClassName 容器 css 类名
func (d diffEditor) ClassName(value string) diffEditor {
	return d.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时是否删除
func (d diffEditor) ClearValueOnHidden(value bool) diffEditor {
	return d.set("clearValueOnHidden", value)
}

// Desc 描述
func (d diffEditor) Desc(value string) diffEditor {
	return d.set("desc", value)
}

// Description 描述内容
func (d diffEditor) Description(value string) diffEditor {
	return d.set("description", value)
}

// DescriptionClassName 描述上的 className
func (d diffEditor) DescriptionClassName(value string) diffEditor {
	return d.set("descriptionClassName", value)
}

// DiffValue 左侧面板的值
func (d diffEditor) DiffValue(value string) diffEditor {
	return d.set("diffValue", value)
}

// Disabled 是否禁用
func (d diffEditor) Disabled(value bool) diffEditor {
	return d.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (d diffEditor) DisabledOn(value string) diffEditor {
	return d.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (d diffEditor) EditorSetting(value string) diffEditor {
	return d.set("editorSetting", value)
}

// ExtraName 额外字段名
func (d diffEditor) ExtraName(value string) diffEditor {
	return d.set("extraName", value)
}

// Hidden 是否隐藏
func (d diffEditor) Hidden(value bool) diffEditor {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (d diffEditor) HiddenOn(value string) diffEditor {
	return d.set("hiddenOn", value)
}

// Hint 输入提示
func (d diffEditor) Hint(value string) diffEditor {
	return d.set("hint", value)
}

// Horizontal 水平布局配置
func (d diffEditor) Horizontal(value string) diffEditor {
	return d.set("horizontal", value)
}

// ID 组件唯一 id
func (d diffEditor) ID(value string) diffEditor {
	return d.set("id", value)
}

// InitAutoFill 初始化自动填充
func (d diffEditor) InitAutoFill(value string) diffEditor {
	return d.set("initAutoFill", value)
}

// Inline 表单控制是否为 inline 模式
func (d diffEditor) Inline(value bool) diffEditor {
	return d.set("inline", value)
}

// InputClassName 配置 input className
func (d diffEditor) InputClassName(value string) diffEditor {
	return d.set("inputClassName", value)
}

// Label 描述标题
func (d diffEditor) Label(value string) diffEditor {
	return d.set("label", value)
}

// LabelAlign 描述标题对齐
func (d diffEditor) LabelAlign(value string) diffEditor {
	return d.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (d diffEditor) LabelClassName(value string) diffEditor {
	return d.set("labelClassName", value)
}

// LabelRemark 显示一个小图标，鼠标放上去的时候显示提示内容
func (d diffEditor) LabelRemark(value string) diffEditor {
	return d.set("labelRemark", value)
}

// LabelWidth label 自定义宽度
func (d diffEditor) LabelWidth(value string) diffEditor {
	return d.set("labelWidth", value)
}

// Language 语言，参考 monaco-editor
func (d diffEditor) Language(value string) diffEditor {
	return d.set("language", value)
}

// Mode 配置当前表单项展示模式
func (d diffEditor) Mode(value string) diffEditor {
	return d.set("mode", value)
}

// Name 字段名
func (d diffEditor) Name(value string) diffEditor {
	return d.set("name", value)
}

// OnEvent 事件动作配置
func (d diffEditor) OnEvent(value string) diffEditor {
	return d.set("onEvent", value)
}

// Options 编辑器配置
func (d diffEditor) Options(value string) diffEditor {
	return d.set("options", value)
}

// Placeholder 占位符
func (d diffEditor) Placeholder(value string) diffEditor {
	return d.set("placeholder", value)
}

// ReadOnly 是否只读
func (d diffEditor) ReadOnly(value bool) diffEditor {
	return d.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (d diffEditor) ReadOnlyOn(value string) diffEditor {
	return d.set("readOnlyOn", value)
}

// Remark 显示一个小图标，鼠标放上去的时候显示提示内容
func (d diffEditor) Remark(value string) diffEditor {
	return d.set("remark", value)
}

// Required 是否为必填
func (d diffEditor) Required(value bool) diffEditor {
	return d.set("required", value)
}

// Row 行配置
func (d diffEditor) Row(value string) diffEditor {
	return d.set("row", value)
}

// SaveImmediately 是否立即保存
func (d diffEditor) SaveImmediately(value bool) diffEditor {
	return d.set("saveImmediately", value)
}

// Size 表单项大小
func (d diffEditor) Size(value string) diffEditor {
	return d.set("size", value)
}

// Static 是否静态展示
func (d diffEditor) Static(value bool) diffEditor {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (d diffEditor) StaticClassName(value string) diffEditor {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (d diffEditor) StaticInputClassName(value string) diffEditor {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (d diffEditor) StaticLabelClassName(value string) diffEditor {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (d diffEditor) StaticOn(value string) diffEditor {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d diffEditor) StaticPlaceholder(value string) diffEditor {
	return d.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (d diffEditor) StaticSchema(value string) diffEditor {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d diffEditor) Style(value string) diffEditor {
	return d.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (d diffEditor) SubmitOnChange(value bool) diffEditor {
	return d.set("submitOnChange", value)
}

// TestIdBuilder 测试 ID 构建器
func (d diffEditor) TestIdBuilder(value string) diffEditor {
	return d.set("testIdBuilder", value)
}

// UseMobileUI 关闭移动端样式
func (d diffEditor) UseMobileUI(value bool) diffEditor {
	return d.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (d diffEditor) ValidateApi(value string) diffEditor {
	return d.set("validateApi", value)
}

// ValidateOnChange 不设置时，每次修改都会触发验证
func (d diffEditor) ValidateOnChange(value bool) diffEditor {
	return d.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (d diffEditor) ValidationErrors(value string) diffEditor {
	return d.set("validationErrors", value)
}

// Validations 验证配置
func (d diffEditor) Validations(value string) diffEditor {
	return d.set("validations", value)
}

// Value 默认值
func (d diffEditor) Value(value string) diffEditor {
	return d.set("value", value)
}

// Visible 是否显示
func (d diffEditor) Visible(value bool) diffEditor {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式
func (d diffEditor) VisibleOn(value string) diffEditor {
	return d.set("visibleOn", value)
}

// Width 在 Table 中调整宽度
func (d diffEditor) Width(value string) diffEditor {
	return d.set("width", value)
}
