package comp

// conditionBuilder 代表一个条件组合控件，支持多种配置项
type conditionBuilder schema

// ConditionBuilder 创建一个新的 ConditionBuilder 实例，并设置默认类型
func ConditionBuilder() conditionBuilder {
	return make(conditionBuilder).set("type", "condition-builder")
}

func (c conditionBuilder) set(key string, value any) conditionBuilder {
	c[key] = value
	return c
}

// AddBtnVisibleOn 设置“添加按钮”显示条件
func (c conditionBuilder) AddBtnVisibleOn(value string) conditionBuilder {
	return c.set("addBtnVisibleOn", value)
}

// AddGroupBtnVisibleOn 设置“添加条件组”按钮显示条件
func (c conditionBuilder) AddGroupBtnVisibleOn(value string) conditionBuilder {
	return c.set("addGroupBtnVisibleOn", value)
}

// AutoFill 设置自动填充选项
func (c conditionBuilder) AutoFill(value string) conditionBuilder {
	return c.set("autoFill", value)
}

// BuilderMode 设置展现模式
func (c conditionBuilder) BuilderMode(value string) conditionBuilder {
	return c.set("builderMode", value)
}

// ClassName 设置容器 CSS 类名
func (c conditionBuilder) ClassName(value string) conditionBuilder {
	return c.set("className", value)
}

// ClearValueOnHidden 设置表单项隐藏时是否删除值
func (c conditionBuilder) ClearValueOnHidden(value bool) conditionBuilder {
	return c.set("clearValueOnHidden", value)
}

// Config 设置其他配置
func (c conditionBuilder) Config(value string) conditionBuilder {
	return c.set("config", value)
}

// Desc 设置描述
func (c conditionBuilder) Desc(value string) conditionBuilder {
	return c.set("desc", value)
}

// Description 设置描述内容，支持 HTML 片段
func (c conditionBuilder) Description(value string) conditionBuilder {
	return c.set("description", value)
}

// DescriptionClassName 设置描述内容的 CSS 类名
func (c conditionBuilder) DescriptionClassName(value string) conditionBuilder {
	return c.set("descriptionClassName", value)
}

