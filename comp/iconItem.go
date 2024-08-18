package comp

// IconItem 图标项渲染器
type IconItem Schema

// NewIconItem 创建一个新的 IconItem 实例
func NewIconItem() IconItem {
	return make(IconItem)
}

func (i IconItem) set(key string, value interface{}) IconItem {
	i[key] = value
	return i
}

// Icon iconfont 里面的类名
func (i IconItem) Icon(value string) IconItem {
	return i.set("icon", value)
}

// Position 位置
func (i IconItem) Position(value string) IconItem {
	return i.set("position", value)
}
