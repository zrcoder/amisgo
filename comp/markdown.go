package comp

// Markdown 渲染
//
// @version 6.7.0
type Markdown Schema

// NewMarkdown 创建一个新的 Markdown 实例
func NewMarkdown() Markdown {
	return Markdown{}.set("type", "markdown")
}

func (m Markdown) set(key string, value interface{}) Markdown {
	m[key] = value
	return m
}

// ClassName 类名
func (m Markdown) ClassName(value string) Markdown {
	m.set("className", value)
	return m
}

// Name 名称
func (m Markdown) Name(value string) Markdown {
	m.set("name", value)
	return m
}

// Options 设置 Markdown 的配置
func (m Markdown) Options(value string) Markdown {
	m.set("options", value)
	return m
}

// Src 外部地址
func (m Markdown) Src(value string) Markdown {
	m.set("src", value)
	return m
}

// Value 静态值
func (m Markdown) Value(value string) Markdown {
	m.set("value", value)
	return m
}
