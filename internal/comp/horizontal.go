package comp

import "github.com/zrcoder/amisgo/schema"

type Horizontal schema.Schema

func NewHorizontal() Horizontal {
	return Horizontal{}
}

func (h Horizontal) set(key string, value any) Horizontal {
	h[key] = value
	return h
}

func (h Horizontal) Left(value string) Horizontal {
	return h.set("left", value)
}

func (h Horizontal) Right(value string) Horizontal {
	return h.set("right", value)
}

func (h Horizontal) Offset(value string) Horizontal {
	return h.set("offset", value)
}

// Justify aligns form content to both ends. Only effective for inline controls.
func (h Horizontal) Justify(value bool) Horizontal {
	return h.set("justify", value)
}
