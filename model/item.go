package model

type PropertyItemSchema Schema

func PropertyItem() PropertyItemSchema {
	return PropertyItemSchema{}
}

func (i PropertyItemSchema) set(key string, value any) PropertyItemSchema {
	i[key] = value
	return i
}

// Label 属性名
func (i PropertyItemSchema) Label(value string) PropertyItemSchema {
	return i.set("label", value)
}

// Content 属性值
func (i PropertyItemSchema) Content(value any) PropertyItemSchema {
	return i.set("content", value)
}

// Span 属性值跨几列
func (i PropertyItemSchema) Span(value int) PropertyItemSchema {
	return i.set("span", value)
}

// VisibleOn 显示表达式
func (i PropertyItemSchema) VisibleOn(value string) PropertyItemSchema {
	return i.set("visibleOn", value)
}

// HiddenOn 隐藏表达式
func (i PropertyItemSchema) HiddenOn(value string) PropertyItemSchema {
	return i.set("hiddenOn", value)
}
