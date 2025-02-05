package model

type Rule Schema

func NewRule() Rule {
	return Rule{}
}

func (r Rule) set(key string, value any) Rule {
	r[key] = value
	return r
}

// Rule sets the validation rule, e.g., "data.user && data.email".
func (r Rule) Rule(value string) Rule {
	return r.set("rule", value)
}

// Message sets the error message for validation failure.
func (r Rule) Message(value string) Rule {
	return r.set("message", value)
}

// Name sets the name of the rule.
func (r Rule) Name(value any) Rule {
	return r.set("name", value)
}
