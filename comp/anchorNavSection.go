package comp

// AnchorNavSection 锚点区域渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/anchor-nav
type AnchorNavSection Schema

// NewAnchorNavSection 创建一个新的 AnchorNavSection 实例
func NewAnchorNavSection() AnchorNavSection {
	return make(AnchorNavSection)
}

func (a AnchorNavSection) set(key string, value interface{}) AnchorNavSection {
	a[key] = value
	return a
}

// 内容 (内容)
func (a AnchorNavSection) Body(value ...interface{}) AnchorNavSection {
	return a.set("body", value)
}

// 子节点
func (a AnchorNavSection) Children(value string) AnchorNavSection {
	return a.set("children", value)
}

// 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a AnchorNavSection) ClassName(value string) AnchorNavSection {
	return a.set("className", value)
}

// 是否禁用
func (a AnchorNavSection) Disabled(value bool) AnchorNavSection {
	return a.set("disabled", value)
}

// 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (a AnchorNavSection) DisabledOn(value string) AnchorNavSection {
	return a.set("disabledOn", value)
}

// 编辑器配置，运行时可以忽略
func (a AnchorNavSection) EditorSetting(value string) AnchorNavSection {
	return a.set("editorSetting", value)
}

// 是否隐藏
func (a AnchorNavSection) Hidden(value bool) AnchorNavSection {
	return a.set("hidden", value)
}

// 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (a AnchorNavSection) HiddenOn(value string) AnchorNavSection {
	return a.set("hiddenOn", value)
}

// 锚点链接
func (a AnchorNavSection) Href(value string) AnchorNavSection {
	return a.set("href", value)
}

// 组件唯一 id，主要用于日志采集
func (a AnchorNavSection) Id(value string) AnchorNavSection {
	return a.set("id", value)
}

// 事件动作配置
func (a AnchorNavSection) OnEvent(value string) AnchorNavSection {
	return a.set("onEvent", value)
}

// 是否静态展示
func (a AnchorNavSection) Static(value bool) AnchorNavSection {
	return a.set("static", value)
}

// 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a AnchorNavSection) StaticClassName(value string) AnchorNavSection {
	return a.set("staticClassName", value)
}

// 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a AnchorNavSection) StaticInputClassName(value string) AnchorNavSection {
	return a.set("staticInputClassName", value)
}

// 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a AnchorNavSection) StaticLabelClassName(value string) AnchorNavSection {
	return a.set("staticLabelClassName", value)
}

// 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (a AnchorNavSection) StaticOn(value string) AnchorNavSection {
	return a.set("staticOn", value)
}

// 静态展示空值占位
func (a AnchorNavSection) StaticPlaceholder(value string) AnchorNavSection {
	return a.set("staticPlaceholder", value)
}

func (a AnchorNavSection) StaticSchema(value string) AnchorNavSection {
	return a.set("staticSchema", value)
}

// 组件样式
func (a AnchorNavSection) Style(value string) AnchorNavSection {
	return a.set("style", value)
}

func (a AnchorNavSection) TestIdBuilder(value string) AnchorNavSection {
	return a.set("testIdBuilder", value)
}

func (a AnchorNavSection) Testid(value string) AnchorNavSection {
	return a.set("testid", value)
}

// 导航文字说明
func (a AnchorNavSection) Title(value string) AnchorNavSection {
	return a.set("title", value)
}

// 可以组件级别用来关闭移动端样式
func (a AnchorNavSection) UseMobileUI(value bool) AnchorNavSection {
	return a.set("useMobileUI", value)
}

// 是否显示
func (a AnchorNavSection) Visible(value bool) AnchorNavSection {
	return a.set("visible", value)
}

// 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (a AnchorNavSection) VisibleOn(value string) AnchorNavSection {
	return a.set("visibleOn", value)
}
