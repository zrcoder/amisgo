package comp

import "github.com/zrcoder/amisgo/model"

// plain is a plain text renderer

type plain model.Schema

// Plain creates a new Plain instance
func Plain() plain {
	return plain{"type": "plain"}
}

func (p plain) set(key string, value any) plain {
	p[key] = value
	return p
}

// ClassName sets the CSS class name
func (p plain) ClassName(value string) plain {
	return p.set("className", value)
}

// Disabled sets the disabled state
func (p plain) Disabled(value bool) plain {
	return p.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (p plain) DisabledOn(value string) plain {
	return p.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (p plain) EditorSetting(value string) plain {
	return p.set("editorSetting", value)
}

// Hidden sets the hidden state
func (p plain) Hidden(value bool) plain {
	return p.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (p plain) HiddenOn(value string) plain {
	return p.set("hiddenOn", value)
}

// ID sets the unique component ID
func (p plain) ID(value string) plain {
	return p.set("id", value)
}

// Inline sets the inline display state
func (p plain) Inline(value bool) plain {
	return p.set("inline", value)
}

// OnEvent sets the event configuration
func (p plain) OnEvent(value any) plain {
	return p.set("onEvent", value)
}

// Placeholder sets the placeholder text
func (p plain) Placeholder(value string) plain {
	return p.set("placeholder", value)
}

// Static sets the static display state
func (p plain) Static(value bool) plain {
	return p.set("static", value)
}

// StaticClassName sets the static display CSS class name
func (p plain) StaticClassName(value string) plain {
	return p.set("staticClassName", value)
}

// StaticInputClassName sets the static input CSS class name
func (p plain) StaticInputClassName(value string) plain {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static label CSS class name
func (p plain) StaticLabelClassName(value string) plain {
	return p.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (p plain) StaticOn(value string) plain {
	return p.set("staticOn", value)
}

// StaticPlaceholder sets the static placeholder text
func (p plain) StaticPlaceholder(value string) plain {
	return p.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (p plain) StaticSchema(value string) plain {
	return p.set("staticSchema", value)
}

// Style sets the component style
func (p plain) Style(value any) plain {
	return p.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (p plain) TestIdBuilder(value string) plain {
	return p.set("testIdBuilder", value)
}

// Testid sets the test ID
func (p plain) Testid(value string) plain {
	return p.set("testid", value)
}

// Text sets the text content
func (p plain) Text(value string) plain {
	return p.set("text", value)
}

// Tpl sets the template content
func (p plain) Tpl(value string) plain {
	return p.set("tpl", value)
}

// UseMobileUI sets the mobile UI state
func (p plain) UseMobileUI(value bool) plain {
	return p.set("useMobileUI", value)
}

// Visible sets the visibility state
func (p plain) Visible(value bool) plain {
	return p.set("visible", value)
}

// VisibleOn sets the visibility expression
func (p plain) VisibleOn(value string) plain {
	return p.set("visibleOn", value)
}
