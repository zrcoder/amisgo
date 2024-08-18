package comp

// SchemaCopyable
//
// @author  slowlyo
// @version 6.7.0
type SchemaCopyable Schema

// NewSchemaCopyable 创建一个新的 SchemaCopyable 实例
func NewSchemaCopyable() SchemaCopyable {
	return SchemaCopyable{}
}

func (s SchemaCopyable) set(key string, value interface{}) SchemaCopyable {
	s[key] = value
	return s
}

// Content 配置复制时的内容模板。 (支持两种语法，但是不能混着用。分别是：1. `${xxx}` 或者 `${xxx|upperCase}` 2. `<%= data.xxx %>`
// 更多文档：https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template)
func (s SchemaCopyable) Content(value string) SchemaCopyable {
	return s.set("content", value)
}

// Icon 可以配置图标 (iconfont 里面的类名。)
func (s SchemaCopyable) Icon(value string) SchemaCopyable {
	return s.set("icon", value)
}

// Tooltip 提示文字内容
func (s SchemaCopyable) Tooltip(value string) SchemaCopyable {
	return s.set("tooltip", value)
}
