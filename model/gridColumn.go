package model

// GridColumn represents a grid column
type GridColumn Schema

func NewGridColumn() GridColumn {
	return make(GridColumn)
}

func (c GridColumn) set(key string, value any) GridColumn {
	c[key] = value
	return c
}

// Body sets the content of the column
func (c GridColumn) Body(value ...any) GridColumn {
	return c.set("body", value)
}

// ColumnClassName sets the CSS class name of the column
func (c GridColumn) ColumnClassName(value string) GridColumn {
	return c.set("columnClassName", value)
}

// Horizontal sets the horizontal width ratio for horizontal layout
func (c GridColumn) Horizontal(value string) GridColumn {
	return c.set("horizontal", value)
}

// ID sets the unique ID of the component
func (c GridColumn) ID(value string) GridColumn {
	return c.set("id", value)
}

// Lg sets the width ratio for large screens (>=1200px)
func (c GridColumn) Lg(value string) GridColumn {
	return c.set("lg", value)
}

// Md sets the width ratio for medium screens (>=992px)
func (c GridColumn) Md(value string) GridColumn {
	return c.set("md", value)
}

// Mode sets the default display mode of sub-form items
func (c GridColumn) Mode(value string) GridColumn {
	return c.set("mode", value)
}

// Sm sets the width ratio for small screens (>=768px)
func (c GridColumn) Sm(value string) GridColumn {
	return c.set("sm", value)
}

// Style sets the style of the column
func (c GridColumn) Style(value any) GridColumn {
	return c.set("style", value)
}

// ThemeCSS sets the theme CSS of the column
func (c GridColumn) ThemeCSS(value any) GridColumn {
	return c.set("themeCss", value)
}

// Valign sets the vertical alignment of the column
func (c GridColumn) Valign(value string) GridColumn {
	return c.set("valign", value)
}

// WrapperCustomStyle sets the custom wrapper style of the column
func (c GridColumn) WrapperCustomStyle(value any) GridColumn {
	return c.set("wrapperCustomStyle", value)
}

// Xs sets the width ratio for extra small screens (<768px)
func (c GridColumn) Xs(value string) GridColumn {
	return c.set("xs", value)
}
