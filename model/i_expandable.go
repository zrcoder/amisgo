package model

// Expandable represents an Expandable renderer
type Expandable Schema

func NewExpandable() Expandable {
	return make(Expandable)
}

func (e Expandable) set(key string, value any) Expandable {
	e[key] = value
	return e
}

// ExpandableOn sets the expression for whether a row is expandable
func (e Expandable) ExpandableOn(value string) Expandable {
	return e.set("expandableOn", value)
}

// ExpandedRowClassNameExpr sets the expression for the expanded row class name
func (e Expandable) ExpandedRowClassNameExpr(value string) Expandable {
	return e.set("expandedRowClassNameExpr", value)
}

// ExpandedRowKeys sets the keys of the expanded rows
func (e Expandable) ExpandedRowKeys(value string) Expandable {
	return e.set("expandedRowKeys", value)
}

// ExpandedRowKeysExpr sets the expression for the keys of the expanded rows
func (e Expandable) ExpandedRowKeysExpr(value string) Expandable {
	return e.set("expandedRowKeysExpr", value)
}

// KeyField sets the key field corresponding to the data source
func (e Expandable) KeyField(value string) Expandable {
	return e.set("keyField", value)
}
