package comp

// chartRadios 图表单选框
type chartRadios schema

// ChartRadios 创建一个新的 ChartRadios 实例
func ChartRadios() chartRadios {
	return make(chartRadios).set("type", "chart-radios")
}

func (c chartRadios) set(key string, value interface{}) chartRadios {
	c[key] = value
	return c
}

// ChartValueField 图表数值字段名
func (c chartRadios) ChartValueField(value string) chartRadios {
	return c.set("chartValueField", value)
}

// Config 图表配置
func (c chartRadios) Config(value string) chartRadios {
	return c.set("config", value)
}

// ShowTooltipOnHighlight 高亮的时候是否显示 tooltip
func (c chartRadios) ShowTooltipOnHighlight(value bool) chartRadios {
	return c.set("showTooltipOnHighlight", value)
}
