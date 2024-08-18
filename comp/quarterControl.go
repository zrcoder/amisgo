package comp

// QuarterControl represents a quarter selection control
//
// @version 6.7.0
type QuarterControl Schema

// NewQuarterControl 创建一个新的 QuarterControl 实例
func NewQuarterControl() QuarterControl {
	return QuarterControl{}.set("type", "input-quarter")
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (q QuarterControl) AutoFill(value string) QuarterControl {
	return q.set("autoFill", value)
}

// BorderMode 边框模式，全边框，还是半边框，或者没边框。 可选值: full | half | none
func (q QuarterControl) BorderMode(value string) QuarterControl {
	return q.set("borderMode", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (q QuarterControl) ClassName(value string) QuarterControl {
	return q.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (q QuarterControl) ClearValueOnHidden(value bool) QuarterControl {
	return q.set("clearValueOnHidden", value)
}

// Clearable 是否显示清除按钮
func (q QuarterControl) Clearable(value bool) QuarterControl {
	return q.set("clearable", value)
}

// Desc
func (q QuarterControl) Desc(value string) QuarterControl {
	return q.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (q QuarterControl) Description(value string) QuarterControl {
	return q.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (q QuarterControl) DescriptionClassName(value string) QuarterControl {
	return q.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (q QuarterControl) Disabled(value bool) QuarterControl {
	return q.set("disabled", value)
}

// DisabledDate 字符串函数，用来决定是否禁用某个日期。(currentDate: moment.Moment, props: any) => boolean;
func (q QuarterControl) DisabledDate(value string) QuarterControl {
	return q.set("disabledDate", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (q QuarterControl) DisabledOn(value string) QuarterControl {
	return q.set("disabledOn", value)
}

// DisplayFormat 日期展示格式(新：替代inputFormat)
func (q QuarterControl) DisplayFormat(value string) QuarterControl {
	return q.set("displayFormat", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (q QuarterControl) EditorSetting(value string) QuarterControl {
	return q.set("editorSetting", value)
}

// Emebed 是否为内联模式？
func (q QuarterControl) Emebed(value bool) QuarterControl {
	return q.set("emebed", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (q QuarterControl) ExtraName(value string) QuarterControl {
	return q.set("extraName", value)
}

// Format 月份存储格式
func (q QuarterControl) Format(value string) QuarterControl {
	return q.set("format", value)
}

// Hidden 是否隐藏
func (q QuarterControl) Hidden(value bool) QuarterControl {
	return q.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (q QuarterControl) HiddenOn(value string) QuarterControl {
	return q.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (q QuarterControl) Hint(value string) QuarterControl {
	return q.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (q QuarterControl) Horizontal(value string) QuarterControl {
	return q.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (q QuarterControl) Id(value string) QuarterControl {
	return q.set("id", value)
}

// InitAutoFill
func (q QuarterControl) InitAutoFill(value string) QuarterControl {
	return q.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (q QuarterControl) Inline(value bool) QuarterControl {
	return q.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (q QuarterControl) InputClassName(value string) QuarterControl {
	return q.set("inputClassName", value)
}

// InputFormat 月份展示格式
func (q QuarterControl) InputFormat(value string) QuarterControl {
	return q.set("inputFormat", value)
}

// Label 描述标题
func (q QuarterControl) Label(value string) QuarterControl {
	return q.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (q QuarterControl) LabelAlign(value string) QuarterControl {
	return q.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (q QuarterControl) LabelClassName(value string) QuarterControl {
	return q.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (q QuarterControl) LabelRemark(value string) QuarterControl {
	return q.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (q QuarterControl) LabelWidth(value string) QuarterControl {
	return q.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (q QuarterControl) Mode(value string) QuarterControl {
	return q.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (q QuarterControl) Name(value string) QuarterControl {
	return q.set("name", value)
}

// OnEvent 事件动作配置
func (q QuarterControl) OnEvent(value string) QuarterControl {
	return q.set("onEvent", value)
}

// Placeholder 占位符
func (q QuarterControl) Placeholder(value string) QuarterControl {
	return q.set("placeholder", value)
}

// ReadOnly 是否只读
func (q QuarterControl) ReadOnly(value bool) QuarterControl {
	return q.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (q QuarterControl) ReadOnlyOn(value string) QuarterControl {
	return q.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (q QuarterControl) Remark(value string) QuarterControl {
	return q.set("remark", value)
}

// Required 是否为必填
func (q QuarterControl) Required(value bool) QuarterControl {
	return q.set("required", value)
}

// Row
func (q QuarterControl) Row(value string) QuarterControl {
	return q.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (q QuarterControl) SaveImmediately(value bool) QuarterControl {
	return q.set("saveImmediately", value)
}

// Shortcuts 日期快捷键
func (q QuarterControl) Shortcuts(value string) QuarterControl {
	return q.set("shortcuts", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (q QuarterControl) Size(value string) QuarterControl {
	return q.set("size", value)
}

// Static 是否静态展示
func (q QuarterControl) Static(value bool) QuarterControl {
	return q.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (q QuarterControl) StaticClassName(value string) QuarterControl {
	return q.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (q QuarterControl) StaticInputClassName(value string) QuarterControl {
	return q.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (q QuarterControl) StaticLabelClassName(value string) QuarterControl {
	return q.set("staticLabelClassName", value)
}

// Template
func (q QuarterControl) Template(value string) QuarterControl {
	return q.set("template", value)
}

// TimeFormat 时间展示格式
func (q QuarterControl) TimeFormat(value string) QuarterControl {
	return q.set("timeFormat", value)
}

// Validation 错误提示
func (q QuarterControl) Validation(value string) QuarterControl {
	return q.set("validation", value)
}

// Value 默认值
func (q QuarterControl) Value(value string) QuarterControl {
	return q.set("value", value)
}

// VerticalAlign 控件竖直对齐方式，可选值: top | middle | bottom
func (q QuarterControl) VerticalAlign(value string) QuarterControl {
	return q.set("verticalAlign", value)
}

// VisibleOn 是否可见表达式 (表达式，语法 `data.xxx > 5`。)
func (q QuarterControl) VisibleOn(value string) QuarterControl {
	return q.set("visibleOn", value)
}

// Width
func (q QuarterControl) Width(value string) QuarterControl {
	return q.set("width", value)
}

// set 设置属性值
func (q QuarterControl) set(key string, value interface{}) QuarterControl {
	// This method should set the property `key` with the given `value` on the `QuarterControl` instance.
	// Assuming we have a way to store and manage these properties.
	return q
}
