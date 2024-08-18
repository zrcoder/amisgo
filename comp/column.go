package comp

// column 表示列渲染器
type column schema

// Column 创建一个新的 Column 实例
func Column() column {
	return make(column)
}

func (c column) set(key string, value interface{}) column {
	c[key] = value
	return c
}

// Align 设置内容对齐方式
func (c column) Align(value string) column {
	return c.set("align", value)
}

// CanAccessSuperData 设置表格列单元格是否可以获取父级数据域值
func (c column) CanAccessSuperData(value bool) column {
	return c.set("canAccessSuperData", value)
}

// Children 设置表头分组
func (c column) Children(value string) column {
	return c.set("children", value)
}

// ClassName 设置列样式
func (c column) ClassName(value string) column {
	return c.set("className", value)
}

// ClassNameExpr 设置单元格样式表达式
func (c column) ClassNameExpr(value string) column {
	return c.set("classNameExpr", value)
}

// ColSpanExpr 设置指定列合并表达式
func (c column) ColSpanExpr(value string) column {
	return c.set("colSpanExpr", value)
}

// Copyable 设置是否可复制
func (c column) Copyable(value bool) column {
	return c.set("copyable", value)
}

// Filterable 设置是否支持筛选
func (c column) Filterable(value bool) column {
	return c.set("filterable", value)
}

// Fixed 设置是否固定在左侧/右侧
func (c column) Fixed(value string) column {
	return c.set("fixed", value)
}

// Name 设置指定列唯一标识
func (c column) Name(value string) column {
	return c.set("name", value)
}

// QuickEdit 配置快速编辑功能
func (c column) QuickEdit(value string) column {
	return c.set("quickEdit", value)
}

// Remark 设置列表头提示
func (c column) Remark(value string) column {
	return c.set("remark", value)
}

// RowSpanExpr 设置指定行合并表达式
func (c column) RowSpanExpr(value string) column {
	return c.set("rowSpanExpr", value)
}

// Searchable 设置是否支持快速搜索
func (c column) Searchable(value bool) column {
	return c.set("searchable", value)
}

// Sortable 设置是否支持快速排序
func (c column) Sortable(value bool) column {
	return c.set("sortable", value)
}

// Sorter 设置快速排序功能
func (c column) Sorter(value bool) column {
	return c.set("sorter", value)
}

// Title 设置指定列标题
func (c column) Title(value string) column {
	return c.set("title", value)
}

// TitleClassName 设置表头单元格样式
func (c column) TitleClassName(value string) column {
	return c.set("titleClassName", value)
}

// Togged 设置当前列是否展示
func (c column) Togged(value bool) column {
	return c.set("toggled", value)
}

// Width 设置列宽
func (c column) Width(value string) column {
	return c.set("width", value)
}
