package comp

import "github.com/zrcoder/amisgo/model"

// list represents a list display component.
type list model.Schema

// NewListRenderer creates a new List instance.
func List() list {
	return list{"type": "list"}
}

// set sets a field value.
func (lr list) set(key string, value any) list {
	lr[key] = value
	return lr
}

// AffixFooter sets whether to fix the footer.
func (lr list) AffixFooter(value bool) list {
	return lr.set("affixFooter", value)
}

// AffixHeader sets whether to fix the header.
func (lr list) AffixHeader(value bool) list {
	return lr.set("affixHeader", value)
}

// CheckOnItemClick sets whether to select on item click.
func (lr list) CheckOnItemClick(value bool) list {
	return lr.set("checkOnItemClick", value)
}

// ClassName sets the container CSS class name.
func (lr list) ClassName(value string) list {
	return lr.set("className", value)
}

// Disabled sets whether the component is disabled.
func (lr list) Disabled(value bool) list {
	return lr.set("disabled", value)
}

// DisabledOn sets the expression to disable the component.
func (lr list) DisabledOn(value string) list {
	return lr.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (lr list) EditorSetting(value string) list {
	return lr.set("editorSetting", value)
}

// Footer sets the footer area.
func (lr list) Footer(value string) list {
	return lr.set("footer", value)
}

// FooterClassName sets the footer area class name.
func (lr list) FooterClassName(value string) list {
	return lr.set("footerClassName", value)
}

// Header sets the header area.
func (lr list) Header(value string) list {
	return lr.set("header", value)
}

// HeaderClassName sets the header area class name.
func (lr list) HeaderClassName(value string) list {
	return lr.set("headerClassName", value)
}

// Hidden sets whether the component is hidden.
func (lr list) Hidden(value bool) list {
	return lr.set("hidden", value)
}

// HiddenOn sets the expression to hide the component.
func (lr list) HiddenOn(value string) list {
	return lr.set("hiddenOn", value)
}

// HideCheckToggler sets whether to hide the check toggler.
func (lr list) HideCheckToggler(value bool) list {
	return lr.set("hideCheckToggler", value)
}

// ID sets the unique ID of the component.
func (lr list) ID(value string) list {
	return lr.set("id", value)
}

// ItemAction sets the action on item click.
func (lr list) ItemAction(value string) list {
	return lr.set("itemAction", value)
}

// ItemCheckableOn sets the expression to make an item checkable.
func (lr list) ItemCheckableOn(value string) list {
	return lr.set("itemCheckableOn", value)
}

// ItemDraggableOn sets the expression to make an item draggable.
func (lr list) ItemDraggableOn(value string) list {
	return lr.set("itemDraggableOn", value)
}

// ListItem sets the configuration for displaying a single item.
func (lr list) ListItem(value string) list {
	return lr.set("listItem", value)
}

// OnEvent sets the event action configuration.
func (lr list) OnEvent(value any) list {
	return lr.set("onEvent", value)
}

// Placeholder sets the placeholder text when no data is available.
func (lr list) Placeholder(value string) list {
	return lr.set("placeholder", value)
}

// ShowFooter sets whether to show the footer.
func (lr list) ShowFooter(value bool) list {
	return lr.set("showFooter", value)
}

// ShowHeader sets whether to show the header.
func (lr list) ShowHeader(value bool) list {
	return lr.set("showHeader", value)
}

// Size sets the size of the component.
func (lr list) Size(value string) list {
	return lr.set("size", value)
}

// Source sets the data source.
func (lr list) Source(value string) list {
	return lr.set("source", value)
}

// Static sets whether the component is static.
func (lr list) Static(value bool) list {
	return lr.set("static", value)
}

// StaticClassName sets the static display form item class name.
func (lr list) StaticClassName(value string) list {
	return lr.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name.
func (lr list) StaticInputClassName(value string) list {
	return lr.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name.
func (lr list) StaticLabelClassName(value string) list {
	return lr.set("staticLabelClassName", value)
}

// StaticOn sets the expression to make the component static.
func (lr list) StaticOn(value string) list {
	return lr.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (lr list) StaticPlaceholder(value string) list {
	return lr.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display mode.
func (lr list) StaticSchema(value string) list {
	return lr.set("staticSchema", value)
}

// Style sets the component style.
func (lr list) Style(value any) list {
	return lr.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (lr list) TestIdBuilder(value string) list {
	return lr.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (lr list) Testid(value string) list {
	return lr.set("testid", value)
}

// Title sets the title.
func (lr list) Title(value any) list {
	return lr.set("title", value)
}

// Type sets the component type.
func (lr list) Type(value string) list {
	return lr.set("type", value)
}

// UseMobileUI sets whether to use mobile UI.
func (lr list) UseMobileUI(value bool) list {
	return lr.set("useMobileUI", value)
}

// ValueField sets the field used as the value.
func (lr list) ValueField(value string) list {
	return lr.set("valueField", value)
}

// Visible sets whether the component is visible.
func (lr list) Visible(value bool) list {
	return lr.set("visible", value)
}

// VisibleOn sets the expression to make the component visible.
func (lr list) VisibleOn(value string) list {
	return lr.set("visibleOn", value)
}
