package comp

// inputMonthRange 月范围控件
// 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/month-range

type inputMonthRange schema

// InputMonthRange 创建一个新的 MonthRangeControl 实例
func InputMonthRange() inputMonthRange {
	return inputMonthRange{}.set("type", "input-month-range")
}

// animation 是否启用游标动画，默认开启
func (m inputMonthRange) Animation(value bool) inputMonthRange {
	return m.set("animation", value)
}

// autoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (m inputMonthRange) AutoFill(value string) inputMonthRange {
	return m.set("autoFill", value)
}

// borderMode 边框模式，全边框，还是半边框，或者没边框。可选值: full | half | none
func (m inputMonthRange) BorderMode(value string) inputMonthRange {
	return m.set("borderMode", value)
}

// className 容器 css 类名
func (m inputMonthRange) ClassName(value string) inputMonthRange {
	return m.set("className", value)
}

// clearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (m inputMonthRange) ClearValueOnHidden(value bool) inputMonthRange {
	return m.set("clearValueOnHidden", value)
}

// delimiter 分割符, 因为有两个值，开始时间和结束时间，所以要有连接符。默认为英文逗号。
func (m inputMonthRange) Delimiter(value string) inputMonthRange {
	return m.set("delimiter", value)
}

// desc 描述内容
func (m inputMonthRange) Desc(value string) inputMonthRange {
	return m.set("desc", value)
}

// description 描述内容，支持 Html 片段
func (m inputMonthRange) Description(value string) inputMonthRange {
	return m.set("description", value)
}

// descriptionClassName 配置描述上的 className
func (m inputMonthRange) DescriptionClassName(value string) inputMonthRange {
	return m.set("descriptionClassName", value)
}

// disabled 是否禁用
func (m inputMonthRange) Disabled(value bool) inputMonthRange {
	return m.set("disabled", value)
}

// disabledOn 是否禁用表达式
func (m inputMonthRange) DisabledOn(value string) inputMonthRange {
	return m.set("disabledOn", value)
}

// displayFormat 用来配置显示的时间格式
func (m inputMonthRange) DisplayFormat(value string) inputMonthRange {
	return m.set("displayFormat", value)
}

// editorSetting 编辑器配置
func (m inputMonthRange) EditorSetting(value string) inputMonthRange {
	return m.set("editorSetting", value)
}

// embed 开启后变成非弹出模式，即内联模式
func (m inputMonthRange) Embed(value bool) inputMonthRange {
	return m.set("embed", value)
}

// endPlaceholder 日期范围结束时间-占位符
func (m inputMonthRange) EndPlaceholder(value string) inputMonthRange {
	return m.set("endPlaceholder", value)
}

// extraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (m inputMonthRange) ExtraName(value string) inputMonthRange {
	return m.set("extraName", value)
}

// format 默认 `X` 即时间戳格式，用来提交的时间格式
func (m inputMonthRange) Format(value string) inputMonthRange {
	return m.set("format", value)
}

// hidden 是否隐藏
func (m inputMonthRange) Hidden(value bool) inputMonthRange {
	return m.set("hidden", value)
}

// hiddenOn 是否隐藏表达式
func (m inputMonthRange) HiddenOn(value string) inputMonthRange {
	return m.set("hiddenOn", value)
}

// hint 输入提示，聚焦的时候显示
func (m inputMonthRange) Hint(value string) inputMonthRange {
	return m.set("hint", value)
}

// horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (m inputMonthRange) Horizontal(value string) inputMonthRange {
	return m.set("horizontal", value)
}

// id 组件唯一 id，主要用于日志采集
func (m inputMonthRange) Id(value string) inputMonthRange {
	return m.set("id", value)
}

// initAutoFill 初始化自动填充
func (m inputMonthRange) InitAutoFill(value string) inputMonthRange {
	return m.set("initAutoFill", value)
}

// inline 表单 control 是否为 inline 模式
func (m inputMonthRange) Inline(value bool) inputMonthRange {
	return m.set("inline", value)
}

// inputClassName 配置 input className
func (m inputMonthRange) InputClassName(value string) inputMonthRange {
	return m.set("inputClassName", value)
}

// inputFormat 默认 `YYYY-MM-DD` 用来配置显示的时间格式
func (m inputMonthRange) InputFormat(value string) inputMonthRange {
	return m.set("inputFormat", value)
}

// joinValues 开启后将选中的选项 value 的值用连接符拼接起来
func (m inputMonthRange) JoinValues(value bool) inputMonthRange {
	return m.set("joinValues", value)
}

// label 描述标题
func (m inputMonthRange) Label(value string) inputMonthRange {
	return m.set("label", value)
}

// labelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (m inputMonthRange) LabelAlign(value string) inputMonthRange {
	return m.set("labelAlign", value)
}

// labelClassName 配置 label className
func (m inputMonthRange) LabelClassName(value string) inputMonthRange {
	return m.set("labelClassName", value)
}

// labelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (m inputMonthRange) LabelRemark(value string) inputMonthRange {
	return m.set("labelRemark", value)
}

// labelWidth label自定义宽度，默认单位为px
func (m inputMonthRange) LabelWidth(value string) inputMonthRange {
	return m.set("labelWidth", value)
}

// maxDate 最大日期限制，支持变量 $xxx 来取值
func (m inputMonthRange) MaxDate(value string) inputMonthRange {
	return m.set("maxDate", value)
}

// maxDuration 最大跨度，比如 2days
func (m inputMonthRange) MaxDuration(value string) inputMonthRange {
	return m.set("maxDuration", value)
}

// minDate 最小日期限制，支持变量 $xxx 来取值
func (m inputMonthRange) MinDate(value string) inputMonthRange {
	return m.set("minDate", value)
}

