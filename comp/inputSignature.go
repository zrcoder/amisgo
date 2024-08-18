package comp

// inputSignature
//

// @version 6.7.0
type inputSignature schema

// InputSignature 创建一个新的 InputSignature 实例，并设置默认的 type
func InputSignature() inputSignature {
	return make(inputSignature).set("type", "input-signature")
}

// set 是一个内部方法，用于设置字段值并返回自身的引用
func (i inputSignature) set(key string, value interface{}) inputSignature {
	i[key] = value
	return i
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (i inputSignature) AutoFill(value string) inputSignature {
	return i.set("autoFill", value)
}

// BgColor 组件背景颜色
func (i inputSignature) BgColor(value string) inputSignature {
	return i.set("bgColor", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (i inputSignature) ClassName(value string) inputSignature {
	return i.set("className", value)
}

// ClearBtnIcon 清空按钮图标
func (i inputSignature) ClearBtnIcon(value string) inputSignature {
	return i.set("clearBtnIcon", value)
}

// ClearBtnLabel 清空按钮名称
func (i inputSignature) ClearBtnLabel(value string) inputSignature {
	return i.set("clearBtnLabel", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (i inputSignature) ClearValueOnHidden(value bool) inputSignature {
	return i.set("clearValueOnHidden", value)
}

// Color 组件字段颜色
func (i inputSignature) Color(value string) inputSignature {
	return i.set("color", value)
}

// ConfirmBtnIcon 确认按钮图标
func (i inputSignature) ConfirmBtnIcon(value string) inputSignature {
	return i.set("confirmBtnIcon", value)
}

// ConfirmBtnLabel 确认按钮名称
func (i inputSignature) ConfirmBtnLabel(value string) inputSignature {
	return i.set("confirmBtnLabel", value)
}

// Desc
func (i inputSignature) Desc(value string) inputSignature {
	return i.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (i inputSignature) Description(value string) inputSignature {
	return i.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (i inputSignature) DescriptionClassName(value string) inputSignature {
	return i.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (i inputSignature) Disabled(value bool) inputSignature {
	return i.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (i inputSignature) DisabledOn(value string) inputSignature {
	return i.set("disabledOn", value)
}

// EbmedCancelIcon 弹窗取消按钮图标
func (i inputSignature) EbmedCancelIcon(value string) inputSignature {
	return i.set("ebmedCancelIcon", value)
}

// EbmedCancelLabel 弹窗取消按钮名称
func (i inputSignature) EbmedCancelLabel(value string) inputSignature {
	return i.set("ebmedCancelLabel", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (i inputSignature) EditorSetting(value string) inputSignature {
	return i.set("editorSetting", value)
}

// Embed 是否内嵌
func (i inputSignature) Embed(value bool) inputSignature {
	return i.set("embed", value)
}

// EmbedBtnIcon 弹窗按钮图标
func (i inputSignature) EmbedBtnIcon(value string) inputSignature {
	return i.set("embedBtnIcon", value)
}

// EmbedBtnLabel 弹窗按钮文案
func (i inputSignature) EmbedBtnLabel(value string) inputSignature {
	return i.set("embedBtnLabel", value)
}

// EmbedConfirmIcon 弹窗确认按钮图标
func (i inputSignature) EmbedConfirmIcon(value string) inputSignature {
	return i.set("embedConfirmIcon", value)
}

// EmbedConfirmLabel 弹窗确认按钮名称
func (i inputSignature) EmbedConfirmLabel(value string) inputSignature {
	return i.set("embedConfirmLabel", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (i inputSignature) ExtraName(value string) inputSignature {
	return i.set("extraName", value)
}

// Height 组件高度，默认占满父容器
func (i inputSignature) Height(value string) inputSignature {
	return i.set("height", value)
}

// Hidden 是否隐藏
func (i inputSignature) Hidden(value bool) inputSignature {
	return i.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (i inputSignature) HiddenOn(value string) inputSignature {
	return i.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (i inputSignature) Hint(value string) inputSignature {
	return i.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (i inputSignature) Horizontal(value string) inputSignature {
	return i.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (i inputSignature) Id(value string) inputSignature {
	return i.set("id", value)
}

// InitAutoFill
func (i inputSignature) InitAutoFill(value string) inputSignature {
	return i.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式
func (i inputSignature) Inline(value bool) inputSignature {
	return i.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (i inputSignature) InputClassName(value string) inputSignature {
	return i.set("inputClassName", value)
}

// Label 描述标题
func (i inputSignature) Label(value string) inputSignature {
	return i.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (i inputSignature) LabelAlign(value string) inputSignature {
	return i.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (i inputSignature) LabelClassName(value string) inputSignature {
	return i.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (i inputSignature) LabelRemark(value string) inputSignature {
	return i.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (i inputSignature) LabelWidth(value string) inputSignature {
	return i.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (i inputSignature) Mode(value string) inputSignature {
	return i.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (i inputSignature) Name(value string) inputSignature {
	return i.set("name", value)
}

// OnEvent 事件动作配置
func (i inputSignature) OnEvent(value string) inputSignature {
	return i.set("onEvent", value)
}

// Placeholder 占位符
func (i inputSignature) Placeholder(value string) inputSignature {
	return i.set("placeholder", value)
}

// ReadOnly 是否只读
func (i inputSignature) ReadOnly(value bool) inputSignature {
	return i.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (i inputSignature) ReadOnlyOn(value string) inputSignature {
	return i.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (i inputSignature) Remark(value string) inputSignature {
	return i.set("remark", value)
}

// Required 是否为必填
func (i inputSignature) Required(value bool) inputSignature {
	return i.set("required", value)
}

// Row
func (i inputSignature) Row(value string) inputSignature {
	return i.set("row", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (i inputSignature) Size(value string) inputSignature {
	return i.set("size", value)
}

// Static 是否静态展示
func (i inputSignature) Static(value bool) inputSignature {
	return i.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (i inputSignature) StaticClassName(value string) inputSignature {
	return i.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (i inputSignature) StaticInputClassName(value string) inputSignature {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (i inputSignature) StaticLabelClassName(value string) inputSignature {
	return i.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (i inputSignature) StaticOn(value string) inputSignature {
	return i.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (i inputSignature) StaticPlaceholder(value string) inputSignature {
	return i.set("staticPlaceholder", value)
}

// StaticSchema
func (i inputSignature) StaticSchema(value string) inputSignature {
	return i.set("staticSchema", value)
}

// Style 组件样式
func (i inputSignature) Style(value string) inputSignature {
	return i.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (i inputSignature) SubmitOnChange(value bool) inputSignature {
	return i.set("submitOnChange", value)
}

// TestIdBuilder
func (i inputSignature) TestIdBuilder(value string) inputSignature {
	return i.set("testIdBuilder", value)
}

// UndoBtnIcon 清空按钮图标
func (i inputSignature) UndoBtnIcon(value string) inputSignature {
	return i.set("undoBtnIcon", value)
}

// UndoBtnLabel 撤销按钮名称
func (i inputSignature) UndoBtnLabel(value string) inputSignature {
	return i.set("undoBtnLabel", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (i inputSignature) UseMobileUI(value bool) inputSignature {
	return i.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (i inputSignature) ValidateApi(value string) inputSignature {
	return i.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证，如果设置了，则由此配置项来决定要不要每次修改都触发验证
func (i inputSignature) ValidateOnChange(value bool) inputSignature {
	return i.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (i inputSignature) ValidationErrors(value string) inputSignature {
	return i.set("validationErrors", value)
}

// Validations
func (i inputSignature) Validations(value string) inputSignature {
	return i.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的
func (i inputSignature) Value(value string) inputSignature {
	return i.set("value", value)
}

// Visible 是否显示
func (i inputSignature) Visible(value bool) inputSignature {
	return i.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (i inputSignature) VisibleOn(value string) inputSignature {
	return i.set("visibleOn", value)
}

// Width 组件宽度，默认占满父容器
func (i inputSignature) Width(value string) inputSignature {
	return i.set("width", value)
}
