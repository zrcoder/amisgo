package comp

import (
	"net/http"

	"github.com/zrcoder/amisgo/internal/servemux"
	"github.com/zrcoder/amisgo/model"
)

// Chart represents a Chart renderer component for data visualization
type Chart model.Schema

func NewChart(mux *http.ServeMux) Chart {
	return Chart{"type": "chart", servemux.Key: mux}
}

func (c Chart) set(key string, value any) Chart {
	c[key] = value
	return c
}

// API sets the configuration endpoint for chart data retrieval
func (c Chart) API(value string) Chart {
	return c.set("api", value)
}

// GetData sets a custom method for retrieving chart data
func (c Chart) GetData(value func() (any, error)) Chart {
	return c.API(servemux.ServeData(c.mux(), value))
}

// ChartTheme configures the visual theme for the chart
func (c Chart) ChartTheme(value string) Chart {
	return c.set("chartTheme", value)
}

// ClassName sets the CSS class name for the container element
func (c Chart) ClassName(value string) Chart {
	return c.set("className", value)
}

// ClickAction configures the behavior when chart elements are clicked, supporting drill-down operations
func (c Chart) ClickAction(value string) Chart {
	return c.set("clickAction", value)
}

// Config sets the ECharts configuration, supporting data mapping
// If data mapping is used, set TrackExpression to ensure synchronous updates
func (c Chart) Config(value any) Chart {
	return c.set("config", value)
}

// DataFilter applies a filter to the chart data
func (c Chart) DataFilter(value string) Chart {
	return c.set("dataFilter", value)
}

// DisableDataMapping disables data mapping in the Config
// By default, data mapping is enabled; set to true to disable
func (c Chart) DisableDataMapping(value bool) Chart {
	return c.set("disableDataMapping", value)
}

// Disabled enables or disables the entire chart component
func (c Chart) Disabled(value bool) Chart {
	return c.set("disabled", value)
}

// DisabledOn sets a conditional expression to dynamically disable the chart
func (c Chart) DisabledOn(value string) Chart {
	return c.set("disabledOn", value)
}

// EditorSetting configures editor-specific settings (can be ignored at runtime)
func (c Chart) EditorSetting(value string) Chart {
	return c.set("editorSetting", value)
}

// Height sets the chart's vertical dimension
func (c Chart) Height(value string) Chart {
	return c.set("height", value)
}

// Hidden controls the visibility of the chart component
func (c Chart) Hidden(value bool) Chart {
	return c.set("hidden", value)
}

// HiddenOn sets a conditional expression to dynamically hide the chart
func (c Chart) HiddenOn(value string) Chart {
	return c.set("hiddenOn", value)
}

// ID sets a unique identifier for the component, primarily used for log collection
func (c Chart) ID(value string) Chart {
	return c.set("id", value)
}

// InitFetch determines whether to load data initially
func (c Chart) InitFetch(value bool) Chart {
	return c.set("initFetch", value)
}

// InitFetchOn sets a conditional expression to configure initial data loading
func (c Chart) InitFetchOn(value string) Chart {
	return c.set("initFetchOn", value)
}

// Interval sets the refresh time for the chart
func (c Chart) Interval(value any) Chart {
	return c.set("interval", value)
}

// LoadBaiduMap enables or disables Baidu Map loading
func (c Chart) LoadBaiduMap(value bool) Chart {
	return c.set("loadBaiduMap", value)
}

// MapName sets the name of the map
func (c Chart) MapName(value string) Chart {
	return c.set("mapName", value)
}

// MapURL specifies the address for retrieving geo JSON files
func (c Chart) MapURL(value string) Chart {
	return c.set("mapURL", value)
}

// Name sets the component name, which can be used for locating and component communication
func (c Chart) Name(value string) Chart {
	return c.set("name", value)
}

// OnEvent configures event-driven actions for the chart
func (c Chart) OnEvent(value any) Chart {
	return c.set("onEvent", value)
}

// ReplaceChartOption determines whether to completely replace the configuration
// By default, new configurations are appended; set to true to replace entirely
func (c Chart) ReplaceChartOption(value bool) Chart {
	return c.set("replaceChartOption", value)
}

// Source sets the data source for the chart
func (c Chart) Source(value string) Chart {
	return c.set("source", value)
}

// Static enables or disables static display mode
func (c Chart) Static(value bool) Chart {
	return c.set("static", value)
}

// StaticClassName sets the CSS class name for static form item display
func (c Chart) StaticClassName(value string) Chart {
	return c.set("staticClassName", value)
}

// StaticInputClassName sets the CSS class name for static input value display
func (c Chart) StaticInputClassName(value string) Chart {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName sets the CSS class name for static label display
func (c Chart) StaticLabelClassName(value string) Chart {
	return c.set("staticLabelClassName", value)
}

// StaticOn sets a conditional expression for static display
func (c Chart) StaticOn(value string) Chart {
	return c.set("staticOn", value)
}

// StaticPlaceholder defines a placeholder for empty static values
func (c Chart) StaticPlaceholder(value string) Chart {
	return c.set("staticPlaceholder", value)
}

// StaticSchema sets the schema for static display
func (c Chart) StaticSchema(value string) Chart {
	return c.set("staticSchema", value)
}

// Style sets custom inline styles for the component
func (c Chart) Style(value any) Chart {
	return c.set("style", value)
}

// TestIdBuilder configures test ID generation
func (c Chart) TestIdBuilder(value string) Chart {
	return c.set("testIdBuilder", value)
}

// TestID sets a specific test identifier
func (c Chart) TestID(value string) Chart {
	return c.set("testid", value)
}

// TrackExpression sets an expression to track changes
// If the expression's result changes, the EChart will be updated
func (c Chart) TrackExpression(value string) Chart {
	return c.set("trackExpression", value)
}

// UnMountOnHidden unmounts the component when it becomes invisible
func (c Chart) UnMountOnHidden(value bool) Chart {
	return c.set("unMountOnHidden", value)
}

// UseMobileUI enables or disables mobile-specific UI styling
func (c Chart) UseMobileUI(value bool) Chart {
	return c.set("useMobileUI", value)
}

// Visible controls the overall visibility of the component
func (c Chart) Visible(value bool) Chart {
	return c.set("visible", value)
}

// VisibleOn sets a conditional expression for component visibility
func (c Chart) VisibleOn(value string) Chart {
	return c.set("visibleOn", value)
}

// Width sets the chart's horizontal dimension
func (c Chart) Width(value string) Chart {
	return c.set("width", value)
}

func (c Chart) mux() *http.ServeMux {
	return c[servemux.Key].(*http.ServeMux)
}
