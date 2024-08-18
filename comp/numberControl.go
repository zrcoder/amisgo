package comp

// NumberControl 数字输入框
//
// @version 6.7.0
type NumberControl Schema

// NewNumberControl 创建一个新的 NumberControl 实例
func NewNumberControl() NumberControl {
	return NumberControl{}.set("type", "input-number")
}

// set 设置字段值
func (nc NumberControl) set(key string, value interface{}) NumberControl {
	nc[key] = value
	return nc
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (nc NumberControl) AutoFill(value string) NumberControl {
	return nc.set("autoFill", value)
}

// Big 是否是大数，如果是的话输入输出都将是字符串
func (nc NumberControl) Big(value bool) NumberControl {
	return nc.set("big", value)
}

// BorderMode 边框模式，全边框，还是半边框，或者没边框。 可选值: full | half | none
func (nc NumberControl) BorderMode(value string) NumberControl {
	return nc.set("borderMode", value)
}

// ClassName 容器 css 类名
func (nc NumberControl) ClassName(value string) NumberControl {
	return nc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (nc NumberControl) ClearValueOnHidden(value bool) NumberControl {
	return nc.set("clearValueOnHidden", value)
}

// Desc 描述
func (nc NumberControl) Desc(value string) NumberControl {
	return nc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (nc NumberControl) Description(value string) NumberControl {
	return nc.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (nc NumberControl) DescriptionClassName(value string) NumberControl {
	return nc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (nc NumberControl) Disabled(value bool) NumberControl {
	return nc.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (nc NumberControl) DisabledOn(value string) NumberControl {
	return nc.set("disabledOn", value)
}

// DisplayMode 输入框为基础输入框还是加强输入框
func (nc NumberControl) DisplayMode(value string) NumberControl {
	return nc.set("displayMode", value)
}

// EditorSetting 编辑器配置
func (nc NumberControl) EditorSetting(value string) NumberControl {
	return nc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (nc NumberControl) ExtraName(value string) NumberControl {
	return nc.set("extraName", value)
}

// Hidden 是否隐藏
func (nc NumberControl) Hidden(value bool) NumberControl {
	return nc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (nc NumberControl) HiddenOn(value string) NumberControl {
	return nc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (nc NumberControl) Hint(value string) NumberControl {
	return nc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (nc NumberControl) Horizontal(value string) NumberControl {
	return nc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (nc NumberControl) Id(value string) NumberControl {
	return nc.set("id", value)
}

// InitAutoFill 初始化自动填充
func (nc NumberControl) InitAutoFill(value string) NumberControl {
	return nc.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式
func (nc NumberControl) Inline(value bool) NumberControl {
	return nc.set("inline", value)
}

// InputClassName 配置 input className
func (nc NumberControl) InputClassName(value string) NumberControl {
	return nc.set("inputClassName", value)
}

// Keyboard 是否启用键盘行为
func (nc NumberControl) Keyboard(value bool) NumberControl {
	return nc.set("keyboard", value)
}

// KilobitSeparator 是否千分分隔
func (nc NumberControl) KilobitSeparator(value bool) NumberControl {
	return nc.set("kilobitSeparator", value)
}

// Label 描述标题
func (nc NumberControl) Label(value string) NumberControl {
	return nc.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (nc NumberControl) LabelAlign(value string) NumberControl {
	return nc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (nc NumberControl) LabelClassName(value string) NumberControl {
	return nc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (nc NumberControl) LabelRemark(value string) NumberControl {
	return nc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (nc NumberControl) LabelWidth(value string) NumberControl {
	return nc.set("labelWidth", value)
}

// Max 最大值
func (nc NumberControl) Max(value string) NumberControl {
	return nc.set("max", value)
}

// Min 最小值
func (nc NumberControl) Min(value string) NumberControl {
	return nc.set("min", value)
}

// Mode 配置当前表单项展示模式
func (nc NumberControl) Mode(value string) NumberControl {
	return nc.set("mode", value)
}

// Name 字段名，表单提交时的 key
func (nc NumberControl) Name(value string) NumberControl {
	return nc.set("name", value)
}

// OnEvent 事件动作配置
func (nc NumberControl) OnEvent(value string) NumberControl {
	return nc.set("onEvent", value)
}

// Placeholder 占位符
func (nc NumberControl) Placeholder(value string) NumberControl {
	return nc.set("placeholder", value)
}

// Precision 精度
func (nc NumberControl) Precision(value string) NumberControl {
	return nc.set("precision", value)
}

// Prefix 前缀
func (nc NumberControl) Prefix(value string) NumberControl {
	return nc.set("prefix", value)
}

// ReadOnly 只读
func (nc NumberControl) ReadOnly(value bool) NumberControl {
	return nc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (nc NumberControl) ReadOnlyOn(value string) NumberControl {
	return nc.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (nc NumberControl) Remark(value string) NumberControl {
	return nc.set("remark", value)
}

// Required 是否为必填
func (nc NumberControl) Required(value bool) NumberControl {
	return nc.set("required", value)
}

// Row
func (nc NumberControl) Row(value string) NumberControl {
	return nc.set("row", value)
}

// SaveImmediately 是否立即保存
func (nc NumberControl) SaveImmediately(value bool) NumberControl {
	return nc.set("saveImmediately", value)
}

// ShowAsPercent 用来开启百分号的展示形式
func (nc NumberControl) ShowAsPercent(value bool) NumberControl {
	return nc.set("showAsPercent", value)
}

// ShowSteps 是否显示上下点击按钮
func (nc NumberControl) ShowSteps(value bool) NumberControl {
	return nc.set("showSteps", value)
}

// Size 表单项大小
func (nc NumberControl) Size(value string) NumberControl {
	return nc.set("size", value)
}

// Static 是否静态展示
func (nc NumberControl) Static(value bool) NumberControl {
	return nc.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (nc NumberControl) StaticClassName(value string) NumberControl {
	return nc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (nc NumberControl) StaticInputClassName(value string) NumberControl {
	return nc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (nc NumberControl) StaticLabelClassName(value string) NumberControl {
	return nc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (nc NumberControl) StaticOn(value string) NumberControl {
	return nc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (nc NumberControl) StaticPlaceholder(value string) NumberControl {
	return nc.set("staticPlaceholder", value)
}

// StaticSchema
func (nc NumberControl) StaticSchema(value string) NumberControl {
	return nc.set("staticSchema", value)
}

// Step 步长
func (nc NumberControl) Step(value string) NumberControl {
	return nc.set("step", value)
}

// Style 组件样式
func (nc NumberControl) Style(value string) NumberControl {
	return nc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (nc NumberControl) SubmitOnChange(value bool) NumberControl {
	return nc.set("submitOnChange", value)
}

// Suffix 后缀
func (nc NumberControl) Suffix(value string) NumberControl {
	return nc.set("suffix", value)
}

// TestIdBuilder
func (nc NumberControl) TestIdBuilder(value string) NumberControl {
	return nc.set("testIdBuilder", value)
}

// UnitOptions 单位列表
func (nc NumberControl) UnitOptions(value string) NumberControl {
	return nc.set("unitOptions", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (nc NumberControl) UseMobileUI(value bool) NumberControl {
	return nc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (nc NumberControl) ValidateApi(value string) NumberControl {
	return nc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (nc NumberControl) ValidateOnChange(value bool) NumberControl {
	return nc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (nc NumberControl) ValidationErrors(value string) NumberControl {
	return nc.set("validationErrors", value)
}

// Validations
func (nc NumberControl) Validations(value string) NumberControl {
	return nc.set("validations", value)
}

// Value 默认值
func (nc NumberControl) Value(value string) NumberControl {
	return nc.set("value", value)
}

// Visible 是否显示
func (nc NumberControl) Visible(value bool) NumberControl {
	return nc.set("visible", value)
}

// VisibleOn 是否显示表达式
func (nc NumberControl) VisibleOn(value string) NumberControl {
	return nc.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (nc NumberControl) Width(value string) NumberControl {
	return nc.set("width", value)
}
