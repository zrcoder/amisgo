package comp

import "github.com/zrcoder/amisgo/schema"

type Column schema.Schema

func NewColumn(component ...any) Column {
	if len(component) == 0 {
		return Column{}
	}
	s := toSchema(component[0])
	return Column(s)
}

// set sets a key-value pair
func (c Column) set(key string, value any) Column {
	c[key] = value
	return c
}

// Align sets the alignment
func (c Column) Align(value string) Column {
	return c.set("align", value)
}

// Body sets the body content
func (c Column) Body(value ...any) Column {
	return c.set("body", value)
}

// CanAccessSuperData sets whether the cell can access parent data
func (c Column) CanAccessSuperData(value bool) Column {
	return c.set("canAccessSuperData", value)
}

// Children sets the header group
func (c Column) Children(value string) Column {
	return c.set("children", value)
}

// ClassName sets the column class name
func (c Column) ClassName(value string) Column {
	return c.set("className", value)
}

// ClassNameExpr sets the cell class name expression
func (c Column) ClassNameExpr(value string) Column {
	return c.set("classNameExpr", value)
}

// ColSpanExpr sets the column span expression
func (c Column) ColSpanExpr(value string) Column {
	return c.set("colSpanExpr", value)
}

// Copyable sets whether the column is copyable
func (c Column) Copyable(value bool) Column {
	return c.set("copyable", value)
}

// Filterable sets whether the column is filterable
func (c Column) Filterable(value bool) Column {
	return c.set("filterable", value)
}

// Fixed sets whether the column is fixed
func (c Column) Fixed(value string) Column {
	return c.set("fixed", value)
}

// Name sets the column unique identifier
func (c Column) Name(value string) Column {
	return c.set("name", value)
}

// Label sets the column label
func (c Column) Label(value string) Column {
	return c.set("label", value)
}

// QuickEdit sets the quick edit configuration
func (c Column) QuickEdit(value any) Column {
	return c.set("quickEdit", value)
}

// Remark sets the header remark
func (c Column) Remark(value string) Column {
	return c.set("remark", value)
}

// RowSpanExpr sets the row span expression
func (c Column) RowSpanExpr(value string) Column {
	return c.set("rowSpanExpr", value)
}

// Searchable sets whether the column is searchable
func (c Column) Searchable(value any) Column {
	return c.set("searchable", value)
}

// Sortable sets whether the column is sortable
func (c Column) Sortable(value bool) Column {
	return c.set("sortable", value)
}

// Sorter sets the sorter
func (c Column) Sorter(value bool) Column {
	return c.set("sorter", value)
}

// Title sets the column title
func (c Column) Title(value any) Column {
	return c.set("title", value)
}

// TitleClassName sets the header cell class name
func (c Column) TitleClassName(value string) Column {
	return c.set("titleClassName", value)
}

// Togged sets whether the column is toggled
func (c Column) Togged(value bool) Column {
	return c.set("toggled", value)
}

// Width sets the column width
func (c Column) Width(value string) Column {
	return c.set("width", value)
}

// ColumnClassName sets the column class name
func (c Column) ColumnClassName(value string) Column {
	return c.set("columnClassName", value)
}

// Xs sets the width ratio for extra small screens
func (c Column) Xs(value any) Column {
	return c.set("xs", value)
}

// Sm sets the width ratio for small screens
func (c Column) Sm(value any) Column {
	return c.set("sm", value)
}

// Md sets the width ratio for medium screens
func (c Column) Md(value any) Column {
	return c.set("md", value)
}

// Lg sets the width ratio for large screens
func (c Column) Lg(value any) Column {
	return c.set("lg", value)
}

// Valign sets the vertical alignment, 'top' | 'middle' | 'bottom' | 'between'
func (c Column) Valign(value string) Column {
	return c.set("valign", value)
}

// Buttons sets the buttons
func (c Column) Buttons(value ...any) Column {
	return c.set("buttons", value)
}

// PopOver sets the pop-over
func (c Column) PopOver(value any) Column {
	return c.set("popOver", value)
}

// Type sets the column type, for example: operation
func (c Column) Type(value string) Column {
	return c.set("type", value)
}
