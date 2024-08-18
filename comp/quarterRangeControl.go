package comp

// QuarterRangeControl 季度范围控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-quarter-range
//
// @author  slowlyo
// @version 6.7.0
type QuarterRangeControl Schema

// NewQuarterRangeControl 创建一个新的 QuarterRangeControl 实例
func NewQuarterRangeControl() QuarterRangeControl {
	return QuarterRangeControl{}.set("type", "input-quarter-range")
}

// set 设置字段的值
func (qrc QuarterRangeControl) set(key string, value interface{}) QuarterRangeControl {
	qrc[key] = value
	return qrc
}

// Animation 是否启用游标动画，默认开启
func (qrc QuarterRangeControl) Animation(value bool) QuarterRangeControl {
	return qrc.set("animation", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (qrc QuarterRangeControl) AutoFill(value string) QuarterRangeControl {
	return qrc.set("autoFill", value)
}

// BorderMode 边框模式，全边框，还是半边框，或者没边框。 可选值: full | half | none
func (qrc QuarterRangeControl) BorderMode(value string) QuarterRangeControl {
	return qrc.set("borderMode", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (qrc QuarterRangeControl) ClassName(value string) QuarterRangeControl {
	return qrc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (qrc QuarterRangeControl) ClearValueOnHidden(value bool) QuarterRangeControl {
	return qrc.set("clearValueOnHidden", value)
}

// Delimiter 分割符, 因为有两个值，开始时间和结束时间，所以要有连接符。默认为英文逗号。
func (qrc QuarterRangeControl) Delimiter(value string) QuarterRangeControl {
	return qrc.set("delimiter", value)
}

// Desc
func (qrc QuarterRangeControl) Desc(value string) QuarterRangeControl {
	return qrc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (qrc QuarterRangeControl) Description(value string) QuarterRangeControl {
	return qrc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (qrc QuarterRangeControl) DescriptionClassName(value string) QuarterRangeControl {
	return qrc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (qrc QuarterRangeControl) Disabled(value bool) QuarterRangeControl {
	return qrc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (qrc QuarterRangeControl) DisabledOn(value string) QuarterRangeControl {
	return qrc.set("disabledOn", value)
}

// DisplayFormat 用来配置显示的时间格式（新：同inputFormat）
func (qrc QuarterRangeControl) DisplayFormat(value string) QuarterRangeControl {
	return qrc.set("displayFormat", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (qrc QuarterRangeControl) EditorSetting(value string) QuarterRangeControl {
	return qrc.set("editorSetting", value)
}

// Embed 开启后变成非弹出模式，即内联模式。
func (qrc QuarterRangeControl) Embed(value bool) QuarterRangeControl {
	return qrc.set("embed", value)
}

// EndPlaceholder 日期范围结束时间-占位符
func (qrc QuarterRangeControl) EndPlaceholder(value string) QuarterRangeControl {
	return qrc.set("endPlaceholder", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (qrc QuarterRangeControl) ExtraName(value string) QuarterRangeControl {
	return qrc.set("extraName", value)
}

// Format 默认 `X` 即时间戳格式，用来提交的时间格式。更多格式类型请参考 moment.
func (qrc QuarterRangeControl) Format(value string) QuarterRangeControl {
	return qrc.set("format", value)
}

// Hidden 是否隐藏
func (qrc QuarterRangeControl) Hidden(value bool) QuarterRangeControl {
	return qrc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (qrc QuarterRangeControl) HiddenOn(value string) QuarterRangeControl {
	return qrc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (qrc QuarterRangeControl) Hint(value string) QuarterRangeControl {
	return qrc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (qrc QuarterRangeControl) Horizontal(value string) QuarterRangeControl {
	return qrc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (qrc QuarterRangeControl) Id(value string) QuarterRangeControl {
	return qrc.set("id", value)
}

// InitAutoFill
func (qrc QuarterRangeControl) InitAutoFill(value string) QuarterRangeControl {
	return qrc.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (qrc QuarterRangeControl) Inline(value bool) QuarterRangeControl {
	return qrc.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (qrc QuarterRangeControl) InputClassName(value string) QuarterRangeControl {
	return qrc.set("inputClassName", value)
}

// InputFormat 默认 `YYYY-MM-DD` 用来配置显示的时间格式。
func (qrc QuarterRangeControl) InputFormat(value string) QuarterRangeControl {
	return qrc.set("inputFormat", value)
}

// JoinValues 开启后将选中的选项 value 的值用连接符拼接起来，作为当前表单项的值。如： `value1,value2` 否则为 `[value1, value2]`
func (qrc QuarterRangeControl) JoinValues(value bool) QuarterRangeControl {
	return qrc.set("joinValues", value)
}

// Label 描述标题
func (qrc QuarterRangeControl) Label(value string) QuarterRangeControl {
	return qrc.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (qrc QuarterRangeControl) LabelAlign(value string) QuarterRangeControl {
	return qrc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (qrc QuarterRangeControl) LabelClassName(value string) QuarterRangeControl {
	return qrc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (qrc QuarterRangeControl) LabelRemark(value string) QuarterRangeControl {
	return qrc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (qrc QuarterRangeControl) LabelWidth(value string) QuarterRangeControl {
	return qrc.set("labelWidth", value)
}

// MaxDate 最大日期限制，支持变量 $xxx 来取值，或者用相对值如：* `-2mins` 2分钟前\n * `+2days` 2天后\n* `-10week` 十周前\n可用单位： `min`、`hour`、`day`、`week`、`month`、`year`。所有单位支持复数形式。
func (qrc QuarterRangeControl) MaxDate(value string) QuarterRangeControl {
	return qrc.set("maxDate", value)
}

// MaxDuration 最大跨度，比如 2days
func (qrc QuarterRangeControl) MaxDuration(value string) QuarterRangeControl {
	return qrc.set("maxDuration", value)
}

// MinDate 最小日期限制，支持变量 $xxx 来取值，或者用相对值如：* `-2mins` 2分钟前\n * `+2days` 2天后\n* `-10week` 十周前\n可用单位： `min`、`hour`、`day`、`week`、`month`、`year`。所有单位支持复数形式。
func (qrc QuarterRangeControl) MinDate(value string) QuarterRangeControl {
	return qrc.set("minDate", value)
}

// MinDuration 最小跨度，比如 2days
func (qrc QuarterRangeControl) MinDuration(value string) QuarterRangeControl {
	return qrc.set("minDuration", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (qrc QuarterRangeControl) Mode(value string) QuarterRangeControl {
	return qrc.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (qrc QuarterRangeControl) Name(value string) QuarterRangeControl {
	return qrc.set("name", value)
}

// OnEvent 事件动作配置
func (qrc QuarterRangeControl) OnEvent(value string) QuarterRangeControl {
	return qrc.set("onEvent", value)
}

// Placeholder 占位符
func (qrc QuarterRangeControl) Placeholder(value string) QuarterRangeControl {
	return qrc.set("placeholder", value)
}

// PopOverContainerSelector 弹窗容器选择器
func (qrc QuarterRangeControl) PopOverContainerSelector(value string) QuarterRangeControl {
	return qrc.set("popOverContainerSelector", value)
}

// Ranges 日期范围快捷键
func (qrc QuarterRangeControl) Ranges(value string) QuarterRangeControl {
	return qrc.set("ranges", value)
}

// ReadOnly 是否只读
func (qrc QuarterRangeControl) ReadOnly(value bool) QuarterRangeControl {
	return qrc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (qrc QuarterRangeControl) ReadOnlyOn(value string) QuarterRangeControl {
	return qrc.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (qrc QuarterRangeControl) Remark(value string) QuarterRangeControl {
	return qrc.set("remark", value)
}

// Required 是否为必填
func (qrc QuarterRangeControl) Required(value bool) QuarterRangeControl {
	return qrc.set("required", value)
}

// Row
func (qrc QuarterRangeControl) Row(value string) QuarterRangeControl {
	return qrc.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (qrc QuarterRangeControl) SaveImmediately(value bool) QuarterRangeControl {
	return qrc.set("saveImmediately", value)
}

// Shortcuts 日期范围快捷键
func (qrc QuarterRangeControl) Shortcuts(value string) QuarterRangeControl {
	return qrc.set("shortcuts", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (qrc QuarterRangeControl) Size(value string) QuarterRangeControl {
	return qrc.set("size", value)
}

// StartPlaceholder 日期范围开始时间-占位符
func (qrc QuarterRangeControl) StartPlaceholder(value string) QuarterRangeControl {
	return qrc.set("startPlaceholder", value)
}

// Static 是否静态展示
func (qrc QuarterRangeControl) Static(value bool) QuarterRangeControl {
	return qrc.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (qrc QuarterRangeControl) StaticClassName(value string) QuarterRangeControl {
	return qrc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (qrc QuarterRangeControl) StaticInputClassName(value string) QuarterRangeControl {
	return qrc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。    className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：    className: {         "red": "data.progress > 80",         "blue": "data.progress > 60"     })
func (qrc QuarterRangeControl) StaticLabelClassName(value string) QuarterRangeControl {
	return qrc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (qrc QuarterRangeControl) StaticOn(value string) QuarterRangeControl {
	return qrc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (qrc QuarterRangeControl) StaticPlaceholder(value string) QuarterRangeControl {
	return qrc.set("staticPlaceholder", value)
}

// StaticSchema
func (qrc QuarterRangeControl) StaticSchema(value string) QuarterRangeControl {
	return qrc.set("staticSchema", value)
}

// Style 组件样式
func (qrc QuarterRangeControl) Style(value string) QuarterRangeControl {
	return qrc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (qrc QuarterRangeControl) SubmitOnChange(value bool) QuarterRangeControl {
	return qrc.set("submitOnChange", value)
}

// TestIdBuilder
func (qrc QuarterRangeControl) TestIdBuilder(value string) QuarterRangeControl {
	return qrc.set("testIdBuilder", value)
}

// Transform 日期数据处理函数，用来处理选择日期之后的的值(value: moment.Moment, config: {type: 'start' | 'end'; originValue: moment.Moment, timeFormat: string}, props: any, data: any, moment: moment) => moment.Moment;
func (qrc QuarterRangeControl) Transform(value string) QuarterRangeControl {
	return qrc.set("transform", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (qrc QuarterRangeControl) UseMobileUI(value bool) QuarterRangeControl {
	return qrc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (qrc QuarterRangeControl) ValidateApi(value string) QuarterRangeControl {
	return qrc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (qrc QuarterRangeControl) ValidateOnChange(value bool) QuarterRangeControl {
	return qrc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (qrc QuarterRangeControl) ValidationErrors(value string) QuarterRangeControl {
	return qrc.set("validationErrors", value)
}

// Validations
func (qrc QuarterRangeControl) Validations(value string) QuarterRangeControl {
	return qrc.set("validations", value)
}

// Value 这里面 value 需要特殊说明一下，因为支持相对值。* `-2mins` 2分钟前\n * `+2days` 2天后\n* `-10week` 十周前\n可用单位： `min`、`hour`、`day`、`week`、`month`、`year`。所有单位支持复数形式。
func (qrc QuarterRangeControl) Value(value string) QuarterRangeControl {
	return qrc.set("value", value)
}

// ValueFormat 用来提交的时间格式。更多格式类型请参考 moment.（新：同format）
func (qrc QuarterRangeControl) ValueFormat(value string) QuarterRangeControl {
	return qrc.set("valueFormat", value)
}

// Visible 是否显示
func (qrc QuarterRangeControl) Visible(value bool) QuarterRangeControl {
	return qrc.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (qrc QuarterRangeControl) VisibleOn(value string) QuarterRangeControl {
	return qrc.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (qrc QuarterRangeControl) Width(value string) QuarterRangeControl {
	return qrc.set("width", value)
}
