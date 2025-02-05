package model

// SchemaMessage represents a message configuration. Remember, this has the lowest priority. If your interface returns a msg, the interface return takes priority.
type SchemaMessage Schema

func NewSchemaMessage() SchemaMessage {
	return SchemaMessage{}
}

func (s SchemaMessage) set(key string, value any) SchemaMessage {
	s[key] = value
	return s
}

// FetchFailed the prompt when fetching fails
func (s SchemaMessage) FetchFailed(value string) SchemaMessage {
	return s.set("fetchFailed", value)
}

// FetchSuccess the prompt when fetching succeeds, default is empty.
func (s SchemaMessage) FetchSuccess(value string) SchemaMessage {
	return s.set("fetchSuccess", value)
}

// SaveFailed the prompt when saving fails.
func (s SchemaMessage) SaveFailed(value string) SchemaMessage {
	return s.set("saveFailed", value)
}

// SaveSuccess the prompt when saving succeeds.
func (s SchemaMessage) SaveSuccess(value string) SchemaMessage {
	return s.set("saveSuccess", value)
}
