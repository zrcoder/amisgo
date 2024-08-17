package comp

// Color 显示渲染器
type Color Schema

// NewColor 创建一个新的 Color 实例
func NewColor() Color {
	return make(Color).set("type", "color")
}

func (c Color) set(key string, value interface{}) Color {
	c[key] = value
	return c
}

// ClassName 设置容器 css 类名
func (c Color) ClassName(value string) Color {
	return c.set("className", value)
}

// DefaultColor 设置默认颜色
func (c Color) DefaultColor(value string) Color {
	return c.set("defaultColor", value)
}

// Disabled 设置是否禁用
func (c Color) Disabled(value bool) Color {
	return c.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
func (c Color) DisabledOn(value string) Color {
	return c.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置
func (c Color) EditorSetting(value string) Color {
	return c.set("editorSetting", value)
}

// Hidden 设置是否隐藏
func (c Color) Hidden(value bool) Color {
	return c.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
func (c Color) HiddenOn(value string) Color {
	return c.set("hiddenOn", value)
}

// ID 设置组件唯一 id
func (c Color) ID(value string) Color {
	return c.set("id", value)
}

// OnEvent 设置事件动作配置
func (c Color) OnEvent(value string) Color {
	return c.set("onEvent", value)
}

// ShowValue 设置是否用文字显示值
func (c Color) ShowValue(value bool) Color {
	return c.set("showValue", value)
}

// Static 设置是否静态展示
func (c Color) Static(value bool) Color {
	return c.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (c Color) StaticClassName(value string) Color {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项Value类名
func (c Color) StaticInputClassName(value string) Color {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项Label类名
func (c Color) StaticLabelClassName(value string) Color {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式
func (c Color) StaticOn(value string) Color {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (c Color) StaticPlaceholder(value string) Color {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示模式
func (c Color) StaticSchema(value string) Color {
	return c.set("staticSchema", value)
}

// Style 设置组件样式
func (c Color) Style(value string) Color {
	return c.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (c Color) TestIdBuilder(value string) Color {
	return c.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (c Color) Testid(value string) Color {
	return c.set("testid", value)
}

// UseMobileUI 设置是否使用移动端样式
func (c Color) UseMobileUI(value bool) Color {
	return c.set("useMobileUI", value)
}

// Visible 设置是否显示
func (c Color) Visible(value bool) Color {
	return c.set("visible", value)
}

// VisibleOn 设置显示条件表达式
func (c Color) VisibleOn(value string) Color {
	return c.set("visibleOn", value)
}
