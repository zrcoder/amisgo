package comp

// markdown 渲染
//
// @version 6.7.0
type markdown schema

// Markdown 创建一个新的 Markdown 实例
func Markdown() markdown {
	return markdown{}.set("type", "markdown")
}

func (m markdown) set(key string, value interface{}) markdown {
	m[key] = value
	return m
}

// ClassName 类名
func (m markdown) ClassName(value string) markdown {
	m.set("className", value)
	return m
}

// Name 名称
func (m markdown) Name(value string) markdown {
	m.set("name", value)
	return m
}

// Options 设置 Markdown 的配置
func (m markdown) Options(value string) markdown {
	m.set("options", value)
	return m
}

// Src 外部地址
func (m markdown) Src(value string) markdown {
	m.set("src", value)
	return m
}

// Value 静态值
func (m markdown) Value(value string) markdown {
	m.set("value", value)
	return m
}
