package comp

type propertyItem schema

func PropertyItem() propertyItem {
	return propertyItem{}
}

func (i propertyItem) set(key string, value any) propertyItem {
	i[key] = value
	return i
}

// Label 属性名
func (i propertyItem) Label(value string) propertyItem {
	return i.set("label", value)
}

// Content 属性值
func (i propertyItem) Content(value any) propertyItem {
	return i.set("content", value)
}

// Span 属性值跨几列
func (i propertyItem) Span(value int) propertyItem {
	return i.set("span", value)
}

// VisibleOn 显示表达式
func (i propertyItem) VisibleOn(value string) propertyItem {
	return i.set("visibleOn", value)
}

// HiddenOn 隐藏表达式
func (i propertyItem) HiddenOn(value string) propertyItem {
	return i.set("hiddenOn", value)
}
