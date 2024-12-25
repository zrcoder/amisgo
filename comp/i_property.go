package comp

type MPropertyItem Schema

func PropertyItem() MPropertyItem {
	return MPropertyItem{}
}

func (i MPropertyItem) set(key string, value any) MPropertyItem {
	i[key] = value
	return i
}

// Label 属性名
func (i MPropertyItem) Label(value string) MPropertyItem {
	return i.set("label", value)
}

// Content 属性值
func (i MPropertyItem) Content(value any) MPropertyItem {
	return i.set("content", value)
}

// Span 属性值跨几列
func (i MPropertyItem) Span(value int) MPropertyItem {
	return i.set("span", value)
}

// VisibleOn 显示表达式
func (i MPropertyItem) VisibleOn(value string) MPropertyItem {
	return i.set("visibleOn", value)
}

// HiddenOn 隐藏表达式
func (i MPropertyItem) HiddenOn(value string) MPropertyItem {
	return i.set("hiddenOn", value)
}
