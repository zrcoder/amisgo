package comp

// portlet

type portlet Schema

// Portlet creates a new Portlet instance
func Portlet() portlet {
	return portlet{}.set("type", "portlet")
}

func (p portlet) set(key string, value any) portlet {
	p[key] = value
	return p
}

// ClassName sets the container CSS class name
func (p portlet) ClassName(value string) portlet {
	return p.set("className", value)
}

// ContentClassName sets the content class name
func (p portlet) ContentClassName(value string) portlet {
	return p.set("contentClassName", value)
}

// Description sets the description on the right side of the title
func (p portlet) Description(value string) portlet {
	return p.set("description", value)
}

// Disabled sets whether the portlet is disabled
func (p portlet) Disabled(value bool) portlet {
	return p.set("disabled", value)
}

// DisabledOn sets the expression to determine if the portlet is disabled
func (p portlet) DisabledOn(value string) portlet {
	return p.set("disabledOn", value)
}

// Divider sets whether to show a divider between the header and content
func (p portlet) Divider(value bool) portlet {
	return p.set("divider", value)
}

// EditorSetting sets the editor configuration
func (p portlet) EditorSetting(value string) portlet {
	return p.set("editorSetting", value)
}

// Hidden sets whether the portlet is hidden
func (p portlet) Hidden(value bool) portlet {
	return p.set("hidden", value)
}

// HiddenOn sets the expression to determine if the portlet is hidden
func (p portlet) HiddenOn(value string) portlet {
	return p.set("hiddenOn", value)
}

// HideHeader sets whether to hide the header
func (p portlet) HideHeader(value bool) portlet {
	return p.set("hideHeader", value)
}

// ID sets the unique component ID
func (p portlet) ID(value string) portlet {
	return p.set("id", value)
}

// LinksClassName sets the class name for the outer links
func (p portlet) LinksClassName(value string) portlet {
	return p.set("linksClassName", value)
}

// MountOnEnter sets whether to load the card only when opened
func (p portlet) MountOnEnter(value bool) portlet {
	return p.set("mountOnEnter", value)
}

// OnEvent sets the event action configuration
func (p portlet) OnEvent(value any) portlet {
	return p.set("onEvent", value)
}

// Scrollable sets whether to support overflow scrolling
func (p portlet) Scrollable(value bool) portlet {
	return p.set("scrollable", value)
}

// Source sets the associated existing data
func (p portlet) Source(value string) portlet {
	return p.set("source", value)
}

// Static sets whether to display statically
func (p portlet) Static(value bool) portlet {
	return p.set("static", value)
}

// StaticClassName sets the class name for static display form items
func (p portlet) StaticClassName(value string) portlet {
	return p.set("staticClassName", value)
}

// StaticInputClassName sets the class name for static display form item values
func (p portlet) StaticInputClassName(value string) portlet {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class name for static display form item labels
func (p portlet) StaticLabelClassName(value string) portlet {
	return p.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the portlet is displayed statically
func (p portlet) StaticOn(value string) portlet {
	return p.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display empty values
func (p portlet) StaticPlaceholder(value string) portlet {
	return p.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema
func (p portlet) StaticSchema(value string) portlet {
	return p.set("staticSchema", value)
}

// Style sets custom styles
func (p portlet) Style(value any) portlet {
	return p.set("style", value)
}

// Tabs sets the tabs
func (p portlet) Tabs(value string) portlet {
	return p.set("tabs", value)
}

// TabsClassName sets the class name for the tabs
func (p portlet) TabsClassName(value string) portlet {
	return p.set("tabsClassName", value)
}

// TabsMode sets the display mode for the tabs
func (p portlet) TabsMode(value string) portlet {
	return p.set("tabsMode", value)
}

// TestIdBuilder sets the test ID builder
func (p portlet) TestIdBuilder(value string) portlet {
	return p.set("testIdBuilder", value)
}

// Testid sets the test ID
func (p portlet) Testid(value string) portlet {
	return p.set("testid", value)
}

// Toolbar sets the toolbar for additional buttons on the right side
func (p portlet) Toolbar(value string) portlet {
	return p.set("toolbar", value)
}

// UnmountOnExit sets whether to destroy the card content when hidden
func (p portlet) UnmountOnExit(value bool) portlet {
	return p.set("unmountOnExit", value)
}

// UseMobileUI sets whether to disable mobile UI styles at the component level
func (p portlet) UseMobileUI(value bool) portlet {
	return p.set("useMobileUI", value)
}

// Visible sets whether the portlet is visible
func (p portlet) Visible(value bool) portlet {
	return p.set("visible", value)
}

// VisibleOn sets the expression to determine if the portlet is visible
func (p portlet) VisibleOn(value string) portlet {
	return p.set("visibleOn", value)
}
