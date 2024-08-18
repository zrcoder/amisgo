package comp

// IconChecked 图标选中渲染器
type IconChecked Schema

// NewIconChecked 创建一个新的 IconChecked 实例
func NewIconChecked() IconChecked {
	return make(IconChecked)
}

func (i IconChecked) set(key string, value interface{}) IconChecked {
	i[key] = value
	return i
}

// ID 组件唯一 id
func (i IconChecked) ID(value string) IconChecked {
	return i.set("id", value)
}

// Name 组件名称
func (i IconChecked) Name(value string) IconChecked {
	return i.set("name", value)
}

// SVG SVG 图标
func (i IconChecked) SVG(value string) IconChecked {
	return i.set("svg", value)
}
