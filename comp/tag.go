package comp

import "github.com/zrcoder/amisgo/model"

// tag represents a tag component

type tag model.Schema

// Tag creates a new Tag instance
func Tag() tag {
	return tag{"type": "tag"}
}

func (t tag) set(key string, value any) tag {
	t[key] = value
	return t
}

// Checkable sets whether the tag is checkable
func (t tag) Checkable(value bool) tag {
	return t.set("checkable", value)
}

// Checked sets whether the tag is checked
func (t tag) Checked(value bool) tag {
	return t.set("checked", value)
}

// ClassName sets the class name of the tag
func (t tag) ClassName(value string) tag {
	return t.set("className", value)
}

// Closable sets whether the tag is closable
func (t tag) Closable(value bool) tag {
	return t.set("closable", value)
}

// CloseIcon sets the close icon of the tag
func (t tag) CloseIcon(value string) tag {
	return t.set("closeIcon", value)
}

// Color sets the color of the tag
func (t tag) Color(value string) tag {
	return t.set("color", value)
}

// Disabled sets whether the tag is disabled
func (t tag) Disabled(value bool) tag {
	return t.set("disabled", value)
}

// DisabledOn sets the expression for disabling the tag
func (t tag) DisabledOn(value string) tag {
	return t.set("disabledOn", value)
}

// DisplayMode sets the display mode of the tag
func (t tag) DisplayMode(value string) tag {
	return t.set("displayMode", value)
}

// EditorSetting sets the editor configuration
func (t tag) EditorSetting(value string) tag {
	return t.set("editorSetting", value)
}

// Hidden sets whether the tag is hidden
func (t tag) Hidden(value bool) tag {
	return t.set("hidden", value)
}

// HiddenOn sets the expression for hiding the tag
func (t tag) HiddenOn(value string) tag {
	return t.set("hiddenOn", value)
}

// Icon sets the icon of the tag
func (t tag) Icon(value string) tag {
	return t.set("icon", value)
}

// ID sets the unique ID of the tag
func (t tag) ID(value string) tag {
	return t.set("id", value)
}

// Label sets the label text of the tag
func (t tag) Label(value string) tag {
	return t.set("label", value)
}

// OnEvent sets the event configuration
func (t tag) OnEvent(value any) tag {
	return t.set("onEvent", value)
}

// Static sets whether the tag is static
func (t tag) Static(value bool) tag {
	return t.set("static", value)
}

// StaticClassName sets the class name for static display
func (t tag) StaticClassName(value string) tag {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the class name for static input display
func (t tag) StaticInputClassName(value string) tag {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class name for static label display
func (t tag) StaticLabelClassName(value string) tag {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (t tag) StaticOn(value string) tag {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (t tag) StaticPlaceholder(value string) tag {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (t tag) StaticSchema(value string) tag {
	return t.set("staticSchema", value)
}

// Style sets the custom style of the tag
func (t tag) Style(value any) tag {
	return t.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (t tag) TestIdBuilder(value string) tag {
	return t.set("testIdBuilder", value)
}

// Testid sets the test ID
func (t tag) Testid(value string) tag {
	return t.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI
func (t tag) UseMobileUI(value bool) tag {
	return t.set("useMobileUI", value)
}

// Visible sets whether the tag is visible
func (t tag) Visible(value bool) tag {
	return t.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (t tag) VisibleOn(value string) tag {
	return t.set("visibleOn", value)
}
