package comp

// ChartRadios 图表单选框
type ChartRadios Schema

// NewChartRadios 创建一个新的 ChartRadios 实例
func NewChartRadios() ChartRadios {
	return make(ChartRadios).set("type", "chart-radios")
}

// set 设置属性值
func (c ChartRadios) set(key string, value interface{}) ChartRadios {
	c[key] = value
	return c
}

// ChartValueField 图表数值字段名
func (c ChartRadios) ChartValueField(value string) ChartRadios {
	return c.set("chartValueField", value)
}

// Config 图表配置
func (c ChartRadios) Config(value string) ChartRadios {
	return c.set("config", value)
}

// ShowTooltipOnHighlight 高亮的时候是否显示 tooltip
func (c ChartRadios) ShowTooltipOnHighlight(value bool) ChartRadios {
	return c.set("showTooltipOnHighlight", value)
}
