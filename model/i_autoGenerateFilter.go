package model

// AutoGenerateFilter represents a renderer for automatically generating filter conditions
type AutoGenerateFilter Schema

func NewAutoGenerateFilter() AutoGenerateFilter {
	return make(AutoGenerateFilter)
}

func (a AutoGenerateFilter) set(key string, value any) AutoGenerateFilter {
	a[key] = value
	return a
}

// ColumnsNum sets the number of columns per row in filter conditions
func (a AutoGenerateFilter) ColumnsNum(value string) AutoGenerateFilter {
	return a.set("columnsNum", value)
}

// DefaultCollapsed determines whether the filter is collapsed by default
func (a AutoGenerateFilter) DefaultCollapsed(value bool) AutoGenerateFilter {
	return a.set("defaultCollapsed", value)
}

// ShowBtnToolbar controls the visibility of query field configuration buttons
func (a AutoGenerateFilter) ShowBtnToolbar(value bool) AutoGenerateFilter {
	return a.set("showBtnToolbar", value)
}
