package comp

// inputMonth 月份选择控件
// 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/Month

type inputMonth schema

func InputMonth() inputMonth {
	return inputMonth{}.set("type", "input-month")
}

func (mc inputMonth) set(key string, value any) inputMonth {
	mc[key] = value
	return mc
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (mc inputMonth) AutoFill(value string) inputMonth {
	return mc.set("autoFill", value)
}

// BorderMode 边框模式，全边框，还是半边框，或者没边框。可选值: full | half | none
func (mc inputMonth) BorderMode(value string) inputMonth {
	return mc.set("borderMode", value)
}

// ClassName 容器 css 类名
func (mc inputMonth) ClassName(value string) inputMonth {
	return mc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (mc inputMonth) ClearValueOnHidden(value bool) inputMonth {
	return mc.set("clearValueOnHidden", value)
}

// Clearable 是否显示清除按钮
func (mc inputMonth) Clearable(value bool) inputMonth {
	return mc.set("clearable", value)
}

// Desc
func (mc inputMonth) Desc(value string) inputMonth {
	return mc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (mc inputMonth) Description(value string) inputMonth {
	return mc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (mc inputMonth) DescriptionClassName(value string) inputMonth {
	return mc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (mc inputMonth) Disabled(value bool) inputMonth {
	return mc.set("disabled", value)
}

// DisabledDate 字符串函数，用来决定是否禁用某个日期
func (mc inputMonth) DisabledDate(value string) inputMonth {
	return mc.set("disabledDate", value)
}

// DisabledOn 是否禁用表达式
func (mc inputMonth) DisabledOn(value string) inputMonth {
	return mc.set("disabledOn", value)
}

// DisplayFormat 日期展示格式
func (mc inputMonth) DisplayFormat(value string) inputMonth {
	return mc.set("displayFormat", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (mc inputMonth) EditorSetting(value string) inputMonth {
	return mc.set("editorSetting", value)
}

// Emebed 是否为内联模式
func (mc inputMonth) Emebed(value bool) inputMonth {
	return mc.set("emebed", value)
}

// ExtraName 额外的字段名
func (mc inputMonth) ExtraName(value string) inputMonth {
	return mc.set("extraName", value)
}

// Format 月份存储格式
func (mc inputMonth) Format(value string) inputMonth {
	return mc.set("format", value)
}

// Hidden 是否隐藏
func (mc inputMonth) Hidden(value bool) inputMonth {
	return mc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (mc inputMonth) HiddenOn(value string) inputMonth {
	return mc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (mc inputMonth) Hint(value string) inputMonth {
	return mc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (mc inputMonth) Horizontal(value string) inputMonth {
	return mc.set("horizontal", value)
}

// ID 组件唯一 id，主要用于日志采集
func (mc inputMonth) ID(value string) inputMonth {
	return mc.set("id", value)
}

// InitAutoFill
func (mc inputMonth) InitAutoFill(value string) inputMonth {
	return mc.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式
func (mc inputMonth) Inline(value bool) inputMonth {
	return mc.set("inline", value)
}

// InputClassName 配置 input className
func (mc inputMonth) InputClassName(value string) inputMonth {
	return mc.set("inputClassName", value)
}

// InputFormat 月份展示格式
func (mc inputMonth) InputFormat(value string) inputMonth {
	return mc.set("inputFormat", value)
}

// Label 描述标题
func (mc inputMonth) Label(value string) inputMonth {
	return mc.set("label", value)
}

// LabelAlign 描述标题可选值: right | left | top | inherit
func (mc inputMonth) LabelAlign(value string) inputMonth {
	return mc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (mc inputMonth) LabelClassName(value string) inputMonth {
	return mc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (mc inputMonth) LabelRemark(value string) inputMonth {
	return mc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (mc inputMonth) LabelWidth(value string) inputMonth {
	return mc.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (mc inputMonth) Mode(value string) inputMonth {
	return mc.set("mode", value)
}

// Name 字段名，表单提交时的 key
func (mc inputMonth) Name(value string) inputMonth {
	return mc.set("name", value)
}

// OnEvent 事件动作配置
func (mc inputMonth) OnEvent(value any) inputMonth {
	return mc.set("onEvent", value)
}

// Placeholder 占位符
func (mc inputMonth) Placeholder(value string) inputMonth {
	return mc.set("placeholder", value)
}

// ReadOnly 是否只读
func (mc inputMonth) ReadOnly(value bool) inputMonth {
	return mc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (mc inputMonth) ReadOnlyOn(value string) inputMonth {
	return mc.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (mc inputMonth) Remark(value string) inputMonth {
	return mc.set("remark", value)
}

// Required 是否为必填
func (mc inputMonth) Required(value bool) inputMonth {
	return mc.set("required", value)
}

// Row
func (mc inputMonth) Row(value string) inputMonth {
	return mc.set("row", value)
}

// SaveImmediately 是否立即保存
func (mc inputMonth) SaveImmediately(value bool) inputMonth {
	return mc.set("saveImmediately", value)
}

// Shortcuts 日期快捷键
func (mc inputMonth) Shortcuts(value string) inputMonth {
	return mc.set("shortcuts", value)
}

// Size 表单项大小
func (mc inputMonth) Size(value string) inputMonth {
	return mc.set("size", value)
}

// Static 是否静态展示
func (mc inputMonth) Static(value bool) inputMonth {
	return mc.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (mc inputMonth) StaticClassName(value string) inputMonth {
	return mc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (mc inputMonth) StaticInputClassName(value string) inputMonth {
	return mc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (mc inputMonth) StaticLabelClassName(value string) inputMonth {
	return mc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (mc inputMonth) StaticOn(value string) inputMonth {
	return mc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (mc inputMonth) StaticPlaceholder(value string) inputMonth {
	return mc.set("staticPlaceholder", value)
}

// StaticSchema
func (mc inputMonth) StaticSchema(value string) inputMonth {
	return mc.set("staticSchema", value)
}

// Style 组件样式
func (mc inputMonth) Style(value any) inputMonth {
	return mc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (mc inputMonth) SubmitOnChange(value bool) inputMonth {
	return mc.set("submitOnChange", value)
}

// TestIdBuilder
func (mc inputMonth) TestIdBuilder(value string) inputMonth {
	return mc.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (mc inputMonth) UseMobileUI(value bool) inputMonth {
	return mc.set("useMobileUI", value)
}

// Utc 设定是否存储 utc 时间
func (mc inputMonth) Utc(value bool) inputMonth {
	return mc.set("utc", value)
}

// ValidateApi 远端校验表单项接口
func (mc inputMonth) ValidateApi(value string) inputMonth {
	return mc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (mc inputMonth) ValidateOnChange(value bool) inputMonth {
	return mc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (mc inputMonth) ValidationErrors(value string) inputMonth {
	return mc.set("validationErrors", value)
}

// Validations
func (mc inputMonth) Validations(value string) inputMonth {
	return mc.set("validations", value)
}

// Value 默认值，切记只能是静态值
func (mc inputMonth) Value(value string) inputMonth {
	return mc.set("value", value)
}

// ValueFormat 替代 format
func (mc inputMonth) ValueFormat(value string) inputMonth {
	return mc.set("valueFormat", value)
}

// Visible 是否显示
func (mc inputMonth) Visible(value bool) inputMonth {
	return mc.set("visible", value)
}

// VisibleOn 是否显示表达式
func (mc inputMonth) VisibleOn(value string) inputMonth {
	return mc.set("visibleOn", value)
}

// Width 在 Table 中调整宽度
func (mc inputMonth) Width(value string) inputMonth {
	return mc.set("width", value)
}
