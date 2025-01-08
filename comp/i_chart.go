package comp

import (
	js "encoding/json"

	"github.com/zrcoder/amisgo/model"
)

type (
	ChartCfg  model.Schema
	ChartSeri model.Schema
	ChartAxis model.Schema
)

func ChartConfig() ChartCfg {
	return ChartCfg{}
}

func ChartSeries() ChartSeri {
	return ChartSeri{}
}

func (c ChartCfg) set(key string, value any) ChartCfg {
	c[key] = value
	return c
}

func (c ChartSeri) set(key string, value any) ChartSeri {
	c[key] = value
	return c
}

func (c ChartCfg) Json() []byte {
	data, _ := js.Marshal(c)
	return data
}

func (c ChartCfg) JsonStr() string {
	return string(c.Json())
}

func (c ChartCfg) Title(value any) ChartCfg {
	return c.set("title", value)
}

func (c ChartCfg) Tooltip(value any) ChartCfg {
	return c.set("tooltip", value)
}

func (c ChartCfg) Polar(value any) ChartCfg {
	return c.set("polar", value)
}

func (c ChartCfg) Radar(value any) ChartCfg {
	return c.set("radar", value)
}

func (c ChartCfg) XAxis(value ChartAxis) ChartCfg {
	return c.set("xAxis", value)
}

func (c ChartCfg) YAxis(value ChartAxis) ChartCfg {
	return c.set("yAxis", value)
}

func (c ChartCfg) RadiusAxis(value ChartAxis) ChartCfg {
	return c.set("radiusAxis", value)
}

func (c ChartCfg) AngleAxis(value ChartAxis) ChartCfg {
	return c.set("angleAxis", value)
}

func (c ChartCfg) Series(value ...ChartSeri) ChartCfg {
	return c.set("series", value)
}

func (c ChartCfg) AnimationDuration(value int) ChartCfg {
	return c.set("animationDuration", value)
}

func (c ChartCfg) Legend(value any) ChartCfg {
	return c.set("legend", value)
}

func (c ChartSeri) Name(value string) ChartSeri {
	return c.set("name", value)
}

// Type line | bar | pie | scatter | radar ...
func (c ChartSeri) Type(value string) ChartSeri {
	return c.set("type", value)
}

func (c ChartSeri) ItemStyle(value any) ChartSeri {
	return c.set("itemStyle", value)
}

func (c ChartSeri) Radius(value string) ChartSeri {
	return c.set("radius", value)
}

func (c ChartSeri) Data(value any) ChartSeri {
	return c.set("data", value)
}

func (c ChartSeri) Label(value any) ChartSeri {
	return c.set("label", value)
}

func (c ChartSeri) Smooth(value bool) ChartSeri {
	return c.set("smooth", value)
}

func (c ChartSeri) LineStyle(value model.Schema) ChartSeri {
	return c.set("lineStyle", value)
}

func (c ChartSeri) AreaStyle(value model.Schema) ChartSeri {
	return c.set("areaStyle", value)
}

func (c ChartSeri) CoordinateSystem(value string) ChartSeri {
	return c.set("coordinateSystem", value)
}

func (c ChartSeri) ShowSymbol(value bool) ChartSeri {
	return c.set("showSymbol", value)
}
