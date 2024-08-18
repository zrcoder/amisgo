package comp

// FormulaControl 公式功能控件
type FormulaControl Schema

// NewFormulaControl 创建一个新的 FormulaControl 实例
func NewFormulaControl() FormulaControl {
	return make(FormulaControl).set("type", "formula")
}

func (fc FormulaControl) set(key string, value interface{}) FormulaControl {
	fc[key] = value
	return fc
}

// AutoFill 自动填充
func (fc FormulaControl) AutoFill(value string) FormulaControl {
	return fc.set("autoFill", value)
}

// AutoSet 是否自动应用
func (fc FormulaControl) AutoSet(value bool) FormulaControl {
	return fc.set("autoSet", value)
}

// ClassName 容器 css 类名
func (fc FormulaControl) ClassName(value string) FormulaControl {
	return fc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时是否删除值
func (fc FormulaControl) ClearValueOnHidden(value bool) FormulaControl {
	return fc.set("clearValueOnHidden", value)
}

// Condition 触发公式的作用条件
func (fc FormulaControl) Condition(value string) FormulaControl {
	return fc.set("condition", value)
}

// Desc 描述内容
func (fc FormulaControl) Desc(value string) FormulaControl {
	return fc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (fc FormulaControl) Description(value string) FormulaControl {
	return fc.set("description", value)
}

// DescriptionClassName 描述上的 className
func (fc FormulaControl) DescriptionClassName(value string) FormulaControl {
	return fc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (fc FormulaControl) Disabled(value bool) FormulaControl {
	return fc.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (fc FormulaControl) DisabledOn(value string) FormulaControl {
	return fc.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (fc FormulaControl) EditorSetting(value string) FormulaControl {
	return fc.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (fc FormulaControl) ExtraName(value string) FormulaControl {
	return fc.set("extraName", value)
}

// Formula 公式
func (fc FormulaControl) Formula(value string) FormulaControl {
	return fc.set("formula", value)
}

// Hidden 是否隐藏
func (fc FormulaControl) Hidden(value bool) FormulaControl {
	return fc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (fc FormulaControl) HiddenOn(value string) FormulaControl {
	return fc.set("hiddenOn", value)
}

// Hint 输入提示
func (fc FormulaControl) Hint(value string) FormulaControl {
	return fc.set("hint", value)
}

// Horizontal 水平布局配置
func (fc FormulaControl) Horizontal(value string) FormulaControl {
	return fc.set("horizontal", value)
}

// ID 触发公式应用的按钮目标
func (fc FormulaControl) ID(value string) FormulaControl {
	return fc.set("id", value)
}

// InitAutoFill 初始化自动填充
func (fc FormulaControl) InitAutoFill(value string) FormulaControl {
	return fc.set("initAutoFill", value)
}

// InitSet 是否初始应用
func (fc FormulaControl) InitSet(value bool) FormulaControl {
	return fc.set("initSet", value)
}

// Inline 表单 control 是否为 inline 模式
func (fc FormulaControl) Inline(value bool) FormulaControl {
	return fc.set("inline", value)
}

// InputClassName 配置 input className
func (fc FormulaControl) InputClassName(value string) FormulaControl {
	return fc.set("inputClassName", value)
}

// Label 描述标题
func (fc FormulaControl) Label(value string) FormulaControl {
	return fc.set("label", value)
}

// LabelAlign 描述标题对齐
func (fc FormulaControl) LabelAlign(value string) FormulaControl {
	return fc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (fc FormulaControl) LabelClassName(value string) FormulaControl {
	return fc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标
func (fc FormulaControl) LabelRemark(value string) FormulaControl {
	return fc.set("labelRemark", value)
}

// LabelWidth label自定义宽度
func (fc FormulaControl) LabelWidth(value string) FormulaControl {
	return fc.set("labelWidth", value)
}

// Mode 组件展示模式
func (fc FormulaControl) Mode(value string) FormulaControl {
	return fc.set("mode", value)
}

// Name 字段名
func (fc FormulaControl) Name(value string) FormulaControl {
	return fc.set("name", value)
}

// OnEvent 事件动作配置
func (fc FormulaControl) OnEvent(value string) FormulaControl {
	return fc.set("onEvent", value)
}

// Placeholder 占位符
func (fc FormulaControl) Placeholder(value string) FormulaControl {
	return fc.set("placeholder", value)
}

// ReadOnly 是否只读
func (fc FormulaControl) ReadOnly(value bool) FormulaControl {
	return fc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (fc FormulaControl) ReadOnlyOn(value string) FormulaControl {
	return fc.set("readOnlyOn", value)
}

// Remark 显示一个小图标
func (fc FormulaControl) Remark(value string) FormulaControl {
	return fc.set("remark", value)
}

// Required 是否为必填
func (fc FormulaControl) Required(value bool) FormulaControl {
	return fc.set("required", value)
}

// Row 行配置
func (fc FormulaControl) Row(value string) FormulaControl {
	return fc.set("row", value)
}

// SaveImmediately 是否立即保存
func (fc FormulaControl) SaveImmediately(value bool) FormulaControl {
	return fc.set("saveImmediately", value)
}

// Size 表单项大小
func (fc FormulaControl) Size(value string) FormulaControl {
	return fc.set("size", value)
}

// Static 是否静态展示
func (fc FormulaControl) Static(value bool) FormulaControl {
	return fc.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (fc FormulaControl) StaticClassName(value string) FormulaControl {
	return fc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (fc FormulaControl) StaticInputClassName(value string) FormulaControl {
	return fc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (fc FormulaControl) StaticLabelClassName(value string) FormulaControl {
	return fc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (fc FormulaControl) StaticOn(value string) FormulaControl {
	return fc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (fc FormulaControl) StaticPlaceholder(value string) FormulaControl {
	return fc.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (fc FormulaControl) StaticSchema(value string) FormulaControl {
	return fc.set("staticSchema", value)
}

// Style 组件样式
func (fc FormulaControl) Style(value string) FormulaControl {
	return fc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (fc FormulaControl) SubmitOnChange(value bool) FormulaControl {
	return fc.set("submitOnChange", value)
}

// TestIdBuilder
func (fc FormulaControl) TestIdBuilder(value string) FormulaControl {
	return fc.set("testIdBuilder", value)
}

// UseMobileUI 关闭移动端样式
func (fc FormulaControl) UseMobileUI(value bool) FormulaControl {
	return fc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (fc FormulaControl) ValidateApi(value string) FormulaControl {
	return fc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (fc FormulaControl) ValidateOnChange(value bool) FormulaControl {
	return fc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (fc FormulaControl) ValidationErrors(value string) FormulaControl {
	return fc.set("validationErrors", value)
}

// Validations
func (fc FormulaControl) Validations(value string) FormulaControl {
	return fc.set("validations", value)
}

// Value 默认值
func (fc FormulaControl) Value(value string) FormulaControl {
	return fc.set("value", value)
}

// Visible 是否显示
func (fc FormulaControl) Visible(value bool) FormulaControl {
	return fc.set("visible", value)
}

// VisibleOn 是否显示表达式
func (fc FormulaControl) VisibleOn(value string) FormulaControl {
	return fc.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (fc FormulaControl) Width(value string) FormulaControl {
	return fc.set("width", value)
}
