package comp

// editor 表示代码编辑器
type editor Schema

func Editor() editor {
	return make(editor).set("type", "editor")
}

func (ec editor) set(key string, value any) editor {
	ec[key] = value
	return ec
}

// AllowFullscreen 是否展示全屏模式开关
func (ec editor) AllowFullscreen(value bool) editor {
	return ec.set("allowFullscreen", value)
}

// AutoFill 自动填充
func (ec editor) AutoFill(value string) editor {
	return ec.set("autoFill", value)
}

// ClassName 容器 CSS 类名
func (ec editor) ClassName(value string) editor {
	return ec.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (ec editor) ClearValueOnHidden(value bool) editor {
	return ec.set("clearValueOnHidden", value)
}

// Desc 描述内容
func (ec editor) Desc(value string) editor {
	return ec.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (ec editor) Description(value string) editor {
	return ec.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (ec editor) DescriptionClassName(value string) editor {
	return ec.set("descriptionClassName", value)
}

// Disabled 是否禁用/只读
func (ec editor) Disabled(value bool) editor {
	return ec.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (ec editor) DisabledOn(value string) editor {
	return ec.set("disabledOn", value)
}

// EditorDidMount 获取编辑器底层实例
func (ec editor) EditorDidMount(value string) editor {
	return ec.set("editorDidMount", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (ec editor) EditorSetting(value string) editor {
	return ec.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (ec editor) ExtraName(value string) editor {
	return ec.set("extraName", value)
}

// Hidden 是否隐藏
func (ec editor) Hidden(value bool) editor {
	return ec.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (ec editor) HiddenOn(value string) editor {
	return ec.set("hiddenOn", value)
}

// Hint 输入提示
func (ec editor) Hint(value string) editor {
	return ec.set("hint", value)
}

// Horizontal 水平布局配置
func (ec editor) Horizontal(value string) editor {
	return ec.set("horizontal", value)
}

// ID 组件唯一 ID
func (ec editor) ID(value string) editor {
	return ec.set("id", value)
}

// InitAutoFill 初始化自动填充
func (ec editor) InitAutoFill(value string) editor {
	return ec.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式
func (ec editor) Inline(value bool) editor {
	return ec.set("inline", value)
}

// InputClassName 配置 input className
func (ec editor) InputClassName(value string) editor {
	return ec.set("inputClassName", value)
}

// Label 描述标题
func (ec editor) Label(value string) editor {
	return ec.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (ec editor) LabelAlign(value string) editor {
	return ec.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (ec editor) LabelClassName(value string) editor {
	return ec.set("labelClassName", value)
}

// LabelRemark 显示一个小图标，鼠标放上去的时候显示提示内容
func (ec editor) LabelRemark(value string) editor {
	return ec.set("labelRemark", value)
}

// LabelWidth label 自定义宽度，默认单位为 px
func (ec editor) LabelWidth(value string) editor {
	return ec.set("labelWidth", value)
}

// Language 语言类型
func (ec editor) Language(value string) editor {
	return ec.set("language", value)
}

// Mode 配置当前表单项展示模式
func (ec editor) Mode(value string) editor {
	return ec.set("mode", value)
}

// Name 字段名，表单提交时的 key
func (ec editor) Name(value string) editor {
	return ec.set("name", value)
}

// OnEvent 事件动作配置
func (ec editor) OnEvent(value any) editor {
	return ec.set("onEvent", value)
}

// Placeholder 占位符
func (ec editor) Placeholder(value string) editor {
	return ec.set("placeholder", value)
}

// ReadOnly 是否只读
func (ec editor) ReadOnly(value bool) editor {
	return ec.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (ec editor) ReadOnlyOn(value string) editor {
	return ec.set("readOnlyOn", value)
}

// Remark 显示一个小图标，鼠标放上去的时候显示提示内容
func (ec editor) Remark(value string) editor {
	return ec.set("remark", value)
}

// Required 是否为必填
func (ec editor) Required(value bool) editor {
	return ec.set("required", value)
}

// Row 配置行数
func (ec editor) Row(value string) editor {
	return ec.set("row", value)
}

// SaveImmediately 是否立即保存
func (ec editor) SaveImmediately(value bool) editor {
	return ec.set("saveImmediately", value)
}

// Size 编辑器高度，取值可以是 md、lg、xl、xxl, 默认 md
func (ec editor) Size(value string) editor {
	return ec.set("size", value)
}

// Static 是否静态展示
func (ec editor) Static(value bool) editor {
	return ec.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (ec editor) StaticClassName(value string) editor {
	return ec.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (ec editor) StaticInputClassName(value string) editor {
	return ec.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (ec editor) StaticLabelClassName(value string) editor {
	return ec.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (ec editor) StaticOn(value string) editor {
	return ec.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (ec editor) StaticPlaceholder(value string) editor {
	return ec.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (ec editor) StaticSchema(value string) editor {
	return ec.set("staticSchema", value)
}

// Style 组件样式
func (ec editor) Style(value any) editor {
	return ec.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (ec editor) SubmitOnChange(value bool) editor {
	return ec.set("submitOnChange", value)
}

// TestIdBuilder 测试 ID 构建器
func (ec editor) TestIdBuilder(value string) editor {
	return ec.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (ec editor) UseMobileUI(value bool) editor {
	return ec.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (ec editor) ValidateApi(value string) editor {
	return ec.set("validateApi", value)
}

// ValidateOnChange 不设置时，表单项每次修改都会触发重新验证
func (ec editor) ValidateOnChange(value bool) editor {
	return ec.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (ec editor) ValidationErrors(value string) editor {
	return ec.set("validationErrors", value)
}

// Validations 验证规则
func (ec editor) Validations(value string) editor {
	return ec.set("validations", value)
}

// Value 默认值
func (ec editor) Value(value string) editor {
	return ec.set("value", value)
}

// Visible 是否显示
func (ec editor) Visible(value bool) editor {
	return ec.set("visible", value)
}

// VisibleOn 是否显示表达式
func (ec editor) VisibleOn(value string) editor {
	return ec.set("visibleOn", value)
}

// Width 在 Table 中调整宽度
func (ec editor) Width(value string) editor {
	return ec.set("width", value)
}

// Options monaco 编辑器的其它配置，比如是否显示行号等，
// 请参考https://microsoft.github.io/monaco-editor/docs.html#interfaces/editor.IEditorOptions.html，
// 不过无法设置 readOnly，只读模式需要使用 disabled: true
func (e editor) Options(value Schema) editor {
	return e.set("options", value)
}
