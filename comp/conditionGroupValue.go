package comp

import "github.com/zrcoder/amisgo/model"

// ConditionGroupValue represents a condition group value with multiple configuration options
type conditionGroupValue model.Schema

// ConditionGroupValue creates a new ConditionGroupValue instance with a default conjunction
func ConditionGroupValue() conditionGroupValue {
	return make(conditionGroupValue).set("conjunction", "and")
}

func (c conditionGroupValue) set(key string, value any) conditionGroupValue {
	c[key] = value
	return c
}

// Children sets the child conditions
func (c conditionGroupValue) Children(value string) conditionGroupValue {
	return c.set("children", value)
}

// Conjunction sets the logical connector, with options: "and" | "or"
func (c conditionGroupValue) Conjunction(value string) conditionGroupValue {
	return c.set("conjunction", value)
}

// ID sets a unique identifier
func (c conditionGroupValue) ID(value string) conditionGroupValue {
	return c.set("id", value)
}

// If sets the condition expression
func (c conditionGroupValue) If(value string) conditionGroupValue {
	return c.set("if", value)
}

// Not enables or disables condition negation
func (c conditionGroupValue) Not(value bool) conditionGroupValue {
	return c.set("not", value)
}
