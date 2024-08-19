package comp

// cards
type cards schema

// Cards 创建一个新的 Cards 实例
func Cards() cards {
	return make(cards).set("type", "cards")
}

func (c cards) set(key string, value any) cards {
	c[key] = value
	return c
}

// AffixFooter 设置 affixFooter 属性
func (c cards) AffixFooter(value bool) cards {
	return c.set("affixFooter", value)
}

// AffixHeader 设置 affixHeader 属性
func (c cards) AffixHeader(value bool) cards {
	return c.set("affixHeader", value)
}

// Card 设置 card 属性
func (c cards) Card(value string) cards {
	return c.set("card", value)
}

// CheckOnItemClick 设置 checkOnItemClick 属性
func (c cards) CheckOnItemClick(value bool) cards {
	return c.set("checkOnItemClick", value)
}

// ClassName 设置 className 属性
func (c cards) ClassName(value string) cards {
	return c.set("className", value)
}

// Disabled 设置 disabled 属性
func (c cards) Disabled(value bool) cards {
	return c.set("disabled", value)
}

// DisabledOn 设置 disabledOn 属性
func (c cards) DisabledOn(value string) cards {
	return c.set("disabledOn", value)
}

// EditorSetting 设置 editorSetting 属性
func (c cards) EditorSetting(value string) cards {
	return c.set("editorSetting", value)
}

// Footer 设置 footer 属性
func (c cards) Footer(value string) cards {
	return c.set("footer", value)
}

// FooterClassName 设置 footerClassName 属性
func (c cards) FooterClassName(value string) cards {
	return c.set("footerClassName", value)
}

// Header 设置 header 属性
func (c cards) Header(value string) cards {
	return c.set("header", value)
}

// HeaderClassName 设置 headerClassName 属性
func (c cards) HeaderClassName(value string) cards {
	return c.set("headerClassName", value)
}

// Hidden 设置 hidden 属性
func (c cards) Hidden(value bool) cards {
	return c.set("hidden", value)
}

// HiddenOn 设置 hiddenOn 属性
func (c cards) HiddenOn(value string) cards {
	return c.set("hiddenOn", value)
}

// HideCheckToggler 设置 hideCheckToggler 属性
func (c cards) HideCheckToggler(value bool) cards {
	return c.set("hideCheckToggler", value)
}

// ID 设置 id 属性
func (c cards) ID(value string) cards {
	return c.set("id", value)
}

// ItemCheckableOn 设置 itemCheckableOn 属性
func (c cards) ItemCheckableOn(value string) cards {
	return c.set("itemCheckableOn", value)
}

// ItemClassName 设置 itemClassName 属性
func (c cards) ItemClassName(value string) cards {
	return c.set("itemClassName", value)
}

// ItemDraggableOn 设置 itemDraggableOn 属性
func (c cards) ItemDraggableOn(value string) cards {
	return c.set("itemDraggableOn", value)
}

// LoadingConfig 设置 loadingConfig 属性
func (c cards) LoadingConfig(value string) cards {
	return c.set("loadingConfig", value)
}

// MasonryLayout 设置 masonryLayout 属性
func (c cards) MasonryLayout(value bool) cards {
	return c.set("masonryLayout", value)
}

// OnEvent 设置 onEvent 属性
func (c cards) OnEvent(value any) cards {
	return c.set("onEvent", value)
}

// Placeholder 设置 placeholder 属性
func (c cards) Placeholder(value string) cards {
	return c.set("placeholder", value)
}

// ShowFooter 设置 showFooter 属性
func (c cards) ShowFooter(value bool) cards {
	return c.set("showFooter", value)
}

// ShowHeader 设置 showHeader 属性
func (c cards) ShowHeader(value bool) cards {
	return c.set("showHeader", value)
}

// Source 设置 source 属性
func (c cards) Source(value string) cards {
	return c.set("source", value)
}

// Static 设置 static 属性
func (c cards) Static(value bool) cards {
	return c.set("static", value)
}

// StaticClassName 设置 staticClassName 属性
func (c cards) StaticClassName(value string) cards {
	return c.set("staticClassName", value)
}

// StaticInputClassName 设置 staticInputClassName 属性
func (c cards) StaticInputClassName(value string) cards {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 设置 staticLabelClassName 属性
func (c cards) StaticLabelClassName(value string) cards {
	return c.set("staticLabelClassName", value)
}

// StaticOn 设置 staticOn 属性
func (c cards) StaticOn(value string) cards {
	return c.set("staticOn", value)
}

// StaticPlaceholder 设置 staticPlaceholder 属性
func (c cards) StaticPlaceholder(value string) cards {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 设置 staticSchema 属性
func (c cards) StaticSchema(value string) cards {
	return c.set("staticSchema", value)
}

// Style 设置 style 属性
func (c cards) Style(value any) cards {
	return c.set("style", value)
}

// TestIdBuilder 设置 testIdBuilder 属性
func (c cards) TestIdBuilder(value string) cards {
	return c.set("testIdBuilder", value)
}

// TestID 设置 testid 属性
func (c cards) TestID(value string) cards {
	return c.set("testid", value)
}

// Title 设置 title 属性
func (c cards) Title(value any) cards {
	return c.set("title", value)
}

// UseMobileUI 设置 useMobileUI 属性
func (c cards) UseMobileUI(value bool) cards {
	return c.set("useMobileUI", value)
}

// ValueField 设置 valueField 属性
func (c cards) ValueField(value string) cards {
	return c.set("valueField", value)
}

// Visible 设置 visible 属性
func (c cards) Visible(value bool) cards {
	return c.set("visible", value)
}

// VisibleOn 设置 visibleOn 属性
func (c cards) VisibleOn(value string) cards {
	return c.set("visibleOn", value)
}
