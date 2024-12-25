package comp

// multilineText 多行文本控件
type multilineText Schema

// MultilineText 创建一个新的 MultilineText 实例
func MultilineText() multilineText {
	return multilineText{}.set("type", "multiline-text")
}

// set 方法用于设置属性并返回自身
func (m multilineText) set(key string, value any) multilineText {
	m[key] = value
	return m
}

// ClassName 设置容器 css 类名
func (m multilineText) ClassName(value string) multilineText {
	return m.set("className", value)
}

// CollapseButtonText 设置收起按钮文本
func (m multilineText) CollapseButtonText(value string) multilineText {
	return m.set("collapseButtonText", value)
}

// Disabled 设置是否禁用
func (m multilineText) Disabled(value bool) multilineText {
	return m.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (m multilineText) DisabledOn(value string) multilineText {
	return m.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置
func (m multilineText) EditorSetting(value string) multilineText {
	return m.set("editorSetting", value)
}

// ExpendButtonText 设置展开按钮文本
func (m multilineText) ExpendButtonText(value string) multilineText {
	return m.set("expendButtonText", value)
}

// Hidden 设置是否隐藏
func (m multilineText) Hidden(value bool) multilineText {
	return m.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (m multilineText) HiddenOn(value string) multilineText {
	return m.set("hiddenOn", value)
}

// Id 设置组件唯一 id
func (m multilineText) ID(value string) multilineText {
	return m.set("id", value)
}

// MaxRows 设置最大行数
func (m multilineText) MaxRows(value string) multilineText {
	return m.set("maxRows", value)
}

// OnEvent 设置事件动作配置
func (m multilineText) OnEvent(value any) multilineText {
	return m.set("onEvent", value)
}

// Static 设置是否静态展示
func (m multilineText) Static(value bool) multilineText {
	return m.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (m multilineText) StaticClassName(value string) multilineText {
	return m.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项 Value 类名
func (m multilineText) StaticInputClassName(value string) multilineText {
	return m.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项 Label 类名
func (m multilineText) StaticLabelClassName(value string) multilineText {
	return m.set("staticLabelClassName", value)
}

// StaticOn 设置静态展示表达式
func (m multilineText) StaticOn(value string) multilineText {
	return m.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (m multilineText) StaticPlaceholder(value string) multilineText {
	return m.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示 schema
func (m multilineText) StaticSchema(value string) multilineText {
	return m.set("staticSchema", value)
}

// Style 设置组件样式
func (m multilineText) Style(value any) multilineText {
	return m.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (m multilineText) TestIdBuilder(value string) multilineText {
	return m.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (m multilineText) Testid(value string) multilineText {
	return m.set("testid", value)
}

// Text 设置文本内容
func (m multilineText) Text(value string) multilineText {
	return m.set("text", value)
}

// UseMobileUI 设置是否关闭移动端样式
func (m multilineText) UseMobileUI(value bool) multilineText {
	return m.set("useMobileUI", value)
}

// Visible 设置是否显示
func (m multilineText) Visible(value bool) multilineText {
	return m.set("visible", value)
}

// VisibleOn 设置显示表达式
func (m multilineText) VisibleOn(value string) multilineText {
	return m.set("visibleOn", value)
}
