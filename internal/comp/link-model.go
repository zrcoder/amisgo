package comp

import "github.com/zrcoder/amisgo/schema"

type NavLink schema.Schema

func NewNavLink() NavLink {
	return NavLink{}
}

func (nl NavLink) set(key string, value any) NavLink {
	nl[key] = value
	return nl
}

// Label sets the link label.
func (nl NavLink) Label(value string) NavLink {
	return nl.set("label", value)
}

// To sets the link URL.
func (nl NavLink) To(value string) NavLink {
	return nl.set("to", value)
}

// Target sets the link target.
func (nl NavLink) Target(value string) NavLink {
	return nl.set("target", value)
}

// Icon sets the link icon.
func (nl NavLink) Icon(value string) NavLink {
	return nl.set("icon", value)
}

// Children sets the child links.
func (nl NavLink) Children(value ...NavLink) NavLink {
	return nl.set("children", value)
}

// Unfolded sets whether the link is initially unfolded.
func (nl NavLink) Unfolded(value ...NavLink) NavLink {
	return nl.set("unfolded", value)
}

// Active sets whether the link is active.
func (nl NavLink) Active(value bool) NavLink {
	return nl.set("active", value)
}

// ActiveOn sets the condition for the link to be active.
func (nl NavLink) ActiveOn(value string) NavLink {
	return nl.set("activeOn", value)
}

// Defer sets whether the link is lazy-loaded.
func (nl NavLink) Defer(value bool) NavLink {
	return nl.set("defer", value)
}

// DeferApi sets the API for lazy-loading.
func (nl NavLink) DeferApi(value any) NavLink {
	return nl.set("deferApi", value)
}

// Disabled sets whether the link is disabled.
func (nl NavLink) Disabled(value bool) NavLink {
	return nl.set("disabled", value)
}

// DisabledTip sets the tooltip text when the link is disabled.
func (nl NavLink) DisabledTip(value string) NavLink {
	return nl.set("disabledTip", value)
}

// ClassName sets the custom CSS class.
func (nl NavLink) ClassName(value string) NavLink {
	return nl.set("className", value)
}

// Mode sets the menu item mode.
func (nl NavLink) Mode(value string) NavLink {
	return nl.set("mode", value)
}

// Overflow sets the overflow property.
func (nl NavLink) Overflow(value any) NavLink {
	return nl.set("overflow", value)
}
