package comp

import "github.com/zrcoder/amisgo/schema"

// Password represents the amis Password renderer
type Password schema.Schema

func NewPassword() Password {
	return Password{"type": "password"}
}

// set sets a field value
func (p Password) set(key string, value any) Password {
	p[key] = value
	return p
}

// ClassName sets the container CSS class name
func (p Password) ClassName(value string) Password {
	return p.set("className", value)
}

// Disabled sets whether the component is disabled
func (p Password) Disabled(value bool) Password {
	return p.set("disabled", value)
}

// DisabledOn sets the expression for disabling the component
func (p Password) DisabledOn(value string) Password {
	return p.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (p Password) EditorSetting(value string) Password {
	return p.set("editorSetting", value)
}

// Hidden sets whether the component is hidden
func (p Password) Hidden(value bool) Password {
	return p.set("hidden", value)
}

// HiddenOn sets the expression for hiding the component
func (p Password) HiddenOn(value string) Password {
	return p.set("hiddenOn", value)
}

// ID sets the unique component ID
func (p Password) ID(value string) Password {
	return p.set("id", value)
}

// MosaicText sets the text for mosaic mode
func (p Password) MosaicText(value string) Password {
	return p.set("mosaicText", value)
}

// OnEvent sets the event action configuration
func (p Password) OnEvent(value Event) Password {
	return p.set("onEvent", value)
}

// Static sets whether the component is statically displayed
func (p Password) Static(value bool) Password {
	return p.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (p Password) StaticClassName(value string) Password {
	return p.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (p Password) StaticInputClassName(value string) Password {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (p Password) StaticLabelClassName(value string) Password {
	return p.set("staticLabelClassName", value)
}

// StaticOn sets the expression for static display
func (p Password) StaticOn(value string) Password {
	return p.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (p Password) StaticPlaceholder(value string) Password {
	return p.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display mode
func (p Password) StaticSchema(value string) Password {
	return p.set("staticSchema", value)
}

// Style sets the component style
func (p Password) Style(value any) Password {
	return p.set("style", value)
}

// TestIdBuilder sets the custom test ID builder
func (p Password) TestIdBuilder(value string) Password {
	return p.set("testIdBuilder", value)
}

// Testid sets the test ID
func (p Password) Testid(value string) Password {
	return p.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI styles
func (p Password) UseMobileUI(value bool) Password {
	return p.set("useMobileUI", value)
}

// Visible sets whether the component is visible
func (p Password) Visible(value bool) Password {
	return p.set("visible", value)
}

// VisibleOn sets the expression for visibility
func (p Password) VisibleOn(value string) Password {
	return p.set("visibleOn", value)
}
