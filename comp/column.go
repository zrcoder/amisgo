package comp

// Column 表示列渲染器
type Column Schema

// NewColumn 创建一个新的 Column 实例
func NewColumn() Column {
	return make(Column)
}

func (c Column) set(key string, value interface{}) Column {
	c[key] = value
	return c
}

// Align 设置内容对齐方式
func (c Column) Align(value string) Column {
	return c.set("align", value)
}

// CanAccessSuperData 设置表格列单元格是否可以获取父级数据域值
func (c Column) CanAccessSuperData(value bool) Column {
	return c.set("canAccessSuperData", value)
}

// Children 设置表头分组
func (c Column) Children(value string) Column {
	return c.set("children", value)
}

// ClassName 设置列样式
func (c Column) ClassName(value string) Column {
	return c.set("className", value)
}

// ClassNameExpr 设置单元格样式表达式
func (c Column) ClassNameExpr(value string) Column {
	return c.set("classNameExpr", value)
}

// ColSpanExpr 设置指定列合并表达式
func (c Column) ColSpanExpr(value string) Column {
	return c.set("colSpanExpr", value)
}

// Copyable 设置是否可复制
func (c Column) Copyable(value bool) Column {
	return c.set("copyable", value)
}

// Filterable 设置是否支持筛选
func (c Column) Filterable(value bool) Column {
	return c.set("filterable", value)
}

// Fixed 设置是否固定在左侧/右侧
func (c Column) Fixed(value string) Column {
	return c.set("fixed", value)
}

// Name 设置指定列唯一标识
func (c Column) Name(value string) Column {
	return c.set("name", value)
}

// QuickEdit 配置快速编辑功能
func (c Column) QuickEdit(value string) Column {
	return c.set("quickEdit", value)
}

// Remark 设置列表头提示
func (c Column) Remark(value string) Column {
	return c.set("remark", value)
}

// RowSpanExpr 设置指定行合并表达式
func (c Column) RowSpanExpr(value string) Column {
	return c.set("rowSpanExpr", value)
}

// Searchable 设置是否支持快速搜索
func (c Column) Searchable(value bool) Column {
	return c.set("searchable", value)
}

// Sortable 设置是否支持快速排序
func (c Column) Sortable(value bool) Column {
	return c.set("sortable", value)
}

// Sorter 设置快速排序功能
func (c Column) Sorter(value bool) Column {
	return c.set("sorter", value)
}

// Title 设置指定列标题
func (c Column) Title(value string) Column {
	return c.set("title", value)
}

// TitleClassName 设置表头单元格样式
func (c Column) TitleClassName(value string) Column {
	return c.set("titleClassName", value)
}

// Togged 设置当前列是否展示
func (c Column) Togged(value bool) Column {
	return c.set("toggled", value)
}

// Width 设置列宽
func (c Column) Width(value string) Column {
	return c.set("width", value)
}
