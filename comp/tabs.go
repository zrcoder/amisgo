package comp

import "github.com/zrcoder/amisgo/model"

// Tabs component. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/tabs

type tabs model.Schema

// Tabs creates a new Tabs instance
func Tabs() tabs {
	return tabs{}.set("type", "tabs")
}

func (t tabs) set(key string, value any) tabs {
	t[key] = value
	return t
}

// ActiveKey sets the active tab by hash or index, supports expressions
func (t tabs) ActiveKey(value any) tabs {
	return t.set("activeKey", value)
}

// AddBtnText sets custom text for the add button
func (t tabs) AddBtnText(value string) tabs {
	return t.set("addBtnText", value)
}

// Addable enables or disables adding new tabs
func (t tabs) Addable(value bool) tabs {
	return t.set("addable", value)
}

// ClassName sets the CSS class name for the container
func (t tabs) ClassName(value string) tabs {
	return t.set("className", value)
}

// Closable enables or disables tab deletion
func (t tabs) Closable(value bool) tabs {
	return t.set("closable", value)
}

// CollapseBtnLabel sets the text for the collapse button
func (t tabs) CollapseBtnLabel(value string) tabs {
	return t.set("collapseBtnLabel", value)
}

// CollapseOnExceed sets the number of tabs before collapsing
func (t tabs) CollapseOnExceed(value string) tabs {
	return t.set("collapseOnExceed", value)
}

// ContentClassName sets the CSS class name for the content
func (t tabs) ContentClassName(value string) tabs {
	return t.set("contentClassName", value)
}

// DefaultKey sets the default active tab by hash or index, supports expressions
func (t tabs) DefaultKey(value any) tabs {
	return t.set("defaultKey", value)
}

// Disabled enables or disables the tabs
func (t tabs) Disabled(value bool) tabs {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to disable the tabs
func (t tabs) DisabledOn(value string) tabs {
	return t.set("disabledOn", value)
}

// Draggable enables or disables tab dragging
func (t tabs) Draggable(value bool) tabs {
	return t.set("draggable", value)
}

// Editable enables or disables tab label editing
func (t tabs) Editable(value bool) tabs {
	return t.set("editable", value)
}

// EditorSetting sets the editor configuration
func (t tabs) EditorSetting(value string) tabs {
	return t.set("editorSetting", value)
}

// Hidden hides or shows the tabs
func (t tabs) Hidden(value bool) tabs {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to hide the tabs
func (t tabs) HiddenOn(value string) tabs {
	return t.set("hiddenOn", value)
}

// ID sets the unique ID for the component
func (t tabs) ID(value string) tabs {
	return t.set("id", value)
}

// LinksClassName sets the CSS class name for the links container
func (t tabs) LinksClassName(value string) tabs {
	return t.set("linksClassName", value)
}

// MountOnEnter loads the tab content only when activated
func (t tabs) MountOnEnter(value bool) tabs {
	return t.set("mountOnEnter", value)
}

// OnEvent sets the event action configuration
func (t tabs) OnEvent(value any) tabs {
	return t.set("onEvent", value)
}

// Scrollable enables or disables scrolling for overflow content
func (t tabs) Scrollable(value bool) tabs {
	return t.set("scrollable", value)
}

// ShowTip enables or disables tooltips
func (t tabs) ShowTip(value bool) tabs {
	return t.set("showTip", value)
}

// ShowTipClassName sets the CSS class name for tooltips
func (t tabs) ShowTipClassName(value string) tabs {
	return t.set("showTipClassName", value)
}

// SidePosition sets the side position for the editor mode (left/right)
func (t tabs) SidePosition(value string) tabs {
	return t.set("sidePosition", value)
}

// Source sets the data source
func (t tabs) Source(value string) tabs {
	return t.set("source", value)
}

// Static enables or disables static display
func (t tabs) Static(value bool) tabs {
	return t.set("static", value)
}

// StaticClassName sets the CSS class name for static display items
func (t tabs) StaticClassName(value string) tabs {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static display item values
func (t tabs) StaticInputClassName(value string) tabs {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static display item labels
func (t tabs) StaticLabelClassName(value string) tabs {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (t tabs) StaticOn(value string) tabs {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display empty values
func (t tabs) StaticPlaceholder(value string) tabs {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (t tabs) StaticSchema(value string) tabs {
	return t.set("staticSchema", value)
}

// Style sets the component style
func (t tabs) Style(value any) tabs {
	return t.set("style", value)
}

// SubFormHorizontal sets the width ratio for horizontal layout
func (t tabs) SubFormHorizontal(value string) tabs {
	return t.set("subFormHorizontal", value)
}

// SubFormMode sets the default display mode for subform items
func (t tabs) SubFormMode(value string) tabs {
	return t.set("subFormMode", value)
}

// Swipeable enables or disables swipe switching (mobile only)
func (t tabs) Swipeable(value bool) tabs {
	return t.set("swipeable", value)
}

// Tabs sets the tab members
func (t tabs) Tabs(value ...any) tabs {
	return t.set("tabs", value)
}

// TabsMode sets the display mode (line, card, radio, vertical, chrome, simple, strong, tiled, sidebar)
func (t tabs) TabsMode(value string) tabs {
	return t.set("tabsMode", value)
}

// TestIdBuilder sets the test ID builder
func (t tabs) TestIdBuilder(value string) tabs {
	return t.set("testIdBuilder", value)
}

// Testid sets the test ID
func (t tabs) Testid(value string) tabs {
	return t.set("testid", value)
}

// Toolbar sets the toolbar configuration
func (t tabs) Toolbar(value string) tabs {
	return t.set("toolbar", value)
}

// UnmountOnExit destroys the tab content when hidden
func (t tabs) UnmountOnExit(value bool) tabs {
	return t.set("unmountOnExit", value)
}

// UseMobileUI enables or disables mobile UI styles
func (t tabs) UseMobileUI(value bool) tabs {
	return t.set("useMobileUI", value)
}

// Visible shows or hides the tabs
func (t tabs) Visible(value bool) tabs {
	return t.set("visible", value)
}

// VisibleOn sets the expression to show the tabs
func (t tabs) VisibleOn(value string) tabs {
	return t.set("visibleOn", value)
}
