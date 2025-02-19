package comp

import "github.com/zrcoder/amisgo/schema"

// SparkLine
type SparkLine  schema.Schema

func NewSparkLine() SparkLine {
	return SparkLine{"type": "sparkline"}
}

func (s SparkLine) set(key string, value any) SparkLine {
	s[key] = value
	return s
}

// ClassName sets the custom sparkLine class
func (s SparkLine) ClassName(value string) SparkLine {
	return s.set("className", value)
}

// ClickAction sets the click action for the sparkLine
func (s SparkLine) ClickAction(value string) SparkLine {
	return s.set("clickAction", value)
}

// Disabled sets whether the sparkLine is disabled
func (s SparkLine) Disabled(value bool) SparkLine {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to disable the sparkLine
func (s SparkLine) DisabledOn(value string) SparkLine {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration for the sparkLine
func (s SparkLine) EditorSetting(value string) SparkLine {
	return s.set("editorSetting", value)
}

// Height sets the height of the sparkLine
func (s SparkLine) Height(value string) SparkLine {
	return s.set("height", value)
}

// Hidden sets whether the sparkLine is hidden
func (s SparkLine) Hidden(value bool) SparkLine {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to hide the sparkLine
func (s SparkLine) HiddenOn(value string) SparkLine {
	return s.set("hiddenOn", value)
}

// Id sets the unique component ID for the sparkLine
func (s SparkLine) ID(value string) SparkLine {
	return s.set("id", value)
}

// Name sets the associated data variable for the sparkLine
func (s SparkLine) Name(value string) SparkLine {
	return s.set("name", value)
}

// OnEvent sets the event action configuration for the sparkLine
func (s SparkLine) OnEvent(value any) SparkLine {
	return s.set("onEvent", value)
}

// Placeholder sets the placeholder text for the sparkLine
func (s SparkLine) Placeholder(value string) SparkLine {
	return s.set("placeholder", value)
}

// Static sets whether to display the sparkLine in static mode
func (s SparkLine) Static(value bool) SparkLine {
	return s.set("static", value)
}

// StaticClassName sets the CSS class for static display of the sparkLine
func (s SparkLine) StaticClassName(value string) SparkLine {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class for static value display of the sparkLine
func (s SparkLine) StaticInputClassName(value string) SparkLine {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class for static label display of the sparkLine
func (s SparkLine) StaticLabelClassName(value string) SparkLine {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the sparkLine should be displayed in static mode
func (s SparkLine) StaticOn(value string) SparkLine {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder text for empty values in static display of the sparkLine
func (s SparkLine) StaticPlaceholder(value string) SparkLine {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display of the sparkLine
func (s SparkLine) StaticSchema(value string) SparkLine {
	return s.set("staticSchema", value)
}

// Style sets the inline styles for the sparkLine
func (s SparkLine) Style(value any) SparkLine {
	return s.set("style", value)
}

// TestIdBuilder sets the test ID builder for the sparkLine
func (s SparkLine) TestIdBuilder(value string) SparkLine {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID for the sparkLine
func (s SparkLine) Testid(value string) SparkLine {
	return s.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles at component level for the sparkLine
func (s SparkLine) UseMobileUI(value bool) SparkLine {
	return s.set("useMobileUI", value)
}

// Value sets the value for the sparkLine
func (s SparkLine) Value(value string) SparkLine {
	return s.set("value", value)
}

// Visible sets whether the sparkLine is visible
func (s SparkLine) Visible(value bool) SparkLine {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine if the sparkLine should be visible
func (s SparkLine) VisibleOn(value string) SparkLine {
	return s.set("visibleOn", value)
}

// Width sets the width of the sparkLine
func (s SparkLine) Width(value string) SparkLine {
	return s.set("width", value)
}
