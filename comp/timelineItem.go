package comp

// TimelineItem 时间轴项渲染器
//
// @author slowlyo
// @version 6.7.0
type TimelineItem Schema

// NewTimelineItem 创建一个新的 TimelineItem 实例
func NewTimelineItem() TimelineItem {
	return TimelineItem{}
}

func (t TimelineItem) set(key string, value interface{}) TimelineItem {
	t[key] = value
	return t
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t TimelineItem) ClassName(value string) TimelineItem {
	return t.set("className", value)
}

// Color 时间点圆圈颜色
func (t TimelineItem) Color(value string) TimelineItem {
	return t.set("color", value)
}

// Detail 详细内容
func (t TimelineItem) Detail(value string) TimelineItem {
	return t.set("detail", value)
}

// DetailClassName 节点详情的CSS类名（优先级高于统一配置的detailClassName）
func (t TimelineItem) DetailClassName(value string) TimelineItem {
	return t.set("detailClassName", value)
}

// DetailCollapsedText detail折叠时文案
func (t TimelineItem) DetailCollapsedText(value string) TimelineItem {
	return t.set("detailCollapsedText", value)
}

// DetailExpandedText detail展开时文案
func (t TimelineItem) DetailExpandedText(value string) TimelineItem {
	return t.set("detailExpandedText", value)
}

// Disabled 是否禁用
func (t TimelineItem) Disabled(value bool) TimelineItem {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (t TimelineItem) DisabledOn(value string) TimelineItem {
	return t.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t TimelineItem) EditorSetting(value string) TimelineItem {
	return t.set("editorSetting", value)
}

// Hidden 是否隐藏
func (t TimelineItem) Hidden(value bool) TimelineItem {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (t TimelineItem) HiddenOn(value string) TimelineItem {
	return t.set("hiddenOn", value)
}

// Icon 图标
func (t TimelineItem) Icon(value string) TimelineItem {
	return t.set("icon", value)
}

// IconClassName 图标的CSS类名
func (t TimelineItem) IconClassName(value string) TimelineItem {
	return t.set("iconClassName", value)
}

// ID 组件唯一 id，主要用于日志采集
func (t TimelineItem) ID(value string) TimelineItem {
	return t.set("id", value)
}

// OnEvent 事件动作配置
func (t TimelineItem) OnEvent(value string) TimelineItem {
	return t.set("onEvent", value)
}

// Static 是否静态展示
func (t TimelineItem) Static(value bool) TimelineItem {
	return t.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t TimelineItem) StaticClassName(value string) TimelineItem {
	return t.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t TimelineItem) StaticInputClassName(value string) TimelineItem {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t TimelineItem) StaticLabelClassName(value string) TimelineItem {
	return t.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (t TimelineItem) StaticOn(value string) TimelineItem {
	return t.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (t TimelineItem) StaticPlaceholder(value string) TimelineItem {
	return t.set("staticPlaceholder", value)
}

// StaticSchema
func (t TimelineItem) StaticSchema(value string) TimelineItem {
	return t.set("staticSchema", value)
}

// Style 组件样式
func (t TimelineItem) Style(value string) TimelineItem {
	return t.set("style", value)
}

// TestIdBuilder
func (t TimelineItem) TestIdBuilder(value string) TimelineItem {
	return t.set("testIdBuilder", value)
}

// Testid
func (t TimelineItem) Testid(value string) TimelineItem {
	return t.set("testid", value)
}

// Time 时间点
func (t TimelineItem) Time(value string) TimelineItem {
	return t.set("time", value)
}

// TimeClassName 节点时间的CSS类名（优先级高于统一配置的timeClassName）
func (t TimelineItem) TimeClassName(value string) TimelineItem {
	return t.set("timeClassName", value)
}

// Title 时间节点标题 (时间节点标题)
func (t TimelineItem) Title(value string) TimelineItem {
	return t.set("title", value)
}

// TitleClassName 节点标题的CSS类名（优先级高于统一配置的titleClassName）
func (t TimelineItem) TitleClassName(value string) TimelineItem {
	return t.set("titleClassName", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (t TimelineItem) UseMobileUI(value bool) TimelineItem {
	return t.set("useMobileUI", value)
}

// Visible 是否显示
func (t TimelineItem) Visible(value bool) TimelineItem {
	return t.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (t TimelineItem) VisibleOn(value string) TimelineItem {
	return t.set("visibleOn", value)
}
