package comp

// tab 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tab
//
// @version 6.7.0
type tab schema

// Tab 创建一个新的 Tab 实例
func Tab() tab {
	return tab{}
}

func (t tab) set(key string, value interface{}) tab {
	t[key] = value
	return t
}

// Badge 徽标
func (t tab) Badge(value string) tab {
	return t.set("badge", value)
}

// Body 内容 (内容)
func (t tab) Body(value ...interface{}) tab {
	return t.set("body", value)
}

// ClassName 容器 css 类名
func (t tab) ClassName(value string) tab {
	return t.set("className", value)
}

// Closable 是否可关闭，优先级高于 tabs 的 closable
func (t tab) Closable(value bool) tab {
	return t.set("closable", value)
}

// Disabled 是否禁用
func (t tab) Disabled(value bool) tab {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (t tab) DisabledOn(value string) tab {
	return t.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t tab) EditorSetting(value string) tab {
	return t.set("editorSetting", value)
}

// Hash 设置以后将跟url的hash对应
func (t tab) Hash(value string) tab {
	return t.set("hash", value)
}

// Hidden 是否隐藏
func (t tab) Hidden(value bool) tab {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (t tab) HiddenOn(value string) tab {
	return t.set("hiddenOn", value)
}

// Horizontal 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
func (t tab) Horizontal(value string) tab {
	return t.set("horizontal", value)
}

// Icon 按钮图标
func (t tab) Icon(value string) tab {
	return t.set("icon", value)
}

// IconPosition 可选值: left | right
func (t tab) IconPosition(value string) tab {
	return t.set("iconPosition", value)
}

// Id 组件唯一 id，主要用于日志采集
func (t tab) Id(value string) tab {
	return t.set("id", value)
}

// Mode 配置子表单项默认的展示方式。 可选值: normal | inline | horizontal
func (t tab) Mode(value string) tab {
	return t.set("mode", value)
}

// MountOnEnter 点开时才加载卡片内容
func (t tab) MountOnEnter(value bool) tab {
	return t.set("mountOnEnter", value)
}

// OnEvent 事件动作配置
func (t tab) OnEvent(value string) tab {
	return t.set("onEvent", value)
}

// Reload 设置以后内容每次都会重新渲染
func (t tab) Reload(value bool) tab {
	return t.set("reload", value)
}

// Static 是否静态展示
func (t tab) Static(value bool) tab {
	return t.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (t tab) StaticClassName(value string) tab {
	return t.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (t tab) StaticInputClassName(value string) tab {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (t tab) StaticLabelClassName(value string) tab {
	return t.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (t tab) StaticOn(value string) tab {
	return t.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (t tab) StaticPlaceholder(value string) tab {
	return t.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (t tab) StaticSchema(value string) tab {
	return t.set("staticSchema", value)
}

// Style 组件样式
func (t tab) Style(value string) tab {
	return t.set("style", value)
}

// Tab 内容 (内容)
func (t tab) Tab(value ...interface{}) tab {
	return t.set("tab", value)
}

// TestIdBuilder
func (t tab) TestIdBuilder(value string) tab {
	return t.set("testIdBuilder", value)
}

// Testid
func (t tab) Testid(value string) tab {
	return t.set("testid", value)
}

// Title Tab 标题
func (t tab) Title(value string) tab {
	return t.set("title", value)
}

// UnmountOnExit 卡片隐藏就销毁卡片节点
func (t tab) UnmountOnExit(value bool) tab {
	return t.set("unmountOnExit", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (t tab) UseMobileUI(value bool) tab {
	return t.set("useMobileUI", value)
}

// Visible 是否显示
func (t tab) Visible(value bool) tab {
	return t.set("visible", value)
}

// VisibleOn 是否显示表达式
func (t tab) VisibleOn(value string) tab {
	return t.set("visibleOn", value)
}
