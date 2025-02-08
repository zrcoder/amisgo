package model

// ListBodyField represents a list body field control
type ListBodyField Schema

func NewListBodyField() ListBodyField {
	return make(ListBodyField)
}

// Copyable enables the copy feature
func (lb ListBodyField) Copyable(value bool) ListBodyField {
	lb.set("copyable", value)
	return lb
}

// InnerClassName sets the CSS class name for the inner component
func (lb ListBodyField) InnerClassName(value string) ListBodyField {
	lb.set("innerClassName", value)
	return lb
}

// Label sets the column header
func (lb ListBodyField) Label(value string) ListBodyField {
	lb.set("label", value)
	return lb
}

// LabelClassName sets the CSS class name for the label
func (lb ListBodyField) LabelClassName(value string) ListBodyField {
	lb.set("labelClassName", value)
	return lb
}

// Name sets the bound field name
func (lb ListBodyField) Name(value string) ListBodyField {
	lb.set("name", value)
	return lb
}

// PopOver sets the pop-over content
func (lb ListBodyField) PopOver(value string) ListBodyField {
	lb.set("popOver", value)
	return lb
}

// QuickEdit enables quick edit feature
func (lb ListBodyField) QuickEdit(value string) ListBodyField {
	lb.set("quickEdit", value)
	return lb
}

// set assigns a value to a key
func (lb ListBodyField) set(key string, value any) ListBodyField {
	lb[key] = value
	return lb
}
