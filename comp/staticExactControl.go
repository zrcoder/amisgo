package comp

// StaticExactControl 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/static
//
// @version 6.7.0
type StaticExactControl Schema

// NewStaticExactControl 创建一个新的 StaticExactControl 实例
func NewStaticExactControl() StaticExactControl {
	return StaticExactControl{}.set("type", "static")
}

func (s StaticExactControl) set(key string, value interface{}) StaticExactControl {
	s[key] = value
	return s
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (s StaticExactControl) AutoFill(value string) StaticExactControl {
	return s.set("autoFill", value)
}

// BorderMode 边框模式，默认是无边框的 可选值: full | half | none
func (s StaticExactControl) BorderMode(value string) StaticExactControl {
	return s.set("borderMode", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s StaticExactControl) ClassName(value string) StaticExactControl {
	return s.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (s StaticExactControl) ClearValueOnHidden(value bool) StaticExactControl {
	return s.set("clearValueOnHidden", value)
}

// Copyable 配置点击复制功能 (配置点击复制功能)
func (s StaticExactControl) Copyable(value bool) StaticExactControl {
	return s.set("copyable", value)
}

// Desc
func (s StaticExactControl) Desc(value string) StaticExactControl {
	return s.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (s StaticExactControl) Description(value string) StaticExactControl {
	return s.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (s StaticExactControl) DescriptionClassName(value string) StaticExactControl {
	return s.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (s StaticExactControl) Disabled(value bool) StaticExactControl {
	return s.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (s StaticExactControl) DisabledOn(value string) StaticExactControl {
	return s.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (s StaticExactControl) EditorSetting(value string) StaticExactControl {
	return s.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (s StaticExactControl) ExtraName(value string) StaticExactControl {
	return s.set("extraName", value)
}

// Hidden 是否隐藏
func (s StaticExactControl) Hidden(value bool) StaticExactControl {
	return s.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (s StaticExactControl) HiddenOn(value string) StaticExactControl {
	return s.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (s StaticExactControl) Hint(value string) StaticExactControl {
	return s.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (s StaticExactControl) Horizontal(value string) StaticExactControl {
	return s.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (s StaticExactControl) Id(value string) StaticExactControl {
	return s.set("id", value)
}

// InitAutoFill
func (s StaticExactControl) InitAutoFill(value string) StaticExactControl {
	return s.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (s StaticExactControl) Inline(value bool) StaticExactControl {
	return s.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (s StaticExactControl) InputClassName(value string) StaticExactControl {
	return s.set("inputClassName", value)
}

// Label 描述标题
func (s StaticExactControl) Label(value string) StaticExactControl {
	return s.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (s StaticExactControl) LabelAlign(value string) StaticExactControl {
	return s.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (s StaticExactControl) LabelClassName(value string) StaticExactControl {
	return s.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (s StaticExactControl) LabelRemark(value string) StaticExactControl {
	return s.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (s StaticExactControl) LabelWidth(value string) StaticExactControl {
	return s.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (s StaticExactControl) Mode(value string) StaticExactControl {
	return s.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (s StaticExactControl) Name(value string) StaticExactControl {
	return s.set("name", value)
}

// OnEvent 事件动作配置
func (s StaticExactControl) OnEvent(value string) StaticExactControl {
	return s.set("onEvent", value)
}

// Placeholder 占位符
func (s StaticExactControl) Placeholder(value string) StaticExactControl {
	return s.set("placeholder", value)
}

// PopOver 配置查看详情功能 (配置查看详情功能)
func (s StaticExactControl) PopOver(value string) StaticExactControl {
	return s.set("popOver", value)
}

// QuickEdit 配置快速编辑功能 (配置快速编辑功能)
func (s StaticExactControl) QuickEdit(value string) StaticExactControl {
	return s.set("quickEdit", value)
}

// ReadOnly 是否只读
func (s StaticExactControl) ReadOnly(value bool) StaticExactControl {
	return s.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (s StaticExactControl) ReadOnlyOn(value string) StaticExactControl {
	return s.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (s StaticExactControl) Remark(value string) StaticExactControl {
	return s.set("remark", value)
}

// Required 是否为必填
func (s StaticExactControl) Required(value bool) StaticExactControl {
	return s.set("required", value)
}

// Row
func (s StaticExactControl) Row(value string) StaticExactControl {
	return s.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (s StaticExactControl) SaveImmediately(value bool) StaticExactControl {
	return s.set("saveImmediately", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (s StaticExactControl) Size(value string) StaticExactControl {
	return s.set("size", value)
}

// Static 是否静态展示
func (s StaticExactControl) Static(value bool) StaticExactControl {
	return s.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s StaticExactControl) StaticClassName(value string) StaticExactControl {
	return s.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s StaticExactControl) StaticInputClassName(value string) StaticExactControl {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (s StaticExactControl) StaticLabelClassName(value string) StaticExactControl {
	return s.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (s StaticExactControl) StaticOn(value string) StaticExactControl {
	return s.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (s StaticExactControl) StaticPlaceholder(value string) StaticExactControl {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s StaticExactControl) StaticSchema(value string) StaticExactControl {
	return s.set("staticSchema", value)
}

// Style 组件样式
func (s StaticExactControl) Style(value string) StaticExactControl {
	return s.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (s StaticExactControl) SubmitOnChange(value bool) StaticExactControl {
	return s.set("submitOnChange", value)
}

// TestIdBuilder
func (s StaticExactControl) TestIdBuilder(value string) StaticExactControl {
	return s.set("testIdBuilder", value)
}

// Text 内容模板，不支持 HTML (支持两种语法，但是不能混着用。分别是：1. `${xxx}` 或者 `${xxx|upperCase}` 2. `<%= data.xxx %>`
// 更多文档：https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template)
func (s StaticExactControl) Text(value string) StaticExactControl {
	return s.set("text", value)
}

// Tpl 内容模板， 支持 HTML (支持两种语法，但是不能混着用。分别是：1. `${xxx}` 或者 `${xxx|upperCase}` 2. `<%= data.xxx %>`
// 更多文档：https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template)
func (s StaticExactControl) Tpl(value string) StaticExactControl {
	return s.set("tpl", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (s StaticExactControl) UseMobileUI(value bool) StaticExactControl {
	return s.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (s StaticExactControl) ValidateApi(value string) StaticExactControl {
	return s.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (s StaticExactControl) ValidateOnChange(value bool) StaticExactControl {
	return s.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (s StaticExactControl) ValidationErrors(value string) StaticExactControl {
	return s.set("validationErrors", value)
}

// Validations
func (s StaticExactControl) Validations(value string) StaticExactControl {
	return s.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (s StaticExactControl) Value(value string) StaticExactControl {
	return s.set("value", value)
}

// Visible 是否显示
func (s StaticExactControl) Visible(value bool) StaticExactControl {
	return s.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (s StaticExactControl) VisibleOn(value string) StaticExactControl {
	return s.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (s StaticExactControl) Width(value string) StaticExactControl {
	return s.set("width", value)
}
