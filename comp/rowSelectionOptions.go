package comp

import "github.com/zrcoder/amisgo/model"

// rowSelectionOptions

type rowSelectionOptions model.Schema

// RowSelectionOptions creates a new RowSelectionOptions instance
func RowSelectionOptions() rowSelectionOptions {
	return rowSelectionOptions{}
}

func (r rowSelectionOptions) set(key string, value any) rowSelectionOptions {
	r[key] = value
	return r
}

// Key sets the selection key
func (r rowSelectionOptions) Key(value string) rowSelectionOptions {
	return r.set("key", value)
}

// Text sets the option display text
func (r rowSelectionOptions) Text(value string) rowSelectionOptions {
	return r.set("text", value)
}
