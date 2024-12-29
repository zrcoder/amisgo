package comp

// carousel
type carousel Schema

// Carousel 创建一个新的 Carousel 实例
func Carousel() carousel {
	return make(carousel).set("type", "carousel")
}

func (c carousel) set(key string, value any) carousel {
	c[key] = value
	return c
}

// AlwaysShowArrow 设置 alwaysShowArrow 属性
func (c carousel) AlwaysShowArrow(value bool) carousel {
	return c.set("alwaysShowArrow", value)
}

// Animation 设置 animation 属性
func (c carousel) Animation(value string) carousel {
	return c.set("animation", value)
}

// Auto 设置 auto 属性
func (c carousel) Auto(value bool) carousel {
	return c.set("auto", value)
}

// ClassName 设置 className 属性
func (c carousel) ClassName(value string) carousel {
	return c.set("className", value)
}

// Controls 设置 controls 属性
func (c carousel) Controls(value string) carousel {
	return c.set("controls", value)
}

// ControlsTheme 设置 controlsTheme 属性
func (c carousel) ControlsTheme(value string) carousel {
	return c.set("controlsTheme", value)
}

// Disabled 设置 disabled 属性
func (c carousel) Disabled(value bool) carousel {
	return c.set("disabled", value)
}

// DisabledOn 设置 disabledOn 属性
func (c carousel) DisabledOn(value string) carousel {
	return c.set("disabledOn", value)
}

// Duration 设置 duration 属性
func (c carousel) Duration(value string) carousel {
	return c.set("duration", value)
}

// EditorSetting 设置 editorSetting 属性
func (c carousel) EditorSetting(value string) carousel {
	return c.set("editorSetting", value)
}

// Height 设置 height 属性
func (c carousel) Height(value string) carousel {
	return c.set("height", value)
}

// Hidden 设置 hidden 属性
func (c carousel) Hidden(value bool) carousel {
	return c.set("hidden", value)
}

// HiddenOn 设置 hiddenOn 属性
func (c carousel) HiddenOn(value string) carousel {
	return c.set("hiddenOn", value)
}

// Icons 设置 icons 属性
func (c carousel) Icons(value string) carousel {
	return c.set("icons", value)
}

// ID 设置 id 属性
func (c carousel) ID(value string) carousel {
	return c.set("id", value)
}

// Interval 设置 interval 属性
func (c carousel) Interval(value string) carousel {
	return c.set("interval", value)
}

// ItemSchema 设置 itemSchema 属性
func (c carousel) ItemSchema(value string) carousel {
	return c.set("itemSchema", value)
}

// Multiple 设置 multiple 属性
func (c carousel) Multiple(value any) carousel {
	return c.set("multiple", value)
}

// Name 设置 name 属性
func (c carousel) Name(value string) carousel {
	return c.set("name", value)
}

// OnEvent 设置 onEvent 属性
func (c carousel) OnEvent(value any) carousel {
	return c.set("onEvent", value)
}

// Options 设置 options 属性
func (c carousel) Options(value ...any) carousel {
	return c.set("options", value)
}

// Placeholder 设置 placeholder 属性
func (c carousel) Placeholder(value string) carousel {
	return c.set("placeholder", value)
}

// Static 设置 static 属性
func (c carousel) Static(value bool) carousel {
	return c.set("static", value)
}

// StaticClassName 设置 staticClassName 属性
func (c carousel) StaticClassName(value string) carousel {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置 staticInputClassName 属性
func (c carousel) StaticInputClassName(value string) carousel {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置 staticLabelClassName 属性
func (c carousel) StaticLabelClassName(value string) carousel {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置 staticOn 属性
func (c carousel) StaticOn(value string) carousel {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置 staticPlaceholder 属性
func (c carousel) StaticPlaceholder(value string) carousel {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置 staticSchema 属性
func (c carousel) StaticSchema(value string) carousel {
	return c.set("staticSchema", value)
}

// Style 设置 style 属性
func (c carousel) Style(value any) carousel {
	return c.set("style", value)
}

// TestIdBuilder 设置 testIdBuilder 属性
func (c carousel) TestIdBuilder(value string) carousel {
	return c.set("testIdBuilder", value)
}

// TestID 设置 testid 属性
func (c carousel) TestID(value string) carousel {
	return c.set("testid", value)
}

// ThumbMode 设置 thumbMode 属性, contain | cover
func (c carousel) ThumbMode(value string) carousel {
	return c.set("thumbMode", value)
}

// UseMobileUI 设置 useMobileUI 属性
func (c carousel) UseMobileUI(value bool) carousel {
	return c.set("useMobileUI", value)
}

// Visible 设置 visible 属性
func (c carousel) Visible(value bool) carousel {
	return c.set("visible", value)
}

// VisibleOn 设置 visibleOn 属性
func (c carousel) VisibleOn(value string) carousel {
	return c.set("visibleOn", value)
}

// Width 设置 width 属性
func (c carousel) Width(value string) carousel {
	return c.set("width", value)
}
