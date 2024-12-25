package comp

// gridColumn 表示一个网格列
type gridColumn Schema

// GridColumn 创建一个新的 GridColumn 实例
func GridColumn() gridColumn {
	return make(gridColumn)
}

func (c gridColumn) set(key string, value any) gridColumn {
	c[key] = value
	return c
}

// Body 设置列的内容
func (c gridColumn) Body(value ...any) gridColumn {
	return c.set("body", value)
}

// ColumnClassName 列类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (c gridColumn) ColumnClassName(value string) gridColumn {
	return c.set("columnClassName", value)
}

// Horizontal 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
func (c gridColumn) Horizontal(value string) gridColumn {
	return c.set("horizontal", value)
}

// ID 组件唯一 id
func (c gridColumn) ID(value string) gridColumn {
	return c.set("id", value)
}

// Lg 大屏时(>=1200px)宽度占比
func (c gridColumn) Lg(value string) gridColumn {
	return c.set("lg", value)
}

// Md 中屏时(>=992px)宽度占比
func (c gridColumn) Md(value string) gridColumn {
	return c.set("md", value)
}

// Mode 配置子表单项默认的展示方式。可选值: normal | inline | horizontal
func (c gridColumn) Mode(value string) gridColumn {
	return c.set("mode", value)
}

// Sm 小屏时（>=768px）宽度占比
func (c gridColumn) Sm(value string) gridColumn {
	return c.set("sm", value)
}

// Style 样式
func (c gridColumn) Style(value any) gridColumn {
	return c.set("style", value)
}

// ThemeCss 主题样式
func (c gridColumn) ThemeCSS(value any) gridColumn {
	return c.set("themeCss", value)
}

// Valign 垂直对齐方式 可选值: top | middle | bottom | between
func (c gridColumn) Valign(value string) gridColumn {
	return c.set("valign", value)
}

// WrapperCustomStyle 自定义包装样式
func (c gridColumn) WrapperCustomStyle(value any) gridColumn {
	return c.set("wrapperCustomStyle", value)
}

// Xs 极小屏（<768px）时宽度占比
func (c gridColumn) Xs(value string) gridColumn {
	return c.set("xs", value)
}
