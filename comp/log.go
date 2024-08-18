package comp

// log 实时日志
//
// @version 6.7.0
type log schema

// Log 创建一个新的 Log 实例
func Log() log {
	return log{}.set("type", "log")
}

func (l log) set(key string, value interface{}) log {
	l[key] = value
	return l
}

// AutoScroll 是否自动滚动到底部
func (l log) AutoScroll(value string) log {
	l.set("autoScroll", value)
	return l
}

// ClassName 外层 CSS 类名
func (l log) ClassName(value string) log {
	l.set("className", value)
	return l
}

// DisableColor 关闭 ANSI 颜色支持
func (l log) DisableColor(value string) log {
	l.set("disableColor", value)
	return l
}

// Encoding 返回内容的字符编码
func (l log) Encoding(value string) log {
	l.set("encoding", value)
	return l
}

// Height 展示区域高度
func (l log) Height(value string) log {
	l.set("height", value)
	return l
}

// MaxLength 最大显示行数
func (l log) MaxLength(value string) log {
	l.set("maxLength", value)
	return l
}

// Operation 可选日志操作：['stop', 'clear', 'showLineNumber', 'filter'] 可选值: stop | clear | showLineNumber | filter
func (l log) Operation(value string) log {
	l.set("operation", value)
	return l
}

// Placeholder 加载中的文字
func (l log) Placeholder(value string) log {
	l.set("placeholder", value)
	return l
}

// RowHeight 设置每行高度，将会开启虚拟渲染
func (l log) RowHeight(value string) log {
	l.set("rowHeight", value)
	return l
}

// Source 接口
func (l log) Source(value string) log {
	l.set("source", value)
	return l
}
