package comp

// Badge 角标渲染器
type Badge Schema

// NewBadge 创建一个新的 Badge 实例
func NewBadge() Badge {
	return make(Badge)
}

// Set 设置属性
func (b Badge) set(key string, value interface{}) Badge {
	b[key] = value
	return b
}

// Animation 设置是否显示动画
func (b Badge) Animation(value bool) Badge {
	return b.set("animation", value)
}

// ClassName 设置类名
func (b Badge) ClassName(value string) Badge {
	return b.set("className", value)
}

// Level 设置提示类型
func (b Badge) Level(value string) Badge {
	return b.set("level", value)
}

// Mode 设置角标类型
func (b Badge) Mode(value string) Badge {
	return b.set("mode", value)
}

// Offset 设置角标位置，相对于position的位置进行偏移
func (b Badge) Offset(value string) Badge {
	return b.set("offset", value)
}

// OverflowCount 设置封顶的数字值
func (b Badge) OverflowCount(value string) Badge {
	return b.set("overflowCount", value)
}

// Position 设置角标位置
func (b Badge) Position(value string) Badge {
	return b.set("position", value)
}

// Size 设置大小
func (b Badge) Size(value string) Badge {
	return b.set("size", value)
}

// Style 设置角标的自定义样式
func (b Badge) Style(value string) Badge {
	return b.set("style", value)
}

// Text 设置文本内容
func (b Badge) Text(value string) Badge {
	return b.set("text", value)
}

// VisibleOn 设置动态控制是否显示
func (b Badge) VisibleOn(value string) Badge {
	return b.set("visibleOn", value)
}
