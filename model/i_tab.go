package model

type Tab Schema

func NewTab() Tab {
	return Tab{}
}

func (t Tab) set(key string, value any) Tab {
	t[key] = value
	return t
}

// Badge sets the badge
func (t Tab) Badge(value string) Tab {
	return t.set("badge", value)
}

// Body sets the content
func (t Tab) Body(value ...any) Tab {
	return t.set("body", value)
}

// ClassName sets the container CSS class name
func (t Tab) ClassName(value string) Tab {
	return t.set("className", value)
}

// Closable sets whether the tab is closable
func (t Tab) Closable(value bool) Tab {
	return t.set("closable", value)
}

// Disabled sets whether the tab is disabled
func (t Tab) Disabled(value bool) Tab {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to disable the tab
func (t Tab) DisabledOn(value string) Tab {
	return t.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (t Tab) EditorSetting(value string) Tab {
	return t.set("editorSetting", value)
}

// Hash sets the URL hash
func (t Tab) Hash(value string) Tab {
	return t.set("hash", value)
}

// Hidden sets whether the tab is hidden
func (t Tab) Hidden(value bool) Tab {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to hide the tab
func (t Tab) HiddenOn(value string) Tab {
	return t.set("hiddenOn", value)
}

// Horizontal sets the horizontal layout ratio
func (t Tab) Horizontal(value string) Tab {
	return t.set("horizontal", value)
}

// Icon sets the button icon
func (t Tab) Icon(value string) Tab {
	return t.set("icon", value)
}

// IconPosition sets the icon position (left or right)
func (t Tab) IconPosition(value string) Tab {
	return t.set("iconPosition", value)
}

// ID sets the unique component ID
func (t Tab) ID(value string) Tab {
	return t.set("id", value)
}

// Mode sets the default display mode (normal, inline, horizontal)
func (t Tab) Mode(value string) Tab {
	return t.set("mode", value)
}

// MountOnEnter sets whether to load content on tab activation
func (t Tab) MountOnEnter(value bool) Tab {
	return t.set("mountOnEnter", value)
}

// OnEvent sets the event action configuration
func (t Tab) OnEvent(value any) Tab {
	return t.set("onEvent", value)
}

// Reload sets whether to reload content on each render
func (t Tab) Reload(value bool) Tab {
	return t.set("reload", value)
}

// Static sets whether the tab is static
func (t Tab) Static(value bool) Tab {
	return t.set("static", value)
}

// StaticClassName sets the static form item class name
func (t Tab) StaticClassName(value string) Tab {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the static form item value class name
func (t Tab) StaticInputClassName(value string) Tab {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static form item label class name
func (t Tab) StaticLabelClassName(value string) Tab {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (t Tab) StaticOn(value string) Tab {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (t Tab) StaticPlaceholder(value string) Tab {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (t Tab) StaticSchema(value string) Tab {
	return t.set("staticSchema", value)
}

// Style sets the component style
func (t Tab) Style(value string) Tab {
	return t.set("style", value)
}

// Tab sets the tab content
func (t Tab) Tab(value ...any) Tab {
	return t.set("tab", value)
}

// TestIdBuilder sets the test ID builder
func (t Tab) TestIdBuilder(value string) Tab {
	return t.set("testIdBuilder", value)
}

// Testid sets the test ID
func (t Tab) Testid(value string) Tab {
	return t.set("testid", value)
}

// Title sets the tab title
func (t Tab) Title(value any) Tab {
	return t.set("title", value)
}

// UnmountOnExit sets whether to destroy the tab node on hide
func (t Tab) UnmountOnExit(value bool) Tab {
	return t.set("unmountOnExit", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (t Tab) UseMobileUI(value bool) Tab {
	return t.set("useMobileUI", value)
}

// Visible sets whether the tab is visible
func (t Tab) Visible(value bool) Tab {
	return t.set("visible", value)
}

// VisibleOn sets the expression to show the tab
func (t Tab) VisibleOn(value string) Tab {
	return t.set("visibleOn", value)
}
