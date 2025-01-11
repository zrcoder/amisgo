package comp

import "github.com/zrcoder/amisgo/model"

// log represents a real-time log component

type log model.Schema

// Log creates a new log instance
func Log() log {
	return log{"type": "log"}
}

func (l log) set(key string, value any) log {
	l[key] = value
	return l
}

// AutoScroll enables auto-scrolling to the bottom
func (l log) AutoScroll(value string) log {
	l.set("autoScroll", value)
	return l
}

// ClassName sets the outer CSS class name
func (l log) ClassName(value string) log {
	l.set("className", value)
	return l
}

// DisableColor disables ANSI color support
func (l log) DisableColor(value string) log {
	l.set("disableColor", value)
	return l
}

// Encoding sets the character encoding of the content
func (l log) Encoding(value string) log {
	l.set("encoding", value)
	return l
}

// Height sets the display area height
func (l log) Height(value string) log {
	l.set("height", value)
	return l
}

// MaxLength sets the maximum number of display lines
func (l log) MaxLength(value string) log {
	l.set("maxLength", value)
	return l
}

// Operation sets optional log operations: stop, clear, showLineNumber, filter
func (l log) Operation(value string) log {
	l.set("operation", value)
	return l
}

// Placeholder sets the loading text
func (l log) Placeholder(value string) log {
	l.set("placeholder", value)
	return l
}

// RowHeight sets the height of each row, enabling virtual rendering
func (l log) RowHeight(value string) log {
	l.set("rowHeight", value)
	return l
}

// Source sets the data source
func (l log) Source(value string) log {
	l.set("source", value)
	return l
}
