package comp

type Collapse Schema

func NewCollapse() Collapse {
	return make(Collapse).set("type", "collapse")
}

func (c Collapse) set(key string, value interface{}) Collapse {
	c[key] = value
	return c
}

// Body 设置 body 属性
func (c Collapse) Body(value ...interface{}) Collapse {
	return c.set("body", value)
}

// BodyClassName 设置 bodyClassName 属性
func (c Collapse) BodyClassName(value string) Collapse {
	return c.set("bodyClassName", value)
}

// ClassName 设置 className 属性
func (c Collapse) ClassName(value string) Collapse {
	return c.set("className", value)
}

// Collapsable 设置 collapsable 属性
func (c Collapse) Collapsable(value bool) Collapse {
	return c.set("collapsable", value)
}

// CollapseHeader 设置 collapseHeader 属性
func (c Collapse) CollapseHeader(value string) Collapse {
	return c.set("collapseHeader", value)
}

// Collapsed 设置 collapsed 属性
func (c Collapse) Collapsed(value bool) Collapse {
	return c.set("collapsed", value)
}

// Disabled 设置 disabled 属性
func (c Collapse) Disabled(value bool) Collapse {
	return c.set("disabled", value)
}

// DisabledOn 设置 disabledOn 属性
func (c Collapse) DisabledOn(value string) Collapse {
	return c.set("disabledOn", value)
}

// DivideLine 设置 divideLine 属性
func (c Collapse) DivideLine(value bool) Collapse {
	return c.set("divideLine", value)
}

// EditorSetting 设置 editorSetting 属性
func (c Collapse) EditorSetting(value string) Collapse {
	return c.set("editorSetting", value)
}

// ExpandIcon 设置 expandIcon 属性
func (c Collapse) ExpandIcon(value string) Collapse {
	return c.set("expandIcon", value)
}

// Header 设置 header 属性
func (c Collapse) Header(value string) Collapse {
	return c.set("header", value)
}

// HeaderPosition 设置 headerPosition 属性
func (c Collapse) HeaderPosition(value string) Collapse {
	return c.set("headerPosition", value)
}

// HeadingClassName 设置 headingClassName 属性
func (c Collapse) HeadingClassName(value string) Collapse {
	return c.set("headingClassName", value)
}

// Hidden 设置 hidden 属性
func (c Collapse) Hidden(value bool) Collapse {
	return c.set("hidden", value)
}

// HiddenOn 设置 hiddenOn 属性
func (c Collapse) HiddenOn(value string) Collapse {
	return c.set("hiddenOn", value)
}

// ID 设置 id 属性
func (c Collapse) ID(value string) Collapse {
	return c.set("id", value)
}

// Key 设置 key 属性
func (c Collapse) Key(value string) Collapse {
	return c.set("key", value)
}

// MountOnEnter 设置 mountOnEnter 属性
func (c Collapse) MountOnEnter(value bool) Collapse {
	return c.set("mountOnEnter", value)
}

// OnEvent 设置 onEvent 属性
func (c Collapse) OnEvent(value string) Collapse {
	return c.set("onEvent", value)
}

// ShowArrow 设置 showArrow 属性
func (c Collapse) ShowArrow(value bool) Collapse {
	return c.set("showArrow", value)
}

// Size 设置 size 属性
func (c Collapse) Size(value string) Collapse {
	return c.set("size", value)
}

// Static 设置 static 属性
func (c Collapse) Static(value bool) Collapse {
	return c.set("static", value)
}

// StaticClassName 设置 staticClassName 属性
func (c Collapse) StaticClassName(value string) Collapse {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置 staticInputClassName 属性
func (c Collapse) StaticInputClassName(value string) Collapse {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置 staticLabelClassName 属性
func (c Collapse) StaticLabelClassName(value string) Collapse {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置 staticOn 属性
func (c Collapse) StaticOn(value string) Collapse {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置 staticPlaceholder 属性
func (c Collapse) StaticPlaceholder(value string) Collapse {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置 staticSchema 属性
func (c Collapse) StaticSchema(value string) Collapse {
	return c.set("staticSchema", value)
}

// Style 设置 style 属性
func (c Collapse) Style(value string) Collapse {
	return c.set("style", value)
}

// TestIdBuilder 设置 testIdBuilder 属性
func (c Collapse) TestIdBuilder(value string) Collapse {
	return c.set("testIdBuilder", value)
}

// Testid 设置 testid 属性
func (c Collapse) Testid(value string) Collapse {
	return c.set("testid", value)
}

// UnmountOnExit 设置 unmountOnExit 属性
func (c Collapse) UnmountOnExit(value bool) Collapse {
	return c.set("unmountOnExit", value)
}

// UseMobileUI 设置 useMobileUI 属性
func (c Collapse) UseMobileUI(value bool) Collapse {
	return c.set("useMobileUI", value)
}

// Visible 设置 visible 属性
func (c Collapse) Visible(value bool) Collapse {
	return c.set("visible", value)
}

// VisibleOn 设置 visibleOn 属性
func (c Collapse) VisibleOn(value string) Collapse {
	return c.set("visibleOn", value)
}
