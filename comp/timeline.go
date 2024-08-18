package comp

// Timeline 时间轴渲染器
//
// @author slowlyo
// @version 6.7.0
type Timeline Schema

// NewTimeline 创建一个新的 Timeline 实例
func NewTimeline() Timeline {
	return Timeline{}.set("type", "timeline")
}

func (t Timeline) set(key string, value interface{}) Timeline {
	t[key] = value
	return t
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t Timeline) ClassName(value string) Timeline {
	return t.set("className", value)
}

// DetailClassName 节点详情的CSS类名
func (t Timeline) DetailClassName(value string) Timeline {
	return t.set("detailClassName", value)
}

// Direction 展示方向 可选值: horizontal | vertical
func (t Timeline) Direction(value string) Timeline {
	return t.set("direction", value)
}

// Disabled 是否禁用
func (t Timeline) Disabled(value bool) Timeline {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (t Timeline) DisabledOn(value string) Timeline {
	return t.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t Timeline) EditorSetting(value string) Timeline {
	return t.set("editorSetting", value)
}

// Hidden 是否隐藏
func (t Timeline) Hidden(value bool) Timeline {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (t Timeline) HiddenOn(value string) Timeline {
	return t.set("hiddenOn", value)
}

// IconClassName 图标的CSS类名
func (t Timeline) IconClassName(value string) Timeline {
	return t.set("iconClassName", value)
}

// ID 组件唯一 id，主要用于日志采集
func (t Timeline) ID(value string) Timeline {
	return t.set("id", value)
}

// ItemTitleSchema 节点title自定一展示模板 (节点title自定一展示模板)
func (t Timeline) ItemTitleSchema(value string) Timeline {
	return t.set("itemTitleSchema", value)
}

// Items 节点数据
func (t Timeline) Items(value string) Timeline {
	return t.set("items", value)
}

// Mode 文字相对于时间轴展示方向 可选值: left | right | alternate
func (t Timeline) Mode(value string) Timeline {
	return t.set("mode", value)
}

// OnEvent 事件动作配置
func (t Timeline) OnEvent(value string) Timeline {
	return t.set("onEvent", value)
}

// Reverse 节点倒序
func (t Timeline) Reverse(value bool) Timeline {
	return t.set("reverse", value)
}

// Source API 或 数据映射
func (t Timeline) Source(value string) Timeline {
	return t.set("source", value)
}

// Static 是否静态展示
func (t Timeline) Static(value bool) Timeline {
	return t.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t Timeline) StaticClassName(value string) Timeline {
	return t.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t Timeline) StaticInputClassName(value string) Timeline {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (t Timeline) StaticLabelClassName(value string) Timeline {
	return t.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (t Timeline) StaticOn(value string) Timeline {
	return t.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (t Timeline) StaticPlaceholder(value string) Timeline {
	return t.set("staticPlaceholder", value)
}

// StaticSchema
func (t Timeline) StaticSchema(value string) Timeline {
	return t.set("staticSchema", value)
}

// Style 组件样式
func (t Timeline) Style(value string) Timeline {
	return t.set("style", value)
}

// TestIdBuilder
func (t Timeline) TestIdBuilder(value string) Timeline {
	return t.set("testIdBuilder", value)
}

// Testid
func (t Timeline) Testid(value string) Timeline {
	return t.set("testid", value)
}

// TimeClassName 节点时间的CSS类名
func (t Timeline) TimeClassName(value string) Timeline {
	return t.set("timeClassName", value)
}

// TitleClassName 节点标题的CSS类名
func (t Timeline) TitleClassName(value string) Timeline {
	return t.set("titleClassName", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (t Timeline) UseMobileUI(value bool) Timeline {
	return t.set("useMobileUI", value)
}

// Visible 是否显示
func (t Timeline) Visible(value bool) Timeline {
	return t.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (t Timeline) VisibleOn(value string) Timeline {
	return t.set("visibleOn", value)
}
