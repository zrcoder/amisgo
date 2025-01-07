package comp

type MPropertyItem Schema

func PropertyItem() MPropertyItem {
	return MPropertyItem{}
}

func (i MPropertyItem) set(key string, value any) MPropertyItem {
	i[key] = value
	return i
}

// Label sets the label property
func (i MPropertyItem) Label(value string) MPropertyItem {
	return i.set("label", value)
}

// Content sets the content property
func (i MPropertyItem) Content(value any) MPropertyItem {
	return i.set("content", value)
}

// Span sets the span property
func (i MPropertyItem) Span(value int) MPropertyItem {
	return i.set("span", value)
}

// VisibleOn sets the visibleOn expression
func (i MPropertyItem) VisibleOn(value string) MPropertyItem {
	return i.set("visibleOn", value)
}

// HiddenOn sets the hiddenOn expression
func (i MPropertyItem) HiddenOn(value string) MPropertyItem {
	return i.set("hiddenOn", value)
}
