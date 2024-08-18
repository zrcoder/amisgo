package comp

// ImageToolbarAction 图片工具栏操作
type ImageToolbarAction Schema

// NewImageToolbarAction 创建一个新的 ImageToolbarAction 实例，并设置默认的 key
func NewImageToolbarAction() ImageToolbarAction {
	return make(ImageToolbarAction).set("key", "ROTATE_RIGHT")
}

func (a ImageToolbarAction) set(key string, value interface{}) ImageToolbarAction {
	a[key] = value
	return a
}

// ConfirmTitle 确认弹窗标题
func (a ImageToolbarAction) ConfirmTitle(value string) ImageToolbarAction {
	return a.set("confirmTitle", value)
}

// Disabled
func (a ImageToolbarAction) Disabled(value bool) ImageToolbarAction {
	return a.set("disabled", value)
}

// Icon
func (a ImageToolbarAction) Icon(value string) ImageToolbarAction {
	return a.set("icon", value)
}

// IconClassName
func (a ImageToolbarAction) IconClassName(value string) ImageToolbarAction {
	return a.set("iconClassName", value)
}

// Key 可选值: ROTATE_RIGHT | ROTATE_LEFT | ZOOM_IN | ZOOM_OUT | SCALE_ORIGIN
func (a ImageToolbarAction) Key(value string) ImageToolbarAction {
	return a.set("key", value)
}

// Label
func (a ImageToolbarAction) Label(value string) ImageToolbarAction {
	return a.set("label", value)
}
