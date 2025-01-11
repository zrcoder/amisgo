package comp

import "github.com/zrcoder/amisgo/model"

// property

type property model.Schema

// Property creates a new Property instance
func Property() property {
	return property{"type": "property"}
}

func (p property) set(key string, value any) property {
	p[key] = value
	return p
}

// ClassName sets the class name of the outer DOM
func (p property) ClassName(value string) property {
	return p.set("className", value)
}

// Column sets the number of columns per row
func (p property) Column(value string) property {
	return p.set("column", value)
}

// ContentStyle sets the style of the property value
func (p property) ContentStyle(value any) property {
	return p.set("contentStyle", value)
}

// Items sets the property items
func (p property) Items(value ...any) property {
	return p.set("items", value)
}

// LabelStyle sets the style of the property label
func (p property) LabelStyle(value any) property {
	return p.set("labelStyle", value)
}

// Mode sets the display mode
func (p property) Mode(value string) property {
	return p.set("mode", value)
}

// Separator sets the separator between property name and value in mode
func (p property) Separator(value string) property {
	return p.set("separator", value)
}

// Source sets the data source
func (p property) Source(value string) property {
	return p.set("source", value)
}

// Style sets the style of the outer DOM
func (p property) Style(value any) property {
	return p.set("style", value)
}

// Title sets the title
func (p property) Title(value any) property {
	return p.set("title", value)
}
