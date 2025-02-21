package comp

import "github.com/zrcoder/amisgo/schema"

// List represents a List display component.
type List schema.Schema

func NewList() List {
	return List{"type": "list"}
}

// set sets a field value.
func (lr List) set(key string, value any) List {
	lr[key] = value
	return lr
}

// AffixFooter sets whether to fix the footer.
func (lr List) AffixFooter(value bool) List {
	return lr.set("affixFooter", value)
}

// AffixHeader sets whether to fix the header.
func (lr List) AffixHeader(value bool) List {
	return lr.set("affixHeader", value)
}

// CheckOnItemClick sets whether to select on item click.
func (lr List) CheckOnItemClick(value bool) List {
	return lr.set("checkOnItemClick", value)
}

// ClassName sets the container CSS class name.
func (lr List) ClassName(value string) List {
	return lr.set("className", value)
}

// Disabled sets whether the component is disabled.
func (lr List) Disabled(value bool) List {
	return lr.set("disabled", value)
}

// DisabledOn sets the expression to disable the component.
func (lr List) DisabledOn(value string) List {
	return lr.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (lr List) EditorSetting(value string) List {
	return lr.set("editorSetting", value)
}

// Footer sets the footer area.
func (lr List) Footer(value string) List {
	return lr.set("footer", value)
}

// FooterClassName sets the footer area class name.
func (lr List) FooterClassName(value string) List {
	return lr.set("footerClassName", value)
}

// Header sets the header area.
func (lr List) Header(value string) List {
	return lr.set("header", value)
}

// HeaderClassName sets the header area class name.
func (lr List) HeaderClassName(value string) List {
	return lr.set("headerClassName", value)
}

// Hidden sets whether the component is hidden.
func (lr List) Hidden(value bool) List {
	return lr.set("hidden", value)
}

// HiddenOn sets the expression to hide the component.
func (lr List) HiddenOn(value string) List {
	return lr.set("hiddenOn", value)
}

// HideCheckToggler sets whether to hide the check toggler.
func (lr List) HideCheckToggler(value bool) List {
	return lr.set("hideCheckToggler", value)
}

// ID sets the unique ID of the component.
func (lr List) ID(value string) List {
	return lr.set("id", value)
}

// ItemAction sets the action on item click.
func (lr List) ItemAction(value string) List {
	return lr.set("itemAction", value)
}

// ItemCheckableOn sets the expression to make an item checkable.
func (lr List) ItemCheckableOn(value string) List {
	return lr.set("itemCheckableOn", value)
}

// ItemDraggableOn sets the expression to make an item draggable.
func (lr List) ItemDraggableOn(value string) List {
	return lr.set("itemDraggableOn", value)
}

// ListItem sets the configuration for displaying a single item.
func (lr List) ListItem(value string) List {
	return lr.set("listItem", value)
}

// OnEvent sets the event action configuration.
func (lr List) OnEvent(value Event) List {
	return lr.set("onEvent", value)
}

// Placeholder sets the placeholder text when no data is available.
func (lr List) Placeholder(value string) List {
	return lr.set("placeholder", value)
}

// ShowFooter sets whether to show the footer.
func (lr List) ShowFooter(value bool) List {
	return lr.set("showFooter", value)
}

// ShowHeader sets whether to show the header.
func (lr List) ShowHeader(value bool) List {
	return lr.set("showHeader", value)
}

// Size sets the size of the component.
func (lr List) Size(value string) List {
	return lr.set("size", value)
}

// Source sets the data source.
func (lr List) Source(value string) List {
	return lr.set("source", value)
}

// Static sets whether the component is static.
func (lr List) Static(value bool) List {
	return lr.set("static", value)
}

// StaticClassName sets the static display form item class name.
func (lr List) StaticClassName(value string) List {
	return lr.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name.
func (lr List) StaticInputClassName(value string) List {
	return lr.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name.
func (lr List) StaticLabelClassName(value string) List {
	return lr.set("staticLabelClassName", value)
}

// StaticOn sets the expression to make the component static.
func (lr List) StaticOn(value string) List {
	return lr.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (lr List) StaticPlaceholder(value string) List {
	return lr.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display mode.
func (lr List) StaticSchema(value string) List {
	return lr.set("staticSchema", value)
}

// Style sets the component style.
func (lr List) Style(value any) List {
	return lr.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (lr List) TestIdBuilder(value string) List {
	return lr.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (lr List) Testid(value string) List {
	return lr.set("testid", value)
}

// Title sets the title.
func (lr List) Title(value any) List {
	return lr.set("title", value)
}

// Type sets the component type.
func (lr List) Type(value string) List {
	return lr.set("type", value)
}

// UseMobileUI sets whether to use mobile UI.
func (lr List) UseMobileUI(value bool) List {
	return lr.set("useMobileUI", value)
}

// ValueField sets the field used as the value.
func (lr List) ValueField(value string) List {
	return lr.set("valueField", value)
}

// Visible sets whether the component is visible.
func (lr List) Visible(value bool) List {
	return lr.set("visible", value)
}

// VisibleOn sets the expression to make the component visible.
func (lr List) VisibleOn(value string) List {
	return lr.set("visibleOn", value)
}
