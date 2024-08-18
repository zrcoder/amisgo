package comp

// Toast 轻提示
//
// @author slowlyo
// @version 6.7.0
type Toast Schema

// NewToast 创建一个新的 Toast 实例
func NewToast() Toast {
	return Toast{}
}

func (t Toast) set(key string, value interface{}) Toast {
	t[key] = value
	return t
}

// Body 内容
func (t Toast) Body(value ...interface{}) Toast {
	return t.set("body", value)
}

// CloseButton 是否显示关闭按钮
func (t Toast) CloseButton(value bool) Toast {
	return t.set("closeButton", value)
}

// Items 轻提示内容
func (t Toast) Items(value string) Toast {
	return t.set("items", value)
}

// Level 展示图标，可选'info'/'success'/'error'/'warning'
func (t Toast) Level(value string) Toast {
	return t.set("level", value)
}

// Position 提示显示位置，可选值: top-right | top-center | top-left | bottom-center | bottom-left | bottom-right | center
func (t Toast) Position(value string) Toast {
	return t.set("position", value)
}

// ShowIcon 是否显示图标
func (t Toast) ShowIcon(value bool) Toast {
	return t.set("showIcon", value)
}

// Timeout 持续时间
func (t Toast) Timeout(value string) Toast {
	return t.set("timeout", value)
}

// Title 标题
func (t Toast) Title(value string) Toast {
	return t.set("title", value)
}
