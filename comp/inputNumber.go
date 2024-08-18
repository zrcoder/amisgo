package comp

// inputNumber 数字输入框
//
// @version 6.7.0
type inputNumber schema

// InputNumber 创建一个新的 NumberControl 实例
func InputNumber() inputNumber {
	return inputNumber{}.set("type", "input-number")
}

// set 设置字段值
func (nc inputNumber) set(key string, value interface{}) inputNumber {
	nc[key] = value
	return nc
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (nc inputNumber) AutoFill(value string) inputNumber {
	return nc.set("autoFill", value)
}

// Big 是否是大数，如果是的话输入输出都将是字符串
func (nc inputNumber) Big(value bool) inputNumber {
	return nc.set("big", value)
}

// BorderMode 边框模式，全边框，还是半边框，或者没边框。 可选值: full | half | none
func (nc inputNumber) BorderMode(value string) inputNumber {
	return nc.set("borderMode", value)
}

// ClassName 容器 css 类名
func (nc inputNumber) ClassName(value string) inputNumber {
	return nc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (nc inputNumber) ClearValueOnHidden(value bool) inputNumber {
	return nc.set("clearValueOnHidden", value)
}

// Desc 描述
func (nc inputNumber) Desc(value string) inputNumber {
	return nc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (nc inputNumber) Description(value string) inputNumber {
	return nc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (nc inputNumber) DescriptionClassName(value string) inputNumber {
	return nc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (nc inputNumber) Disabled(value bool) inputNumber {
	return nc.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (nc inputNumber) DisabledOn(value string) inputNumber {
	return nc.set("disabledOn", value)
}

// DisplayMode 输入框为基础输入框还是加强输入框
func (nc inputNumber) DisplayMode(value string) inputNumber {
	return nc.set("displayMode", value)
}

// EditorSetting 编辑器配置
func (nc inputNumber) EditorSetting(value string) inputNumber {
	return nc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (nc inputNumber) ExtraName(value string) inputNumber {
	return nc.set("extraName", value)
}

// Hidden 是否隐藏
func (nc inputNumber) Hidden(value bool) inputNumber {
	return nc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (nc inputNumber) HiddenOn(value string) inputNumber {
	return nc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (nc inputNumber) Hint(value string) inputNumber {
	return nc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (nc inputNumber) Horizontal(value string) inputNumber {
	return nc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (nc inputNumber) Id(value string) inputNumber {
	return nc.set("id", value)
}

// InitAutoFill 初始化自动填充
func (nc inputNumber) InitAutoFill(value string) inputNumber {
	return nc.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式
func (nc inputNumber) Inline(value bool) inputNumber {
	return nc.set("inline", value)
}

// InputClassName 配置 input className
func (nc inputNumber) InputClassName(value string) inputNumber {
	return nc.set("inputClassName", value)
}

// Keyboard 是否启用键盘行为
func (nc inputNumber) Keyboard(value bool) inputNumber {
	return nc.set("keyboard", value)
}

// KilobitSeparator 是否千分分隔
func (nc inputNumber) KilobitSeparator(value bool) inputNumber {
	return nc.set("kilobitSeparator", value)
}

// Label 描述标题
func (nc inputNumber) Label(value string) inputNumber {
	return nc.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (nc inputNumber) LabelAlign(value string) inputNumber {
	return nc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (nc inputNumber) LabelClassName(value string) inputNumber {
	return nc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (nc inputNumber) LabelRemark(value string) inputNumber {
	return nc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (nc inputNumber) LabelWidth(value string) inputNumber {
	return nc.set("labelWidth", value)
}

// Max 最大值
func (nc inputNumber) Max(value string) inputNumber {
	return nc.set("max", value)
}

// Min 最小值
func (nc inputNumber) Min(value string) inputNumber {
	return nc.set("min", value)
}

// Mode 配置当前表单项展示模式
func (nc inputNumber) Mode(value string) inputNumber {
	return nc.set("mode", value)
}

// Name 字段名，表单提交时的 key
func (nc inputNumber) Name(value string) inputNumber {
	return nc.set("name", value)
}

// OnEvent 事件动作配置
func (nc inputNumber) OnEvent(value string) inputNumber {
	return nc.set("onEvent", value)
}

// Placeholder 占位符
func (nc inputNumber) Placeholder(value string) inputNumber {
	return nc.set("placeholder", value)
}

// Precision 精度
func (nc inputNumber) Precision(value string) inputNumber {
	return nc.set("precision", value)
}

// Prefix 前缀
func (nc inputNumber) Prefix(value string) inputNumber {
	return nc.set("prefix", value)
}

// ReadOnly 只读
func (nc inputNumber) ReadOnly(value bool) inputNumber {
	return nc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (nc inputNumber) ReadOnlyOn(value string) inputNumber {
	return nc.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (nc inputNumber) Remark(value string) inputNumber {
	return nc.set("remark", value)
}

// Required 是否为必填
func (nc inputNumber) Required(value bool) inputNumber {
	return nc.set("required", value)
}

// Row
func (nc inputNumber) Row(value string) inputNumber {
	return nc.set("row", value)
}

// SaveImmediately 是否立即保存
func (nc inputNumber) SaveImmediately(value bool) inputNumber {
	return nc.set("saveImmediately", value)
}

// ShowAsPercent 用来开启百分号的展示形式
func (nc inputNumber) ShowAsPercent(value bool) inputNumber {
	return nc.set("showAsPercent", value)
}

// ShowSteps 是否显示上下点击按钮
func (nc inputNumber) ShowSteps(value bool) inputNumber {
	return nc.set("showSteps", value)
}

// Size 表单项大小
func (nc inputNumber) Size(value string) inputNumber {
	return nc.set("size", value)
}

// Static 是否静态展示
func (nc inputNumber) Static(value bool) inputNumber {
	return nc.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (nc inputNumber) StaticClassName(value string) inputNumber {
	return nc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (nc inputNumber) StaticInputClassName(value string) inputNumber {
	return nc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (nc inputNumber) StaticLabelClassName(value string) inputNumber {
	return nc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (nc inputNumber) StaticOn(value string) inputNumber {
	return nc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (nc inputNumber) StaticPlaceholder(value string) inputNumber {
	return nc.set("staticPlaceholder", value)
}

// StaticSchema
func (nc inputNumber) StaticSchema(value string) inputNumber {
	return nc.set("staticSchema", value)
}

// Step 步长
func (nc inputNumber) Step(value string) inputNumber {
	return nc.set("step", value)
}

// Style 组件样式
func (nc inputNumber) Style(value string) inputNumber {
	return nc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (nc inputNumber) SubmitOnChange(value bool) inputNumber {
	return nc.set("submitOnChange", value)
}

// Suffix 后缀
func (nc inputNumber) Suffix(value string) inputNumber {
	return nc.set("suffix", value)
}

// TestIdBuilder
func (nc inputNumber) TestIdBuilder(value string) inputNumber {
	return nc.set("testIdBuilder", value)
}

// UnitOptions 单位列表
func (nc inputNumber) UnitOptions(value string) inputNumber {
	return nc.set("unitOptions", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (nc inputNumber) UseMobileUI(value bool) inputNumber {
	return nc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (nc inputNumber) ValidateApi(value string) inputNumber {
	return nc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (nc inputNumber) ValidateOnChange(value bool) inputNumber {
	return nc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (nc inputNumber) ValidationErrors(value string) inputNumber {
	return nc.set("validationErrors", value)
}

// Validations
func (nc inputNumber) Validations(value string) inputNumber {
	return nc.set("validations", value)
}

// Value 默认值
func (nc inputNumber) Value(value string) inputNumber {
	return nc.set("value", value)
}

// Visible 是否显示
func (nc inputNumber) Visible(value bool) inputNumber {
	return nc.set("visible", value)
}

// VisibleOn 是否显示表达式
func (nc inputNumber) VisibleOn(value string) inputNumber {
	return nc.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (nc inputNumber) Width(value string) inputNumber {
	return nc.set("width", value)
}
