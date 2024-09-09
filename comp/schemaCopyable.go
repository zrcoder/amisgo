package comp

// schemaCopyable

type schemaCopyable schema

// SchemaCopyable 创建一个新的 SchemaCopyable 实例
func SchemaCopyable() schemaCopyable {
	return schemaCopyable{}
}

func (s schemaCopyable) set(key string, value any) schemaCopyable {
	s[key] = value
	return s
}

// Content 配置复制时的内容模板。 (支持两种语法，但是不能混着用。分别是：1. `${xxx}` 或者 `${xxx|upperCase}` 2. `<%= data.xxx %>`
// 更多文档：https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template)
func (s schemaCopyable) Content(value string) schemaCopyable {
	return s.set("content", value)
}

// Icon 可以配置图标 (iconfont 里面的类名。)
func (s schemaCopyable) Icon(value string) schemaCopyable {
	return s.set("icon", value)
}

// Tooltip 提示文字内容
func (s schemaCopyable) Tooltip(value string) schemaCopyable {
	return s.set("tooltip", value)
}
