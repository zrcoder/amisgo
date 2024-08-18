package comp

// MultilineText 多行文本控件
type MultilineText Schema

// NewMultilineText 创建一个新的 MultilineText 实例
func NewMultilineText() MultilineText {
	return MultilineText{}.set("type", "multiline-text")
}

// set 方法用于设置属性并返回自身
func (m MultilineText) set(key string, value interface{}) MultilineText {
	m[key] = value
	return m
}

// ClassName 设置容器 css 类名
func (m MultilineText) ClassName(value string) MultilineText {
	return m.set("className", value)
}

// CollapseButtonText 设置收起按钮文本
func (m MultilineText) CollapseButtonText(value string) MultilineText {
	return m.set("collapseButtonText", value)
}

// Disabled 设置是否禁用
func (m MultilineText) Disabled(value bool) MultilineText {
	return m.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (m MultilineText) DisabledOn(value string) MultilineText {
	return m.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置
func (m MultilineText) EditorSetting(value string) MultilineText {
	return m.set("editorSetting", value)
}

// ExpendButtonText 设置展开按钮文本
func (m MultilineText) ExpendButtonText(value string) MultilineText {
	return m.set("expendButtonText", value)
}

// Hidden 设置是否隐藏
func (m MultilineText) Hidden(value bool) MultilineText {
	return m.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (m MultilineText) HiddenOn(value string) MultilineText {
	return m.set("hiddenOn", value)
}

// Id 设置组件唯一 id
func (m MultilineText) Id(value string) MultilineText {
	return m.set("id", value)
}

// MaxRows 设置最大行数
func (m MultilineText) MaxRows(value string) MultilineText {
	return m.set("maxRows", value)
}

// OnEvent 设置事件动作配置
func (m MultilineText) OnEvent(value string) MultilineText {
	return m.set("onEvent", value)
}

// Static 设置是否静态展示
func (m MultilineText) Static(value bool) MultilineText {
	return m.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (m MultilineText) StaticClassName(value string) MultilineText {
	return m.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项 Value 类名
func (m MultilineText) StaticInputClassName(value string) MultilineText {
	return m.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项 Label 类名
func (m MultilineText) StaticLabelClassName(value string) MultilineText {
	return m.set("staticLabelClassName", value)
}

// StaticOn 设置静态展示表达式
func (m MultilineText) StaticOn(value string) MultilineText {
	return m.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (m MultilineText) StaticPlaceholder(value string) MultilineText {
	return m.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示 schema
func (m MultilineText) StaticSchema(value string) MultilineText {
	return m.set("staticSchema", value)
}

// Style 设置组件样式
func (m MultilineText) Style(value string) MultilineText {
	return m.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (m MultilineText) TestIdBuilder(value string) MultilineText {
	return m.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (m MultilineText) Testid(value string) MultilineText {
	return m.set("testid", value)
}

// Text 设置文本内容
func (m MultilineText) Text(value string) MultilineText {
	return m.set("text", value)
}

// UseMobileUI 设置是否关闭移动端样式
func (m MultilineText) UseMobileUI(value bool) MultilineText {
	return m.set("useMobileUI", value)
}

// Visible 设置是否显示
func (m MultilineText) Visible(value bool) MultilineText {
	return m.set("visible", value)
}

// VisibleOn 设置显示表达式
func (m MultilineText) VisibleOn(value string) MultilineText {
	return m.set("visibleOn", value)
}
