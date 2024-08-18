package comp

// card
type card schema

// Card 创建一个新的 Card 实例
func Card() card {
	return make(card).set("type", "card")
}

func (c card) set(key string, value interface{}) card {
	c[key] = value
	return c
}

// Actions 设置 actions 属性
func (c card) Actions(value string) card {
	return c.set("actions", value)
}

// Body 设置 body 属性
func (c card) Body(value ...interface{}) card {
	return c.set("body", value)
}

// BodyClassName 设置 bodyClassName 属性
func (c card) BodyClassName(value string) card {
	return c.set("bodyClassName", value)
}

// CheckOnItemClick 设置 checkOnItemClick 属性
func (c card) CheckOnItemClick(value bool) card {
	return c.set("checkOnItemClick", value)
}

// ClassName 设置 className 属性
func (c card) ClassName(value string) card {
	return c.set("className", value)
}

// Disabled 设置 disabled 属性
func (c card) Disabled(value bool) card {
	return c.set("disabled", value)
}

// DisabledOn 设置 disabledOn 属性
func (c card) DisabledOn(value string) card {
	return c.set("disabledOn", value)
}

// EditorSetting 设置 editorSetting 属性
func (c card) EditorSetting(value string) card {
	return c.set("editorSetting", value)
}

// Header 设置 header 属性
func (c card) Header(value string) card {
	return c.set("header", value)
}

// Hidden 设置 hidden 属性
func (c card) Hidden(value bool) card {
	return c.set("hidden", value)
}

// HiddenOn 设置 hiddenOn 属性
func (c card) HiddenOn(value string) card {
	return c.set("hiddenOn", value)
}

// HideCheckToggler 设置 hideCheckToggler 属性
func (c card) HideCheckToggler(value bool) card {
	return c.set("hideCheckToggler", value)
}

// ID 设置 id 属性
func (c card) ID(value string) card {
	return c.set("id", value)
}

// ItemAction 设置 itemAction 属性
func (c card) ItemAction(value string) card {
	return c.set("itemAction", value)
}

// Media 设置 media 属性
func (c card) Media(value string) card {
	return c.set("media", value)
}

// OnEvent 设置 onEvent 属性
func (c card) OnEvent(value string) card {
	return c.set("onEvent", value)
}

// Secondary 设置 secondary 属性
func (c card) Secondary(value string) card {
	return c.set("secondary", value)
}

// Static 设置 static 属性
func (c card) Static(value bool) card {
	return c.set("static", value)
}

// StaticClassName 设置 staticClassName 属性
func (c card) StaticClassName(value string) card {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置 staticInputClassName 属性
func (c card) StaticInputClassName(value string) card {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置 staticLabelClassName 属性
func (c card) StaticLabelClassName(value string) card {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置 staticOn 属性
func (c card) StaticOn(value string) card {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置 staticPlaceholder 属性
func (c card) StaticPlaceholder(value string) card {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置 staticSchema 属性
func (c card) StaticSchema(value string) card {
	return c.set("staticSchema", value)
}

// Style 设置 style 属性
func (c card) Style(value string) card {
	return c.set("style", value)
}

// TestIdBuilder 设置 testIdBuilder 属性
func (c card) TestIdBuilder(value string) card {
	return c.set("testIdBuilder", value)
}

// TestID 设置 testid 属性
func (c card) TestID(value string) card {
	return c.set("testid", value)
}

// Toolbar 设置 toolbar 属性
func (c card) Toolbar(value string) card {
	return c.set("toolbar", value)
}

// UseCardLabel 设置 useCardLabel 属性
func (c card) UseCardLabel(value bool) card {
	return c.set("useCardLabel", value)
}

// UseMobileUI 设置 useMobileUI 属性
func (c card) UseMobileUI(value bool) card {
	return c.set("useMobileUI", value)
}

// Visible 设置 visible 属性
func (c card) Visible(value bool) card {
	return c.set("visible", value)
}

// VisibleOn 设置 visibleOn 属性
func (c card) VisibleOn(value string) card {
	return c.set("visibleOn", value)
}

// WrapperComponent 设置 wrapperComponent 属性
func (c card) WrapperComponent(value string) card {
	return c.set("wrapperComponent", value)
}
