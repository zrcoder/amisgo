package comp

import "github.com/zrcoder/amisgo/schema"

// Property
type Property schema.Schema

func NewProperty() Property {
	return Property{"type": "property"}
}

func (p Property) set(key string, value any) Property {
	p[key] = value
	return p
}

// ClassName sets the class name of the outer DOM
func (p Property) ClassName(value string) Property {
	return p.set("className", value)
}

// Column sets the number of columns per row
func (p Property) Column(value string) Property {
	return p.set("column", value)
}

// ContentStyle sets the style of the property value
func (p Property) ContentStyle(value any) Property {
	return p.set("contentStyle", value)
}

// Items sets the property items
func (p Property) Items(value ...any) Property {
	return p.set("items", value)
}

// LabelStyle sets the style of the property label
func (p Property) LabelStyle(value any) Property {
	return p.set("labelStyle", value)
}

// Mode sets the display mode
func (p Property) Mode(value string) Property {
	return p.set("mode", value)
}

// Separator sets the separator between property name and value in mode
func (p Property) Separator(value string) Property {
	return p.set("separator", value)
}

// Source sets the data source
func (p Property) Source(value string) Property {
	return p.set("source", value)
}

// Style sets the style of the outer DOM
func (p Property) Style(value any) Property {
	return p.set("style", value)
}

// Title sets the title
func (p Property) Title(value any) Property {
	return p.set("title", value)
}
