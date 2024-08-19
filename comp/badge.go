package comp

// badge 角标渲染器
type badge schema

// Badge 创建一个新的 Badge 实例
func Badge() badge {
	return make(badge)
}

func (b badge) set(key string, value any) badge {
	b[key] = value
	return b
}

// Animation 设置是否显示动画
func (b badge) Animation(value bool) badge {
	return b.set("animation", value)
}

// ClassName 设置类名
func (b badge) ClassName(value string) badge {
	return b.set("className", value)
}

// Level 设置提示类型
func (b badge) Level(value string) badge {
	return b.set("level", value)
}

// Mode 设置角标类型
func (b badge) Mode(value string) badge {
	return b.set("mode", value)
}

// Offset 设置角标位置，相对于position的位置进行偏移
func (b badge) Offset(value string) badge {
	return b.set("offset", value)
}

// OverflowCount 设置封顶的数字值
func (b badge) OverflowCount(value string) badge {
	return b.set("overflowCount", value)
}

// Position 设置角标位置
func (b badge) Position(value string) badge {
	return b.set("position", value)
}

// Size 设置大小
func (b badge) Size(value string) badge {
	return b.set("size", value)
}

// Style 设置角标的自定义样式
func (b badge) Style(value any) badge {
	return b.set("style", value)
}

// Text 设置文本内容
func (b badge) Text(value string) badge {
	return b.set("text", value)
}

// VisibleOn 设置动态控制是否显示
func (b badge) VisibleOn(value string) badge {
	return b.set("visibleOn", value)
}
