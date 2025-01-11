package comp

import "github.com/zrcoder/amisgo/model"

type tableView model.Schema

// TableView creates a new TableView instance
func TableView() tableView {
	return tableView{"type": "table-view"}
}

// set assigns a value to a key in the tableView schema
func (tv tableView) set(key string, value any) tableView {
	tv[key] = value
	return tv
}

// TODO
