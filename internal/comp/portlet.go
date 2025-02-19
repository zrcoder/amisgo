package comp

import "github.com/zrcoder/amisgo/schema"

// Portlet
type Portlet schema.Schema

func NewPortlet() Portlet {
	return Portlet{"type": "portlet"}
}

func (p Portlet) set(key string, value any) Portlet {
	p[key] = value
	return p
}

// ClassName sets the container CSS class name
func (p Portlet) ClassName(value string) Portlet {
	return p.set("className", value)
}

// ContentClassName sets the content class name
func (p Portlet) ContentClassName(value string) Portlet {
	return p.set("contentClassName", value)
}

// Description sets the description on the right side of the title
func (p Portlet) Description(value string) Portlet {
	return p.set("description", value)
}

// Disabled sets whether the portlet is disabled
func (p Portlet) Disabled(value bool) Portlet {
	return p.set("disabled", value)
}

// DisabledOn sets the expression to determine if the portlet is disabled
func (p Portlet) DisabledOn(value string) Portlet {
	return p.set("disabledOn", value)
}

// Divider sets whether to show a divider between the header and content
func (p Portlet) Divider(value bool) Portlet {
	return p.set("divider", value)
}

// EditorSetting sets the editor configuration
func (p Portlet) EditorSetting(value string) Portlet {
	return p.set("editorSetting", value)
}

// Hidden sets whether the portlet is hidden
func (p Portlet) Hidden(value bool) Portlet {
	return p.set("hidden", value)
}

// HiddenOn sets the expression to determine if the portlet is hidden
func (p Portlet) HiddenOn(value string) Portlet {
	return p.set("hiddenOn", value)
}

// HideHeader sets whether to hide the header
func (p Portlet) HideHeader(value bool) Portlet {
	return p.set("hideHeader", value)
}

// ID sets the unique component ID
func (p Portlet) ID(value string) Portlet {
	return p.set("id", value)
}

// LinksClassName sets the class name for the outer links
func (p Portlet) LinksClassName(value string) Portlet {
	return p.set("linksClassName", value)
}

// MountOnEnter sets whether to load the card only when opened
func (p Portlet) MountOnEnter(value bool) Portlet {
	return p.set("mountOnEnter", value)
}

// OnEvent sets the event action configuration
func (p Portlet) OnEvent(value any) Portlet {
	return p.set("onEvent", value)
}

// Scrollable sets whether to support overflow scrolling
func (p Portlet) Scrollable(value bool) Portlet {
	return p.set("scrollable", value)
}

// Source sets the associated existing data
func (p Portlet) Source(value string) Portlet {
	return p.set("source", value)
}

// Static sets whether to display statically
func (p Portlet) Static(value bool) Portlet {
	return p.set("static", value)
}

// StaticClassName sets the class name for static display form items
func (p Portlet) StaticClassName(value string) Portlet {
	return p.set("staticClassName", value)
}

// StaticInputClassName sets the class name for static display form item values
func (p Portlet) StaticInputClassName(value string) Portlet {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class name for static display form item labels
func (p Portlet) StaticLabelClassName(value string) Portlet {
	return p.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the portlet is displayed statically
func (p Portlet) StaticOn(value string) Portlet {
	return p.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display empty values
func (p Portlet) StaticPlaceholder(value string) Portlet {
	return p.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema.Schema
func (p Portlet) StaticSchema(value string) Portlet {
	return p.set("staticSchema", value)
}

// Style sets custom styles
func (p Portlet) Style(value any) Portlet {
	return p.set("style", value)
}

// Tabs sets the tabs
func (p Portlet) Tabs(value string) Portlet {
	return p.set("tabs", value)
}

// TabsClassName sets the class name for the tabs
func (p Portlet) TabsClassName(value string) Portlet {
	return p.set("tabsClassName", value)
}

// TabsMode sets the display mode for the tabs
func (p Portlet) TabsMode(value string) Portlet {
	return p.set("tabsMode", value)
}

// TestIdBuilder sets the test ID builder
func (p Portlet) TestIdBuilder(value string) Portlet {
	return p.set("testIdBuilder", value)
}

// Testid sets the test ID
func (p Portlet) Testid(value string) Portlet {
	return p.set("testid", value)
}

// Toolbar sets the toolbar for additional buttons on the right side
func (p Portlet) Toolbar(value string) Portlet {
	return p.set("toolbar", value)
}

// UnmountOnExit sets whether to destroy the card content when hidden
func (p Portlet) UnmountOnExit(value bool) Portlet {
	return p.set("unmountOnExit", value)
}

// UseMobileUI sets whether to disable mobile UI styles at the component level
func (p Portlet) UseMobileUI(value bool) Portlet {
	return p.set("useMobileUI", value)
}

// Visible sets whether the portlet is visible
func (p Portlet) Visible(value bool) Portlet {
	return p.set("visible", value)
}

// VisibleOn sets the expression to determine if the portlet is visible
func (p Portlet) VisibleOn(value string) Portlet {
	return p.set("visibleOn", value)
}
