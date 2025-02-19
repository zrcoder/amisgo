package comp

import "github.com/zrcoder/amisgo/schema"

// Tag represents a tag component
type Tag schema.Schema

func NewTag() Tag {
	return Tag{"type": "tag"}
}

func (t Tag) set(key string, value any) Tag {
	t[key] = value
	return t
}

// Checkable sets whether the tag is checkable
func (t Tag) Checkable(value bool) Tag {
	return t.set("checkable", value)
}

// Checked sets whether the tag is checked
func (t Tag) Checked(value bool) Tag {
	return t.set("checked", value)
}

// ClassName sets the class name of the tag
func (t Tag) ClassName(value string) Tag {
	return t.set("className", value)
}

// Closable sets whether the tag is closable
func (t Tag) Closable(value bool) Tag {
	return t.set("closable", value)
}

// CloseIcon sets the close icon of the tag
func (t Tag) CloseIcon(value string) Tag {
	return t.set("closeIcon", value)
}

// Color sets the color of the tag
func (t Tag) Color(value string) Tag {
	return t.set("color", value)
}

// Disabled sets whether the tag is disabled
func (t Tag) Disabled(value bool) Tag {
	return t.set("disabled", value)
}

// DisabledOn sets the expression for disabling the tag
func (t Tag) DisabledOn(value string) Tag {
	return t.set("disabledOn", value)
}

// DisplayMode sets the display mode of the tag
func (t Tag) DisplayMode(value string) Tag {
	return t.set("displayMode", value)
}

// EditorSetting sets the editor configuration
func (t Tag) EditorSetting(value string) Tag {
	return t.set("editorSetting", value)
}

// Hidden sets whether the tag is hidden
func (t Tag) Hidden(value bool) Tag {
	return t.set("hidden", value)
}

// HiddenOn sets the expression for hiding the tag
func (t Tag) HiddenOn(value string) Tag {
	return t.set("hiddenOn", value)
}

// Icon sets the icon of the tag
func (t Tag) Icon(value string) Tag {
	return t.set("icon", value)
}

// ID sets the unique ID of the tag
func (t Tag) ID(value string) Tag {
	return t.set("id", value)
}

// Label sets the label text of the tag
func (t Tag) Label(value string) Tag {
	return t.set("label", value)
}

// OnEvent sets the event configuration
func (t Tag) OnEvent(value any) Tag {
	return t.set("onEvent", value)
}

// Static sets whether the tag is static
func (t Tag) Static(value bool) Tag {
	return t.set("static", value)
}

// StaticClassName sets the class name for static display
func (t Tag) StaticClassName(value string) Tag {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the class name for static input display
func (t Tag) StaticInputClassName(value string) Tag {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the class name for static label display
func (t Tag) StaticLabelClassName(value string) Tag {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (t Tag) StaticOn(value string) Tag {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (t Tag) StaticPlaceholder(value string) Tag {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (t Tag) StaticSchema(value string) Tag {
	return t.set("staticSchema", value)
}

// Style sets the custom style of the tag
func (t Tag) Style(value any) Tag {
	return t.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (t Tag) TestIdBuilder(value string) Tag {
	return t.set("testIdBuilder", value)
}

// Testid sets the test ID
func (t Tag) Testid(value string) Tag {
	return t.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI
func (t Tag) UseMobileUI(value bool) Tag {
	return t.set("useMobileUI", value)
}

// Visible sets whether the tag is visible
func (t Tag) Visible(value bool) Tag {
	return t.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (t Tag) VisibleOn(value string) Tag {
	return t.set("visibleOn", value)
}
