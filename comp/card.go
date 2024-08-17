package comp

// Card 继承自 BaseRenderer
type Card BaseRenderer

// NewCard 创建一个新的 Card 实例
func NewCard() Card {
	c := Card(make(BaseRenderer))
	return c.set("type", "card")
}

func (c Card) set(key string, value interface{}) Card {
	c[key] = value
	return c
}

// Actions 设置 actions 属性
func (c Card) Actions(value string) Card {
	return c.set("actions", value)
}

// Body 设置 body 属性
func (c Card) Body(value ...interface{}) Card {
	return c.set("body", value)
}

// ClassName 设置 className 属性
func (c Card) ClassName(value string) Card {
	return c.set("className", value)
}

// Disabled 设置 disabled 属性
func (c Card) Disabled(value bool) Card {
	return c.set("disabled", value)
}

// DisabledOn 设置 disabledOn 属性
func (c Card) DisabledOn(value string) Card {
	return c.set("disabledOn", value)
}

// EditorSetting 设置 editorSetting 属性
func (c Card) EditorSetting(value string) Card {
	return c.set("editorSetting", value)
}

// Header 设置 header 属性
func (c Card) Header(value string) Card {
	return c.set("header", value)
}

// Hidden 设置 hidden 属性
func (c Card) Hidden(value bool) Card {
	return c.set("hidden", value)
}

// HiddenOn 设置 hiddenOn 属性
func (c Card) HiddenOn(value string) Card {
	return c.set("hiddenOn", value)
}

// ID 设置 id 属性
func (c Card) ID(value string) Card {
	return c.set("id", value)
}

// ItemAction 设置 itemAction 属性
func (c Card) ItemAction(value string) Card {
	return c.set("itemAction", value)
}

// Media 设置 media 属性
func (c Card) Media(value string) Card {
	return c.set("media", value)
}

// OnEvent 设置 onEvent 属性
func (c Card) OnEvent(value string) Card {
	return c.set("onEvent", value)
}

// Secondary 设置 secondary 属性
func (c Card) Secondary(value string) Card {
	return c.set("secondary", value)
}

// Static 设置 static 属性
func (c Card) Static(value bool) Card {
	return c.set("static", value)
}

// StaticClassName 设置 staticClassName 属性
func (c Card) StaticClassName(value string) Card {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置 staticInputClassName 属性
func (c Card) StaticInputClassName(value string) Card {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置 staticLabelClassName 属性
func (c Card) StaticLabelClassName(value string) Card {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置 staticOn 属性
func (c Card) StaticOn(value string) Card {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置 staticPlaceholder 属性
func (c Card) StaticPlaceholder(value string) Card {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置 staticSchema 属性
func (c Card) StaticSchema(value string) Card {
	return c.set("staticSchema", value)
}

// Style 设置 style 属性
func (c Card) Style(value string) Card {
	return c.set("style", value)
}

// TestIdBuilder 设置 testIdBuilder 属性
func (c Card) TestIdBuilder(value string) Card {
	return c.set("testIdBuilder", value)
}

// TestID 设置 testid 属性
func (c Card) TestID(value string) Card {
	return c.set("testid", value)
}

// Toolbar 设置 toolbar 属性
func (c Card) Toolbar(value string) Card {
	return c.set("toolbar", value)
}

// UseCardLabel 设置 useCardLabel 属性
func (c Card) UseCardLabel(value bool) Card {
	return c.set("useCardLabel", value)
}

// UseMobileUI 设置 useMobileUI 属性
func (c Card) UseMobileUI(value bool) Card {
	return c.set("useMobileUI", value)
}

// Visible 设置 visible 属性
func (c Card) Visible(value bool) Card {
	return c.set("visible", value)
}

// VisibleOn 设置 visibleOn 属性
func (c Card) VisibleOn(value string) Card {
	return c.set("visibleOn", value)
}
