package comp

// timelineItem 时间轴项渲染器
//
// @version 6.7.0
type timelineItem schema

// NewTimelineItem 创建一个新的 TimelineItem 实例
func NewTimelineItem() timelineItem {
	return timelineItem{}
}

func (t timelineItem) set(key string, value any) timelineItem {
	t[key] = value
	return t
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t timelineItem) ClassName(value string) timelineItem {
	return t.set("className", value)
}

// Color 时间点圆圈颜色
func (t timelineItem) Color(value string) timelineItem {
	return t.set("color", value)
}

// Detail 详细内容
func (t timelineItem) Detail(value string) timelineItem {
	return t.set("detail", value)
}

// DetailClassName 节点详情的CSS类名（优先级高于统一配置的detailClassName）
func (t timelineItem) DetailClassName(value string) timelineItem {
	return t.set("detailClassName", value)
}

// DetailCollapsedText detail折叠时文案
func (t timelineItem) DetailCollapsedText(value string) timelineItem {
	return t.set("detailCollapsedText", value)
}

// DetailExpandedText detail展开时文案
func (t timelineItem) DetailExpandedText(value string) timelineItem {
	return t.set("detailExpandedText", value)
}

// Disabled 是否禁用
func (t timelineItem) Disabled(value bool) timelineItem {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (t timelineItem) DisabledOn(value string) timelineItem {
	return t.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t timelineItem) EditorSetting(value string) timelineItem {
	return t.set("editorSetting", value)
}

// Hidden 是否隐藏
func (t timelineItem) Hidden(value bool) timelineItem {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (t timelineItem) HiddenOn(value string) timelineItem {
	return t.set("hiddenOn", value)
}

// Icon 图标
func (t timelineItem) Icon(value string) timelineItem {
	return t.set("icon", value)
}

// IconClassName 图标的CSS类名
func (t timelineItem) IconClassName(value string) timelineItem {
	return t.set("iconClassName", value)
}

// ID 组件唯一 id，主要用于日志采集
func (t timelineItem) ID(value string) timelineItem {
	return t.set("id", value)
}

// OnEvent 事件动作配置
func (t timelineItem) OnEvent(value any) timelineItem {
	return t.set("onEvent", value)
}

// Static 是否静态展示
func (t timelineItem) Static(value bool) timelineItem {
	return t.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t timelineItem) StaticClassName(value string) timelineItem {
	return t.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t timelineItem) StaticInputClassName(value string) timelineItem {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t timelineItem) StaticLabelClassName(value string) timelineItem {
	return t.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (t timelineItem) StaticOn(value string) timelineItem {
	return t.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (t timelineItem) StaticPlaceholder(value string) timelineItem {
	return t.set("staticPlaceholder", value)
}

// StaticSchema
func (t timelineItem) StaticSchema(value string) timelineItem {
	return t.set("staticSchema", value)
}

// Style 组件样式
func (t timelineItem) Style(value any) timelineItem {
	return t.set("style", value)
}

// TestIdBuilder
func (t timelineItem) TestIdBuilder(value string) timelineItem {
	return t.set("testIdBuilder", value)
}

// Testid
func (t timelineItem) Testid(value string) timelineItem {
	return t.set("testid", value)
}

// Time 时间点
func (t timelineItem) Time(value string) timelineItem {
	return t.set("time", value)
}

// TimeClassName 节点时间的CSS类名（优先级高于统一配置的timeClassName）
func (t timelineItem) TimeClassName(value string) timelineItem {
	return t.set("timeClassName", value)
}

// Title 时间节点标题 (时间节点标题)
func (t timelineItem) Title(value any) timelineItem {
	return t.set("title", value)
}

// TitleClassName 节点标题的CSS类名（优先级高于统一配置的titleClassName）
func (t timelineItem) TitleClassName(value string) timelineItem {
	return t.set("titleClassName", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (t timelineItem) UseMobileUI(value bool) timelineItem {
	return t.set("useMobileUI", value)
}

// Visible 是否显示
func (t timelineItem) Visible(value bool) timelineItem {
	return t.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (t timelineItem) VisibleOn(value string) timelineItem {
	return t.set("visibleOn", value)
}
