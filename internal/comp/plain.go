package comp

import "github.com/zrcoder/amisgo/model"

// Plain is a Plain text renderer
type Plain model.Schema

func NewPlain() Plain {
	return Plain{"type": "plain"}
}

func (p Plain) set(key string, value any) Plain {
	p[key] = value
	return p
}

// ClassName sets the CSS class name
func (p Plain) ClassName(value string) Plain {
	return p.set("className", value)
}

// Disabled sets the disabled state
func (p Plain) Disabled(value bool) Plain {
	return p.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (p Plain) DisabledOn(value string) Plain {
	return p.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (p Plain) EditorSetting(value string) Plain {
	return p.set("editorSetting", value)
}

// Hidden sets the hidden state
func (p Plain) Hidden(value bool) Plain {
	return p.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (p Plain) HiddenOn(value string) Plain {
	return p.set("hiddenOn", value)
}

// ID sets the unique component ID
func (p Plain) ID(value string) Plain {
	return p.set("id", value)
}

// Inline sets the inline display state
func (p Plain) Inline(value bool) Plain {
	return p.set("inline", value)
}

// OnEvent sets the event configuration
func (p Plain) OnEvent(value any) Plain {
	return p.set("onEvent", value)
}

// Placeholder sets the placeholder text
func (p Plain) Placeholder(value string) Plain {
	return p.set("placeholder", value)
}

// Static sets the static display state
func (p Plain) Static(value bool) Plain {
	return p.set("static", value)
}

// StaticClassName sets the static display CSS class name
func (p Plain) StaticClassName(value string) Plain {
	return p.set("staticClassName", value)
}

// StaticInputClassName sets the static input CSS class name
func (p Plain) StaticInputClassName(value string) Plain {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label CSS class name
func (p Plain) StaticLabelClassName(value string) Plain {
	return p.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (p Plain) StaticOn(value string) Plain {
	return p.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder text
func (p Plain) StaticPlaceholder(value string) Plain {
	return p.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (p Plain) StaticSchema(value string) Plain {
	return p.set("staticSchema", value)
}

// Style sets the component style
func (p Plain) Style(value any) Plain {
	return p.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (p Plain) TestIdBuilder(value string) Plain {
	return p.set("testIdBuilder", value)
}

// Testid sets the test ID
func (p Plain) Testid(value string) Plain {
	return p.set("testid", value)
}

// Text sets the text content
func (p Plain) Text(value string) Plain {
	return p.set("text", value)
}

// Tpl sets the template content
func (p Plain) Tpl(value string) Plain {
	return p.set("tpl", value)
}

// UseMobileUI sets the mobile UI state
func (p Plain) UseMobileUI(value bool) Plain {
	return p.set("useMobileUI", value)
}

// Visible sets the visibility state
func (p Plain) Visible(value bool) Plain {
	return p.set("visible", value)
}

// VisibleOn sets the visibility expression
func (p Plain) VisibleOn(value string) Plain {
	return p.set("visibleOn", value)
}
