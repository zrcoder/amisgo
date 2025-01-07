package comp

// toast represents a lightweight notification

type toast Schema

// Toast creates a new Toast instance
func Toast() toast {
	return toast{}
}

func (t toast) set(key string, value any) toast {
	t[key] = value
	return t
}

// Body sets the content
func (t toast) Body(value ...any) toast {
	return t.set("body", value)
}

// CloseButton shows or hides the close button
func (t toast) CloseButton(value bool) toast {
	return t.set("closeButton", value)
}

// Items sets the notification content
func (t toast) Items(value ...any) toast {
	return t.set("items", value)
}

// Level sets the icon type: 'info', 'success', 'error', 'warning'
func (t toast) Level(value string) toast {
	return t.set("level", value)
}

// Position sets the display position: top-right, top-center, top-left, bottom-center, bottom-left, bottom-right, center
func (t toast) Position(value string) toast {
	return t.set("position", value)
}

// ShowIcon shows or hides the icon
func (t toast) ShowIcon(value bool) toast {
	return t.set("showIcon", value)
}

// Timeout sets the duration
func (t toast) Timeout(value string) toast {
	return t.set("timeout", value)
}

// Title sets the title
func (t toast) Title(value any) toast {
	return t.set("title", value)
}
