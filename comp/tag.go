package comp

// tag 代表一个标签组件

type tag Schema

// Tag 创建一个新的 Tag 实例
func Tag() tag {
	return tag{}.set("type", "tag")
}

func (t tag) set(key string, value any) tag {
	t[key] = value
	return t
}

// Checkable 设置是否为可选标签
func (t tag) Checkable(value bool) tag {
	return t.set("checkable", value)
}

// Checked 设置是否选中
func (t tag) Checked(value bool) tag {
	return t.set("checked", value)
}

// ClassName 设置标签的类名
func (t tag) ClassName(value string) tag {
	return t.set("className", value)
}

// Closable 设置是否展示关闭按钮
func (t tag) Closable(value bool) tag {
	return t.set("closable", value)
}

// CloseIcon 设置关闭图标
func (t tag) CloseIcon(value string) tag {
	return t.set("closeIcon", value)
}

// Color 设置标签颜色
func (t tag) Color(value string) tag {
	return t.set("color", value)
}

// Disabled 设置是否禁用
func (t tag) Disabled(value bool) tag {
	return t.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (t tag) DisabledOn(value string) tag {
	return t.set("disabledOn", value)
}

// DisplayMode 设置标签的显示模式
func (t tag) DisplayMode(value string) tag {
	return t.set("displayMode", value)
}

// EditorSetting 设置编辑器配置
func (t tag) EditorSetting(value string) tag {
	return t.set("editorSetting", value)
}

// Hidden 设置是否隐藏
func (t tag) Hidden(value bool) tag {
	return t.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (t tag) HiddenOn(value string) tag {
	return t.set("hiddenOn", value)
}

// Icon 设置前置图标
func (t tag) Icon(value string) tag {
	return t.set("icon", value)
}

// ID 设置组件唯一 ID
func (t tag) ID(value string) tag {
	return t.set("id", value)
}

// Label 设置标签文本内容
func (t tag) Label(value string) tag {
	return t.set("label", value)
}

// OnEvent 设置事件动作配置
func (t tag) OnEvent(value any) tag {
	return t.set("onEvent", value)
}

// Static 设置是否静态展示
func (t tag) Static(value bool) tag {
	return t.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (t tag) StaticClassName(value string) tag {
	return t.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项值类名
func (t tag) StaticInputClassName(value string) tag {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项标签类名
func (t tag) StaticLabelClassName(value string) tag {
	return t.set("staticLabelClassName", value)
}

// StaticOn 设置静态展示表达式
func (t tag) StaticOn(value string) tag {
	return t.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (t tag) StaticPlaceholder(value string) tag {
	return t.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示的 schema
func (t tag) StaticSchema(value string) tag {
	return t.set("staticSchema", value)
}

// Style 设置自定义样式
func (t tag) Style(value any) tag {
	return t.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (t tag) TestIdBuilder(value string) tag {
	return t.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (t tag) Testid(value string) tag {
	return t.set("testid", value)
}

// UseMobileUI 设置是否使用移动端样式
func (t tag) UseMobileUI(value bool) tag {
	return t.set("useMobileUI", value)
}

// Visible 设置是否显示
func (t tag) Visible(value bool) tag {
	return t.set("visible", value)
}

// VisibleOn 设置显示表达式
func (t tag) VisibleOn(value string) tag {
	return t.set("visibleOn", value)
}
