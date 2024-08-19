package comp

// date 展示渲染器
type date schema

// Date 创建一个新的 Date 实例，并设置默认的 type
func Date() date {
	return make(date).set("type", "date")
}

func (d date) set(key string, value any) date {
	d[key] = value
	return d
}

// ClassName 容器 css 类名
func (d date) ClassName(value string) date {
	return d.set("className", value)
}

// Disabled 是否禁用
func (d date) Disabled(value bool) date {
	return d.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (d date) DisabledOn(value string) date {
	return d.set("disabledOn", value)
}

// DisplayFormat 展示的时间格式
func (d date) DisplayFormat(value string) date {
	return d.set("displayFormat", value)
}

// DisplayTimeZone 时区
func (d date) DisplayTimeZone(value string) date {
	return d.set("displayTimeZone", value)
}

// EditorSetting 编辑器配置
func (d date) EditorSetting(value string) date {
	return d.set("editorSetting", value)
}

// Format 展示的时间格式
func (d date) Format(value string) date {
	return d.set("format", value)
}

// FromNow 显示成相对时间
func (d date) FromNow(value bool) date {
	return d.set("fromNow", value)
}

// Hidden 是否隐藏
func (d date) Hidden(value bool) date {
	return d.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (d date) HiddenOn(value string) date {
	return d.set("hiddenOn", value)
}

// ID 组件唯一 id
func (d date) ID(value string) date {
	return d.set("id", value)
}

// OnEvent 事件动作配置
func (d date) OnEvent(value any) date {
	return d.set("onEvent", value)
}

// Placeholder 占位符
func (d date) Placeholder(value string) date {
	return d.set("placeholder", value)
}

// Static 是否静态展示
func (d date) Static(value bool) date {
	return d.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (d date) StaticClassName(value string) date {
	return d.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (d date) StaticInputClassName(value string) date {
	return d.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (d date) StaticLabelClassName(value string) date {
	return d.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (d date) StaticOn(value string) date {
	return d.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (d date) StaticPlaceholder(value string) date {
	return d.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (d date) StaticSchema(value string) date {
	return d.set("staticSchema", value)
}

// Style 组件样式
func (d date) Style(value any) date {
	return d.set("style", value)
}

// TestIdBuilder 测试 id 构建器
func (d date) TestIdBuilder(value string) date {
	return d.set("testIdBuilder", value)
}

// TestID 测试 id
func (d date) TestID(value string) date {
	return d.set("testid", value)
}

// UpdateFrequency 更新频率
func (d date) UpdateFrequency(value string) date {
	return d.set("updateFrequency", value)
}

// UseMobileUI 关闭移动端样式
func (d date) UseMobileUI(value bool) date {
	return d.set("useMobileUI", value)
}

// ValueFormat 值的时间格式
func (d date) ValueFormat(value string) date {
	return d.set("valueFormat", value)
}

// Visible 是否显示
func (d date) Visible(value bool) date {
	return d.set("visible", value)
}

// VisibleOn 是否显示表达式
func (d date) VisibleOn(value string) date {
	return d.set("visibleOn", value)
}
