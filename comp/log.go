package comp

// Log 实时日志
//
// @version 6.7.0
type Log Schema

// NewLog 创建一个新的 Log 实例
func NewLog() Log {
	return Log{}.set("type", "log")
}

func (l Log) set(key string, value interface{}) Log {
	l[key] = value
	return l
}

// AutoScroll 是否自动滚动到底部
func (l Log) AutoScroll(value string) Log {
	l.set("autoScroll", value)
	return l
}

// ClassName 外层 CSS 类名
func (l Log) ClassName(value string) Log {
	l.set("className", value)
	return l
}

// DisableColor 关闭 ANSI 颜色支持
func (l Log) DisableColor(value string) Log {
	l.set("disableColor", value)
	return l
}

// Encoding 返回内容的字符编码
func (l Log) Encoding(value string) Log {
	l.set("encoding", value)
	return l
}

// Height 展示区域高度
func (l Log) Height(value string) Log {
	l.set("height", value)
	return l
}

// MaxLength 最大显示行数
func (l Log) MaxLength(value string) Log {
	l.set("maxLength", value)
	return l
}

// Operation 可选日志操作：['stop', 'clear', 'showLineNumber', 'filter'] 可选值: stop | clear | showLineNumber | filter
func (l Log) Operation(value string) Log {
	l.set("operation", value)
	return l
}

// Placeholder 加载中的文字
func (l Log) Placeholder(value string) Log {
	l.set("placeholder", value)
	return l
}

// RowHeight 设置每行高度，将会开启虚拟渲染
func (l Log) RowHeight(value string) Log {
	l.set("rowHeight", value)
	return l
}

// Source 接口
func (l Log) Source(value string) Log {
	l.set("source", value)
	return l
}
