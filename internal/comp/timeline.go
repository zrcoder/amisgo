package comp

import "github.com/zrcoder/amisgo/schema"

// Timeline represents a Timeline renderer
type Timeline schema.Schema

func NewTimeline() Timeline {
	return Timeline{"type": "timeline"}
}

func (t Timeline) set(key string, value any) Timeline {
	t[key] = value
	return t
}

// ClassName sets the CSS class name for the container
func (t Timeline) ClassName(value string) Timeline {
	return t.set("className", value)
}

// DetailClassName sets the CSS class name for node details
func (t Timeline) DetailClassName(value string) Timeline {
	return t.set("detailClassName", value)
}

// Direction sets the display direction (horizontal or vertical)
func (t Timeline) Direction(value string) Timeline {
	return t.set("direction", value)
}

// Disabled sets whether the timeline is disabled
func (t Timeline) Disabled(value bool) Timeline {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to determine if the timeline is disabled
func (t Timeline) DisabledOn(value string) Timeline {
	return t.set("disabledOn", value)
}

// EditorSetting sets the editor configuration (ignored at runtime)
func (t Timeline) EditorSetting(value string) Timeline {
	return t.set("editorSetting", value)
}

// Hidden sets whether the timeline is hidden
func (t Timeline) Hidden(value bool) Timeline {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to determine if the timeline is hidden
func (t Timeline) HiddenOn(value string) Timeline {
	return t.set("hiddenOn", value)
}

// IconClassName sets the CSS class name for icons
func (t Timeline) IconClassName(value string) Timeline {
	return t.set("iconClassName", value)
}

// ID sets the unique ID for the component
func (t Timeline) ID(value string) Timeline {
	return t.set("id", value)
}

// ItemTitleSchema sets the custom display template for node titles
func (t Timeline) ItemTitleSchema(value string) Timeline {
	return t.set("itemTitleSchema", value)
}

// Items sets the node data
func (t Timeline) Items(value ...any) Timeline {
	return t.set("items", value)
}

// Mode sets the text display direction relative to the timeline (left, right, alternate)
func (t Timeline) Mode(value string) Timeline {
	return t.set("mode", value)
}

// OnEvent sets the event action configuration
func (t Timeline) OnEvent(value any) Timeline {
	return t.set("onEvent", value)
}

// Reverse sets whether the nodes are in reverse order
func (t Timeline) Reverse(value bool) Timeline {
	return t.set("reverse", value)
}

// Source sets the API or data mapping
func (t Timeline) Source(value string) Timeline {
	return t.set("source", value)
}

// Static sets whether the timeline is displayed statically
func (t Timeline) Static(value bool) Timeline {
	return t.set("static", value)
}

// StaticClassName sets the CSS class name for static display form items
func (t Timeline) StaticClassName(value string) Timeline {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static display form item values
func (t Timeline) StaticInputClassName(value string) Timeline {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static display form item labels
func (t Timeline) StaticLabelClassName(value string) Timeline {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the timeline is displayed statically
func (t Timeline) StaticOn(value string) Timeline {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display empty values
func (t Timeline) StaticPlaceholder(value string) Timeline {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema.Schema
func (t Timeline) StaticSchema(value string) Timeline {
	return t.set("staticSchema", value)
}

// Style sets the component style
func (t Timeline) Style(value any) Timeline {
	return t.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (t Timeline) TestIdBuilder(value string) Timeline {
	return t.set("testIdBuilder", value)
}

// Testid sets the test ID
func (t Timeline) Testid(value string) Timeline {
	return t.set("testid", value)
}

// TimeClassName sets the CSS class name for node times
func (t Timeline) TimeClassName(value string) Timeline {
	return t.set("timeClassName", value)
}

// TitleClassName sets the CSS class name for node titles
func (t Timeline) TitleClassName(value string) Timeline {
	return t.set("titleClassName", value)
}

// UseMobileUI sets whether to disable mobile styles at the component level
func (t Timeline) UseMobileUI(value bool) Timeline {
	return t.set("useMobileUI", value)
}

// Visible sets whether the timeline is visible
func (t Timeline) Visible(value bool) Timeline {
	return t.set("visible", value)
}

// VisibleOn sets the expression to determine if the timeline is visible
func (t Timeline) VisibleOn(value string) Timeline {
	return t.set("visibleOn", value)
}
