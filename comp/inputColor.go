package comp

// inputColor 颜色选择框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-color
type inputColor schema

func InputColor() inputColor {
	return make(inputColor).set("type", "input-color")
}

func (i inputColor) set(key string, value interface{}) inputColor {
	i[key] = value
	return i
}

// AllowCustomColor 是否允许用户输入颜色。
func (i inputColor) AllowCustomColor(value bool) inputColor {
	return i.set("allowCustomColor", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (i inputColor) AutoFill(value string) inputColor {
	return i.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i inputColor) ClassName(value string) inputColor {
	return i.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (i inputColor) ClearValueOnHidden(value bool) inputColor {
	return i.set("clearValueOnHidden", value)
}

// Clearable 是否显示清除按钮
func (i inputColor) Clearable(value bool) inputColor {
	return i.set("clearable", value)
}

// CloseOnSelect 选中颜色后是否关闭弹出层。
func (i inputColor) CloseOnSelect(value bool) inputColor {
	return i.set("closeOnSelect", value)
}

// Desc
func (i inputColor) Desc(value string) inputColor {
	return i.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (i inputColor) Description(value string) inputColor {
	return i.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (i inputColor) DescriptionClassName(value string) inputColor {
	return i.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (i inputColor) Disabled(value bool) inputColor {
	return i.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (i inputColor) DisabledOn(value string) inputColor {
	return i.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (i inputColor) EditorSetting(value string) inputColor {
	return i.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (i inputColor) ExtraName(value string) inputColor {
	return i.set("extraName", value)
}

// Format 颜色格式 可选值: hex | rgb | rgba | hsl
func (i inputColor) Format(value string) inputColor {
	return i.set("format", value)
}

// Hidden 是否隐藏
func (i inputColor) Hidden(value bool) inputColor {
	return i.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (i inputColor) HiddenOn(value string) inputColor {
	return i.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (i inputColor) Hint(value string) inputColor {
	return i.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (i inputColor) Horizontal(value string) inputColor {
	return i.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (i inputColor) Id(value string) inputColor {
	return i.set("id", value)
}

// InitAutoFill
func (i inputColor) InitAutoFill(value string) inputColor {
	return i.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (i inputColor) Inline(value bool) inputColor {
	return i.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (i inputColor) InputClassName(value string) inputColor {
	return i.set("inputClassName", value)
}

// Label 描述标题
func (i inputColor) Label(value string) inputColor {
	return i.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (i inputColor) LabelAlign(value string) inputColor {
	return i.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (i inputColor) LabelClassName(value string) inputColor {
	return i.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (i inputColor) LabelRemark(value string) inputColor {
	return i.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (i inputColor) LabelWidth(value string) inputColor {
	return i.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (i inputColor) Mode(value string) inputColor {
	return i.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (i inputColor) Name(value string) inputColor {
	return i.set("name", value)
}

// OnEvent 事件动作配置
func (i inputColor) OnEvent(value string) inputColor {
	return i.set("onEvent", value)
}

// Placeholder 占位符
func (i inputColor) Placeholder(value string) inputColor {
	return i.set("placeholder", value)
}

// PopOverContainerSelector 弹窗容器选择器
func (i inputColor) PopOverContainerSelector(value string) inputColor {
	return i.set("popOverContainerSelector", value)
}

// PresetColors 预设颜色，用户可以直接从预设中选。
func (i inputColor) PresetColors(value string) inputColor {
	return i.set("presetColors", value)
}

// ReadOnly 是否只读
func (i inputColor) ReadOnly(value bool) inputColor {
	return i.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (i inputColor) ReadOnlyOn(value string) inputColor {
	return i.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (i inputColor) Remark(value string) inputColor {
	return i.set("remark", value)
}

// Required 是否为必填
func (i inputColor) Required(value bool) inputColor {
	return i.set("required", value)
}

// Row
func (i inputColor) Row(value string) inputColor {
	return i.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (i inputColor) SaveImmediately(value bool) inputColor {
	return i.set("saveImmediately", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (i inputColor) Size(value string) inputColor {
	return i.set("size", value)
}

// Static 是否静态展示
func (i inputColor) Static(value bool) inputColor {
	return i.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i inputColor) StaticClassName(value string) inputColor {
	return i.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i inputColor) StaticInputClassName(value string) inputColor {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i inputColor) StaticLabelClassName(value string) inputColor {
	return i.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (i inputColor) StaticOn(value string) inputColor {
	return i.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (i inputColor) StaticPlaceholder(value string) inputColor {
	return i.set("staticPlaceholder", value)
}

// StaticSchema
func (i inputColor) StaticSchema(value string) inputColor {
	return i.set("staticSchema", value)
}

// Style 组件样式
func (i inputColor) Style(value string) inputColor {
	return i.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (i inputColor) SubmitOnChange(value bool) inputColor {
	return i.set("submitOnChange", value)
}

// TestIdBuilder
func (i inputColor) TestIdBuilder(value string) inputColor {
	return i.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (i inputColor) UseMobileUI(value bool) inputColor {
	return i.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (i inputColor) ValidateApi(value string) inputColor {
	return i.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (i inputColor) ValidateOnChange(value bool) inputColor {
	return i.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (i inputColor) ValidationErrors(value string) inputColor {
	return i.set("validationErrors", value)
}

// Validations
func (i inputColor) Validations(value string) inputColor {
	return i.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (i inputColor) Value(value string) inputColor {
	return i.set("value", value)
}

// Visible 是否显示
func (i inputColor) Visible(value bool) inputColor {
	return i.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (i inputColor) VisibleOn(value string) inputColor {
	return i.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (i inputColor) Width(value string) inputColor {
	return i.set("width", value)
}
