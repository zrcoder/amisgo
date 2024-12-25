package comp

type MNavLink schema

func NavLink() MNavLink {
	return MNavLink{}
}

func (nl MNavLink) set(key string, value any) MNavLink {
	nl[key] = value
	return nl
}

// Label 名称
func (nl MNavLink) Label(value string) MNavLink {
	return nl.set("label", value)
}

// To 链接地址
func (nl MNavLink) To(value string) MNavLink {
	return nl.set("to", value)
}

// Target 链接关系
func (nl MNavLink) Target(value string) MNavLink {
	return nl.set("target", value)
}

// Icon 图标
func (nl MNavLink) Icon(value string) MNavLink {
	return nl.set("icon", value)
}

// Children 子链接
func (nl MNavLink) Children(value ...MNavLink) MNavLink {
	return nl.set("children", value)
}

// Unfolded 初始是否展开
func (nl MNavLink) Unfolded(value ...MNavLink) MNavLink {
	return nl.set("unfolded", value)
}

// Active 是否激活
func (nl MNavLink) Active(value bool) MNavLink {
	return nl.set("active", value)
}

// ActiveOn 是否激活的条件，留空将自动分析链接地址
func (nl MNavLink) ActiveOn(value string) MNavLink {
	return nl.set("activeOn", value)
}

// Defer 标记是否为懒加载项
func (nl MNavLink) Defer(value bool) MNavLink {
	return nl.set("defer", value)
}

// DeferApi 可以不配置，如果配置优先级更高
func (nl MNavLink) DeferApi(value any) MNavLink {
	return nl.set("deferApi", value)
}

// Diaabled 是否禁用
func (nl MNavLink) Disabled(value bool) MNavLink {
	return nl.set("disabled", value)
}

// DisabledTip 禁用提示信息
func (nl MNavLink) DisabledTip(value string) MNavLink {
	return nl.set("disabledTip", value)
}

// ClassName 自定义样式
func (nl MNavLink) ClassName(value string) MNavLink {
	return nl.set("className", value)
}

// Mode 菜菜单项模式 "" | "group" | "divider"
func (nl MNavLink) Mode(value string) MNavLink {
	return nl.set("mode", value)
}

func (nl MNavLink) Overflow(value any) MNavLink {
	return nl.set("overflow", value)
}
