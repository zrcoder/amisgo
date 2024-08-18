package comp

// ListenerAction 监听器操作
//
// @version 6.7.0
type ListenerAction Schema

// NewListenerAction 创建一个新的 ListenerAction 实例
func NewListenerAction() ListenerAction {
	return make(ListenerAction)
}

// set 设置字段值
func (la ListenerAction) set(key string, value interface{}) ListenerAction {
	la[key] = value
	return la
}

// ActionType 设置操作类型
func (la ListenerAction) ActionType(value string) ListenerAction {
	return la.set("actionType", value)
}

// Args 设置参数
func (la ListenerAction) Args(value string) ListenerAction {
	return la.set("args", value)
}

// ComponentId 设置组件 ID
func (la ListenerAction) ComponentId(value string) ListenerAction {
	return la.set("componentId", value)
}

// ComponentName 设置组件名称
func (la ListenerAction) ComponentName(value string) ListenerAction {
	return la.set("componentName", value)
}

// ConfirmTitle 设置确认弹窗标题
func (la ListenerAction) ConfirmTitle(value string) ListenerAction {
	return la.set("confirmTitle", value)
}

// Data 设置数据
func (la ListenerAction) Data(value string) ListenerAction {
	return la.set("data", value)
}

// DataMergeMode 设置数据合并模式
// 可选值: merge | override
func (la ListenerAction) DataMergeMode(value string) ListenerAction {
	return la.set("dataMergeMode", value)
}

// Description 设置描述
func (la ListenerAction) Description(value string) ListenerAction {
	return la.set("description", value)
}

// ExecOn 设置执行条件
func (la ListenerAction) ExecOn(value string) ListenerAction {
	return la.set("execOn", value)
}

// Expression 设置表达式
func (la ListenerAction) Expression(value string) ListenerAction {
	return la.set("expression", value)
}

// IgnoreError 设置是否忽略错误
func (la ListenerAction) IgnoreError(value bool) ListenerAction {
	return la.set("ignoreError", value)
}

// OutputVar 设置输出变量
func (la ListenerAction) OutputVar(value string) ListenerAction {
	return la.set("outputVar", value)
}

// PreventDefault 设置是否阻止默认行为
func (la ListenerAction) PreventDefault(value bool) ListenerAction {
	return la.set("preventDefault", value)
}

// StopPropagation 设置是否阻止事件传播
func (la ListenerAction) StopPropagation(value bool) ListenerAction {
	return la.set("stopPropagation", value)
}
