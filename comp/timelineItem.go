package comp

// timelineItem represents a timeline item renderer

type timelineItem Schema

// NewTimelineItem creates a new TimelineItem instance
func NewTimelineItem() timelineItem {
	return timelineItem{}
}

func (t timelineItem) set(key string, value any) timelineItem {
	t[key] = value
	return t
}

// ClassName sets the CSS class name
func (t timelineItem) ClassName(value string) timelineItem {
	return t.set("className", value)
}

// Color sets the color of the timeline point
func (t timelineItem) Color(value string) timelineItem {
	return t.set("color", value)
}

// Detail sets the detailed content
func (t timelineItem) Detail(value string) timelineItem {
	return t.set("detail", value)
}

// DetailClassName sets the CSS class name for the detail
func (t timelineItem) DetailClassName(value string) timelineItem {
	return t.set("detailClassName", value)
}

// DetailCollapsedText sets the text when detail is collapsed
func (t timelineItem) DetailCollapsedText(value string) timelineItem {
	return t.set("detailCollapsedText", value)
}

// DetailExpandedText sets the text when detail is expanded
func (t timelineItem) DetailExpandedText(value string) timelineItem {
	return t.set("detailExpandedText", value)
}

// Disabled sets whether the item is disabled
func (t timelineItem) Disabled(value bool) timelineItem {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to determine if the item is disabled
func (t timelineItem) DisabledOn(value string) timelineItem {
	return t.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (t timelineItem) EditorSetting(value string) timelineItem {
	return t.set("editorSetting", value)
}

// Hidden sets whether the item is hidden
func (t timelineItem) Hidden(value bool) timelineItem {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to determine if the item is hidden
func (t timelineItem) HiddenOn(value string) timelineItem {
	return t.set("hiddenOn", value)
}

// Icon sets the icon
func (t timelineItem) Icon(value string) timelineItem {
	return t.set("icon", value)
}

// IconClassName sets the CSS class name for the icon
func (t timelineItem) IconClassName(value string) timelineItem {
	return t.set("iconClassName", value)
}

// ID sets the unique ID for the component
func (t timelineItem) ID(value string) timelineItem {
	return t.set("id", value)
}

// OnEvent sets the event action configuration
func (t timelineItem) OnEvent(value any) timelineItem {
	return t.set("onEvent", value)
}

// Static sets whether the item is displayed statically
func (t timelineItem) Static(value bool) timelineItem {
	return t.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (t timelineItem) StaticClassName(value string) timelineItem {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (t timelineItem) StaticInputClassName(value string) timelineItem {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (t timelineItem) StaticLabelClassName(value string) timelineItem {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the item is displayed statically
func (t timelineItem) StaticOn(value string) timelineItem {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (t timelineItem) StaticPlaceholder(value string) timelineItem {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (t timelineItem) StaticSchema(value string) timelineItem {
	return t.set("staticSchema", value)
}

// Style sets the component style
func (t timelineItem) Style(value any) timelineItem {
	return t.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (t timelineItem) TestIdBuilder(value string) timelineItem {
	return t.set("testIdBuilder", value)
}

// Testid sets the test ID
func (t timelineItem) Testid(value string) timelineItem {
	return t.set("testid", value)
}

// Time sets the time point
func (t timelineItem) Time(value string) timelineItem {
	return t.set("time", value)
}

// TimeClassName sets the CSS class name for the time point
func (t timelineItem) TimeClassName(value string) timelineItem {
	return t.set("timeClassName", value)
}

// Title sets the title of the time point
func (t timelineItem) Title(value any) timelineItem {
	return t.set("title", value)
}

// TitleClassName sets the CSS class name for the title
func (t timelineItem) TitleClassName(value string) timelineItem {
	return t.set("titleClassName", value)
}

// UseMobileUI sets whether to use mobile UI
func (t timelineItem) UseMobileUI(value bool) timelineItem {
	return t.set("useMobileUI", value)
}

// Visible sets whether the item is visible
func (t timelineItem) Visible(value bool) timelineItem {
	return t.set("visible", value)
}

// VisibleOn sets the expression to determine if the item is visible
func (t timelineItem) VisibleOn(value string) timelineItem {
	return t.set("visibleOn", value)
}
