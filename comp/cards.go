package comp

// Cards 继承自 BaseRenderer
type Cards Schema

// NewCards 创建一个新的 Cards 实例
func NewCards() Cards {
	return make(Cards).set("type", "cards")
}

func (c Cards) set(key string, value interface{}) Cards {
	c[key] = value
	return c
}

// AffixFooter 设置 affixFooter 属性
func (c Cards) AffixFooter(value bool) Cards {
	return c.set("affixFooter", value)
}

// AffixHeader 设置 affixHeader 属性
func (c Cards) AffixHeader(value bool) Cards {
	return c.set("affixHeader", value)
}

// Card 设置 card 属性
func (c Cards) Card(value string) Cards {
	return c.set("card", value)
}

// CheckOnItemClick 设置 checkOnItemClick 属性
func (c Cards) CheckOnItemClick(value bool) Cards {
	return c.set("checkOnItemClick", value)
}

// ClassName 设置 className 属性
func (c Cards) ClassName(value string) Cards {
	return c.set("className", value)
}

// Disabled 设置 disabled 属性
func (c Cards) Disabled(value bool) Cards {
	return c.set("disabled", value)
}

// DisabledOn 设置 disabledOn 属性
func (c Cards) DisabledOn(value string) Cards {
	return c.set("disabledOn", value)
}

// EditorSetting 设置 editorSetting 属性
func (c Cards) EditorSetting(value string) Cards {
	return c.set("editorSetting", value)
}

// Footer 设置 footer 属性
func (c Cards) Footer(value string) Cards {
	return c.set("footer", value)
}

// FooterClassName 设置 footerClassName 属性
func (c Cards) FooterClassName(value string) Cards {
	return c.set("footerClassName", value)
}

// Header 设置 header 属性
func (c Cards) Header(value string) Cards {
	return c.set("header", value)
}

// HeaderClassName 设置 headerClassName 属性
func (c Cards) HeaderClassName(value string) Cards {
	return c.set("headerClassName", value)
}

// Hidden 设置 hidden 属性
func (c Cards) Hidden(value bool) Cards {
	return c.set("hidden", value)
}

// HiddenOn 设置 hiddenOn 属性
func (c Cards) HiddenOn(value string) Cards {
	return c.set("hiddenOn", value)
}

// HideCheckToggler 设置 hideCheckToggler 属性
func (c Cards) HideCheckToggler(value bool) Cards {
	return c.set("hideCheckToggler", value)
}

// ID 设置 id 属性
func (c Cards) ID(value string) Cards {
	return c.set("id", value)
}

// ItemCheckableOn 设置 itemCheckableOn 属性
func (c Cards) ItemCheckableOn(value string) Cards {
	return c.set("itemCheckableOn", value)
}

// ItemClassName 设置 itemClassName 属性
func (c Cards) ItemClassName(value string) Cards {
	return c.set("itemClassName", value)
}

// ItemDraggableOn 设置 itemDraggableOn 属性
func (c Cards) ItemDraggableOn(value string) Cards {
	return c.set("itemDraggableOn", value)
}

// LoadingConfig 设置 loadingConfig 属性
func (c Cards) LoadingConfig(value string) Cards {
	return c.set("loadingConfig", value)
}

// MasonryLayout 设置 masonryLayout 属性
func (c Cards) MasonryLayout(value bool) Cards {
	return c.set("masonryLayout", value)
}

// OnEvent 设置 onEvent 属性
func (c Cards) OnEvent(value string) Cards {
	return c.set("onEvent", value)
}

// Placeholder 设置 placeholder 属性
func (c Cards) Placeholder(value string) Cards {
	return c.set("placeholder", value)
}

// ShowFooter 设置 showFooter 属性
func (c Cards) ShowFooter(value bool) Cards {
	return c.set("showFooter", value)
}

// ShowHeader 设置 showHeader 属性
func (c Cards) ShowHeader(value bool) Cards {
	return c.set("showHeader", value)
}

// Source 设置 source 属性
func (c Cards) Source(value string) Cards {
	return c.set("source", value)
}

// Static 设置 static 属性
func (c Cards) Static(value bool) Cards {
	return c.set("static", value)
}

// StaticClassName 设置 staticClassName 属性
func (c Cards) StaticClassName(value string) Cards {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置 staticInputClassName 属性
func (c Cards) StaticInputClassName(value string) Cards {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置 staticLabelClassName 属性
func (c Cards) StaticLabelClassName(value string) Cards {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置 staticOn 属性
func (c Cards) StaticOn(value string) Cards {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置 staticPlaceholder 属性
func (c Cards) StaticPlaceholder(value string) Cards {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置 staticSchema 属性
func (c Cards) StaticSchema(value string) Cards {
	return c.set("staticSchema", value)
}

// Style 设置 style 属性
func (c Cards) Style(value string) Cards {
	return c.set("style", value)
}

// TestIdBuilder 设置 testIdBuilder 属性
func (c Cards) TestIdBuilder(value string) Cards {
	return c.set("testIdBuilder", value)
}

// TestID 设置 testid 属性
func (c Cards) TestID(value string) Cards {
	return c.set("testid", value)
}

// Title 设置 title 属性
func (c Cards) Title(value string) Cards {
	return c.set("title", value)
}

// UseMobileUI 设置 useMobileUI 属性
func (c Cards) UseMobileUI(value bool) Cards {
	return c.set("useMobileUI", value)
}

// ValueField 设置 valueField 属性
func (c Cards) ValueField(value string) Cards {
	return c.set("valueField", value)
}

// Visible 设置 visible 属性
func (c Cards) Visible(value bool) Cards {
	return c.set("visible", value)
}

// VisibleOn 设置 visibleOn 属性
func (c Cards) VisibleOn(value string) Cards {
	return c.set("visibleOn", value)
}
