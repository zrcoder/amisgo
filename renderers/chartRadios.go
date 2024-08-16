package renderers

// ChartRadios 图表单选框
type ChartRadios struct {
	*BaseRenderer
}

// NewChartRadios 创建一个新的 ChartRadios 实例
func NewChartRadios() *ChartRadios {
	c := &ChartRadios{
		BaseRenderer: NewBaseRenderer(),
	}
	c.set("type", "chart-radios")
	return c
}

// set 设置属性值
func (c *ChartRadios) set(key string, value interface{}) *ChartRadios {
	c.BaseRenderer.set(key, value)
	return c
}

// ChartValueField 图表数值字段名
func (c *ChartRadios) ChartValueField(value string) *ChartRadios {
	return c.set("chartValueField", value)
}

// Config 图表配置
func (c *ChartRadios) Config(value string) *ChartRadios {
	return c.set("config", value)
}

// ShowTooltipOnHighlight 高亮的时候是否显示 tooltip
func (c *ChartRadios) ShowTooltipOnHighlight(value bool) *ChartRadios {
	return c.set("showTooltipOnHighlight", value)
}
