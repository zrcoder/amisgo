package comp

// Card2
type Card2 Schema

// NewCard2 创建一个新的 Card2 实例
func NewCard2() Card2 {
	return make(Card2).set("type", "card2")
}

func (c2 Card2) set(key string, value interface{}) Card2 {
	c2[key] = value
	return c2
}

// Body 设置 body 属性
func (c2 Card2) Body(value ...interface{}) Card2 {
	return c2.set("body", value)
}

// BodyClassName 设置 bodyClassName 属性
func (c2 Card2) BodyClassName(value string) Card2 {
	return c2.set("bodyClassName", value)
}

// CheckOnItemClick 设置 checkOnItemClick 属性
func (c2 Card2) CheckOnItemClick(value bool) Card2 {
	return c2.set("checkOnItemClick", value)
}

// ClassName 设置 className 属性
func (c2 Card2) ClassName(value string) Card2 {
	return c2.set("className", value)
}

// Disabled 设置 disabled 属性
func (c2 Card2) Disabled(value bool) Card2 {
	return c2.set("disabled", value)
}

// DisabledOn 设置 disabledOn 属性
func (c2 Card2) DisabledOn(value string) Card2 {
	return c2.set("disabledOn", value)
}

// EditorSetting 设置 editorSetting 属性
func (c2 Card2) EditorSetting(value string) Card2 {
	return c2.set("editorSetting", value)
}

// Hidden 设置 hidden 属性
func (c2 Card2) Hidden(value bool) Card2 {
	return c2.set("hidden", value)
}

// HiddenOn 设置 hiddenOn 属性
func (c2 Card2) HiddenOn(value string) Card2 {
	return c2.set("hiddenOn", value)
}

// HideCheckToggler 设置 hideCheckToggler 属性
func (c2 Card2) HideCheckToggler(value bool) Card2 {
	return c2.set("hideCheckToggler", value)
}

// ID 设置 id 属性
func (c2 Card2) ID(value string) Card2 {
	return c2.set("id", value)
}

// OnEvent 设置 onEvent 属性
func (c2 Card2) OnEvent(value string) Card2 {
	return c2.set("onEvent", value)
}

// Static 设置 static 属性
func (c2 Card2) Static(value bool) Card2 {
	return c2.set("static", value)
}

// StaticClassName 设置 staticClassName 属性
func (c2 Card2) StaticClassName(value string) Card2 {
	return c2.set("staticClassName", value)
}

// StaticInputClassName 设置 staticInputClassName 属性
func (c2 Card2) StaticInputClassName(value string) Card2 {
	return c2.set("staticInputClassName", value)
}

// StaticLabelClassName 设置 staticLabelClassName 属性
func (c2 Card2) StaticLabelClassName(value string) Card2 {
	return c2.set("staticLabelClassName", value)
}

// StaticOn 设置 staticOn 属性
func (c2 Card2) StaticOn(value string) Card2 {
	return c2.set("staticOn", value)
}

// StaticPlaceholder 设置 staticPlaceholder 属性
func (c2 Card2) StaticPlaceholder(value string) Card2 {
	return c2.set("staticPlaceholder", value)
}

// StaticSchema 设置 staticSchema 属性
func (c2 Card2) StaticSchema(value string) Card2 {
	return c2.set("staticSchema", value)
}

// Style 设置 style 属性
func (c2 Card2) Style(value string) Card2 {
	return c2.set("style", value)
}

// TestIdBuilder 设置 testIdBuilder 属性
func (c2 Card2) TestIdBuilder(value string) Card2 {
	return c2.set("testIdBuilder", value)
}

// TestID 设置 testid 属性
func (c2 Card2) TestID(value string) Card2 {
	return c2.set("testid", value)
}

// UseMobileUI 设置 useMobileUI 属性
func (c2 Card2) UseMobileUI(value bool) Card2 {
	return c2.set("useMobileUI", value)
}

// Visible 设置 visible 属性
func (c2 Card2) Visible(value bool) Card2 {
	return c2.set("visible", value)
}

// VisibleOn 设置 visibleOn 属性
func (c2 Card2) VisibleOn(value string) Card2 {
	return c2.set("visibleOn", value)
}

// WrapperComponent 设置 wrapperComponent 属性
func (c2 Card2) WrapperComponent(value string) Card2 {
	return c2.set("wrapperComponent", value)
}
