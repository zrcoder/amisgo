package comp

// GridNav 宫格导航
type GridNav Schema

// NewGridNav 创建一个新的 GridNav 实例
func NewGridNav() GridNav {
	return make(GridNav).set("type", "grid-nav")
}

func (g GridNav) set(key string, value interface{}) GridNav {
	g[key] = value
	return g
}

// Border 是否显示列表项边框
func (g GridNav) Border(value bool) GridNav {
	return g.set("border", value)
}

// Center 是否将列表项内容居中显示
func (g GridNav) Center(value bool) GridNav {
	return g.set("center", value)
}

// ClassName 外层 CSS 类名
func (g GridNav) ClassName(value string) GridNav {
	return g.set("className", value)
}

// ColumnNum 列数
func (g GridNav) ColumnNum(value string) GridNav {
	return g.set("columnNum", value)
}

// Direction 列表项内容排列的方向，可选值为 horizontal 、vertical
func (g GridNav) Direction(value string) GridNav {
	return g.set("direction", value)
}

// Gutter 列表项之间的间距，默认单位为px
func (g GridNav) Gutter(value string) GridNav {
	return g.set("gutter", value)
}

// IconRatio 图标宽度占比，单位%
func (g GridNav) IconRatio(value string) GridNav {
	return g.set("iconRatio", value)
}

// ItemClassName 列表项 css 类名
func (g GridNav) ItemClassName(value string) GridNav {
	return g.set("itemClassName", value)
}

// Options 列表项图标
func (g GridNav) Options(value string) GridNav {
	return g.set("options", value)
}

// Reverse 是否调换图标和文本的位置
func (g GridNav) Reverse(value bool) GridNav {
	return g.set("reverse", value)
}

// Source 数据源
func (g GridNav) Source(value string) GridNav {
	return g.set("source", value)
}

// Square 是否将列表项固定为正方形
func (g GridNav) Square(value bool) GridNav {
	return g.set("square", value)
}

// Value 图片数组
func (g GridNav) Value(value string) GridNav {
	return g.set("value", value)
}
