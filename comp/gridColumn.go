package comp

// GridColumn 表示一个网格列
type GridColumn Schema

// NewGridColumn 创建一个新的 GridColumn 实例
func NewGridColumn() GridColumn {
	return make(GridColumn)
}

// set 设置属性并返回当前对象
func (c GridColumn) set(key string, value interface{}) GridColumn {
	c[key] = value
	return c
}

// Body 设置列的内容
func (c GridColumn) Body(value ...interface{}) GridColumn {
	return c.set("body", value)
}

// ColumnClassName 列类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (c GridColumn) ColumnClassName(value string) GridColumn {
	return c.set("columnClassName", value)
}

// Horizontal 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
func (c GridColumn) Horizontal(value string) GridColumn {
	return c.set("horizontal", value)
}

// ID 组件唯一 id
func (c GridColumn) ID(value string) GridColumn {
	return c.set("id", value)
}

// Lg 大屏时(>=1200px)宽度占比
func (c GridColumn) Lg(value string) GridColumn {
	return c.set("lg", value)
}

// Md 中屏时(>=992px)宽度占比
func (c GridColumn) Md(value string) GridColumn {
	return c.set("md", value)
}

// Mode 配置子表单项默认的展示方式。可选值: normal | inline | horizontal
func (c GridColumn) Mode(value string) GridColumn {
	return c.set("mode", value)
}

// Sm 小屏时（>=768px）宽度占比
func (c GridColumn) Sm(value string) GridColumn {
	return c.set("sm", value)
}

// Style 样式
func (c GridColumn) Style(value string) GridColumn {
	return c.set("style", value)
}

// ThemeCss 主题样式
func (c GridColumn) ThemeCss(value string) GridColumn {
	return c.set("themeCss", value)
}

// Valign 垂直对齐方式 可选值: top | middle | bottom | between
func (c GridColumn) Valign(value string) GridColumn {
	return c.set("valign", value)
}

// WrapperCustomStyle 自定义包装样式
func (c GridColumn) WrapperCustomStyle(value string) GridColumn {
	return c.set("wrapperCustomStyle", value)
}

// Xs 极小屏（<768px）时宽度占比
func (c GridColumn) Xs(value string) GridColumn {
	return c.set("xs", value)
}
