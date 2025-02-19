package comp

import "github.com/zrcoder/amisgo/schema"

// Iframe renderer. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/Iframe
type Iframe schema.Schema

func NewIframe() Iframe {
	return Iframe{"type": "iframe"}
}

// Allow sets the allow attribute
func (i Iframe) Allow(value string) Iframe {
	return i.set("allow", value)
}

// ClassName sets the container CSS class name
func (i Iframe) ClassName(value string) Iframe {
	return i.set("className", value)
}

// Disabled sets the disabled state
func (i Iframe) Disabled(value bool) Iframe {
	return i.set("disabled", value)
}

// DisabledOn sets the disabled expression
func (i Iframe) DisabledOn(value string) Iframe {
	return i.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (i Iframe) EditorSetting(value string) Iframe {
	return i.set("editorSetting", value)
}

// Events sets the event responses
func (i Iframe) Events(value string) Iframe {
	return i.set("events", value)
}

// Height sets the height
func (i Iframe) Height(value string) Iframe {
	return i.set("height", value)
}

// Hidden sets the hidden state
func (i Iframe) Hidden(value bool) Iframe {
	return i.set("hidden", value)
}

// HiddenOn sets the hidden expression
func (i Iframe) HiddenOn(value string) Iframe {
	return i.set("hiddenOn", value)
}

// ID sets the unique component ID
func (i Iframe) ID(value string) Iframe {
	return i.set("id", value)
}

// Name sets the name
func (i Iframe) Name(value string) Iframe {
	return i.set("name", value)
}

// OnEvent sets the event action configuration
func (i Iframe) OnEvent(value any) Iframe {
	return i.set("onEvent", value)
}

// Referrerpolicy sets the referrer policy
func (i Iframe) Referrerpolicy(value string) Iframe {
	return i.set("referrerpolicy", value)
}

// Sandbox sets the sandbox attribute
func (i Iframe) Sandbox(value string) Iframe {
	return i.set("sandbox", value)
}

// Src sets the page URL
func (i Iframe) Src(value string) Iframe {
	return i.set("src", value)
}

// Static sets the static display state
func (i Iframe) Static(value bool) Iframe {
	return i.set("static", value)
}

// StaticClassName sets the static display form item class name
func (i Iframe) StaticClassName(value string) Iframe {
	return i.set("staticClassName", value)
}

// StaticInputClassName sets the static display form item value class name
func (i Iframe) StaticInputClassName(value string) Iframe {
	return i.set("staticInputClassName", value)
}

// StaticLabelClassName sets the static display form item label class name
func (i Iframe) StaticLabelClassName(value string) Iframe {
	return i.set("staticLabelClassName", value)
}

// StaticOn sets the static display expression
func (i Iframe) StaticOn(value string) Iframe {
	return i.set("staticOn", value)
}

// StaticPlaceholder sets the static display placeholder
func (i Iframe) StaticPlaceholder(value string) Iframe {
	return i.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema
func (i Iframe) StaticSchema(value string) Iframe {
	return i.set("staticSchema", value)
}

// Style sets the component style
func (i Iframe) Style(value any) Iframe {
	return i.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (i Iframe) TestIdBuilder(value string) Iframe {
	return i.set("testIdBuilder", value)
}

// Testid sets the test ID
func (i Iframe) Testid(value string) Iframe {
	return i.set("testid", value)
}

// UseMobileUI sets the mobile UI usage
func (i Iframe) UseMobileUI(value bool) Iframe {
	return i.set("useMobileUI", value)
}

// Visible sets the visibility
func (i Iframe) Visible(value bool) Iframe {
	return i.set("visible", value)
}

// VisibleOn sets the visibility expression
func (i Iframe) VisibleOn(value string) Iframe {
	return i.set("visibleOn", value)
}

// Width sets the width
func (i Iframe) Width(value string) Iframe {
	return i.set("width", value)
}

func (i Iframe) set(key string, value any) Iframe {
	i[key] = value
	return i
}
