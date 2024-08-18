package comp

// color 显示渲染器
type color schema

// Color 创建一个新的 Color 实例
func Color() color {
	return make(color).set("type", "color")
}

func (c color) set(key string, value interface{}) color {
	c[key] = value
	return c
}

// ClassName 设置容器 css 类名
func (c color) ClassName(value string) color {
	return c.set("className", value)
}

// DefaultColor 设置默认颜色
func (c color) DefaultColor(value string) color {
	return c.set("defaultColor", value)
}

// Disabled 设置是否禁用
func (c color) Disabled(value bool) color {
	return c.set("disabled", value)
}

// DisabledOn 设置是否禁用表达式
func (c color) DisabledOn(value string) color {
	return c.set("disabledOn", value)
}

// EditorSetting 设置编辑器配置
func (c color) EditorSetting(value string) color {
	return c.set("editorSetting", value)
}

// Hidden 设置是否隐藏
func (c color) Hidden(value bool) color {
	return c.set("hidden", value)
}

// HiddenOn 设置是否隐藏表达式
func (c color) HiddenOn(value string) color {
	return c.set("hiddenOn", value)
}

// ID 设置组件唯一 id
func (c color) ID(value string) color {
	return c.set("id", value)
}

// OnEvent 设置事件动作配置
func (c color) OnEvent(value string) color {
	return c.set("onEvent", value)
}

// ShowValue 设置是否用文字显示值
func (c color) ShowValue(value bool) color {
	return c.set("showValue", value)
}

// Static 设置是否静态展示
func (c color) Static(value bool) color {
	return c.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (c color) StaticClassName(value string) color {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项Value类名
func (c color) StaticInputClassName(value string) color {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项Label类名
func (c color) StaticLabelClassName(value string) color {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置是否静态展示表达式
func (c color) StaticOn(value string) color {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (c color) StaticPlaceholder(value string) color {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示模式
func (c color) StaticSchema(value string) color {
	return c.set("staticSchema", value)
}

// Style 设置组件样式
func (c color) Style(value string) color {
	return c.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (c color) TestIdBuilder(value string) color {
	return c.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (c color) Testid(value string) color {
	return c.set("testid", value)
}

// UseMobileUI 设置是否使用移动端样式
func (c color) UseMobileUI(value bool) color {
	return c.set("useMobileUI", value)
}

// Visible 设置是否显示
func (c color) Visible(value bool) color {
	return c.set("visible", value)
}

// VisibleOn 设置显示条件表达式
func (c color) VisibleOn(value string) color {
	return c.set("visibleOn", value)
}
