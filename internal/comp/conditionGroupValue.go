package comp

import "github.com/zrcoder/amisgo/model"

// ConditionGroupValue represents a condition group value with multiple configuration options
type ConditionGroupValue model.Schema

func NewConditionGroupValue() ConditionGroupValue {
	return ConditionGroupValue{"conjunction": "and"}
}

func (c ConditionGroupValue) set(key string, value any) ConditionGroupValue {
	c[key] = value
	return c
}

// Children sets the child conditions
func (c ConditionGroupValue) Children(value string) ConditionGroupValue {
	return c.set("children", value)
}

// Conjunction sets the logical connector, with options: "and" | "or"
func (c ConditionGroupValue) Conjunction(value string) ConditionGroupValue {
	return c.set("conjunction", value)
}

// ID sets a unique identifier
func (c ConditionGroupValue) ID(value string) ConditionGroupValue {
	return c.set("id", value)
}

// If sets the condition expression
func (c ConditionGroupValue) If(value string) ConditionGroupValue {
	return c.set("if", value)
}

// Not enables or disables condition negation
func (c ConditionGroupValue) Not(value bool) ConditionGroupValue {
	return c.set("not", value)
}
