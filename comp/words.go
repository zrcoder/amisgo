package comp

// words represents a text display component
type words Schema

// Words creates a new Words instance
func Words() words {
	return words{}.set("type", "words")
}

func (w words) set(key string, value any) words {
	w[key] = value
	return w
}

// ClassName sets the CSS class name
func (w words) ClassName(value string) words {
	return w.set("className", value)
}

// CollapseButton sets the collapse button text
func (w words) CollapseButton(value string) words {
	return w.set("collapseButton", value)
}

// CollapseButtonText sets the collapse button text
func (w words) CollapseButtonText(value string) words {
	return w.set("collapseButtonText", value)
}

// Delimiter sets the delimiter
func (w words) Delimiter(value string) words {
	return w.set("delimiter", value)
}

// Disabled sets the disabled state
func (w words) Disabled(value bool) words {
	return w.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (w words) DisabledOn(value string) words {
	return w.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (w words) EditorSetting(value string) words {
	return w.set("editorSetting", value)
}

// ExpendButton sets the expand button text
func (w words) ExpendButton(value string) words {
	return w.set("expendButton", value)
}

// ExpendButtonText sets the expand button text
func (w words) ExpendButtonText(value string) words {
	return w.set("expendButtonText", value)
}

// Hidden sets the hidden state
func (w words) Hidden(value bool) words {
	return w.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (w words) HiddenOn(value string) words {
	return w.set("hiddenOn", value)
}

// ID sets the unique component ID
func (w words) ID(value string) words {
	return w.set("id", value)
}

// InTag sets the tag display mode for arrays
func (w words) InTag(value string) words {
	return w.set("inTag", value)
}

// Limit sets the display limit
func (w words) Limit(value string) words {
	return w.set("limit", value)
}

// OnEvent sets the event configuration
func (w words) OnEvent(value any) words {
	return w.set("onEvent", value)
}

// Static sets the static display state
func (w words) Static(value bool) words {
	return w.set("static", value)
}

// StaticClassName sets the static display CSS class name
func (w words) StaticClassName(value string) words {
	return w.set("staticClassName", value)
}

// StaticInputClassName sets the static input CSS class name
func (w words) StaticInputClassName(value string) words {
	return w.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label CSS class name
func (w words) StaticLabelClassName(value string) words {
	return w.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (w words) StaticOn(value string) words {
	return w.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder
func (w words) StaticPlaceholder(value string) words {
	return w.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (w words) StaticSchema(value string) words {
	return w.set("staticSchema", value)
}

// Style sets the component style
func (w words) Style(value any) words {
	return w.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (w words) TestIdBuilder(value string) words {
	return w.set("testIdBuilder", value)
}

// Testid sets the test ID
func (w words) Testid(value string) words {
	return w.set("testid", value)
}

// UseMobileUI sets the mobile UI usage
func (w words) UseMobileUI(value bool) words {
	return w.set("useMobileUI", value)
}

// Visible sets the visibility state
func (w words) Visible(value bool) words {
	return w.set("visible", value)
}

// VisibleOn sets the visibility expression
func (w words) VisibleOn(value string) words {
	return w.set("visibleOn", value)
}

// Words sets the tags data
func (w words) Words(value string) words {
	return w.set("words", value)
}
