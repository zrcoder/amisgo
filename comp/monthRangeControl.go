package comp

// MonthRangeControl 月范围控件
//
// 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/month-range
// @version 6.7.0
type MonthRangeControl Schema

// NewMonthRangeControl 创建一个新的 MonthRangeControl 实例
func NewMonthRangeControl() MonthRangeControl {
	return MonthRangeControl{}.set("type", "input-month-range")
}

// animation 是否启用游标动画，默认开启
func (m MonthRangeControl) Animation(value bool) MonthRangeControl {
	return m.set("animation", value)
}

// autoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (m MonthRangeControl) AutoFill(value string) MonthRangeControl {
	return m.set("autoFill", value)
}

// borderMode 边框模式，全边框，还是半边框，或者没边框。可选值: full | half | none
func (m MonthRangeControl) BorderMode(value string) MonthRangeControl {
	return m.set("borderMode", value)
}

// className 容器 css 类名
func (m MonthRangeControl) ClassName(value string) MonthRangeControl {
	return m.set("className", value)
}

// clearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (m MonthRangeControl) ClearValueOnHidden(value bool) MonthRangeControl {
	return m.set("clearValueOnHidden", value)
}

// delimiter 分割符, 因为有两个值，开始时间和结束时间，所以要有连接符。默认为英文逗号。
func (m MonthRangeControl) Delimiter(value string) MonthRangeControl {
	return m.set("delimiter", value)
}

// desc 描述内容
func (m MonthRangeControl) Desc(value string) MonthRangeControl {
	return m.set("desc", value)
}

// description 描述内容，支持 Html 片段
func (m MonthRangeControl) Description(value string) MonthRangeControl {
	return m.set("description", value)
}

// descriptionClassName 配置描述上的 className
func (m MonthRangeControl) DescriptionClassName(value string) MonthRangeControl {
	return m.set("descriptionClassName", value)
}

// disabled 是否禁用
func (m MonthRangeControl) Disabled(value bool) MonthRangeControl {
	return m.set("disabled", value)
}

// disabledOn 是否禁用表达式
func (m MonthRangeControl) DisabledOn(value string) MonthRangeControl {
	return m.set("disabledOn", value)
}

// displayFormat 用来配置显示的时间格式
func (m MonthRangeControl) DisplayFormat(value string) MonthRangeControl {
	return m.set("displayFormat", value)
}

// editorSetting 编辑器配置
func (m MonthRangeControl) EditorSetting(value string) MonthRangeControl {
	return m.set("editorSetting", value)
}

// embed 开启后变成非弹出模式，即内联模式
func (m MonthRangeControl) Embed(value bool) MonthRangeControl {
	return m.set("embed", value)
}

// endPlaceholder 日期范围结束时间-占位符
func (m MonthRangeControl) EndPlaceholder(value string) MonthRangeControl {
	return m.set("endPlaceholder", value)
}

// extraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (m MonthRangeControl) ExtraName(value string) MonthRangeControl {
	return m.set("extraName", value)
}

// format 默认 `X` 即时间戳格式，用来提交的时间格式
func (m MonthRangeControl) Format(value string) MonthRangeControl {
	return m.set("format", value)
}

// hidden 是否隐藏
func (m MonthRangeControl) Hidden(value bool) MonthRangeControl {
	return m.set("hidden", value)
}

// hiddenOn 是否隐藏表达式
func (m MonthRangeControl) HiddenOn(value string) MonthRangeControl {
	return m.set("hiddenOn", value)
}

// hint 输入提示，聚焦的时候显示
func (m MonthRangeControl) Hint(value string) MonthRangeControl {
	return m.set("hint", value)
}

// horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (m MonthRangeControl) Horizontal(value string) MonthRangeControl {
	return m.set("horizontal", value)
}

// id 组件唯一 id，主要用于日志采集
func (m MonthRangeControl) Id(value string) MonthRangeControl {
	return m.set("id", value)
}

// initAutoFill 初始化自动填充
func (m MonthRangeControl) InitAutoFill(value string) MonthRangeControl {
	return m.set("initAutoFill", value)
}

