package comp

// Expandable 表示可展开的渲染器
type Expandable Schema

// NewExpandable 创建一个新的 Expandable 实例
func NewExpandable() Expandable {
	return make(Expandable)
}

func (e Expandable) set(key string, value interface{}) Expandable {
	e[key] = value
	return e
}

// ExpandableOn 行是否可展开表达式
func (e Expandable) ExpandableOn(value string) Expandable {
	return e.set("expandableOn", value)
}

// ExpandedRowClassNameExpr 展开行自定义样式表达式
func (e Expandable) ExpandedRowClassNameExpr(value string) Expandable {
	return e.set("expandedRowClassNameExpr", value)
}

// ExpandedRowKeys 已展开的 key 值
func (e Expandable) ExpandedRowKeys(value string) Expandable {
	return e.set("expandedRowKeys", value)
}

// ExpandedRowKeysExpr 已展开的 key 值表达式
func (e Expandable) ExpandedRowKeysExpr(value string) Expandable {
	return e.set("expandedRowKeysExpr", value)
}

// KeyField 对应数据源的 key 值
func (e Expandable) KeyField(value string) Expandable {
	return e.set("keyField", value)
}

// Type 对应渲染器类型
func (e Expandable) Type(value string) Expandable {
	return e.set("type", value)
}
