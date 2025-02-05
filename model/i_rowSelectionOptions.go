package model

// RowSelectionOptions
type RowSelectionOptions Schema

func NewRowSelectionOptions() RowSelectionOptions {
	return RowSelectionOptions{}
}

func (r RowSelectionOptions) set(key string, value any) RowSelectionOptions {
	r[key] = value
	return r
}

// Key sets the selection key
func (r RowSelectionOptions) Key(value string) RowSelectionOptions {
	return r.set("key", value)
}

// Text sets the option display text
func (r RowSelectionOptions) Text(value string) RowSelectionOptions {
	return r.set("text", value)
}
