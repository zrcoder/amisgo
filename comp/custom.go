package comp

// Custom 自定义组件
type Custom Schema

// NewCustom 创建一个新的 Custom 实例，并设置默认的 type
func NewCustom() Custom {
	return make(Custom).set("type", "custom")
}

func (c Custom) set(key string, value interface{}) Custom {
	c[key] = value
	return c
}

// ClassName 节点 class
func (c Custom) ClassName(value string) Custom {
	return c.set("className", value)
}

// HTML 初始化节点 html
func (c Custom) HTML(value string) Custom {
	return c.set("html", value)
}

// ID 节点 id
func (c Custom) ID(value string) Custom {
	return c.set("id", value)
}

// Inline 默认使用 div 标签，如果 true 就使用 span 标签
func (c Custom) Inline(value bool) Custom {
	return c.set("inline", value)
}

// Name 节点 名称
func (c Custom) Name(value string) Custom {
	return c.set("name", value)
}

// OnMount 节点初始化之后调用的函数
func (c Custom) OnMount(value string) Custom {
	return c.set("onMount", value)
}

// OnUnmount 节点销毁的时候调用的函数
func (c Custom) OnUnmount(value string) Custom {
	return c.set("onUnmount", value)
}

// OnUpdate 数据有更新的时候调用的函数
func (c Custom) OnUpdate(value string) Custom {
	return c.set("onUpdate", value)
}
