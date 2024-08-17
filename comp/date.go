package comp

// Date 展示渲染器
type Date BaseRenderer

// NewDate 创建一个新的 Date 实例，并设置默认的 type
func NewDate() Date {
	d := Date(make(BaseRenderer))
	return d.set("type", "date")
}

func (d Date) set(key string, value interface{}) Date {
	d[key] = value
	return d
}

// ClassName 容器 css 类名
func (d Date) ClassName(value string) Date {
	return d.set("className", value)
}

// Disabled 是否禁用
func (d Date) Disabled(value bool) Date {
	return d.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (d Date) DisabledOn(value string) Date {
	return d.set("disabledOn", value)
}

// DisplayFormat 展示的时间格式
func (d Date) DisplayFormat(value string) Date {
	return d.set("displayFormat", value)
}

// DisplayTimeZone 时区
func (d Date) DisplayTimeZone(value string) Date {
	return d.set("displayTimeZone", value)
}

// EditorSetting 编辑器配置
func (d Date) EditorSetting(value string) Date {
	return d.set("editorSetting", value)
}

// Format 展示的时间格式
func (d Date) Format(value string) Date {
	return d.set("format", value)
}

// FromNow 显示成相对时间
func (d Date) FromNow(value bool) Date {
	return d.set("fromNow", value)
}

// Hidden 是否隐藏
func (d Date) Hidden(value bool) Date {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (d Date) HiddenOn(value string) Date {
	return d.set("hiddenOn", value)
}

// ID 组件唯一 id
func (d Date) ID(value string) Date {
	return d.set("id", value)
}

// OnEvent 事件动作配置
func (d Date) OnEvent(value string) Date {
	return d.set("onEvent", value)
}

// Placeholder 占位符
func (d Date) Placeholder(value string) Date {
	return d.set("placeholder", value)
}

// Static 是否静态展示
func (d Date) Static(value bool) Date {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (d Date) StaticClassName(value string) Date {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (d Date) StaticInputClassName(value string) Date {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (d Date) StaticLabelClassName(value string) Date {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (d Date) StaticOn(value string) Date {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d Date) StaticPlaceholder(value string) Date {
	return d.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (d Date) StaticSchema(value string) Date {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d Date) Style(value string) Date {
	return d.set("style", value)
}

// TestIdBuilder 测试 id 构建器
func (d Date) TestIdBuilder(value string) Date {
	return d.set("testIdBuilder", value)
}

// TestID 测试 id
func (d Date) TestID(value string) Date {
	return d.set("testid", value)
}

// UpdateFrequency 更新频率
func (d Date) UpdateFrequency(value string) Date {
	return d.set("updateFrequency", value)
}

// UseMobileUI 关闭移动端样式
func (d Date) UseMobileUI(value bool) Date {
	return d.set("useMobileUI", value)
}

// ValueFormat 值的时间格式
func (d Date) ValueFormat(value string) Date {
	return d.set("valueFormat", value)
}

// Visible 是否显示
func (d Date) Visible(value bool) Date {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式
func (d Date) VisibleOn(value string) Date {
	return d.set("visibleOn", value)
}
