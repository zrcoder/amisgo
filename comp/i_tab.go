package comp

import "github.com/zrcoder/amisgo/model"

type MTab model.Schema

// Tab creates a new Tab instance
func Tab() MTab {
	return MTab{}
}

func (t MTab) set(key string, value any) MTab {
	t[key] = value
	return t
}

// Badge sets the badge
func (t MTab) Badge(value string) MTab {
	return t.set("badge", value)
}

// Body sets the content
func (t MTab) Body(value ...any) MTab {
	return t.set("body", value)
}

// ClassName sets the container CSS class name
func (t MTab) ClassName(value string) MTab {
	return t.set("className", value)
}

// Closable sets whether the tab is closable
func (t MTab) Closable(value bool) MTab {
	return t.set("closable", value)
}

// Disabled sets whether the tab is disabled
func (t MTab) Disabled(value bool) MTab {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to disable the tab
func (t MTab) DisabledOn(value string) MTab {
	return t.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (t MTab) EditorSetting(value string) MTab {
	return t.set("editorSetting", value)
}

// Hash sets the URL hash
func (t MTab) Hash(value string) MTab {
	return t.set("hash", value)
}

// Hidden sets whether the tab is hidden
func (t MTab) Hidden(value bool) MTab {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to hide the tab
func (t MTab) HiddenOn(value string) MTab {
	return t.set("hiddenOn", value)
}

// Horizontal sets the horizontal layout ratio
func (t MTab) Horizontal(value string) MTab {
	return t.set("horizontal", value)
}

// Icon sets the button icon
func (t MTab) Icon(value string) MTab {
	return t.set("icon", value)
}

// IconPosition sets the icon position (left or right)
func (t MTab) IconPosition(value string) MTab {
	return t.set("iconPosition", value)
}

// ID sets the unique component ID
func (t MTab) ID(value string) MTab {
	return t.set("id", value)
}

// Mode sets the default display mode (normal, inline, horizontal)
func (t MTab) Mode(value string) MTab {
	return t.set("mode", value)
}

// MountOnEnter sets whether to load content on tab activation
func (t MTab) MountOnEnter(value bool) MTab {
	return t.set("mountOnEnter", value)
}

// OnEvent sets the event action configuration
func (t MTab) OnEvent(value any) MTab {
	return t.set("onEvent", value)
}

// Reload sets whether to reload content on each render
func (t MTab) Reload(value bool) MTab {
	return t.set("reload", value)
}

// Static sets whether the tab is static
func (t MTab) Static(value bool) MTab {
	return t.set("static", value)
}

// StaticClassName sets the static form item class name
func (t MTab) StaticClassName(value string) MTab {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the static form item value class name
func (t MTab) StaticInputClassName(value string) MTab {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static form item label class name
func (t MTab) StaticLabelClassName(value string) MTab {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (t MTab) StaticOn(value string) MTab {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (t MTab) StaticPlaceholder(value string) MTab {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (t MTab) StaticSchema(value string) MTab {
	return t.set("staticSchema", value)
}

// Style sets the component style
func (t MTab) Style(value string) MTab {
	return t.set("style", value)
}

// Tab sets the tab content
func (t MTab) Tab(value ...any) MTab {
	return t.set("tab", value)
}

// TestIdBuilder sets the test ID builder
func (t MTab) TestIdBuilder(value string) MTab {
	return t.set("testIdBuilder", value)
}

// Testid sets the test ID
func (t MTab) Testid(value string) MTab {
	return t.set("testid", value)
}

// Title sets the tab title
func (t MTab) Title(value any) MTab {
	return t.set("title", value)
}

// UnmountOnExit sets whether to destroy the tab node on hide
func (t MTab) UnmountOnExit(value bool) MTab {
	return t.set("unmountOnExit", value)
}

// UseMobileUI sets whether to disable mobile UI styles
func (t MTab) UseMobileUI(value bool) MTab {
	return t.set("useMobileUI", value)
}

// Visible sets whether the tab is visible
func (t MTab) Visible(value bool) MTab {
	return t.set("visible", value)
}

// VisibleOn sets the expression to show the tab
func (t MTab) VisibleOn(value string) MTab {
	return t.set("visibleOn", value)
}
