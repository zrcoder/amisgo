package comp

import "github.com/zrcoder/amisgo/schema"

// MultilineText represents a multiline text component
type MultilineText schema.Schema

func NewMultilineText() MultilineText {
	return MultilineText{"type": "multiline-text"}
}

// set sets a property and returns the instance
func (m MultilineText) set(key string, value any) MultilineText {
	m[key] = value
	return m
}

// ClassName sets the CSS class name
func (m MultilineText) ClassName(value string) MultilineText {
	return m.set("className", value)
}

// CollapseButtonText sets the collapse button text
func (m MultilineText) CollapseButtonText(value string) MultilineText {
	return m.set("collapseButtonText", value)
}

// Disabled sets the disabled state
func (m MultilineText) Disabled(value bool) MultilineText {
	return m.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (m MultilineText) DisabledOn(value string) MultilineText {
	return m.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (m MultilineText) EditorSetting(value string) MultilineText {
	return m.set("editorSetting", value)
}

// ExpendButtonText sets the expand button text
func (m MultilineText) ExpendButtonText(value string) MultilineText {
	return m.set("expendButtonText", value)
}

// Hidden sets the hidden state
func (m MultilineText) Hidden(value bool) MultilineText {
	return m.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (m MultilineText) HiddenOn(value string) MultilineText {
	return m.set("hiddenOn", value)
}

// ID sets the unique component ID
func (m MultilineText) ID(value string) MultilineText {
	return m.set("id", value)
}

// MaxRows sets the maximum number of rows
func (m MultilineText) MaxRows(value string) MultilineText {
	return m.set("maxRows", value)
}

// OnEvent sets the event configuration
func (m MultilineText) OnEvent(value any) MultilineText {
	return m.set("onEvent", value)
}

// Static sets the static display state
func (m MultilineText) Static(value bool) MultilineText {
	return m.set("static", value)
}

// StaticClassName sets the static display class name
func (m MultilineText) StaticClassName(value string) MultilineText {
	return m.set("staticClassName", value)
}

// StaticInputClassName sets the static input class name
func (m MultilineText) StaticInputClassName(value string) MultilineText {
	return m.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label class name
func (m MultilineText) StaticLabelClassName(value string) MultilineText {
	return m.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (m MultilineText) StaticOn(value string) MultilineText {
	return m.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder
func (m MultilineText) StaticPlaceholder(value string) MultilineText {
	return m.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema
func (m MultilineText) StaticSchema(value string) MultilineText {
	return m.set("staticSchema", value)
}

// Style sets the component style
func (m MultilineText) Style(value any) MultilineText {
	return m.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (m MultilineText) TestIdBuilder(value string) MultilineText {
	return m.set("testIdBuilder", value)
}

// Testid sets the test ID
func (m MultilineText) Testid(value string) MultilineText {
	return m.set("testid", value)
}

// Text sets the text content
func (m MultilineText) Text(value string) MultilineText {
	return m.set("text", value)
}

// UseMobileUI sets the mobile UI state
func (m MultilineText) UseMobileUI(value bool) MultilineText {
	return m.set("useMobileUI", value)
}

// Visible sets the visibility state
func (m MultilineText) Visible(value bool) MultilineText {
	return m.set("visible", value)
}

// VisibleOn sets the visibility expression
func (m MultilineText) VisibleOn(value string) MultilineText {
	return m.set("visibleOn", value)
}
