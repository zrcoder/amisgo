package comp

import "github.com/zrcoder/amisgo/schema"

// Nav represents a navigation renderer
type Nav schema.Schema

func NewNav() Nav {
	return Nav{"type": "nav"}
}

// set sets a property and returns the nav instance
func (n Nav) set(key string, value any) Nav {
	n[key] = value
	return n
}

// Accordion sets accordion mode, only supported in vertical inline mode
func (n Nav) Accordion(value bool) Nav {
	return n.set("accordion", value)
}

// Badge sets the badge
func (n Nav) Badge(value string) Nav {
	return n.set("badge", value)
}

// ClassName sets the container CSS class name
func (n Nav) ClassName(value string) Nav {
	return n.set("className", value)
}

// Collapsed sets the collapsed state
func (n Nav) Collapsed(value bool) Nav {
	return n.set("collapsed", value)
}

// DefaultOpenLevel sets the default open level
func (n Nav) DefaultOpenLevel(value string) Nav {
	return n.set("defaultOpenLevel", value)
}

// DeferApi sets the lazy load API
func (n Nav) DeferApi(value string) Nav {
	return n.set("deferApi", value)
}

// Disabled sets the disabled state
func (n Nav) Disabled(value bool) Nav {
	return n.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (n Nav) DisabledOn(value string) Nav {
	return n.set("disabledOn", value)
}

// DragOnSameLevel sets whether dragging is allowed only on the same level
func (n Nav) DragOnSameLevel(value bool) Nav {
	return n.set("dragOnSameLevel", value)
}

// Draggable sets the draggable state
func (n Nav) Draggable(value bool) Nav {
	return n.set("draggable", value)
}

// EditorSetting sets the editor configuration
func (n Nav) EditorSetting(value string) Nav {
	return n.set("editorSetting", value)
}

// ExpandIcon sets the custom expand icon
func (n Nav) ExpandIcon(value string) Nav {
	return n.set("expandIcon", value)
}

// ExpandPosition sets the custom expand icon position
func (n Nav) ExpandPosition(value string) Nav {
	return n.set("expandPosition", value)
}

// Hidden sets the hidden state
func (n Nav) Hidden(value bool) Nav {
	return n.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (n Nav) HiddenOn(value string) Nav {
	return n.set("hiddenOn", value)
}

// ID sets the unique component ID
func (n Nav) ID(value string) Nav {
	return n.set("id", value)
}

// IndentSize sets the indent size
func (n Nav) IndentSize(value string) Nav {
	return n.set("indentSize", value)
}

// ItemActions sets the item actions menu list
func (n Nav) ItemActions(value string) Nav {
	return n.set("itemActions", value)
}

// ItemBadge sets the item badge
func (n Nav) ItemBadge(value string) Nav {
	return n.set("itemBadge", value)
}

// Level sets the maximum display level
func (n Nav) Level(value string) Nav {
	return n.set("level", value)
}

// Links sets the links
func (n Nav) Links(value ...any) Nav {
	return n.set("links", value)
}

// Mode sets the menu open mode in vertical mode: inline | float (default: inline)
func (n Nav) Mode(value string) Nav {
	return n.set("mode", value)
}

// OnEvent sets the event action configuration
func (n Nav) OnEvent(value any) Nav {
	return n.set("onEvent", value)
}

// Overflow sets the auto-collapse configuration for horizontal navigation
func (n Nav) Overflow(value string) Nav {
	return n.set("overflow", value)
}

// PopupClassName sets the submenu popup layer style
func (n Nav) PopupClassName(value string) Nav {
	return n.set("popupClassName", value)
}

// SaveOrderApi sets the API for saving the order
func (n Nav) SaveOrderApi(value string) Nav {
	return n.set("saveOrderApi", value)
}

// SearchConfig sets the search box configuration
func (n Nav) SearchConfig(value string) Nav {
	return n.set("searchConfig", value)
}

// Searchable sets whether search is enabled
func (n Nav) Searchable(value bool) Nav {
	return n.set("searchable", value)
}

// ShowKey sets the key to display only the submenu items under the specified key
func (n Nav) ShowKey(value string) Nav {
	return n.set("showKey", value)
}

// Source sets the API for fetching data
func (n Nav) Source(value string) Nav {
	return n.set("source", value)
}

// Stacked sets the layout: true for vertical, false for horizontal
func (n Nav) Stacked(value bool) Nav {
	return n.set("stacked", value)
}

// Static sets whether to display statically
func (n Nav) Static(value bool) Nav {
	return n.set("static", value)
}

// StaticClassName sets the static display form item class name
func (n Nav) StaticClassName(value string) Nav {
	return n.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (n Nav) StaticInputClassName(value string) Nav {
	return n.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (n Nav) StaticLabelClassName(value string) Nav {
	return n.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (n Nav) StaticOn(value string) Nav {
	return n.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder for empty values
func (n Nav) StaticPlaceholder(value string) Nav {
	return n.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema.Schema
func (n Nav) StaticSchema(value string) Nav {
	return n.set("staticSchema", value)
}

// Style sets the component style
func (n Nav) Style(value any) Nav {
	return n.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (n Nav) TestIdBuilder(value string) Nav {
	return n.set("testIdBuilder", value)
}

// Testid sets the test ID
func (n Nav) Testid(value string) Nav {
	return n.set("testid", value)
}

// ThemeColor sets the theme color
func (n Nav) ThemeColor(value string) Nav {
	return n.set("themeColor", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (n Nav) UseMobileUI(value bool) Nav {
	return n.set("useMobileUI", value)
}

// Visible sets the visibility
func (n Nav) Visible(value bool) Nav {
	return n.set("visible", value)
}

// VisibleOn sets the visibility expression
func (n Nav) VisibleOn(value string) Nav {
	return n.set("visibleOn", value)
}