// inline 表单 control 是否为 inline 模式
func (m MonthRangeControl) Inline(value bool) MonthRangeControl {
	return m.set("inline", value)
}

// inputClassName 配置 input className
func (m MonthRangeControl) InputClassName(value string) MonthRangeControl {
	return m.set("inputClassName", value)
}

// inputFormat 默认 `YYYY-MM-DD` 用来配置显示的时间格式
func (m MonthRangeControl) InputFormat(value string) MonthRangeControl {
	return m.set("inputFormat", value)
}

// joinValues 开启后将选中的选项 value 的值用连接符拼接起来
func (m MonthRangeControl) JoinValues(value bool) MonthRangeControl {
	return m.set("joinValues", value)
}

// label 描述标题
func (m MonthRangeControl) Label(value string) MonthRangeControl {
	return m.set("label", value)
}

// labelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (m MonthRangeControl) LabelAlign(value string) MonthRangeControl {
	return m.set("labelAlign", value)
}

// labelClassName 配置 label className
func (m MonthRangeControl) LabelClassName(value string) MonthRangeControl {
	return m.set("labelClassName", value)
}

// labelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (m MonthRangeControl) LabelRemark(value string) MonthRangeControl {
	return m.set("labelRemark", value)
}

// labelWidth label自定义宽度，默认单位为px
func (m MonthRangeControl) LabelWidth(value string) MonthRangeControl {
	return m.set("labelWidth", value)
}

// maxDate 最大日期限制，支持变量 $xxx 来取值
func (m MonthRangeControl) MaxDate(value string) MonthRangeControl {
	return m.set("maxDate", value)
}

// maxDuration 最大跨度，比如 2days
func (m MonthRangeControl) MaxDuration(value string) MonthRangeControl {
	return m.set("maxDuration", value)
}

// minDate 最小日期限制，支持变量 $xxx 来取值
func (m MonthRangeControl) MinDate(value string) MonthRangeControl {
	return m.set("minDate", value)
}

// minDuration 最小跨度，比如 2days
func (m MonthRangeControl) MinDuration(value string) MonthRangeControl {
	return m.set("minDuration", value)
}

// mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (m MonthRangeControl) Mode(value string) MonthRangeControl {
	return m.set("mode", value)
}

// name 字段名，表单提交时的 key，支持多层级，用.连接
func (m MonthRangeControl) Name(value string) MonthRangeControl {
	return m.set("name", value)
}

// onEvent 事件动作配置
func (m MonthRangeControl) OnEvent(value string) MonthRangeControl {
	return m.set("onEvent", value)
}

// placeholder 占位符
func (m MonthRangeControl) Placeholder(value string) MonthRangeControl {
	return m.set("placeholder", value)
}

// popOverContainerSelector 弹窗容器选择器
func (m MonthRangeControl) PopOverContainerSelector(value string) MonthRangeControl {
	return m.set("popOverContainerSelector", value)
}

// ranges 日期范围快捷键
func (m MonthRangeControl) Ranges(value string) MonthRangeControl {
	return m.set("ranges", value)
}

// readOnly 是否只读
func (m MonthRangeControl) ReadOnly(value bool) MonthRangeControl {
	return m.set("readOnly", value)
}

// readOnlyOn 只读条件
func (m MonthRangeControl) ReadOnlyOn(value string) MonthRangeControl {
	return m.set("readOnlyOn", value)
}

// remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (m MonthRangeControl) Remark(value string) MonthRangeControl {
	return m.set("remark", value)
}

// required 是否为必填
func (m MonthRangeControl) Required(value bool) MonthRangeControl {
	return m.set("required", value)
}

// row 配置行
func (m MonthRangeControl) Row(value string) MonthRangeControl {
	return m.set("row", value)
}

// saveImmediately 是否立即保存
func (m MonthRangeControl) SaveImmediately(value bool) MonthRangeControl {
	return m.set("saveImmediately", value)
}

