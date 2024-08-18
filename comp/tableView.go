package comp

// tableView 表格展现
//
// @version 6.7.0
type tableView schema

// TableView 创建一个新的 TableView 实例
func TableView() tableView {
	return tableView{}.set("type", "table-view")
}

func (tv tableView) set(key string, value interface{}) tableView {
	tv[key] = value
	return tv
}

// TODO