// minDuration 最小跨度，比如 2days
func (m inputMonthRange) MinDuration(value string) inputMonthRange {
	return m.set("minDuration", value)
}

// mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (m inputMonthRange) Mode(value string) inputMonthRange {
	return m.set("mode", value)
}

// name 字段名，表单提交时的 key，支持多层级，用.连接
func (m inputMonthRange) Name(value string) inputMonthRange {
	return m.set("name", value)
}

// onEvent 事件动作配置
func (m inputMonthRange) OnEvent(value any) inputMonthRange {
	return m.set("onEvent", value)
}

// placeholder 占位符
func (m inputMonthRange) Placeholder(value string) inputMonthRange {
	return m.set("placeholder", value)
}

// popOverContainerSelector 弹窗容器选择器
func (m inputMonthRange) PopOverContainerSelector(value string) inputMonthRange {
	return m.set("popOverContainerSelector", value)
}

// ranges 日期范围快捷键
func (m inputMonthRange) Ranges(value string) inputMonthRange {
	return m.set("ranges", value)
}

// readOnly 是否只读
func (m inputMonthRange) ReadOnly(value bool) inputMonthRange {
	return m.set("readOnly", value)
}

// readOnlyOn 只读条件
func (m inputMonthRange) ReadOnlyOn(value string) inputMonthRange {
	return m.set("readOnlyOn", value)
}

// remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (m inputMonthRange) Remark(value string) inputMonthRange {
	return m.set("remark", value)
}

// required 是否为必填
func (m inputMonthRange) Required(value bool) inputMonthRange {
	return m.set("required", value)
}

// row 配置行
func (m inputMonthRange) Row(value string) inputMonthRange {
	return m.set("row", value)
}

// saveImmediately 是否立即保存
func (m inputMonthRange) SaveImmediately(value bool) inputMonthRange {
	return m.set("saveImmediately", value)
}

// shortcuts 日期范围快捷键
func (m inputMonthRange) Shortcuts(value string) inputMonthRange {
	return m.set("shortcuts", value)
}

// size 表单项大小 可选值: xs | sm | md | lg | full
func (m inputMonthRange) Size(value string) inputMonthRange {
	return m.set("size", value)
}

// startPlaceholder 日期范围开始时间-占位符
func (m inputMonthRange) StartPlaceholder(value string) inputMonthRange {
	return m.set("startPlaceholder", value)
}

// static 是否静态展示
func (m inputMonthRange) Static(value bool) inputMonthRange {
	return m.set("static", value)
}

// staticClassName 静态展示表单项类名
func (m inputMonthRange) StaticClassName(value string) inputMonthRange {
	return m.set("staticClassName", value)
}

// staticInputClassName 静态展示表单项Value类名
func (m inputMonthRange) StaticInputClassName(value string) inputMonthRange {
	return m.set("staticInputClassName", value)
}

// staticLabelClassName 静态展示表单项Label类名
func (m inputMonthRange) StaticLabelClassName(value string) inputMonthRange {
	return m.set("staticLabelClassName", value)
}

// staticOn 是否静态展示表达式
func (m inputMonthRange) StaticOn(value string) inputMonthRange {
	return m.set("staticOn", value)
}

// staticPlaceholder 静态展示空值占位
func (m inputMonthRange) StaticPlaceholder(value string) inputMonthRange {
	return m.set("staticPlaceholder", value)
}

// staticSchema 静态展示模式的 Schema
func (m inputMonthRange) StaticSchema(value string) inputMonthRange {
	return m.set("staticSchema", value)
}

// style 组件样式
func (m inputMonthRange) Style(value any) inputMonthRange {
	return m.set("style", value)
}

// submitOnChange 当修改完的时候是否提交表单
func (m inputMonthRange) SubmitOnChange(value bool) inputMonthRange {
	return m.set("submitOnChange", value)
}

// testIdBuilder 测试 ID 构建器
func (m inputMonthRange) TestIdBuilder(value string) inputMonthRange {
	return m.set("testIdBuilder", value)
}

// transform 日期数据处理函数
func (m inputMonthRange) Transform(value string) inputMonthRange {
	return m.set("transform", value)
}

// useMobileUI 组件级别用来关闭移动端样式
func (m inputMonthRange) UseMobileUI(value bool) inputMonthRange {
	return m.set("useMobileUI", value)
}

// validateApi 远端校验表单项接口
func (m inputMonthRange) ValidateApi(value string) inputMonthRange {
	return m.set("validateApi", value)
}

// validateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (m inputMonthRange) ValidateOnChange(value bool) inputMonthRange {
	return m.set("validateOnChange", value)
}

// validationErrors 验证失败的提示信息
func (m inputMonthRange) ValidationErrors(value string) inputMonthRange {
	return m.set("validationErrors", value)
}

// validations 验证规则
func (m inputMonthRange) Validations(value string) inputMonthRange {
	return m.set("validations", value)
}

// value 这里面 value 需要特殊说明一下，因为支持相对值
func (m inputMonthRange) Value(value string) inputMonthRange {
	return m.set("value", value)
}

// valueFormat 用来提交的时间格式
func (m inputMonthRange) ValueFormat(value string) inputMonthRange {
	return m.set("valueFormat", value)
}

// visible 是否显示
func (m inputMonthRange) Visible(value bool) inputMonthRange {
	return m.set("visible", value)
}

// visibleOn 是否显示表达式
func (m inputMonthRange) VisibleOn(value string) inputMonthRange {
	return m.set("visibleOn", value)
}

// width 在 Table 中调整宽度
func (m inputMonthRange) Width(value string) inputMonthRange {
	return m.set("width", value)
}

func (m inputMonthRange) set(key string, value any) inputMonthRange {
	m[key] = value
	return m
}
