package comp

import "github.com/zrcoder/amisgo/schema"

// ChartRadios represents a chart-based radio button selection component
type ChartRadios schema.Schema

func NewChartRadios() ChartRadios {
	return ChartRadios{"type": "chart-radios"}
}

func (c ChartRadios) set(key string, value any) ChartRadios {
	c[key] = value
	return c
}

// ChartValueField sets the field name for numerical values in the chart
func (c ChartRadios) ChartValueField(value string) ChartRadios {
	return c.set("chartValueField", value)
}

// Config sets the configuration for the chart
func (c ChartRadios) Config(value string) ChartRadios {
	return c.set("config", value)
}

// ShowTooltipOnHighlight enables or disables tooltip display when an item is highlighted
func (c ChartRadios) ShowTooltipOnHighlight(value bool) ChartRadios {
	return c.set("showTooltipOnHighlight", value)
}
