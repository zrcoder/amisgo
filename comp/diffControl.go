package comp

// DiffControl 编辑器
type DiffControl Schema

// NewDiffControl 创建一个新的 DiffControl 实例，并设置默认的 type
func NewDiffControl() DiffControl {
	return make(DiffControl).set("type", "diff-editor")
}

func (d DiffControl) set(key string, value interface{}) DiffControl {
	d[key] = value
	return d
}

// AutoFill 自动填充
func (d DiffControl) AutoFill(value string) DiffControl {
	return d.set("autoFill", value)
}

// ClassName 容器 css 类名
func (d DiffControl) ClassName(value string) DiffControl {
	return d.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时是否删除
func (d DiffControl) ClearValueOnHidden(value bool) DiffControl {
	return d.set("clearValueOnHidden", value)
}

// Desc 描述
func (d DiffControl) Desc(value string) DiffControl {
	return d.set("desc", value)
}

// Description 描述内容
func (d DiffControl) Description(value string) DiffControl {
	return d.set("description", value)
}

// DescriptionClassName 描述上的 className
func (d DiffControl) DescriptionClassName(value string) DiffControl {
	return d.set("descriptionClassName", value)
}

// DiffValue 左侧面板的值
func (d DiffControl) DiffValue(value string) DiffControl {
	return d.set("diffValue", value)
}

// Disabled 是否禁用
func (d DiffControl) Disabled(value bool) DiffControl {
	return d.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (d DiffControl) DisabledOn(value string) DiffControl {
	return d.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (d DiffControl) EditorSetting(value string) DiffControl {
	return d.set("editorSetting", value)
}

// ExtraName 额外字段名
func (d DiffControl) ExtraName(value string) DiffControl {
	return d.set("extraName", value)
}

// Hidden 是否隐藏
func (d DiffControl) Hidden(value bool) DiffControl {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (d DiffControl) HiddenOn(value string) DiffControl {
	return d.set("hiddenOn", value)
}

// Hint 输入提示
func (d DiffControl) Hint(value string) DiffControl {
	return d.set("hint", value)
}

// Horizontal 水平布局配置
func (d DiffControl) Horizontal(value string) DiffControl {
	return d.set("horizontal", value)
}

// ID 组件唯一 id
func (d DiffControl) ID(value string) DiffControl {
	return d.set("id", value)
}

// InitAutoFill 初始化自动填充
func (d DiffControl) InitAutoFill(value string) DiffControl {
	return d.set("initAutoFill", value)
}

// Inline 表单控制是否为 inline 模式
func (d DiffControl) Inline(value bool) DiffControl {
	return d.set("inline", value)
}

// InputClassName 配置 input className
func (d DiffControl) InputClassName(value string) DiffControl {
	return d.set("inputClassName", value)
}

// Label 描述标题
func (d DiffControl) Label(value string) DiffControl {
	return d.set("label", value)
}

// LabelAlign 描述标题对齐
func (d DiffControl) LabelAlign(value string) DiffControl {
	return d.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (d DiffControl) LabelClassName(value string) DiffControl {
	return d.set("labelClassName", value)
}

// LabelRemark 显示一个小图标，鼠标放上去的时候显示提示内容
func (d DiffControl) LabelRemark(value string) DiffControl {
	return d.set("labelRemark", value)
}

// LabelWidth label 自定义宽度
func (d DiffControl) LabelWidth(value string) DiffControl {
	return d.set("labelWidth", value)
}

// Language 语言，参考 monaco-editor
func (d DiffControl) Language(value string) DiffControl {
	return d.set("language", value)
}

// Mode 配置当前表单项展示模式
func (d DiffControl) Mode(value string) DiffControl {
	return d.set("mode", value)
}

// Name 字段名
func (d DiffControl) Name(value string) DiffControl {
	return d.set("name", value)
}

// OnEvent 事件动作配置
func (d DiffControl) OnEvent(value string) DiffControl {
	return d.set("onEvent", value)
}

// Options 编辑器配置
func (d DiffControl) Options(value string) DiffControl {
	return d.set("options", value)
}

// Placeholder 占位符
func (d DiffControl) Placeholder(value string) DiffControl {
	return d.set("placeholder", value)
}

// ReadOnly 是否只读
func (d DiffControl) ReadOnly(value bool) DiffControl {
	return d.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (d DiffControl) ReadOnlyOn(value string) DiffControl {
	return d.set("readOnlyOn", value)
}

// Remark 显示一个小图标，鼠标放上去的时候显示提示内容
func (d DiffControl) Remark(value string) DiffControl {
	return d.set("remark", value)
}

// Required 是否为必填
func (d DiffControl) Required(value bool) DiffControl {
	return d.set("required", value)
}

// Row 行配置
func (d DiffControl) Row(value string) DiffControl {
	return d.set("row", value)
}

// SaveImmediately 是否立即保存
func (d DiffControl) SaveImmediately(value bool) DiffControl {
	return d.set("saveImmediately", value)
}

// Size 表单项大小
func (d DiffControl) Size(value string) DiffControl {
	return d.set("size", value)
}

// Static 是否静态展示
func (d DiffControl) Static(value bool) DiffControl {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (d DiffControl) StaticClassName(value string) DiffControl {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (d DiffControl) StaticInputClassName(value string) DiffControl {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (d DiffControl) StaticLabelClassName(value string) DiffControl {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (d DiffControl) StaticOn(value string) DiffControl {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d DiffControl) StaticPlaceholder(value string) DiffControl {
	return d.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (d DiffControl) StaticSchema(value string) DiffControl {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d DiffControl) Style(value string) DiffControl {
	return d.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (d DiffControl) SubmitOnChange(value bool) DiffControl {
	return d.set("submitOnChange", value)
}

// TestIdBuilder 测试 ID 构建器
func (d DiffControl) TestIdBuilder(value string) DiffControl {
	return d.set("testIdBuilder", value)
}

// UseMobileUI 关闭移动端样式
func (d DiffControl) UseMobileUI(value bool) DiffControl {
	return d.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (d DiffControl) ValidateApi(value string) DiffControl {
	return d.set("validateApi", value)
}

// ValidateOnChange 不设置时，每次修改都会触发验证
func (d DiffControl) ValidateOnChange(value bool) DiffControl {
	return d.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (d DiffControl) ValidationErrors(value string) DiffControl {
	return d.set("validationErrors", value)
}

// Validations 验证配置
func (d DiffControl) Validations(value string) DiffControl {
	return d.set("validations", value)
}

// Value 默认值
func (d DiffControl) Value(value string) DiffControl {
	return d.set("value", value)
}

// Visible 是否显示
func (d DiffControl) Visible(value bool) DiffControl {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式
func (d DiffControl) VisibleOn(value string) DiffControl {
	return d.set("visibleOn", value)
}

// Width 在 Table 中调整宽度
func (d DiffControl) Width(value string) DiffControl {
	return d.set("width", value)
}
