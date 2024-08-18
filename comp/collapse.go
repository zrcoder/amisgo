package comp

type collapse schema

func Collapse() collapse {
	return make(collapse).set("type", "collapse")
}

func (c collapse) set(key string, value interface{}) collapse {
	c[key] = value
	return c
}

// Body 设置 body 属性
func (c collapse) Body(value ...interface{}) collapse {
	return c.set("body", value)
}

// BodyClassName 设置 bodyClassName 属性
func (c collapse) BodyClassName(value string) collapse {
	return c.set("bodyClassName", value)
}

// ClassName 设置 className 属性
func (c collapse) ClassName(value string) collapse {
	return c.set("className", value)
}

// Collapsable 设置 collapsable 属性
func (c collapse) Collapsable(value bool) collapse {
	return c.set("collapsable", value)
}

// CollapseHeader 设置 collapseHeader 属性
func (c collapse) CollapseHeader(value string) collapse {
	return c.set("collapseHeader", value)
}

// Collapsed 设置 collapsed 属性
func (c collapse) Collapsed(value bool) collapse {
	return c.set("collapsed", value)
}

// Disabled 设置 disabled 属性
func (c collapse) Disabled(value bool) collapse {
	return c.set("disabled", value)
}

// DisabledOn 设置 disabledOn 属性
func (c collapse) DisabledOn(value string) collapse {
	return c.set("disabledOn", value)
}

// DivideLine 设置 divideLine 属性
func (c collapse) DivideLine(value bool) collapse {
	return c.set("divideLine", value)
}

// EditorSetting 设置 editorSetting 属性
func (c collapse) EditorSetting(value string) collapse {
	return c.set("editorSetting", value)
}

// ExpandIcon 设置 expandIcon 属性
func (c collapse) ExpandIcon(value string) collapse {
	return c.set("expandIcon", value)
}

// Header 设置 header 属性
func (c collapse) Header(value string) collapse {
	return c.set("header", value)
}

// HeaderPosition 设置 headerPosition 属性
func (c collapse) HeaderPosition(value string) collapse {
	return c.set("headerPosition", value)
}

// HeadingClassName 设置 headingClassName 属性
func (c collapse) HeadingClassName(value string) collapse {
	return c.set("headingClassName", value)
}

// Hidden 设置 hidden 属性
func (c collapse) Hidden(value bool) collapse {
	return c.set("hidden", value)
}

// HiddenOn 设置 hiddenOn 属性
func (c collapse) HiddenOn(value string) collapse {
	return c.set("hiddenOn", value)
}

// ID 设置 id 属性
func (c collapse) ID(value string) collapse {
	return c.set("id", value)
}

// Key 设置 key 属性
func (c collapse) Key(value string) collapse {
	return c.set("key", value)
}

// MountOnEnter 设置 mountOnEnter 属性
func (c collapse) MountOnEnter(value bool) collapse {
	return c.set("mountOnEnter", value)
}

// OnEvent 设置 onEvent 属性
func (c collapse) OnEvent(value string) collapse {
	return c.set("onEvent", value)
}

// ShowArrow 设置 showArrow 属性
func (c collapse) ShowArrow(value bool) collapse {
	return c.set("showArrow", value)
}

// Size 设置 size 属性
func (c collapse) Size(value string) collapse {
	return c.set("size", value)
}

// Static 设置 static 属性
func (c collapse) Static(value bool) collapse {
	return c.set("static", value)
}

// StaticClassName 设置 staticClassName 属性
func (c collapse) StaticClassName(value string) collapse {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置 staticInputClassName 属性
func (c collapse) StaticInputClassName(value string) collapse {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置 staticLabelClassName 属性
func (c collapse) StaticLabelClassName(value string) collapse {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置 staticOn 属性
func (c collapse) StaticOn(value string) collapse {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置 staticPlaceholder 属性
func (c collapse) StaticPlaceholder(value string) collapse {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置 staticSchema 属性
func (c collapse) StaticSchema(value string) collapse {
	return c.set("staticSchema", value)
}

// Style 设置 style 属性
func (c collapse) Style(value string) collapse {
	return c.set("style", value)
}

// TestIdBuilder 设置 testIdBuilder 属性
func (c collapse) TestIdBuilder(value string) collapse {
	return c.set("testIdBuilder", value)
}

// Testid 设置 testid 属性
func (c collapse) Testid(value string) collapse {
	return c.set("testid", value)
}

// UnmountOnExit 设置 unmountOnExit 属性
func (c collapse) UnmountOnExit(value bool) collapse {
	return c.set("unmountOnExit", value)
}

// UseMobileUI 设置 useMobileUI 属性
func (c collapse) UseMobileUI(value bool) collapse {
	return c.set("useMobileUI", value)
}

// Visible 设置 visible 属性
func (c collapse) Visible(value bool) collapse {
	return c.set("visible", value)
}

// VisibleOn 设置 visibleOn 属性
func (c collapse) VisibleOn(value string) collapse {
	return c.set("visibleOn", value)
}
