package comp

import "github.com/zrcoder/amisgo/schema"

// IconItem represents an icon item renderer
type IconItem schema.Schema

func NewIconItem() IconItem {
	return make(IconItem)
}

func (i IconItem) set(key string, value any) IconItem {
	i[key] = value
	return i
}

// Icon sets the iconfont class name
func (i IconItem) Icon(value string) IconItem {
	return i.set("icon", value)
}

// Position sets the position
func (i IconItem) Position(value string) IconItem {
	return i.set("position", value)
}
