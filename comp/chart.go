package comp

import (
	"github.com/zrcoder/amisgo/internal/servermux"
	"github.com/zrcoder/amisgo/model"
)

// Chart represents a chart renderer component for data visualization
type chart model.Schema

// Chart creates a new Chart instance
func Chart() chart {
	return chart{"type": "chart"}
}

func (c chart) set(key string, value any) chart {
	c[key] = value
	return c
}

// API sets the configuration endpoint for chart data retrieval
func (c chart) API(value string) chart {
	return c.set("api", value)
}

// GetData sets a custom method for retrieving chart data
func (c chart) GetData(value func() (any, error)) chart {
	return c.API(servermux.ServeData(value))
}

// ChartTheme configures the visual theme for the chart
func (c chart) ChartTheme(value string) chart {
	return c.set("chartTheme", value)
}

// ClassName sets the CSS class name for the container element
func (c chart) ClassName(value string) chart {
	return c.set("className", value)
}

// ClickAction configures the behavior when chart elements are clicked, supporting drill-down operations
func (c chart) ClickAction(value string) chart {
	return c.set("clickAction", value)
}

// Config sets the ECharts configuration, supporting data mapping
// If data mapping is used, set TrackExpression to ensure synchronous updates
func (c chart) Config(value any) chart {
	return c.set("config", value)
}

// DataFilter applies a filter to the chart data
func (c chart) DataFilter(value string) chart {
	return c.set("dataFilter", value)
}

// DisableDataMapping disables data mapping in the Config
// By default, data mapping is enabled; set to true to disable
func (c chart) DisableDataMapping(value bool) chart {
	return c.set("disableDataMapping", value)
}

// Disabled enables or disables the entire chart component
func (c chart) Disabled(value bool) chart {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the chart
func (c chart) DisabledOn(value string) chart {
	return c.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings (can be ignored at runtime)
func (c chart) EditorSetting(value string) chart {
	return c.set("editorSetting", value)
}

// Height sets the chart's vertical dimension
func (c chart) Height(value string) chart {
	return c.set("height", value)
}

// Hidden controls the visibility of the chart component
func (c chart) Hidden(value bool) chart {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the chart
func (c chart) HiddenOn(value string) chart {
	return c.set("hiddenOn", value)
}

// ID sets a unique identifier for the component, primarily used for log collection
func (c chart) ID(value string) chart {
	return c.set("id", value)
}

// InitFetch determines whether to load data initially
func (c chart) InitFetch(value bool) chart {
	return c.set("initFetch", value)
}

// InitFetchOn sets a conditional expression to configure initial data loading
func (c chart) InitFetchOn(value string) chart {
	return c.set("initFetchOn", value)
}

// Interval sets the refresh time for the chart
func (c chart) Interval(value any) chart {
	return c.set("interval", value)
}

// LoadBaiduMap enables or disables Baidu Map loading
func (c chart) LoadBaiduMap(value bool) chart {
	return c.set("loadBaiduMap", value)
}

// MapName sets the name of the map
func (c chart) MapName(value string) chart {
	return c.set("mapName", value)
}

// MapURL specifies the address for retrieving geo JSON files
func (c chart) MapURL(value string) chart {
	return c.set("mapURL", value)
}

// Name sets the component name, which can be used for locating and component communication
func (c chart) Name(value string) chart {
	return c.set("name", value)
}

// OnEvent configures event-driven actions for the chart
func (c chart) OnEvent(value any) chart {
	return c.set("onEvent", value)
}

// ReplaceChartOption determines whether to completely replace the configuration
// By default, new configurations are appended; set to true to replace entirely
func (c chart) ReplaceChartOption(value bool) chart {
	return c.set("replaceChartOption", value)
}

// Source sets the data source for the chart
func (c chart) Source(value string) chart {
	return c.set("source", value)
}

// Static enables or disables static display mode
func (c chart) Static(value bool) chart {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display
func (c chart) StaticClassName(value string) chart {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display
func (c chart) StaticInputClassName(value string) chart {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c chart) StaticLabelClassName(value string) chart {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c chart) StaticOn(value string) chart {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c chart) StaticPlaceholder(value string) chart {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (c chart) StaticSchema(value string) chart {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c chart) Style(value any) chart {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c chart) TestIdBuilder(value string) chart {
	return c.set("testIdBuilder", value)
}

// TestID sets a specific test identifier
func (c chart) TestID(value string) chart {
	return c.set("testid", value)
}

// TrackExpression sets an expression to track changes
// If the expression's result changes, the EChart will be updated
func (c chart) TrackExpression(value string) chart {
	return c.set("trackExpression", value)
}

// UnMountOnHidden unmounts the component when it becomes invisible
func (c chart) UnMountOnHidden(value bool) chart {
	return c.set("unMountOnHidden", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c chart) UseMobileUI(value bool) chart {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the component
func (c chart) Visible(value bool) chart {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c chart) VisibleOn(value string) chart {
	return c.set("visibleOn", value)
}

// Width sets the chart's horizontal dimension
func (c chart) Width(value string) chart {
	return c.set("width", value)
}
