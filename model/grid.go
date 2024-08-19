package model

type GridSchema Schema

func Grid() GridSchema {
	return GridSchema{}
}

func (g GridSchema) set(key string, value any) GridSchema {
	g[key] = value
	return g
}

// X 格子起始位置的横坐标
func (g GridSchema) X(value int) GridSchema {
	return g.set("x", value)
}

// Y 格子起始位置的纵坐标
func (g GridSchema) Y(value int) GridSchema {
	return g.set("y", value)
}

// W 格子横跨几个宽度
func (g GridSchema) W(value int) GridSchema {
	return g.set("w", value)
}

// H 格子纵跨几个高度
func (g GridSchema) H(value int) GridSchema {
	return g.set("h", value)
}

// Width 格子所在列的宽度 int/string/"auto"
func (g GridSchema) Width(value any) GridSchema {
	return g.set("width", value)
}

// Height 格子所在行的高度 int/string/"auto"
func (g GridSchema) Height(value any) GridSchema {
	return g.set("width", value)
}

// Align 格子内容水平布局 left｜ center ｜ right ｜ auto
func (g GridSchema) Align(value string) GridSchema {
	return g.set("align", value)
}

// Valign 格子内容垂直布局 top | bottom | middle ｜ auto
func (g GridSchema) Valign(value string) GridSchema {
	return g.set("valign", value)
}
