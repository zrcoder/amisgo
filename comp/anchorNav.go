package comp

// AnchorNav 锚点导航渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/anchor-nav
type AnchorNav Schema

// NewAnchorNav 创建一个新的 AnchorNav 实例
func NewAnchorNav() AnchorNav {
	return make(AnchorNav).set("type", "anchor-nav")
}

func (a AnchorNav) set(key string, value interface{}) AnchorNav {
	a[key] = value
	return a
}

// 被激活（定位）的楼层
func (a AnchorNav) Active(value string) AnchorNav {
	return a.set("active", value)
}

// 样式名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a AnchorNav) ClassName(value string) AnchorNav {
	return a.set("className", value)
}

// 可选值: vertical | horizontal
func (a AnchorNav) Direction(value string) AnchorNav {
	return a.set("direction", value)
}

// 是否禁用
func (a AnchorNav) Disabled(value bool) AnchorNav {
	return a.set("disabled", value)
}

// 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (a AnchorNav) DisabledOn(value string) AnchorNav {
	return a.set("disabledOn", value)
}

// 编辑器配置，运行时可以忽略
func (a AnchorNav) EditorSetting(value string) AnchorNav {
	return a.set("editorSetting", value)
}

// 是否隐藏
func (a AnchorNav) Hidden(value bool) AnchorNav {
	return a.set("hidden", value)
}

// 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (a AnchorNav) HiddenOn(value string) AnchorNav {
	return a.set("hiddenOn", value)
}

// 组件唯一 id，主要用于日志采集
func (a AnchorNav) Id(value string) AnchorNav {
	return a.set("id", value)
}

// 导航样式名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a AnchorNav) LinkClassName(value string) AnchorNav {
	return a.set("linkClassName", value)
}

// 楼层集合
func (a AnchorNav) Links(value string) AnchorNav {
	return a.set("links", value)
}

// 事件动作配置
func (a AnchorNav) OnEvent(value string) AnchorNav {
	return a.set("onEvent", value)
}

// 楼层样式名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a AnchorNav) SectionClassName(value string) AnchorNav {
	return a.set("sectionClassName", value)
}

// 是否静态展示
func (a AnchorNav) Static(value bool) AnchorNav {
	return a.set("static", value)
}

// 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a AnchorNav) StaticClassName(value string) AnchorNav {
	return a.set("staticClassName", value)
}

// 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a AnchorNav) StaticInputClassName(value string) AnchorNav {
	return a.set("staticInputClassName", value)
}

// 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a AnchorNav) StaticLabelClassName(value string) AnchorNav {
	return a.set("staticLabelClassName", value)
}

// 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (a AnchorNav) StaticOn(value string) AnchorNav {
	return a.set("staticOn", value)
}

// 静态展示空值占位
func (a AnchorNav) StaticPlaceholder(value string) AnchorNav {
	return a.set("staticPlaceholder", value)
}

func (a AnchorNav) StaticSchema(value string) AnchorNav {
	return a.set("staticSchema", value)
}

// 组件样式
func (a AnchorNav) Style(value string) AnchorNav {
	return a.set("style", value)
}

func (a AnchorNav) TestIdBuilder(value string) AnchorNav {
	return a.set("testIdBuilder", value)
}

func (a AnchorNav) Testid(value string) AnchorNav {
	return a.set("testid", value)
}

// 可以组件级别用来关闭移动端样式
func (a AnchorNav) UseMobileUI(value bool) AnchorNav {
	return a.set("useMobileUI", value)
}

// 是否显示
func (a AnchorNav) Visible(value bool) AnchorNav {
	return a.set("visible", value)
}

// 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (a AnchorNav) VisibleOn(value string) AnchorNav {
	return a.set("visibleOn", value)
}
