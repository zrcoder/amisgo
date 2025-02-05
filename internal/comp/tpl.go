package comp

import "github.com/zrcoder/amisgo/model"

// Tpl renderer
type Tpl model.Schema

func NewTpl() Tpl {
	return Tpl{"type": "tpl"}
}

func NewHtml() Tpl {
	return Tpl{"type": "html"}
}

func (t Tpl) set(key string, value any) Tpl {
	t[key] = value
	return t
}

// Badge sets the badge value
func (t Tpl) Badge(value string) Tpl {
	return t.set("badge", value)
}

// ClassName sets the container CSS class name
func (t Tpl) ClassName(value string) Tpl {
	return t.set("className", value)
}

// Disabled sets whether the component is disabled
func (t Tpl) Disabled(value bool) Tpl {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled
func (t Tpl) DisabledOn(value string) Tpl {
	return t.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (t Tpl) EditorSetting(value string) Tpl {
	return t.set("editorSetting", value)
}

// Hidden sets whether the component is hidden
func (t Tpl) Hidden(value bool) Tpl {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden
func (t Tpl) HiddenOn(value string) Tpl {
	return t.set("hiddenOn", value)
}

// Html sets the HTML content
func (t Tpl) Html(value string) Tpl {
	return t.set("html", value)
}

// ID sets the unique component ID
func (t Tpl) ID(value string) Tpl {
	return t.set("id", value)
}

// Inline sets whether the component is displayed inline
func (t Tpl) Inline(value bool) Tpl {
	return t.set("inline", value)
}

// OnEvent sets the event action configuration
func (t Tpl) OnEvent(value any) Tpl {
	return t.set("onEvent", value)
}

// Raw sets the raw content
func (t Tpl) Raw(value string) Tpl {
	return t.set("raw", value)
}

// Static sets whether the component is displayed statically
func (t Tpl) Static(value bool) Tpl {
	return t.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (t Tpl) StaticClassName(value string) Tpl {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (t Tpl) StaticInputClassName(value string) Tpl {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (t Tpl) StaticLabelClassName(value string) Tpl {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is displayed statically
func (t Tpl) StaticOn(value string) Tpl {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (t Tpl) StaticPlaceholder(value string) Tpl {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (t Tpl) StaticSchema(value string) Tpl {
	return t.set("staticSchema", value)
}

// Style sets the custom style
func (t Tpl) Style(value any) Tpl {
	return t.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (t Tpl) TestIdBuilder(value string) Tpl {
	return t.set("testIdBuilder", value)
}

// Testid sets the test ID
func (t Tpl) Testid(value string) Tpl {
	return t.set("testid", value)
}

// Text sets the text content
func (t Tpl) Text(value string) Tpl {
	return t.set("text", value)
}

// Tpl sets the template content
func (t Tpl) Tpl(value any) Tpl {
	return t.set("tpl", value)
}

// UseMobileUI sets whether to use mobile UI
func (t Tpl) UseMobileUI(value bool) Tpl {
	return t.set("useMobileUI", value)
}

// Visible sets whether the component is visible
func (t Tpl) Visible(value bool) Tpl {
	return t.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible
func (t Tpl) VisibleOn(value string) Tpl {
	return t.set("visibleOn", value)
}

// WrapperComponent sets the wrapper component type
func (t Tpl) WrapperComponent(value string) Tpl {
	return t.set("wrapperComponent", value)
}
