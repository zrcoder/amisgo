package comp

type MColumn Schema

func Column() MColumn {
	return make(MColumn)
}

// Set 设置键值对
func (c MColumn) Set(key string, value any) MColumn {
	c[key] = value
	return c
}

// Align 设置内容对齐方式
func (c MColumn) Align(value string) MColumn {
	return c.Set("align", value)
}

// Body 主体内容
func (c MColumn) Body(value ...any) MColumn {
	return c.Set("body", value)
}

// CanAccessSuperData 设置表格列单元格是否可以获取父级数据域值
func (c MColumn) CanAccessSuperData(value bool) MColumn {
	return c.Set("canAccessSuperData", value)
}

// Children 设置表头分组
func (c MColumn) Children(value string) MColumn {
	return c.Set("children", value)
}

// ClassName 设置列样式
func (c MColumn) ClassName(value string) MColumn {
	return c.Set("className", value)
}

// ClassNameExpr 设置单元格样式表达式
func (c MColumn) ClassNameExpr(value string) MColumn {
	return c.Set("classNameExpr", value)
}

// ColSpanExpr 设置指定列合并表达式
func (c MColumn) ColSpanExpr(value string) MColumn {
	return c.Set("colSpanExpr", value)
}

// Copyable 设置是否可复制
func (c MColumn) Copyable(value bool) MColumn {
	return c.Set("copyable", value)
}

// Filterable 设置是否支持筛选
func (c MColumn) Filterable(value bool) MColumn {
	return c.Set("filterable", value)
}

// Fixed 设置是否固定在左侧/右侧
func (c MColumn) Fixed(value string) MColumn {
	return c.Set("fixed", value)
}

// Type
func (c MColumn) Type(value string) MColumn {
	return c.Set("type", value)
}

// Name 设置指定列唯一标识
func (c MColumn) Name(value string) MColumn {
	return c.Set("name", value)
}

// Label 设置标签
func (c MColumn) Label(value string) MColumn {
	return c.Set("label", value)
}

// QuickEdit 配置快速编辑功能
func (c MColumn) QuickEdit(value string) MColumn {
	return c.Set("quickEdit", value)
}

// Remark 设置列表头提示
func (c MColumn) Remark(value string) MColumn {
	return c.Set("remark", value)
}

// RowSpanExpr 设置指定行合并表达式
func (c MColumn) RowSpanExpr(value string) MColumn {
	return c.Set("rowSpanExpr", value)
}

// Searchable 设置是否支持快速搜索(bool), 或自定义搜索组件如 switch
func (c MColumn) Searchable(value any) MColumn {
	return c.Set("searchable", value)
}

// Sortable 设置是否支持快速排序
func (c MColumn) Sortable(value bool) MColumn {
	return c.Set("sortable", value)
}

// Sorter 设置快速排序功能
func (c MColumn) Sorter(value bool) MColumn {
	return c.Set("sorter", value)
}

// Title 设置指定列标题
func (c MColumn) Title(value any) MColumn {
	return c.Set("title", value)
}

// TitleClassName 设置表头单元格样式
func (c MColumn) TitleClassName(value string) MColumn {
	return c.Set("titleClassName", value)
}

// Togged 设置当前列是否展示
func (c MColumn) Togged(value bool) MColumn {
	return c.Set("toggled", value)
}

// Width 设置列宽
func (c MColumn) Width(value string) MColumn {
	return c.Set("width", value)
}

// ColumnClassName 列类名
func (c MColumn) ColumnClassName(value string) MColumn {
	return c.Set("columnClassName", value)
}

// Xs 宽度占比 1-12 (int) | "auto" (string)
func (c MColumn) Xs(value any) MColumn {
	return c.Set("xs", value)
}

// Sm 宽度占比 1-12 (int) | "auto" (string)
func (c MColumn) Sm(value any) MColumn {
	return c.Set("xs", value)
}

// Md 宽度占比 1-12 (int) | "auto" (string)
func (c MColumn) Md(value any) MColumn {
	return c.Set("xs", value)
}

// Lg 宽度占比 1-12 (int) | "auto" (string)
func (c MColumn) Lg(value any) MColumn {
	return c.Set("xs", value)
}

// Valign 当前列内垂直对齐 'top' | 'middle' | 'bottom' | 'between'
func (c MColumn) Valign(value string) MColumn {
	return c.Set("valign", value)
}

// Buttons
func (c MColumn) Buttons(value ...action) MColumn {
	return c.Set("buttons", value)
}
