package comp

// anchorNav 锚点导航渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/anchor-nav
type anchorNav schema

// AnchorNav 创建一个新的 AnchorNav 实例
func AnchorNav() anchorNav {
	return make(anchorNav).set("type", "anchor-nav")
}

func (a anchorNav) set(key string, value any) anchorNav {
	a[key] = value
	return a
}

// 被激活（定位）的楼层
func (a anchorNav) Active(value string) anchorNav {
	return a.set("active", value)
}

// 样式名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a anchorNav) ClassName(value string) anchorNav {
	return a.set("className", value)
}

// 可选值: vertical | horizontal
func (a anchorNav) Direction(value string) anchorNav {
	return a.set("direction", value)
}

// 是否禁用
func (a anchorNav) Disabled(value bool) anchorNav {
	return a.set("disabled", value)
}

// 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (a anchorNav) DisabledOn(value string) anchorNav {
	return a.set("disabledOn", value)
}

// 编辑器配置，运行时可以忽略
func (a anchorNav) EditorSetting(value string) anchorNav {
	return a.set("editorSetting", value)
}

// 是否隐藏
func (a anchorNav) Hidden(value bool) anchorNav {
	return a.set("hidden", value)
}

// 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (a anchorNav) HiddenOn(value string) anchorNav {
	return a.set("hiddenOn", value)
}

// 组件唯一 id，主要用于日志采集
func (a anchorNav) ID(value string) anchorNav {
	return a.set("id", value)
}

// 导航样式名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a anchorNav) LinkClassName(value string) anchorNav {
	return a.set("linkClassName", value)
}

// 楼层集合
func (a anchorNav) Links(value string) anchorNav {
	return a.set("links", value)
}

// 事件动作配置
func (a anchorNav) OnEvent(value any) anchorNav {
	return a.set("onEvent", value)
}

// 楼层样式名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a anchorNav) SectionClassName(value string) anchorNav {
	return a.set("sectionClassName", value)
}

// 是否静态展示
func (a anchorNav) Static(value bool) anchorNav {
	return a.set("static", value)
}

// 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a anchorNav) StaticClassName(value string) anchorNav {
	return a.set("staticClassName", value)
}

// 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a anchorNav) StaticInputClassName(value string) anchorNav {
	return a.set("staticInputClassName", value)
}

// 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a anchorNav) StaticLabelClassName(value string) anchorNav {
	return a.set("staticLabelClassName", value)
}

// 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (a anchorNav) StaticOn(value string) anchorNav {
	return a.set("staticOn", value)
}

// 静态展示空值占位
func (a anchorNav) StaticPlaceholder(value string) anchorNav {
	return a.set("staticPlaceholder", value)
}

func (a anchorNav) StaticSchema(value string) anchorNav {
	return a.set("staticSchema", value)
}

// 组件样式
func (a anchorNav) Style(value any) anchorNav {
	return a.set("style", value)
}

func (a anchorNav) TestIdBuilder(value string) anchorNav {
	return a.set("testIdBuilder", value)
}

func (a anchorNav) Testid(value string) anchorNav {
	return a.set("testid", value)
}

// 可以组件级别用来关闭移动端样式
func (a anchorNav) UseMobileUI(value bool) anchorNav {
	return a.set("useMobileUI", value)
}

// 是否显示
func (a anchorNav) Visible(value bool) anchorNav {
	return a.set("visible", value)
}

// 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (a anchorNav) VisibleOn(value string) anchorNav {
	return a.set("visibleOn", value)
}
