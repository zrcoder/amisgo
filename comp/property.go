package comp

// Property
//
// @version 6.7.0
type Property Schema

// NewProperty 创建一个新的 Property 实例
func NewProperty() Property {
	return Property{}.
		set("type", "property")
}

func (p Property) set(key string, value interface{}) Property {
	p[key] = value
	return p
}

// ClassName 外层 dom 的类名
func (p Property) ClassName(value string) Property {
	return p.set("className", value)
}

// Column 每行几列
func (p Property) Column(value string) Property {
	return p.set("column", value)
}

// ContentStyle 属性值的样式
func (p Property) ContentStyle(value string) Property {
	return p.set("contentStyle", value)
}

// Items 属性项
func (p Property) Items(value string) Property {
	return p.set("items", value)
}

// LabelStyle 属性名的样式
func (p Property) LabelStyle(value string) Property {
	return p.set("labelStyle", value)
}

// Mode 显示模式
func (p Property) Mode(value string) Property {
	return p.set("mode", value)
}

// Separator 模式下属性名和值之间的分隔符
func (p Property) Separator(value string) Property {
	return p.set("separator", value)
}

// Source 数据源
func (p Property) Source(value string) Property {
	return p.set("source", value)
}

// Style 外层 dom 的样式
func (p Property) Style(value string) Property {
	return p.set("style", value)
}

// Title 标题
func (p Property) Title(value string) Property {
	return p.set("title", value)
}
