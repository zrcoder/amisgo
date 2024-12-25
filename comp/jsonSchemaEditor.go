package comp

// jsonSchemaEditor JSON Schema Editor 控件
// 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/json-schema-editor

type jsonSchemaEditor Schema

func JsonSchemaEditor() jsonSchemaEditor {
	return make(jsonSchemaEditor).set("type", "json-schema-editor")
}

// set 是一个内部方法，用于设置字段值并返回自身的引用
func (j jsonSchemaEditor) set(key string, value any) jsonSchemaEditor {
	j[key] = value
	return j
}

// AdvancedSettings 自定义详情配置面板
func (j jsonSchemaEditor) AdvancedSettings(value string) jsonSchemaEditor {
	return j.set("advancedSettings", value)
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内
func (j jsonSchemaEditor) AutoFill(value string) jsonSchemaEditor {
	return j.set("autoFill", value)
}

// ClassName 容器 css 类名
func (j jsonSchemaEditor) ClassName(value string) jsonSchemaEditor {
	return j.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (j jsonSchemaEditor) ClearValueOnHidden(value bool) jsonSchemaEditor {
	return j.set("clearValueOnHidden", value)
}

// Definitions 可以理解为类型模板，方便快速定义复杂类型
func (j jsonSchemaEditor) Definitions(value string) jsonSchemaEditor {
	return j.set("definitions", value)
}

// Desc 描述内容
func (j jsonSchemaEditor) Desc(value string) jsonSchemaEditor {
	return j.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (j jsonSchemaEditor) Description(value string) jsonSchemaEditor {
	return j.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (j jsonSchemaEditor) DescriptionClassName(value string) jsonSchemaEditor {
	return j.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (j jsonSchemaEditor) Disabled(value bool) jsonSchemaEditor {
	return j.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (j jsonSchemaEditor) DisabledOn(value string) jsonSchemaEditor {
	return j.set("disabledOn", value)
}

// DisabledTypes 禁用类型，默认禁用了 null 类型
func (j jsonSchemaEditor) DisabledTypes(value string) jsonSchemaEditor {
	return j.set("disabledTypes", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (j jsonSchemaEditor) EditorSetting(value string) jsonSchemaEditor {
	return j.set("editorSetting", value)
}

// EnableAdvancedSetting 开启详情配置
func (j jsonSchemaEditor) EnableAdvancedSetting(value bool) jsonSchemaEditor {
	return j.set("enableAdvancedSetting", value)
}

// ExtraName 额外的字段名
func (j jsonSchemaEditor) ExtraName(value string) jsonSchemaEditor {
	return j.set("extraName", value)
}

// Hidden 是否隐藏
func (j jsonSchemaEditor) Hidden(value bool) jsonSchemaEditor {
	return j.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (j jsonSchemaEditor) HiddenOn(value string) jsonSchemaEditor {
	return j.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (j jsonSchemaEditor) Hint(value string) jsonSchemaEditor {
	return j.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配
func (j jsonSchemaEditor) Horizontal(value string) jsonSchemaEditor {
	return j.set("horizontal", value)
}

// ID 组件唯一 id
func (j jsonSchemaEditor) ID(value string) jsonSchemaEditor {
	return j.set("id", value)
}

// InitAutoFill 初始化自动填充
func (j jsonSchemaEditor) InitAutoFill(value string) jsonSchemaEditor {
	return j.set("initAutoFill", value)
}

// Inline 表单  是否为 inline 模式
func (j jsonSchemaEditor) Inline(value bool) jsonSchemaEditor {
	return j.set("inline", value)
}

// InputClassName 配置 input className
func (j jsonSchemaEditor) InputClassName(value string) jsonSchemaEditor {
	return j.set("inputClassName", value)
}

// Label 描述标题
func (j jsonSchemaEditor) Label(value string) jsonSchemaEditor {
	return j.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (j jsonSchemaEditor) LabelAlign(value string) jsonSchemaEditor {
	return j.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (j jsonSchemaEditor) LabelClassName(value string) jsonSchemaEditor {
	return j.set("labelClassName", value)
}

// LabelRemark 显示一个小图标，鼠标放上去的时候显示提示内容
func (j jsonSchemaEditor) LabelRemark(value string) jsonSchemaEditor {
	return j.set("labelRemark", value)
}

// LabelWidth label 自定义宽度，默认单位为 px
func (j jsonSchemaEditor) LabelWidth(value string) jsonSchemaEditor {
	return j.set("labelWidth", value)
}

// Mini 是否为迷你模式，会隐藏一些不必要的元素
func (j jsonSchemaEditor) Mini(value bool) jsonSchemaEditor {
	return j.set("mini", value)
}

// Mode 配置当前表单项展示模式
func (j jsonSchemaEditor) Mode(value string) jsonSchemaEditor {
	return j.set("mode", value)
}

// Name 字段名，表单提交时的 key
func (j jsonSchemaEditor) Name(value string) jsonSchemaEditor {
	return j.set("name", value)
}

// OnEvent 事件动作配置
func (j jsonSchemaEditor) OnEvent(value any) jsonSchemaEditor {
	return j.set("onEvent", value)
}

// Placeholder 各属性输入控件的占位提示文本
func (j jsonSchemaEditor) Placeholder(value string) jsonSchemaEditor {
	return j.set("placeholder", value)
}

// ReadOnly 是否只读
func (j jsonSchemaEditor) ReadOnly(value bool) jsonSchemaEditor {
	return j.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (j jsonSchemaEditor) ReadOnlyOn(value string) jsonSchemaEditor {
	return j.set("readOnlyOn", value)
}

// Remark 显示一个小图标，鼠标放上去的时候显示提示内容
func (j jsonSchemaEditor) Remark(value string) jsonSchemaEditor {
	return j.set("remark", value)
}

// Required 是否为必填
func (j jsonSchemaEditor) Required(value bool) jsonSchemaEditor {
	return j.set("required", value)
}

// RootTypeMutable 顶层是否允许修改类型
func (j jsonSchemaEditor) RootTypeMutable(value bool) jsonSchemaEditor {
	return j.set("rootTypeMutable", value)
}

// Row 行配置
func (j jsonSchemaEditor) Row(value string) jsonSchemaEditor {
	return j.set("row", value)
}

// SaveImmediately 是否立即保存 (TableColumn 中使用)
func (j jsonSchemaEditor) SaveImmediately(value bool) jsonSchemaEditor {
	return j.set("saveImmediately", value)
}

// ShowRootInfo 顶层类型信息是否隐藏
func (j jsonSchemaEditor) ShowRootInfo(value bool) jsonSchemaEditor {
	return j.set("showRootInfo", value)
}

// Size 表单项大小
func (j jsonSchemaEditor) Size(value string) jsonSchemaEditor {
	return j.set("size", value)
}

// Static 是否静态展示
func (j jsonSchemaEditor) Static(value bool) jsonSchemaEditor {
	return j.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (j jsonSchemaEditor) StaticClassName(value string) jsonSchemaEditor {
	return j.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (j jsonSchemaEditor) StaticInputClassName(value string) jsonSchemaEditor {
	return j.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (j jsonSchemaEditor) StaticLabelClassName(value string) jsonSchemaEditor {
	return j.set("staticLabelClassName", value)
}

// StaticOn 静态展示表达式
func (j jsonSchemaEditor) StaticOn(value string) jsonSchemaEditor {
	return j.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (j jsonSchemaEditor) StaticPlaceholder(value string) jsonSchemaEditor {
	return j.set("staticPlaceholder", value)
}

// StaticSchema 静态展示表单项 Schema
func (j jsonSchemaEditor) StaticSchema(value string) jsonSchemaEditor {
	return j.set("staticSchema", value)
}

// Style 组件样式
func (j jsonSchemaEditor) Style(value any) jsonSchemaEditor {
	return j.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (j jsonSchemaEditor) SubmitOnChange(value bool) jsonSchemaEditor {
	return j.set("submitOnChange", value)
}

// TestIdBuilder 测试 ID 构建器
func (j jsonSchemaEditor) TestIdBuilder(value string) jsonSchemaEditor {
	return j.set("testIdBuilder", value)
}

// UseMobileUI 组件级别用来关闭移动端样式
func (j jsonSchemaEditor) UseMobileUI(value bool) jsonSchemaEditor {
	return j.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (j jsonSchemaEditor) ValidateApi(value string) jsonSchemaEditor {
	return j.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证
func (j jsonSchemaEditor) ValidateOnChange(value bool) jsonSchemaEditor {
	return j.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (j jsonSchemaEditor) ValidationErrors(value string) jsonSchemaEditor {
	return j.set("validationErrors", value)
}

// Validations 验证规则
func (j jsonSchemaEditor) Validations(value string) jsonSchemaEditor {
	return j.set("validations", value)
}

// Value 默认值
func (j jsonSchemaEditor) Value(value string) jsonSchemaEditor {
	return j.set("value", value)
}

// Visible 是否显示
func (j jsonSchemaEditor) Visible(value bool) jsonSchemaEditor {
	return j.set("visible", value)
}

// VisibleOn 是否显示表达式
func (j jsonSchemaEditor) VisibleOn(value string) jsonSchemaEditor {
	return j.set("visibleOn", value)
}

// Width 在 Table 中调整宽度
func (j jsonSchemaEditor) Width(value string) jsonSchemaEditor {
	return j.set("width", value)
}
