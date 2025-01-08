package comp

import "github.com/zrcoder/amisgo/model"

// switchContainer is a state container renderer. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/state-container

type switchContainer model.Schema

// SwitchContainer creates a new SwitchContainer instance
func SwitchContainer() switchContainer {
	return switchContainer{}.set("type", "switch-container")
}

func (s switchContainer) set(key string, value any) switchContainer {
	s[key] = value
	return s
}

// ClassName sets the CSS class name for the container
func (s switchContainer) ClassName(value string) switchContainer {
	return s.set("className", value)
}

// Disabled sets whether the container is disabled
func (s switchContainer) Disabled(value bool) switchContainer {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to determine if the container is disabled
func (s switchContainer) DisabledOn(value string) switchContainer {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (s switchContainer) EditorSetting(value string) switchContainer {
	return s.set("editorSetting", value)
}

// Hidden sets whether the container is hidden
func (s switchContainer) Hidden(value bool) switchContainer {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to determine if the container is hidden
func (s switchContainer) HiddenOn(value string) switchContainer {
	return s.set("hiddenOn", value)
}

// ID sets the unique ID for the component, mainly for logging
func (s switchContainer) ID(value string) switchContainer {
	return s.set("id", value)
}

// Items sets the list of state items
func (s switchContainer) Items(value ...any) switchContainer {
	return s.set("items", value)
}

// OnEvent sets the event action configuration
func (s switchContainer) OnEvent(value any) switchContainer {
	return s.set("onEvent", value)
}

// Static sets whether the container is displayed statically
func (s switchContainer) Static(value bool) switchContainer {
	return s.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (s switchContainer) StaticClassName(value string) switchContainer {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (s switchContainer) StaticInputClassName(value string) switchContainer {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (s switchContainer) StaticLabelClassName(value string) switchContainer {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the container is displayed statically
func (s switchContainer) StaticOn(value string) switchContainer {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (s switchContainer) StaticPlaceholder(value string) switchContainer {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (s switchContainer) StaticSchema(value string) switchContainer {
	return s.set("staticSchema", value)
}

// Style sets custom styles for the container
func (s switchContainer) Style(value any) switchContainer {
	return s.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (s switchContainer) TestIdBuilder(value string) switchContainer {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID
func (s switchContainer) Testid(value string) switchContainer {
	return s.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI styles
func (s switchContainer) UseMobileUI(value bool) switchContainer {
	return s.set("useMobileUI", value)
}

// Visible sets whether the container is visible
func (s switchContainer) Visible(value bool) switchContainer {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine if the container is visible
func (s switchContainer) VisibleOn(value string) switchContainer {
	return s.set("visibleOn", value)
}
