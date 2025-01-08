package comp

import "github.com/zrcoder/amisgo/model"

// status Document: https://aisuda.bce.baidu.com/amis/zh-CN/components/status

type status model.Schema

// Status creates a new Status instance
func Status() status {
	return status{}.set("type", "status")
}

func (s status) set(key string, value any) status {
	s[key] = value
	return s
}

// ClassName sets the container CSS class name
// Can be a string or an object for conditional classes
// Example: {"red": "data.progress > 80", "blue": "data.progress > 60"}
func (s status) ClassName(value string) status {
	return s.set("className", value)
}

// Disabled sets whether the status is disabled
func (s status) Disabled(value bool) status {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to determine if the status should be disabled
func (s status) DisabledOn(value string) status {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, can be ignored at runtime
func (s status) EditorSetting(value string) status {
	return s.set("editorSetting", value)
}

// Hidden sets whether the status is hidden
func (s status) Hidden(value bool) status {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to determine if the status should be hidden
func (s status) HiddenOn(value string) status {
	return s.set("hiddenOn", value)
}

// Id sets the unique component ID, mainly used for log collection
func (s status) ID(value string) status {
	return s.set("id", value)
}

// LabelMap sets the label mapping relationship
func (s status) LabelMap(value string) status {
	return s.set("labelMap", value)
}

// Map sets the status icon mapping relationship
func (s status) Map(value string) status {
	return s.set("map", value)
}

// OnEvent sets the event action configuration
func (s status) OnEvent(value any) status {
	return s.set("onEvent", value)
}

// Placeholder sets the placeholder
func (s status) Placeholder(value string) status {
	return s.set("placeholder", value)
}

// Source sets the field of the new version configuration mapping source
func (s status) Source(value string) status {
	return s.set("source", value)
}

// Static sets whether to display in static mode
func (s status) Static(value bool) status {
	return s.set("static", value)
}

// StaticClassName sets the CSS class for static display
func (s status) StaticClassName(value string) status {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class for static value display
func (s status) StaticInputClassName(value string) status {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class for static label display
func (s status) StaticLabelClassName(value string) status {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the status should be displayed in static mode
func (s status) StaticOn(value string) status {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder text for empty values in static display
func (s status) StaticPlaceholder(value string) status {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s status) StaticSchema(value string) status {
	return s.set("staticSchema", value)
}

// Style sets the component inline styles
func (s status) Style(value any) status {
	return s.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (s status) TestIdBuilder(value string) status {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID
func (s status) Testid(value string) status {
	return s.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles at component level
func (s status) UseMobileUI(value bool) status {
	return s.set("useMobileUI", value)
}

// Visible sets whether the status is visible
func (s status) Visible(value bool) status {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine if the status should be visible
func (s status) VisibleOn(value string) status {
	return s.set("visibleOn", value)
}
