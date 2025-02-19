package comp

import "github.com/zrcoder/amisgo/schema"

// ImageToolbarAction represents an image toolbar action
type ImageToolbarAction schema.Schema

func NewImageToolbarAction() ImageToolbarAction {
	return ImageToolbarAction{"key": "ROTATE_RIGHT"}
}

func (a ImageToolbarAction) set(key string, value any) ImageToolbarAction {
	a[key] = value
	return a
}

// ConfirmTitle sets the confirm dialog title
func (a ImageToolbarAction) ConfirmTitle(value any) ImageToolbarAction {
	return a.set("confirmTitle", value)
}

// Disabled sets the disabled state
func (a ImageToolbarAction) Disabled(value bool) ImageToolbarAction {
	return a.set("disabled", value)
}

// Icon sets the icon
func (a ImageToolbarAction) Icon(value string) ImageToolbarAction {
	return a.set("icon", value)
}

// IconClassName sets the icon class name
func (a ImageToolbarAction) IconClassName(value string) ImageToolbarAction {
	return a.set("iconClassName", value)
}

// Key sets the action key (ROTATE_RIGHT, ROTATE_LEFT, ZOOM_IN, ZOOM_OUT, SCALE_ORIGIN)
func (a ImageToolbarAction) Key(value string) ImageToolbarAction {
	return a.set("key", value)
}

// Label sets the label
func (a ImageToolbarAction) Label(value string) ImageToolbarAction {
	return a.set("label", value)
}
