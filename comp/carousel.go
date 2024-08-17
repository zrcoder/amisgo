package comp

// Carousel 继承自 BaseRenderer
type Carousel Schema

// NewCarousel 创建一个新的 Carousel 实例
func NewCarousel() Carousel {
	return make(Carousel).set("type", "carousel")
}

func (c Carousel) set(key string, value interface{}) Carousel {
	c[key] = value
	return c
}

// AlwaysShowArrow 设置 alwaysShowArrow 属性
func (c Carousel) AlwaysShowArrow(value bool) Carousel {
	return c.set("alwaysShowArrow", value)
}

// Animation 设置 animation 属性
func (c Carousel) Animation(value string) Carousel {
	return c.set("animation", value)
}

// Auto 设置 auto 属性
func (c Carousel) Auto(value bool) Carousel {
	return c.set("auto", value)
}

// ClassName 设置 className 属性
func (c Carousel) ClassName(value string) Carousel {
	return c.set("className", value)
}

// Controls 设置 controls 属性
func (c Carousel) Controls(value string) Carousel {
	return c.set("controls", value)
}

// ControlsTheme 设置 controlsTheme 属性
func (c Carousel) ControlsTheme(value string) Carousel {
	return c.set("controlsTheme", value)
}

// Disabled 设置 disabled 属性
func (c Carousel) Disabled(value bool) Carousel {
	return c.set("disabled", value)
}

// DisabledOn 设置 disabledOn 属性
func (c Carousel) DisabledOn(value string) Carousel {
	return c.set("disabledOn", value)
}

// Duration 设置 duration 属性
func (c Carousel) Duration(value string) Carousel {
	return c.set("duration", value)
}

// EditorSetting 设置 editorSetting 属性
func (c Carousel) EditorSetting(value string) Carousel {
	return c.set("editorSetting", value)
}

// Height 设置 height 属性
func (c Carousel) Height(value string) Carousel {
	return c.set("height", value)
}

// Hidden 设置 hidden 属性
func (c Carousel) Hidden(value bool) Carousel {
	return c.set("hidden", value)
}

// HiddenOn 设置 hiddenOn 属性
func (c Carousel) HiddenOn(value string) Carousel {
	return c.set("hiddenOn", value)
}

// Icons 设置 icons 属性
func (c Carousel) Icons(value string) Carousel {
	return c.set("icons", value)
}

// ID 设置 id 属性
func (c Carousel) ID(value string) Carousel {
	return c.set("id", value)
}

// Interval 设置 interval 属性
func (c Carousel) Interval(value string) Carousel {
	return c.set("interval", value)
}

// ItemSchema 设置 itemSchema 属性
func (c Carousel) ItemSchema(value string) Carousel {
	return c.set("itemSchema", value)
}

// Multiple 设置 multiple 属性
func (c Carousel) Multiple(value string) Carousel {
	return c.set("multiple", value)
}

// Name 设置 name 属性
func (c Carousel) Name(value string) Carousel {
	return c.set("name", value)
}

// OnEvent 设置 onEvent 属性
func (c Carousel) OnEvent(value string) Carousel {
	return c.set("onEvent", value)
}

// Options 设置 options 属性
func (c Carousel) Options(value string) Carousel {
	return c.set("options", value)
}

// Placeholder 设置 placeholder 属性
func (c Carousel) Placeholder(value string) Carousel {
	return c.set("placeholder", value)
}

// Static 设置 static 属性
func (c Carousel) Static(value bool) Carousel {
	return c.set("static", value)
}

// StaticClassName 设置 staticClassName 属性
func (c Carousel) StaticClassName(value string) Carousel {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置 staticInputClassName 属性
func (c Carousel) StaticInputClassName(value string) Carousel {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置 staticLabelClassName 属性
func (c Carousel) StaticLabelClassName(value string) Carousel {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置 staticOn 属性
func (c Carousel) StaticOn(value string) Carousel {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置 staticPlaceholder 属性
func (c Carousel) StaticPlaceholder(value string) Carousel {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置 staticSchema 属性
func (c Carousel) StaticSchema(value string) Carousel {
	return c.set("staticSchema", value)
}

// Style 设置 style 属性
func (c Carousel) Style(value string) Carousel {
	return c.set("style", value)
}

// TestIdBuilder 设置 testIdBuilder 属性
func (c Carousel) TestIdBuilder(value string) Carousel {
	return c.set("testIdBuilder", value)
}

// TestID 设置 testid 属性
func (c Carousel) TestID(value string) Carousel {
	return c.set("testid", value)
}

// ThumbMode 设置 thumbMode 属性
func (c Carousel) ThumbMode(value string) Carousel {
	return c.set("thumbMode", value)
}

// UseMobileUI 设置 useMobileUI 属性
func (c Carousel) UseMobileUI(value bool) Carousel {
	return c.set("useMobileUI", value)
}

// Visible 设置 visible 属性
func (c Carousel) Visible(value bool) Carousel {
	return c.set("visible", value)
}

// VisibleOn 设置 visibleOn 属性
func (c Carousel) VisibleOn(value string) Carousel {
	return c.set("visibleOn", value)
}

// Width 设置 width 属性
func (c Carousel) Width(value string) Carousel {
	return c.set("width", value)
}
