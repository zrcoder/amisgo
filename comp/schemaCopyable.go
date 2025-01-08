package comp

import "github.com/zrcoder/amisgo/model"

// schemaCopyable

type schemaCopyable model.Schema

// model.SchemaCopyable creates a new model.SchemaCopyable instance
func SchemaCopyable() schemaCopyable {
	return schemaCopyable{}
}

func (s schemaCopyable) set(key string, value any) schemaCopyable {
	s[key] = value
	return s
}

// Content configures the content template when copying. (Supports two syntaxes, but cannot be mixed. They are: 1. `${xxx}` or `${xxx|upperCase}` 2. `<%= data.xxx %>`)
// More documentation: https://aisuda.bce.baidu.com/amis/zh-CN/docs/concepts/template)
func (s schemaCopyable) Content(value string) schemaCopyable {
	return s.set("content", value)
}

// Icon can configure the icon (class name in iconfont.)
func (s schemaCopyable) Icon(value string) schemaCopyable {
	return s.set("icon", value)
}

// Tooltip sets the tooltip text content
func (s schemaCopyable) Tooltip(value string) schemaCopyable {
	return s.set("tooltip", value)
}
