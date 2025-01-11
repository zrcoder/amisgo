package comp

import "github.com/zrcoder/amisgo/model"

// schemaMessage represents a message configuration. Remember, this has the lowest priority. If your interface returns a msg, the interface return takes priority.

type schemaMessage model.Schema

// model.SchemaMessage creates a new model.SchemaMessage instance
func SchemaMessage() schemaMessage {
	return schemaMessage{}
}

func (s schemaMessage) set(key string, value any) schemaMessage {
	s[key] = value
	return s
}

// FetchFailed the prompt when fetching fails
func (s schemaMessage) FetchFailed(value string) schemaMessage {
	return s.set("fetchFailed", value)
}

// FetchSuccess the prompt when fetching succeeds, default is empty.
func (s schemaMessage) FetchSuccess(value string) schemaMessage {
	return s.set("fetchSuccess", value)
}

// SaveFailed the prompt when saving fails.
func (s schemaMessage) SaveFailed(value string) schemaMessage {
	return s.set("saveFailed", value)
}

// SaveSuccess the prompt when saving succeeds.
func (s schemaMessage) SaveSuccess(value string) schemaMessage {
	return s.set("saveSuccess", value)
}
