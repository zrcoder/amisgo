package comp

// ListRenderer represents a list display component.
type ListRenderer Schema

// NewListRenderer creates a new ListRenderer instance.
func NewListRenderer() ListRenderer {
	lr := make(ListRenderer)
	lr.set("type", "list")
	return lr
}

// set sets a field value.
func (lr ListRenderer) set(key string, value any) ListRenderer {
	lr[key] = value
	return lr
}

// AffixFooter sets whether to fix the footer.
func (lr ListRenderer) AffixFooter(value bool) ListRenderer {
	return lr.set("affixFooter", value)
}

// AffixHeader sets whether to fix the header.
func (lr ListRenderer) AffixHeader(value bool) ListRenderer {
	return lr.set("affixHeader", value)
}

// CheckOnItemClick sets whether to select on item click.
func (lr ListRenderer) CheckOnItemClick(value bool) ListRenderer {
	return lr.set("checkOnItemClick", value)
}

// ClassName sets the container CSS class name.
func (lr ListRenderer) ClassName(value string) ListRenderer {
	return lr.set("className", value)
}

// Disabled sets whether the component is disabled.
func (lr ListRenderer) Disabled(value bool) ListRenderer {
	return lr.set("disabled", value)
}

// DisabledOn sets the expression to disable the component.
func (lr ListRenderer) DisabledOn(value string) ListRenderer {
	return lr.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (lr ListRenderer) EditorSetting(value string) ListRenderer {
	return lr.set("editorSetting", value)
}

// Footer sets the footer area.
func (lr ListRenderer) Footer(value string) ListRenderer {
	return lr.set("footer", value)
}

// FooterClassName sets the footer area class name.
func (lr ListRenderer) FooterClassName(value string) ListRenderer {
	return lr.set("footerClassName", value)
}

// Header sets the header area.
func (lr ListRenderer) Header(value string) ListRenderer {
	return lr.set("header", value)
}

// HeaderClassName sets the header area class name.
func (lr ListRenderer) HeaderClassName(value string) ListRenderer {
	return lr.set("headerClassName", value)
}

// Hidden sets whether the component is hidden.
func (lr ListRenderer) Hidden(value bool) ListRenderer {
	return lr.set("hidden", value)
}

// HiddenOn sets the expression to hide the component.
func (lr ListRenderer) HiddenOn(value string) ListRenderer {
	return lr.set("hiddenOn", value)
}

// HideCheckToggler sets whether to hide the check toggler.
func (lr ListRenderer) HideCheckToggler(value bool) ListRenderer {
	return lr.set("hideCheckToggler", value)
}

// ID sets the unique ID of the component.
func (lr ListRenderer) ID(value string) ListRenderer {
	return lr.set("id", value)
}

// ItemAction sets the action on item click.
func (lr ListRenderer) ItemAction(value string) ListRenderer {
	return lr.set("itemAction", value)
}

// ItemCheckableOn sets the expression to make an item checkable.
func (lr ListRenderer) ItemCheckableOn(value string) ListRenderer {
	return lr.set("itemCheckableOn", value)
}

// ItemDraggableOn sets the expression to make an item draggable.
func (lr ListRenderer) ItemDraggableOn(value string) ListRenderer {
	return lr.set("itemDraggableOn", value)
}

// ListItem sets the configuration for displaying a single item.
func (lr ListRenderer) ListItem(value string) ListRenderer {
	return lr.set("listItem", value)
}

// OnEvent sets the event action configuration.
func (lr ListRenderer) OnEvent(value any) ListRenderer {
	return lr.set("onEvent", value)
}

// Placeholder sets the placeholder text when no data is available.
func (lr ListRenderer) Placeholder(value string) ListRenderer {
	return lr.set("placeholder", value)
}

// ShowFooter sets whether to show the footer.
func (lr ListRenderer) ShowFooter(value bool) ListRenderer {
	return lr.set("showFooter", value)
}

// ShowHeader sets whether to show the header.
func (lr ListRenderer) ShowHeader(value bool) ListRenderer {
	return lr.set("showHeader", value)
}

// Size sets the size of the component.
func (lr ListRenderer) Size(value string) ListRenderer {
	return lr.set("size", value)
}

// Source sets the data source.
func (lr ListRenderer) Source(value string) ListRenderer {
	return lr.set("source", value)
}

// Static sets whether the component is static.
func (lr ListRenderer) Static(value bool) ListRenderer {
	return lr.set("static", value)
}

// StaticClassName sets the static display form item class name.
func (lr ListRenderer) StaticClassName(value string) ListRenderer {
	return lr.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name.
func (lr ListRenderer) StaticInputClassName(value string) ListRenderer {
	return lr.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name.
func (lr ListRenderer) StaticLabelClassName(value string) ListRenderer {
	return lr.set("staticLabelClassName", value)
}

// StaticOn sets the expression to make the component static.
func (lr ListRenderer) StaticOn(value string) ListRenderer {
	return lr.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display.
func (lr ListRenderer) StaticPlaceholder(value string) ListRenderer {
	return lr.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display mode.
func (lr ListRenderer) StaticSchema(value string) ListRenderer {
	return lr.set("staticSchema", value)
}

// Style sets the component style.
func (lr ListRenderer) Style(value any) ListRenderer {
	return lr.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (lr ListRenderer) TestIdBuilder(value string) ListRenderer {
	return lr.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (lr ListRenderer) Testid(value string) ListRenderer {
	return lr.set("testid", value)
}

// Title sets the title.
func (lr ListRenderer) Title(value any) ListRenderer {
	return lr.set("title", value)
}

// Type sets the component type.
func (lr ListRenderer) Type(value string) ListRenderer {
	return lr.set("type", value)
}

// UseMobileUI sets whether to use mobile UI.
func (lr ListRenderer) UseMobileUI(value bool) ListRenderer {
	return lr.set("useMobileUI", value)
}

// ValueField sets the field used as the value.
func (lr ListRenderer) ValueField(value string) ListRenderer {
	return lr.set("valueField", value)
}

// Visible sets whether the component is visible.
func (lr ListRenderer) Visible(value bool) ListRenderer {
	return lr.set("visible", value)
}

// VisibleOn sets the expression to make the component visible.
func (lr ListRenderer) VisibleOn(value string) ListRenderer {
	return lr.set("visibleOn", value)
}
