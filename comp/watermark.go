package comp

// Watermark 水印
//
// @author  slowlyo
// @version 6.7.0
type Watermark Schema

// NewWatermark 创建一个新的 Watermark 实例
func NewWatermark() Watermark {
	return Watermark{}.set("type", "custom-watermark")
}

func (wm Watermark) set(key string, value interface{}) Watermark {
	wm[key] = value
	return wm
}

// Body 内容
func (wm Watermark) Body(value string) Watermark {
	return wm.set("body", value)
}

// ClassName 设置类名
func (wm Watermark) ClassName(value string) Watermark {
	return wm.set("className", value)
}

// Content 水印文字内容
func (wm Watermark) Content(value string) Watermark {
	return wm.set("content", value)
}

// Font 文字样式
func (wm Watermark) Font(value string) Watermark {
	return wm.set("font", value)
}

// Gap 水印之间的间距, 默认: [100, 100]
func (wm Watermark) Gap(value string) Watermark {
	return wm.set("gap", value)
}

// Height 水印的高度，content 的默认值为自身的高度, 默认: 64
func (wm Watermark) Height(value string) Watermark {
	return wm.set("height", value)
}

// Image 图片源，建议导出 2 倍或 3 倍图，优先级高 (支持 base64 格式)
func (wm Watermark) Image(value string) Watermark {
	return wm.set("image", value)
}

// Inherit 是否将水印传导给弹出组件如 Modal、Drawer
func (wm Watermark) Inherit(value bool) Watermark {
	return wm.set("inherit", value)
}

// Offset 水印距离容器左上角的偏移量，默认为 gap/2
func (wm Watermark) Offset(value string) Watermark {
	return wm.set("offset", value)
}

// Rotate 水印绘制时，旋转的角度，单位 °, 默认: -22
func (wm Watermark) Rotate(value string) Watermark {
	return wm.set("rotate", value)
}

// Type 指定为 watermark 渲染器。文档: https://ant-design.antgroup.com/components/watermark-cn#watermark
func (wm Watermark) Type(value string) Watermark {
	return wm.set("type", value)
}

// Width 水印的宽度，content 的默认值为自身的宽度, 默认: 120
func (wm Watermark) Width(value string) Watermark {
	return wm.set("width", value)
}

// ZIndex 追加的水印元素的 z-index, 默认: 9
func (wm Watermark) ZIndex(value string) Watermark {
	return wm.set("zIndex", value)
}
