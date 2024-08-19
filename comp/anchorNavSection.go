package comp

// anchorNavSection 锚点区域渲染器 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/anchor-nav
type anchorNavSection schema

// AnchorNavSection 创建一个新的 AnchorNavSection 实例
func AnchorNavSection() anchorNavSection {
	return make(anchorNavSection)
}

func (a anchorNavSection) set(key string, value any) anchorNavSection {
	a[key] = value
	return a
}

// 内容 (内容)
func (a anchorNavSection) Body(value ...any) anchorNavSection {
	return a.set("body", value)
}

// 子节点
func (a anchorNavSection) Children(value string) anchorNavSection {
	return a.set("children", value)
}

// 容器 css 类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a anchorNavSection) ClassName(value string) anchorNavSection {
	return a.set("className", value)
}

// 是否禁用
func (a anchorNavSection) Disabled(value bool) anchorNavSection {
	return a.set("disabled", value)
}

// 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (a anchorNavSection) DisabledOn(value string) anchorNavSection {
	return a.set("disabledOn", value)
}

// 编辑器配置，运行时可以忽略
func (a anchorNavSection) EditorSetting(value string) anchorNavSection {
	return a.set("editorSetting", value)
}

// 是否隐藏
func (a anchorNavSection) Hidden(value bool) anchorNavSection {
	return a.set("hidden", value)
}

// 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (a anchorNavSection) HiddenOn(value string) anchorNavSection {
	return a.set("hiddenOn", value)
}

// 锚点链接
func (a anchorNavSection) Href(value string) anchorNavSection {
	return a.set("href", value)
}

// 组件唯一 id，主要用于日志采集
func (a anchorNavSection) Id(value string) anchorNavSection {
	return a.set("id", value)
}

// 事件动作配置
func (a anchorNavSection) OnEvent(value any) anchorNavSection {
	return a.set("onEvent", value)
}

// 是否静态展示
func (a anchorNavSection) Static(value bool) anchorNavSection {
	return a.set("static", value)
}

// 静态展示表单项类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a anchorNavSection) StaticClassName(value string) anchorNavSection {
	return a.set("staticClassName", value)
}

// 静态展示表单项Value类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a anchorNavSection) StaticInputClassName(value string) anchorNavSection {
	return a.set("staticInputClassName", value)
}

// 静态展示表单项Label类名 (css类名，配置字符串，或者对象。 className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如： className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (a anchorNavSection) StaticLabelClassName(value string) anchorNavSection {
	return a.set("staticLabelClassName", value)
}

// 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (a anchorNavSection) StaticOn(value string) anchorNavSection {
	return a.set("staticOn", value)
}

// 静态展示空值占位
func (a anchorNavSection) StaticPlaceholder(value string) anchorNavSection {
	return a.set("staticPlaceholder", value)
}

func (a anchorNavSection) StaticSchema(value string) anchorNavSection {
	return a.set("staticSchema", value)
}

// 组件样式
func (a anchorNavSection) Style(value any) anchorNavSection {
	return a.set("style", value)
}

func (a anchorNavSection) TestIdBuilder(value string) anchorNavSection {
	return a.set("testIdBuilder", value)
}

func (a anchorNavSection) Testid(value string) anchorNavSection {
	return a.set("testid", value)
}

// 导航文字说明
func (a anchorNavSection) Title(value any) anchorNavSection {
	return a.set("title", value)
}

// 可以组件级别用来关闭移动端样式
func (a anchorNavSection) UseMobileUI(value bool) anchorNavSection {
	return a.set("useMobileUI", value)
}

// 是否显示
func (a anchorNavSection) Visible(value bool) anchorNavSection {
	return a.set("visible", value)
}

// 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (a anchorNavSection) VisibleOn(value string) anchorNavSection {
	return a.set("visibleOn", value)
}
