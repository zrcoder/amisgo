package model

type HorizontalSchema Schema

func Horizontal() HorizontalSchema {
	return HorizontalSchema{}
}

func (h HorizontalSchema) set(key string, value any) HorizontalSchema {
	h[key] = value
	return h
}

func (h HorizontalSchema) Left(value string) HorizontalSchema {
	return h.set("left", value)
}

// LeftFixed bool or string('xs' | 'sm' | 'md' | 'lg')
func (h HorizontalSchema) LeftFixed(value any) HorizontalSchema {
	return h.set("leftFixed", value)
}

func (h HorizontalSchema) Right(value string) HorizontalSchema {
	return h.set("right", value)
}

func (h HorizontalSchema) Offset(value string) HorizontalSchema {
	return h.set("offset", value)
}

// Justify 有时表单内容需要两端对齐，可在 horizontal 中增加 justify 配置，注意只对内联控件生效
func (h HorizontalSchema) Justify(value bool) HorizontalSchema {
	return h.set("justify", value)
}
