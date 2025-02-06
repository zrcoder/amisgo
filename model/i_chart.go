package model

import (
	js "encoding/json"
)

type (
	ChartCfg  Schema
	ChartSeri Schema
	ChartAxis Schema
)

func NewChartConfig() ChartCfg {
	return ChartCfg{}
}

func NewChartSeries() ChartSeri {
	return ChartSeri{}
}

func NewChartAxis() ChartAxis {
	return ChartAxis{}
}

func (c ChartCfg) Set(key string, value any) ChartCfg {
	c[key] = value
	return c
}

func (c ChartSeri) Set(key string, value any) ChartSeri {
	c[key] = value
	return c
}

func (c ChartAxis) Set(key string, value any) ChartAxis {
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
	return c.Set("title", value)
}

func (c ChartCfg) Tooltip(value any) ChartCfg {
	return c.Set("tooltip", value)
}

func (c ChartCfg) Polar(value any) ChartCfg {
	return c.Set("polar", value)
}

func (c ChartCfg) Radar(value any) ChartCfg {
	return c.Set("radar", value)
}

func (c ChartCfg) XAxis(value ChartAxis) ChartCfg {
	return c.Set("xAxis", value)
}

func (c ChartCfg) YAxis(value ChartAxis) ChartCfg {
	return c.Set("yAxis", value)
}

func (c ChartCfg) RadiusAxis(value ChartAxis) ChartCfg {
	return c.Set("radiusAxis", value)
}

func (c ChartCfg) AngleAxis(value ChartAxis) ChartCfg {
	return c.Set("angleAxis", value)
}

func (c ChartCfg) Series(value ...ChartSeri) ChartCfg {
	return c.Set("series", value)
}

func (c ChartCfg) AnimationDuration(value int) ChartCfg {
	return c.Set("animationDuration", value)
}

func (c ChartCfg) Legend(value any) ChartCfg {
	return c.Set("legend", value)
}

func (c ChartSeri) Name(value string) ChartSeri {
	return c.Set("name", value)
}

// Type line | bar | pie | scatter | radar ...
func (c ChartSeri) Type(value string) ChartSeri {
	return c.Set("type", value)
}

func (c ChartSeri) ItemStyle(value any) ChartSeri {
	return c.Set("itemStyle", value)
}

func (c ChartSeri) Radius(value string) ChartSeri {
	return c.Set("radius", value)
}

func (c ChartSeri) Data(value any) ChartSeri {
	return c.Set("data", value)
}

func (c ChartSeri) Label(value any) ChartSeri {
	return c.Set("label", value)
}

func (c ChartSeri) Smooth(value bool) ChartSeri {
	return c.Set("smooth", value)
}

func (c ChartSeri) LineStyle(value Schema) ChartSeri {
	return c.Set("lineStyle", value)
}

func (c ChartSeri) AreaStyle(value Schema) ChartSeri {
	return c.Set("areaStyle", value)
}

func (c ChartSeri) CoordinateSystem(value string) ChartSeri {
	return c.Set("coordinateSystem", value)
}

func (c ChartSeri) ShowSymbol(value bool) ChartSeri {
	return c.Set("showSymbol", value)
}

func (c ChartAxis) Type(value string) ChartAxis {
	return c.Set("type", value)
}

func (c ChartAxis) Data(value any) ChartAxis {
	return c.Set("data", value)
}

func (c ChartAxis) StartAngle(value float64) ChartAxis {
	return c.Set("startAngle", value)
}

// Min number | string
func (c ChartAxis) Min(value any) ChartAxis {
	return c.Set("min", value)
}

// Max number | string
func (c ChartAxis) Max(value any) ChartAxis {
	return c.Set("max", value)
}
