package comp

// tableColumn 表格列，不指定类型时默认为文本类型。

type tableColumn schema

// TableColumn 创建一个新的 TableColumn 实例
func TableColumn() tableColumn {
	return tableColumn{}
}

func (c tableColumn) set(key string, value any) tableColumn {
	c[key] = value
	return c
}

// Align 列对齐方式
func (c tableColumn) Align(value string) tableColumn {
	return c.set("align", value)
}

// Breakpoint 结合表格的 footable 一起使用
func (c tableColumn) Breakpoint(value string) tableColumn {
	return c.set("breakpoint", value)
}

// CanAccessSuperData 表格列单元格是否可以获取父级数据域值
func (c tableColumn) CanAccessSuperData(value bool) tableColumn {
	return c.set("canAccessSuperData", value)
}

// ClassName 列样式表
func (c tableColumn) ClassName(value string) tableColumn {
	return c.set("className", value)
}

// ClassNameExpr 单元格样式表达式
func (c tableColumn) ClassNameExpr(value string) tableColumn {
	return c.set("classNameExpr", value)
}

// Copyable 配置点击复制功能
func (c tableColumn) Copyable(value bool) tableColumn {
	return c.set("copyable", value)
}

// Filterable todo
func (c tableColumn) Filterable(value bool) tableColumn {
	return c.set("filterable", value)
}

// Fixed 配置是否固定当前列
func (c tableColumn) Fixed(value string) tableColumn {
	return c.set("fixed", value)
}

// HeaderAlign 标题左右对齐方式
func (c tableColumn) HeaderAlign(value string) tableColumn {
	return c.set("headerAlign", value)
}

// InnerStyle 单元格内部组件自定义样式
func (c tableColumn) InnerStyle(value any) tableColumn {
	return c.set("innerStyle", value)
}

// Label 列标题
func (c tableColumn) Label(value string) tableColumn {
	return c.set("label", value)
}

// LabelClassName 列头样式表
func (c tableColumn) LabelClassName(value string) tableColumn {
	return c.set("labelClassName", value)
}

// LazyRenderAfter 当一次性渲染太多列上有用
func (c tableColumn) LazyRenderAfter(value string) tableColumn {
	return c.set("lazyRenderAfter", value)
}

// Name 绑定字段名
func (c tableColumn) Name(value string) tableColumn {
	return c.set("name", value)
}

// PopOver 配置查看详情功能
func (c tableColumn) PopOver(value string) tableColumn {
	return c.set("popOver", value)
}

// QuickEdit 配置快速编辑功能
func (c tableColumn) QuickEdit(value string) tableColumn {
	return c.set("quickEdit", value)
}

// QuickEditOnUpdate 作为表单项时，可以单独配置编辑时的快速编辑面板
func (c tableColumn) QuickEditOnUpdate(value string) tableColumn {
	return c.set("quickEditOnUpdate", value)
}

// Remark 提示信息
func (c tableColumn) Remark(value string) tableColumn {
	return c.set("remark", value)
}

// Searchable 是否可快速搜索
func (c tableColumn) Searchable(value bool) tableColumn {
	return c.set("searchable", value)
}

// Sortable 配置是否可以排序
func (c tableColumn) Sortable(value bool) tableColumn {
	return c.set("sortable", value)
}

// Toggled 配置是否默认展示
func (c tableColumn) Toggled(value bool) tableColumn {
	return c.set("toggled", value)
}

// Unique 是否唯一
func (c tableColumn) Unique(value bool) tableColumn {
	return c.set("unique", value)
}

// VAlign 列垂直对齐方式
func (c tableColumn) VAlign(value string) tableColumn {
	return c.set("vAlign", value)
}

// Value 默认值
func (c tableColumn) Value(value string) tableColumn {
	return c.set("value", value)
}

// Width 列宽度
func (c tableColumn) Width(value string) tableColumn {
	return c.set("width", value)
}
