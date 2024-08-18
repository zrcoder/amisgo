package comp

// RangeControl 组件的范围控制
//
// @author  slowlyo
// @version 6.7.0
type RangeControl Schema

// NewRangeControl 创建一个新的 RangeControl 实例
func NewRangeControl() RangeControl {
	return RangeControl{}.set("type", "input-range")
}

func (rc RangeControl) set(key string, value interface{}) RangeControl {
	rc[key] = value
	return rc
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (rc RangeControl) AutoFill(value string) RangeControl {
	return rc.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。)
func (rc RangeControl) ClassName(value string) RangeControl {
	return rc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (rc RangeControl) ClearValueOnHidden(value bool) RangeControl {
	return rc.set("clearValueOnHidden", value)
}

// Clearable 输入框是否可清除
func (rc RangeControl) Clearable(value bool) RangeControl {
	return rc.set("clearable", value)
}

// Delimiter 分隔符
func (rc RangeControl) Delimiter(value string) RangeControl {
	return rc.set("delimiter", value)
}

// Desc 描述内容
func (rc RangeControl) Desc(value string) RangeControl {
	return rc.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (rc RangeControl) Description(value string) RangeControl {
	return rc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (rc RangeControl) DescriptionClassName(value string) RangeControl {
	return rc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (rc RangeControl) Disabled(value bool) RangeControl {
	return rc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (rc RangeControl) DisabledOn(value string) RangeControl {
	return rc.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (rc RangeControl) EditorSetting(value string) RangeControl {
	return rc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (rc RangeControl) ExtraName(value string) RangeControl {
	return rc.set("extraName", value)
}

// Hidden 是否隐藏
func (rc RangeControl) Hidden(value bool) RangeControl {
	return rc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (rc RangeControl) HiddenOn(value string) RangeControl {
	return rc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (rc RangeControl) Hint(value string) RangeControl {
	return rc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。
func (rc RangeControl) Horizontal(value string) RangeControl {
	return rc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (rc RangeControl) Id(value string) RangeControl {
	return rc.set("id", value)
}

// InitAutoFill
func (rc RangeControl) InitAutoFill(value string) RangeControl {
	return rc.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (rc RangeControl) Inline(value bool) RangeControl {
	return rc.set("inline", value)
}

// InputClassName 配置 input className
func (rc RangeControl) InputClassName(value string) RangeControl {
	return rc.set("inputClassName", value)
}

// JoinValues 是否通过分隔符连接
func (rc RangeControl) JoinValues(value bool) RangeControl {
	return rc.set("joinValues", value)
}

// Label 描述标题
func (rc RangeControl) Label(value string) RangeControl {
	return rc.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (rc RangeControl) LabelAlign(value string) RangeControl {
	return rc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (rc RangeControl) LabelClassName(value string) RangeControl {
	return rc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
func (rc RangeControl) LabelRemark(value string) RangeControl {
	return rc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (rc RangeControl) LabelWidth(value string) RangeControl {
	return rc.set("labelWidth", value)
}

// Marks 刻度
func (rc RangeControl) Marks(value string) RangeControl {
	return rc.set("marks", value)
}

// Max 最大值
func (rc RangeControl) Max(value string) RangeControl {
	return rc.set("max", value)
}

// Min 最小值
func (rc RangeControl) Min(value string) RangeControl {
	return rc.set("min", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (rc RangeControl) Mode(value string) RangeControl {
	return rc.set("mode", value)
}

// Multiple 是否为双滑块
func (rc RangeControl) Multiple(value bool) RangeControl {
	return rc.set("multiple", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (rc RangeControl) Name(value string) RangeControl {
	return rc.set("name", value)
}

// OnEvent 事件动作配置
func (rc RangeControl) OnEvent(value string) RangeControl {
	return rc.set("onEvent", value)
}

// Parts 分割块数
func (rc RangeControl) Parts(value string) RangeControl {
	return rc.set("parts", value)
}

// Placeholder 占位符
func (rc RangeControl) Placeholder(value string) RangeControl {
	return rc.set("placeholder", value)
}

// ReadOnly 是否只读
func (rc RangeControl) ReadOnly(value bool) RangeControl {
	return rc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (rc RangeControl) ReadOnlyOn(value string) RangeControl {
	return rc.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (rc RangeControl) Remark(value string) RangeControl {
	return rc.set("remark", value)
}

// Required 是否为必填
func (rc RangeControl) Required(value bool) RangeControl {
	return rc.set("required", value)
}

// Row
func (rc RangeControl) Row(value string) RangeControl {
	return rc.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (rc RangeControl) SaveImmediately(value bool) RangeControl {
	return rc.set("saveImmediately", value)
}

// ShowInput 是否展示输入框
func (rc RangeControl) ShowInput(value bool) RangeControl {
	return rc.set("showInput", value)
}

// ShowSteps 是否展示步长
func (rc RangeControl) ShowSteps(value bool) RangeControl {
	return rc.set("showSteps", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (rc RangeControl) Size(value string) RangeControl {
	return rc.set("size", value)
}

// Static 是否静态展示
func (rc RangeControl) Static(value bool) RangeControl {
	return rc.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (rc RangeControl) StaticClassName(value string) RangeControl {
	return rc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (rc RangeControl) StaticInputClassName(value string) RangeControl {
	return rc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (rc RangeControl) StaticLabelClassName(value string) RangeControl {
	return rc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (rc RangeControl) StaticOn(value string) RangeControl {
	return rc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (rc RangeControl) StaticPlaceholder(value string) RangeControl {
	return rc.set("staticPlaceholder", value)
}

// StaticSchema
func (rc RangeControl) StaticSchema(value string) RangeControl {
	return rc.set("staticSchema", value)
}

// Step 步长
func (rc RangeControl) Step(value string) RangeControl {
	return rc.set("step", value)
}

// Style 组件样式
func (rc RangeControl) Style(value string) RangeControl {
	return rc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (rc RangeControl) SubmitOnChange(value bool) RangeControl {
	return rc.set("submitOnChange", value)
}

// TestIdBuilder
func (rc RangeControl) TestIdBuilder(value string) RangeControl {
	return rc.set("testIdBuilder", value)
}

// TooltipPlacement 标签方向 (标签方向) 可选值: auto | top | right | bottom | left
func (rc RangeControl) TooltipPlacement(value string) RangeControl {
	return rc.set("tooltipPlacement", value)
}

// TooltipVisible 是否展示标签
func (rc RangeControl) TooltipVisible(value bool) RangeControl {
	return rc.set("tooltipVisible", value)
}

// Unit 单位
func (rc RangeControl) Unit(value string) RangeControl {
	return rc.set("unit", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (rc RangeControl) UseMobileUI(value bool) RangeControl {
	return rc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (rc RangeControl) ValidateApi(value string) RangeControl {
	return rc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (rc RangeControl) ValidateOnChange(value bool) RangeControl {
	return rc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (rc RangeControl) ValidationErrors(value string) RangeControl {
	return rc.set("validationErrors", value)
}

// Validations
func (rc RangeControl) Validations(value string) RangeControl {
	return rc.set("validations", value)
}

// Value 滑块值 (Range 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/range)
func (rc RangeControl) Value(value string) RangeControl {
	return rc.set("value", value)
}

// Visible 是否显示
func (rc RangeControl) Visible(value bool) RangeControl {
	return rc.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (rc RangeControl) VisibleOn(value string) RangeControl {
	return rc.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (rc RangeControl) Width(value string) RangeControl {
	return rc.set("width", value)
}
