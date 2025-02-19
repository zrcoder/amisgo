package comp

import "github.com/zrcoder/amisgo/schema"

type GridItem schema.Schema

func NewGridItem() GridItem {
	return GridItem{}
}

func (g GridItem) set(key string, value any) GridItem {
	g[key] = value
	return g
}

// X sets the starting X coordinate
func (g GridItem) X(value int) GridItem {
	return g.set("x", value)
}

// Y sets the starting Y coordinate
func (g GridItem) Y(value int) GridItem {
	return g.set("y", value)
}

// W sets the width span
func (g GridItem) W(value int) GridItem {
	return g.set("w", value)
}

// H sets the height span
func (g GridItem) H(value int) GridItem {
	return g.set("h", value)
}

// Width sets the column width (int/string/"auto")
func (g GridItem) Width(value any) GridItem {
	return g.set("width", value)
}

// Height sets the row height (int/string/"auto")
func (g GridItem) Height(value any) GridItem {
	return g.set("height", value)
}

// Align sets the horizontal alignment (left/center/right/auto)
func (g GridItem) Align(value string) GridItem {
	return g.set("align", value)
}

// Valign sets the vertical alignment (top/bottom/middle/auto)
func (g GridItem) Valign(value string) GridItem {
	return g.set("valign", value)
}
