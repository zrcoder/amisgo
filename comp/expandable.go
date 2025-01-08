package comp

import "github.com/zrcoder/amisgo/model"

// expandable represents an expandable renderer
type expandable model.Schema

// Expandable creates a new instance of expandable
func Expandable() expandable {
	return make(expandable)
}

func (e expandable) set(key string, value any) expandable {
	e[key] = value
	return e
}

// ExpandableOn sets the expression for whether a row is expandable
func (e expandable) ExpandableOn(value string) expandable {
	return e.set("expandableOn", value)
}

// ExpandedRowClassNameExpr sets the expression for the expanded row class name
func (e expandable) ExpandedRowClassNameExpr(value string) expandable {
	return e.set("expandedRowClassNameExpr", value)
}

// ExpandedRowKeys sets the keys of the expanded rows
func (e expandable) ExpandedRowKeys(value string) expandable {
	return e.set("expandedRowKeys", value)
}

// ExpandedRowKeysExpr sets the expression for the keys of the expanded rows
func (e expandable) ExpandedRowKeysExpr(value string) expandable {
	return e.set("expandedRowKeysExpr", value)
}

// KeyField sets the key field corresponding to the data source
func (e expandable) KeyField(value string) expandable {
	return e.set("keyField", value)
}
