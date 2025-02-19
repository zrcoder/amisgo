package comp

import "github.com/zrcoder/amisgo/schema"

// Log represents a real-time Log component
type Log schema.Schema

func NewLog() Log {
	return Log{"type": "log"}
}

func (l Log) set(key string, value any) Log {
	l[key] = value
	return l
}

// AutoScroll enables auto-scrolling to the bottom
func (l Log) AutoScroll(value string) Log {
	l.set("autoScroll", value)
	return l
}

// ClassName sets the outer CSS class name
func (l Log) ClassName(value string) Log {
	l.set("className", value)
	return l
}

// DisableColor disables ANSI color support
func (l Log) DisableColor(value string) Log {
	l.set("disableColor", value)
	return l
}

// Encoding sets the character encoding of the content
func (l Log) Encoding(value string) Log {
	l.set("encoding", value)
	return l
}

// Height sets the display area height
func (l Log) Height(value string) Log {
	l.set("height", value)
	return l
}

// MaxLength sets the maximum number of display lines
func (l Log) MaxLength(value string) Log {
	l.set("maxLength", value)
	return l
}

// Operation sets optional log operations: stop, clear, showLineNumber, filter
func (l Log) Operation(value string) Log {
	l.set("operation", value)
	return l
}

// Placeholder sets the loading text
func (l Log) Placeholder(value string) Log {
	l.set("placeholder", value)
	return l
}

// RowHeight sets the height of each row, enabling virtual rendering
func (l Log) RowHeight(value string) Log {
	l.set("rowHeight", value)
	return l
}

// Source sets the data source
func (l Log) Source(value string) Log {
	l.set("source", value)
	return l
}
