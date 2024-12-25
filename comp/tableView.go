package comp

// tableView 表格展现

type tableView Schema

// TableView 创建一个新的 TableView 实例
func TableView() tableView {
	return tableView{}.set("type", "table-view")
}

func (tv tableView) set(key string, value any) tableView {
	tv[key] = value
	return tv
}

// TODO
