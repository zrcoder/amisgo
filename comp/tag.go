package comp

// Tag 代表一个标签组件
//
// @version 6.7.0
type Tag Schema

// NewTag 创建一个新的 Tag 实例
func NewTag() Tag {
	return Tag{}.set("type", "tag")
}

func (t Tag) set(key string, value interface{}) Tag {
	t[key] = value
	return t
}

// Checkable 设置是否为可选标签
func (t Tag) Checkable(value bool) Tag {
	return t.set("checkable", value)
}

// Checked 设置是否选中
func (t Tag) Checked(value bool) Tag {
	return t.set("checked", value)
}

// ClassName 设置标签的类名
func (t Tag) ClassName(value string) Tag {
	return t.set("className", value)
}

// Closable 设置是否展示关闭按钮
func (t Tag) Closable(value bool) Tag {
	return t.set("closable", value)
}

// CloseIcon 设置关闭图标
func (t Tag) CloseIcon(value string) Tag {
	return t.set("closeIcon", value)
}

// Color 设置标签颜色
func (t Tag) Color(value string) Tag {
	return t.set("color", value)
}

// Disabled 设置是否禁用
func (t Tag) Disabled(value bool) Tag {
	return t.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (t Tag) DisabledOn(value string) Tag {
	return t.set("disabledOn", value)
}

// DisplayMode 设置标签的显示模式
func (t Tag) DisplayMode(value string) Tag {
	return t.set("displayMode", value)
}

// EditorSetting 设置编辑器配置
func (t Tag) EditorSetting(value string) Tag {
	return t.set("editorSetting", value)
}

// Hidden 设置是否隐藏
func (t Tag) Hidden(value bool) Tag {
	return t.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (t Tag) HiddenOn(value string) Tag {
	return t.set("hiddenOn", value)
}

// Icon 设置前置图标
func (t Tag) Icon(value string) Tag {
	return t.set("icon", value)
}

// ID 设置组件唯一 ID
func (t Tag) ID(value string) Tag {
	return t.set("id", value)
}

// Label 设置标签文本内容
func (t Tag) Label(value string) Tag {
	return t.set("label", value)
}

// OnEvent 设置事件动作配置
func (t Tag) OnEvent(value string) Tag {
	return t.set("onEvent", value)
}

// Static 设置是否静态展示
func (t Tag) Static(value bool) Tag {
	return t.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (t Tag) StaticClassName(value string) Tag {
	return t.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项值类名
func (t Tag) StaticInputClassName(value string) Tag {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项标签类名
func (t Tag) StaticLabelClassName(value string) Tag {
	return t.set("staticLabelClassName", value)
}

// StaticOn 设置静态展示表达式
func (t Tag) StaticOn(value string) Tag {
	return t.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (t Tag) StaticPlaceholder(value string) Tag {
	return t.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示的 schema
func (t Tag) StaticSchema(value string) Tag {
	return t.set("staticSchema", value)
}

// Style 设置自定义样式
func (t Tag) Style(value string) Tag {
	return t.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (t Tag) TestIdBuilder(value string) Tag {
	return t.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (t Tag) Testid(value string) Tag {
	return t.set("testid", value)
}

// UseMobileUI 设置是否使用移动端样式
func (t Tag) UseMobileUI(value bool) Tag {
	return t.set("useMobileUI", value)
}

// Visible 设置是否显示
func (t Tag) Visible(value bool) Tag {
	return t.set("visible", value)
}

// VisibleOn 设置显示表达式
func (t Tag) VisibleOn(value string) Tag {
	return t.set("visibleOn", value)
}
