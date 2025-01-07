package comp

// tableColumn represents a table column with default type as text.

type tableColumn Schema

// TableColumn creates a new TableColumn instance.
func TableColumn() tableColumn {
	return tableColumn{}
}

func (c tableColumn) set(key string, value any) tableColumn {
	c[key] = value
	return c
}

// Align sets the column alignment.
func (c tableColumn) Align(value string) tableColumn {
	return c.set("align", value)
}

// Breakpoint sets the breakpoint for footable.
func (c tableColumn) Breakpoint(value string) tableColumn {
	return c.set("breakpoint", value)
}

// CanAccessSuperData allows the column to access parent data.
func (c tableColumn) CanAccessSuperData(value bool) tableColumn {
	return c.set("canAccessSuperData", value)
}

// ClassName sets the column's CSS class.
func (c tableColumn) ClassName(value string) tableColumn {
	return c.set("className", value)
}

// ClassNameExpr sets the cell's CSS class expression.
func (c tableColumn) ClassNameExpr(value string) tableColumn {
	return c.set("classNameExpr", value)
}

// Copyable enables copy functionality for the column.
func (c tableColumn) Copyable(value bool) tableColumn {
	return c.set("copyable", value)
}

// Filterable enables filtering for the column.
func (c tableColumn) Filterable(value bool) tableColumn {
	return c.set("filterable", value)
}

// Fixed sets the column as fixed.
func (c tableColumn) Fixed(value string) tableColumn {
	return c.set("fixed", value)
}

// HeaderAlign sets the header alignment.
func (c tableColumn) HeaderAlign(value string) tableColumn {
	return c.set("headerAlign", value)
}

// InnerStyle sets custom styles for inner components.
func (c tableColumn) InnerStyle(value any) tableColumn {
	return c.set("innerStyle", value)
}

// Label sets the column label.
func (c tableColumn) Label(value string) tableColumn {
	return c.set("label", value)
}

// LabelClassName sets the CSS class for the column header.
func (c tableColumn) LabelClassName(value string) tableColumn {
	return c.set("labelClassName", value)
}

// LazyRenderAfter sets lazy rendering for many columns.
func (c tableColumn) LazyRenderAfter(value string) tableColumn {
	return c.set("lazyRenderAfter", value)
}

// Name sets the field name the column binds to.
func (c tableColumn) Name(value string) tableColumn {
	return c.set("name", value)
}

// PopOver enables the popover for details.
func (c tableColumn) PopOver(value string) tableColumn {
	return c.set("popOver", value)
}

// QuickEdit enables quick edit functionality.
func (c tableColumn) QuickEdit(value string) tableColumn {
	return c.set("quickEdit", value)
}

// QuickEditOnUpdate sets quick edit panel for updates.
func (c tableColumn) QuickEditOnUpdate(value string) tableColumn {
	return c.set("quickEditOnUpdate", value)
}

// Remark sets the tooltip text.
func (c tableColumn) Remark(value string) tableColumn {
	return c.set("remark", value)
}

// Searchable enables quick search for the column.
func (c tableColumn) Searchable(value bool) tableColumn {
	return c.set("searchable", value)
}

// Sortable enables sorting for the column.
func (c tableColumn) Sortable(value bool) tableColumn {
	return c.set("sortable", value)
}

// Toggled sets the column's default visibility.
func (c tableColumn) Toggled(value bool) tableColumn {
	return c.set("toggled", value)
}

// Unique sets the column as unique.
func (c tableColumn) Unique(value bool) tableColumn {
	return c.set("unique", value)
}

// VAlign sets the vertical alignment.
func (c tableColumn) VAlign(value string) tableColumn {
	return c.set("vAlign", value)
}

// Value sets the default value for the column.
func (c tableColumn) Value(value string) tableColumn {
	return c.set("value", value)
}

// Width sets the column width.
func (c tableColumn) Width(value string) tableColumn {
	return c.set("width", value)
}
