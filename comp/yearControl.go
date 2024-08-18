package comp

// YearControl 年份选择控件
//
// @author  slowlyo
// @version 6.7.0
type YearControl Schema

// NewYearControl 创建一个新的 YearControl 实例
func NewYearControl() YearControl {
	return YearControl{}.set("type", "input-year")
}

func (yc YearControl) set(key string, value interface{}) YearControl {
	yc[key] = value
	return yc
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (yc YearControl) AutoFill(value string) YearControl {
	return yc.set("autoFill", value)
}

// BorderMode 边框模式，全边框，还是半边框，或者没边框。 可选值: full | half | none
func (yc YearControl) BorderMode(value string) YearControl {
	return yc.set("borderMode", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (yc YearControl) ClassName(value string) YearControl {
	return yc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (yc YearControl) ClearValueOnHidden(value bool) YearControl {
	return yc.set("clearValueOnHidden", value)
}

// Clearable 是否显示清除按钮
func (yc YearControl) Clearable(value bool) YearControl {
	return yc.set("clearable", value)
}

// Desc
func (yc YearControl) Desc(value string) YearControl {
	return yc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (yc YearControl) Description(value string) YearControl {
	return yc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (yc YearControl) DescriptionClassName(value string) YearControl {
	return yc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (yc YearControl) Disabled(value bool) YearControl {
	return yc.set("disabled", value)
}

// DisabledDate 字符串函数，用来决定是否禁用某个日期。(currentDate: moment.Moment, props: any) => boolean;
func (yc YearControl) DisabledDate(value string) YearControl {
	return yc.set("disabledDate", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (yc YearControl) DisabledOn(value string) YearControl {
	return yc.set("disabledOn", value)
}

// DisplayFormat 日期展示格式(新：替代inputFormat)
func (yc YearControl) DisplayFormat(value string) YearControl {
	return yc.set("displayFormat", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (yc YearControl) EditorSetting(value string) YearControl {
	return yc.set("editorSetting", value)
}

// Emebed 是否为内联模式？
func (yc YearControl) Emebed(value bool) YearControl {
	return yc.set("emebed", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (yc YearControl) ExtraName(value string) YearControl {
	return yc.set("extraName", value)
}

// Format 月份存储格式
func (yc YearControl) Format(value string) YearControl {
	return yc.set("format", value)
}

// Hidden 是否隐藏
func (yc YearControl) Hidden(value bool) YearControl {
	return yc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (yc YearControl) HiddenOn(value string) YearControl {
	return yc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (yc YearControl) Hint(value string) YearControl {
	return yc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (yc YearControl) Horizontal(value string) YearControl {
	return yc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (yc YearControl) Id(value string) YearControl {
	return yc.set("id", value)
}

// InitAutoFill
func (yc YearControl) InitAutoFill(value string) YearControl {
	return yc.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (yc YearControl) Inline(value bool) YearControl {
	return yc.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (yc YearControl) InputClassName(value string) YearControl {
	return yc.set("inputClassName", value)
}

// InputFormat 月份展示格式
func (yc YearControl) InputFormat(value string) YearControl {
	return yc.set("inputFormat", value)
}

// Label 描述标题
func (yc YearControl) Label(value string) YearControl {
	return yc.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (yc YearControl) LabelAlign(value string) YearControl {
	return yc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (yc YearControl) LabelClassName(value string) YearControl {
	return yc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (yc YearControl) LabelRemark(value string) YearControl {
	return yc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (yc YearControl) LabelWidth(value string) YearControl {
	return yc.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (yc YearControl) Mode(value string) YearControl {
	return yc.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (yc YearControl) Name(value string) YearControl {
	return yc.set("name", value)
}

// OnEvent 事件动作配置
func (yc YearControl) OnEvent(value string) YearControl {
	return yc.set("onEvent", value)
}

// Placeholder 占位符
func (yc YearControl) Placeholder(value string) YearControl {
	return yc.set("placeholder", value)
}

// ReadOnly 是否只读
func (yc YearControl) ReadOnly(value bool) YearControl {
	return yc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (yc YearControl) ReadOnlyOn(value string) YearControl {
	return yc.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (yc YearControl) Remark(value string) YearControl {
	return yc.set("remark", value)
}

// Required 是否为必填
func (yc YearControl) Required(value bool) YearControl {
	return yc.set("required", value)
}

// Row
func (yc YearControl) Row(value string) YearControl {
	return yc.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (yc YearControl) SaveImmediately(value bool) YearControl {
	return yc.set("saveImmediately", value)
}

// Shortcuts 日期快捷键
func (yc YearControl) Shortcuts(value string) YearControl {
	return yc.set("shortcuts", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (yc YearControl) Size(value string) YearControl {
	return yc.set("size", value)
}

// Static 是否静态展示
func (yc YearControl) Static(value bool) YearControl {
	return yc.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (yc YearControl) StaticClassName(value string) YearControl {
	return yc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (yc YearControl) StaticInputClassName(value string) YearControl {
	return yc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (yc YearControl) StaticLabelClassName(value string) YearControl {
	return yc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (yc YearControl) StaticOn(value string) YearControl {
	return yc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (yc YearControl) StaticPlaceholder(value string) YearControl {
	return yc.set("staticPlaceholder", value)
}

// StaticSchema
func (yc YearControl) StaticSchema(value string) YearControl {
	return yc.set("staticSchema", value)
}

// Style 组件样式
func (yc YearControl) Style(value string) YearControl {
	return yc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (yc YearControl) SubmitOnChange(value bool) YearControl {
	return yc.set("submitOnChange", value)
}

// TestIdBuilder
func (yc YearControl) TestIdBuilder(value string) YearControl {
	return yc.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (yc YearControl) UseMobileUI(value bool) YearControl {
	return yc.set("useMobileUI", value)
}

// Utc 设定是否存储 utc 时间。
func (yc YearControl) Utc(value bool) YearControl {
	return yc.set("utc", value)
}

// ValidateApi 远端校验表单项接口
func (yc YearControl) ValidateApi(value string) YearControl {
	return yc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (yc YearControl) ValidateOnChange(value bool) YearControl {
	return yc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (yc YearControl) ValidationErrors(value string) YearControl {
	return yc.set("validationErrors", value)
}

// Validations
func (yc YearControl) Validations(value string) YearControl {
	return yc.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (yc YearControl) Value(value string) YearControl {
	return yc.set("value", value)
}

// ValueFormat 替代format
func (yc YearControl) ValueFormat(value string) YearControl {
	return yc.set("valueFormat", value)
}

// Visible 是否显示
func (yc YearControl) Visible(value bool) YearControl {
	return yc.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (yc YearControl) VisibleOn(value string) YearControl {
	return yc.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (yc YearControl) Width(value string) YearControl {
	return yc.set("width", value)
}
