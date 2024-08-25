package comp

type text schema

func Text() text {
	return text{}.set("type", "text")
}

func (t text) set(key string, value any) text {
	t[key] = value
	return t
}

func (t text) Label(value string) text {
	return t.set("label", value)
}
