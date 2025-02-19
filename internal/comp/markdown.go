package comp

import "github.com/zrcoder/amisgo/schema"

// Markdown rendering
type Markdown schema.Schema

func NewMarkdown() Markdown {
	return Markdown{"type": "markdown"}
}

func (m Markdown) set(key string, value any) Markdown {
	m[key] = value
	return m
}

// ClassName sets the class name
func (m Markdown) ClassName(value string) Markdown {
	m.set("className", value)
	return m
}

// Name sets the name
func (m Markdown) Name(value string) Markdown {
	m.set("name", value)
	return m
}

// Options sets the Markdown options
func (m Markdown) Options(value schema.Schema) Markdown {
	m.set("options", value)
	return m
}

// Src sets the external source URL
func (m Markdown) Src(value string) Markdown {
	m.set("src", value)
	return m
}

// Value sets the static value
func (m Markdown) Value(value string) Markdown {
	m.set("value", value)
	return m
}
