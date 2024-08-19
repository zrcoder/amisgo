package comp

// toast 轻提示
//

// @version 6.7.0
type toast schema

// Toast 创建一个新的 Toast 实例
func Toast() toast {
	return toast{}
}

func (t toast) set(key string, value any) toast {
	t[key] = value
	return t
}

// Body 内容
func (t toast) Body(value ...any) toast {
	return t.set("body", value)
}

// CloseButton 是否显示关闭按钮
func (t toast) CloseButton(value bool) toast {
	return t.set("closeButton", value)
}

// Items 轻提示内容
func (t toast) Items(value ...any) toast {
	return t.set("items", value)
}

// Level 展示图标，可选'info'/'success'/'error'/'warning'
func (t toast) Level(value string) toast {
	return t.set("level", value)
}

// Position 提示显示位置，可选值: top-right | top-center | top-left | bottom-center | bottom-left | bottom-right | center
func (t toast) Position(value string) toast {
	return t.set("position", value)
}

// ShowIcon 是否显示图标
func (t toast) ShowIcon(value bool) toast {
	return t.set("showIcon", value)
}

// Timeout 持续时间
func (t toast) Timeout(value string) toast {
	return t.set("timeout", value)
}

// Title 标题
func (t toast) Title(value any) toast {
	return t.set("title", value)
}
