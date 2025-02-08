package comp

import "github.com/zrcoder/amisgo/model"

type TableView model.Schema

func NewTableView() TableView {
	return TableView{"type": "table-view"}
}

// set assigns a value to a key in the tableView schema
func (tv TableView) set(key string, value any) TableView {
	tv[key] = value
	return tv
}

// TODO
