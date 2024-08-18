package comp

// hidden 隐藏域。功能性组件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/hidden
type hidden schema

func Hidden() hidden {
	return make(hidden).set("type", "hidden")
}

func (h hidden) set(key string, value interface{}) hidden {
	h[key] = value
	return h
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (h hidden) AutoFill(value string) hidden {
	return h.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (h hidden) ClassName(value string) hidden {
	return h.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (h hidden) ClearValueOnHidden(value bool) hidden {
	return h.set("clearValueOnHidden", value)
}

// Desc
func (h hidden) Desc(value string) hidden {
	return h.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (h hidden) Description(value string) hidden {
	return h.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (h hidden) DescriptionClassName(value string) hidden {
	return h.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (h hidden) Disabled(value bool) hidden {
	return h.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (h hidden) DisabledOn(value string) hidden {
	return h.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (h hidden) EditorSetting(value string) hidden {
	return h.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (h hidden) ExtraName(value string) hidden {
	return h.set("extraName", value)
}

// Hidden 是否隐藏
func (h hidden) Hidden(value bool) hidden {
	return h.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (h hidden) HiddenOn(value string) hidden {
	return h.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (h hidden) Hint(value string) hidden {
	return h.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (h hidden) Horizontal(value string) hidden {
	return h.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (h hidden) Id(value string) hidden {
	return h.set("id", value)
}

// InitAutoFill
func (h hidden) InitAutoFill(value string) hidden {
	return h.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (h hidden) Inline(value bool) hidden {
	return h.set("inline", value)
}

// InputClassName 配置 input className
func (h hidden) InputClassName(value string) hidden {
	return h.set("inputClassName", value)
}

// Label 描述标题
func (h hidden) Label(value string) hidden {
	return h.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (h hidden) LabelAlign(value string) hidden {
	return h.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (h hidden) LabelClassName(value string) hidden {
	return h.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (h hidden) LabelRemark(value string) hidden {
	return h.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (h hidden) LabelWidth(value string) hidden {
	return h.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (h hidden) Mode(value string) hidden {
	return h.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (h hidden) Name(value string) hidden {
	return h.set("name", value)
}

// OnEvent 事件动作配置
func (h hidden) OnEvent(value string) hidden {
	return h.set("onEvent", value)
}

// Placeholder 占位符
func (h hidden) Placeholder(value string) hidden {
	return h.set("placeholder", value)
}

// ReadOnly 是否只读
func (h hidden) ReadOnly(value bool) hidden {
	return h.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (h hidden) ReadOnlyOn(value string) hidden {
	return h.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (h hidden) Remark(value string) hidden {
	return h.set("remark", value)
}

// Required 是否为必填
func (h hidden) Required(value bool) hidden {
	return h.set("required", value)
}

// Row
func (h hidden) Row(value string) hidden {
	return h.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (h hidden) SaveImmediately(value bool) hidden {
	return h.set("saveImmediately", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (h hidden) Size(value string) hidden {
	return h.set("size", value)
}

// Static 是否静态展示
func (h hidden) Static(value bool) hidden {
	return h.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (h hidden) StaticClassName(value string) hidden {
	return h.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (h hidden) StaticInputClassName(value string) hidden {
	return h.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (h hidden) StaticLabelClassName(value string) hidden {
	return h.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (h hidden) StaticOn(value string) hidden {
	return h.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (h hidden) StaticPlaceholder(value string) hidden {
	return h.set("staticPlaceholder", value)
}

// StaticSchema
func (h hidden) StaticSchema(value string) hidden {
	return h.set("staticSchema", value)
}

// Style 组件样式
func (h hidden) Style(value string) hidden {
	return h.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (h hidden) SubmitOnChange(value bool) hidden {
	return h.set("submitOnChange", value)
}

// TestIdBuilder
func (h hidden) TestIdBuilder(value string) hidden {
	return h.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (h hidden) UseMobileUI(value bool) hidden {
	return h.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (h hidden) ValidateApi(value string) hidden {
	return h.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (h hidden) ValidateOnChange(value bool) hidden {
	return h.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (h hidden) ValidationErrors(value string) hidden {
	return h.set("validationErrors", value)
}

// Validations
func (h hidden) Validations(value string) hidden {
	return h.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (h hidden) Value(value string) hidden {
	return h.set("value", value)
}

// Visible 是否显示
func (h hidden) Visible(value bool) hidden {
	return h.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (h hidden) VisibleOn(value string) hidden {
	return h.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (h hidden) Width(value string) hidden {
	return h.set("width", value)
}
