package comp

// RowSelectionOptions
//
// @author  slowlyo
// @version 6.7.0
type RowSelectionOptions Schema

// NewRowSelectionOptions 创建一个新的 RowSelectionOptions 实例
func NewRowSelectionOptions() RowSelectionOptions {
	return RowSelectionOptions{}
}

func (r RowSelectionOptions) set(key string, value interface{}) RowSelectionOptions {
	r[key] = value
	return r
}

// Key 选择类型 选择全部
func (r RowSelectionOptions) Key(value string) RowSelectionOptions {
	return r.set("key", value)
}

// Text 选项显示文本
func (r RowSelectionOptions) Text(value string) RowSelectionOptions {
	return r.set("text", value)
}
