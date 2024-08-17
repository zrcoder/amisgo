package comp

// Component 代表一个灵活的组件，可以配置各种属性
type Component BaseRenderer

// NewComponent 创建一个新的 Component 实例，并设置默认类型
func NewComponent() Component {
	return Component(make(BaseRenderer))
}

// Set 用于设置组件的其他属性
func (c Component) Set(key string, value interface{}) Component {
	return c.set(key, value)
}

func (c Component) set(key string, value interface{}) Component {
	c[key] = value
	return c
}
