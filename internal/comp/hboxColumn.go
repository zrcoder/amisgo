package comp

import "github.com/zrcoder/amisgo/schema"

// HBoxColumn represents a column renderer
type HBoxColumn schema.Schema

func NewHBoxColumn() HBoxColumn {
	return make(HBoxColumn)
}

func (h HBoxColumn) set(key string, value any) HBoxColumn {
	h[key] = value
	return h
}

// Body sets the content area
func (h HBoxColumn) Body(value ...any) HBoxColumn {
	return h.set("body", value)
}

// ColumnClassName sets the CSS class name for the column
func (h HBoxColumn) ColumnClassName(value string) HBoxColumn {
	return h.set("columnClassName", value)
}

// Height sets the height
func (h HBoxColumn) Height(value string) HBoxColumn {
	return h.set("height", value)
}

// Horizontal sets the horizontal layout ratio
func (h HBoxColumn) Horizontal(value string) HBoxColumn {
	return h.set("horizontal", value)
}

// Mode sets the display mode for sub-form items
func (h HBoxColumn) Mode(value string) HBoxColumn {
	return h.set("mode", value)
}

// Style sets additional styles
func (h HBoxColumn) Style(value any) HBoxColumn {
	return h.set("style", value)
}

// Valign sets the vertical alignment
func (h HBoxColumn) Valign(value string) HBoxColumn {
	return h.set("valign", value)
}

// Visible sets the visibility
func (h HBoxColumn) Visible(value bool) HBoxColumn {
	return h.set("visible", value)
}

// VisibleOn sets the visibility expression
func (h HBoxColumn) VisibleOn(value string) HBoxColumn {
	return h.set("visibleOn", value)
}

// Width sets the width
func (h HBoxColumn) Width(value string) HBoxColumn {
	return h.set("width", value)
}
