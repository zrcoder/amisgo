package comp

// DateRange 展示渲染器
type DateRange Schema

// NewDateRange 创建一个新的 DateRange 实例，并设置默认的 type
func NewDateRange() DateRange {
	return make(DateRange).set("type", "date-range")
}

func (d DateRange) set(key string, value interface{}) DateRange {
	d[key] = value
	return d
}

// ClassName 容器 css 类名
func (d DateRange) ClassName(value string) DateRange {
	return d.set("className", value)
}

// Connector 连接符
func (d DateRange) Connector(value string) DateRange {
	return d.set("connector", value)
}

// Delimiter 分割符
func (d DateRange) Delimiter(value string) DateRange {
	return d.set("delimiter", value)
}

// Disabled 是否禁用
func (d DateRange) Disabled(value bool) DateRange {
	return d.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (d DateRange) DisabledOn(value string) DateRange {
	return d.set("disabledOn", value)
}

// DisplayFormat 展示的时间格式
func (d DateRange) DisplayFormat(value string) DateRange {
	return d.set("displayFormat", value)
}

// EditorSetting 编辑器配置
func (d DateRange) EditorSetting(value string) DateRange {
	return d.set("editorSetting", value)
}

// Format 展示的时间格式
func (d DateRange) Format(value string) DateRange {
	return d.set("format", value)
}

// Hidden 是否隐藏
func (d DateRange) Hidden(value bool) DateRange {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (d DateRange) HiddenOn(value string) DateRange {
	return d.set("hiddenOn", value)
}

// ID 组件唯一 id
func (d DateRange) ID(value string) DateRange {
	return d.set("id", value)
}

// OnEvent 事件动作配置
func (d DateRange) OnEvent(value string) DateRange {
	return d.set("onEvent", value)
}

// Static 是否静态展示
func (d DateRange) Static(value bool) DateRange {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (d DateRange) StaticClassName(value string) DateRange {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (d DateRange) StaticInputClassName(value string) DateRange {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (d DateRange) StaticLabelClassName(value string) DateRange {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (d DateRange) StaticOn(value string) DateRange {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d DateRange) StaticPlaceholder(value string) DateRange {
	return d.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (d DateRange) StaticSchema(value string) DateRange {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d DateRange) Style(value string) DateRange {
	return d.set("style", value)
}

// TestIdBuilder 测试 id 构建器
func (d DateRange) TestIdBuilder(value string) DateRange {
	return d.set("testIdBuilder", value)
}

// TestID 测试 id
func (d DateRange) TestID(value string) DateRange {
	return d.set("testid", value)
}

// UseMobileUI 关闭移动端样式
func (d DateRange) UseMobileUI(value bool) DateRange {
	return d.set("useMobileUI", value)
}

// ValueFormat 值的时间格式
func (d DateRange) ValueFormat(value string) DateRange {
	return d.set("valueFormat", value)
}

// Visible 是否显示
func (d DateRange) Visible(value bool) DateRange {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式
func (d DateRange) VisibleOn(value string) DateRange {
	return d.set("visibleOn", value)
}
