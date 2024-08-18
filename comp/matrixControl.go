package comp

// MatrixControl 选择控件，适合做权限勾选
//
// @version 6.7.0
type MatrixControl Schema

// NewMatrixControl 创建一个新的 MatrixControl 实例
func NewMatrixControl() MatrixControl {
	return MatrixControl{}.set("type", "matrix-checkboxes")
}

func (mc MatrixControl) set(key string, value interface{}) MatrixControl {
	mc[key] = value
	return mc
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (mc MatrixControl) AutoFill(value string) MatrixControl {
	mc.set("autoFill", value)
	return mc
}

// ClassName 容器 css 类名
func (mc MatrixControl) ClassName(value string) MatrixControl {
	mc.set("className", value)
	return mc
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (mc MatrixControl) ClearValueOnHidden(value bool) MatrixControl {
	mc.set("clearValueOnHidden", value)
	return mc
}

// Columns
func (mc MatrixControl) Columns(value string) MatrixControl {
	mc.set("columns", value)
	return mc
}

// Desc
func (mc MatrixControl) Desc(value string) MatrixControl {
	mc.set("desc", value)
	return mc
}

// Description 描述内容，支持 Html 片段
func (mc MatrixControl) Description(value string) MatrixControl {
	mc.set("description", value)
	return mc
}

// DescriptionClassName 配置描述上的 className
func (mc MatrixControl) DescriptionClassName(value string) MatrixControl {
	mc.set("descriptionClassName", value)
	return mc
}

// Disabled 是否禁用
func (mc MatrixControl) Disabled(value bool) MatrixControl {
	mc.set("disabled", value)
	return mc
}

// DisabledOn 是否禁用表达式
func (mc MatrixControl) DisabledOn(value string) MatrixControl {
	mc.set("disabledOn", value)
	return mc
}

// EditorSetting 编辑器配置，运行时可以忽略
func (mc MatrixControl) EditorSetting(value string) MatrixControl {
	mc.set("editorSetting", value)
	return mc
}

// ExtraName 额外的字段名
func (mc MatrixControl) ExtraName(value string) MatrixControl {
	mc.set("extraName", value)
	return mc
}

// Hidden 是否隐藏
func (mc MatrixControl) Hidden(value bool) MatrixControl {
	mc.set("hidden", value)
	return mc
}

// HiddenOn 是否隐藏表达式
func (mc MatrixControl) HiddenOn(value string) MatrixControl {
	mc.set("hiddenOn", value)
	return mc
}

// Hint 输入提示，聚焦的时候显示
func (mc MatrixControl) Hint(value string) MatrixControl {
	mc.set("hint", value)
	return mc
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (mc MatrixControl) Horizontal(value string) MatrixControl {
	mc.set("horizontal", value)
	return mc
}

// ID 组件唯一 id，主要用于日志采集
func (mc MatrixControl) ID(value string) MatrixControl {
	mc.set("id", value)
	return mc
}

// InitAutoFill
func (mc MatrixControl) InitAutoFill(value string) MatrixControl {
	mc.set("initAutoFill", value)
	return mc
}

// Inline 表单 control 是否为 inline 模式
func (mc MatrixControl) Inline(value bool) MatrixControl {
	mc.set("inline", value)
	return mc
}

// InputClassName 配置 input className
func (mc MatrixControl) InputClassName(value string) MatrixControl {
	mc.set("inputClassName", value)
	return mc
}

// Label 描述标题
func (mc MatrixControl) Label(value string) MatrixControl {
	mc.set("label", value)
	return mc
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (mc MatrixControl) LabelAlign(value string) MatrixControl {
	mc.set("labelAlign", value)
	return mc
}

// LabelClassName 配置 label className
func (mc MatrixControl) LabelClassName(value string) MatrixControl {
	mc.set("labelClassName", value)
	return mc
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
func (mc MatrixControl) LabelRemark(value string) MatrixControl {
	mc.set("labelRemark", value)
	return mc
}

// LabelWidth label自定义宽度，默认单位为px
func (mc MatrixControl) LabelWidth(value string) MatrixControl {
	mc.set("labelWidth", value)
	return mc
}

// Mode 配置当前表单项展示模式
func (mc MatrixControl) Mode(value string) MatrixControl {
	mc.set("mode", value)
	return mc
}

// Multiple 配置singleSelectMode时设置为false
func (mc MatrixControl) Multiple(value bool) MatrixControl {
	mc.set("multiple", value)
	return mc
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (mc MatrixControl) Name(value string) MatrixControl {
	mc.set("name", value)
	return mc
}

// OnEvent 事件动作配置
func (mc MatrixControl) OnEvent(value string) MatrixControl {
	mc.set("onEvent", value)
	return mc
}

// Placeholder 占位符
func (mc MatrixControl) Placeholder(value string) MatrixControl {
	mc.set("placeholder", value)
	return mc
}

// ReadOnly 是否只读
func (mc MatrixControl) ReadOnly(value bool) MatrixControl {
	mc.set("readOnly", value)
	return mc
}

// ReadOnlyOn 只读条件
func (mc MatrixControl) ReadOnlyOn(value string) MatrixControl {
	mc.set("readOnlyOn", value)
	return mc
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (mc MatrixControl) Remark(value string) MatrixControl {
	mc.set("remark", value)
	return mc
}

// Required 是否为必填
func (mc MatrixControl) Required(value bool) MatrixControl {
	mc.set("required", value)
	return mc
}

// Row
func (mc MatrixControl) Row(value string) MatrixControl {
	mc.set("row", value)
	return mc
}

// RowLabel 行标题说明
func (mc MatrixControl) RowLabel(value string) MatrixControl {
	mc.set("rowLabel", value)
	return mc
}

// Rows
func (mc MatrixControl) Rows(value string) MatrixControl {
	mc.set("rows", value)
	return mc
}

// SaveImmediately 是否立即保存 (TableColumn中使用)
func (mc MatrixControl) SaveImmediately(value bool) MatrixControl {
	mc.set("saveImmediately", value)
	return mc
}

// SingleSelectMode 设置单选模式，multiple为false时有效
func (mc MatrixControl) SingleSelectMode(value bool) MatrixControl {
	mc.set("singleSelectMode", value)
	return mc
}

// Size 表单项大小
func (mc MatrixControl) Size(value string) MatrixControl {
	mc.set("size", value)
	return mc
}

// Source 可用来通过 API 拉取 options
func (mc MatrixControl) Source(value string) MatrixControl {
	mc.set("source", value)
	return mc
}

// Static 是否静态展示
func (mc MatrixControl) Static(value bool) MatrixControl {
	mc.set("static", value)
	return mc
}

// StaticClassName 静态展示表单项类名
func (mc MatrixControl) StaticClassName(value string) MatrixControl {
	mc.set("staticClassName", value)
	return mc
}

// StaticInputClassName 静态展示表单项Value类名
func (mc MatrixControl) StaticInputClassName(value string) MatrixControl {
	mc.set("staticInputClassName", value)
	return mc
}

// StaticLabelClassName 静态展示表单项Label类名
func (mc MatrixControl) StaticLabelClassName(value string) MatrixControl {
	mc.set("staticLabelClassName", value)
	return mc
}

// StaticOn 是否静态展示表达式
func (mc MatrixControl) StaticOn(value string) MatrixControl {
	mc.set("staticOn", value)
	return mc
}

// StaticPlaceholder 静态展示空值占位
func (mc MatrixControl) StaticPlaceholder(value string) MatrixControl {
	mc.set("staticPlaceholder", value)
	return mc
}

// StaticSchema
func (mc MatrixControl) StaticSchema(value string) MatrixControl {
	mc.set("staticSchema", value)
	return mc
}

// Style 组件样式
func (mc MatrixControl) Style(value string) MatrixControl {
	mc.set("style", value)
	return mc
}

// SubmitOnChange 当修改完的时候是否提交表单
func (mc MatrixControl) SubmitOnChange(value bool) MatrixControl {
	mc.set("submitOnChange", value)
	return mc
}

// TestIdBuilder
func (mc MatrixControl) TestIdBuilder(value string) MatrixControl {
	mc.set("testIdBuilder", value)
	return mc
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (mc MatrixControl) UseMobileUI(value bool) MatrixControl {
	mc.set("useMobileUI", value)
	return mc
}

// ValidateApi 远端校验表单项接口
func (mc MatrixControl) ValidateApi(value string) MatrixControl {
	mc.set("validateApi", value)
	return mc
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (mc MatrixControl) ValidateOnChange(value bool) MatrixControl {
	mc.set("validateOnChange", value)
	return mc
}

// ValidationErrors 验证失败的提示信息
func (mc MatrixControl) ValidationErrors(value string) MatrixControl {
	mc.set("validationErrors", value)
	return mc
}

// Validations
func (mc MatrixControl) Validations(value string) MatrixControl {
	mc.set("validations", value)
	return mc
}

// Value 默认值，切记只能是静态值
func (mc MatrixControl) Value(value string) MatrixControl {
	mc.set("value", value)
	return mc
}

// Visible 是否显示
func (mc MatrixControl) Visible(value bool) MatrixControl {
	mc.set("visible", value)
	return mc
}

// VisibleOn 是否显示表达式
func (mc MatrixControl) VisibleOn(value string) MatrixControl {
	mc.set("visibleOn", value)
	return mc
}

// Width 在 Table 中调整宽度
func (mc MatrixControl) Width(value string) MatrixControl {
	mc.set("width", value)
	return mc
}
