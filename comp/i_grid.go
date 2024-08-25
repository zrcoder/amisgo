package comp

type gridItem schema

func GridItem() gridItem {
	return gridItem{}
}

func (g gridItem) set(key string, value any) gridItem {
	g[key] = value
	return g
}

// X 格子起始位置的横坐标
func (g gridItem) X(value int) gridItem {
	return g.set("x", value)
}

// Y 格子起始位置的纵坐标
func (g gridItem) Y(value int) gridItem {
	return g.set("y", value)
}

// W 格子横跨几个宽度
func (g gridItem) W(value int) gridItem {
	return g.set("w", value)
}

// H 格子纵跨几个高度
func (g gridItem) H(value int) gridItem {
	return g.set("h", value)
}

// Width 格子所在列的宽度 int/string/"auto"
func (g gridItem) Width(value any) gridItem {
	return g.set("width", value)
}

// Height 格子所在行的高度 int/string/"auto"
func (g gridItem) Height(value any) gridItem {
	return g.set("width", value)
}

// Align 格子内容水平布局 left｜ center ｜ right ｜ auto
func (g gridItem) Align(value string) gridItem {
	return g.set("align", value)
}

// Valign 格子内容垂直布局 top | bottom | middle ｜ auto
func (g gridItem) Valign(value string) gridItem {
	return g.set("valign", value)
}
