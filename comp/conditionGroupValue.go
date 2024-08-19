package comp

// conditionGroupValue 代表一个条件组合值，支持多种配置项
type conditionGroupValue schema

// ConditionGroupValue 创建一个新的 ConditionGroupValue 实例，并设置默认的 conjunction
func ConditionGroupValue() conditionGroupValue {
	return make(conditionGroupValue).set("conjunction", "and")
}

func (c conditionGroupValue) set(key string, value any) conditionGroupValue {
	c[key] = value
	return c
}

// Children 设置子条件
func (c conditionGroupValue) Children(value string) conditionGroupValue {
	return c.set("children", value)
}

// Conjunction 设置连接词，可选值: and | or
func (c conditionGroupValue) Conjunction(value string) conditionGroupValue {
	return c.set("conjunction", value)
}

// ID 设置 ID
func (c conditionGroupValue) ID(value string) conditionGroupValue {
	return c.set("id", value)
}

// If 设置条件
func (c conditionGroupValue) If(value string) conditionGroupValue {
	return c.set("if", value)
}

// Not 设置是否取反
func (c conditionGroupValue) Not(value bool) conditionGroupValue {
	return c.set("not", value)
}
