package comp

// TimeControl 时间选择控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/time
//
// @author slowlyo
// @version 6.7.0
type TimeControl Schema

// NewTimeControl 创建一个新的 TimeControl 实例
func NewTimeControl() TimeControl {
	return TimeControl{}.set("type", "input-time")
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (tc TimeControl) AutoFill(value string) TimeControl {
	return tc.set("autoFill", value)
}

// BorderMode 边框模式，全边框，还是半边框，或者没边框。可选值: full | half | none
func (tc TimeControl) BorderMode(value string) TimeControl {
	return tc.set("borderMode", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (tc TimeControl) ClassName(value string) TimeControl {
	return tc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (tc TimeControl) ClearValueOnHidden(value bool) TimeControl {
	return tc.set("clearValueOnHidden", value)
}

// Clearable 是否显示清除按钮
func (tc TimeControl) Clearable(value bool) TimeControl {
	return tc.set("clearable", value)
}

// Desc
func (tc TimeControl) Desc(value string) TimeControl {
	return tc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (tc TimeControl) Description(value string) TimeControl {
	return tc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (tc TimeControl) DescriptionClassName(value string) TimeControl {
	return tc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (tc TimeControl) Disabled(value bool) TimeControl {
	return tc.set("disabled", value)
}

// DisabledDate 字符串函数，用来决定是否禁用某个日期。(currentDate: moment.Moment, props: any) => boolean;
func (tc TimeControl) DisabledDate(value string) TimeControl {
	return tc.set("disabledDate", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (tc TimeControl) DisabledOn(value string) TimeControl {
	return tc.set("disabledOn", value)
}

// DisplayFormat 日期展示格式(新：替代inputFormat)
func (tc TimeControl) DisplayFormat(value string) TimeControl {
	return tc.set("displayFormat", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (tc TimeControl) EditorSetting(value string) TimeControl {
	return tc.set("editorSetting", value)
}

// Emebed 是否为内联模式？
func (tc TimeControl) Emebed(value bool) TimeControl {
	return tc.set("emebed", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (tc TimeControl) ExtraName(value string) TimeControl {
	return tc.set("extraName", value)
}

// Format 日期存储格式
func (tc TimeControl) Format(value string) TimeControl {
	return tc.set("format", value)
}

// Hidden 是否隐藏
func (tc TimeControl) Hidden(value bool) TimeControl {
	return tc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (tc TimeControl) HiddenOn(value string) TimeControl {
	return tc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (tc TimeControl) Hint(value string) TimeControl {
	return tc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (tc TimeControl) Horizontal(value string) TimeControl {
	return tc.set("horizontal", value)
}

// ID 组件唯一 id，主要用于日志采集
func (tc TimeControl) ID(value string) TimeControl {
	return tc.set("id", value)
}

// InitAutoFill
func (tc TimeControl) InitAutoFill(value string) TimeControl {
	return tc.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (tc TimeControl) Inline(value bool) TimeControl {
	return tc.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (tc TimeControl) InputClassName(value string) TimeControl {
	return tc.set("inputClassName", value)
}

// InputFormat 日期展示格式
func (tc TimeControl) InputFormat(value string) TimeControl {
	return tc.set("inputFormat", value)
}

// Label 描述标题
func (tc TimeControl) Label(value string) TimeControl {
	return tc.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (tc TimeControl) LabelAlign(value string) TimeControl {
	return tc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (tc TimeControl) LabelClassName(value string) TimeControl {
	return tc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (tc TimeControl) LabelRemark(value string) TimeControl {
	return tc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (tc TimeControl) LabelWidth(value string) TimeControl {
	return tc.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (tc TimeControl) Mode(value string) TimeControl {
	return tc.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (tc TimeControl) Name(value string) TimeControl {
	return tc.set("name", value)
}

// OnEvent 事件动作配置
func (tc TimeControl) OnEvent(value string) TimeControl {
	return tc.set("onEvent", value)
}

// Placeholder 占位符
func (tc TimeControl) Placeholder(value string) TimeControl {
	return tc.set("placeholder", value)
}

// ReadOnly 是否只读
func (tc TimeControl) ReadOnly(value bool) TimeControl {
	return tc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (tc TimeControl) ReadOnlyOn(value string) TimeControl {
	return tc.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (tc TimeControl) Remark(value string) TimeControl {
	return tc.set("remark", value)
}

// Required 是否为必填
func (tc TimeControl) Required(value bool) TimeControl {
	return tc.set("required", value)
}

// Row
func (tc TimeControl) Row(value string) TimeControl {
	return tc.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (tc TimeControl) SaveImmediately(value bool) TimeControl {
	return tc.set("saveImmediately", value)
}

// Shortcuts 日期快捷键
func (tc TimeControl) Shortcuts(value string) TimeControl {
	return tc.set("shortcuts", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (tc TimeControl) Size(value string) TimeControl {
	return tc.set("size", value)
}

// Static 是否静态展示
func (tc TimeControl) Static(value bool) TimeControl {
	return tc.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (tc TimeControl) StaticClassName(value string) TimeControl {
	return tc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (tc TimeControl) StaticInputClassName(value string) TimeControl {
	return tc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (tc TimeControl) StaticLabelClassName(value string) TimeControl {
	return tc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (tc TimeControl) StaticOn(value string) TimeControl {
	return tc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (tc TimeControl) StaticPlaceholder(value string) TimeControl {
	return tc.set("staticPlaceholder", value)
}

// StaticSchema
func (tc TimeControl) StaticSchema(value string) TimeControl {
	return tc.set("staticSchema", value)
}

// Style 组件样式
func (tc TimeControl) Style(value string) TimeControl {
	return tc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (tc TimeControl) SubmitOnChange(value bool) TimeControl {
	return tc.set("submitOnChange", value)
}

// TestIdBuilder
func (tc TimeControl) TestIdBuilder(value string) TimeControl {
	return tc.set("testIdBuilder", value)
}

// TimeConstraints 时间输入范围限制
func (tc TimeControl) TimeConstraints(value string) TimeControl {
	return tc.set("timeConstraints", value)
}

// TimeFormat 时间的格式。
func (tc TimeControl) TimeFormat(value string) TimeControl {
	return tc.set("timeFormat", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (tc TimeControl) UseMobileUI(value bool) TimeControl {
	return tc.set("useMobileUI", value)
}

// UTC 设定是否存储 utc 时间。
func (tc TimeControl) UTC(value bool) TimeControl {
	return tc.set("utc", value)
}

// ValidateApi 远端校验表单项接口
func (tc TimeControl) ValidateApi(value string) TimeControl {
	return tc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (tc TimeControl) ValidateOnChange(value bool) TimeControl {
	return tc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (tc TimeControl) ValidationErrors(value string) TimeControl {
	return tc.set("validationErrors", value)
}

// Validations
func (tc TimeControl) Validations(value string) TimeControl {
	return tc.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (tc TimeControl) Value(value string) TimeControl {
	return tc.set("value", value)
}

// ValueFormat 替代format
func (tc TimeControl) ValueFormat(value string) TimeControl {
	return tc.set("valueFormat", value)
}

// Visible 是否显示
func (tc TimeControl) Visible(value bool) TimeControl {
	return tc.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (tc TimeControl) VisibleOn(value string) TimeControl {
	return tc.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (tc TimeControl) Width(value string) TimeControl {
	return tc.set("width", value)
}

// set 设置键值对（私有方法）
func (tc TimeControl) set(key string, value interface{}) TimeControl {
	// 这里假设 Schema 类型有一个 set 方法可以设置键值对
	tc[key] = value
	return tc
}
