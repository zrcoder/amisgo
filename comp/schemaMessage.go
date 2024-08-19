package comp

// schemaMessage 消息文案配置，记住这个优先级是最低的，如果你的接口返回了 msg，接口返回的优先。

// @version 6.7.0
type schemaMessage schema

// SchemaMessage 创建一个新的 SchemaMessage 实例
func SchemaMessage() schemaMessage {
	return schemaMessage{}
}

func (s schemaMessage) set(key string, value any) schemaMessage {
	s[key] = value
	return s
}

// FetchFailed 获取失败时的提示
func (s schemaMessage) FetchFailed(value string) schemaMessage {
	return s.set("fetchFailed", value)
}

// FetchSuccess 获取成功的提示，默认为空。
func (s schemaMessage) FetchSuccess(value string) schemaMessage {
	return s.set("fetchSuccess", value)
}

// SaveFailed 保存失败时的提示。
func (s schemaMessage) SaveFailed(value string) schemaMessage {
	return s.set("saveFailed", value)
}

// SaveSuccess 保存成功时的提示。
func (s schemaMessage) SaveSuccess(value string) schemaMessage {
	return s.set("saveSuccess", value)
}
