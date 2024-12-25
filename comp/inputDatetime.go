package comp

// inputDatetime 日期时间选择控件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/input-datetime
type inputDatetime Schema

func InputDatetime() inputDatetime {
	return make(inputDatetime).set("type", "input-datetime")
}

func (d inputDatetime) set(key string, value any) inputDatetime {
	d[key] = value
	return d
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (d inputDatetime) AutoFill(value string) inputDatetime {
	return d.set("autoFill", value)
}

// BorderMode 边框模式，全边框，还是半边框，或者没边框。 可选值: full | half | none
func (d inputDatetime) BorderMode(value string) inputDatetime {
	return d.set("borderMode", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d inputDatetime) ClassName(value string) inputDatetime {
	return d.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (d inputDatetime) ClearValueOnHidden(value bool) inputDatetime {
	return d.set("clearValueOnHidden", value)
}

// Clearable 是否显示清除按钮
func (d inputDatetime) Clearable(value bool) inputDatetime {
	return d.set("clearable", value)
}

// Desc
func (d inputDatetime) Desc(value string) inputDatetime {
	return d.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (d inputDatetime) Description(value string) inputDatetime {
	return d.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (d inputDatetime) DescriptionClassName(value string) inputDatetime {
	return d.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (d inputDatetime) Disabled(value bool) inputDatetime {
	return d.set("disabled", value)
}

// DisabledDate 字符串函数，用来决定是否禁用某个日期。(currentDate: moment.Moment, props: any) => boolean;
func (d inputDatetime) DisabledDate(value string) inputDatetime {
	return d.set("disabledDate", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (d inputDatetime) DisabledOn(value string) inputDatetime {
	return d.set("disabledOn", value)
}

// DisplayFormat 日期展示格式(新：替代inputFormat)
func (d inputDatetime) DisplayFormat(value string) inputDatetime {
	return d.set("displayFormat", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (d inputDatetime) EditorSetting(value string) inputDatetime {
	return d.set("editorSetting", value)
}

// Emebed 是否为内联模式？
func (d inputDatetime) Emebed(value bool) inputDatetime {
	return d.set("emebed", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (d inputDatetime) ExtraName(value string) inputDatetime {
	return d.set("extraName", value)
}

// Format 日期存储格式
func (d inputDatetime) Format(value string) inputDatetime {
	return d.set("format", value)
}

// Hidden 是否隐藏
func (d inputDatetime) Hidden(value bool) inputDatetime {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (d inputDatetime) HiddenOn(value string) inputDatetime {
	return d.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (d inputDatetime) Hint(value string) inputDatetime {
	return d.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (d inputDatetime) Horizontal(value string) inputDatetime {
	return d.set("horizontal", value)
}

// ID 组件唯一 id，主要用于日志采集
func (d inputDatetime) ID(value string) inputDatetime {
	return d.set("id", value)
}

// InitAutoFill
func (d inputDatetime) InitAutoFill(value string) inputDatetime {
	return d.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (d inputDatetime) Inline(value bool) inputDatetime {
	return d.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (d inputDatetime) InputClassName(value string) inputDatetime {
	return d.set("inputClassName", value)
}

// InputFormat 日期展示格式
func (d inputDatetime) InputFormat(value string) inputDatetime {
	return d.set("inputFormat", value)
}

// IsEndDate 是否为结束时间，如果是，那么会自动加上 23:59:59
func (d inputDatetime) IsEndDate(value bool) inputDatetime {
	return d.set("isEndDate", value)
}

// Label 描述标题
func (d inputDatetime) Label(value string) inputDatetime {
	return d.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (d inputDatetime) LabelAlign(value string) inputDatetime {
	return d.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (d inputDatetime) LabelClassName(value string) inputDatetime {
	return d.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (d inputDatetime) LabelRemark(value string) inputDatetime {
	return d.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (d inputDatetime) LabelWidth(value string) inputDatetime {
	return d.set("labelWidth", value)
}

// MaxDate 限制最大日期
func (d inputDatetime) MaxDate(value string) inputDatetime {
	return d.set("maxDate", value)
}

// MinDate 限制最小日期
func (d inputDatetime) MinDate(value string) inputDatetime {
	return d.set("minDate", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (d inputDatetime) Mode(value string) inputDatetime {
	return d.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (d inputDatetime) Name(value string) inputDatetime {
	return d.set("name", value)
}

// OnEvent 事件动作配置
func (d inputDatetime) OnEvent(value any) inputDatetime {
	return d.set("onEvent", value)
}

// Placeholder 占位符
func (d inputDatetime) Placeholder(value string) inputDatetime {
	return d.set("placeholder", value)
}

// ReadOnly 是否只读
func (d inputDatetime) ReadOnly(value bool) inputDatetime {
	return d.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (d inputDatetime) ReadOnlyOn(value string) inputDatetime {
	return d.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (d inputDatetime) Remark(value string) inputDatetime {
	return d.set("remark", value)
}

// Required 是否为必填
func (d inputDatetime) Required(value bool) inputDatetime {
	return d.set("required", value)
}

// Row
func (d inputDatetime) Row(value string) inputDatetime {
	return d.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (d inputDatetime) SaveImmediately(value bool) inputDatetime {
	return d.set("saveImmediately", value)
}

// Shortcuts 日期快捷键
func (d inputDatetime) Shortcuts(value string) inputDatetime {
	return d.set("shortcuts", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (d inputDatetime) Size(value string) inputDatetime {
	return d.set("size", value)
}

// Static 是否静态展示
func (d inputDatetime) Static(value bool) inputDatetime {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d inputDatetime) StaticClassName(value string) inputDatetime {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d inputDatetime) StaticInputClassName(value string) inputDatetime {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (d inputDatetime) StaticLabelClassName(value string) inputDatetime {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (d inputDatetime) StaticOn(value string) inputDatetime {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d inputDatetime) StaticPlaceholder(value string) inputDatetime {
	return d.set("staticPlaceholder", value)
}

// StaticSchema
func (d inputDatetime) StaticSchema(value string) inputDatetime {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d inputDatetime) Style(value any) inputDatetime {
	return d.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (d inputDatetime) SubmitOnChange(value bool) inputDatetime {
	return d.set("submitOnChange", value)
}

// TestIdBuilder
func (d inputDatetime) TestIdBuilder(value string) inputDatetime {
	return d.set("testIdBuilder", value)
}

// TimeConstraints 时间输入范围限制
func (d inputDatetime) TimeConstraints(value string) inputDatetime {
	return d.set("timeConstraints", value)
}

// TimeFormat 时间的格式。
func (d inputDatetime) TimeFormat(value string) inputDatetime {
	return d.set("timeFormat", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (d inputDatetime) UseMobileUI(value bool) inputDatetime {
	return d.set("useMobileUI", value)
}

// Utc 设定是否存储 utc 时间。
func (d inputDatetime) Utc(value bool) inputDatetime {
	return d.set("utc", value)
}

// ValidateApi 远端校验表单项接口
func (d inputDatetime) ValidateApi(value string) inputDatetime {
	return d.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证，如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (d inputDatetime) ValidateOnChange(value bool) inputDatetime {
	return d.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (d inputDatetime) ValidationErrors(value string) inputDatetime {
	return d.set("validationErrors", value)
}

// Validations
func (d inputDatetime) Validations(value string) inputDatetime {
	return d.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (d inputDatetime) Value(value string) inputDatetime {
	return d.set("value", value)
}

// ValueFormat 替代format
func (d inputDatetime) ValueFormat(value string) inputDatetime {
	return d.set("valueFormat", value)
}

// Visible 是否显示
func (d inputDatetime) Visible(value bool) inputDatetime {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (d inputDatetime) VisibleOn(value string) inputDatetime {
	return d.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (d inputDatetime) Width(value string) inputDatetime {
	return d.set("width", value)
}
