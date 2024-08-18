package comp

// fieldSet 表单项集合
type fieldSet schema

func FieldSet() fieldSet {
	return make(fieldSet).set("type", "fieldset").set("titlePosition", "top")
}

func (f fieldSet) set(key string, value interface{}) fieldSet {
	f[key] = value
	return f
}

// AutoFill 自动填充
func (f fieldSet) AutoFill(value string) fieldSet {
	return f.set("autoFill", value)
}

// Body 内容区域
func (f fieldSet) Body(value ...interface{}) fieldSet {
	return f.set("body", value)
}

// BodyClassName 配置 Body 容器 className
func (f fieldSet) BodyClassName(value string) fieldSet {
	return f.set("bodyClassName", value)
}

// ClassName 容器 css 类名
func (f fieldSet) ClassName(value string) fieldSet {
	return f.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值
func (f fieldSet) ClearValueOnHidden(value bool) fieldSet {
	return f.set("clearValueOnHidden", value)
}

// Collapsable 是否可折叠
func (f fieldSet) Collapsable(value bool) fieldSet {
	return f.set("collapsable", value)
}

// CollapseHeader 收起的标题
func (f fieldSet) CollapseHeader(value string) fieldSet {
	return f.set("collapseHeader", value)
}

// CollapseTitle 收起的标题
func (f fieldSet) CollapseTitle(value string) fieldSet {
	return f.set("collapseTitle", value)
}

// Collapsed 默认是否折叠
func (f fieldSet) Collapsed(value bool) fieldSet {
	return f.set("collapsed", value)
}

// Desc 描述
func (f fieldSet) Desc(value string) fieldSet {
	return f.set("desc", value)
}

// Description 描述内容，支持 Html 片段
func (f fieldSet) Description(value string) fieldSet {
	return f.set("description", value)
}

// DescriptionClassName 配置描述上的 className
func (f fieldSet) DescriptionClassName(value string) fieldSet {
	return f.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (f fieldSet) Disabled(value bool) fieldSet {
	return f.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (f fieldSet) DisabledOn(value string) fieldSet {
	return f.set("disabledOn", value)
}

// DivideLine 标题内容分割线
func (f fieldSet) DivideLine(value bool) fieldSet {
	return f.set("divideLine", value)
}

// EditorSetting 编辑器配置
func (f fieldSet) EditorSetting(value string) fieldSet {
	return f.set("editorSetting", value)
}

// ExpandIcon 自定义切换图标
func (f fieldSet) ExpandIcon(value string) fieldSet {
	return f.set("expandIcon", value)
}

// ExtraName 额外的字段名
func (f fieldSet) ExtraName(value string) fieldSet {
	return f.set("extraName", value)
}

// Header 标题
func (f fieldSet) Header(value string) fieldSet {
	return f.set("header", value)
}

// HeaderPosition 标题展示位置
func (f fieldSet) HeaderPosition(value string) fieldSet {
	return f.set("headerPosition", value)
}

// HeadingClassName 标题 CSS 类名
func (f fieldSet) HeadingClassName(value string) fieldSet {
	return f.set("headingClassName", value)
}

// Hidden 是否隐藏
func (f fieldSet) Hidden(value bool) fieldSet {
	return f.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (f fieldSet) HiddenOn(value string) fieldSet {
	return f.set("hiddenOn", value)
}

// Hint 输入提示
func (f fieldSet) Hint(value string) fieldSet {
	return f.set("hint", value)
}

// Horizontal 水平布局具体的左右分配
func (f fieldSet) Horizontal(value string) fieldSet {
	return f.set("horizontal", value)
}

// ID 组件唯一 id
func (f fieldSet) ID(value string) fieldSet {
	return f.set("id", value)
}

// InitAutoFill 初始化自动填充
func (f fieldSet) InitAutoFill(value string) fieldSet {
	return f.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式
func (f fieldSet) Inline(value bool) fieldSet {
	return f.set("inline", value)
}

// InputClassName 配置 input className
func (f fieldSet) InputClassName(value string) fieldSet {
	return f.set("inputClassName", value)
}

// Key 标识
func (f fieldSet) Key(value string) fieldSet {
	return f.set("key", value)
}

// Label 描述标题
func (f fieldSet) Label(value string) fieldSet {
	return f.set("label", value)
}

// LabelAlign 描述标题对齐方式
func (f fieldSet) LabelAlign(value string) fieldSet {
	return f.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (f fieldSet) LabelClassName(value string) fieldSet {
	return f.set("labelClassName", value)
}

// LabelRemark 显示一个小图标，鼠标放上去的时候显示提示内容
func (f fieldSet) LabelRemark(value string) fieldSet {
	return f.set("labelRemark", value)
}

// LabelWidth label 自定义宽度
func (f fieldSet) LabelWidth(value string) fieldSet {
	return f.set("labelWidth", value)
}

// Mode 配置当前表单项展示模式
func (f fieldSet) Mode(value string) fieldSet {
	return f.set("mode", value)
}

// MountOnEnter 点开时才加载内容
func (f fieldSet) MountOnEnter(value bool) fieldSet {
	return f.set("mountOnEnter", value)
}

// Name 字段名
func (f fieldSet) Name(value string) fieldSet {
	return f.set("name", value)
}

// OnEvent 事件动作配置
func (f fieldSet) OnEvent(value string) fieldSet {
	return f.set("onEvent", value)
}

// Placeholder 占位符
func (f fieldSet) Placeholder(value string) fieldSet {
	return f.set("placeholder", value)
}

// ReadOnly 是否只读
func (f fieldSet) ReadOnly(value bool) fieldSet {
	return f.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (f fieldSet) ReadOnlyOn(value string) fieldSet {
	return f.set("readOnlyOn", value)
}

// Remark 显示一个小图标，鼠标放上去的时候显示提示内容
func (f fieldSet) Remark(value string) fieldSet {
	return f.set("remark", value)
}

// Required 是否为必填
func (f fieldSet) Required(value bool) fieldSet {
	return f.set("required", value)
}

// Row
func (f fieldSet) Row(value string) fieldSet {
	return f.set("row", value)
}

// SaveImmediately 是否立即保存
func (f fieldSet) SaveImmediately(value bool) fieldSet {
	return f.set("saveImmediately", value)
}

// ShowArrow 图标是否展示
func (f fieldSet) ShowArrow(value bool) fieldSet {
	return f.set("showArrow", value)
}

// Size 控件大小
func (f fieldSet) Size(value string) fieldSet {
	return f.set("size", value)
}

// Static 是否静态展示
func (f fieldSet) Static(value bool) fieldSet {
	return f.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (f fieldSet) StaticClassName(value string) fieldSet {
	return f.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项 Value 类名
func (f fieldSet) StaticInputClassName(value string) fieldSet {
	return f.set("staticInputClassName", value)
}

// SubmitOnChange 是否每次更改都提交
func (f fieldSet) SubmitOnChange(value bool) fieldSet {
	return f.set("submitOnChange", value)
}

// Title 标题
func (f fieldSet) Title(value string) fieldSet {
	return f.set("title", value)
}

// TitlePosition 标题展示位置
func (f fieldSet) TitlePosition(value string) fieldSet {
	return f.set("titlePosition", value)
}

// Visible 组件显示条件
func (f fieldSet) Visible(value bool) fieldSet {
	return f.set("visible", value)
}

// VisibleOn 组件显示条件表达式
func (f fieldSet) VisibleOn(value string) fieldSet {
	return f.set("visibleOn", value)
}
