package comp

// iconItem 图标项渲染器
type iconItem schema

// IconItem 创建一个新的 IconItem 实例
func IconItem() iconItem {
	return make(iconItem)
}

func (i iconItem) set(key string, value any) iconItem {
	i[key] = value
	return i
}

// Icon iconfont 里面的类名
func (i iconItem) Icon(value string) iconItem {
	return i.set("icon", value)
}

// Position 位置
func (i iconItem) Position(value string) iconItem {
	return i.set("position", value)
}
