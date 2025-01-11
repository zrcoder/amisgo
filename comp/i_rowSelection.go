package comp

import "github.com/zrcoder/amisgo/model"

// rowSelection represents the configuration for row selection in a table.
type rowSelection model.Schema

// RowSelection creates and returns a new rowSelection instance.
func RowSelection() rowSelection {
	return rowSelection{}
}

// set is a helper method to set a key-value pair in the rowSelection.
func (r rowSelection) set(key string, value any) rowSelection {
	r[key] = value
	return r
}

// ColumnWidth sets the column width for the row selection.
func (r rowSelection) ColumnWidth(value string) rowSelection {
	return r.set("columnWidth", value)
}

// DisableOn sets the condition expression for disabling rows.
func (r rowSelection) DisableOn(value string) rowSelection {
	return r.set("disableOn", value)
}

// KeyField sets the key field that corresponds to the data source.
func (r rowSelection) KeyField(value string) rowSelection {
	return r.set("keyField", value)
}

// RowClick sets whether to trigger selection or deselection when clicking on a row.
func (r rowSelection) RowClick(value bool) rowSelection {
	return r.set("rowClick", value)
}

// SelectedRowKeys sets the selected row keys.
func (r rowSelection) SelectedRowKeys(value string) rowSelection {
	return r.set("selectedRowKeys", value)
}

// SelectedRowKeysExpr sets the expression for selected row keys.
func (r rowSelection) SelectedRowKeysExpr(value string) rowSelection {
	return r.set("selectedRowKeysExpr", value)
}

// Selections sets the custom selection menu.
func (r rowSelection) Selections(value string) rowSelection {
	return r.set("selections", value)
}

// Type sets the selection type, either single or multiple.
func (r rowSelection) Type(value string) rowSelection {
	return r.set("type", value)
}
