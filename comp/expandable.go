package comp

// expandable 表示可展开的渲染器
type expandable schema

// Expandable 创建一个新的 Expandable 实例
func Expandable() expandable {
	return make(expandable)
}

func (e expandable) set(key string, value any) expandable {
	e[key] = value
	return e
}

// ExpandableOn 行是否可展开表达式
func (e expandable) ExpandableOn(value string) expandable {
	return e.set("expandableOn", value)
}

// ExpandedRowClassNameExpr 展开行自定义样式表达式
func (e expandable) ExpandedRowClassNameExpr(value string) expandable {
	return e.set("expandedRowClassNameExpr", value)
}

// ExpandedRowKeys 已展开的 key 值
func (e expandable) ExpandedRowKeys(value string) expandable {
	return e.set("expandedRowKeys", value)
}

// ExpandedRowKeysExpr 已展开的 key 值表达式
func (e expandable) ExpandedRowKeysExpr(value string) expandable {
	return e.set("expandedRowKeysExpr", value)
}

// KeyField 对应数据源的 key 值
func (e expandable) KeyField(value string) expandable {
	return e.set("keyField", value)
}
