package comp

// TextareaControl 多行文本输入框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/textarea
//
// @author  slowlyo
// @version 6.7.0
type TextareaControl Schema

// NewTextareaControl 创建一个新的 TextareaControl 实例
func NewTextareaControl() TextareaControl {
	return TextareaControl{}.set("type", "textarea")
}

func (t TextareaControl) set(key string, value interface{}) TextareaControl {
	t[key] = value
	return t
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (t TextareaControl) AutoFill(value string) TextareaControl {
	return t.set("autoFill", value)
}

// BorderMode 边框模式，全边框，还是半边框，或者没边框。 可选值: full | half | none
func (t TextareaControl) BorderMode(value string) TextareaControl {
	return t.set("borderMode", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t TextareaControl) ClassName(value string) TextareaControl {
	return t.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (t TextareaControl) ClearValueOnHidden(value bool) TextareaControl {
	return t.set("clearValueOnHidden", value)
}

// Clearable 输入内容是否可清除
func (t TextareaControl) Clearable(value bool) TextareaControl {
	return t.set("clearable", value)
}

// Desc
func (t TextareaControl) Desc(value string) TextareaControl {
	return t.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (t TextareaControl) Description(value string) TextareaControl {
	return t.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (t TextareaControl) DescriptionClassName(value string) TextareaControl {
	return t.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (t TextareaControl) Disabled(value bool) TextareaControl {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (t TextareaControl) DisabledOn(value string) TextareaControl {
	return t.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t TextareaControl) EditorSetting(value string) TextareaControl {
	return t.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (t TextareaControl) ExtraName(value string) TextareaControl {
	return t.set("extraName", value)
}

// Hidden 是否隐藏
func (t TextareaControl) Hidden(value bool) TextareaControl {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (t TextareaControl) HiddenOn(value string) TextareaControl {
	return t.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (t TextareaControl) Hint(value string) TextareaControl {
	return t.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (t TextareaControl) Horizontal(value string) TextareaControl {
	return t.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (t TextareaControl) Id(value string) TextareaControl {
	return t.set("id", value)
}

// InitAutoFill
func (t TextareaControl) InitAutoFill(value string) TextareaControl {
	return t.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (t TextareaControl) Inline(value bool) TextareaControl {
	return t.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (t TextareaControl) InputClassName(value string) TextareaControl {
	return t.set("inputClassName", value)
}

// Label 描述标题
func (t TextareaControl) Label(value string) TextareaControl {
	return t.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (t TextareaControl) LabelAlign(value string) TextareaControl {
	return t.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (t TextareaControl) LabelClassName(value string) TextareaControl {
	return t.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (t TextareaControl) LabelRemark(value string) TextareaControl {
	return t.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (t TextareaControl) LabelWidth(value string) TextareaControl {
	return t.set("labelWidth", value)
}

// MaxLength 限制文字个数
func (t TextareaControl) MaxLength(value string) TextareaControl {
	return t.set("maxLength", value)
}

// MaxRows 最大行数
func (t TextareaControl) MaxRows(value string) TextareaControl {
	return t.set("maxRows", value)
}

// MinRows 最小行数
func (t TextareaControl) MinRows(value string) TextareaControl {
	return t.set("minRows", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (t TextareaControl) Mode(value string) TextareaControl {
	return t.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (t TextareaControl) Name(value string) TextareaControl {
	return t.set("name", value)
}

// OnEvent 事件动作配置
func (t TextareaControl) OnEvent(value string) TextareaControl {
	return t.set("onEvent", value)
}

// Placeholder 占位符
func (t TextareaControl) Placeholder(value string) TextareaControl {
	return t.set("placeholder", value)
}

// ReadOnly 是否只读
func (t TextareaControl) ReadOnly(value bool) TextareaControl {
	return t.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (t TextareaControl) ReadOnlyOn(value string) TextareaControl {
	return t.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (t TextareaControl) Remark(value string) TextareaControl {
	return t.set("remark", value)
}

// Required 是否为必填
func (t TextareaControl) Required(value bool) TextareaControl {
	return t.set("required", value)
}

// ResetValue 重置值
func (t TextareaControl) ResetValue(value string) TextareaControl {
	return t.set("resetValue", value)
}

// Row
func (t TextareaControl) Row(value string) TextareaControl {
	return t.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (t TextareaControl) SaveImmediately(value bool) TextareaControl {
	return t.set("saveImmediately", value)
}

// ShowCounter 是否显示计数
func (t TextareaControl) ShowCounter(value bool) TextareaControl {
	return t.set("showCounter", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (t TextareaControl) Size(value string) TextareaControl {
	return t.set("size", value)
}

// Static 是否静态展示
func (t TextareaControl) Static(value bool) TextareaControl {
	return t.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t TextareaControl) StaticClassName(value string) TextareaControl {
	return t.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t TextareaControl) StaticInputClassName(value string) TextareaControl {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t TextareaControl) StaticLabelClassName(value string) TextareaControl {
	return t.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (t TextareaControl) StaticOn(value string) TextareaControl {
	return t.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (t TextareaControl) StaticPlaceholder(value string) TextareaControl {
	return t.set("staticPlaceholder", value)
}

// StaticSchema
func (t TextareaControl) StaticSchema(value string) TextareaControl {
	return t.set("staticSchema", value)
}

// Style 组件样式
func (t TextareaControl) Style(value string) TextareaControl {
	return t.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (t TextareaControl) SubmitOnChange(value bool) TextareaControl {
	return t.set("submitOnChange", value)
}

// TestIdBuilder
func (t TextareaControl) TestIdBuilder(value string) TextareaControl {
	return t.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (t TextareaControl) UseMobileUI(value bool) TextareaControl {
	return t.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (t TextareaControl) ValidateApi(value string) TextareaControl {
	return t.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (t TextareaControl) ValidateOnChange(value bool) TextareaControl {
	return t.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (t TextareaControl) ValidationErrors(value string) TextareaControl {
	return t.set("validationErrors", value)
}

// Validations
func (t TextareaControl) Validations(value string) TextareaControl {
	return t.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (t TextareaControl) Value(value string) TextareaControl {
	return t.set("value", value)
}

// Visible 是否显示
func (t TextareaControl) Visible(value bool) TextareaControl {
	return t.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (t TextareaControl) VisibleOn(value string) TextareaControl {
	return t.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (t TextareaControl) Width(value string) TextareaControl {
	return t.set("width", value)
}
