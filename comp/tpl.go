package comp

import "github.com/zrcoder/amisgo/model"

// tpl renderer
type tpl model.Schema

// Tpl creates a new Tpl instance
func Tpl() tpl {
	return tpl{"type": "tpl"}
}

// Html creates a template with type set to "html".
func Html() tpl {
	return tpl{"type": "html"}
}

func (t tpl) set(key string, value any) tpl {
	t[key] = value
	return t
}

// Badge sets the badge value
func (t tpl) Badge(value string) tpl {
	return t.set("badge", value)
}

// ClassName sets the container CSS class name
func (t tpl) ClassName(value string) tpl {
	return t.set("className", value)
}

// Disabled sets whether the component is disabled
func (t tpl) Disabled(value bool) tpl {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to determine if the component is disabled
func (t tpl) DisabledOn(value string) tpl {
	return t.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (t tpl) EditorSetting(value string) tpl {
	return t.set("editorSetting", value)
}

// Hidden sets whether the component is hidden
func (t tpl) Hidden(value bool) tpl {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to determine if the component is hidden
func (t tpl) HiddenOn(value string) tpl {
	return t.set("hiddenOn", value)
}

// Html sets the HTML content
func (t tpl) Html(value string) tpl {
	return t.set("html", value)
}

// ID sets the unique component ID
func (t tpl) ID(value string) tpl {
	return t.set("id", value)
}

// Inline sets whether the component is displayed inline
func (t tpl) Inline(value bool) tpl {
	return t.set("inline", value)
}

// OnEvent sets the event action configuration
func (t tpl) OnEvent(value any) tpl {
	return t.set("onEvent", value)
}

// Raw sets the raw content
func (t tpl) Raw(value string) tpl {
	return t.set("raw", value)
}

// Static sets whether the component is displayed statically
func (t tpl) Static(value bool) tpl {
	return t.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (t tpl) StaticClassName(value string) tpl {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (t tpl) StaticInputClassName(value string) tpl {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (t tpl) StaticLabelClassName(value string) tpl {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the component is displayed statically
func (t tpl) StaticOn(value string) tpl {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (t tpl) StaticPlaceholder(value string) tpl {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (t tpl) StaticSchema(value string) tpl {
	return t.set("staticSchema", value)
}

// Style sets the custom style
func (t tpl) Style(value any) tpl {
	return t.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (t tpl) TestIdBuilder(value string) tpl {
	return t.set("testIdBuilder", value)
}

// Testid sets the test ID
func (t tpl) Testid(value string) tpl {
	return t.set("testid", value)
}

// Text sets the text content
func (t tpl) Text(value string) tpl {
	return t.set("text", value)
}

// Tpl sets the template content
func (t tpl) Tpl(value any) tpl {
	return t.set("tpl", value)
}

// UseMobileUI sets whether to use mobile UI
func (t tpl) UseMobileUI(value bool) tpl {
	return t.set("useMobileUI", value)
}

// Visible sets whether the component is visible
func (t tpl) Visible(value bool) tpl {
	return t.set("visible", value)
}

// VisibleOn sets the expression to determine if the component is visible
func (t tpl) VisibleOn(value string) tpl {
	return t.set("visibleOn", value)
}

// WrapperComponent sets the wrapper component type
func (t tpl) WrapperComponent(value string) tpl {
	return t.set("wrapperComponent", value)
}
