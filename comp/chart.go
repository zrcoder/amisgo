package comp

// Chart 图表渲染器
type Chart Schema

// NewChart 创建一个新的 Chart 实例
func NewChart() Chart {
	return make(Chart).set("type", "chart")
}

// set 设置属性值
func (c Chart) set(key string, value interface{}) Chart {
	c[key] = value
	return c
}

// API 图表配置接口 (图表配置接口)
func (c Chart) API(value string) Chart {
	return c.set("api", value)
}

// ChartTheme Chart 主题配置
func (c Chart) ChartTheme(value string) Chart {
	return c.set("chartTheme", value)
}

// ClassName 容器 css 类名
func (c Chart) ClassName(value string) Chart {
	return c.set("className", value)
}

// ClickAction 点击行为配置，可以用来满足下钻操作等。
func (c Chart) ClickAction(value string) Chart {
	return c.set("clickAction", value)
}

// Config 配置echart的config，支持数据映射。如果用了数据映射，为了同步更新，请设置 trackExpression
func (c Chart) Config(value string) Chart {
	return c.set("config", value)
}

// DataFilter 数据过滤器
func (c Chart) DataFilter(value string) Chart {
	return c.set("dataFilter", value)
}

// DisableDataMapping 默认开启 Config 中的数据映射，如果想关闭，请开启此功能。
func (c Chart) DisableDataMapping(value bool) Chart {
	return c.set("disableDataMapping", value)
}

// Disabled 是否禁用
func (c Chart) Disabled(value bool) Chart {
	return c.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (c Chart) DisabledOn(value string) Chart {
	return c.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (c Chart) EditorSetting(value string) Chart {
	return c.set("editorSetting", value)
}

// Height 高度设置
func (c Chart) Height(value string) Chart {
	return c.set("height", value)
}

// Hidden 是否隐藏
func (c Chart) Hidden(value bool) Chart {
	return c.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (c Chart) HiddenOn(value string) Chart {
	return c.set("hiddenOn", value)
}

// ID 组件唯一 id，主要用于日志采集
func (c Chart) ID(value string) Chart {
	return c.set("id", value)
}

// InitFetch 是否初始加载
func (c Chart) InitFetch(value bool) Chart {
	return c.set("initFetch", value)
}

// InitFetchOn 是否初始加载用表达式来配置
func (c Chart) InitFetchOn(value string) Chart {
	return c.set("initFetchOn", value)
}

// Interval 刷新时间
func (c Chart) Interval(value string) Chart {
	return c.set("interval", value)
}

// LoadBaiduMap 加载百度地图
func (c Chart) LoadBaiduMap(value bool) Chart {
	return c.set("loadBaiduMap", value)
}

// MapName 地图名称
func (c Chart) MapName(value string) Chart {
	return c.set("mapName", value)
}

// MapURL 获取 geo json 文件的地址
func (c Chart) MapURL(value string) Chart {
	return c.set("mapURL", value)
}

// Name 组件名字，这个名字可以用来定位，用于组件通信
func (c Chart) Name(value string) Chart {
	return c.set("name", value)
}

// OnEvent 事件动作配置
func (c Chart) OnEvent(value string) Chart {
	return c.set("onEvent", value)
}

// ReplaceChartOption 默认配置时追加的，如果更新配置想完全替换配置请配置为 true.
func (c Chart) ReplaceChartOption(value bool) Chart {
	return c.set("replaceChartOption", value)
}

// Source 数据源
func (c Chart) Source(value string) Chart {
	return c.set("source", value)
}

// Static 是否静态展示
func (c Chart) Static(value bool) Chart {
	return c.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (c Chart) StaticClassName(value string) Chart {
	return c.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (c Chart) StaticInputClassName(value string) Chart {
	return c.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (c Chart) StaticLabelClassName(value string) Chart {
	return c.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (c Chart) StaticOn(value string) Chart {
	return c.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (c Chart) StaticPlaceholder(value string) Chart {
	return c.set("staticPlaceholder", value)
}

// StaticSchema 静态展示表单项 Schema
func (c Chart) StaticSchema(value string) Chart {
	return c.set("staticSchema", value)
}

// Style style样式
func (c Chart) Style(value string) Chart {
	return c.set("style", value)
}

// TestIdBuilder 测试 ID 构建器
func (c Chart) TestIdBuilder(value string) Chart {
	return c.set("testIdBuilder", value)
}

// TestID 测试 ID
func (c Chart) TestID(value string) Chart {
	return c.set("testid", value)
}

// TrackExpression 跟踪表达式，如果这个表达式的运行结果发生变化了，则会更新 Echart
func (c Chart) TrackExpression(value string) Chart {
	return c.set("trackExpression", value)
}

// UnMountOnHidden 不可见的时候隐藏
func (c Chart) UnMountOnHidden(value bool) Chart {
	return c.set("unMountOnHidden", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (c Chart) UseMobileUI(value bool) Chart {
	return c.set("useMobileUI", value)
}

// Visible 是否显示
func (c Chart) Visible(value bool) Chart {
	return c.set("visible", value)
}

// VisibleOn 是否显示表达式
func (c Chart) VisibleOn(value string) Chart {
	return c.set("visibleOn", value)
}

// Width 宽度设置
func (c Chart) Width(value string) Chart {
	return c.set("width", value)
}
