package comp

type MEventAction schema

// EventAction 创建一个新的 EventAction 实例
func EventAction() MEventAction {
	return MEventAction{}
}

// set 设置字段值
func (la MEventAction) set(key string, value any) MEventAction {
	la[key] = value
	return la
}

// ActionType 设置操作类型
// 可选值：
// toast ｜ ajax ｜ dialog ｜ closeDialog ｜ drawer ｜ closeDrawer ｜ confirmDialog ｜ alert
// url | link | goBack | refresh | copy | print | email | setEventData | wait
// setValue | reload | show | hidden | enabled | disabled | static | nonstatic
// broadcast | loop | break | continue | switch | parallel |
// custom
func (la MEventAction) ActionType(value string) MEventAction {
	return la.set("actionType", value)
}

// Script 在 actionType 为 custom 时设置自定义脚本
func (ea MEventAction) Script(value string) MEventAction {
	return ea.set("script", value)
}

// Args 设置参数
func (la MEventAction) Args(value Schema) MEventAction {
	return la.set("args", value)
}

// ComponentId 设置组件 ID
func (la MEventAction) ComponentID(value string) MEventAction {
	return la.set("componentId", value)
}

// ComponentName 设置组件名称
func (la MEventAction) ComponentName(value string) MEventAction {
	return la.set("componentName", value)
}

// Value
func (ea MEventAction) Value(value string) MEventAction {
	return ea.set("value", value)
}

// ConfirmTitle 设置确认弹窗标题
func (la MEventAction) ConfirmTitle(value any) MEventAction {
	return la.set("confirmTitle", value)
}

// Data 设置数据
func (la MEventAction) Data(value Data) MEventAction {
	return la.set("data", value)
}

// DataMergeMode 设置数据合并模式
// 可选值: merge | override
func (la MEventAction) DataMergeMode(value string) MEventAction {
	return la.set("dataMergeMode", value)
}

// Description 设置描述
func (la MEventAction) Description(value string) MEventAction {
	return la.set("description", value)
}

// ExecOn 设置执行条件
func (la MEventAction) ExecOn(value string) MEventAction {
	return la.set("execOn", value)
}

// Expression 执行条件 boolean|表达式|ConditionBuilder
func (la MEventAction) Expression(value any) MEventAction {
	return la.set("expression", value)
}

// IgnoreError 设置是否忽略错误
func (la MEventAction) IgnoreError(value bool) MEventAction {
	return la.set("ignoreError", value)
}

// OutputVar 设置输出变量
func (la MEventAction) OutputVar(value string) MEventAction {
	return la.set("outputVar", value)
}

// PreventDefault 设置是否阻止默认行为 boolean|表达式|ConditionBuilder
func (la MEventAction) PreventDefault(value any) MEventAction {
	return la.set("preventDefault", value)
}

// StopPropagation 设置是否阻止事件传播 boolean|表达式|ConditionBuilder
func (la MEventAction) StopPropagation(value any) MEventAction {
	return la.set("stopPropagation", value)
}
