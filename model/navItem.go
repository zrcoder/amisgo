package model

// NavItem represents a navigation item
type NavItem Schema

func NewNavItem() NavItem {
	return NavItem{}
}

// set sets a property and returns the navItem
func (n NavItem) set(key string, value any) NavItem {
	n[key] = value
	return n
}

// Active sets the active state
func (n NavItem) Active(value bool) NavItem {
	return n.set("active", value)
}

// Children sets the children
func (n NavItem) Children(value string) NavItem {
	return n.set("children", value)
}

// ClassName sets the CSS class name
func (n NavItem) ClassName(value string) NavItem {
	return n.set("className", value)
}

// Defer sets the defer state
func (n NavItem) Defer(value bool) NavItem {
	return n.set("defer", value)
}

// DeferApi sets the defer API
func (n NavItem) DeferApi(value string) NavItem {
	return n.set("deferApi", value)
}

// Disabled sets the disabled state
func (n NavItem) Disabled(value bool) NavItem {
	return n.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (n NavItem) DisabledOn(value string) NavItem {
	return n.set("disabledOn", value)
}

// DisabledTip sets the disabled tip
func (n NavItem) DisabledTip(value string) NavItem {
	return n.set("disabledTip", value)
}

// EditorSetting sets the editor setting
func (n NavItem) EditorSetting(value string) NavItem {
	return n.set("editorSetting", value)
}

// Hidden sets the hidden state
func (n NavItem) Hidden(value bool) NavItem {
	return n.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (n NavItem) HiddenOn(value string) NavItem {
	return n.set("hiddenOn", value)
}

// Icon sets the icon class name
func (n NavItem) Icon(value string) NavItem {
	return n.set("icon", value)
}

// ID sets the unique ID
func (n NavItem) ID(value string) NavItem {
	return n.set("id", value)
}

// Key sets the unique key
func (n NavItem) Key(value string) NavItem {
	return n.set("key", value)
}

// Label sets the label
func (n NavItem) Label(value string) NavItem {
	return n.set("label", value)
}

// Mode sets the mode
func (n NavItem) Mode(value string) NavItem {
	return n.set("mode", value)
}

// OnEvent sets the event configuration
func (n NavItem) OnEvent(value any) NavItem {
	return n.set("onEvent", value)
}

// Static sets the static display state
func (n NavItem) Static(value bool) NavItem {
	return n.set("static", value)
}

// StaticClassName sets the static display class name
func (n NavItem) StaticClassName(value string) NavItem {
	return n.set("staticClassName", value)
}

// StaticInputClassName sets the static input class name
func (n NavItem) StaticInputClassName(value string) NavItem {
	return n.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label class name
func (n NavItem) StaticLabelClassName(value string) NavItem {
	return n.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (n NavItem) StaticOn(value string) NavItem {
	return n.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder
func (n NavItem) StaticPlaceholder(value string) NavItem {
	return n.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (n NavItem) StaticSchema(value string) NavItem {
	return n.set("staticSchema", value)
}

// Style sets the component style
func (n NavItem) Style(value any) NavItem {
	return n.set("style", value)
}

// Target sets the target
func (n NavItem) Target(value string) NavItem {
	return n.set("target", value)
}

// TestIdBuilder sets the test ID builder
func (n NavItem) TestIdBuilder(value string) NavItem {
	return n.set("testIdBuilder", value)
}

// Testid sets the test ID
func (n NavItem) Testid(value string) NavItem {
	return n.set("testid", value)
}

// To sets the navigation target
func (n NavItem) To(value string) NavItem {
	return n.set("to", value)
}

// Unfolded sets the unfolded state
func (n NavItem) Unfolded(value bool) NavItem {
	return n.set("unfolded", value)
}

// UseMobileUI sets the mobile UI state
func (n NavItem) UseMobileUI(value bool) NavItem {
	return n.set("useMobileUI", value)
}

// Visible sets the visibility
func (n NavItem) Visible(value bool) NavItem {
	return n.set("visible", value)
}

// VisibleOn sets the visibility expression
func (n NavItem) VisibleOn(value string) NavItem {
	return n.set("visibleOn", value)
}
