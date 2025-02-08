package comp

import "github.com/zrcoder/amisgo/model"

// Wrapper represents a container renderer
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/Wrapper
type Wrapper model.Schema

func NewWrapper() Wrapper {
	return Wrapper{"type": "wrapper"}
}

func (w Wrapper) set(key string, value any) Wrapper {
	w[key] = value
	return w
}

// Body sets the content
func (w Wrapper) Body(value ...any) Wrapper {
	return w.set("body", value)
}

// ClassName sets the container CSS class name
func (w Wrapper) ClassName(value string) Wrapper {
	return w.set("className", value)
}

// Disabled sets whether the container is disabled
func (w Wrapper) Disabled(value bool) Wrapper {
	return w.set("disabled", value)
}

// DisabledOn sets the expression to determine if the container is disabled
func (w Wrapper) DisabledOn(value string) Wrapper {
	return w.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (w Wrapper) EditorSetting(value string) Wrapper {
	return w.set("editorSetting", value)
}

// Hidden sets whether the container is hidden
func (w Wrapper) Hidden(value bool) Wrapper {
	return w.set("hidden", value)
}

// HiddenOn sets the expression to determine if the container is hidden
func (w Wrapper) HiddenOn(value string) Wrapper {
	return w.set("hiddenOn", value)
}

// ID sets the unique component ID, mainly for logging
func (w Wrapper) ID(value string) Wrapper {
	return w.set("id", value)
}

// OnEvent sets the event action configuration
func (w Wrapper) OnEvent(value any) Wrapper {
	return w.set("onEvent", value)
}

// Size sets the size (xs | sm | md | lg | none)
func (w Wrapper) Size(value string) Wrapper {
	return w.set("size", value)
}

// Static sets whether the container is displayed statically
func (w Wrapper) Static(value bool) Wrapper {
	return w.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (w Wrapper) StaticClassName(value string) Wrapper {
	return w.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (w Wrapper) StaticInputClassName(value string) Wrapper {
	return w.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (w Wrapper) StaticLabelClassName(value string) Wrapper {
	return w.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the container is displayed statically
func (w Wrapper) StaticOn(value string) Wrapper {
	return w.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (w Wrapper) StaticPlaceholder(value string) Wrapper {
	return w.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (w Wrapper) StaticSchema(value string) Wrapper {
	return w.set("staticSchema", value)
}

// Style sets custom styles
func (w Wrapper) Style(value any) Wrapper {
	return w.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (w Wrapper) TestIdBuilder(value string) Wrapper {
	return w.set("testIdBuilder", value)
}

// Testid sets the test ID
func (w Wrapper) Testid(value string) Wrapper {
	return w.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI styles
func (w Wrapper) UseMobileUI(value bool) Wrapper {
	return w.set("useMobileUI", value)
}

// Visible sets whether the container is visible
func (w Wrapper) Visible(value bool) Wrapper {
	return w.set("visible", value)
}

// VisibleOn sets the expression to determine if the container is visible
func (w Wrapper) VisibleOn(value string) Wrapper {
	return w.set("visibleOn", value)
}

// Wrap sets whether the container should wrap
func (w Wrapper) Wrap(value bool) Wrapper {
	return w.set("wrap", value)
}
