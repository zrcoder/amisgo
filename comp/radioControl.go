package comp

// RadioControl 单选框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/radios
//
// @version 6.7.0
type RadioControl Schema

// NewRadioControl 创建一个新的 RadioControl 实例
func NewRadioControl() RadioControl {
	return RadioControl{}.set("type", "radio")
}

// set 是一个私有方法，用于设置字段值
func (rc RadioControl) set(key string, value interface{}) RadioControl {
	rc[key] = value
	return rc
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (rc RadioControl) AutoFill(value string) RadioControl {
	return rc.set("autoFill", value)
}

// Badge 角标 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (rc RadioControl) Badge(value string) RadioControl {
	return rc.set("badge", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (rc RadioControl) ClassName(value string) RadioControl {
	return rc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (rc RadioControl) ClearValueOnHidden(value bool) RadioControl {
	return rc.set("clearValueOnHidden", value)
}

// Desc
func (rc RadioControl) Desc(value string) RadioControl {
	return rc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (rc RadioControl) Description(value string) RadioControl {
	return rc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (rc RadioControl) DescriptionClassName(value string) RadioControl {
	return rc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (rc RadioControl) Disabled(value bool) RadioControl {
	return rc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (rc RadioControl) DisabledOn(value string) RadioControl {
	return rc.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (rc RadioControl) EditorSetting(value string) RadioControl {
	return rc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (rc RadioControl) ExtraName(value string) RadioControl {
	return rc.set("extraName", value)
}

// FalseValue 未勾选值
func (rc RadioControl) FalseValue(value string) RadioControl {
	return rc.set("falseValue", value)
}

// Hidden 是否隐藏
func (rc RadioControl) Hidden(value bool) RadioControl {
	return rc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (rc RadioControl) HiddenOn(value string) RadioControl {
	return rc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (rc RadioControl) Hint(value string) RadioControl {
	return rc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (rc RadioControl) Horizontal(value string) RadioControl {
	return rc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (rc RadioControl) Id(value string) RadioControl {
	return rc.set("id", value)
}

// InitAutoFill
func (rc RadioControl) InitAutoFill(value string) RadioControl {
	return rc.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (rc RadioControl) Inline(value bool) RadioControl {
	return rc.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (rc RadioControl) InputClassName(value string) RadioControl {
	return rc.set("inputClassName", value)
}

// Label 描述标题
func (rc RadioControl) Label(value string) RadioControl {
	return rc.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (rc RadioControl) LabelAlign(value string) RadioControl {
	return rc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (rc RadioControl) LabelClassName(value string) RadioControl {
	return rc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (rc RadioControl) LabelRemark(value string) RadioControl {
	return rc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (rc RadioControl) LabelWidth(value string) RadioControl {
	return rc.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (rc RadioControl) Mode(value string) RadioControl {
	return rc.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (rc RadioControl) Name(value string) RadioControl {
	return rc.set("name", value)
}

// OnEvent 事件动作配置
func (rc RadioControl) OnEvent(value string) RadioControl {
	return rc.set("onEvent", value)
}

// Option 选项说明
func (rc RadioControl) Option(value string) RadioControl {
	return rc.set("option", value)
}

// OptionType 可选值: default | button
func (rc RadioControl) OptionType(value string) RadioControl {
	return rc.set("optionType", value)
}

// Partial
func (rc RadioControl) Partial(value bool) RadioControl {
	return rc.set("partial", value)
}

// Placeholder 占位符
func (rc RadioControl) Placeholder(value string) RadioControl {
	return rc.set("placeholder", value)
}

// ReadOnly 是否只读
func (rc RadioControl) ReadOnly(value bool) RadioControl {
	return rc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (rc RadioControl) ReadOnlyOn(value string) RadioControl {
	return rc.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (rc RadioControl) Remark(value string) RadioControl {
	return rc.set("remark", value)
}

// Required 是否为必填
func (rc RadioControl) Required(value bool) RadioControl {
	return rc.set("required", value)
}

// Row
func (rc RadioControl) Row(value string) RadioControl {
	return rc.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (rc RadioControl) SaveImmediately(value bool) RadioControl {
	return rc.set("saveImmediately", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (rc RadioControl) Size(value string) RadioControl {
	return rc.set("size", value)
}

// Static 是否静态展示
func (rc RadioControl) Static(value bool) RadioControl {
	return rc.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (rc RadioControl) StaticClassName(value string) RadioControl {
	return rc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (rc RadioControl) StaticInputClassName(value string) RadioControl {
	return rc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (rc RadioControl) StaticLabelClassName(value string) RadioControl {
	return rc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (rc RadioControl) StaticOn(value string) RadioControl {
	return rc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (rc RadioControl) StaticPlaceholder(value string) RadioControl {
	return rc.set("staticPlaceholder", value)
}

// StaticSchema
func (rc RadioControl) StaticSchema(value string) RadioControl {
	return rc.set("staticSchema", value)
}

// Style 组件样式
func (rc RadioControl) Style(value string) RadioControl {
	return rc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (rc RadioControl) SubmitOnChange(value bool) RadioControl {
	return rc.set("submitOnChange", value)
}

// TestIdBuilder
func (rc RadioControl) TestIdBuilder(value string) RadioControl {
	return rc.set("testIdBuilder", value)
}

// TrueValue 勾选值
func (rc RadioControl) TrueValue(value string) RadioControl {
	return rc.set("trueValue", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (rc RadioControl) UseMobileUI(value bool) RadioControl {
	return rc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (rc RadioControl) ValidateApi(value string) RadioControl {
	return rc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (rc RadioControl) ValidateOnChange(value bool) RadioControl {
	return rc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (rc RadioControl) ValidationErrors(value string) RadioControl {
	return rc.set("validationErrors", value)
}

// Validations
func (rc RadioControl) Validations(value string) RadioControl {
	return rc.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (rc RadioControl) Value(value string) RadioControl {
	return rc.set("value", value)
}

// Visible 是否显示
func (rc RadioControl) Visible(value bool) RadioControl {
	return rc.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (rc RadioControl) VisibleOn(value string) RadioControl {
	return rc.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (rc RadioControl) Width(value string) RadioControl {
	return rc.set("width", value)
}
