package comp

// custom 自定义组件
type custom Schema

// Custom 创建一个新的 Custom 实例，并设置默认的 type
func Custom() custom {
	return make(custom).set("type", "custom")
}

func (c custom) set(key string, value any) custom {
	c[key] = value
	return c
}

// ClassName 节点 class
func (c custom) ClassName(value string) custom {
	return c.set("className", value)
}

// HTML 初始化节点 html
func (c custom) HTML(value string) custom {
	return c.set("html", value)
}

// ID 节点 id
func (c custom) ID(value string) custom {
	return c.set("id", value)
}

// Inline 默认使用 div 标签，如果 true 就使用 span 标签
func (c custom) Inline(value bool) custom {
	return c.set("inline", value)
}

// Name 节点 名称
func (c custom) Name(value string) custom {
	return c.set("name", value)
}

// OnMount 节点初始化之后调用的函数
func (c custom) OnMount(value string) custom {
	return c.set("onMount", value)
}

// OnUnmount 节点销毁的时候调用的函数
func (c custom) OnUnmount(value string) custom {
	return c.set("onUnmount", value)
}

// OnUpdate 数据有更新的时候调用的函数
func (c custom) OnUpdate(value string) custom {
	return c.set("onUpdate", value)
}
