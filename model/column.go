package model

// ColumnSchema 表示列渲染器
type ColumnSchema Schema

// Column 创建一个新的 Column 实例
func Column() ColumnSchema {
	return make(ColumnSchema)
}

// Set 设置键值对
func (c ColumnSchema) Set(key string, value any) ColumnSchema {
	c[key] = value
	return c
}

// Align 设置内容对齐方式
func (c ColumnSchema) Align(value string) ColumnSchema {
	return c.Set("align", value)
}

// CanAccessSuperData 设置表格列单元格是否可以获取父级数据域值
func (c ColumnSchema) CanAccessSuperData(value bool) ColumnSchema {
	return c.Set("canAccessSuperData", value)
}

// Children 设置表头分组
func (c ColumnSchema) Children(value string) ColumnSchema {
	return c.Set("children", value)
}

// ClassName 设置列样式
func (c ColumnSchema) ClassName(value string) ColumnSchema {
	return c.Set("className", value)
}

// ClassNameExpr 设置单元格样式表达式
func (c ColumnSchema) ClassNameExpr(value string) ColumnSchema {
	return c.Set("classNameExpr", value)
}

// ColSpanExpr 设置指定列合并表达式
func (c ColumnSchema) ColSpanExpr(value string) ColumnSchema {
	return c.Set("colSpanExpr", value)
}

// Copyable 设置是否可复制
func (c ColumnSchema) Copyable(value bool) ColumnSchema {
	return c.Set("copyable", value)
}

// Filterable 设置是否支持筛选
func (c ColumnSchema) Filterable(value bool) ColumnSchema {
	return c.Set("filterable", value)
}

// Fixed 设置是否固定在左侧/右侧
func (c ColumnSchema) Fixed(value string) ColumnSchema {
	return c.Set("fixed", value)
}

// Name 设置指定列唯一标识
func (c ColumnSchema) Name(value string) ColumnSchema {
	return c.Set("name", value)
}

// Label 设置标签
func (c ColumnSchema) Label(value string) ColumnSchema {
	return c.Set("label", value)
}

// QuickEdit 配置快速编辑功能
func (c ColumnSchema) QuickEdit(value string) ColumnSchema {
	return c.Set("quickEdit", value)
}

// Remark 设置列表头提示
func (c ColumnSchema) Remark(value string) ColumnSchema {
	return c.Set("remark", value)
}

// RowSpanExpr 设置指定行合并表达式
func (c ColumnSchema) RowSpanExpr(value string) ColumnSchema {
	return c.Set("rowSpanExpr", value)
}

// Searchable 设置是否支持快速搜索
func (c ColumnSchema) Searchable(value bool) ColumnSchema {
	return c.Set("searchable", value)
}

// Sortable 设置是否支持快速排序
func (c ColumnSchema) Sortable(value bool) ColumnSchema {
	return c.Set("sortable", value)
}

// Sorter 设置快速排序功能
func (c ColumnSchema) Sorter(value bool) ColumnSchema {
	return c.Set("sorter", value)
}

// Title 设置指定列标题
func (c ColumnSchema) Title(value any) ColumnSchema {
	return c.Set("title", value)
}

// TitleClassName 设置表头单元格样式
func (c ColumnSchema) TitleClassName(value string) ColumnSchema {
	return c.Set("titleClassName", value)
}

// Togged 设置当前列是否展示
func (c ColumnSchema) Togged(value bool) ColumnSchema {
	return c.Set("toggled", value)
}

// Width 设置列宽
func (c ColumnSchema) Width(value string) ColumnSchema {
	return c.Set("width", value)
}

// ColumnClassName 列类名
func (c ColumnSchema) ColumnClassName(value string) ColumnSchema {
	return c.Set("columnClassName", value)
}

// Xs 宽度占比 1-12 (int) | "auto" (string)
func (c ColumnSchema) Xs(value any) ColumnSchema {
	return c.Set("xs", value)
}

// Sm 宽度占比 1-12 (int) | "auto" (string)
func (c ColumnSchema) Sm(value any) ColumnSchema {
	return c.Set("xs", value)
}

// Md 宽度占比 1-12 (int) | "auto" (string)
func (c ColumnSchema) Md(value any) ColumnSchema {
	return c.Set("xs", value)
}

// Lg 宽度占比 1-12 (int) | "auto" (string)
func (c ColumnSchema) Lg(value any) ColumnSchema {
	return c.Set("xs", value)
}

// Valign 当前列内垂直对齐 'top' | 'middle' | 'bottom' | 'between'
func (c ColumnSchema) Valign(value string) ColumnSchema {
	return c.Set("valign", value)
}
