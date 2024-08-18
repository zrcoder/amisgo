package comp

// dateRange 展示渲染器
type dateRange schema

// DateRange 创建一个新的 DateRange 实例，并设置默认的 type
func DateRange() dateRange {
	return make(dateRange).set("type", "date-range")
}

func (d dateRange) set(key string, value interface{}) dateRange {
	d[key] = value
	return d
}

// ClassName 容器 css 类名
func (d dateRange) ClassName(value string) dateRange {
	return d.set("className", value)
}

// Connector 连接符
func (d dateRange) Connector(value string) dateRange {
	return d.set("connector", value)
}

// Delimiter 分割符
func (d dateRange) Delimiter(value string) dateRange {
	return d.set("delimiter", value)
}

// Disabled 是否禁用
func (d dateRange) Disabled(value bool) dateRange {
	return d.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (d dateRange) DisabledOn(value string) dateRange {
	return d.set("disabledOn", value)
}

// DisplayFormat 展示的时间格式
func (d dateRange) DisplayFormat(value string) dateRange {
	return d.set("displayFormat", value)
}

// EditorSetting 编辑器配置
func (d dateRange) EditorSetting(value string) dateRange {
	return d.set("editorSetting", value)
}

// Format 展示的时间格式
func (d dateRange) Format(value string) dateRange {
	return d.set("format", value)
}

// Hidden 是否隐藏
func (d dateRange) Hidden(value bool) dateRange {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (d dateRange) HiddenOn(value string) dateRange {
	return d.set("hiddenOn", value)
}

// ID 组件唯一 id
func (d dateRange) ID(value string) dateRange {
	return d.set("id", value)
}

// OnEvent 事件动作配置
func (d dateRange) OnEvent(value string) dateRange {
	return d.set("onEvent", value)
}

// Static 是否静态展示
func (d dateRange) Static(value bool) dateRange {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (d dateRange) StaticClassName(value string) dateRange {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (d dateRange) StaticInputClassName(value string) dateRange {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (d dateRange) StaticLabelClassName(value string) dateRange {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (d dateRange) StaticOn(value string) dateRange {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d dateRange) StaticPlaceholder(value string) dateRange {
	return d.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (d dateRange) StaticSchema(value string) dateRange {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d dateRange) Style(value string) dateRange {
	return d.set("style", value)
}

// TestIdBuilder 测试 id 构建器
func (d dateRange) TestIdBuilder(value string) dateRange {
	return d.set("testIdBuilder", value)
}

// TestID 测试 id
func (d dateRange) TestID(value string) dateRange {
	return d.set("testid", value)
}

// UseMobileUI 关闭移动端样式
func (d dateRange) UseMobileUI(value bool) dateRange {
	return d.set("useMobileUI", value)
}

// ValueFormat 值的时间格式
func (d dateRange) ValueFormat(value string) dateRange {
	return d.set("valueFormat", value)
}

// Visible 是否显示
func (d dateRange) Visible(value bool) dateRange {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式
func (d dateRange) VisibleOn(value string) dateRange {
	return d.set("visibleOn", value)
}
