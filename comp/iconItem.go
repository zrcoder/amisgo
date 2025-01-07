package comp

// iconItem represents an icon item renderer
type iconItem Schema

// IconItem creates a new IconItem instance
func IconItem() iconItem {
	return make(iconItem)
}

func (i iconItem) set(key string, value any) iconItem {
	i[key] = value
	return i
}

// Icon sets the iconfont class name
func (i iconItem) Icon(value string) iconItem {
	return i.set("icon", value)
}

// Position sets the position
func (i iconItem) Position(value string) iconItem {
	return i.set("position", value)
}
