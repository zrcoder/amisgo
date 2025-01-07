package comp

// link represents a link display component.

type link Schema

// Link creates a new Link instance with default type.
func Link() link {
	return make(link).set("type", "link")
}

// Badge sets the badge value.
func (l link) Badge(value string) link {
	return l.set("badge", value)
}

// Blank sets whether to open in a new window.
func (l link) Blank(value bool) link {
	return l.set("blank", value)
}

// Body sets the link content.
func (l link) Body(value ...any) link {
	return l.set("body", value)
}

// ClassName sets the container CSS class name.
func (l link) ClassName(value string) link {
	return l.set("className", value)
}

// Disabled sets whether the link is disabled.
func (l link) Disabled(value bool) link {
	return l.set("disabled", value)
}

// DisabledOn sets the expression to disable the link.
func (l link) DisabledOn(value string) link {
	return l.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (l link) EditorSetting(value string) link {
	return l.set("editorSetting", value)
}

// Hidden sets whether the link is hidden.
func (l link) Hidden(value bool) link {
	return l.set("hidden", value)
}

// HiddenOn sets the expression to hide the link.
func (l link) HiddenOn(value string) link {
	return l.set("hiddenOn", value)
}

// Href sets the link address.
func (l link) Href(value string) link {
	return l.set("href", value)
}

// HtmlTarget sets the native target attribute of the <a> tag.
func (l link) HtmlTarget(value string) link {
	return l.set("htmlTarget", value)
}

// Icon sets the icon.
func (l link) Icon(value string) link {
	return l.set("icon", value)
}

// ID sets the unique component ID.
func (l link) ID(value string) link {
	return l.set("id", value)
}

// OnEvent sets the event action configuration.
func (l link) OnEvent(value any) link {
	return l.set("onEvent", value)
}

// RightIcon sets the right-side icon.
func (l link) RightIcon(value string) link {
	return l.set("rightIcon", value)
}

// Static sets whether to display statically.
func (l link) Static(value bool) link {
	return l.set("static", value)
}

// StaticClassName sets the static display form item class name.
func (l link) StaticClassName(value string) link {
	return l.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name.
func (l link) StaticInputClassName(value string) link {
	return l.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name.
func (l link) StaticLabelClassName(value string) link {
	return l.set("staticLabelClassName", value)
}

// StaticOn sets the expression to display statically.
func (l link) StaticOn(value string) link {
	return l.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder.
func (l link) StaticPlaceholder(value string) link {
	return l.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema.
func (l link) StaticSchema(value string) link {
	return l.set("staticSchema", value)
}

// Style sets the component style.
func (l link) Style(value any) link {
	return l.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (l link) TestIdBuilder(value string) link {
	return l.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (l link) Testid(value string) link {
	return l.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles.
func (l link) UseMobileUI(value bool) link {
	return l.set("useMobileUI", value)
}

// Visible sets whether the link is visible.
func (l link) Visible(value bool) link {
	return l.set("visible", value)
}

// VisibleOn sets the expression to display the link.
func (l link) VisibleOn(value string) link {
	return l.set("visibleOn", value)
}

// set sets a key-value pair.
func (l link) set(key string, value any) link {
	l[key] = value
	return l
}
