package model

// SchemaCopyable
type SchemaCopyable Schema

func NewSchemaCopyable() SchemaCopyable {
	return SchemaCopyable{}
}

func (s SchemaCopyable) set(key string, value any) SchemaCopyable {
	s[key] = value
	return s
}

// Content configures the content template when copying. (Supports two syntaxes, but cannot be mixed. They are: 1. `${xxx}` or `${xxx|upperCase}` 2. `<%= data.xxx %>`)
// More documentation: https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template)
func (s SchemaCopyable) Content(value string) SchemaCopyable {
	return s.set("content", value)
}

// Icon can configure the icon (class name in iconfont.)
func (s SchemaCopyable) Icon(value string) SchemaCopyable {
	return s.set("icon", value)
}

// Tooltip sets the tooltip text content
func (s SchemaCopyable) Tooltip(value string) SchemaCopyable {
	return s.set("tooltip", value)
}
