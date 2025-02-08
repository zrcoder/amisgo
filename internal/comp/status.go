package comp

import "github.com/zrcoder/amisgo/model"

// Status Document: https://aisuda.bce.baidu.com/amis/zh-CN/components/Status
type Status model.Schema

func NewStatus() Status {
	return Status{"type": "status"}
}

func (s Status) set(key string, value any) Status {
	s[key] = value
	return s
}

// ClassName sets the container CSS class name
// Can be a string or an object for conditional classes
// Example: {"red": "data.progress > 80", "blue": "data.progress > 60"}
func (s Status) ClassName(value string) Status {
	return s.set("className", value)
}

// Disabled sets whether the status is disabled
func (s Status) Disabled(value bool) Status {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to determine if the status should be disabled
func (s Status) DisabledOn(value string) Status {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration, can be ignored at runtime
func (s Status) EditorSetting(value string) Status {
	return s.set("editorSetting", value)
}

// Hidden sets whether the status is hidden
func (s Status) Hidden(value bool) Status {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to determine if the status should be hidden
func (s Status) HiddenOn(value string) Status {
	return s.set("hiddenOn", value)
}

// Id sets the unique component ID, mainly used for log collection
func (s Status) ID(value string) Status {
	return s.set("id", value)
}

// LabelMap sets the label mapping relationship
func (s Status) LabelMap(value string) Status {
	return s.set("labelMap", value)
}

// Map sets the status icon mapping relationship
func (s Status) Map(value string) Status {
	return s.set("map", value)
}

// OnEvent sets the event action configuration
func (s Status) OnEvent(value any) Status {
	return s.set("onEvent", value)
}

// Placeholder sets the placeholder
func (s Status) Placeholder(value string) Status {
	return s.set("placeholder", value)
}

// Source sets the field of the new version configuration mapping source
func (s Status) Source(value string) Status {
	return s.set("source", value)
}

// Static sets whether to display in static mode
func (s Status) Static(value bool) Status {
	return s.set("static", value)
}

// StaticClassName sets the CSS class for static display
func (s Status) StaticClassName(value string) Status {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class for static value display
func (s Status) StaticInputClassName(value string) Status {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class for static label display
func (s Status) StaticLabelClassName(value string) Status {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the status should be displayed in static mode
func (s Status) StaticOn(value string) Status {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder text for empty values in static display
func (s Status) StaticPlaceholder(value string) Status {
	return s.set("staticPlaceholder", value)
}

// StaticSchema
func (s Status) StaticSchema(value string) Status {
	return s.set("staticSchema", value)
}

// Style sets the component inline styles
func (s Status) Style(value any) Status {
	return s.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (s Status) TestIdBuilder(value string) Status {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID
func (s Status) Testid(value string) Status {
	return s.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles at component level
func (s Status) UseMobileUI(value bool) Status {
	return s.set("useMobileUI", value)
}

// Visible sets whether the status is visible
func (s Status) Visible(value bool) Status {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine if the status should be visible
func (s Status) VisibleOn(value string) Status {
	return s.set("visibleOn", value)
}
