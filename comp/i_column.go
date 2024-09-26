package comp

// column 表示列渲染器
type column schema

// Column 创建一个新的 Column 实例
func Column() column {
	return make(column)
}

// Set 设置键值对
func (c column) Set(key string, value any) column {
	c[key] = value
	return c
}

// Align 设置内容对齐方式
func (c column) Align(value string) column {
	return c.Set("align", value)
}

// Body 主体内容
func (c column) Body(value ...any) column {
	return c.Set("body", value)
}

// CanAccessSuperData 设置表格列单元格是否可以获取父级数据域值
func (c column) CanAccessSuperData(value bool) column {
	return c.Set("canAccessSuperData", value)
}

// Children 设置表头分组
func (c column) Children(value string) column {
	return c.Set("children", value)
}

// ClassName 设置列样式
func (c column) ClassName(value string) column {
	return c.Set("className", value)
}

// ClassNameExpr 设置单元格样式表达式
func (c column) ClassNameExpr(value string) column {
	return c.Set("classNameExpr", value)
}

// ColSpanExpr 设置指定列合并表达式
func (c column) ColSpanExpr(value string) column {
	return c.Set("colSpanExpr", value)
}

// Copyable 设置是否可复制
func (c column) Copyable(value bool) column {
	return c.Set("copyable", value)
}

// Filterable 设置是否支持筛选
func (c column) Filterable(value bool) column {
	return c.Set("filterable", value)
}

// Fixed 设置是否固定在左侧/右侧
func (c column) Fixed(value string) column {
	return c.Set("fixed", value)
}

// Type
func (c column) Type(value string) column {
	return c.Set("type", value)
}

// Name 设置指定列唯一标识
func (c column) Name(value string) column {
	return c.Set("name", value)
}

// Label 设置标签
func (c column) Label(value string) column {
	return c.Set("label", value)
}

// QuickEdit 配置快速编辑功能
func (c column) QuickEdit(value string) column {
	return c.Set("quickEdit", value)
}

// Remark 设置列表头提示
func (c column) Remark(value string) column {
	return c.Set("remark", value)
}

// RowSpanExpr 设置指定行合并表达式
func (c column) RowSpanExpr(value string) column {
	return c.Set("rowSpanExpr", value)
}

// Searchable 设置是否支持快速搜索
func (c column) Searchable(value bool) column {
	return c.Set("searchable", value)
}

// Sortable 设置是否支持快速排序
func (c column) Sortable(value bool) column {
	return c.Set("sortable", value)
}

// Sorter 设置快速排序功能
func (c column) Sorter(value bool) column {
	return c.Set("sorter", value)
}

// Title 设置指定列标题
func (c column) Title(value any) column {
	return c.Set("title", value)
}

// TitleClassName 设置表头单元格样式
func (c column) TitleClassName(value string) column {
	return c.Set("titleClassName", value)
}

// Togged 设置当前列是否展示
func (c column) Togged(value bool) column {
	return c.Set("toggled", value)
}

// Width 设置列宽
func (c column) Width(value string) column {
	return c.Set("width", value)
}

// ColumnClassName 列类名
func (c column) ColumnClassName(value string) column {
	return c.Set("columnClassName", value)
}

// Xs 宽度占比 1-12 (int) | "auto" (string)
func (c column) Xs(value any) column {
	return c.Set("xs", value)
}

// Sm 宽度占比 1-12 (int) | "auto" (string)
func (c column) Sm(value any) column {
	return c.Set("xs", value)
}

// Md 宽度占比 1-12 (int) | "auto" (string)
func (c column) Md(value any) column {
	return c.Set("xs", value)
}

// Lg 宽度占比 1-12 (int) | "auto" (string)
func (c column) Lg(value any) column {
	return c.Set("xs", value)
}

// Valign 当前列内垂直对齐 'top' | 'middle' | 'bottom' | 'between'
func (c column) Valign(value string) column {
	return c.Set("valign", value)
}

// Buttons
func (c column) Buttons(value ...action) column {
	return c.Set("buttons", value)
}
