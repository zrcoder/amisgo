package comp

// TableView 表格展现
//
// @version 6.7.0
type TableView Schema

// NewTableView 创建一个新的 TableView 实例
func NewTableView() TableView {
	return TableView{}.set("type", "table-view")
}

func (tv TableView) set(key string, value interface{}) TableView {
	tv[key] = value
	return tv
}

// TODO
