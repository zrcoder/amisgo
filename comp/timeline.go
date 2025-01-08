package comp

import "github.com/zrcoder/amisgo/model"

// timeline represents a timeline renderer

type timeline model.Schema

// NewTimeline creates a new Timeline instance
func NewTimeline() timeline {
	return timeline{}.set("type", "timeline")
}

func (t timeline) set(key string, value any) timeline {
	t[key] = value
	return t
}

// ClassName sets the CSS class name for the container
func (t timeline) ClassName(value string) timeline {
	return t.set("className", value)
}

// DetailClassName sets the CSS class name for node details
func (t timeline) DetailClassName(value string) timeline {
	return t.set("detailClassName", value)
}

// Direction sets the display direction (horizontal or vertical)
func (t timeline) Direction(value string) timeline {
	return t.set("direction", value)
}

// Disabled sets whether the timeline is disabled
func (t timeline) Disabled(value bool) timeline {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to determine if the timeline is disabled
func (t timeline) DisabledOn(value string) timeline {
	return t.set("disabledOn", value)
}

// EditorSetting sets the editor configuration (ignored at runtime)
func (t timeline) EditorSetting(value string) timeline {
	return t.set("editorSetting", value)
}

// Hidden sets whether the timeline is hidden
func (t timeline) Hidden(value bool) timeline {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to determine if the timeline is hidden
func (t timeline) HiddenOn(value string) timeline {
	return t.set("hiddenOn", value)
}

// IconClassName sets the CSS class name for icons
func (t timeline) IconClassName(value string) timeline {
	return t.set("iconClassName", value)
}

// ID sets the unique ID for the component
func (t timeline) ID(value string) timeline {
	return t.set("id", value)
}

// ItemTitleSchema sets the custom display template for node titles
func (t timeline) ItemTitleSchema(value string) timeline {
	return t.set("itemTitleSchema", value)
}

// Items sets the node data
func (t timeline) Items(value ...any) timeline {
	return t.set("items", value)
}

// Mode sets the text display direction relative to the timeline (left, right, alternate)
func (t timeline) Mode(value string) timeline {
	return t.set("mode", value)
}

// OnEvent sets the event action configuration
func (t timeline) OnEvent(value any) timeline {
	return t.set("onEvent", value)
}

// Reverse sets whether the nodes are in reverse order
func (t timeline) Reverse(value bool) timeline {
	return t.set("reverse", value)
}

// Source sets the API or data mapping
func (t timeline) Source(value string) timeline {
	return t.set("source", value)
}

// Static sets whether the timeline is displayed statically
func (t timeline) Static(value bool) timeline {
	return t.set("static", value)
}

// StaticClassName sets the CSS class name for static display form items
func (t timeline) StaticClassName(value string) timeline {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static display form item values
func (t timeline) StaticInputClassName(value string) timeline {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static display form item labels
func (t timeline) StaticLabelClassName(value string) timeline {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the timeline is displayed statically
func (t timeline) StaticOn(value string) timeline {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display empty values
func (t timeline) StaticPlaceholder(value string) timeline {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (t timeline) StaticSchema(value string) timeline {
	return t.set("staticSchema", value)
}

// Style sets the component style
func (t timeline) Style(value any) timeline {
	return t.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (t timeline) TestIdBuilder(value string) timeline {
	return t.set("testIdBuilder", value)
}

// Testid sets the test ID
func (t timeline) Testid(value string) timeline {
	return t.set("testid", value)
}

// TimeClassName sets the CSS class name for node times
func (t timeline) TimeClassName(value string) timeline {
	return t.set("timeClassName", value)
}

// TitleClassName sets the CSS class name for node titles
func (t timeline) TitleClassName(value string) timeline {
	return t.set("titleClassName", value)
}

// UseMobileUI sets whether to disable mobile styles at the component level
func (t timeline) UseMobileUI(value bool) timeline {
	return t.set("useMobileUI", value)
}

// Visible sets whether the timeline is visible
func (t timeline) Visible(value bool) timeline {
	return t.set("visible", value)
}

// VisibleOn sets the expression to determine if the timeline is visible
func (t timeline) VisibleOn(value string) timeline {
	return t.set("visibleOn", value)
}
