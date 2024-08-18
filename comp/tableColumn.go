package comp

// TableColumn 表格列，不指定类型时默认为文本类型。
//
// @version 6.7.0
type TableColumn Schema

// NewTableColumn 创建一个新的 TableColumn 实例
func NewTableColumn() TableColumn {
	return TableColumn{}
}

func (c TableColumn) set(key string, value interface{}) TableColumn {
	c[key] = value
	return c
}

// Align 列对齐方式
func (c TableColumn) Align(value string) TableColumn {
	return c.set("align", value)
}

// Breakpoint 结合表格的 footable 一起使用
func (c TableColumn) Breakpoint(value string) TableColumn {
	return c.set("breakpoint", value)
}

// CanAccessSuperData 表格列单元格是否可以获取父级数据域值
func (c TableColumn) CanAccessSuperData(value bool) TableColumn {
	return c.set("canAccessSuperData", value)
}

// ClassName 列样式表
func (c TableColumn) ClassName(value string) TableColumn {
	return c.set("className", value)
}

// ClassNameExpr 单元格样式表达式
func (c TableColumn) ClassNameExpr(value string) TableColumn {
	return c.set("classNameExpr", value)
}

// Copyable 配置点击复制功能
func (c TableColumn) Copyable(value bool) TableColumn {
	return c.set("copyable", value)
}

// Filterable todo
func (c TableColumn) Filterable(value bool) TableColumn {
	return c.set("filterable", value)
}

// Fixed 配置是否固定当前列
func (c TableColumn) Fixed(value string) TableColumn {
	return c.set("fixed", value)
}

// HeaderAlign 标题左右对齐方式
func (c TableColumn) HeaderAlign(value string) TableColumn {
	return c.set("headerAlign", value)
}

// InnerStyle 单元格内部组件自定义样式
func (c TableColumn) InnerStyle(value string) TableColumn {
	return c.set("innerStyle", value)
}

// Label 列标题
func (c TableColumn) Label(value string) TableColumn {
	return c.set("label", value)
}

// LabelClassName 列头样式表
func (c TableColumn) LabelClassName(value string) TableColumn {
	return c.set("labelClassName", value)
}

// LazyRenderAfter 当一次性渲染太多列上有用
func (c TableColumn) LazyRenderAfter(value string) TableColumn {
	return c.set("lazyRenderAfter", value)
}

// Name 绑定字段名
func (c TableColumn) Name(value string) TableColumn {
	return c.set("name", value)
}

// PopOver 配置查看详情功能
func (c TableColumn) PopOver(value string) TableColumn {
	return c.set("popOver", value)
}

// QuickEdit 配置快速编辑功能
func (c TableColumn) QuickEdit(value string) TableColumn {
	return c.set("quickEdit", value)
}

// QuickEditOnUpdate 作为表单项时，可以单独配置编辑时的快速编辑面板
func (c TableColumn) QuickEditOnUpdate(value string) TableColumn {
	return c.set("quickEditOnUpdate", value)
}

// Remark 提示信息
func (c TableColumn) Remark(value string) TableColumn {
	return c.set("remark", value)
}

// Searchable 是否可快速搜索
func (c TableColumn) Searchable(value bool) TableColumn {
	return c.set("searchable", value)
}

// Sortable 配置是否可以排序
func (c TableColumn) Sortable(value bool) TableColumn {
	return c.set("sortable", value)
}

// Toggled 配置是否默认展示
func (c TableColumn) Toggled(value bool) TableColumn {
	return c.set("toggled", value)
}

// Unique 是否唯一
func (c TableColumn) Unique(value bool) TableColumn {
	return c.set("unique", value)
}

// VAlign 列垂直对齐方式
func (c TableColumn) VAlign(value string) TableColumn {
	return c.set("vAlign", value)
}

// Value 默认值
func (c TableColumn) Value(value string) TableColumn {
	return c.set("value", value)
}

// Width 列宽度
func (c TableColumn) Width(value string) TableColumn {
	return c.set("width", value)
}
