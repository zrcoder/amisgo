package comp

import "github.com/zrcoder/amisgo/model"

// AutoGenerateFilter represents a renderer for automatically generating filter conditions
type autoGenerateFilter model.Schema

// AutoGenerateFilter creates a new AutoGenerateFilter instance
func AutoGenerateFilter() autoGenerateFilter {
	return make(autoGenerateFilter)
}

func (a autoGenerateFilter) set(key string, value any) autoGenerateFilter {
	a[key] = value
	return a
}

// ColumnsNum sets the number of columns per row in filter conditions
func (a autoGenerateFilter) ColumnsNum(value string) autoGenerateFilter {
	return a.set("columnsNum", value)
}

// DefaultCollapsed determines whether the filter is collapsed by default
func (a autoGenerateFilter) DefaultCollapsed(value bool) autoGenerateFilter {
	return a.set("defaultCollapsed", value)
}

// ShowBtnToolbar controls the visibility of query field configuration buttons
func (a autoGenerateFilter) ShowBtnToolbar(value bool) autoGenerateFilter {
	return a.set("showBtnToolbar", value)
}
