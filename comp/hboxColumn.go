package comp

// HBoxColumn 列渲染器
type HBoxColumn Schema

// NewHBoxColumn 创建一个新的 HBoxColumn 实例
func NewHBoxColumn() HBoxColumn {
	return make(HBoxColumn)
}

// set 设置属性
func (h HBoxColumn) set(key string, value interface{}) HBoxColumn {
	h[key] = value
	return h
}

// body 内容区
func (h HBoxColumn) Body(value ...interface{}) HBoxColumn {
	return h.set("body", value)
}

// columnClassName 列上 CSS 类名
func (h HBoxColumn) ColumnClassName(value string) HBoxColumn {
	return h.set("columnClassName", value)
}

// height 高度
func (h HBoxColumn) Height(value string) HBoxColumn {
	return h.set("height", value)
}

// horizontal 水平排版左右宽度占比
func (h HBoxColumn) Horizontal(value string) HBoxColumn {
	return h.set("horizontal", value)
}

// mode 子表单项展示方式
func (h HBoxColumn) Mode(value string) HBoxColumn {
	return h.set("mode", value)
}

// style 其他样式
func (h HBoxColumn) Style(value string) HBoxColumn {
	return h.set("style", value)
}

// valign 垂直对齐方式
func (h HBoxColumn) Valign(value string) HBoxColumn {
	return h.set("valign", value)
}

// visible 是否显示
func (h HBoxColumn) Visible(value bool) HBoxColumn {
	return h.set("visible", value)
}

// visibleOn 是否显示表达式
func (h HBoxColumn) VisibleOn(value string) HBoxColumn {
	return h.set("visibleOn", value)
}

// width 宽度
func (h HBoxColumn) Width(value string) HBoxColumn {
	return h.set("width", value)
}
