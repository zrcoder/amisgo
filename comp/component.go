package comp

// component 代表一个灵活的组件，可以配置各种属性
type component Schema

// Component 创建一个新的 Component 实例，并设置默认类型
func Component() component {
	return make(component)
}

// Set 用于设置组件的其他属性
func (c component) Set(key string, value any) component {
	return c.set(key, value)
}

func (c component) set(key string, value any) component {
	c[key] = value
	return c
}
