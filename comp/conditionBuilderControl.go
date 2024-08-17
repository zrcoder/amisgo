package comp

// ConditionBuilderControl 代表一个条件组合控件，支持多种配置项
type ConditionBuilderControl BaseRenderer

// NewConditionBuilderControl 创建一个新的 ConditionBuilderControl 实例，并设置默认类型
func NewConditionBuilderControl() ConditionBuilderControl {
	cc := ConditionBuilderControl(make(BaseRenderer))
	return cc.set("type", "condition-builder")
}

func (c ConditionBuilderControl) set(key string, value interface{}) ConditionBuilderControl {
	c[key] = value
	return c
}

// AddBtnVisibleOn 设置“添加按钮”显示条件
func (c ConditionBuilderControl) AddBtnVisibleOn(value string) ConditionBuilderControl {
	return c.set("addBtnVisibleOn", value)
}

// AddGroupBtnVisibleOn 设置“添加条件组”按钮显示条件
func (c ConditionBuilderControl) AddGroupBtnVisibleOn(value string) ConditionBuilderControl {
	return c.set("addGroupBtnVisibleOn", value)
}

// AutoFill 设置自动填充选项
func (c ConditionBuilderControl) AutoFill(value string) ConditionBuilderControl {
	return c.set("autoFill", value)
}

// BuilderMode 设置展现模式
func (c ConditionBuilderControl) BuilderMode(value string) ConditionBuilderControl {
	return c.set("builderMode", value)
}

// ClassName 设置容器 CSS 类名
func (c ConditionBuilderControl) ClassName(value string) ConditionBuilderControl {
	return c.set("className", value)
}

// ClearValueOnHidden 设置表单项隐藏时是否删除值
func (c ConditionBuilderControl) ClearValueOnHidden(value bool) ConditionBuilderControl {
	return c.set("clearValueOnHidden", value)
}

// Config 设置其他配置
func (c ConditionBuilderControl) Config(value string) ConditionBuilderControl {
	return c.set("config", value)
}

// Desc 设置描述
func (c ConditionBuilderControl) Desc(value string) ConditionBuilderControl {
	return c.set("desc", value)
}

// Description 设置描述内容，支持 HTML 片段
func (c ConditionBuilderControl) Description(value string) ConditionBuilderControl {
	return c.set("description", value)
}

// DescriptionClassName 设置描述内容的 CSS 类名
func (c ConditionBuilderControl) DescriptionClassName(value string) ConditionBuilderControl {
	return c.set("descriptionClassName", value)
}

