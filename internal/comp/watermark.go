package comp

import "github.com/zrcoder/amisgo/model"

// Watermark represents a Watermark schema
type Watermark model.Schema

func NewWatermark() Watermark {
	return Watermark{"type": "custom-watermark"}
}

func (wm Watermark) set(key string, value any) Watermark {
	wm[key] = value
	return wm
}

// Body sets the content
func (wm Watermark) Body(value ...any) Watermark {
	return wm.set("body", value)
}

// ClassName sets the class name
func (wm Watermark) ClassName(value string) Watermark {
	return wm.set("className", value)
}

// Content sets the watermark text
func (wm Watermark) Content(value string) Watermark {
	return wm.set("content", value)
}

// Font sets the font style
func (wm Watermark) Font(value string) Watermark {
	return wm.set("font", value)
}

// Gap sets the gap between watermarks, default: [100, 100]
func (wm Watermark) Gap(value string) Watermark {
	return wm.set("gap", value)
}

// Height sets the watermark height, default: 64
func (wm Watermark) Height(value string) Watermark {
	return wm.set("height", value)
}

// Image sets the image source, supports base64
func (wm Watermark) Image(value string) Watermark {
	return wm.set("image", value)
}

// Inherit sets whether to inherit the watermark to pop-up components
func (wm Watermark) Inherit(value bool) Watermark {
	return wm.set("inherit", value)
}

// Offset sets the offset from the top-left corner, default: gap/2
func (wm Watermark) Offset(value string) Watermark {
	return wm.set("offset", value)
}

// Rotate sets the rotation angle, default: -22
func (wm Watermark) Rotate(value string) Watermark {
	return wm.set("rotate", value)
}

// Type sets the renderer type to watermark
func (wm Watermark) Type(value string) Watermark {
	return wm.set("type", value)
}

// Width sets the watermark width, default: 120
func (wm Watermark) Width(value string) Watermark {
	return wm.set("width", value)
}

// ZIndex sets the z-index of the watermark element, default: 9
func (wm Watermark) ZIndex(value string) Watermark {
	return wm.set("zIndex", value)
}
