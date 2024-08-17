package comp

// FieldSetControl 表单项集合
type FieldSetControl Schema

// NewFieldSetControl 创建一个新的 FieldSetControl 实例，并设置默认的 type
func NewFieldSetControl() FieldSetControl {
	return make(FieldSetControl).set("type", "fieldset").set("titlePosition", "top")
}

func (f FieldSetControl) set(key string, value interface{}) FieldSetControl {
	f[key] = value
	return f
}

// AutoFill 自动填充
func (f FieldSetControl) AutoFill(value string) FieldSetControl {
	return f.set("autoFill", value)
}

// Body 内容区域
func (f FieldSetControl) Body(value ...interface{}) FieldSetControl {
	return f.set("body", value)
}

// BodyClassName 配置 Body 容器 className
func (f FieldSetControl) BodyClassName(value string) FieldSetControl {
	return f.set("bodyClassName", value)
}

// ClassName 容器 css 类名
func (f FieldSetControl) ClassName(value string) FieldSetControl {
	return f.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (f FieldSetControl) ClearValueOnHidden(value bool) FieldSetControl {
	return f.set("clearValueOnHidden", value)
}

// Collapsable 是否可折叠
func (f FieldSetControl) Collapsable(value bool) FieldSetControl {
	return f.set("collapsable", value)
}

// CollapseHeader 收起的标题
func (f FieldSetControl) CollapseHeader(value string) FieldSetControl {
	return f.set("collapseHeader", value)
}

// CollapseTitle 收起的标题
func (f FieldSetControl) CollapseTitle(value string) FieldSetControl {
	return f.set("collapseTitle", value)
}

// Collapsed 默认是否折叠
func (f FieldSetControl) Collapsed(value bool) FieldSetControl {
	return f.set("collapsed", value)
}

// Desc 描述
func (f FieldSetControl) Desc(value string) FieldSetControl {
	return f.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (f FieldSetControl) Description(value string) FieldSetControl {
	return f.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (f FieldSetControl) DescriptionClassName(value string) FieldSetControl {
	return f.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (f FieldSetControl) Disabled(value bool) FieldSetControl {
	return f.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (f FieldSetControl) DisabledOn(value string) FieldSetControl {
	return f.set("disabledOn", value)
}

// DivideLine 标题内容分割线
func (f FieldSetControl) DivideLine(value bool) FieldSetControl {
	return f.set("divideLine", value)
}

// EditorSetting 编辑器配置
func (f FieldSetControl) EditorSetting(value string) FieldSetControl {
	return f.set("editorSetting", value)
}

// ExpandIcon 自定义切换图标
func (f FieldSetControl) ExpandIcon(value string) FieldSetControl {
	return f.set("expandIcon", value)
}

// ExtraName 额外的字段名
func (f FieldSetControl) ExtraName(value string) FieldSetControl {
	return f.set("extraName", value)
}

// Header 标题
func (f FieldSetControl) Header(value string) FieldSetControl {
	return f.set("header", value)
}

// HeaderPosition 标题展示位置
func (f FieldSetControl) HeaderPosition(value string) FieldSetControl {
	return f.set("headerPosition", value)
}

// HeadingClassName 标题 CSS 类名
func (f FieldSetControl) HeadingClassName(value string) FieldSetControl {
	return f.set("headingClassName", value)
}

// Hidden 是否隐藏
func (f FieldSetControl) Hidden(value bool) FieldSetControl {
	return f.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (f FieldSetControl) HiddenOn(value string) FieldSetControl {
	return f.set("hiddenOn", value)
}

// Hint 输入提示
func (f FieldSetControl) Hint(value string) FieldSetControl {
	return f.set("hint", value)
}

// Horizontal 水平布局具体的左右分配
func (f FieldSetControl) Horizontal(value string) FieldSetControl {
	return f.set("horizontal", value)
}

// ID 组件唯一 id
func (f FieldSetControl) ID(value string) FieldSetControl {
	return f.set("id", value)
}

// InitAutoFill 初始化自动填充
func (f FieldSetControl) InitAutoFill(value string) FieldSetControl {
	return f.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式
func (f FieldSetControl) Inline(value bool) FieldSetControl {
	return f.set("inline", value)
}

// InputClassName 配置 input className
func (f FieldSetControl) InputClassName(value string) FieldSetControl {
	return f.set("inputClassName", value)
}

// Key 标识
func (f FieldSetControl) Key(value string) FieldSetControl {
	return f.set("key", value)
}

// Label 描述标题
func (f FieldSetControl) Label(value string) FieldSetControl {
	return f.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (f FieldSetControl) LabelAlign(value string) FieldSetControl {
	return f.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (f FieldSetControl) LabelClassName(value string) FieldSetControl {
	return f.set("labelClassName", value)
}

// LabelRemark 显示一个小图标，鼠标放上去的时候显示提示内容
func (f FieldSetControl) LabelRemark(value string) FieldSetControl {
	return f.set("labelRemark", value)
}

// LabelWidth label 自定义宽度
func (f FieldSetControl) LabelWidth(value string) FieldSetControl {
	return f.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式
func (f FieldSetControl) Mode(value string) FieldSetControl {
	return f.set("mode", value)
}

// MountOnEnter 点开时才加载内容
func (f FieldSetControl) MountOnEnter(value bool) FieldSetControl {
	return f.set("mountOnEnter", value)
}

// Name 字段名
func (f FieldSetControl) Name(value string) FieldSetControl {
	return f.set("name", value)
}

// OnEvent 事件动作配置
func (f FieldSetControl) OnEvent(value string) FieldSetControl {
	return f.set("onEvent", value)
}

// Placeholder 占位符
func (f FieldSetControl) Placeholder(value string) FieldSetControl {
	return f.set("placeholder", value)
}

// ReadOnly 是否只读
func (f FieldSetControl) ReadOnly(value bool) FieldSetControl {
	return f.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (f FieldSetControl) ReadOnlyOn(value string) FieldSetControl {
	return f.set("readOnlyOn", value)
}

// Remark 显示一个小图标，鼠标放上去的时候显示提示内容
func (f FieldSetControl) Remark(value string) FieldSetControl {
	return f.set("remark", value)
}

// Required 是否为必填
func (f FieldSetControl) Required(value bool) FieldSetControl {
	return f.set("required", value)
}

// Row
func (f FieldSetControl) Row(value string) FieldSetControl {
	return f.set("row", value)
}

// SaveImmediately 是否立即保存
func (f FieldSetControl) SaveImmediately(value bool) FieldSetControl {
	return f.set("saveImmediately", value)
}

// ShowArrow 图标是否展示
func (f FieldSetControl) ShowArrow(value bool) FieldSetControl {
	return f.set("showArrow", value)
}

// Size 控件大小
func (f FieldSetControl) Size(value string) FieldSetControl {
	return f.set("size", value)
}

// Static 是否静态展示
func (f FieldSetControl) Static(value bool) FieldSetControl {
	return f.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (f FieldSetControl) StaticClassName(value string) FieldSetControl {
	return f.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (f FieldSetControl) StaticInputClassName(value string) FieldSetControl {
	return f.set("staticInputClassName", value)
}

// SubmitOnChange 是否每次更改都提交
func (f FieldSetControl) SubmitOnChange(value bool) FieldSetControl {
	return f.set("submitOnChange", value)
}

// Title 标题
func (f FieldSetControl) Title(value string) FieldSetControl {
	return f.set("title", value)
}

// TitlePosition 标题展示位置
func (f FieldSetControl) TitlePosition(value string) FieldSetControl {
	return f.set("titlePosition", value)
}

// Visible 组件显示条件
func (f FieldSetControl) Visible(value bool) FieldSetControl {
	return f.set("visible", value)
}

// VisibleOn 组件显示条件表达式
func (f FieldSetControl) VisibleOn(value string) FieldSetControl {
	return f.set("visibleOn", value)
}