// Disabled 设置是否禁用
func (c ConditionBuilderControl) Disabled(value bool) ConditionBuilderControl {
	return c.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (c ConditionBuilderControl) DisabledOn(value string) ConditionBuilderControl {
	return c.set("disabledOn", value)
}

// Draggable 设置是否可拖拽
func (c ConditionBuilderControl) Draggable(value bool) ConditionBuilderControl {
	return c.set("draggable", value)
}

// EditorSetting 设置编辑器配置
func (c ConditionBuilderControl) EditorSetting(value string) ConditionBuilderControl {
	return c.set("editorSetting", value)
}

// Embed 设置内嵌模式
func (c ConditionBuilderControl) Embed(value bool) ConditionBuilderControl {
	return c.set("embed", value)
}

// ExtraName 设置额外字段名
func (c ConditionBuilderControl) ExtraName(value string) ConditionBuilderControl {
	return c.set("extraName", value)
}

// Fields 设置字段集合
func (c ConditionBuilderControl) Fields(value string) ConditionBuilderControl {
	return c.set("fields", value)
}

// Formula 设置字段输入控件为公式编辑器
func (c ConditionBuilderControl) Formula(value string) ConditionBuilderControl {
	return c.set("formula", value)
}

// FormulaForIf 设置 if 里面公式编辑器配置
func (c ConditionBuilderControl) FormulaForIf(value string) ConditionBuilderControl {
	return c.set("formulaForIf", value)
}

// Funcs 设置函数集合
func (c ConditionBuilderControl) Funcs(value string) ConditionBuilderControl {
	return c.set("funcs", value)
}

// Hidden 设置是否隐藏
func (c ConditionBuilderControl) Hidden(value bool) ConditionBuilderControl {
	return c.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (c ConditionBuilderControl) HiddenOn(value string) ConditionBuilderControl {
	return c.set("hiddenOn", value)
}

// Hint 设置输入提示
func (c ConditionBuilderControl) Hint(value string) ConditionBuilderControl {
	return c.set("hint", value)
}

// Horizontal 设置水平布局配置
func (c ConditionBuilderControl) Horizontal(value string) ConditionBuilderControl {
	return c.set("horizontal", value)
}

// ID 设置组件唯一 ID
func (c ConditionBuilderControl) ID(value string) ConditionBuilderControl {
	return c.set("id", value)
}

// InitAutoFill 设置初始化自动填充
func (c ConditionBuilderControl) InitAutoFill(value string) ConditionBuilderControl {
	return c.set("initAutoFill", value)
}

// Inline 设置表单控制是否为 inline 模式
func (c ConditionBuilderControl) Inline(value bool) ConditionBuilderControl {
	return c.set("inline", value)
}

// InputClassName 设置输入的 CSS 类名
func (c ConditionBuilderControl) InputClassName(value string) ConditionBuilderControl {
	return c.set("inputClassName", value)
}

// Label 设置描述标题
func (c ConditionBuilderControl) Label(value string) ConditionBuilderControl {
	return c.set("label", value)
}

// LabelAlign 设置描述标题对齐
func (c ConditionBuilderControl) LabelAlign(value string) ConditionBuilderControl {
	return c.set("labelAlign", value)
}

// LabelClassName 设置描述标题的 CSS 类名
func (c ConditionBuilderControl) LabelClassName(value string) ConditionBuilderControl {
	return c.set("labelClassName", value)
}

// LabelRemark 设置描述标题旁的小图标备注
func (c ConditionBuilderControl) LabelRemark(value string) ConditionBuilderControl {
	return c.set("labelRemark", value)
}

// LabelWidth 设置描述标题的宽度
func (c ConditionBuilderControl) LabelWidth(value string) ConditionBuilderControl {
	return c.set("labelWidth", value)
}

// Mode 设置表单项展示模式
func (c ConditionBuilderControl) Mode(value string) ConditionBuilderControl {
	return c.set("mode", value)
}

// Name 设置字段名
func (c ConditionBuilderControl) Name(value string) ConditionBuilderControl {
	return c.set("name", value)
}

// OnEvent 设置事件动作配置
func (c ConditionBuilderControl) OnEvent(value string) ConditionBuilderControl {
	return c.set("onEvent", value)
}

// PickerIcon 设置弹窗触发图标
func (c ConditionBuilderControl) PickerIcon(value string) ConditionBuilderControl {
	return c.set("pickerIcon", value)
}

// Placeholder 设置占位符
func (c ConditionBuilderControl) Placeholder(value string) ConditionBuilderControl {
	return c.set("placeholder", value)
}

// ReadOnly 设置是否只读
func (c ConditionBuilderControl) ReadOnly(value bool) ConditionBuilderControl {
	return c.set("readOnly", value)
}

// ReadOnlyOn 设置只读条件
func (c ConditionBuilderControl) ReadOnlyOn(value string) ConditionBuilderControl {
	return c.set("readOnlyOn", value)
}

// Remark 设置提示内容
func (c ConditionBuilderControl) Remark(value string) ConditionBuilderControl {
	return c.set("remark", value)
}

// Required 设置是否必填
func (c ConditionBuilderControl) Required(value bool) ConditionBuilderControl {
	return c.set("required", value)
}

// Row 设置行数
func (c ConditionBuilderControl) Row(value string) ConditionBuilderControl {
	return c.set("row", value)
}

// SaveImmediately 设置是否立即保存
func (c ConditionBuilderControl) SaveImmediately(value bool) ConditionBuilderControl {
	return c.set("saveImmediately", value)
}

// ShowANDOR 设置是否显示并或切换键按钮
func (c ConditionBuilderControl) ShowANDOR(value bool) ConditionBuilderControl {
	return c.set("showANDOR", value)
}

// Size 设置表单项大小
func (c ConditionBuilderControl) Size(value string) ConditionBuilderControl {
	return c.set("size", value)
}

// Source 设置远程配置项
func (c ConditionBuilderControl) Source(value string) ConditionBuilderControl {
	return c.set("source", value)
}

// Static 设置是否静态展示
func (c ConditionBuilderControl) Static(value bool) ConditionBuilderControl {
	return c.set("static", value)
}

// StaticClassName 设置静态展示的 CSS 类名
func (c ConditionBuilderControl) StaticClassName(value string) ConditionBuilderControl {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示的输入类名
func (c ConditionBuilderControl) StaticInputClassName(value string) ConditionBuilderControl {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示的标签类名
func (c ConditionBuilderControl) StaticLabelClassName(value string) ConditionBuilderControl {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置静态展示的表达式
func (c ConditionBuilderControl) StaticOn(value string) ConditionBuilderControl {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (c ConditionBuilderControl) StaticPlaceholder(value string) ConditionBuilderControl {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示模式
func (c ConditionBuilderControl) StaticSchema(value string) ConditionBuilderControl {
	return c.set("staticSchema", value)
}

// Style 设置组件样式
func (c ConditionBuilderControl) Style(value string) ConditionBuilderControl {
	return c.set("style", value)
}

// SubmitOnChange 设置修改时是否提交表单
func (c ConditionBuilderControl) SubmitOnChange(value bool) ConditionBuilderControl {
	return c.set("submitOnChange", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (c ConditionBuilderControl) TestIdBuilder(value string) ConditionBuilderControl {
	return c.set("testIdBuilder", value)
}

// UseMobileUI 设置是否使用移动端样式
func (c ConditionBuilderControl) UseMobileUI(value bool) ConditionBuilderControl {
	return c.set("useMobileUI", value)
}

// ValidateApi 设置远端校验表单项接口
func (c ConditionBuilderControl) ValidateApi(value string) ConditionBuilderControl {
	return c.set("validateApi", value)
}

// ValidateOnChange 设置是否在每次修改时触发验证
func (c ConditionBuilderControl) ValidateOnChange(value bool) ConditionBuilderControl {
	return c.set("validateOnChange", value)
}

// ValidationErrors 设置验证失败提示信息
func (c ConditionBuilderControl) ValidationErrors(value string) ConditionBuilderControl {
	return c.set("validationErrors", value)
}

// Validations 设置验证规则
func (c ConditionBuilderControl) Validations(value string) ConditionBuilderControl {
	return c.set("validations", value)
}

// Value 设置默认值
func (c ConditionBuilderControl) Value(value string) ConditionBuilderControl {
	return c.set("value", value)
}

// Visible 设置是否显示
func (c ConditionBuilderControl) Visible(value bool) ConditionBuilderControl {
	return c.set("visible", value)
}

// VisibleOn 设置显示表达式
func (c ConditionBuilderControl) VisibleOn(value string) ConditionBuilderControl {
	return c.set("visibleOn", value)
}

// Width 设置宽度
func (c ConditionBuilderControl) Width(value string) ConditionBuilderControl {
	return c.set("width", value)
}
