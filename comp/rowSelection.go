package comp

// RowSelection
//
// @author  slowlyo
// @version 6.7.0
type RowSelection Schema

// NewRowSelection 创建一个新的 RowSelection 实例
func NewRowSelection() RowSelection {
	return RowSelection{}
}

func (r RowSelection) set(key string, value interface{}) RowSelection {
	r[key] = value
	return r
}

// ColumnWidth 已选择的key值表达式
func (r RowSelection) ColumnWidth(value string) RowSelection {
	return r.set("columnWidth", value)
}

// DisableOn 行是否禁用表达式
func (r RowSelection) DisableOn(value string) RowSelection {
	return r.set("disableOn", value)
}

// KeyField 对应数据源的key值
func (r RowSelection) KeyField(value string) RowSelection {
	return r.set("keyField", value)
}

// RowClick 是否点击行触发选中或取消选中
func (r RowSelection) RowClick(value bool) RowSelection {
	return r.set("rowClick", value)
}

// SelectedRowKeys 已选择的key值
func (r RowSelection) SelectedRowKeys(value string) RowSelection {
	return r.set("selectedRowKeys", value)
}

// SelectedRowKeysExpr 已选择的key值表达式
func (r RowSelection) SelectedRowKeysExpr(value string) RowSelection {
	return r.set("selectedRowKeysExpr", value)
}

// Selections 自定义选择菜单
func (r RowSelection) Selections(value string) RowSelection {
	return r.set("selections", value)
}

// Type 选择类型 单选/多选
func (r RowSelection) Type(value string) RowSelection {
	return r.set("type", value)
}
