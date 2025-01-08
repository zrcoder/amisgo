package comp

import "github.com/zrcoder/amisgo/model"

type MColumn model.Schema

func Column() MColumn {
	return make(MColumn)
}

// Set sets a key-value pair
func (c MColumn) Set(key string, value any) MColumn {
	c[key] = value
	return c
}

// Align sets the alignment
func (c MColumn) Align(value string) MColumn {
	return c.Set("align", value)
}

// Body sets the body content
func (c MColumn) Body(value ...any) MColumn {
	return c.Set("body", value)
}

// CanAccessSuperData sets whether the cell can access parent data
func (c MColumn) CanAccessSuperData(value bool) MColumn {
	return c.Set("canAccessSuperData", value)
}

// Children sets the header group
func (c MColumn) Children(value string) MColumn {
	return c.Set("children", value)
}

// ClassName sets the column class name
func (c MColumn) ClassName(value string) MColumn {
	return c.Set("className", value)
}

// ClassNameExpr sets the cell class name expression
func (c MColumn) ClassNameExpr(value string) MColumn {
	return c.Set("classNameExpr", value)
}

// ColSpanExpr sets the column span expression
func (c MColumn) ColSpanExpr(value string) MColumn {
	return c.Set("colSpanExpr", value)
}

// Copyable sets whether the column is copyable
func (c MColumn) Copyable(value bool) MColumn {
	return c.Set("copyable", value)
}

// Filterable sets whether the column is filterable
func (c MColumn) Filterable(value bool) MColumn {
	return c.Set("filterable", value)
}

// Fixed sets whether the column is fixed
func (c MColumn) Fixed(value string) MColumn {
	return c.Set("fixed", value)
}

// Type sets the column type
func (c MColumn) Type(value string) MColumn {
	return c.Set("type", value)
}

// Name sets the column unique identifier
func (c MColumn) Name(value string) MColumn {
	return c.Set("name", value)
}

// Label sets the column label
func (c MColumn) Label(value string) MColumn {
	return c.Set("label", value)
}

// QuickEdit sets the quick edit configuration
func (c MColumn) QuickEdit(value string) MColumn {
	return c.Set("quickEdit", value)
}

// Remark sets the header remark
func (c MColumn) Remark(value string) MColumn {
	return c.Set("remark", value)
}

// RowSpanExpr sets the row span expression
func (c MColumn) RowSpanExpr(value string) MColumn {
	return c.Set("rowSpanExpr", value)
}

// Searchable sets whether the column is searchable
func (c MColumn) Searchable(value any) MColumn {
	return c.Set("searchable", value)
}

// Sortable sets whether the column is sortable
func (c MColumn) Sortable(value bool) MColumn {
	return c.Set("sortable", value)
}

// Sorter sets the sorter
func (c MColumn) Sorter(value bool) MColumn {
	return c.Set("sorter", value)
}

// Title sets the column title
func (c MColumn) Title(value any) MColumn {
	return c.Set("title", value)
}

// TitleClassName sets the header cell class name
func (c MColumn) TitleClassName(value string) MColumn {
	return c.Set("titleClassName", value)
}

// Togged sets whether the column is toggled
func (c MColumn) Togged(value bool) MColumn {
	return c.Set("toggled", value)
}

// Width sets the column width
func (c MColumn) Width(value string) MColumn {
	return c.Set("width", value)
}

// ColumnClassName sets the column class name
func (c MColumn) ColumnClassName(value string) MColumn {
	return c.Set("columnClassName", value)
}

// Xs sets the width ratio for extra small screens
func (c MColumn) Xs(value any) MColumn {
	return c.Set("xs", value)
}

// Sm sets the width ratio for small screens
func (c MColumn) Sm(value any) MColumn {
	return c.Set("sm", value)
}

// Md sets the width ratio for medium screens
func (c MColumn) Md(value any) MColumn {
	return c.Set("md", value)
}

// Lg sets the width ratio for large screens
func (c MColumn) Lg(value any) MColumn {
	return c.Set("lg", value)
}

// Valign sets the vertical alignment
func (c MColumn) Valign(value string) MColumn {
	return c.Set("valign", value)
}

// Buttons sets the buttons
func (c MColumn) Buttons(value ...action) MColumn {
	return c.Set("buttons", value)
}
