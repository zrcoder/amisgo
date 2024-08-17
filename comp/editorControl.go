package comp

// EditorControl 表示代码编辑器
type EditorControl Schema

// NewEditorControl 创建一个新的 EditorControl 实例，并设置默认的 type 为 'editor'
func NewEditorControl() EditorControl {
	return make(EditorControl).set("type", "editor")
}

func (ec EditorControl) set(key string, value interface{}) EditorControl {
	ec[key] = value
	return ec
}

// AllowFullscreen 是否展示全屏模式开关
func (ec EditorControl) AllowFullscreen(value bool) EditorControl {
	return ec.set("allowFullscreen", value)
}

// AutoFill 自动填充
func (ec EditorControl) AutoFill(value string) EditorControl {
	return ec.set("autoFill", value)
}

// ClassName 容器 CSS 类名
func (ec EditorControl) ClassName(value string) EditorControl {
	return ec.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (ec EditorControl) ClearValueOnHidden(value bool) EditorControl {
	return ec.set("clearValueOnHidden", value)
}

// Desc 描述内容
func (ec EditorControl) Desc(value string) EditorControl {
	return ec.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (ec EditorControl) Description(value string) EditorControl {
	return ec.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (ec EditorControl) DescriptionClassName(value string) EditorControl {
	return ec.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (ec EditorControl) Disabled(value bool) EditorControl {
	return ec.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (ec EditorControl) DisabledOn(value string) EditorControl {
	return ec.set("disabledOn", value)
}

// EditorDidMount 获取编辑器底层实例
func (ec EditorControl) EditorDidMount(value string) EditorControl {
	return ec.set("editorDidMount", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (ec EditorControl) EditorSetting(value string) EditorControl {
	return ec.set("editorSetting", value)
}

// ExtraName 额外的字段名
func (ec EditorControl) ExtraName(value string) EditorControl {
	return ec.set("extraName", value)
}

// Hidden 是否隐藏
func (ec EditorControl) Hidden(value bool) EditorControl {
	return ec.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (ec EditorControl) HiddenOn(value string) EditorControl {
	return ec.set("hiddenOn", value)
}

// Hint 输入提示
func (ec EditorControl) Hint(value string) EditorControl {
	return ec.set("hint", value)
}

// Horizontal 水平布局配置
func (ec EditorControl) Horizontal(value string) EditorControl {
	return ec.set("horizontal", value)
}

// ID 组件唯一 ID
func (ec EditorControl) ID(value string) EditorControl {
	return ec.set("id", value)
}

// InitAutoFill 初始化自动填充
func (ec EditorControl) InitAutoFill(value string) EditorControl {
	return ec.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式
func (ec EditorControl) Inline(value bool) EditorControl {
	return ec.set("inline", value)
}

// InputClassName 配置 input className
func (ec EditorControl) InputClassName(value string) EditorControl {
	return ec.set("inputClassName", value)
}

// Label 描述标题
func (ec EditorControl) Label(value string) EditorControl {
	return ec.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (ec EditorControl) LabelAlign(value string) EditorControl {
	return ec.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (ec EditorControl) LabelClassName(value string) EditorControl {
	return ec.set("labelClassName", value)
}

// LabelRemark 显示一个小图标，鼠标放上去的时候显示提示内容
func (ec EditorControl) LabelRemark(value string) EditorControl {
	return ec.set("labelRemark", value)
}

// LabelWidth label 自定义宽度，默认单位为 px
func (ec EditorControl) LabelWidth(value string) EditorControl {
	return ec.set("labelWidth", value)
}

// Language 语言类型
func (ec EditorControl) Language(value string) EditorControl {
	return ec.set("language", value)
}

// Mode 配置当前表单项展示模式
func (ec EditorControl) Mode(value string) EditorControl {
	return ec.set("mode", value)
}

// Name 字段名，表单提交时的 key
func (ec EditorControl) Name(value string) EditorControl {
	return ec.set("name", value)
}

// OnEvent 事件动作配置
func (ec EditorControl) OnEvent(value string) EditorControl {
	return ec.set("onEvent", value)
}

// Placeholder 占位符
func (ec EditorControl) Placeholder(value string) EditorControl {
	return ec.set("placeholder", value)
}

// ReadOnly 是否只读
func (ec EditorControl) ReadOnly(value bool) EditorControl {
	return ec.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (ec EditorControl) ReadOnlyOn(value string) EditorControl {
	return ec.set("readOnlyOn", value)
}

// Remark 显示一个小图标，鼠标放上去的时候显示提示内容
func (ec EditorControl) Remark(value string) EditorControl {
	return ec.set("remark", value)
}

// Required 是否为必填
func (ec EditorControl) Required(value bool) EditorControl {
	return ec.set("required", value)
}

// Row 配置行数
func (ec EditorControl) Row(value string) EditorControl {
	return ec.set("row", value)
}

// SaveImmediately 是否立即保存
func (ec EditorControl) SaveImmediately(value bool) EditorControl {
	return ec.set("saveImmediately", value)
}

// Size 编辑器大小
func (ec EditorControl) Size(value string) EditorControl {
	return ec.set("size", value)
}

// Static 是否静态展示
func (ec EditorControl) Static(value bool) EditorControl {
	return ec.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (ec EditorControl) StaticClassName(value string) EditorControl {
	return ec.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (ec EditorControl) StaticInputClassName(value string) EditorControl {
	return ec.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项 Label 类名
func (ec EditorControl) StaticLabelClassName(value string) EditorControl {
	return ec.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (ec EditorControl) StaticOn(value string) EditorControl {
	return ec.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (ec EditorControl) StaticPlaceholder(value string) EditorControl {
	return ec.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (ec EditorControl) StaticSchema(value string) EditorControl {
	return ec.set("staticSchema", value)
}

// Style 组件样式
func (ec EditorControl) Style(value string) EditorControl {
	return ec.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单
func (ec EditorControl) SubmitOnChange(value bool) EditorControl {
	return ec.set("submitOnChange", value)
}

// TestIdBuilder 测试 ID 构建器
func (ec EditorControl) TestIdBuilder(value string) EditorControl {
	return ec.set("testIdBuilder", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (ec EditorControl) UseMobileUI(value bool) EditorControl {
	return ec.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (ec EditorControl) ValidateApi(value string) EditorControl {
	return ec.set("validateApi", value)
}

// ValidateOnChange 不设置时，表单项每次修改都会触发重新验证
func (ec EditorControl) ValidateOnChange(value bool) EditorControl {
	return ec.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (ec EditorControl) ValidationErrors(value string) EditorControl {
	return ec.set("validationErrors", value)
}

// Validations 验证规则
func (ec EditorControl) Validations(value string) EditorControl {
	return ec.set("validations", value)
}

// Value 默认值
func (ec EditorControl) Value(value string) EditorControl {
	return ec.set("value", value)
}

// Visible 是否显示
func (ec EditorControl) Visible(value bool) EditorControl {
	return ec.set("visible", value)
}

// VisibleOn 是否显示表达式
func (ec EditorControl) VisibleOn(value string) EditorControl {
	return ec.set("visibleOn", value)
}

// Width 在 Table 中调整宽度
func (ec EditorControl) Width(value string) EditorControl {
	return ec.set("width", value)
}
