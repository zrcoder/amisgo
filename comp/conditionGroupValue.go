package comp

// ConditionGroupValue 代表一个条件组合值，支持多种配置项
type ConditionGroupValue Schema

// NewConditionGroupValue 创建一个新的 ConditionGroupValue 实例，并设置默认的 conjunction
func NewConditionGroupValue() ConditionGroupValue {
	return make(ConditionGroupValue).set("conjunction", "and")
}

func (c ConditionGroupValue) set(key string, value interface{}) ConditionGroupValue {
	c[key] = value
	return c
}

// Children 设置子条件
func (c ConditionGroupValue) Children(value string) ConditionGroupValue {
	return c.set("children", value)
}

// Conjunction 设置连接词，可选值: and | or
func (c ConditionGroupValue) Conjunction(value string) ConditionGroupValue {
	return c.set("conjunction", value)
}

// ID 设置 ID
func (c ConditionGroupValue) ID(value string) ConditionGroupValue {
	return c.set("id", value)
}

// If 设置条件
func (c ConditionGroupValue) If(value string) ConditionGroupValue {
	return c.set("if", value)
}

// Not 设置是否取反
func (c ConditionGroupValue) Not(value bool) ConditionGroupValue {
	return c.set("not", value)
}
