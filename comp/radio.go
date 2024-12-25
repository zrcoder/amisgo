package comp

// radio 单选框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/radios

type radio Schema

// Radio 创建一个新的 RadioControl 实例
func Radio() radio {
	return radio{}.set("type", "radio")
}

// set 是一个私有方法，用于设置字段值
func (rc radio) set(key string, value any) radio {
	rc[key] = value
	return rc
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (rc radio) AutoFill(value string) radio {
	return rc.set("autoFill", value)
}

// Badge 角标 (Badge 角标。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/badge)
func (rc radio) Badge(value string) radio {
	return rc.set("badge", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (rc radio) ClassName(value string) radio {
	return rc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (rc radio) ClearValueOnHidden(value bool) radio {
	return rc.set("clearValueOnHidden", value)
}

// Desc
func (rc radio) Desc(value string) radio {
	return rc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (rc radio) Description(value string) radio {
	return rc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (rc radio) DescriptionClassName(value string) radio {
	return rc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (rc radio) Disabled(value bool) radio {
	return rc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (rc radio) DisabledOn(value string) radio {
	return rc.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (rc radio) EditorSetting(value string) radio {
	return rc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (rc radio) ExtraName(value string) radio {
	return rc.set("extraName", value)
}

// FalseValue 未勾选值
func (rc radio) FalseValue(value string) radio {
	return rc.set("falseValue", value)
}

// Hidden 是否隐藏
func (rc radio) Hidden(value bool) radio {
	return rc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (rc radio) HiddenOn(value string) radio {
	return rc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (rc radio) Hint(value string) radio {
	return rc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (rc radio) Horizontal(value string) radio {
	return rc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (rc radio) ID(value string) radio {
	return rc.set("id", value)
}

// InitAutoFill
func (rc radio) InitAutoFill(value string) radio {
	return rc.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (rc radio) Inline(value bool) radio {
	return rc.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (rc radio) InputClassName(value string) radio {
	return rc.set("inputClassName", value)
}

// Label 描述标题
func (rc radio) Label(value string) radio {
	return rc.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (rc radio) LabelAlign(value string) radio {
	return rc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (rc radio) LabelClassName(value string) radio {
	return rc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (rc radio) LabelRemark(value string) radio {
	return rc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (rc radio) LabelWidth(value string) radio {
	return rc.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (rc radio) Mode(value string) radio {
	return rc.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (rc radio) Name(value string) radio {
	return rc.set("name", value)
}

// OnEvent 事件动作配置
func (rc radio) OnEvent(value any) radio {
	return rc.set("onEvent", value)
}

// Option 选项说明
func (rc radio) Option(value string) radio {
	return rc.set("option", value)
}

// OptionType 可选值: default | button
func (rc radio) OptionType(value string) radio {
	return rc.set("optionType", value)
}

// Partial
func (rc radio) Partial(value bool) radio {
	return rc.set("partial", value)
}

// Placeholder 占位符
func (rc radio) Placeholder(value string) radio {
	return rc.set("placeholder", value)
}

// ReadOnly 是否只读
func (rc radio) ReadOnly(value bool) radio {
	return rc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (rc radio) ReadOnlyOn(value string) radio {
	return rc.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (rc radio) Remark(value string) radio {
	return rc.set("remark", value)
}

// Required 是否为必填
func (rc radio) Required(value bool) radio {
	return rc.set("required", value)
}

// Row
func (rc radio) Row(value string) radio {
	return rc.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (rc radio) SaveImmediately(value bool) radio {
	return rc.set("saveImmediately", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (rc radio) Size(value string) radio {
	return rc.set("size", value)
}

// Static 是否静态展示
func (rc radio) Static(value bool) radio {
	return rc.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (rc radio) StaticClassName(value string) radio {
	return rc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (rc radio) StaticInputClassName(value string) radio {
	return rc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (rc radio) StaticLabelClassName(value string) radio {
	return rc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (rc radio) StaticOn(value string) radio {
	return rc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (rc radio) StaticPlaceholder(value string) radio {
	return rc.set("staticPlaceholder", value)
}

// StaticSchema
func (rc radio) StaticSchema(value string) radio {
	return rc.set("staticSchema", value)
}

// Style 组件样式
func (rc radio) Style(value any) radio {
	return rc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (rc radio) SubmitOnChange(value bool) radio {
	return rc.set("submitOnChange", value)
}

// TestIdBuilder
func (rc radio) TestIdBuilder(value string) radio {
	return rc.set("testIdBuilder", value)
}

// TrueValue 勾选值
func (rc radio) TrueValue(value string) radio {
	return rc.set("trueValue", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (rc radio) UseMobileUI(value bool) radio {
	return rc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (rc radio) ValidateApi(value string) radio {
	return rc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (rc radio) ValidateOnChange(value bool) radio {
	return rc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (rc radio) ValidationErrors(value string) radio {
	return rc.set("validationErrors", value)
}

// Validations
func (rc radio) Validations(value string) radio {
	return rc.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (rc radio) Value(value string) radio {
	return rc.set("value", value)
}

// Visible 是否显示
func (rc radio) Visible(value bool) radio {
	return rc.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (rc radio) VisibleOn(value string) radio {
	return rc.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (rc radio) Width(value string) radio {
	return rc.set("width", value)
}
