package comp

// SchemaMessage 消息文案配置，记住这个优先级是最低的，如果你的接口返回了 msg，接口返回的优先。
//
// @author  slowlyo
// @version 6.7.0
type SchemaMessage Schema

// NewSchemaMessage 创建一个新的 SchemaMessage 实例
func NewSchemaMessage() SchemaMessage {
	return SchemaMessage{}
}

func (s SchemaMessage) set(key string, value interface{}) SchemaMessage {
	s[key] = value
	return s
}

// FetchFailed 获取失败时的提示
func (s SchemaMessage) FetchFailed(value string) SchemaMessage {
	return s.set("fetchFailed", value)
}

// FetchSuccess 获取成功的提示，默认为空。
func (s SchemaMessage) FetchSuccess(value string) SchemaMessage {
	return s.set("fetchSuccess", value)
}

// SaveFailed 保存失败时的提示。
func (s SchemaMessage) SaveFailed(value string) SchemaMessage {
	return s.set("saveFailed", value)
}

// SaveSuccess 保存成功时的提示。
func (s SchemaMessage) SaveSuccess(value string) SchemaMessage {
	return s.set("saveSuccess", value)
}
