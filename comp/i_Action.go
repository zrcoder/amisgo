package comp

type eventAction schema

// EventAction 创建一个新的 EventAction 实例
func EventAction() eventAction {
	return eventAction{}
}

// set 设置字段值
func (la eventAction) set(key string, value any) eventAction {
	la[key] = value
	return la
}

// ActionType 设置操作类型
// 可选值：
// toast ｜ ajax ｜ dialog ｜ closeDialog ｜ drawer ｜ closeDrawer ｜ confirmDialog ｜ alert
// url | link | goBack | refresh | copy | print | email | setEventData | wait | custom
// setValue | reload | show | hidden | enabled | disabled | static | nonstatic
// broadcast | loop | break | continue | switch | parallel |
func (la eventAction) ActionType(value string) eventAction {
	return la.set("actionType", value)
}

// Args 设置参数
func (la eventAction) Args(value Schema) eventAction {
	return la.set("args", value)
}

// ComponentId 设置组件 ID
func (la eventAction) ComponentId(value string) eventAction {
	return la.set("componentId", value)
}

// ComponentName 设置组件名称
func (la eventAction) ComponentName(value string) eventAction {
	return la.set("componentName", value)
}

// ConfirmTitle 设置确认弹窗标题
func (la eventAction) ConfirmTitle(value any) eventAction {
	return la.set("confirmTitle", value)
}

// Data 设置数据
func (la eventAction) Data(value Data) eventAction {
	return la.set("data", value)
}

// DataMergeMode 设置数据合并模式
// 可选值: merge | override
func (la eventAction) DataMergeMode(value string) eventAction {
	return la.set("dataMergeMode", value)
}

// Description 设置描述
func (la eventAction) Description(value string) eventAction {
	return la.set("description", value)
}

// ExecOn 设置执行条件
func (la eventAction) ExecOn(value string) eventAction {
	return la.set("execOn", value)
}

// Expression 执行条件 boolean|表达式|ConditionBuilder
func (la eventAction) Expression(value any) eventAction {
	return la.set("expression", value)
}

// IgnoreError 设置是否忽略错误
func (la eventAction) IgnoreError(value bool) eventAction {
	return la.set("ignoreError", value)
}

// OutputVar 设置输出变量
func (la eventAction) OutputVar(value string) eventAction {
	return la.set("outputVar", value)
}

// PreventDefault 设置是否阻止默认行为 boolean|表达式|ConditionBuilder
func (la eventAction) PreventDefault(value any) eventAction {
	return la.set("preventDefault", value)
}

// StopPropagation 设置是否阻止事件传播 boolean|表达式|ConditionBuilder
func (la eventAction) StopPropagation(value any) eventAction {
	return la.set("stopPropagation", value)
}
