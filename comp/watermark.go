package comp

import "github.com/zrcoder/amisgo/model"

// watermark represents a watermark schema
type watermark model.Schema

// Watermark creates a new Watermark instance
func Watermark() watermark {
	return watermark{"type": "custom-watermark"}
}

func (wm watermark) set(key string, value any) watermark {
	wm[key] = value
	return wm
}

// Body sets the content
func (wm watermark) Body(value ...any) watermark {
	return wm.set("body", value)
}

// ClassName sets the class name
func (wm watermark) ClassName(value string) watermark {
	return wm.set("className", value)
}

// Content sets the watermark text
func (wm watermark) Content(value string) watermark {
	return wm.set("content", value)
}

// Font sets the font style
func (wm watermark) Font(value string) watermark {
	return wm.set("font", value)
}

// Gap sets the gap between watermarks, default: [100, 100]
func (wm watermark) Gap(value string) watermark {
	return wm.set("gap", value)
}

// Height sets the watermark height, default: 64
func (wm watermark) Height(value string) watermark {
	return wm.set("height", value)
}

// Image sets the image source, supports base64
func (wm watermark) Image(value string) watermark {
	return wm.set("image", value)
}

// Inherit sets whether to inherit the watermark to pop-up components
func (wm watermark) Inherit(value bool) watermark {
	return wm.set("inherit", value)
}

// Offset sets the offset from the top-left corner, default: gap/2
func (wm watermark) Offset(value string) watermark {
	return wm.set("offset", value)
}

// Rotate sets the rotation angle, default: -22
func (wm watermark) Rotate(value string) watermark {
	return wm.set("rotate", value)
}

// Type sets the renderer type to watermark
func (wm watermark) Type(value string) watermark {
	return wm.set("type", value)
}

// Width sets the watermark width, default: 120
func (wm watermark) Width(value string) watermark {
	return wm.set("width", value)
}

// ZIndex sets the z-index of the watermark element, default: 9
func (wm watermark) ZIndex(value string) watermark {
	return wm.set("zIndex", value)
}
