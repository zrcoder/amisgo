package comp

type MGridItem Schema

func GridItem() MGridItem {
	return MGridItem{}
}

func (g MGridItem) set(key string, value any) MGridItem {
	g[key] = value
	return g
}

// X 格子起始位置的横坐标
func (g MGridItem) X(value int) MGridItem {
	return g.set("x", value)
}

// Y 格子起始位置的纵坐标
func (g MGridItem) Y(value int) MGridItem {
	return g.set("y", value)
}

// W 格子横跨几个宽度
func (g MGridItem) W(value int) MGridItem {
	return g.set("w", value)
}

// H 格子纵跨几个高度
func (g MGridItem) H(value int) MGridItem {
	return g.set("h", value)
}

// Width 格子所在列的宽度 int/string/"auto"
func (g MGridItem) Width(value any) MGridItem {
	return g.set("width", value)
}

// Height 格子所在行的高度 int/string/"auto"
func (g MGridItem) Height(value any) MGridItem {
	return g.set("width", value)
}

// Align 格子内容水平布局 left｜ center ｜ right ｜ auto
func (g MGridItem) Align(value string) MGridItem {
	return g.set("align", value)
}

// Valign 格子内容垂直布局 top | bottom | middle ｜ auto
func (g MGridItem) Valign(value string) MGridItem {
	return g.set("valign", value)
}
