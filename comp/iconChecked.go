package comp

// iconChecked 图标选中渲染器
type iconChecked schema

// IconChecked 创建一个新的 IconChecked 实例
func IconChecked() iconChecked {
	return make(iconChecked)
}

func (i iconChecked) set(key string, value interface{}) iconChecked {
	i[key] = value
	return i
}

// ID 组件唯一 id
func (i iconChecked) ID(value string) iconChecked {
	return i.set("id", value)
}

// Name 组件名称
func (i iconChecked) Name(value string) iconChecked {
	return i.set("name", value)
}

// SVG SVG 图标
func (i iconChecked) SVG(value string) iconChecked {
	return i.set("svg", value)
}
