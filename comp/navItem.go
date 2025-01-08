package comp

import "github.com/zrcoder/amisgo/model"

// navItem represents a navigation item
type navItem model.Schema

// NavItem creates a new navItem instance
func NavItem() navItem {
	return navItem{}
}

// set sets a property and returns the navItem
func (n navItem) set(key string, value any) navItem {
	n[key] = value
	return n
}

// Active sets the active state
func (n navItem) Active(value bool) navItem {
	return n.set("active", value)
}

// Children sets the children
func (n navItem) Children(value string) navItem {
	return n.set("children", value)
}

// ClassName sets the CSS class name
func (n navItem) ClassName(value string) navItem {
	return n.set("className", value)
}

// Defer sets the defer state
func (n navItem) Defer(value bool) navItem {
	return n.set("defer", value)
}

// DeferApi sets the defer API
func (n navItem) DeferApi(value string) navItem {
	return n.set("deferApi", value)
}

// Disabled sets the disabled state
func (n navItem) Disabled(value bool) navItem {
	return n.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (n navItem) DisabledOn(value string) navItem {
	return n.set("disabledOn", value)
}

// DisabledTip sets the disabled tip
func (n navItem) DisabledTip(value string) navItem {
	return n.set("disabledTip", value)
}

// EditorSetting sets the editor setting
func (n navItem) EditorSetting(value string) navItem {
	return n.set("editorSetting", value)
}

// Hidden sets the hidden state
func (n navItem) Hidden(value bool) navItem {
	return n.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (n navItem) HiddenOn(value string) navItem {
	return n.set("hiddenOn", value)
}

// Icon sets the icon class name
func (n navItem) Icon(value string) navItem {
	return n.set("icon", value)
}

// ID sets the unique ID
func (n navItem) ID(value string) navItem {
	return n.set("id", value)
}

// Key sets the unique key
func (n navItem) Key(value string) navItem {
	return n.set("key", value)
}

// Label sets the label
func (n navItem) Label(value string) navItem {
	return n.set("label", value)
}

// Mode sets the mode
func (n navItem) Mode(value string) navItem {
	return n.set("mode", value)
}

// OnEvent sets the event configuration
func (n navItem) OnEvent(value any) navItem {
	return n.set("onEvent", value)
}

// Static sets the static display state
func (n navItem) Static(value bool) navItem {
	return n.set("static", value)
}

// StaticClassName sets the static display class name
func (n navItem) StaticClassName(value string) navItem {
	return n.set("staticClassName", value)
}

// StaticInputClassName sets the static input class name
func (n navItem) StaticInputClassName(value string) navItem {
	return n.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label class name
func (n navItem) StaticLabelClassName(value string) navItem {
	return n.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (n navItem) StaticOn(value string) navItem {
	return n.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder
func (n navItem) StaticPlaceholder(value string) navItem {
	return n.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (n navItem) StaticSchema(value string) navItem {
	return n.set("staticSchema", value)
}

// Style sets the component style
func (n navItem) Style(value any) navItem {
	return n.set("style", value)
}

// Target sets the target
func (n navItem) Target(value string) navItem {
	return n.set("target", value)
}

// TestIdBuilder sets the test ID builder
func (n navItem) TestIdBuilder(value string) navItem {
	return n.set("testIdBuilder", value)
}

// Testid sets the test ID
func (n navItem) Testid(value string) navItem {
	return n.set("testid", value)
}

// To sets the navigation target
func (n navItem) To(value string) navItem {
	return n.set("to", value)
}

// Unfolded sets the unfolded state
func (n navItem) Unfolded(value bool) navItem {
	return n.set("unfolded", value)
}

// UseMobileUI sets the mobile UI state
func (n navItem) UseMobileUI(value bool) navItem {
	return n.set("useMobileUI", value)
}

// Visible sets the visibility
func (n navItem) Visible(value bool) navItem {
	return n.set("visible", value)
}

// VisibleOn sets the visibility expression
func (n navItem) VisibleOn(value string) navItem {
	return n.set("visibleOn", value)
}
