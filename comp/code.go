package comp

import "github.com/zrcoder/amisgo/model"

// Code represents a code display and editing component
type code model.Schema

// Code creates a new Code instance
func Code() code {
	return make(code).set("type", "code")
}

func (c code) set(key string, value any) code {
	c[key] = value
	return c
}

// ClassName sets the CSS class name for the code component
func (c code) ClassName(value string) code {
	return c.set("className", value)
}

// EditorTheme sets the syntax highlighting theme for the code editor
func (c code) EditorTheme(value string) code {
	return c.set("editorTheme", value)
}

// Language sets the programming language for syntax highlighting
func (c code) Language(value string) code {
	return c.set("language", value)
}

// Name sets the field name for the code component
func (c code) Name(value string) code {
	return c.set("name", value)
}

// TabSize sets the number of spaces for a single tab indentation
func (c code) TabSize(value string) code {
	return c.set("tabSize", value)
}

// Value sets the initial code content to be displayed
func (c code) Value(value string) code {
	return c.set("value", value)
}

// WordWrap enables or disables automatic line wrapping in the code editor
func (c code) WordWrap(value bool) code {
	return c.set("wordWrap", value)
}
