package comp

// MonthControl 月份选择控件
//
// 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/Month
// @version 6.7.0
type MonthControl Schema

// NewMonthControl 创建一个新的 MonthControl 实例
func NewMonthControl() MonthControl {
	return MonthControl{}.set("type", "input-month")
}

func (mc MonthControl) set(key string, value interface{}) MonthControl {
	mc[key] = value
	return mc
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (mc MonthControl) AutoFill(value string) MonthControl {
	return mc.set("autoFill", value)
}

// BorderMode 边框模式，全边框，还是半边框，或者没边框。可选值: full | half | none
func (mc MonthControl) BorderMode(value string) MonthControl {
	return mc.set("borderMode", value)
}

// ClassName 容器 css 类名
func (mc MonthControl) ClassName(value string) MonthControl {
	return mc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (mc MonthControl) ClearValueOnHidden(value bool) MonthControl {
	return mc.set("clearValueOnHidden", value)
}

// Clearable 是否显示清除按钮
func (mc MonthControl) Clearable(value bool) MonthControl {
	return mc.set("clearable", value)
}

// Desc
func (mc MonthControl) Desc(value string) MonthControl {
	return mc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (mc MonthControl) Description(value string) MonthControl {
	return mc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (mc MonthControl) DescriptionClassName(value string) MonthControl {
	return mc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (mc MonthControl) Disabled(value bool) MonthControl {
	return mc.set("disabled", value)
}

// DisabledDate 字符串函数，用来决定是否禁用某个日期
func (mc MonthControl) DisabledDate(value string) MonthControl {
	return mc.set("disabledDate", value)
}

// DisabledOn 是否禁用表达式
func (mc MonthControl) DisabledOn(value string) MonthControl {
	return mc.set("disabledOn", value)
}

// DisplayFormat 日期展示格式
func (mc MonthControl) DisplayFormat(value string) MonthControl {
	return mc.set("displayFormat", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (mc MonthControl) EditorSetting(value string) MonthControl {
	return mc.set("editorSetting", value)
}

// Emebed 是否为内联模式
func (mc MonthControl) Emebed(value bool) MonthControl {
	return mc.set("emebed", value)
}

// ExtraName 额外的字段名
func (mc MonthControl) ExtraName(value string) MonthControl {
	return mc.set("extraName", value)
}

// Format 月份存储格式
func (mc MonthControl) Format(value string) MonthControl {
	return mc.set("format", value)
}

// Hidden 是否隐藏
func (mc MonthControl) Hidden(value bool) MonthControl {
	return mc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (mc MonthControl) HiddenOn(value string) MonthControl {
	return mc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (mc MonthControl) Hint(value string) MonthControl {
	return mc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (mc MonthControl) Horizontal(value string) MonthControl {
	return mc.set("horizontal", value)
}

// ID 组件唯一 id，主要用于日志采集
func (mc MonthControl) ID(value string) MonthControl {
	return mc.set("id", value)
}

// InitAutoFill
func (mc MonthControl) InitAutoFill(value string) MonthControl {
	return mc.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式
func (mc MonthControl) Inline(value bool) MonthControl {
	return mc.set("inline", value)
}

// InputClassName 配置 input className
func (mc MonthControl) InputClassName(value string) MonthControl {
	return mc.set("inputClassName", value)
}

// InputFormat 月份展示格式
func (mc MonthControl) InputFormat(value string) MonthControl {
	return mc.set("inputFormat", value)
}

// Label 描述标题
func (mc MonthControl) Label(value string) MonthControl {
	return mc.set("label", value)
}

// LabelAlign 描述标题可选值: right | left | top | inherit
func (mc MonthControl) LabelAlign(value string) MonthControl {
	return mc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (mc MonthControl) LabelClassName(value string) MonthControl {
	return mc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (mc MonthControl) LabelRemark(value string) MonthControl {
	return mc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (mc MonthControl) LabelWidth(value string) MonthControl {
	return mc.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (mc MonthControl) Mode(value string) MonthControl {
	return mc.set("mode", value)
}

// Name 字段名，表单提交时的 key
func (mc MonthControl) Name(value string) MonthControl {
	return mc.set("name", value)
}

// OnEvent 事件动作配置
func (mc MonthControl) OnEvent(value string) MonthControl {
	return mc.set("onEvent", value)
}

// Placeholder 占位符
func (mc MonthControl) Placeholder(value string) MonthControl {
	return mc.set("placeholder", value)
}

// ReadOnly 是否只读
func (mc MonthControl) ReadOnly(value bool) MonthControl {
	return mc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (mc MonthControl) ReadOnlyOn(value string) MonthControl {
	return mc.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (mc MonthControl) Remark(value string) MonthControl {
	return mc.set("remark", value)
}

// Required 是否为必填
func (mc MonthControl) Required(value bool) MonthControl {
	return mc.set("required", value)
}

// Row
func (mc MonthControl) Row(value string) MonthControl {
	return mc.set("row", value)
}

// SaveImmediately 是否立即保存
func (mc MonthControl) SaveImmediately(value bool) MonthControl {
	return mc.set("saveImmediately", value)
}

// Shortcuts 日期快捷键
func (mc MonthControl) Shortcuts(value string) MonthControl {
	return mc.set("shortcuts", value)
}

// Size 表单项大小
func (mc MonthControl) Size(value string) MonthControl {
	return mc.set("size", value)
}

// Static 是否静态展示
func (mc MonthControl) Static(value bool) MonthControl {
	return mc.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (mc MonthControl) StaticClassName(value string) MonthControl {
	return mc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (mc MonthControl) StaticInputClassName(value string) MonthControl {
	return mc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (mc MonthControl) StaticLabelClassName(value string) MonthControl {
	return mc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (mc MonthControl) StaticOn(value string) MonthControl {
	return mc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (mc MonthControl) StaticPlaceholder(value string) MonthControl {
	return mc.set("staticPlaceholder", value)
}

// StaticSchema
func (mc MonthControl) StaticSchema(value string) MonthControl {
	return mc.set("staticSchema", value)
}

// Style 组件样式
func (mc MonthControl) Style(value string) MonthControl {
	return mc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (mc MonthControl) SubmitOnChange(value bool) MonthControl {
	return mc.set("submitOnChange", value)
}

// TestIdBuilder
func (mc MonthControl) TestIdBuilder(value string) MonthControl {
	return mc.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (mc MonthControl) UseMobileUI(value bool) MonthControl {
	return mc.set("useMobileUI", value)
}

// Utc 设定是否存储 utc 时间
func (mc MonthControl) Utc(value bool) MonthControl {
	return mc.set("utc", value)
}

// ValidateApi 远端校验表单项接口
func (mc MonthControl) ValidateApi(value string) MonthControl {
	return mc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (mc MonthControl) ValidateOnChange(value bool) MonthControl {
	return mc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (mc MonthControl) ValidationErrors(value string) MonthControl {
	return mc.set("validationErrors", value)
}

// Validations
func (mc MonthControl) Validations(value string) MonthControl {
	return mc.set("validations", value)
}

// Value 默认值，切记只能是静态值
func (mc MonthControl) Value(value string) MonthControl {
	return mc.set("value", value)
}

// ValueFormat 替代 format
func (mc MonthControl) ValueFormat(value string) MonthControl {
	return mc.set("valueFormat", value)
}

// Visible 是否显示
func (mc MonthControl) Visible(value bool) MonthControl {
	return mc.set("visible", value)
}

// VisibleOn 是否显示表达式
func (mc MonthControl) VisibleOn(value string) MonthControl {
	return mc.set("visibleOn", value)
}

// Width 在 Table 中调整宽度
func (mc MonthControl) Width(value string) MonthControl {
	return mc.set("width", value)
}
