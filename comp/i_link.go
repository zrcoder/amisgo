package comp

type navLink schema

func NavLink() navLink {
	return navLink{}
}

func (nl navLink) set(key string, value any) navLink {
	nl[key] = value
	return nl
}

// Label 名称
func (nl navLink) Label(value string) navLink {
	return nl.set("label", value)
}

// To 链接地址
func (nl navLink) To(value string) navLink {
	return nl.set("to", value)
}

// Target 链接关系
func (nl navLink) Target(value string) navLink {
	return nl.set("target", value)
}

// Icon 图标
func (nl navLink) Icon(value string) navLink {
	return nl.set("icon", value)
}

// Children 子链接
func (nl navLink) Children(value ...navLink) navLink {
	return nl.set("children", value)
}

// Unfolded 初始是否展开
func (nl navLink) Unfolded(value ...navLink) navLink {
	return nl.set("unfolded", value)
}

// Active 是否激活
func (nl navLink) Active(value bool) navLink {
	return nl.set("active", value)
}

// ActiveOn 是否激活的条件，留空将自动分析链接地址
func (nl navLink) ActiveOn(value string) navLink {
	return nl.set("activeOn", value)
}

// Defer 标记是否为懒加载项
func (nl navLink) Defer(value bool) navLink {
	return nl.set("defer", value)
}

// DeferApi 可以不配置，如果配置优先级更高
func (nl navLink) DeferApi(value any) navLink {
	return nl.set("deferApi", value)
}

// Diaabled 是否禁用
func (nl navLink) Disabled(value bool) navLink {
	return nl.set("disabled", value)
}

// DisabledTip 禁用提示信息
func (nl navLink) DisabledTip(value string) navLink {
	return nl.set("disabledTip", value)
}

// ClassName 自定义样式
func (nl navLink) ClassName(value string) navLink {
	return nl.set("className", value)
}

// Mode 菜菜单项模式 "" | "group" | "divider"
func (nl navLink) Mode(value string) navLink {
	return nl.set("mode", value)
}

func (nl navLink) Overflow(value any) navLink {
	return nl.set("overflow", value)
}
