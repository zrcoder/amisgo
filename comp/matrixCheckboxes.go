package comp

// matrixCheckboxes 选择控件，适合做权限勾选
//
// @version 6.7.0
type matrixCheckboxes schema

// MatrixCheckboxes 创建一个新的 MatrixControl 实例
func MatrixCheckboxes() matrixCheckboxes {
	return matrixCheckboxes{}.set("type", "matrix-checkboxes")
}

func (mc matrixCheckboxes) set(key string, value any) matrixCheckboxes {
	mc[key] = value
	return mc
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (mc matrixCheckboxes) AutoFill(value string) matrixCheckboxes {
	mc.set("autoFill", value)
	return mc
}

// ClassName 容器 css 类名
func (mc matrixCheckboxes) ClassName(value string) matrixCheckboxes {
	mc.set("className", value)
	return mc
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (mc matrixCheckboxes) ClearValueOnHidden(value bool) matrixCheckboxes {
	mc.set("clearValueOnHidden", value)
	return mc
}

// Columns
func (mc matrixCheckboxes) Columns(value ...any) matrixCheckboxes {
	mc.set("columns", value)
	return mc
}

// Desc
func (mc matrixCheckboxes) Desc(value string) matrixCheckboxes {
	mc.set("desc", value)
	return mc
}

// Description 描述内容，支持 Html 片段
func (mc matrixCheckboxes) Description(value string) matrixCheckboxes {
	mc.set("description", value)
	return mc
}

// DescriptionClassName 配置描述上的 className
func (mc matrixCheckboxes) DescriptionClassName(value string) matrixCheckboxes {
	mc.set("descriptionClassName", value)
	return mc
}

// Disabled 是否禁用
func (mc matrixCheckboxes) Disabled(value bool) matrixCheckboxes {
	mc.set("disabled", value)
	return mc
}

// DisabledOn 是否禁用表达式
func (mc matrixCheckboxes) DisabledOn(value string) matrixCheckboxes {
	mc.set("disabledOn", value)
	return mc
}

// EditorSetting 编辑器配置，运行时可以忽略
func (mc matrixCheckboxes) EditorSetting(value string) matrixCheckboxes {
	mc.set("editorSetting", value)
	return mc
}

// ExtraName 额外的字段名
func (mc matrixCheckboxes) ExtraName(value string) matrixCheckboxes {
	mc.set("extraName", value)
	return mc
}

// Hidden 是否隐藏
func (mc matrixCheckboxes) Hidden(value bool) matrixCheckboxes {
	mc.set("hidden", value)
	return mc
}

// HiddenOn 是否隐藏表达式
func (mc matrixCheckboxes) HiddenOn(value string) matrixCheckboxes {
	mc.set("hiddenOn", value)
	return mc
}

// Hint 输入提示，聚焦的时候显示
func (mc matrixCheckboxes) Hint(value string) matrixCheckboxes {
	mc.set("hint", value)
	return mc
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (mc matrixCheckboxes) Horizontal(value string) matrixCheckboxes {
	mc.set("horizontal", value)
	return mc
}

// ID 组件唯一 id，主要用于日志采集
func (mc matrixCheckboxes) ID(value string) matrixCheckboxes {
	mc.set("id", value)
	return mc
}

// InitAutoFill
func (mc matrixCheckboxes) InitAutoFill(value string) matrixCheckboxes {
	mc.set("initAutoFill", value)
	return mc
}

// Inline 表单 control 是否为 inline 模式
func (mc matrixCheckboxes) Inline(value bool) matrixCheckboxes {
	mc.set("inline", value)
	return mc
}

// InputClassName 配置 input className
func (mc matrixCheckboxes) InputClassName(value string) matrixCheckboxes {
	mc.set("inputClassName", value)
	return mc
}

// Label 描述标题
func (mc matrixCheckboxes) Label(value string) matrixCheckboxes {
	mc.set("label", value)
	return mc
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (mc matrixCheckboxes) LabelAlign(value string) matrixCheckboxes {
	mc.set("labelAlign", value)
	return mc
}

// LabelClassName 配置 label className
func (mc matrixCheckboxes) LabelClassName(value string) matrixCheckboxes {
	mc.set("labelClassName", value)
	return mc
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起
func (mc matrixCheckboxes) LabelRemark(value string) matrixCheckboxes {
	mc.set("labelRemark", value)
	return mc
}

// LabelWidth label自定义宽度，默认单位为px
func (mc matrixCheckboxes) LabelWidth(value string) matrixCheckboxes {
	mc.set("labelWidth", value)
	return mc
}

// Mode 配置当前表单项展示模式
func (mc matrixCheckboxes) Mode(value string) matrixCheckboxes {
	mc.set("mode", value)
	return mc
}

// Multiple 配置singleSelectMode时设置为false
func (mc matrixCheckboxes) Multiple(value bool) matrixCheckboxes {
	mc.set("multiple", value)
	return mc
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (mc matrixCheckboxes) Name(value string) matrixCheckboxes {
	mc.set("name", value)
	return mc
}

// OnEvent 事件动作配置
func (mc matrixCheckboxes) OnEvent(value any) matrixCheckboxes {
	mc.set("onEvent", value)
	return mc
}

// Placeholder 占位符
func (mc matrixCheckboxes) Placeholder(value string) matrixCheckboxes {
	mc.set("placeholder", value)
	return mc
}

// ReadOnly 是否只读
func (mc matrixCheckboxes) ReadOnly(value bool) matrixCheckboxes {
	mc.set("readOnly", value)
	return mc
}

// ReadOnlyOn 只读条件
func (mc matrixCheckboxes) ReadOnlyOn(value string) matrixCheckboxes {
	mc.set("readOnlyOn", value)
	return mc
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容
func (mc matrixCheckboxes) Remark(value string) matrixCheckboxes {
	mc.set("remark", value)
	return mc
}

// Required 是否为必填
func (mc matrixCheckboxes) Required(value bool) matrixCheckboxes {
	mc.set("required", value)
	return mc
}

// Row
func (mc matrixCheckboxes) Row(value string) matrixCheckboxes {
	mc.set("row", value)
	return mc
}

// RowLabel 行标题说明
func (mc matrixCheckboxes) RowLabel(value string) matrixCheckboxes {
	mc.set("rowLabel", value)
	return mc
}

// Rows
func (mc matrixCheckboxes) Rows(value string) matrixCheckboxes {
	mc.set("rows", value)
	return mc
}

// SaveImmediately 是否立即保存 (TableColumn中使用)
func (mc matrixCheckboxes) SaveImmediately(value bool) matrixCheckboxes {
	mc.set("saveImmediately", value)
	return mc
}

// SingleSelectMode 设置单选模式，multiple为false时有效
func (mc matrixCheckboxes) SingleSelectMode(value bool) matrixCheckboxes {
	mc.set("singleSelectMode", value)
	return mc
}

// Size 表单项大小
func (mc matrixCheckboxes) Size(value string) matrixCheckboxes {
	mc.set("size", value)
	return mc
}

// Source 可用来通过 API 拉取 options
func (mc matrixCheckboxes) Source(value string) matrixCheckboxes {
	mc.set("source", value)
	return mc
}

// Static 是否静态展示
func (mc matrixCheckboxes) Static(value bool) matrixCheckboxes {
	mc.set("static", value)
	return mc
}

// StaticClassName 静态展示表单项类名
func (mc matrixCheckboxes) StaticClassName(value string) matrixCheckboxes {
	mc.set("staticClassName", value)
	return mc
}

// StaticInputClassName 静态展示表单项Value类名
func (mc matrixCheckboxes) StaticInputClassName(value string) matrixCheckboxes {
	mc.set("staticInputClassName", value)
	return mc
}

// StaticLabelClassName 静态展示表单项Label类名
func (mc matrixCheckboxes) StaticLabelClassName(value string) matrixCheckboxes {
	mc.set("staticLabelClassName", value)
	return mc
}

// StaticOn 是否静态展示表达式
func (mc matrixCheckboxes) StaticOn(value string) matrixCheckboxes {
	mc.set("staticOn", value)
	return mc
}

// StaticPlaceholder 静态展示空值占位
func (mc matrixCheckboxes) StaticPlaceholder(value string) matrixCheckboxes {
	mc.set("staticPlaceholder", value)
	return mc
}

// StaticSchema
func (mc matrixCheckboxes) StaticSchema(value string) matrixCheckboxes {
	mc.set("staticSchema", value)
	return mc
}

// Style 组件样式
func (mc matrixCheckboxes) Style(value any) matrixCheckboxes {
	mc.set("style", value)
	return mc
}

// SubmitOnChange 当修改完的时候是否提交表单
func (mc matrixCheckboxes) SubmitOnChange(value bool) matrixCheckboxes {
	mc.set("submitOnChange", value)
	return mc
}

// TestIdBuilder
func (mc matrixCheckboxes) TestIdBuilder(value string) matrixCheckboxes {
	mc.set("testIdBuilder", value)
	return mc
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (mc matrixCheckboxes) UseMobileUI(value bool) matrixCheckboxes {
	mc.set("useMobileUI", value)
	return mc
}

// ValidateApi 远端校验表单项接口
func (mc matrixCheckboxes) ValidateApi(value string) matrixCheckboxes {
	mc.set("validateApi", value)
	return mc
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (mc matrixCheckboxes) ValidateOnChange(value bool) matrixCheckboxes {
	mc.set("validateOnChange", value)
	return mc
}

// ValidationErrors 验证失败的提示信息
func (mc matrixCheckboxes) ValidationErrors(value string) matrixCheckboxes {
	mc.set("validationErrors", value)
	return mc
}

// Validations
func (mc matrixCheckboxes) Validations(value string) matrixCheckboxes {
	mc.set("validations", value)
	return mc
}

// Value 默认值，切记只能是静态值
func (mc matrixCheckboxes) Value(value string) matrixCheckboxes {
	mc.set("value", value)
	return mc
}

// Visible 是否显示
func (mc matrixCheckboxes) Visible(value bool) matrixCheckboxes {
	mc.set("visible", value)
	return mc
}

// VisibleOn 是否显示表达式
func (mc matrixCheckboxes) VisibleOn(value string) matrixCheckboxes {
	mc.set("visibleOn", value)
	return mc
}

// Width 在 Table 中调整宽度
func (mc matrixCheckboxes) Width(value string) matrixCheckboxes {
	mc.set("width", value)
	return mc
}
