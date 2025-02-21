package comp

import "github.com/zrcoder/amisgo/schema"

// SwitchContainer is a state container renderer. Documentation: https://aisuda.bce.baidu.com/amis/zh-CN/components/state-container
type SwitchContainer schema.Schema

func NewSwitchContainer() SwitchContainer {
	return SwitchContainer{"type": "switch-container"}
}

func (s SwitchContainer) set(key string, value any) SwitchContainer {
	s[key] = value
	return s
}

// ClassName sets the CSS class name for the container
func (s SwitchContainer) ClassName(value string) SwitchContainer {
	return s.set("className", value)
}

// Disabled sets whether the container is disabled
func (s SwitchContainer) Disabled(value bool) SwitchContainer {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to determine if the container is disabled
func (s SwitchContainer) DisabledOn(value string) SwitchContainer {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, ignored at runtime
func (s SwitchContainer) EditorSetting(value string) SwitchContainer {
	return s.set("editorSetting", value)
}

// Hidden sets whether the container is hidden
func (s SwitchContainer) Hidden(value bool) SwitchContainer {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to determine if the container is hidden
func (s SwitchContainer) HiddenOn(value string) SwitchContainer {
	return s.set("hiddenOn", value)
}

// ID sets the unique ID for the component, mainly for logging
func (s SwitchContainer) ID(value string) SwitchContainer {
	return s.set("id", value)
}

// Items sets the list of state items
func (s SwitchContainer) Items(value ...any) SwitchContainer {
	return s.set("items", value)
}

// OnEvent sets the event action configuration
func (s SwitchContainer) OnEvent(value Event) SwitchContainer {
	return s.set("onEvent", value)
}

// Static sets whether the container is displayed statically
func (s SwitchContainer) Static(value bool) SwitchContainer {
	return s.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (s SwitchContainer) StaticClassName(value string) SwitchContainer {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (s SwitchContainer) StaticInputClassName(value string) SwitchContainer {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (s SwitchContainer) StaticLabelClassName(value string) SwitchContainer {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the container is displayed statically
func (s SwitchContainer) StaticOn(value string) SwitchContainer {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (s SwitchContainer) StaticPlaceholder(value string) SwitchContainer {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (s SwitchContainer) StaticSchema(value string) SwitchContainer {
	return s.set("staticSchema", value)
}

// Style sets custom styles for the container
func (s SwitchContainer) Style(value any) SwitchContainer {
	return s.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (s SwitchContainer) TestIdBuilder(value string) SwitchContainer {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID
func (s SwitchContainer) Testid(value string) SwitchContainer {
	return s.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI styles
func (s SwitchContainer) UseMobileUI(value bool) SwitchContainer {
	return s.set("useMobileUI", value)
}

// Visible sets whether the container is visible
func (s SwitchContainer) Visible(value bool) SwitchContainer {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine if the container is visible
func (s SwitchContainer) VisibleOn(value string) SwitchContainer {
	return s.set("visibleOn", value)
}
