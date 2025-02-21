package comp

import "github.com/zrcoder/amisgo/schema"

// Tabs component. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/Tabs
type Tabs schema.Schema

func NewTabs() Tabs {
	return Tabs{"type": "tabs"}
}

func (t Tabs) set(key string, value any) Tabs {
	t[key] = value
	return t
}

// ActiveKey sets the active tab by hash or index, supports expressions
func (t Tabs) ActiveKey(value any) Tabs {
	return t.set("activeKey", value)
}

// AddBtnText sets custom text for the add button
func (t Tabs) AddBtnText(value string) Tabs {
	return t.set("addBtnText", value)
}

// Addable enables or disables adding new tabs
func (t Tabs) Addable(value bool) Tabs {
	return t.set("addable", value)
}

// ClassName sets the CSS class name for the container
func (t Tabs) ClassName(value string) Tabs {
	return t.set("className", value)
}

// Closable enables or disables tab deletion
func (t Tabs) Closable(value bool) Tabs {
	return t.set("closable", value)
}

// CollapseBtnLabel sets the text for the collapse button
func (t Tabs) CollapseBtnLabel(value string) Tabs {
	return t.set("collapseBtnLabel", value)
}

// CollapseOnExceed sets the number of tabs before collapsing
func (t Tabs) CollapseOnExceed(value string) Tabs {
	return t.set("collapseOnExceed", value)
}

// ContentClassName sets the CSS class name for the content
func (t Tabs) ContentClassName(value string) Tabs {
	return t.set("contentClassName", value)
}

// DefaultKey sets the default active tab by hash or index, supports expressions
func (t Tabs) DefaultKey(value any) Tabs {
	return t.set("defaultKey", value)
}

// Disabled enables or disables the tabs
func (t Tabs) Disabled(value bool) Tabs {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to disable the tabs
func (t Tabs) DisabledOn(value string) Tabs {
	return t.set("disabledOn", value)
}

// Draggable enables or disables tab dragging
func (t Tabs) Draggable(value bool) Tabs {
	return t.set("draggable", value)
}

// Editable enables or disables tab label editing
func (t Tabs) Editable(value bool) Tabs {
	return t.set("editable", value)
}

// EditorSetting sets the editor configuration
func (t Tabs) EditorSetting(value string) Tabs {
	return t.set("editorSetting", value)
}

// Hidden hides or shows the tabs
func (t Tabs) Hidden(value bool) Tabs {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to hide the tabs
func (t Tabs) HiddenOn(value string) Tabs {
	return t.set("hiddenOn", value)
}

// ID sets the unique ID for the component
func (t Tabs) ID(value string) Tabs {
	return t.set("id", value)
}

// LinksClassName sets the CSS class name for the links container
func (t Tabs) LinksClassName(value string) Tabs {
	return t.set("linksClassName", value)
}

// MountOnEnter loads the tab content only when activated
func (t Tabs) MountOnEnter(value bool) Tabs {
	return t.set("mountOnEnter", value)
}

// OnEvent sets the event action configuration
func (t Tabs) OnEvent(value Event) Tabs {
	return t.set("onEvent", value)
}

// Scrollable enables or disables scrolling for overflow content
func (t Tabs) Scrollable(value bool) Tabs {
	return t.set("scrollable", value)
}

// ShowTip enables or disables tooltips
func (t Tabs) ShowTip(value bool) Tabs {
	return t.set("showTip", value)
}

// ShowTipClassName sets the CSS class name for tooltips
func (t Tabs) ShowTipClassName(value string) Tabs {
	return t.set("showTipClassName", value)
}

// SidePosition sets the side position for the editor mode (left/right)
func (t Tabs) SidePosition(value string) Tabs {
	return t.set("sidePosition", value)
}

// Source sets the data source
func (t Tabs) Source(value string) Tabs {
	return t.set("source", value)
}

// Static enables or disables static display
func (t Tabs) Static(value bool) Tabs {
	return t.set("static", value)
}

// StaticClassName sets the CSS class name for static display items
func (t Tabs) StaticClassName(value string) Tabs {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static display item values
func (t Tabs) StaticInputClassName(value string) Tabs {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static display item labels
func (t Tabs) StaticLabelClassName(value string) Tabs {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (t Tabs) StaticOn(value string) Tabs {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display empty values
func (t Tabs) StaticPlaceholder(value string) Tabs {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (t Tabs) StaticSchema(value string) Tabs {
	return t.set("staticSchema", value)
}

// Style sets the component style
func (t Tabs) Style(value any) Tabs {
	return t.set("style", value)
}

// SubFormHorizontal sets the width ratio for horizontal layout
func (t Tabs) SubFormHorizontal(value string) Tabs {
	return t.set("subFormHorizontal", value)
}

// SubFormMode sets the default display mode for subform items
func (t Tabs) SubFormMode(value string) Tabs {
	return t.set("subFormMode", value)
}

// Swipeable enables or disables swipe switching (mobile only)
func (t Tabs) Swipeable(value bool) Tabs {
	return t.set("swipeable", value)
}

// Tabs sets the tab members
func (t Tabs) Tabs(value ...any) Tabs {
	return t.set("tabs", value)
}

// TabsMode sets the display mode (line, card, radio, vertical, chrome, simple, strong, tiled, sidebar)
func (t Tabs) TabsMode(value string) Tabs {
	return t.set("tabsMode", value)
}

// TestIdBuilder sets the test ID builder
func (t Tabs) TestIdBuilder(value string) Tabs {
	return t.set("testIdBuilder", value)
}

// Testid sets the test ID
func (t Tabs) Testid(value string) Tabs {
	return t.set("testid", value)
}

// Toolbar sets the toolbar configuration
func (t Tabs) Toolbar(value string) Tabs {
	return t.set("toolbar", value)
}

// UnmountOnExit destroys the tab content when hidden
func (t Tabs) UnmountOnExit(value bool) Tabs {
	return t.set("unmountOnExit", value)
}

// UseMobileUI enables or disables mobile UI styles
func (t Tabs) UseMobileUI(value bool) Tabs {
	return t.set("useMobileUI", value)
}

// Visible shows or hides the tabs
func (t Tabs) Visible(value bool) Tabs {
	return t.set("visible", value)
}

// VisibleOn sets the expression to show the tabs
func (t Tabs) VisibleOn(value string) Tabs {
	return t.set("visibleOn", value)
}
