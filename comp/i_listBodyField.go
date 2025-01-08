package comp

import "github.com/zrcoder/amisgo/model"

// MListBodyField represents a list body field control
type MListBodyField model.Schema

// ListBodyField creates a new ListBodyField instance
func ListBodyField() MListBodyField {
	return make(MListBodyField)
}

// Copyable enables the copy feature
func (lb MListBodyField) Copyable(value bool) MListBodyField {
	lb.set("copyable", value)
	return lb
}

// InnerClassName sets the CSS class name for the inner component
func (lb MListBodyField) InnerClassName(value string) MListBodyField {
	lb.set("innerClassName", value)
	return lb
}

// Label sets the column header
func (lb MListBodyField) Label(value string) MListBodyField {
	lb.set("label", value)
	return lb
}

// LabelClassName sets the CSS class name for the label
func (lb MListBodyField) LabelClassName(value string) MListBodyField {
	lb.set("labelClassName", value)
	return lb
}

// Name sets the bound field name
func (lb MListBodyField) Name(value string) MListBodyField {
	lb.set("name", value)
	return lb
}

// PopOver sets the pop-over content
func (lb MListBodyField) PopOver(value string) MListBodyField {
	lb.set("popOver", value)
	return lb
}

// QuickEdit enables quick edit feature
func (lb MListBodyField) QuickEdit(value string) MListBodyField {
	lb.set("quickEdit", value)
	return lb
}

// set assigns a value to a key
func (lb MListBodyField) set(key string, value any) MListBodyField {
	lb[key] = value
	return lb
}
