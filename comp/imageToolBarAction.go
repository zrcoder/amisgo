package comp

// imageToolbarAction represents an image toolbar action
type imageToolbarAction Schema

// ImageToolbarAction creates a new ImageToolbarAction instance with a default key
func ImageToolbarAction() imageToolbarAction {
	return make(imageToolbarAction).set("key", "ROTATE_RIGHT")
}

func (a imageToolbarAction) set(key string, value any) imageToolbarAction {
	a[key] = value
	return a
}

// ConfirmTitle sets the confirm dialog title
func (a imageToolbarAction) ConfirmTitle(value any) imageToolbarAction {
	return a.set("confirmTitle", value)
}

// Disabled sets the disabled state
func (a imageToolbarAction) Disabled(value bool) imageToolbarAction {
	return a.set("disabled", value)
}

// Icon sets the icon
func (a imageToolbarAction) Icon(value string) imageToolbarAction {
	return a.set("icon", value)
}

// IconClassName sets the icon class name
func (a imageToolbarAction) IconClassName(value string) imageToolbarAction {
	return a.set("iconClassName", value)
}

// Key sets the action key (ROTATE_RIGHT, ROTATE_LEFT, ZOOM_IN, ZOOM_OUT, SCALE_ORIGIN)
func (a imageToolbarAction) Key(value string) imageToolbarAction {
	return a.set("key", value)
}

// Label sets the label
func (a imageToolbarAction) Label(value string) imageToolbarAction {
	return a.set("label", value)
}
