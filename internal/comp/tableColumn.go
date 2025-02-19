package comp

import "github.com/zrcoder/amisgo/schema"

// TableColumn represents a table column with default type as text.
type TableColumn schema.Schema

func NewTableColumn() TableColumn {
	return TableColumn{}
}

func (c TableColumn) set(key string, value any) TableColumn {
	c[key] = value
	return c
}

// Align sets the column alignment.
func (c TableColumn) Align(value string) TableColumn {
	return c.set("align", value)
}

// Breakpoint sets the breakpoint for footable.
func (c TableColumn) Breakpoint(value string) TableColumn {
	return c.set("breakpoint", value)
}

// CanAccessSuperData allows the column to access parent data.
func (c TableColumn) CanAccessSuperData(value bool) TableColumn {
	return c.set("canAccessSuperData", value)
}

// ClassName sets the column's CSS class.
func (c TableColumn) ClassName(value string) TableColumn {
	return c.set("className", value)
}

// ClassNameExpr sets the cell's CSS class expression.
func (c TableColumn) ClassNameExpr(value string) TableColumn {
	return c.set("classNameExpr", value)
}

// Copyable enables copy functionality for the column.
func (c TableColumn) Copyable(value bool) TableColumn {
	return c.set("copyable", value)
}

// Filterable enables filtering for the column.
func (c TableColumn) Filterable(value bool) TableColumn {
	return c.set("filterable", value)
}

// Fixed sets the column as fixed.
func (c TableColumn) Fixed(value string) TableColumn {
	return c.set("fixed", value)
}

// HeaderAlign sets the header alignment.
func (c TableColumn) HeaderAlign(value string) TableColumn {
	return c.set("headerAlign", value)
}

// InnerStyle sets custom styles for inner components.
func (c TableColumn) InnerStyle(value any) TableColumn {
	return c.set("innerStyle", value)
}

// Label sets the column label.
func (c TableColumn) Label(value string) TableColumn {
	return c.set("label", value)
}

// LabelClassName sets the CSS class for the column header.
func (c TableColumn) LabelClassName(value string) TableColumn {
	return c.set("labelClassName", value)
}

// LazyRenderAfter sets lazy rendering for many columns.
func (c TableColumn) LazyRenderAfter(value string) TableColumn {
	return c.set("lazyRenderAfter", value)
}

// Name sets the field name the column binds to.
func (c TableColumn) Name(value string) TableColumn {
	return c.set("name", value)
}

// PopOver enables the popover for details.
func (c TableColumn) PopOver(value string) TableColumn {
	return c.set("popOver", value)
}

// QuickEdit enables quick edit functionality.
func (c TableColumn) QuickEdit(value string) TableColumn {
	return c.set("quickEdit", value)
}

// QuickEditOnUpdate sets quick edit panel for updates.
func (c TableColumn) QuickEditOnUpdate(value string) TableColumn {
	return c.set("quickEditOnUpdate", value)
}

// Remark sets the tooltip text.
func (c TableColumn) Remark(value string) TableColumn {
	return c.set("remark", value)
}

// Searchable enables quick search for the column.
func (c TableColumn) Searchable(value bool) TableColumn {
	return c.set("searchable", value)
}

// Sortable enables sorting for the column.
func (c TableColumn) Sortable(value bool) TableColumn {
	return c.set("sortable", value)
}

// Toggled sets the column's default visibility.
func (c TableColumn) Toggled(value bool) TableColumn {
	return c.set("toggled", value)
}

// Unique sets the column as unique.
func (c TableColumn) Unique(value bool) TableColumn {
	return c.set("unique", value)
}

// VAlign sets the vertical alignment.
func (c TableColumn) VAlign(value string) TableColumn {
	return c.set("vAlign", value)
}

// Value sets the default value for the column.
func (c TableColumn) Value(value string) TableColumn {
	return c.set("value", value)
}

// Width sets the column width.
func (c TableColumn) Width(value string) TableColumn {
	return c.set("width", value)
}
