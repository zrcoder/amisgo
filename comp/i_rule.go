package comp

type rule schema

// Rule 表单校验规则
func Rule() rule {
	return rule{}
}

func (r rule) set(key string, value any) rule {
	r[key] = value
	return r
}

// Rule 校验规则 如 "data.user && data.email"
func (r rule) Rule(value string) rule {
	return r.set("rule", value)
}

// Message 校验不通过时的提示语
func (r rule) Message(value string) rule {
	return r.set("message", value)
}

// Name
func (r rule) Name(value any) rule {
	return r.set("name", value)
}
