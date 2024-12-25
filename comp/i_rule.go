package comp

type MRule schema

// Rule 表单校验规则
func Rule() MRule {
	return MRule{}
}

func (r MRule) set(key string, value any) MRule {
	r[key] = value
	return r
}

// Rule 校验规则 如 "data.user && data.email"
func (r MRule) Rule(value string) MRule {
	return r.set("rule", value)
}

// Message 校验不通过时的提示语
func (r MRule) Message(value string) MRule {
	return r.set("message", value)
}

// Name
func (r MRule) Name(value any) MRule {
	return r.set("name", value)
}
