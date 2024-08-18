package comp

// listenerAction 监听器操作
//
// @version 6.7.0
type listenerAction schema

// ListenerAction 创建一个新的 ListenerAction 实例
func ListenerAction() listenerAction {
	return make(listenerAction)
}

// set 设置字段值
func (la listenerAction) set(key string, value interface{}) listenerAction {
	la[key] = value
	return la
}

// ActionType 设置操作类型
func (la listenerAction) ActionType(value string) listenerAction {
	return la.set("actionType", value)
}

// Args 设置参数
func (la listenerAction) Args(value string) listenerAction {
	return la.set("args", value)
}

// ComponentId 设置组件 ID
func (la listenerAction) ComponentId(value string) listenerAction {
	return la.set("componentId", value)
}

// ComponentName 设置组件名称
func (la listenerAction) ComponentName(value string) listenerAction {
	return la.set("componentName", value)
}

// ConfirmTitle 设置确认弹窗标题
func (la listenerAction) ConfirmTitle(value string) listenerAction {
	return la.set("confirmTitle", value)
}

// Data 设置数据
func (la listenerAction) Data(value string) listenerAction {
	return la.set("data", value)
}

// DataMergeMode 设置数据合并模式
// 可选值: merge | override
func (la listenerAction) DataMergeMode(value string) listenerAction {
	return la.set("dataMergeMode", value)
}

// Description 设置描述
func (la listenerAction) Description(value string) listenerAction {
	return la.set("description", value)
}

// ExecOn 设置执行条件
func (la listenerAction) ExecOn(value string) listenerAction {
	return la.set("execOn", value)
}

// Expression 设置表达式
func (la listenerAction) Expression(value string) listenerAction {
	return la.set("expression", value)
}

// IgnoreError 设置是否忽略错误
func (la listenerAction) IgnoreError(value bool) listenerAction {
	return la.set("ignoreError", value)
}

// OutputVar 设置输出变量
func (la listenerAction) OutputVar(value string) listenerAction {
	return la.set("outputVar", value)
}

// PreventDefault 设置是否阻止默认行为
func (la listenerAction) PreventDefault(value bool) listenerAction {
	return la.set("preventDefault", value)
}

// StopPropagation 设置是否阻止事件传播
func (la listenerAction) StopPropagation(value bool) listenerAction {
	return la.set("stopPropagation", value)
}
