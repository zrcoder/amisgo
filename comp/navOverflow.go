package comp

// navOverflow 导航溢出配置
type navOverflow schema

// NavOverflow 创建一个新的 NavOverflow 实例
func NavOverflow() navOverflow {
	return navOverflow{}
}

// set 方法用于设置属性并返回自身
func (n navOverflow) set(key string, value any) navOverflow {
	n[key] = value
	return n
}

// Enable 设置是否开启响应式收纳
func (n navOverflow) Enable(value bool) navOverflow {
	return n.set("enable", value)
}

// ItemWidth 设置导航项目宽度
func (n navOverflow) ItemWidth(value string) navOverflow {
	return n.set("itemWidth", value)
}

// MaxVisibleCount 设置最大可显示数量
func (n navOverflow) MaxVisibleCount(value string) navOverflow {
	return n.set("maxVisibleCount", value)
}

// OverflowClassName 设置菜单触发按钮 CSS 类名
func (n navOverflow) OverflowClassName(value string) navOverflow {
	return n.set("overflowClassName", value)
}

// OverflowIndicator 设置菜单触发按钮的图标
func (n navOverflow) OverflowIndicator(value string) navOverflow {
	return n.set("overflowIndicator", value)
}

// OverflowLabel 设置菜单触发按钮的文字
func (n navOverflow) OverflowLabel(value string) navOverflow {
	return n.set("overflowLabel", value)
}

// OverflowListClassName 设置菜单外层 CSS 类名
func (n navOverflow) OverflowListClassName(value string) navOverflow {
	return n.set("overflowListClassName", value)
}

// OverflowPopoverClassName 设置 Popover 浮层 CSS 类名
func (n navOverflow) OverflowPopoverClassName(value string) navOverflow {
	return n.set("overflowPopoverClassName", value)
}

// OverflowSuffix 设置导航列表后缀节点
func (n navOverflow) OverflowSuffix(value string) navOverflow {
	return n.set("overflowSuffix", value)
}

// Style 设置自定义样式
func (n navOverflow) Style(value any) navOverflow {
	return n.set("style", value)
}

// WrapperComponent 设置包裹导航的外层标签名
func (n navOverflow) WrapperComponent(value string) navOverflow {
	return n.set("wrapperComponent", value)
}
