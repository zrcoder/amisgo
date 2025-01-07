package comp

type MRule Schema

// Rule returns an empty validation rule.
func Rule() MRule {
	return MRule{}
}

func (r MRule) set(key string, value any) MRule {
	r[key] = value
	return r
}

// Rule sets the validation rule, e.g., "data.user && data.email".
func (r MRule) Rule(value string) MRule {
	return r.set("rule", value)
}

// Message sets the error message for validation failure.
func (r MRule) Message(value string) MRule {
	return r.set("message", value)
}

// Name sets the name of the rule.
func (r MRule) Name(value any) MRule {
	return r.set("name", value)
}
