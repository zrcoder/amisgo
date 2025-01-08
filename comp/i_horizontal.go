package comp

import "github.com/zrcoder/amisgo/model"

type MHorizontal model.Schema

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

func (h MHorizontal) Right(value string) MHorizontal {
	return h.set("right", value)
}

func (h MHorizontal) Offset(value string) MHorizontal {
	return h.set("offset", value)
}

// Justify aligns form content to both ends. Only effective for inline controls.
func (h MHorizontal) Justify(value bool) MHorizontal {
	return h.set("justify", value)
}
