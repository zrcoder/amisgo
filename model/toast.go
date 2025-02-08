package model

// Toast represents a lightweight notification
type Toast Schema

func NewToast() Toast {
	return Toast{}
}

func (t Toast) set(key string, value any) Toast {
	t[key] = value
	return t
}

// Body sets the content
func (t Toast) Body(value ...any) Toast {
	return t.set("body", value)
}

// CloseButton shows or hides the close button
func (t Toast) CloseButton(value bool) Toast {
	return t.set("closeButton", value)
}

// Items sets the notification content
func (t Toast) Items(value ...any) Toast {
	return t.set("items", value)
}

// Level sets the icon type: 'info', 'success', 'error', 'warning'
func (t Toast) Level(value string) Toast {
	return t.set("level", value)
}

// Position sets the display position: top-right, top-center, top-left, bottom-center, bottom-left, bottom-right, center
func (t Toast) Position(value string) Toast {
	return t.set("position", value)
}

// ShowIcon shows or hides the icon
func (t Toast) ShowIcon(value bool) Toast {
	return t.set("showIcon", value)
}

// Timeout sets the duration
func (t Toast) Timeout(value string) Toast {
	return t.set("timeout", value)
}

// Title sets the title
func (t Toast) Title(value any) Toast {
	return t.set("title", value)
}
