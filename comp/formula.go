package comp

// formula 公式功能控件
type formula schema

// Formula 创建一个新的 FormulaControl 实例
func Formula() formula {
	return make(formula).set("type", "formula")
}

func (fc formula) set(key string, value interface{}) formula {
	fc[key] = value
	return fc
}

// AutoFill 自动填充
func (fc formula) AutoFill(value string) formula {
	return fc.set("autoFill", value)
}

// AutoSet 是否自动应用
func (fc formula) AutoSet(value bool) formula {
	return fc.set("autoSet", value)
}

// ClassName 容器 css 类名
func (fc formula) ClassName(value string) formula {
	return fc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时是否删除值
func (fc formula) ClearValueOnHidden(value bool) formula {
	return fc.set("clearValueOnHidden", value)
}

// Condition 触发公式的作用条件
func (fc formula) Condition(value string) formula {
	return fc.set("condition", value)
}

// Desc 描述内容
func (fc formula) Desc(value string) formula {
	return fc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (fc formula) Description(value string) formula {
	return fc.set("description", value)
}

// DescriptionClassName 描述上的 className
func (fc formula) DescriptionClassName(value string) formula {
	return fc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (fc formula) Disabled(value bool) formula {
	return fc.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (fc formula) DisabledOn(value string) formula {
	return fc.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (fc formula) EditorSetting(value string) formula {
	return fc.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (fc formula) ExtraName(value string) formula {
	return fc.set("extraName", value)
}

// Formula 公式
func (fc formula) Formula(value string) formula {
	return fc.set("formula", value)
}

// Hidden 是否隐藏
func (fc formula) Hidden(value bool) formula {
	return fc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (fc formula) HiddenOn(value string) formula {
	return fc.set("hiddenOn", value)
}

// Hint 输入提示
func (fc formula) Hint(value string) formula {
	return fc.set("hint", value)
}

// Horizontal 水平布局配置
func (fc formula) Horizontal(value string) formula {
	return fc.set("horizontal", value)
}

// ID 触发公式应用的按钮目标
func (fc formula) ID(value string) formula {
	return fc.set("id", value)
}

// InitAutoFill 初始化自动填充
func (fc formula) InitAutoFill(value string) formula {
	return fc.set("initAutoFill", value)
}

// InitSet 是否初始应用
func (fc formula) InitSet(value bool) formula {
	return fc.set("initSet", value)
}

// Inline 表单 control 是否为 inline 模式
func (fc formula) Inline(value bool) formula {
	return fc.set("inline", value)
}

// InputClassName 配置 input className
func (fc formula) InputClassName(value string) formula {
	return fc.set("inputClassName", value)
}

// Label 描述标题
func (fc formula) Label(value string) formula {
	return fc.set("label", value)
}

// LabelAlign 描述标题对齐
func (fc formula) LabelAlign(value string) formula {
	return fc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (fc formula) LabelClassName(value string) formula {
	return fc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标
func (fc formula) LabelRemark(value string) formula {
	return fc.set("labelRemark", value)
}

// LabelWidth label自定义宽度
func (fc formula) LabelWidth(value string) formula {
	return fc.set("labelWidth", value)
}

// Mode 组件展示模式
func (fc formula) Mode(value string) formula {
	return fc.set("mode", value)
}

// Name 字段名
func (fc formula) Name(value string) formula {
	return fc.set("name", value)
}

// OnEvent 事件动作配置
func (fc formula) OnEvent(value string) formula {
	return fc.set("onEvent", value)
}

// Placeholder 占位符
func (fc formula) Placeholder(value string) formula {
	return fc.set("placeholder", value)
}

// ReadOnly 是否只读
func (fc formula) ReadOnly(value bool) formula {
	return fc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (fc formula) ReadOnlyOn(value string) formula {
	return fc.set("readOnlyOn", value)
}

// Remark 显示一个小图标
func (fc formula) Remark(value string) formula {
	return fc.set("remark", value)
}

// Required 是否为必填
func (fc formula) Required(value bool) formula {
	return fc.set("required", value)
}

// Row 行配置
func (fc formula) Row(value string) formula {
	return fc.set("row", value)
}

// SaveImmediately 是否立即保存
func (fc formula) SaveImmediately(value bool) formula {
	return fc.set("saveImmediately", value)
}

// Size 表单项大小
func (fc formula) Size(value string) formula {
	return fc.set("size", value)
}

// Static 是否静态展示
func (fc formula) Static(value bool) formula {
	return fc.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (fc formula) StaticClassName(value string) formula {
	return fc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (fc formula) StaticInputClassName(value string) formula {
	return fc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (fc formula) StaticLabelClassName(value string) formula {
	return fc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (fc formula) StaticOn(value string) formula {
	return fc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (fc formula) StaticPlaceholder(value string) formula {
	return fc.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (fc formula) StaticSchema(value string) formula {
	return fc.set("staticSchema", value)
}

// Style 组件样式
func (fc formula) Style(value string) formula {
	return fc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (fc formula) SubmitOnChange(value bool) formula {
	return fc.set("submitOnChange", value)
}

// TestIdBuilder
func (fc formula) TestIdBuilder(value string) formula {
	return fc.set("testIdBuilder", value)
}

// UseMobileUI 关闭移动端样式
func (fc formula) UseMobileUI(value bool) formula {
	return fc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (fc formula) ValidateApi(value string) formula {
	return fc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (fc formula) ValidateOnChange(value bool) formula {
	return fc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (fc formula) ValidationErrors(value string) formula {
	return fc.set("validationErrors", value)
}

// Validations
func (fc formula) Validations(value string) formula {
	return fc.set("validations", value)
}

// Value 默认值
func (fc formula) Value(value string) formula {
	return fc.set("value", value)
}

// Visible 是否显示
func (fc formula) Visible(value bool) formula {
	return fc.set("visible", value)
}

// VisibleOn 是否显示表达式
func (fc formula) VisibleOn(value string) formula {
	return fc.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (fc formula) Width(value string) formula {
	return fc.set("width", value)
}
