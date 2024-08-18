package comp

// NavOverflow 导航溢出配置
type NavOverflow Schema

// NewNavOverflow 创建一个新的 NavOverflow 实例
func NewNavOverflow() NavOverflow {
	return NavOverflow{}
}

// set 方法用于设置属性并返回自身
func (n NavOverflow) set(key string, value interface{}) NavOverflow {
	n[key] = value
	return n
}

// Enable 设置是否开启响应式收纳
func (n NavOverflow) Enable(value bool) NavOverflow {
	return n.set("enable", value)
}

// ItemWidth 设置导航项目宽度
func (n NavOverflow) ItemWidth(value string) NavOverflow {
	return n.set("itemWidth", value)
}

// MaxVisibleCount 设置最大可显示数量
func (n NavOverflow) MaxVisibleCount(value string) NavOverflow {
	return n.set("maxVisibleCount", value)
}

// OverflowClassName 设置菜单触发按钮 CSS 类名
func (n NavOverflow) OverflowClassName(value string) NavOverflow {
	return n.set("overflowClassName", value)
}

// OverflowIndicator 设置菜单触发按钮的图标
func (n NavOverflow) OverflowIndicator(value string) NavOverflow {
	return n.set("overflowIndicator", value)
}

// OverflowLabel 设置菜单触发按钮的文字
func (n NavOverflow) OverflowLabel(value string) NavOverflow {
	return n.set("overflowLabel", value)
}

// OverflowListClassName 设置菜单外层 CSS 类名
func (n NavOverflow) OverflowListClassName(value string) NavOverflow {
	return n.set("overflowListClassName", value)
}

// OverflowPopoverClassName 设置 Popover 浮层 CSS 类名
func (n NavOverflow) OverflowPopoverClassName(value string) NavOverflow {
	return n.set("overflowPopoverClassName", value)
}

// OverflowSuffix 设置导航列表后缀节点
func (n NavOverflow) OverflowSuffix(value string) NavOverflow {
	return n.set("overflowSuffix", value)
}

// Style 设置自定义样式
func (n NavOverflow) Style(value string) NavOverflow {
	return n.set("style", value)
}

// WrapperComponent 设置包裹导航的外层标签名
func (n NavOverflow) WrapperComponent(value string) NavOverflow {
	return n.set("wrapperComponent", value)
}
