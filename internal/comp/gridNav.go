package comp

import "github.com/zrcoder/amisgo/schema"

// GridNav represents a grid navigation component
type GridNav schema.Schema

func NewGridNav() GridNav {
	return GridNav{"type": "grid-nav"}
}

func (g GridNav) set(key string, value any) GridNav {
	g[key] = value
	return g
}

// Border sets whether to show item borders
func (g GridNav) Border(value bool) GridNav {
	return g.set("border", value)
}

// Center sets whether to center item content
func (g GridNav) Center(value bool) GridNav {
	return g.set("center", value)
}

// ClassName sets the outer CSS class name
func (g GridNav) ClassName(value string) GridNav {
	return g.set("className", value)
}

// ColumnNum sets the number of columns
func (g GridNav) ColumnNum(value string) GridNav {
	return g.set("columnNum", value)
}

// Direction sets the item content direction (horizontal or vertical)
func (g GridNav) Direction(value string) GridNav {
	return g.set("direction", value)
}

// Gutter sets the spacing between items (default unit is px)
func (g GridNav) Gutter(value string) GridNav {
	return g.set("gutter", value)
}

// IconRatio sets the icon width ratio (unit is %)
func (g GridNav) IconRatio(value string) GridNav {
	return g.set("iconRatio", value)
}

// ItemClassName sets the item CSS class name
func (g GridNav) ItemClassName(value string) GridNav {
	return g.set("itemClassName", value)
}

// Options sets the item icons
func (g GridNav) Options(value ...any) GridNav {
	return g.set("options", value)
}

// Reverse sets whether to swap the icon and text positions
func (g GridNav) Reverse(value bool) GridNav {
	return g.set("reverse", value)
}

// Source sets the data source
func (g GridNav) Source(value string) GridNav {
	return g.set("source", value)
}

// Square sets whether to make items square
func (g GridNav) Square(value bool) GridNav {
	return g.set("square", value)
}

// Value sets the image array
func (g GridNav) Value(value string) GridNav {
	return g.set("value", value)
}
