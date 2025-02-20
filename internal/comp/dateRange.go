package comp

import "github.com/zrcoder/amisgo/schema"

// DateRange represents a date range renderer
type DateRange schema.Schema

func NewDateRange() DateRange {
	return DateRange{"type": "date-range"}
}

func (d DateRange) set(key string, value any) DateRange {
	d[key] = value
	return d
}

// ClassName sets the CSS class name for the container
func (d DateRange) ClassName(value string) DateRange {
	return d.set("className", value)
}

// Connector sets the connector for the date range
func (d DateRange) Connector(value string) DateRange {
	return d.set("connector", value)
}

// Delimiter sets the delimiter for the date range
func (d DateRange) Delimiter(value string) DateRange {
	return d.set("delimiter", value)
}

// Disabled sets whether the date range is disabled
func (d DateRange) Disabled(value bool) DateRange {
	return d.set("disabled", value)
}

// DisabledOn sets a conditional expression for disabling the date range
func (d DateRange) DisabledOn(value string) DateRange {
	return d.set("disabledOn", value)
}

// DisplayFormat sets the display format for the date range
func (d DateRange) DisplayFormat(value string) DateRange {
	return d.set("displayFormat", value)
}

// EditorSetting configures editor-specific settings
func (d DateRange) EditorSetting(value string) DateRange {
	return d.set("editorSetting", value)
}

// Format sets the format for the date range
func (d DateRange) Format(value string) DateRange {
	return d.set("format", value)
}

// Hidden sets whether the date range is hidden
func (d DateRange) Hidden(value bool) DateRange {
	return d.set("hidden", value)
}

// HiddenOn sets a conditional expression for hiding the date range
func (d DateRange) HiddenOn(value string) DateRange {
	return d.set("hiddenOn", value)
}

// ID sets the unique identifier for the date range component
func (d DateRange) ID(value string) DateRange {
	return d.set("id", value)
}

// OnEvent configures event-driven actions
func (d DateRange) OnEvent(value Event) DateRange {
	return d.set("onEvent", value)
}

// Static sets whether the date range is statically displayed
func (d DateRange) Static(value bool) DateRange {
	return d.set("static", value)
}

// StaticClassName sets the CSS class name for static display
func (d DateRange) StaticClassName(value string) DateRange {
	return d.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input display
func (d DateRange) StaticInputClassName(value string) DateRange {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (d DateRange) StaticLabelClassName(value string) DateRange {
	return d.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (d DateRange) StaticOn(value string) DateRange {
	return d.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (d DateRange) StaticPlaceholder(value string) DateRange {
	return d.set("staticPlaceholder", value)
}

// StaticSchema sets the schema.Schema for static display
func (d DateRange) StaticSchema(value string) DateRange {
	return d.set("staticSchema", value)
}

// Style sets custom inline styles
func (d DateRange) Style(value any) DateRange {
	return d.set("style", value)
}

// TestIdBuilder configures test ID generation
func (d DateRange) TestIdBuilder(value string) DateRange {
	return d.set("testIdBuilder", value)
}

// TestID sets a specific test identifier
func (d DateRange) TestID(value string) DateRange {
	return d.set("testid", value)
}

// UseMobileUI sets whether to use mobile UI
func (d DateRange) UseMobileUI(value bool) DateRange {
	return d.set("useMobileUI", value)
}

// ValueFormat sets the format for the date range value
func (d DateRange) ValueFormat(value string) DateRange {
	return d.set("valueFormat", value)
}

// Visible sets whether the date range is visible
func (d DateRange) Visible(value bool) DateRange {
	return d.set("visible", value)
}

// VisibleOn sets a conditional expression for visibility
func (d DateRange) VisibleOn(value string) DateRange {
	return d.set("visibleOn", value)
}
