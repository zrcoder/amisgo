package comp

// imageToolbarAction 图片工具栏操作
type imageToolbarAction Schema

// ImageToolbarAction 创建一个新的 ImageToolbarAction 实例，并设置默认的 key
func ImageToolbarAction() imageToolbarAction {
	return make(imageToolbarAction).set("key", "ROTATE_RIGHT")
}

func (a imageToolbarAction) set(key string, value any) imageToolbarAction {
	a[key] = value
	return a
}

// ConfirmTitle 确认弹窗标题
func (a imageToolbarAction) ConfirmTitle(value any) imageToolbarAction {
	return a.set("confirmTitle", value)
}

// Disabled
func (a imageToolbarAction) Disabled(value bool) imageToolbarAction {
	return a.set("disabled", value)
}

// Icon
func (a imageToolbarAction) Icon(value string) imageToolbarAction {
	return a.set("icon", value)
}

// IconClassName
func (a imageToolbarAction) IconClassName(value string) imageToolbarAction {
	return a.set("iconClassName", value)
}

// Key 可选值: ROTATE_RIGHT | ROTATE_LEFT | ZOOM_IN | ZOOM_OUT | SCALE_ORIGIN
func (a imageToolbarAction) Key(value string) imageToolbarAction {
	return a.set("key", value)
}

// Label
func (a imageToolbarAction) Label(value string) imageToolbarAction {
	return a.set("label", value)
}
