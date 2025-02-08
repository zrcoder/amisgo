package model

type PortletTab Schema

func NewPortletTab() PortletTab {
	return PortletTab{}
}

func (p PortletTab) set(key string, value any) PortletTab {
	p[key] = value
	return p
}

// Body sets the body content
func (p PortletTab) Body(value ...any) PortletTab {
	return p.set("body", value)
}

// ClassName sets the container CSS class name
func (p PortletTab) ClassName(value string) PortletTab {
	return p.set("className", value)
}

// Disabled sets the disabled state
func (p PortletTab) Disabled(value bool) PortletTab {
	return p.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (p PortletTab) DisabledOn(value string) PortletTab {
	return p.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (p PortletTab) EditorSetting(value string) PortletTab {
	return p.set("editorSetting", value)
}

// Hidden sets the hidden state
func (p PortletTab) Hidden(value bool) PortletTab {
	return p.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (p PortletTab) HiddenOn(value string) PortletTab {
	return p.set("hiddenOn", value)
}

// Icon sets the button icon
func (p PortletTab) Icon(value string) PortletTab {
	return p.set("icon", value)
}

// IconPosition sets the icon position
func (p PortletTab) IconPosition(value string) PortletTab {
	return p.set("iconPosition", value)
}

// ID sets the unique component ID
func (p PortletTab) ID(value string) PortletTab {
	return p.set("id", value)
}

// MountOnEnter sets whether to load content on tab click
func (p PortletTab) MountOnEnter(value bool) PortletTab {
	return p.set("mountOnEnter", value)
}

// OnEvent sets the event action configuration
func (p PortletTab) OnEvent(value any) PortletTab {
	return p.set("onEvent", value)
}

// Reload sets whether to re-render content each time
func (p PortletTab) Reload(value bool) PortletTab {
	return p.set("reload", value)
}

// Static sets the static display state
func (p PortletTab) Static(value bool) PortletTab {
	return p.set("static", value)
}

// StaticClassName sets the static display form item class name
func (p PortletTab) StaticClassName(value string) PortletTab {
	return p.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (p PortletTab) StaticInputClassName(value string) PortletTab {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (p PortletTab) StaticLabelClassName(value string) PortletTab {
	return p.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (p PortletTab) StaticOn(value string) PortletTab {
	return p.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder
func (p PortletTab) StaticPlaceholder(value string) PortletTab {
	return p.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema
func (p PortletTab) StaticSchema(value string) PortletTab {
	return p.set("staticSchema", value)
}

// Style sets the component style
func (p PortletTab) Style(value any) PortletTab {
	return p.set("style", value)
}

// Tab sets the tab content
func (p PortletTab) Tab(value string) PortletTab {
	return p.set("tab", value)
}

// TestIdBuilder sets the test ID builder
func (p PortletTab) TestIdBuilder(value string) PortletTab {
	return p.set("testIdBuilder", value)
}

// Testid sets the test ID
func (p PortletTab) Testid(value string) PortletTab {
	return p.set("testid", value)
}

// Title sets the tab title
func (p PortletTab) Title(value any) PortletTab {
	return p.set("title", value)
}

// Toolbar sets the toolbar configuration
func (p PortletTab) Toolbar(value string) PortletTab {
	return p.set("toolbar", value)
}

// UnmountOnExit sets whether to destroy the tab node on hide
func (p PortletTab) UnmountOnExit(value bool) PortletTab {
	return p.set("unmountOnExit", value)
}

// UseMobileUI sets whether to disable mobile styles
func (p PortletTab) UseMobileUI(value bool) PortletTab {
	return p.set("useMobileUI", value)
}

// Visible sets the visibility state
func (p PortletTab) Visible(value bool) PortletTab {
	return p.set("visible", value)
}

// VisibleOn sets the visibility expression
func (p PortletTab) VisibleOn(value string) PortletTab {
	return p.set("visibleOn", value)
}
