package comp

// property

type property Schema

// Property 创建一个新的 Property 实例
func Property() property {
	return property{}.
		set("type", "property")
}

func (p property) set(key string, value any) property {
	p[key] = value
	return p
}

// ClassName 外层 dom 的类名
func (p property) ClassName(value string) property {
	return p.set("className", value)
}

// Column 每行几列
func (p property) Column(value string) property {
	return p.set("column", value)
}

// ContentStyle 属性值的样式
func (p property) ContentStyle(value any) property {
	return p.set("contentStyle", value)
}

// Items 属性项
func (p property) Items(value ...any) property {
	return p.set("items", value)
}

// LabelStyle 属性名的样式
func (p property) LabelStyle(value any) property {
	return p.set("labelStyle", value)
}

// Mode 显示模式
func (p property) Mode(value string) property {
	return p.set("mode", value)
}

// Separator 模式下属性名和值之间的分隔符
func (p property) Separator(value string) property {
	return p.set("separator", value)
}

// Source 数据源
func (p property) Source(value string) property {
	return p.set("source", value)
}

// Style 外层 dom 的样式
func (p property) Style(value any) property {
	return p.set("style", value)
}

// Title 标题
func (p property) Title(value any) property {
	return p.set("title", value)
}
