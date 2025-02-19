package comp

import "github.com/zrcoder/amisgo/schema"

type PropertyItem schema.Schema

func NewPropertyItem() PropertyItem {
	return PropertyItem{}
}

func (i PropertyItem) set(key string, value any) PropertyItem {
	i[key] = value
	return i
}

// Label sets the label property
func (i PropertyItem) Label(value string) PropertyItem {
	return i.set("label", value)
}

// Content sets the content property
func (i PropertyItem) Content(value any) PropertyItem {
	return i.set("content", value)
}

// Span sets the span property
func (i PropertyItem) Span(value int) PropertyItem {
	return i.set("span", value)
}

// VisibleOn sets the visibleOn expression
func (i PropertyItem) VisibleOn(value string) PropertyItem {
	return i.set("visibleOn", value)
}

// HiddenOn sets the hiddenOn expression
func (i PropertyItem) HiddenOn(value string) PropertyItem {
	return i.set("hiddenOn", value)
}
