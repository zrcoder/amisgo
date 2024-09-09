package comp

// rowSelectionOptions

type rowSelectionOptions schema

// RowSelectionOptions 创建一个新的 RowSelectionOptions 实例
func RowSelectionOptions() rowSelectionOptions {
	return rowSelectionOptions{}
}

func (r rowSelectionOptions) set(key string, value any) rowSelectionOptions {
	r[key] = value
	return r
}

// Key 选择类型 选择全部
func (r rowSelectionOptions) Key(value string) rowSelectionOptions {
	return r.set("key", value)
}

// Text 选项显示文本
func (r rowSelectionOptions) Text(value string) rowSelectionOptions {
	return r.set("text", value)
}
