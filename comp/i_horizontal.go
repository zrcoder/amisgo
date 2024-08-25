package comp

type horizontal schema

func Horizontal() horizontal {
	return horizontal{}
}

func (h horizontal) set(key string, value any) horizontal {
	h[key] = value
	return h
}

func (h horizontal) Left(value string) horizontal {
	return h.set("left", value)
}

// LeftFixed bool or string('xs' | 'sm' | 'md' | 'lg')
func (h horizontal) LeftFixed(value any) horizontal {
	return h.set("leftFixed", value)
}

func (h horizontal) Right(value string) horizontal {
	return h.set("right", value)
}

func (h horizontal) Offset(value string) horizontal {
	return h.set("offset", value)
}

// Justify 有时表单内容需要两端对齐，可在 horizontal 中增加 justify 配置，注意只对内联控件生效
func (h horizontal) Justify(value bool) horizontal {
	return h.set("justify", value)
}
