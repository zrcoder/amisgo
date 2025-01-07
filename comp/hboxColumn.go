package comp

// hBoxColumn represents a column renderer
type hBoxColumn Schema

// HBoxColumn creates a new hBoxColumn instance
func HBoxColumn() hBoxColumn {
	return make(hBoxColumn)
}

func (h hBoxColumn) set(key string, value any) hBoxColumn {
	h[key] = value
	return h
}

// Body sets the content area
func (h hBoxColumn) Body(value ...any) hBoxColumn {
	return h.set("body", value)
}

// ColumnClassName sets the CSS class name for the column
func (h hBoxColumn) ColumnClassName(value string) hBoxColumn {
	return h.set("columnClassName", value)
}

// Height sets the height
func (h hBoxColumn) Height(value string) hBoxColumn {
	return h.set("height", value)
}

// Horizontal sets the horizontal layout ratio
func (h hBoxColumn) Horizontal(value string) hBoxColumn {
	return h.set("horizontal", value)
}

// Mode sets the display mode for sub-form items
func (h hBoxColumn) Mode(value string) hBoxColumn {
	return h.set("mode", value)
}

// Style sets additional styles
func (h hBoxColumn) Style(value any) hBoxColumn {
	return h.set("style", value)
}

// Valign sets the vertical alignment
func (h hBoxColumn) Valign(value string) hBoxColumn {
	return h.set("valign", value)
}

// Visible sets the visibility
func (h hBoxColumn) Visible(value bool) hBoxColumn {
	return h.set("visible", value)
}

// VisibleOn sets the visibility expression
func (h hBoxColumn) VisibleOn(value string) hBoxColumn {
	return h.set("visibleOn", value)
}

// Width sets the width
func (h hBoxColumn) Width(value string) hBoxColumn {
	return h.set("width", value)
}