// shortcuts 日期范围快捷键
func (m MonthRangeControl) Shortcuts(value string) MonthRangeControl {
	return m.set("shortcuts", value)
}

// size 表单项大小 可选值: xs | sm | md | lg | full
func (m MonthRangeControl) Size(value string) MonthRangeControl {
	return m.set("size", value)
}

// startPlaceholder 日期范围开始时间-占位符
func (m MonthRangeControl) StartPlaceholder(value string) MonthRangeControl {
	return m.set("startPlaceholder", value)
}

// static 是否静态展示
func (m MonthRangeControl) Static(value bool) MonthRangeControl {
	return m.set("static", value)
}

// staticClassName 静态展示表单项类名
func (m MonthRangeControl) StaticClassName(value string) MonthRangeControl {
	return m.set("staticClassName", value)
}

// staticInputClassName 静态展示表单项Value类名
func (m MonthRangeControl) StaticInputClassName(value string) MonthRangeControl {
	return m.set("staticInputClassName", value)
}

// staticLabelClassName 静态展示表单项Label类名
func (m MonthRangeControl) StaticLabelClassName(value string) MonthRangeControl {
	return m.set("staticLabelClassName", value)
}

// staticOn 是否静态展示表达式
func (m MonthRangeControl) StaticOn(value string) MonthRangeControl {
	return m.set("staticOn", value)
}

// staticPlaceholder 静态展示空值占位
func (m MonthRangeControl) StaticPlaceholder(value string) MonthRangeControl {
	return m.set("staticPlaceholder", value)
}

// staticSchema 静态展示模式的 Schema
func (m MonthRangeControl) StaticSchema(value string) MonthRangeControl {
	return m.set("staticSchema", value)
}

// style 组件样式
func (m MonthRangeControl) Style(value string) MonthRangeControl {
	return m.set("style", value)
}

// submitOnChange 当修改完的时候是否提交表单
func (m MonthRangeControl) SubmitOnChange(value bool) MonthRangeControl {
	return m.set("submitOnChange", value)
}

// testIdBuilder 测试 ID 构建器
func (m MonthRangeControl) TestIdBuilder(value string) MonthRangeControl {
	return m.set("testIdBuilder", value)
}

// transform 日期数据处理函数
func (m MonthRangeControl) Transform(value string) MonthRangeControl {
	return m.set("transform", value)
}

// useMobileUI 组件级别用来关闭移动端样式
func (m MonthRangeControl) UseMobileUI(value bool) MonthRangeControl {
	return m.set("useMobileUI", value)
}

// validateApi 远端校验表单项接口
func (m MonthRangeControl) ValidateApi(value string) MonthRangeControl {
	return m.set("validateApi", value)
}

// validateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (m MonthRangeControl) ValidateOnChange(value bool) MonthRangeControl {
	return m.set("validateOnChange", value)
}

// validationErrors 验证失败的提示信息
func (m MonthRangeControl) ValidationErrors(value string) MonthRangeControl {
	return m.set("validationErrors", value)
}

// validations 验证规则
func (m MonthRangeControl) Validations(value string) MonthRangeControl {
	return m.set("validations", value)
}

// value 这里面 value 需要特殊说明一下，因为支持相对值
func (m MonthRangeControl) Value(value string) MonthRangeControl {
	return m.set("value", value)
}

// valueFormat 用来提交的时间格式
func (m MonthRangeControl) ValueFormat(value string) MonthRangeControl {
	return m.set("valueFormat", value)
}

// visible 是否显示
func (m MonthRangeControl) Visible(value bool) MonthRangeControl {
	return m.set("visible", value)
}

// visibleOn 是否显示表达式
func (m MonthRangeControl) VisibleOn(value string) MonthRangeControl {
	return m.set("visibleOn", value)
}

// width 在 Table 中调整宽度
func (m MonthRangeControl) Width(value string) MonthRangeControl {
	return m.set("width", value)
}

func (m MonthRangeControl) set(key string, value interface{}) MonthRangeControl {
	m[key] = value
	return m
}
