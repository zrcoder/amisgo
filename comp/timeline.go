package comp

// timeline 时间轴渲染器

type timeline schema

// NewTimeline 创建一个新的 Timeline 实例
func NewTimeline() timeline {
	return timeline{}.set("type", "timeline")
}

func (t timeline) set(key string, value any) timeline {
	t[key] = value
	return t
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t timeline) ClassName(value string) timeline {
	return t.set("className", value)
}

// DetailClassName 节点详情的CSS类名
func (t timeline) DetailClassName(value string) timeline {
	return t.set("detailClassName", value)
}

// Direction 展示方向 可选值: horizontal | vertical
func (t timeline) Direction(value string) timeline {
	return t.set("direction", value)
}

// Disabled 是否禁用
func (t timeline) Disabled(value bool) timeline {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (t timeline) DisabledOn(value string) timeline {
	return t.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t timeline) EditorSetting(value string) timeline {
	return t.set("editorSetting", value)
}

// Hidden 是否隐藏
func (t timeline) Hidden(value bool) timeline {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (t timeline) HiddenOn(value string) timeline {
	return t.set("hiddenOn", value)
}

// IconClassName 图标的CSS类名
func (t timeline) IconClassName(value string) timeline {
	return t.set("iconClassName", value)
}

// ID 组件唯一 id，主要用于日志采集
func (t timeline) ID(value string) timeline {
	return t.set("id", value)
}

// ItemTitleSchema 节点title自定一展示模板 (节点title自定一展示模板)
func (t timeline) ItemTitleSchema(value string) timeline {
	return t.set("itemTitleSchema", value)
}

// Items 节点数据
func (t timeline) Items(value ...any) timeline {
	return t.set("items", value)
}

// Mode 文字相对于时间轴展示方向 可选值: left | right | alternate
func (t timeline) Mode(value string) timeline {
	return t.set("mode", value)
}

// OnEvent 事件动作配置
func (t timeline) OnEvent(value any) timeline {
	return t.set("onEvent", value)
}

// Reverse 节点倒序
func (t timeline) Reverse(value bool) timeline {
	return t.set("reverse", value)
}

// Source API 或 数据映射
func (t timeline) Source(value string) timeline {
	return t.set("source", value)
}

// Static 是否静态展示
func (t timeline) Static(value bool) timeline {
	return t.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t timeline) StaticClassName(value string) timeline {
	return t.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t timeline) StaticInputClassName(value string) timeline {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t timeline) StaticLabelClassName(value string) timeline {
	return t.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (t timeline) StaticOn(value string) timeline {
	return t.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (t timeline) StaticPlaceholder(value string) timeline {
	return t.set("staticPlaceholder", value)
}

// StaticSchema
func (t timeline) StaticSchema(value string) timeline {
	return t.set("staticSchema", value)
}

// Style 组件样式
func (t timeline) Style(value any) timeline {
	return t.set("style", value)
}

// TestIdBuilder
func (t timeline) TestIdBuilder(value string) timeline {
	return t.set("testIdBuilder", value)
}

// Testid
func (t timeline) Testid(value string) timeline {
	return t.set("testid", value)
}

// TimeClassName 节点时间的CSS类名
func (t timeline) TimeClassName(value string) timeline {
	return t.set("timeClassName", value)
}

// TitleClassName 节点标题的CSS类名
func (t timeline) TitleClassName(value string) timeline {
	return t.set("titleClassName", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (t timeline) UseMobileUI(value bool) timeline {
	return t.set("useMobileUI", value)
}

// Visible 是否显示
func (t timeline) Visible(value bool) timeline {
	return t.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (t timeline) VisibleOn(value string) timeline {
	return t.set("visibleOn", value)
}
