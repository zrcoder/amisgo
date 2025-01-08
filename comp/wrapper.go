package comp

import "github.com/zrcoder/amisgo/model"

// wrapper represents a container renderer
// Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/wrapper

type wrapper model.Schema

// Wrapper creates a new Wrapper instance
func Wrapper() wrapper {
	return wrapper{}.set("type", "wrapper")
}

func (w wrapper) set(key string, value any) wrapper {
	w[key] = value
	return w
}

// Body sets the content
func (w wrapper) Body(value ...any) wrapper {
	return w.set("body", value)
}

// ClassName sets the container CSS class name
func (w wrapper) ClassName(value string) wrapper {
	return w.set("className", value)
}

// Disabled sets whether the container is disabled
func (w wrapper) Disabled(value bool) wrapper {
	return w.set("disabled", value)
}

// DisabledOn sets the expression to determine if the container is disabled
func (w wrapper) DisabledOn(value string) wrapper {
	return w.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (w wrapper) EditorSetting(value string) wrapper {
	return w.set("editorSetting", value)
}

// Hidden sets whether the container is hidden
func (w wrapper) Hidden(value bool) wrapper {
	return w.set("hidden", value)
}

// HiddenOn sets the expression to determine if the container is hidden
func (w wrapper) HiddenOn(value string) wrapper {
	return w.set("hiddenOn", value)
}

// ID sets the unique component ID, mainly for logging
func (w wrapper) ID(value string) wrapper {
	return w.set("id", value)
}

// OnEvent sets the event action configuration
func (w wrapper) OnEvent(value any) wrapper {
	return w.set("onEvent", value)
}

// Size sets the size (xs | sm | md | lg | none)
func (w wrapper) Size(value string) wrapper {
	return w.set("size", value)
}

// Static sets whether the container is displayed statically
func (w wrapper) Static(value bool) wrapper {
	return w.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (w wrapper) StaticClassName(value string) wrapper {
	return w.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (w wrapper) StaticInputClassName(value string) wrapper {
	return w.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (w wrapper) StaticLabelClassName(value string) wrapper {
	return w.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the container is displayed statically
func (w wrapper) StaticOn(value string) wrapper {
	return w.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (w wrapper) StaticPlaceholder(value string) wrapper {
	return w.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (w wrapper) StaticSchema(value string) wrapper {
	return w.set("staticSchema", value)
}

// Style sets custom styles
func (w wrapper) Style(value any) wrapper {
	return w.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (w wrapper) TestIdBuilder(value string) wrapper {
	return w.set("testIdBuilder", value)
}

// Testid sets the test ID
func (w wrapper) Testid(value string) wrapper {
	return w.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI styles
func (w wrapper) UseMobileUI(value bool) wrapper {
	return w.set("useMobileUI", value)
}

// Visible sets whether the container is visible
func (w wrapper) Visible(value bool) wrapper {
	return w.set("visible", value)
}

// VisibleOn sets the expression to determine if the container is visible
func (w wrapper) VisibleOn(value string) wrapper {
	return w.set("visibleOn", value)
}

// Wrap sets whether the container should wrap
func (w wrapper) Wrap(value bool) wrapper {
	return w.set("wrap", value)
}
