package model

// TimelineItem represents a timeline item renderer
type TimelineItem Schema

func NewTimelineItem() TimelineItem {
	return TimelineItem{}
}

func (t TimelineItem) set(key string, value any) TimelineItem {
	t[key] = value
	return t
}

// ClassName sets the CSS class name
func (t TimelineItem) ClassName(value string) TimelineItem {
	return t.set("className", value)
}

// Color sets the color of the timeline point
func (t TimelineItem) Color(value string) TimelineItem {
	return t.set("color", value)
}

// Detail sets the detailed content
func (t TimelineItem) Detail(value string) TimelineItem {
	return t.set("detail", value)
}

// DetailClassName sets the CSS class name for the detail
func (t TimelineItem) DetailClassName(value string) TimelineItem {
	return t.set("detailClassName", value)
}

// DetailCollapsedText sets the text when detail is collapsed
func (t TimelineItem) DetailCollapsedText(value string) TimelineItem {
	return t.set("detailCollapsedText", value)
}

// DetailExpandedText sets the text when detail is expanded
func (t TimelineItem) DetailExpandedText(value string) TimelineItem {
	return t.set("detailExpandedText", value)
}

// Disabled sets whether the item is disabled
func (t TimelineItem) Disabled(value bool) TimelineItem {
	return t.set("disabled", value)
}

// DisabledOn sets the expression to determine if the item is disabled
func (t TimelineItem) DisabledOn(value string) TimelineItem {
	return t.set("disabledOn", value)
}

// EditorSetting sets the editor configuration
func (t TimelineItem) EditorSetting(value string) TimelineItem {
	return t.set("editorSetting", value)
}

// Hidden sets whether the item is hidden
func (t TimelineItem) Hidden(value bool) TimelineItem {
	return t.set("hidden", value)
}

// HiddenOn sets the expression to determine if the item is hidden
func (t TimelineItem) HiddenOn(value string) TimelineItem {
	return t.set("hiddenOn", value)
}

// Icon sets the icon
func (t TimelineItem) Icon(value string) TimelineItem {
	return t.set("icon", value)
}

// IconClassName sets the CSS class name for the icon
func (t TimelineItem) IconClassName(value string) TimelineItem {
	return t.set("iconClassName", value)
}

// ID sets the unique ID for the component
func (t TimelineItem) ID(value string) TimelineItem {
	return t.set("id", value)
}

// OnEvent sets the event action configuration
func (t TimelineItem) OnEvent(value any) TimelineItem {
	return t.set("onEvent", value)
}

// Static sets whether the item is displayed statically
func (t TimelineItem) Static(value bool) TimelineItem {
	return t.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (t TimelineItem) StaticClassName(value string) TimelineItem {
	return t.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (t TimelineItem) StaticInputClassName(value string) TimelineItem {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (t TimelineItem) StaticLabelClassName(value string) TimelineItem {
	return t.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the item is displayed statically
func (t TimelineItem) StaticOn(value string) TimelineItem {
	return t.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder for static display
func (t TimelineItem) StaticPlaceholder(value string) TimelineItem {
	return t.set("staticPlaceholder", value)
}

// StaticSchema sets the static schema
func (t TimelineItem) StaticSchema(value string) TimelineItem {
	return t.set("staticSchema", value)
}

// Style sets the component style
func (t TimelineItem) Style(value any) TimelineItem {
	return t.set("style", value)
}

// TestIdBuilder sets the test ID builder
func (t TimelineItem) TestIdBuilder(value string) TimelineItem {
	return t.set("testIdBuilder", value)
}

// Testid sets the test ID
func (t TimelineItem) Testid(value string) TimelineItem {
	return t.set("testid", value)
}

// Time sets the time point
func (t TimelineItem) Time(value string) TimelineItem {
	return t.set("time", value)
}

// TimeClassName sets the CSS class name for the time point
func (t TimelineItem) TimeClassName(value string) TimelineItem {
	return t.set("timeClassName", value)
}

// Title sets the title of the time point
func (t TimelineItem) Title(value any) TimelineItem {
	return t.set("title", value)
}

// TitleClassName sets the CSS class name for the title
func (t TimelineItem) TitleClassName(value string) TimelineItem {
	return t.set("titleClassName", value)
}

// UseMobileUI sets whether to use mobile UI
func (t TimelineItem) UseMobileUI(value bool) TimelineItem {
	return t.set("useMobileUI", value)
}

// Visible sets whether the item is visible
func (t TimelineItem) Visible(value bool) TimelineItem {
	return t.set("visible", value)
}

// VisibleOn sets the expression to determine if the item is visible
func (t TimelineItem) VisibleOn(value string) TimelineItem {
	return t.set("visibleOn", value)
}
