package comp

// rowSelection

type rowSelection schema

// RowSelection 创建一个新的 RowSelection 实例
func RowSelection() rowSelection {
	return rowSelection{}
}

func (r rowSelection) set(key string, value any) rowSelection {
	r[key] = value
	return r
}

// ColumnWidth 已选择的key值表达式
func (r rowSelection) ColumnWidth(value string) rowSelection {
	return r.set("columnWidth", value)
}

// DisableOn 行是否禁用表达式
func (r rowSelection) DisableOn(value string) rowSelection {
	return r.set("disableOn", value)
}

// KeyField 对应数据源的key值
func (r rowSelection) KeyField(value string) rowSelection {
	return r.set("keyField", value)
}

// RowClick 是否点击行触发选中或取消选中
func (r rowSelection) RowClick(value bool) rowSelection {
	return r.set("rowClick", value)
}

// SelectedRowKeys 已选择的key值
func (r rowSelection) SelectedRowKeys(value string) rowSelection {
	return r.set("selectedRowKeys", value)
}

// SelectedRowKeysExpr 已选择的key值表达式
func (r rowSelection) SelectedRowKeysExpr(value string) rowSelection {
	return r.set("selectedRowKeysExpr", value)
}

// Selections 自定义选择菜单
func (r rowSelection) Selections(value string) rowSelection {
	return r.set("selections", value)
}

// Type 选择类型 单选/多选
func (r rowSelection) Type(value string) rowSelection {
	return r.set("type", value)
}
