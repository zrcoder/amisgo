package model

// RowSelection represents the configuration for row selection in a table.
type RowSelection Schema

func NewRowSelection() RowSelection {
	return RowSelection{}
}

// set is a helper method to set a key-value pair in the rowSelection.
func (r RowSelection) set(key string, value any) RowSelection {
	r[key] = value
	return r
}

// ColumnWidth sets the column width for the row selection.
func (r RowSelection) ColumnWidth(value string) RowSelection {
	return r.set("columnWidth", value)
}

// DisableOn sets the condition expression for disabling rows.
func (r RowSelection) DisableOn(value string) RowSelection {
	return r.set("disableOn", value)
}

// KeyField sets the key field that corresponds to the data source.
func (r RowSelection) KeyField(value string) RowSelection {
	return r.set("keyField", value)
}

// RowClick sets whether to trigger selection or deselection when clicking on a row.
func (r RowSelection) RowClick(value bool) RowSelection {
	return r.set("rowClick", value)
}

// SelectedRowKeys sets the selected row keys.
func (r RowSelection) SelectedRowKeys(value string) RowSelection {
	return r.set("selectedRowKeys", value)
}

// SelectedRowKeysExpr sets the expression for selected row keys.
func (r RowSelection) SelectedRowKeysExpr(value string) RowSelection {
	return r.set("selectedRowKeysExpr", value)
}

// Selections sets the custom selection menu.
func (r RowSelection) Selections(value string) RowSelection {
	return r.set("selections", value)
}

// Type sets the selection type, either single or multiple.
func (r RowSelection) Type(value string) RowSelection {
	return r.set("type", value)
}
