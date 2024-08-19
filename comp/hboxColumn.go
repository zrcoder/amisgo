package comp

// hBoxColumn 列渲染器
type hBoxColumn schema

// HBoxColumn 创建一个新的 HBoxColumn 实例
func HBoxColumn() hBoxColumn {
	return make(hBoxColumn)
}

func (h hBoxColumn) set(key string, value any) hBoxColumn {
	h[key] = value
	return h
}

// body 内容区
func (h hBoxColumn) Body(value ...any) hBoxColumn {
	return h.set("body", value)
}

// columnClassName 列上 CSS 类名
func (h hBoxColumn) ColumnClassName(value string) hBoxColumn {
	return h.set("columnClassName", value)
}

// height 高度
func (h hBoxColumn) Height(value string) hBoxColumn {
	return h.set("height", value)
}

// horizontal 水平排版左右宽度占比
func (h hBoxColumn) Horizontal(value string) hBoxColumn {
	return h.set("horizontal", value)
}

// mode 子表单项展示方式
func (h hBoxColumn) Mode(value string) hBoxColumn {
	return h.set("mode", value)
}

// style 其他样式
func (h hBoxColumn) Style(value any) hBoxColumn {
	return h.set("style", value)
}

// valign 垂直对齐方式
func (h hBoxColumn) Valign(value string) hBoxColumn {
	return h.set("valign", value)
}

// visible 是否显示
func (h hBoxColumn) Visible(value bool) hBoxColumn {
	return h.set("visible", value)
}

// visibleOn 是否显示表达式
func (h hBoxColumn) VisibleOn(value string) hBoxColumn {
	return h.set("visibleOn", value)
}

// width 宽度
func (h hBoxColumn) Width(value string) hBoxColumn {
	return h.set("width", value)
}
