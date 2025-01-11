package comp

import "github.com/zrcoder/amisgo/model"

// multilineText represents a multiline text component
type multilineText model.Schema

// MultilineText creates a new multilineText instance
func MultilineText() multilineText {
	return multilineText{"type": "multiline-text"}
}

// set sets a property and returns the instance
func (m multilineText) set(key string, value any) multilineText {
	m[key] = value
	return m
}

// ClassName sets the CSS class name
func (m multilineText) ClassName(value string) multilineText {
	return m.set("className", value)
}

// CollapseButtonText sets the collapse button text
func (m multilineText) CollapseButtonText(value string) multilineText {
	return m.set("collapseButtonText", value)
}

// Disabled sets the disabled state
func (m multilineText) Disabled(value bool) multilineText {
	return m.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (m multilineText) DisabledOn(value string) multilineText {
	return m.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (m multilineText) EditorSetting(value string) multilineText {
	return m.set("editorSetting", value)
}

// ExpendButtonText sets the expand button text
func (m multilineText) ExpendButtonText(value string) multilineText {
	return m.set("expendButtonText", value)
}

// Hidden sets the hidden state
func (m multilineText) Hidden(value bool) multilineText {
	return m.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (m multilineText) HiddenOn(value string) multilineText {
	return m.set("hiddenOn", value)
}

// ID sets the unique component ID
func (m multilineText) ID(value string) multilineText {
	return m.set("id", value)
}

// MaxRows sets the maximum number of rows
func (m multilineText) MaxRows(value string) multilineText {
	return m.set("maxRows", value)
}

// OnEvent sets the event configuration
func (m multilineText) OnEvent(value any) multilineText {
	return m.set("onEvent", value)
}

// Static sets the static display state
func (m multilineText) Static(value bool) multilineText {
	return m.set("static", value)
}

// StaticClassName sets the static display class name
func (m multilineText) StaticClassName(value string) multilineText {
	return m.set("staticClassName", value)
}

// StaticInputClassName sets the static input class name
func (m multilineText) StaticInputClassName(value string) multilineText {
	return m.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label class name
func (m multilineText) StaticLabelClassName(value string) multilineText {
	return m.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (m multilineText) StaticOn(value string) multilineText {
	return m.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder
func (m multilineText) StaticPlaceholder(value string) multilineText {
	return m.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (m multilineText) StaticSchema(value string) multilineText {
	return m.set("staticSchema", value)
}

// Style sets the component style
func (m multilineText) Style(value any) multilineText {
	return m.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (m multilineText) TestIdBuilder(value string) multilineText {
	return m.set("testIdBuilder", value)
}

// Testid sets the test ID
func (m multilineText) Testid(value string) multilineText {
	return m.set("testid", value)
}

// Text sets the text content
func (m multilineText) Text(value string) multilineText {
	return m.set("text", value)
}

// UseMobileUI sets the mobile UI state
func (m multilineText) UseMobileUI(value bool) multilineText {
	return m.set("useMobileUI", value)
}

// Visible sets the visibility state
func (m multilineText) Visible(value bool) multilineText {
	return m.set("visible", value)
}

// VisibleOn sets the visibility expression
func (m multilineText) VisibleOn(value string) multilineText {
	return m.set("visibleOn", value)
}
