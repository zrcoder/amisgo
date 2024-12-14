package comp

// uuid UUID 功能性组件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/uuid

type uuid schema

func UUID() uuid {
	return uuid{}.set("type", "uuid")
}

func (uc uuid) set(key string, value any) uuid {
	uc[key] = value
	return uc
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (uc uuid) AutoFill(value string) uuid {
	return uc.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (uc uuid) ClassName(value string) uuid {
	return uc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (uc uuid) ClearValueOnHidden(value bool) uuid {
	return uc.set("clearValueOnHidden", value)
}

// Desc
func (uc uuid) Desc(value string) uuid {
	return uc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (uc uuid) Description(value string) uuid {
	return uc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (uc uuid) DescriptionClassName(value string) uuid {
	return uc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (uc uuid) Disabled(value bool) uuid {
	return uc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (uc uuid) DisabledOn(value string) uuid {
	return uc.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (uc uuid) EditorSetting(value string) uuid {
	return uc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (uc uuid) ExtraName(value string) uuid {
	return uc.set("extraName", value)
}

// Hidden 是否隐藏
func (uc uuid) Hidden(value bool) uuid {
	return uc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (uc uuid) HiddenOn(value string) uuid {
	return uc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (uc uuid) Hint(value string) uuid {
	return uc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (uc uuid) Horizontal(value string) uuid {
	return uc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (uc uuid) ID(value string) uuid {
	return uc.set("id", value)
}

// InitAutoFill
func (uc uuid) InitAutoFill(value string) uuid {
	return uc.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (uc uuid) Inline(value bool) uuid {
	return uc.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (uc uuid) InputClassName(value string) uuid {
	return uc.set("inputClassName", value)
}

// Label 描述标题
func (uc uuid) Label(value string) uuid {
	return uc.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (uc uuid) LabelAlign(value string) uuid {
	return uc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (uc uuid) LabelClassName(value string) uuid {
	return uc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (uc uuid) LabelRemark(value string) uuid {
	return uc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (uc uuid) LabelWidth(value string) uuid {
	return uc.set("labelWidth", value)
}

// Length 长度，默认 uuid 的长度是 36，如果不需要那么长，可以设置这个来缩短
func (uc uuid) Length(value string) uuid {
	return uc.set("length", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (uc uuid) Mode(value string) uuid {
	return uc.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (uc uuid) Name(value string) uuid {
	return uc.set("name", value)
}

// OnEvent 事件动作配置
func (uc uuid) OnEvent(value any) uuid {
	return uc.set("onEvent", value)
}

// Placeholder 占位符
func (uc uuid) Placeholder(value string) uuid {
	return uc.set("placeholder", value)
}

// ReadOnly 是否只读
func (uc uuid) ReadOnly(value bool) uuid {
	return uc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (uc uuid) ReadOnlyOn(value string) uuid {
	return uc.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (uc uuid) Remark(value string) uuid {
	return uc.set("remark", value)
}

// Required 是否为必填
func (uc uuid) Required(value bool) uuid {
	return uc.set("required", value)
}

// Row
func (uc uuid) Row(value string) uuid {
	return uc.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (uc uuid) SaveImmediately(value bool) uuid {
	return uc.set("saveImmediately", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (uc uuid) Size(value string) uuid {
	return uc.set("size", value)
}

// Static 是否静态展示
func (uc uuid) Static(value bool) uuid {
	return uc.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (uc uuid) StaticClassName(value string) uuid {
	return uc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (uc uuid) StaticInputClassName(value string) uuid {
	return uc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (uc uuid) StaticLabelClassName(value string) uuid {
	return uc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (uc uuid) StaticOn(value string) uuid {
	return uc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (uc uuid) StaticPlaceholder(value string) uuid {
	return uc.set("staticPlaceholder", value)
}

// StaticSchema
func (uc uuid) StaticSchema(value string) uuid {
	return uc.set("staticSchema", value)
}

// Style 组件样式
func (uc uuid) Style(value any) uuid {
	return uc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (uc uuid) SubmitOnChange(value bool) uuid {
	return uc.set("submitOnChange", value)
}

// TestIdBuilder
func (uc uuid) TestIdBuilder(value string) uuid {
	return uc.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (uc uuid) UseMobileUI(value bool) uuid {
	return uc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (uc uuid) ValidateApi(value string) uuid {
	return uc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (uc uuid) ValidateOnChange(value bool) uuid {
	return uc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (uc uuid) ValidationErrors(value string) uuid {
	return uc.set("validationErrors", value)
}

// Validations
func (uc uuid) Validations(value string) uuid {
	return uc.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (uc uuid) Value(value string) uuid {
	return uc.set("value", value)
}

// Visible 是否显示
func (uc uuid) Visible(value bool) uuid {
	return uc.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (uc uuid) VisibleOn(value string) uuid {
	return uc.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (uc uuid) Width(value string) uuid {
	return uc.set("width", value)
}
