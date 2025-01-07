package comp

// password represents the amis password renderer

type password Schema

// Password creates a new Password instance
func Password() password {
	return password{}.set("type", "password")
}

// set sets a field value
func (p password) set(key string, value any) password {
	p[key] = value
	return p
}

// ClassName sets the container CSS class name
func (p password) ClassName(value string) password {
	return p.set("className", value)
}

// Disabled sets whether the component is disabled
func (p password) Disabled(value bool) password {
	return p.set("disabled", value)
}

// DisabledOn sets the expression for disabling the component
func (p password) DisabledOn(value string) password {
	return p.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (p password) EditorSetting(value string) password {
	return p.set("editorSetting", value)
}

// Hidden sets whether the component is hidden
func (p password) Hidden(value bool) password {
	return p.set("hidden", value)
}

// HiddenOn sets the expression for hiding the component
func (p password) HiddenOn(value string) password {
	return p.set("hiddenOn", value)
}

// ID sets the unique component ID
func (p password) ID(value string) password {
	return p.set("id", value)
}

// MosaicText sets the text for mosaic mode
func (p password) MosaicText(value string) password {
	return p.set("mosaicText", value)
}

// OnEvent sets the event action configuration
func (p password) OnEvent(value any) password {
	return p.set("onEvent", value)
}

// Static sets whether the component is statically displayed
func (p password) Static(value bool) password {
	return p.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (p password) StaticClassName(value string) password {
	return p.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (p password) StaticInputClassName(value string) password {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (p password) StaticLabelClassName(value string) password {
	return p.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (p password) StaticOn(value string) password {
	return p.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (p password) StaticPlaceholder(value string) password {
	return p.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display mode
func (p password) StaticSchema(value string) password {
	return p.set("staticSchema", value)
}

// Style sets the component style
func (p password) Style(value any) password {
	return p.set("style", value)
}

// TestIdBuilder sets the custom test ID builder
func (p password) TestIdBuilder(value string) password {
	return p.set("testIdBuilder", value)
}

// Testid sets the test ID
func (p password) Testid(value string) password {
	return p.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI styles
func (p password) UseMobileUI(value bool) password {
	return p.set("useMobileUI", value)
}

// Visible sets whether the component is visible
func (p password) Visible(value bool) password {
	return p.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (p password) VisibleOn(value string) password {
	return p.set("visibleOn", value)
}
