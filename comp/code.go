package comp

import "github.com/zrcoder/amisgo/model"

// Code represents a Code display and editing component
type Code model.Schema

func NewCode() Code {
	return Code{"type": "code"}
}

func (c Code) set(key string, value any) Code {
	c[key] = value
	return c
}

// ClassName sets the CSS class name for the code component
func (c Code) ClassName(value string) Code {
	return c.set("className", value)
}

// EditorTheme sets the syntax highlighting theme for the code editor
func (c Code) EditorTheme(value string) Code {
	return c.set("editorTheme", value)
}

// Language sets the programming language for syntax highlighting
func (c Code) Language(value string) Code {
	return c.set("language", value)
}

// Name sets the field name for the code component
func (c Code) Name(value string) Code {
	return c.set("name", value)
}

// TabSize sets the number of spaces for a single tab indentation
func (c Code) TabSize(value string) Code {
	return c.set("tabSize", value)
}

// Value sets the initial code content to be displayed
func (c Code) Value(value string) Code {
	return c.set("value", value)
}

// WordWrap enables or disables automatic line wrapping in the code editor
func (c Code) WordWrap(value bool) Code {
	return c.set("wordWrap", value)
}
