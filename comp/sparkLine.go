package comp

import "github.com/zrcoder/amisgo/model"

// sparkLine

type sparkLine model.Schema

// SparkLine creates a new SparkLine instance
func SparkLine() sparkLine {
	return sparkLine{"type": "sparkline"}
}

func (s sparkLine) set(key string, value any) sparkLine {
	s[key] = value
	return s
}

// ClassName sets the custom sparkLine class
func (s sparkLine) ClassName(value string) sparkLine {
	return s.set("className", value)
}

// ClickAction sets the click action for the sparkLine
func (s sparkLine) ClickAction(value string) sparkLine {
	return s.set("clickAction", value)
}

// Disabled sets whether the sparkLine is disabled
func (s sparkLine) Disabled(value bool) sparkLine {
	return s.set("disabled", value)
}

// DisabledOn sets the expression to disable the sparkLine
func (s sparkLine) DisabledOn(value string) sparkLine {
	return s.set("disabledOn", value)
}

// EditorSetting sets the editor configuration for the sparkLine
func (s sparkLine) EditorSetting(value string) sparkLine {
	return s.set("editorSetting", value)
}

// Height sets the height of the sparkLine
func (s sparkLine) Height(value string) sparkLine {
	return s.set("height", value)
}

// Hidden sets whether the sparkLine is hidden
func (s sparkLine) Hidden(value bool) sparkLine {
	return s.set("hidden", value)
}

// HiddenOn sets the expression to hide the sparkLine
func (s sparkLine) HiddenOn(value string) sparkLine {
	return s.set("hiddenOn", value)
}

// Id sets the unique component ID for the sparkLine
func (s sparkLine) ID(value string) sparkLine {
	return s.set("id", value)
}

// Name sets the associated data variable for the sparkLine
func (s sparkLine) Name(value string) sparkLine {
	return s.set("name", value)
}

// OnEvent sets the event action configuration for the sparkLine
func (s sparkLine) OnEvent(value any) sparkLine {
	return s.set("onEvent", value)
}

// Placeholder sets the placeholder text for the sparkLine
func (s sparkLine) Placeholder(value string) sparkLine {
	return s.set("placeholder", value)
}

// Static sets whether to display the sparkLine in static mode
func (s sparkLine) Static(value bool) sparkLine {
	return s.set("static", value)
}

// StaticClassName sets the CSS class for static display of the sparkLine
func (s sparkLine) StaticClassName(value string) sparkLine {
	return s.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class for static value display of the sparkLine
func (s sparkLine) StaticInputClassName(value string) sparkLine {
	return s.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class for static label display of the sparkLine
func (s sparkLine) StaticLabelClassName(value string) sparkLine {
	return s.set("staticLabelClassName", value)
}

// StaticOn sets the expression to determine if the sparkLine should be displayed in static mode
func (s sparkLine) StaticOn(value string) sparkLine {
	return s.set("staticOn", value)
}

// StaticPlaceholder sets the placeholder text for empty values in static display of the sparkLine
func (s sparkLine) StaticPlaceholder(value string) sparkLine {
	return s.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display of the sparkLine
func (s sparkLine) StaticSchema(value string) sparkLine {
	return s.set("staticSchema", value)
}

// Style sets the inline styles for the sparkLine
func (s sparkLine) Style(value any) sparkLine {
	return s.set("style", value)
}

// TestIdBuilder sets the test ID builder for the sparkLine
func (s sparkLine) TestIdBuilder(value string) sparkLine {
	return s.set("testIdBuilder", value)
}

// Testid sets the test ID for the sparkLine
func (s sparkLine) Testid(value string) sparkLine {
	return s.set("testid", value)
}

// UseMobileUI sets whether to disable mobile UI styles at component level for the sparkLine
func (s sparkLine) UseMobileUI(value bool) sparkLine {
	return s.set("useMobileUI", value)
}

// Value sets the value for the sparkLine
func (s sparkLine) Value(value string) sparkLine {
	return s.set("value", value)
}

// Visible sets whether the sparkLine is visible
func (s sparkLine) Visible(value bool) sparkLine {
	return s.set("visible", value)
}

// VisibleOn sets the expression to determine if the sparkLine should be visible
func (s sparkLine) VisibleOn(value string) sparkLine {
	return s.set("visibleOn", value)
}

// Width sets the width of the sparkLine
func (s sparkLine) Width(value string) sparkLine {
	return s.set("width", value)
}
