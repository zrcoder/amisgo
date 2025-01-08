package comp

import "github.com/zrcoder/amisgo/model"

type MGridItem model.Schema

func GridItem() MGridItem {
	return MGridItem{}
}

func (g MGridItem) set(key string, value any) MGridItem {
	g[key] = value
	return g
}

// X sets the starting X coordinate
func (g MGridItem) X(value int) MGridItem {
	return g.set("x", value)
}

// Y sets the starting Y coordinate
func (g MGridItem) Y(value int) MGridItem {
	return g.set("y", value)
}

// W sets the width span
func (g MGridItem) W(value int) MGridItem {
	return g.set("w", value)
}

// H sets the height span
func (g MGridItem) H(value int) MGridItem {
	return g.set("h", value)
}

// Width sets the column width (int/string/"auto")
func (g MGridItem) Width(value any) MGridItem {
	return g.set("width", value)
}

// Height sets the row height (int/string/"auto")
func (g MGridItem) Height(value any) MGridItem {
	return g.set("height", value)
}

// Align sets the horizontal alignment (left/center/right/auto)
func (g MGridItem) Align(value string) MGridItem {
	return g.set("align", value)
}

// Valign sets the vertical alignment (top/bottom/middle/auto)
func (g MGridItem) Valign(value string) MGridItem {
	return g.set("valign", value)
}
