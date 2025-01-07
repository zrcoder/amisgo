package comp

// dateRange represents a date range renderer
type dateRange Schema

// DateRange creates a new DateRange instance with default type set to "date-range"
func DateRange() dateRange {
	return make(dateRange).set("type", "date-range")
}

func (d dateRange) set(key string, value any) dateRange {
	d[key] = value
	return d
}

// ClassName sets the CSS class name for the container
func (d dateRange) ClassName(value string) dateRange {
	return d.set("className", value)
}

// Connector sets the connector for the date range
func (d dateRange) Connector(value string) dateRange {
	return d.set("connector", value)
}

// Delimiter sets the delimiter for the date range
func (d dateRange) Delimiter(value string) dateRange {
	return d.set("delimiter", value)
}

// Disabled sets whether the date range is disabled
func (d dateRange) Disabled(value bool) dateRange {
	return d.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the date range
func (d dateRange) DisabledOn(value string) dateRange {
	return d.set("disabledOn", value)
}

// DisplayFormat sets the display format for the date range
func (d dateRange) DisplayFormat(value string) dateRange {
	return d.set("displayFormat", value)
}

// EditorSetting configures editor-specific settings
func (d dateRange) EditorSetting(value string) dateRange {
	return d.set("editorSetting", value)
}

// Format sets the format for the date range
func (d dateRange) Format(value string) dateRange {
	return d.set("format", value)
}

// Hidden sets whether the date range is hidden
func (d dateRange) Hidden(value bool) dateRange {
	return d.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the date range
func (d dateRange) HiddenOn(value string) dateRange {
	return d.set("hiddenOn", value)
}

// ID sets the unique identifier for the date range component
func (d dateRange) ID(value string) dateRange {
	return d.set("id", value)
}

// OnEvent configures event-driven actions
func (d dateRange) OnEvent(value any) dateRange {
	return d.set("onEvent", value)
}

// Static sets whether the date range is statically displayed
func (d dateRange) Static(value bool) dateRange {
	return d.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (d dateRange) StaticClassName(value string) dateRange {
	return d.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (d dateRange) StaticInputClassName(value string) dateRange {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (d dateRange) StaticLabelClassName(value string) dateRange {
	return d.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (d dateRange) StaticOn(value string) dateRange {
	return d.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (d dateRange) StaticPlaceholder(value string) dateRange {
	return d.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (d dateRange) StaticSchema(value string) dateRange {
	return d.set("staticSchema", value)
}

// Style sets custom inline styles
func (d dateRange) Style(value any) dateRange {
	return d.set("style", value)
}

// TestIdBuilder configures test ID generation
func (d dateRange) TestIdBuilder(value string) dateRange {
	return d.set("testIdBuilder", value)
}

// TestID sets a specific test identifier
func (d dateRange) TestID(value string) dateRange {
	return d.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI
func (d dateRange) UseMobileUI(value bool) dateRange {
	return d.set("useMobileUI", value)
}

// ValueFormat sets the format for the date range value
func (d dateRange) ValueFormat(value string) dateRange {
	return d.set("valueFormat", value)
}

// Visible sets whether the date range is visible
func (d dateRange) Visible(value bool) dateRange {
	return d.set("visible", value)
}

// VisibleOn sets a conditional expression for visibility
func (d dateRange) VisibleOn(value string) dateRange {
	return d.set("visibleOn", value)
}
