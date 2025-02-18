package model

type ToastItem Schema

func NewToastItem() ToastItem {
	return ToastItem{}
}

func (i ToastItem) set(key string, value any) ToastItem {
	i[key] = value
	return i
}

// Title sets the title
func (i ToastItem) Title(value any) ToastItem {
	return i.set("title", value)
}

// Body sets the body
func (i ToastItem) Body(value any) ToastItem {
	return i.set("body", value)
}

// Level sets the level
// Valid levels include: 'info', 'success', 'error', 'warning'
func (i ToastItem) Level(value string) ToastItem {
	return i.set("level", value)
}

// Position sets the position property of ToastItem to the specified value.
// Valid positions include: 'top-right', 'top-center', 'top-left', 'bottom-center', 'bottom-left', 'bottom-right', 'center'.
func (i ToastItem) Position(value string) ToastItem {
	return i.set("position", value)
}

// CloseButton sets whether to show a close button on the toast notification.
func (i ToastItem) CloseButton(value bool) ToastItem {
	return i.set("closeButton", value)
}

// ShowIcon sets whether to show the icon
func (i ToastItem) ShowIcon(value bool) ToastItem {
	return i.set("showIcon", value)
}

// Timeout sets the duration in milliseconds before the toast message disappears.
func (i ToastItem) Timeout(value float64) ToastItem {
	return i.set("timeout", value)
}

// AllowHtml sets whether HTML content is allowed in the toast notification.
func (i ToastItem) AllowHtml(value bool) ToastItem {
	return i.set("allowHtml", value)
}
