package comp

import "github.com/zrcoder/amisgo/schema"

// ListItem represents a list item.
type ListItem schema.Schema

func NewListItem() ListItem {
	return make(ListItem)
}

// set sets a field value.
func (li ListItem) set(key string, value any) ListItem {
	li[key] = value
	return li
}

// Actions sets actions.
func (li ListItem) Actions(value string) ListItem {
	return li.set("actions", value)
}

// ActionsPosition sets the position of actions (left or right).
func (li ListItem) ActionsPosition(value string) ListItem {
	return li.set("actionsPosition", value)
}

// Avatar sets the image URL.
func (li ListItem) Avatar(value string) ListItem {
	return li.set("avatar", value)
}

// Body sets the content area.
func (li ListItem) Body(value ...any) ListItem {
	return li.set("body", value)
}

// ClassName sets the container CSS class name.
func (li ListItem) ClassName(value string) ListItem {
	return li.set("className", value)
}

// Desc sets the description.
func (li ListItem) Desc(value string) ListItem {
	return li.set("desc", value)
}

// Disabled sets whether the item is disabled.
func (li ListItem) Disabled(value bool) ListItem {
	return li.set("disabled", value)
}

// DisabledOn sets the expression to determine if the item is disabled.
func (li ListItem) DisabledOn(value string) ListItem {
	return li.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (li ListItem) EditorSetting(value string) ListItem {
	return li.set("editorSetting", value)
}

// Hidden sets whether the item is hidden.
func (li ListItem) Hidden(value bool) ListItem {
	return li.set("hidden", value)
}

// HiddenOn sets the expression to determine if the item is hidden.
func (li ListItem) HiddenOn(value string) ListItem {
	return li.set("hiddenOn", value)
}

// ID sets the unique component ID.
func (li ListItem) ID(value string) ListItem {
	return li.set("id", value)
}

// OnEvent sets the event action configuration.
func (li ListItem) OnEvent(value any) ListItem {
	return li.set("onEvent", value)
}

// Remark sets the tooltip description.
func (li ListItem) Remark(value string) ListItem {
	return li.set("remark", value)
}

// Static sets whether the item is displayed statically.
func (li ListItem) Static(value bool) ListItem {
	return li.set("static", value)
}

// StaticClassName sets the CSS class name for static display.
func (li ListItem) StaticClassName(value string) ListItem {
	return li.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display.
func (li ListItem) StaticInputClassName(value string) ListItem {
	return li.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display.
func (li ListItem) StaticLabelClassName(value string) ListItem {
	return li.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the item is displayed statically.
func (li ListItem) StaticOn(value string) ListItem {
	return li.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (li ListItem) StaticPlaceholder(value string) ListItem {
	return li.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display mode.
func (li ListItem) StaticSchema(value string) ListItem {
	return li.set("staticSchema", value)
}

// Style sets the component style.
func (li ListItem) Style(value any) ListItem {
	return li.set("style", value)
}

// SubTitle sets the subtitle.
func (li ListItem) SubTitle(value any) ListItem {
	return li.set("subTitle", value)
}

// TestIdBuilder sets the test ID builder.
func (li ListItem) TestIdBuilder(value string) ListItem {
	return li.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (li ListItem) Testid(value string) ListItem {
	return li.set("testid", value)
}

// Title sets the title.
func (li ListItem) Title(value any) ListItem {
	return li.set("title", value)
}

// UseMobileUI sets whether to use mobile UI.
func (li ListItem) UseMobileUI(value bool) ListItem {
	return li.set("useMobileUI", value)
}

// Visible sets whether the item is visible.
func (li ListItem) Visible(value bool) ListItem {
	return li.set("visible", value)
}

// VisibleOn sets the expression to determine if the item is visible.
func (li ListItem) VisibleOn(value string) ListItem {
	return li.set("visibleOn", value)
}
