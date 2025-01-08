package comp

import "github.com/zrcoder/amisgo/model"

// gridColumn represents a grid column
type gridColumn model.Schema

// GridColumn creates a new gridColumn instance
func GridColumn() gridColumn {
	return make(gridColumn)
}

func (c gridColumn) set(key string, value any) gridColumn {
	c[key] = value
	return c
}

// Body sets the content of the column
func (c gridColumn) Body(value ...any) gridColumn {
	return c.set("body", value)
}

// ColumnClassName sets the CSS class name of the column
func (c gridColumn) ColumnClassName(value string) gridColumn {
	return c.set("columnClassName", value)
}

// Horizontal sets the horizontal width ratio for horizontal layout
func (c gridColumn) Horizontal(value string) gridColumn {
	return c.set("horizontal", value)
}

// ID sets the unique ID of the component
func (c gridColumn) ID(value string) gridColumn {
	return c.set("id", value)
}

// Lg sets the width ratio for large screens (>=1200px)
func (c gridColumn) Lg(value string) gridColumn {
	return c.set("lg", value)
}

// Md sets the width ratio for medium screens (>=992px)
func (c gridColumn) Md(value string) gridColumn {
	return c.set("md", value)
}

// Mode sets the default display mode of sub-form items
func (c gridColumn) Mode(value string) gridColumn {
	return c.set("mode", value)
}

// Sm sets the width ratio for small screens (>=768px)
func (c gridColumn) Sm(value string) gridColumn {
	return c.set("sm", value)
}

// Style sets the style of the column
func (c gridColumn) Style(value any) gridColumn {
	return c.set("style", value)
}

// ThemeCSS sets the theme CSS of the column
func (c gridColumn) ThemeCSS(value any) gridColumn {
	return c.set("themeCss", value)
}

// Valign sets the vertical alignment of the column
func (c gridColumn) Valign(value string) gridColumn {
	return c.set("valign", value)
}

// WrapperCustomStyle sets the custom wrapper style of the column
func (c gridColumn) WrapperCustomStyle(value any) gridColumn {
	return c.set("wrapperCustomStyle", value)
}

// Xs sets the width ratio for extra small screens (<768px)
func (c gridColumn) Xs(value string) gridColumn {
	return c.set("xs", value)
}
