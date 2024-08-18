package comp

// InputColorControl 颜色选择框。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/color
type InputColorControl Schema

// NewInputColorControl 创建一个新的 InputColorControl 实例，并设置默认的 type
func NewInputColorControl() InputColorControl {
	return make(InputColorControl).set("type", "input-color")
}

func (i InputColorControl) set(key string, value interface{}) InputColorControl {
	i[key] = value
	return i
}

// AllowCustomColor 是否允许用户输入颜色。
func (i InputColorControl) AllowCustomColor(value bool) InputColorControl {
	return i.set("allowCustomColor", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (i InputColorControl) AutoFill(value string) InputColorControl {
	return i.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i InputColorControl) ClassName(value string) InputColorControl {
	return i.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (i InputColorControl) ClearValueOnHidden(value bool) InputColorControl {
	return i.set("clearValueOnHidden", value)
}

// Clearable 是否显示清除按钮
func (i InputColorControl) Clearable(value bool) InputColorControl {
	return i.set("clearable", value)
}

// CloseOnSelect 选中颜色后是否关闭弹出层。
func (i InputColorControl) CloseOnSelect(value bool) InputColorControl {
	return i.set("closeOnSelect", value)
}

// Desc
func (i InputColorControl) Desc(value string) InputColorControl {
	return i.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (i InputColorControl) Description(value string) InputColorControl {
	return i.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (i InputColorControl) DescriptionClassName(value string) InputColorControl {
	return i.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (i InputColorControl) Disabled(value bool) InputColorControl {
	return i.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (i InputColorControl) DisabledOn(value string) InputColorControl {
	return i.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (i InputColorControl) EditorSetting(value string) InputColorControl {
	return i.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (i InputColorControl) ExtraName(value string) InputColorControl {
	return i.set("extraName", value)
}

// Format 颜色格式 可选值: hex | rgb | rgba | hsl
func (i InputColorControl) Format(value string) InputColorControl {
	return i.set("format", value)
}

// Hidden 是否隐藏
func (i InputColorControl) Hidden(value bool) InputColorControl {
	return i.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (i InputColorControl) HiddenOn(value string) InputColorControl {
	return i.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (i InputColorControl) Hint(value string) InputColorControl {
	return i.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (i InputColorControl) Horizontal(value string) InputColorControl {
	return i.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (i InputColorControl) Id(value string) InputColorControl {
	return i.set("id", value)
}

// InitAutoFill
func (i InputColorControl) InitAutoFill(value string) InputColorControl {
	return i.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (i InputColorControl) Inline(value bool) InputColorControl {
	return i.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (i InputColorControl) InputClassName(value string) InputColorControl {
	return i.set("inputClassName", value)
}

// Label 描述标题
func (i InputColorControl) Label(value string) InputColorControl {
	return i.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (i InputColorControl) LabelAlign(value string) InputColorControl {
	return i.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (i InputColorControl) LabelClassName(value string) InputColorControl {
	return i.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (i InputColorControl) LabelRemark(value string) InputColorControl {
	return i.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (i InputColorControl) LabelWidth(value string) InputColorControl {
	return i.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (i InputColorControl) Mode(value string) InputColorControl {
	return i.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (i InputColorControl) Name(value string) InputColorControl {
	return i.set("name", value)
}

// OnEvent 事件动作配置
func (i InputColorControl) OnEvent(value string) InputColorControl {
	return i.set("onEvent", value)
}

// Placeholder 占位符
func (i InputColorControl) Placeholder(value string) InputColorControl {
	return i.set("placeholder", value)
}

// PopOverContainerSelector 弹窗容器选择器
func (i InputColorControl) PopOverContainerSelector(value string) InputColorControl {
	return i.set("popOverContainerSelector", value)
}

// PresetColors 预设颜色，用户可以直接从预设中选。
func (i InputColorControl) PresetColors(value string) InputColorControl {
	return i.set("presetColors", value)
}

// ReadOnly 是否只读
func (i InputColorControl) ReadOnly(value bool) InputColorControl {
	return i.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (i InputColorControl) ReadOnlyOn(value string) InputColorControl {
	return i.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (i InputColorControl) Remark(value string) InputColorControl {
	return i.set("remark", value)
}

// Required 是否为必填
func (i InputColorControl) Required(value bool) InputColorControl {
	return i.set("required", value)
}

// Row
func (i InputColorControl) Row(value string) InputColorControl {
	return i.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (i InputColorControl) SaveImmediately(value bool) InputColorControl {
	return i.set("saveImmediately", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (i InputColorControl) Size(value string) InputColorControl {
	return i.set("size", value)
}

// Static 是否静态展示
func (i InputColorControl) Static(value bool) InputColorControl {
	return i.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i InputColorControl) StaticClassName(value string) InputColorControl {
	return i.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i InputColorControl) StaticInputClassName(value string) InputColorControl {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (i InputColorControl) StaticLabelClassName(value string) InputColorControl {
	return i.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (i InputColorControl) StaticOn(value string) InputColorControl {
	return i.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (i InputColorControl) StaticPlaceholder(value string) InputColorControl {
	return i.set("staticPlaceholder", value)
}

// StaticSchema
func (i InputColorControl) StaticSchema(value string) InputColorControl {
	return i.set("staticSchema", value)
}

// Style 组件样式
func (i InputColorControl) Style(value string) InputColorControl {
	return i.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (i InputColorControl) SubmitOnChange(value bool) InputColorControl {
	return i.set("submitOnChange", value)
}

// TestIdBuilder
func (i InputColorControl) TestIdBuilder(value string) InputColorControl {
	return i.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (i InputColorControl) UseMobileUI(value bool) InputColorControl {
	return i.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (i InputColorControl) ValidateApi(value string) InputColorControl {
	return i.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (i InputColorControl) ValidateOnChange(value bool) InputColorControl {
	return i.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (i InputColorControl) ValidationErrors(value string) InputColorControl {
	return i.set("validationErrors", value)
}

// Validations
func (i InputColorControl) Validations(value string) InputColorControl {
	return i.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (i InputColorControl) Value(value string) InputColorControl {
	return i.set("value", value)
}

// Visible 是否显示
func (i InputColorControl) Visible(value bool) InputColorControl {
	return i.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (i InputColorControl) VisibleOn(value string) InputColorControl {
	return i.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (i InputColorControl) Width(value string) InputColorControl {
	return i.set("width", value)
}
