package comp

// gridNav 宫格导航
type gridNav Schema

// GridNav 创建一个新的 GridNav 实例
func GridNav() gridNav {
	return make(gridNav).set("type", "grid-nav")
}

func (g gridNav) set(key string, value any) gridNav {
	g[key] = value
	return g
}

// Border 是否显示列表项边框
func (g gridNav) Border(value bool) gridNav {
	return g.set("border", value)
}

// Center 是否将列表项内容居中显示
func (g gridNav) Center(value bool) gridNav {
	return g.set("center", value)
}

// ClassName 外层 CSS 类名
func (g gridNav) ClassName(value string) gridNav {
	return g.set("className", value)
}

// ColumnNum 列数
func (g gridNav) ColumnNum(value string) gridNav {
	return g.set("columnNum", value)
}

// Direction 列表项内容排列的方向，可选值为 horizontal 、vertical
func (g gridNav) Direction(value string) gridNav {
	return g.set("direction", value)
}

// Gutter 列表项之间的间距，默认单位为px
func (g gridNav) Gutter(value string) gridNav {
	return g.set("gutter", value)
}

// IconRatio 图标宽度占比，单位%
func (g gridNav) IconRatio(value string) gridNav {
	return g.set("iconRatio", value)
}

// ItemClassName 列表项 css 类名
func (g gridNav) ItemClassName(value string) gridNav {
	return g.set("itemClassName", value)
}

// Options 列表项图标
func (g gridNav) Options(value ...any) gridNav {
	return g.set("options", value)
}

// Reverse 是否调换图标和文本的位置
func (g gridNav) Reverse(value bool) gridNav {
	return g.set("reverse", value)
}

// Source 数据源
func (g gridNav) Source(value string) gridNav {
	return g.set("source", value)
}

// Square 是否将列表项固定为正方形
func (g gridNav) Square(value bool) gridNav {
	return g.set("square", value)
}

// Value 图片数组
func (g gridNav) Value(value string) gridNav {
	return g.set("value", value)
}
