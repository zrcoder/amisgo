package comp

import "github.com/zrcoder/amisgo/model"

// nav represents a navigation renderer
type nav model.Schema

// Nav creates a new Nav instance
func Nav() nav {
	return nav{}.set("type", "nav")
}

// set sets a property and returns the nav instance
func (n nav) set(key string, value any) nav {
	n[key] = value
	return n
}

// Accordion sets accordion mode, only supported in vertical inline mode
func (n nav) Accordion(value bool) nav {
	return n.set("accordion", value)
}

// Badge sets the badge
func (n nav) Badge(value string) nav {
	return n.set("badge", value)
}

// ClassName sets the container CSS class name
func (n nav) ClassName(value string) nav {
	return n.set("className", value)
}

// Collapsed sets the collapsed state
func (n nav) Collapsed(value bool) nav {
	return n.set("collapsed", value)
}

// DefaultOpenLevel sets the default open level
func (n nav) DefaultOpenLevel(value string) nav {
	return n.set("defaultOpenLevel", value)
}

// DeferApi sets the lazy load API
func (n nav) DeferApi(value string) nav {
	return n.set("deferApi", value)
}

// Disabled sets the disabled state
func (n nav) Disabled(value bool) nav {
	return n.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (n nav) DisabledOn(value string) nav {
	return n.set("disabledOn", value)
}

// DragOnSameLevel sets whether dragging is allowed only on the same level
func (n nav) DragOnSameLevel(value bool) nav {
	return n.set("dragOnSameLevel", value)
}

// Draggable sets the draggable state
func (n nav) Draggable(value bool) nav {
	return n.set("draggable", value)
}

// EditorSetting sets the editor configuration
func (n nav) EditorSetting(value string) nav {
	return n.set("editorSetting", value)
}

// ExpandIcon sets the custom expand icon
func (n nav) ExpandIcon(value string) nav {
	return n.set("expandIcon", value)
}

// ExpandPosition sets the custom expand icon position
func (n nav) ExpandPosition(value string) nav {
	return n.set("expandPosition", value)
}

// Hidden sets the hidden state
func (n nav) Hidden(value bool) nav {
	return n.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (n nav) HiddenOn(value string) nav {
	return n.set("hiddenOn", value)
}

// ID sets the unique component ID
func (n nav) ID(value string) nav {
	return n.set("id", value)
}

// IndentSize sets the indent size
func (n nav) IndentSize(value string) nav {
	return n.set("indentSize", value)
}

// ItemActions sets the item actions menu list
func (n nav) ItemActions(value string) nav {
	return n.set("itemActions", value)
}

// ItemBadge sets the item badge
func (n nav) ItemBadge(value string) nav {
	return n.set("itemBadge", value)
}

// Level sets the maximum display level
func (n nav) Level(value string) nav {
	return n.set("level", value)
}

// Links sets the links
func (n nav) Links(value ...any) nav {
	return n.set("links", value)
}

// Mode sets the menu open mode in vertical mode: inline | float (default: inline)
func (n nav) Mode(value string) nav {
	return n.set("mode", value)
}

// OnEvent sets the event action configuration
func (n nav) OnEvent(value any) nav {
	return n.set("onEvent", value)
}

// Overflow sets the auto-collapse configuration for horizontal navigation
func (n nav) Overflow(value string) nav {
	return n.set("overflow", value)
}

// PopupClassName sets the submenu popup layer style
func (n nav) PopupClassName(value string) nav {
	return n.set("popupClassName", value)
}

// SaveOrderApi sets the API for saving the order
func (n nav) SaveOrderApi(value string) nav {
	return n.set("saveOrderApi", value)
}

// SearchConfig sets the search box configuration
func (n nav) SearchConfig(value string) nav {
	return n.set("searchConfig", value)
}

// Searchable sets whether search is enabled
func (n nav) Searchable(value bool) nav {
	return n.set("searchable", value)
}

// ShowKey sets the key to display only the submenu items under the specified key
func (n nav) ShowKey(value string) nav {
	return n.set("showKey", value)
}

// Source sets the API for fetching data
func (n nav) Source(value string) nav {
	return n.set("source", value)
}

// Stacked sets the layout: true for vertical, false for horizontal
func (n nav) Stacked(value bool) nav {
	return n.set("stacked", value)
}

// Static sets whether to display statically
func (n nav) Static(value bool) nav {
	return n.set("static", value)
}

// StaticClassName sets the static display form item class name
func (n nav) StaticClassName(value string) nav {
	return n.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (n nav) StaticInputClassName(value string) nav {
	return n.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (n nav) StaticLabelClassName(value string) nav {
	return n.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (n nav) StaticOn(value string) nav {
	return n.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder for empty values
func (n nav) StaticPlaceholder(value string) nav {
	return n.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema
func (n nav) StaticSchema(value string) nav {
	return n.set("staticSchema", value)
}

// Style sets the component style
func (n nav) Style(value any) nav {
	return n.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (n nav) TestIdBuilder(value string) nav {
	return n.set("testIdBuilder", value)
}

// Testid sets the test ID
func (n nav) Testid(value string) nav {
	return n.set("testid", value)
}

// ThemeColor sets the theme color
func (n nav) ThemeColor(value string) nav {
	return n.set("themeColor", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (n nav) UseMobileUI(value bool) nav {
	return n.set("useMobileUI", value)
}

// Visible sets the visibility
func (n nav) Visible(value bool) nav {
	return n.set("visible", value)
}

// VisibleOn sets the visibility expression
func (n nav) VisibleOn(value string) nav {
	return n.set("visibleOn", value)
}
