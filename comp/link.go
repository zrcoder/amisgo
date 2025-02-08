package comp

import "github.com/zrcoder/amisgo/model"

// Link represents a Link display component.
type Link model.Schema

func NewLink() Link {
	return Link{"type": "link"}
}

// Badge sets the badge value.
func (l Link) Badge(value string) Link {
	return l.set("badge", value)
}

// Blank sets whether to open in a new window.
func (l Link) Blank(value bool) Link {
	return l.set("blank", value)
}

// Body sets the link content.
func (l Link) Body(value ...any) Link {
	return l.set("body", value)
}

// ClassName sets the container CSS class name.
func (l Link) ClassName(value string) Link {
	return l.set("className", value)
}

// Disabled sets whether the link is disabled.
func (l Link) Disabled(value bool) Link {
	return l.set("disabled", value)
}

// DisabledOn sets the expression to disable the link.
func (l Link) DisabledOn(value string) Link {
	return l.set("disabledOn", value)
}

// EditorSetting sets the editor configuration.
func (l Link) EditorSetting(value string) Link {
	return l.set("editorSetting", value)
}

// Hidden sets whether the link is hidden.
func (l Link) Hidden(value bool) Link {
	return l.set("hidden", value)
}

// HiddenOn sets the expression to hide the link.
func (l Link) HiddenOn(value string) Link {
	return l.set("hiddenOn", value)
}

// Href sets the link address.
func (l Link) Href(value string) Link {
	return l.set("href", value)
}

// HtmlTarget sets the native target attribute of the <a> tag.
func (l Link) HtmlTarget(value string) Link {
	return l.set("htmlTarget", value)
}

// Icon sets the icon.
func (l Link) Icon(value string) Link {
	return l.set("icon", value)
}

// ID sets the unique component ID.
func (l Link) ID(value string) Link {
	return l.set("id", value)
}

// OnEvent sets the event action configuration.
func (l Link) OnEvent(value any) Link {
	return l.set("onEvent", value)
}

// RightIcon sets the right-side icon.
func (l Link) RightIcon(value string) Link {
	return l.set("rightIcon", value)
}

// Static sets whether to display statically.
func (l Link) Static(value bool) Link {
	return l.set("static", value)
}

// StaticClassName sets the static display form item class name.
func (l Link) StaticClassName(value string) Link {
	return l.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name.
func (l Link) StaticInputClassName(value string) Link {
	return l.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name.
func (l Link) StaticLabelClassName(value string) Link {
	return l.set("staticLabelClassName", value)
}

// StaticOn sets the expression to display statically.
func (l Link) StaticOn(value string) Link {
	return l.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder.
func (l Link) StaticPlaceholder(value string) Link {
	return l.set("staticPlaceholder", value)
}

// StaticSchema sets the static display schema.
func (l Link) StaticSchema(value string) Link {
	return l.set("staticSchema", value)
}

// Style sets the component style.
func (l Link) Style(value any) Link {
	return l.set("style", value)
}

// TestIdBuilder sets the test ID builder.
func (l Link) TestIdBuilder(value string) Link {
	return l.set("testIdBuilder", value)
}

// Testid sets the test ID.
func (l Link) Testid(value string) Link {
	return l.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles.
func (l Link) UseMobileUI(value bool) Link {
	return l.set("useMobileUI", value)
}

// Visible sets whether the link is visible.
func (l Link) Visible(value bool) Link {
	return l.set("visible", value)
}

// VisibleOn sets the expression to display the link.
func (l Link) VisibleOn(value string) Link {
	return l.set("visibleOn", value)
}

// set sets a key-value pair.
func (l Link) set(key string, value any) Link {
	l[key] = value
	return l
}
