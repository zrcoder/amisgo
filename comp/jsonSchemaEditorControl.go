package comp

// JSONSchemaEditorControl JSON Schema Editor 控件
//
// 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/json-schema-editor
// @version 6.7.0
type JSONSchemaEditorControl Schema

// NewJSONSchemaEditorControl 创建一个新的 JSONSchemaEditorControl 实例，并设置默认的 type
func NewJSONSchemaEditorControl() JSONSchemaEditorControl {
	return make(JSONSchemaEditorControl).set("type", "json-schema-editor")
}

// set 是一个内部方法，用于设置字段值并返回自身的引用
func (j JSONSchemaEditorControl) set(key string, value interface{}) JSONSchemaEditorControl {
	j[key] = value
	return j
}

// AdvancedSettings 自定义详情配置面板
func (j JSONSchemaEditorControl) AdvancedSettings(value string) JSONSchemaEditorControl {
	return j.set("advancedSettings", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (j JSONSchemaEditorControl) AutoFill(value string) JSONSchemaEditorControl {
	return j.set("autoFill", value)
}

// ClassName 容器 css 类名
func (j JSONSchemaEditorControl) ClassName(value string) JSONSchemaEditorControl {
	return j.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (j JSONSchemaEditorControl) ClearValueOnHidden(value bool) JSONSchemaEditorControl {
	return j.set("clearValueOnHidden", value)
}

// Definitions 可以理解为类型模板，方便快速定义复杂类型
func (j JSONSchemaEditorControl) Definitions(value string) JSONSchemaEditorControl {
	return j.set("definitions", value)
}

// Desc 描述内容
func (j JSONSchemaEditorControl) Desc(value string) JSONSchemaEditorControl {
	return j.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (j JSONSchemaEditorControl) Description(value string) JSONSchemaEditorControl {
	return j.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (j JSONSchemaEditorControl) DescriptionClassName(value string) JSONSchemaEditorControl {
	return j.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (j JSONSchemaEditorControl) Disabled(value bool) JSONSchemaEditorControl {
	return j.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (j JSONSchemaEditorControl) DisabledOn(value string) JSONSchemaEditorControl {
	return j.set("disabledOn", value)
}

// DisabledTypes 禁用类型，默认禁用了 null 类型
func (j JSONSchemaEditorControl) DisabledTypes(value string) JSONSchemaEditorControl {
	return j.set("disabledTypes", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (j JSONSchemaEditorControl) EditorSetting(value string) JSONSchemaEditorControl {
	return j.set("editorSetting", value)
}

// EnableAdvancedSetting 开启详情配置
func (j JSONSchemaEditorControl) EnableAdvancedSetting(value bool) JSONSchemaEditorControl {
	return j.set("enableAdvancedSetting", value)
}

// ExtraName 额外的字段名
func (j JSONSchemaEditorControl) ExtraName(value string) JSONSchemaEditorControl {
	return j.set("extraName", value)
}

// Hidden 是否隐藏
func (j JSONSchemaEditorControl) Hidden(value bool) JSONSchemaEditorControl {
	return j.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (j JSONSchemaEditorControl) HiddenOn(value string) JSONSchemaEditorControl {
	return j.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (j JSONSchemaEditorControl) Hint(value string) JSONSchemaEditorControl {
	return j.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (j JSONSchemaEditorControl) Horizontal(value string) JSONSchemaEditorControl {
	return j.set("horizontal", value)
}

// ID 组件唯一 id
func (j JSONSchemaEditorControl) ID(value string) JSONSchemaEditorControl {
	return j.set("id", value)
}

// InitAutoFill 初始化自动填充
func (j JSONSchemaEditorControl) InitAutoFill(value string) JSONSchemaEditorControl {
	return j.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式
func (j JSONSchemaEditorControl) Inline(value bool) JSONSchemaEditorControl {
	return j.set("inline", value)
}

// InputClassName 配置 input className
func (j JSONSchemaEditorControl) InputClassName(value string) JSONSchemaEditorControl {
	return j.set("inputClassName", value)
}

// Label 描述标题
func (j JSONSchemaEditorControl) Label(value string) JSONSchemaEditorControl {
	return j.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (j JSONSchemaEditorControl) LabelAlign(value string) JSONSchemaEditorControl {
	return j.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (j JSONSchemaEditorControl) LabelClassName(value string) JSONSchemaEditorControl {
	return j.set("labelClassName", value)
}

// LabelRemark 显示一个小图标，鼠标放上去的时候显示提示内容
func (j JSONSchemaEditorControl) LabelRemark(value string) JSONSchemaEditorControl {
	return j.set("labelRemark", value)
}

// LabelWidth label 自定义宽度，默认单位为 px
func (j JSONSchemaEditorControl) LabelWidth(value string) JSONSchemaEditorControl {
	return j.set("labelWidth", value)
}

// Mini 是否为迷你模式，会隐藏一些不必要的元素
func (j JSONSchemaEditorControl) Mini(value bool) JSONSchemaEditorControl {
	return j.set("mini", value)
}

// Mode 配置当前表单项展示模式
func (j JSONSchemaEditorControl) Mode(value string) JSONSchemaEditorControl {
	return j.set("mode", value)
}

// Name 字段名，表单提交时的 key
func (j JSONSchemaEditorControl) Name(value string) JSONSchemaEditorControl {
	return j.set("name", value)
}

// OnEvent 事件动作配置
func (j JSONSchemaEditorControl) OnEvent(value string) JSONSchemaEditorControl {
	return j.set("onEvent", value)
}

// Placeholder 各属性输入控件的占位提示文本
func (j JSONSchemaEditorControl) Placeholder(value string) JSONSchemaEditorControl {
	return j.set("placeholder", value)
}

// ReadOnly 是否只读
func (j JSONSchemaEditorControl) ReadOnly(value bool) JSONSchemaEditorControl {
	return j.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (j JSONSchemaEditorControl) ReadOnlyOn(value string) JSONSchemaEditorControl {
	return j.set("readOnlyOn", value)
}

// Remark 显示一个小图标，鼠标放上去的时候显示提示内容
func (j JSONSchemaEditorControl) Remark(value string) JSONSchemaEditorControl {
	return j.set("remark", value)
}

// Required 是否为必填
func (j JSONSchemaEditorControl) Required(value bool) JSONSchemaEditorControl {
	return j.set("required", value)
}

// RootTypeMutable 顶层是否允许修改类型
func (j JSONSchemaEditorControl) RootTypeMutable(value bool) JSONSchemaEditorControl {
	return j.set("rootTypeMutable", value)
}

// Row 行配置
func (j JSONSchemaEditorControl) Row(value string) JSONSchemaEditorControl {
	return j.set("row", value)
}

// SaveImmediately 是否立即保存 (TableColumn 中使用)
func (j JSONSchemaEditorControl) SaveImmediately(value bool) JSONSchemaEditorControl {
	return j.set("saveImmediately", value)
}

// ShowRootInfo 顶层类型信息是否隐藏
func (j JSONSchemaEditorControl) ShowRootInfo(value bool) JSONSchemaEditorControl {
	return j.set("showRootInfo", value)
}

// Size 表单项大小
func (j JSONSchemaEditorControl) Size(value string) JSONSchemaEditorControl {
	return j.set("size", value)
}

// Static 是否静态展示
func (j JSONSchemaEditorControl) Static(value bool) JSONSchemaEditorControl {
	return j.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (j JSONSchemaEditorControl) StaticClassName(value string) JSONSchemaEditorControl {
	return j.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (j JSONSchemaEditorControl) StaticInputClassName(value string) JSONSchemaEditorControl {
	return j.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (j JSONSchemaEditorControl) StaticLabelClassName(value string) JSONSchemaEditorControl {
	return j.set("staticLabelClassName", value)
}

// StaticOn 静态展示表达式
func (j JSONSchemaEditorControl) StaticOn(value string) JSONSchemaEditorControl {
	return j.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (j JSONSchemaEditorControl) StaticPlaceholder(value string) JSONSchemaEditorControl {
	return j.set("staticPlaceholder", value)
}

// StaticSchema 静态展示表单项 Schema
func (j JSONSchemaEditorControl) StaticSchema(value string) JSONSchemaEditorControl {
	return j.set("staticSchema", value)
}

// Style 组件样式
func (j JSONSchemaEditorControl) Style(value string) JSONSchemaEditorControl {
	return j.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (j JSONSchemaEditorControl) SubmitOnChange(value bool) JSONSchemaEditorControl {
	return j.set("submitOnChange", value)
}

// TestIdBuilder 测试 ID 构建器
func (j JSONSchemaEditorControl) TestIdBuilder(value string) JSONSchemaEditorControl {
	return j.set("testIdBuilder", value)
}

// UseMobileUI 组件级别用来关闭移动端样式
func (j JSONSchemaEditorControl) UseMobileUI(value bool) JSONSchemaEditorControl {
	return j.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (j JSONSchemaEditorControl) ValidateApi(value string) JSONSchemaEditorControl {
	return j.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (j JSONSchemaEditorControl) ValidateOnChange(value bool) JSONSchemaEditorControl {
	return j.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (j JSONSchemaEditorControl) ValidationErrors(value string) JSONSchemaEditorControl {
	return j.set("validationErrors", value)
}

// Validations 验证规则
func (j JSONSchemaEditorControl) Validations(value string) JSONSchemaEditorControl {
	return j.set("validations", value)
}

// Value 默认值
func (j JSONSchemaEditorControl) Value(value string) JSONSchemaEditorControl {
	return j.set("value", value)
}

// Visible 是否显示
func (j JSONSchemaEditorControl) Visible(value bool) JSONSchemaEditorControl {
	return j.set("visible", value)
}

// VisibleOn 是否显示表达式
func (j JSONSchemaEditorControl) VisibleOn(value string) JSONSchemaEditorControl {
	return j.set("visibleOn", value)
}

// Width 在 Table 中调整宽度
func (j JSONSchemaEditorControl) Width(value string) JSONSchemaEditorControl {
	return j.set("width", value)
}
