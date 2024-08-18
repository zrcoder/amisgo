package comp

// SwitchControl 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/switch
//
// @version 6.7.0
type SwitchControl Schema

// NewSwitchControl 创建一个新的 SwitchControl 实例
func NewSwitchControl() SwitchControl {
	return SwitchControl{}.set("type", "switch")
}

func (s SwitchControl) set(key string, value interface{}) SwitchControl {
	s[key] = value
	return s
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (s SwitchControl) AutoFill(value string) SwitchControl {
	return s.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s SwitchControl) ClassName(value string) SwitchControl {
	return s.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (s SwitchControl) ClearValueOnHidden(value bool) SwitchControl {
	return s.set("clearValueOnHidden", value)
}

// Desc
func (s SwitchControl) Desc(value string) SwitchControl {
	return s.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (s SwitchControl) Description(value string) SwitchControl {
	return s.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (s SwitchControl) DescriptionClassName(value string) SwitchControl {
	return s.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (s SwitchControl) Disabled(value bool) SwitchControl {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s SwitchControl) DisabledOn(value string) SwitchControl {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (s SwitchControl) EditorSetting(value string) SwitchControl {
	return s.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (s SwitchControl) ExtraName(value string) SwitchControl {
	return s.set("extraName", value)
}

// FalseValue 未勾选值
func (s SwitchControl) FalseValue(value string) SwitchControl {
	return s.set("falseValue", value)
}

// Hidden 是否隐藏
func (s SwitchControl) Hidden(value bool) SwitchControl {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s SwitchControl) HiddenOn(value string) SwitchControl {
	return s.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (s SwitchControl) Hint(value string) SwitchControl {
	return s.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (s SwitchControl) Horizontal(value string) SwitchControl {
	return s.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s SwitchControl) Id(value string) SwitchControl {
	return s.set("id", value)
}

// InitAutoFill
func (s SwitchControl) InitAutoFill(value string) SwitchControl {
	return s.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (s SwitchControl) Inline(value bool) SwitchControl {
	return s.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (s SwitchControl) InputClassName(value string) SwitchControl {
	return s.set("inputClassName", value)
}

// Label 描述标题
func (s SwitchControl) Label(value string) SwitchControl {
	return s.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (s SwitchControl) LabelAlign(value string) SwitchControl {
	return s.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (s SwitchControl) LabelClassName(value string) SwitchControl {
	return s.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (s SwitchControl) LabelRemark(value string) SwitchControl {
	return s.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (s SwitchControl) LabelWidth(value string) SwitchControl {
	return s.set("labelWidth", value)
}

// Loading 是否处于加载状态
func (s SwitchControl) Loading(value bool) SwitchControl {
	return s.set("loading", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (s SwitchControl) Mode(value string) SwitchControl {
	return s.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (s SwitchControl) Name(value string) SwitchControl {
	return s.set("name", value)
}

// OffText 关闭时显示的内容
func (s SwitchControl) OffText(value string) SwitchControl {
	return s.set("offText", value)
}

// OnEvent 事件动作配置
func (s SwitchControl) OnEvent(value string) SwitchControl {
	return s.set("onEvent", value)
}

// OnText 开启时显示的内容
func (s SwitchControl) OnText(value string) SwitchControl {
	return s.set("onText", value)
}

// Option 选项说明
func (s SwitchControl) Option(value string) SwitchControl {
	return s.set("option", value)
}

// Placeholder 占位符
func (s SwitchControl) Placeholder(value string) SwitchControl {
	return s.set("placeholder", value)
}

// ReadOnly 是否只读
func (s SwitchControl) ReadOnly(value bool) SwitchControl {
	return s.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (s SwitchControl) ReadOnlyOn(value string) SwitchControl {
	return s.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (s SwitchControl) Remark(value string) SwitchControl {
	return s.set("remark", value)
}

// Required 是否为必填
func (s SwitchControl) Required(value bool) SwitchControl {
	return s.set("required", value)
}

// Row
func (s SwitchControl) Row(value string) SwitchControl {
	return s.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (s SwitchControl) SaveImmediately(value bool) SwitchControl {
	return s.set("saveImmediately", value)
}

// Size 开关尺寸 可选值: sm | md
func (s SwitchControl) Size(value string) SwitchControl {
	return s.set("size", value)
}

// Static 是否静态展示
func (s SwitchControl) Static(value bool) SwitchControl {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s SwitchControl) StaticClassName(value string) SwitchControl {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s SwitchControl) StaticInputClassName(value string) SwitchControl {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (s SwitchControl) StaticLabelClassName(value string) SwitchControl {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (s SwitchControl) StaticOn(value string) SwitchControl {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s SwitchControl) StaticPlaceholder(value string) SwitchControl {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s SwitchControl) StaticSchema(value string) SwitchControl {
	return s.set("staticSchema", value)
}

// Style 组件样式
func (s SwitchControl) Style(value string) SwitchControl {
	return s.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (s SwitchControl) SubmitOnChange(value bool) SwitchControl {
	return s.set("submitOnChange", value)
}

// TestIdBuilder
func (s SwitchControl) TestIdBuilder(value string) SwitchControl {
	return s.set("testIdBuilder", value)
}

// TrueValue 勾选值
func (s SwitchControl) TrueValue(value string) SwitchControl {
	return s.set("trueValue", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s SwitchControl) UseMobileUI(value bool) SwitchControl {
	return s.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (s SwitchControl) ValidateApi(value string) SwitchControl {
	return s.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (s SwitchControl) ValidateOnChange(value bool) SwitchControl {
	return s.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (s SwitchControl) ValidationErrors(value string) SwitchControl {
	return s.set("validationErrors", value)
}

// Validations
func (s SwitchControl) Validations(value string) SwitchControl {
	return s.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (s SwitchControl) Value(value string) SwitchControl {
	return s.set("value", value)
}

// Visible 是否显示
func (s SwitchControl) Visible(value bool) SwitchControl {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (s SwitchControl) VisibleOn(value string) SwitchControl {
	return s.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (s SwitchControl) Width(value string) SwitchControl {
	return s.set("width", value)
}
