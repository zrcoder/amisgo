package comp

import "github.com/zrcoder/amisgo/model"

// gridNav represents a grid navigation component
type gridNav model.Schema

// GridNav creates a new GridNav instance
func GridNav() gridNav {
	return gridNav{"type": "grid-nav"}
}

func (g gridNav) set(key string, value any) gridNav {
	g[key] = value
	return g
}

// Border sets whether to show item borders
func (g gridNav) Border(value bool) gridNav {
	return g.set("border", value)
}

// Center sets whether to center item content
func (g gridNav) Center(value bool) gridNav {
	return g.set("center", value)
}

// ClassName sets the outer CSS class name
func (g gridNav) ClassName(value string) gridNav {
	return g.set("className", value)
}

// ColumnNum sets the number of columns
func (g gridNav) ColumnNum(value string) gridNav {
	return g.set("columnNum", value)
}

// Direction sets the item content direction (horizontal or vertical)
func (g gridNav) Direction(value string) gridNav {
	return g.set("direction", value)
}

// Gutter sets the spacing between items (default unit is px)
func (g gridNav) Gutter(value string) gridNav {
	return g.set("gutter", value)
}

// IconRatio sets the icon width ratio (unit is %)
func (g gridNav) IconRatio(value string) gridNav {
	return g.set("iconRatio", value)
}

// ItemClassName sets the item CSS class name
func (g gridNav) ItemClassName(value string) gridNav {
	return g.set("itemClassName", value)
}

// Options sets the item icons
func (g gridNav) Options(value ...any) gridNav {
	return g.set("options", value)
}

// Reverse sets whether to swap the icon and text positions
func (g gridNav) Reverse(value bool) gridNav {
	return g.set("reverse", value)
}

// Source sets the data source
func (g gridNav) Source(value string) gridNav {
	return g.set("source", value)
}

// Square sets whether to make items square
func (g gridNav) Square(value bool) gridNav {
	return g.set("square", value)
}

// Value sets the image array
func (g gridNav) Value(value string) gridNav {
	return g.set("value", value)
}
