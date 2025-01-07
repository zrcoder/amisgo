package comp

type MNavLink Schema

func NavLink() MNavLink {
	return MNavLink{}
}

func (nl MNavLink) set(key string, value any) MNavLink {
	nl[key] = value
	return nl
}

// Label sets the link label.
func (nl MNavLink) Label(value string) MNavLink {
	return nl.set("label", value)
}

// To sets the link URL.
func (nl MNavLink) To(value string) MNavLink {
	return nl.set("to", value)
}

// Target sets the link target.
func (nl MNavLink) Target(value string) MNavLink {
	return nl.set("target", value)
}

// Icon sets the link icon.
func (nl MNavLink) Icon(value string) MNavLink {
	return nl.set("icon", value)
}

// Children sets the child links.
func (nl MNavLink) Children(value ...MNavLink) MNavLink {
	return nl.set("children", value)
}

// Unfolded sets whether the link is initially unfolded.
func (nl MNavLink) Unfolded(value ...MNavLink) MNavLink {
	return nl.set("unfolded", value)
}

// Active sets whether the link is active.
func (nl MNavLink) Active(value bool) MNavLink {
	return nl.set("active", value)
}

// ActiveOn sets the condition for the link to be active.
func (nl MNavLink) ActiveOn(value string) MNavLink {
	return nl.set("activeOn", value)
}

// Defer sets whether the link is lazy-loaded.
func (nl MNavLink) Defer(value bool) MNavLink {
	return nl.set("defer", value)
}

// DeferApi sets the API for lazy-loading.
func (nl MNavLink) DeferApi(value any) MNavLink {
	return nl.set("deferApi", value)
}

// Disabled sets whether the link is disabled.
func (nl MNavLink) Disabled(value bool) MNavLink {
	return nl.set("disabled", value)
}

// DisabledTip sets the tooltip text when the link is disabled.
func (nl MNavLink) DisabledTip(value string) MNavLink {
	return nl.set("disabledTip", value)
}

// ClassName sets the custom CSS class.
func (nl MNavLink) ClassName(value string) MNavLink {
	return nl.set("className", value)
}

// Mode sets the menu item mode.
func (nl MNavLink) Mode(value string) MNavLink {
	return nl.set("mode", value)
}

// Overflow sets the overflow property.
func (nl MNavLink) Overflow(value any) MNavLink {
	return nl.set("overflow", value)
}