// Disabled 设置是否禁用
func (c conditionBuilder) Disabled(value bool) conditionBuilder {
	return c.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (c conditionBuilder) DisabledOn(value string) conditionBuilder {
	return c.set("disabledOn", value)
}

// Draggable 设置是否可拖拽
func (c conditionBuilder) Draggable(value bool) conditionBuilder {
	return c.set("draggable", value)
}

// EditorSetting 设置编辑器配置
func (c conditionBuilder) EditorSetting(value string) conditionBuilder {
	return c.set("editorSetting", value)
}

// Embed 设置内嵌模式
func (c conditionBuilder) Embed(value bool) conditionBuilder {
	return c.set("embed", value)
}

// ExtraName 设置额外字段名
func (c conditionBuilder) ExtraName(value string) conditionBuilder {
	return c.set("extraName", value)
}

// Fields 设置字段集合
func (c conditionBuilder) Fields(value string) conditionBuilder {
	return c.set("fields", value)
}

// Formula 设置字段输入控件为公式编辑器
func (c conditionBuilder) Formula(value string) conditionBuilder {
	return c.set("formula", value)
}

// FormulaForIf 设置 if 里面公式编辑器配置
func (c conditionBuilder) FormulaForIf(value string) conditionBuilder {
	return c.set("formulaForIf", value)
}

// Funcs 设置函数集合
func (c conditionBuilder) Funcs(value string) conditionBuilder {
	return c.set("funcs", value)
}

// Hidden 设置是否隐藏
func (c conditionBuilder) Hidden(value bool) conditionBuilder {
	return c.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (c conditionBuilder) HiddenOn(value string) conditionBuilder {
	return c.set("hiddenOn", value)
}

// Hint 设置输入提示
func (c conditionBuilder) Hint(value string) conditionBuilder {
	return c.set("hint", value)
}

// Horizontal 设置水平布局配置
func (c conditionBuilder) Horizontal(value string) conditionBuilder {
	return c.set("horizontal", value)
}

// ID 设置组件唯一 ID
func (c conditionBuilder) ID(value string) conditionBuilder {
	return c.set("id", value)
}

// InitAutoFill 设置初始化自动填充
func (c conditionBuilder) InitAutoFill(value string) conditionBuilder {
	return c.set("initAutoFill", value)
}

// Inline 设置表单控制是否为 inline 模式
func (c conditionBuilder) Inline(value bool) conditionBuilder {
	return c.set("inline", value)
}

// InputClassName 设置输入的 CSS 类名
func (c conditionBuilder) InputClassName(value string) conditionBuilder {
	return c.set("inputClassName", value)
}

// Label 设置描述标题
func (c conditionBuilder) Label(value string) conditionBuilder {
	return c.set("label", value)
}

// LabelAlign 设置描述标题对齐
func (c conditionBuilder) LabelAlign(value string) conditionBuilder {
	return c.set("labelAlign", value)
}

// LabelClassName 设置描述标题的 CSS 类名
func (c conditionBuilder) LabelClassName(value string) conditionBuilder {
	return c.set("labelClassName", value)
}

// LabelRemark 设置描述标题旁的小图标备注
func (c conditionBuilder) LabelRemark(value string) conditionBuilder {
	return c.set("labelRemark", value)
}

// LabelWidth 设置描述标题的宽度
func (c conditionBuilder) LabelWidth(value string) conditionBuilder {
	return c.set("labelWidth", value)
}

// Mode 设置表单项展示模式
func (c conditionBuilder) Mode(value string) conditionBuilder {
	return c.set("mode", value)
}

// Name 设置字段名
func (c conditionBuilder) Name(value string) conditionBuilder {
	return c.set("name", value)
}

// OnEvent 设置事件动作配置
func (c conditionBuilder) OnEvent(value any) conditionBuilder {
	return c.set("onEvent", value)
}

// PickerIcon 设置弹窗触发图标
func (c conditionBuilder) PickerIcon(value string) conditionBuilder {
	return c.set("pickerIcon", value)
}

// Placeholder 设置占位符
func (c conditionBuilder) Placeholder(value string) conditionBuilder {
	return c.set("placeholder", value)
}

// ReadOnly 设置是否只读
func (c conditionBuilder) ReadOnly(value bool) conditionBuilder {
	return c.set("readOnly", value)
}

// ReadOnlyOn 设置只读条件
func (c conditionBuilder) ReadOnlyOn(value string) conditionBuilder {
	return c.set("readOnlyOn", value)
}

// Remark 设置提示内容
func (c conditionBuilder) Remark(value string) conditionBuilder {
	return c.set("remark", value)
}

// Required 设置是否必填
func (c conditionBuilder) Required(value bool) conditionBuilder {
	return c.set("required", value)
}

// Row 设置行数
func (c conditionBuilder) Row(value string) conditionBuilder {
	return c.set("row", value)
}

// SaveImmediately 设置是否立即保存
func (c conditionBuilder) SaveImmediately(value bool) conditionBuilder {
	return c.set("saveImmediately", value)
}

// ShowANDOR 设置是否显示并或切换键按钮
func (c conditionBuilder) ShowANDOR(value bool) conditionBuilder {
	return c.set("showANDOR", value)
}

// Size 设置表单项大小
func (c conditionBuilder) Size(value string) conditionBuilder {
	return c.set("size", value)
}

// Source 设置远程配置项
func (c conditionBuilder) Source(value string) conditionBuilder {
	return c.set("source", value)
}

// Static 设置是否静态展示
func (c conditionBuilder) Static(value bool) conditionBuilder {
	return c.set("static", value)
}

// StaticClassName 设置静态展示的 CSS 类名
func (c conditionBuilder) StaticClassName(value string) conditionBuilder {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示的输入类名
func (c conditionBuilder) StaticInputClassName(value string) conditionBuilder {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示的标签类名
func (c conditionBuilder) StaticLabelClassName(value string) conditionBuilder {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置静态展示的表达式
func (c conditionBuilder) StaticOn(value string) conditionBuilder {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (c conditionBuilder) StaticPlaceholder(value string) conditionBuilder {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示模式
func (c conditionBuilder) StaticSchema(value string) conditionBuilder {
	return c.set("staticSchema", value)
}

// Style 设置组件样式
func (c conditionBuilder) Style(value any) conditionBuilder {
	return c.set("style", value)
}

// SubmitOnChange 设置修改时是否提交表单
func (c conditionBuilder) SubmitOnChange(value bool) conditionBuilder {
	return c.set("submitOnChange", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (c conditionBuilder) TestIdBuilder(value string) conditionBuilder {
	return c.set("testIdBuilder", value)
}

// UseMobileUI 设置是否使用移动端样式
func (c conditionBuilder) UseMobileUI(value bool) conditionBuilder {
	return c.set("useMobileUI", value)
}

// ValidateApi 设置远端校验表单项接口
func (c conditionBuilder) ValidateApi(value string) conditionBuilder {
	return c.set("validateApi", value)
}

// ValidateOnChange 设置是否在每次修改时触发验证
func (c conditionBuilder) ValidateOnChange(value bool) conditionBuilder {
	return c.set("validateOnChange", value)
}

// ValidationErrors 设置验证失败提示信息
func (c conditionBuilder) ValidationErrors(value string) conditionBuilder {
	return c.set("validationErrors", value)
}

// Validations 设置验证规则
func (c conditionBuilder) Validations(value string) conditionBuilder {
	return c.set("validations", value)
}

// Value 设置默认值
func (c conditionBuilder) Value(value string) conditionBuilder {
	return c.set("value", value)
}

// Visible 设置是否显示
func (c conditionBuilder) Visible(value bool) conditionBuilder {
	return c.set("visible", value)
}

// VisibleOn 设置显示表达式
func (c conditionBuilder) VisibleOn(value string) conditionBuilder {
	return c.set("visibleOn", value)
}

// Width 设置宽度
func (c conditionBuilder) Width(value string) conditionBuilder {
	return c.set("width", value)
}
