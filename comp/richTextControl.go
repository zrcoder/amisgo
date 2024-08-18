package comp

// RichTextControl 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-rich-text
//
// @version 6.7.0
type RichTextControl Schema

// NewRichTextControl 创建一个新的 RichTextControl 实例
func NewRichTextControl() RichTextControl {
	return RichTextControl{}.set("type", "input-rich-text")
}

func (rc RichTextControl) set(key string, value interface{}) RichTextControl {
	rc[key] = value
	return rc
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (rc RichTextControl) AutoFill(value string) RichTextControl {
	return rc.set("autoFill", value)
}

// BorderMode 边框模式，全边框，还是半边框，或者没边框。 可选值: full | half | none
func (rc RichTextControl) BorderMode(value string) RichTextControl {
	return rc.set("borderMode", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (rc RichTextControl) ClassName(value string) RichTextControl {
	return rc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (rc RichTextControl) ClearValueOnHidden(value bool) RichTextControl {
	return rc.set("clearValueOnHidden", value)
}

// Desc
func (rc RichTextControl) Desc(value string) RichTextControl {
	return rc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (rc RichTextControl) Description(value string) RichTextControl {
	return rc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (rc RichTextControl) DescriptionClassName(value string) RichTextControl {
	return rc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (rc RichTextControl) Disabled(value bool) RichTextControl {
	return rc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (rc RichTextControl) DisabledOn(value string) RichTextControl {
	return rc.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (rc RichTextControl) EditorSetting(value string) RichTextControl {
	return rc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (rc RichTextControl) ExtraName(value string) RichTextControl {
	return rc.set("extraName", value)
}

// FileField 接收器的字段名
func (rc RichTextControl) FileField(value string) RichTextControl {
	return rc.set("fileField", value)
}

// Hidden 是否隐藏
func (rc RichTextControl) Hidden(value bool) RichTextControl {
	return rc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (rc RichTextControl) HiddenOn(value string) RichTextControl {
	return rc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (rc RichTextControl) Hint(value string) RichTextControl {
	return rc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (rc RichTextControl) Horizontal(value string) RichTextControl {
	return rc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (rc RichTextControl) Id(value string) RichTextControl {
	return rc.set("id", value)
}

// InitAutoFill
func (rc RichTextControl) InitAutoFill(value string) RichTextControl {
	return rc.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (rc RichTextControl) Inline(value bool) RichTextControl {
	return rc.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (rc RichTextControl) InputClassName(value string) RichTextControl {
	return rc.set("inputClassName", value)
}

// Label 描述标题
func (rc RichTextControl) Label(value string) RichTextControl {
	return rc.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (rc RichTextControl) LabelAlign(value string) RichTextControl {
	return rc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (rc RichTextControl) LabelClassName(value string) RichTextControl {
	return rc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (rc RichTextControl) LabelRemark(value string) RichTextControl {
	return rc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (rc RichTextControl) LabelWidth(value string) RichTextControl {
	return rc.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (rc RichTextControl) Mode(value string) RichTextControl {
	return rc.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (rc RichTextControl) Name(value string) RichTextControl {
	return rc.set("name", value)
}

// OnEvent 事件动作配置
func (rc RichTextControl) OnEvent(value string) RichTextControl {
	return rc.set("onEvent", value)
}

// Options tinymce 或 froala 的配置
func (rc RichTextControl) Options(value string) RichTextControl {
	return rc.set("options", value)
}

// Placeholder 占位符
func (rc RichTextControl) Placeholder(value string) RichTextControl {
	return rc.set("placeholder", value)
}

// ReadOnly 是否只读
func (rc RichTextControl) ReadOnly(value bool) RichTextControl {
	return rc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (rc RichTextControl) ReadOnlyOn(value string) RichTextControl {
	return rc.set("readOnlyOn", value)
}

// Receiver 图片保存 API (图片保存 API)
func (rc RichTextControl) Receiver(value string) RichTextControl {
	return rc.set("receiver", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (rc RichTextControl) Remark(value string) RichTextControl {
	return rc.set("remark", value)
}

// Required 是否为必填
func (rc RichTextControl) Required(value bool) RichTextControl {
	return rc.set("required", value)
}

// Row
func (rc RichTextControl) Row(value string) RichTextControl {
	return rc.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (rc RichTextControl) SaveImmediately(value bool) RichTextControl {
	return rc.set("saveImmediately", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (rc RichTextControl) Size(value string) RichTextControl {
	return rc.set("size", value)
}

// Static 是否静态展示
func (rc RichTextControl) Static(value bool) RichTextControl {
	return rc.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (rc RichTextControl) StaticClassName(value string) RichTextControl {
	return rc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (rc RichTextControl) StaticInputClassName(value string) RichTextControl {
	return rc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: {"red": "data.progress > 80", "blue": "data.progress > 60"})
func (rc RichTextControl) StaticLabelClassName(value string) RichTextControl {
	return rc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (rc RichTextControl) StaticOn(value string) RichTextControl {
	return rc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (rc RichTextControl) StaticPlaceholder(value string) RichTextControl {
	return rc.set("staticPlaceholder", value)
}

// StaticSchema
func (rc RichTextControl) StaticSchema(value string) RichTextControl {
	return rc.set("staticSchema", value)
}

// Style 组件样式
func (rc RichTextControl) Style(value string) RichTextControl {
	return rc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (rc RichTextControl) SubmitOnChange(value bool) RichTextControl {
	return rc.set("submitOnChange", value)
}

// TestIdBuilder
func (rc RichTextControl) TestIdBuilder(value string) RichTextControl {
	return rc.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (rc RichTextControl) UseMobileUI(value bool) RichTextControl {
	return rc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (rc RichTextControl) ValidateApi(value string) RichTextControl {
	return rc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (rc RichTextControl) ValidateOnChange(value bool) RichTextControl {
	return rc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (rc RichTextControl) ValidationErrors(value string) RichTextControl {
	return rc.set("validationErrors", value)
}

// Validations
func (rc RichTextControl) Validations(value string) RichTextControl {
	return rc.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (rc RichTextControl) Value(value string) RichTextControl {
	return rc.set("value", value)
}

// Vendor 编辑器类型 可选值: froala | tinymce
func (rc RichTextControl) Vendor(value string) RichTextControl {
	return rc.set("vendor", value)
}

// VideoReceiver 视频保存 API (视频保存 API)
func (rc RichTextControl) VideoReceiver(value string) RichTextControl {
	return rc.set("videoReceiver", value)
}

// Visible 是否显示
func (rc RichTextControl) Visible(value bool) RichTextControl {
	return rc.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (rc RichTextControl) VisibleOn(value string) RichTextControl {
	return rc.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (rc RichTextControl) Width(value string) RichTextControl {
	return rc.set("width", value)
}
