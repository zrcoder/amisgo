package comp

type MHorizontal Schema

func Horizontal() MHorizontal {
	return MHorizontal{}
}

func (h MHorizontal) set(key string, value any) MHorizontal {
	h[key] = value
	return h
}

func (h MHorizontal) Left(value string) MHorizontal {
	return h.set("left", value)
}

// LeftFixed bool or string('xs' | 'sm' | 'md' | 'lg')
func (h MHorizontal) LeftFixed(value any) MHorizontal {
	return h.set("leftFixed", value)
}

func (h MHorizontal) Right(value string) MHorizontal {
	return h.set("right", value)
}

func (h MHorizontal) Offset(value string) MHorizontal {
	return h.set("offset", value)
}

// Justify 有时表单内容需要两端对齐，可在 horizontal 中增加 justify 配置，注意只对内联控件生效
func (h MHorizontal) Justify(value bool) MHorizontal {
	return h.set("justify", value)
}
