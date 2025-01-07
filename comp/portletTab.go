package comp

// portletTab

type portletTab Schema

// PortletTab creates a new PortletTab instance
func PortletTab() portletTab {
	return portletTab{}
}

func (p portletTab) set(key string, value any) portletTab {
	p[key] = value
	return p
}

// Body sets the body content
func (p portletTab) Body(value ...any) portletTab {
	return p.set("body", value)
}

// ClassName sets the container CSS class name
func (p portletTab) ClassName(value string) portletTab {
	return p.set("className", value)
}

// Disabled sets the disabled state
func (p portletTab) Disabled(value bool) portletTab {
	return p.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (p portletTab) DisabledOn(value string) portletTab {
	return p.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (p portletTab) EditorSetting(value string) portletTab {
	return p.set("editorSetting", value)
}

// Hidden sets the hidden state
func (p portletTab) Hidden(value bool) portletTab {
	return p.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (p portletTab) HiddenOn(value string) portletTab {
	return p.set("hiddenOn", value)
}

// Icon sets the button icon
func (p portletTab) Icon(value string) portletTab {
	return p.set("icon", value)
}

// IconPosition sets the icon position
func (p portletTab) IconPosition(value string) portletTab {
	return p.set("iconPosition", value)
}

// ID sets the unique component ID
func (p portletTab) ID(value string) portletTab {
	return p.set("id", value)
}

// MountOnEnter sets whether to load content on tab click
func (p portletTab) MountOnEnter(value bool) portletTab {
	return p.set("mountOnEnter", value)
}

// OnEvent sets the event action configuration
func (p portletTab) OnEvent(value any) portletTab {
	return p.set("onEvent", value)
}

// Reload sets whether to re-render content each time
func (p portletTab) Reload(value bool) portletTab {
	return p.set("reload", value)
}

// Static sets the static display state
func (p portletTab) Static(value bool) portletTab {
	return p.set("static", value)
}

// StaticClassName sets the static display form item class name
func (p portletTab) StaticClassName(value string) portletTab {
	return p.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (p portletTab) StaticInputClassName(value string) portletTab {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (p portletTab) StaticLabelClassName(value string) portletTab {
	return p.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (p portletTab) StaticOn(value string) portletTab {
	return p.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder
func (p portletTab) StaticPlaceholder(value string) portletTab {
	return p.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema
func (p portletTab) StaticSchema(value string) portletTab {
	return p.set("staticSchema", value)
}

// Style sets the component style
func (p portletTab) Style(value any) portletTab {
	return p.set("style", value)
}

// Tab sets the tab content
func (p portletTab) Tab(value string) portletTab {
	return p.set("tab", value)
}

// TestIdBuilder sets the test ID builder
func (p portletTab) TestIdBuilder(value string) portletTab {
	return p.set("testIdBuilder", value)
}

// Testid sets the test ID
func (p portletTab) Testid(value string) portletTab {
	return p.set("testid", value)
}

// Title sets the tab title
func (p portletTab) Title(value any) portletTab {
	return p.set("title", value)
}

// Toolbar sets the toolbar configuration
func (p portletTab) Toolbar(value string) portletTab {
	return p.set("toolbar", value)
}

// UnmountOnExit sets whether to destroy the tab node on hide
func (p portletTab) UnmountOnExit(value bool) portletTab {
	return p.set("unmountOnExit", value)
}

// UseMobileUI sets whether to disable mobile styles
func (p portletTab) UseMobileUI(value bool) portletTab {
	return p.set("useMobileUI", value)
}

// Visible sets the visibility state
func (p portletTab) Visible(value bool) portletTab {
	return p.set("visible", value)
}

// VisibleOn sets the visibility expression
func (p portletTab) VisibleOn(value string) portletTab {
	return p.set("visibleOn", value)
}
