package comp

import "github.com/zrcoder/amisgo/model"

// ChartRadios represents a chart-based radio button selection component
type chartRadios model.Schema

// ChartRadios creates a new ChartRadios instance
func ChartRadios() chartRadios {
	return make(chartRadios).set("type", "chart-radios")
}

func (c chartRadios) set(key string, value any) chartRadios {
	c[key] = value
	return c
}

// ChartValueField sets the field name for numerical values in the chart
func (c chartRadios) ChartValueField(value string) chartRadios {
	return c.set("chartValueField", value)
}

// Config sets the configuration for the chart
func (c chartRadios) Config(value string) chartRadios {
	return c.set("config", value)
}

// ShowTooltipOnHighlight enables or disables tooltip display when an item is highlighted
func (c chartRadios) ShowTooltipOnHighlight(value bool) chartRadios {
	return c.set("showTooltipOnHighlight", value)
}
