package comp

// Tab 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/tab
//
// @version 6.7.0
type Tab Schema

// NewTab 创建一个新的 Tab 实例
func NewTab() Tab {
	return Tab{}.set("type", "tab")
}

func (t Tab) set(key string, value interface{}) Tab {
	t[key] = value
	return t
}

// Badge 徽标
func (t Tab) Badge(value string) Tab {
	return t.set("badge", value)
}

// Body 内容 (内容)
func (t Tab) Body(value ...interface{}) Tab {
	return t.set("body", value)
}

// ClassName 容器 css 类名
func (t Tab) ClassName(value string) Tab {
	return t.set("className", value)
}

// Closable 是否可关闭，优先级高于 tabs 的 closable
func (t Tab) Closable(value bool) Tab {
	return t.set("closable", value)
}

// Disabled 是否禁用
func (t Tab) Disabled(value bool) Tab {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (t Tab) DisabledOn(value string) Tab {
	return t.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t Tab) EditorSetting(value string) Tab {
	return t.set("editorSetting", value)
}

// Hash 设置以后将跟url的hash对应
func (t Tab) Hash(value string) Tab {
	return t.set("hash", value)
}

// Hidden 是否隐藏
func (t Tab) Hidden(value bool) Tab {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (t Tab) HiddenOn(value string) Tab {
	return t.set("hiddenOn", value)
}

// Horizontal 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
func (t Tab) Horizontal(value string) Tab {
	return t.set("horizontal", value)
}

// Icon 按钮图标
func (t Tab) Icon(value string) Tab {
	return t.set("icon", value)
}

// IconPosition 可选值: left | right
func (t Tab) IconPosition(value string) Tab {
	return t.set("iconPosition", value)
}

// Id 组件唯一 id，主要用于日志采集
func (t Tab) Id(value string) Tab {
	return t.set("id", value)
}

// Mode 配置子表单项默认的展示方式。 可选值: normal | inline | horizontal
func (t Tab) Mode(value string) Tab {
	return t.set("mode", value)
}

// MountOnEnter 点开时才加载卡片内容
func (t Tab) MountOnEnter(value bool) Tab {
	return t.set("mountOnEnter", value)
}

// OnEvent 事件动作配置
func (t Tab) OnEvent(value string) Tab {
	return t.set("onEvent", value)
}

// Reload 设置以后内容每次都会重新渲染
func (t Tab) Reload(value bool) Tab {
	return t.set("reload", value)
}

// Static 是否静态展示
func (t Tab) Static(value bool) Tab {
	return t.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (t Tab) StaticClassName(value string) Tab {
	return t.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (t Tab) StaticInputClassName(value string) Tab {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (t Tab) StaticLabelClassName(value string) Tab {
	return t.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (t Tab) StaticOn(value string) Tab {
	return t.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (t Tab) StaticPlaceholder(value string) Tab {
	return t.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (t Tab) StaticSchema(value string) Tab {
	return t.set("staticSchema", value)
}

// Style 组件样式
func (t Tab) Style(value string) Tab {
	return t.set("style", value)
}

// Tab 内容 (内容)
func (t Tab) Tab(value string) Tab {
	return t.set("tab", value)
}

// TestIdBuilder
func (t Tab) TestIdBuilder(value string) Tab {
	return t.set("testIdBuilder", value)
}

// Testid
func (t Tab) Testid(value string) Tab {
	return t.set("testid", value)
}

// Title Tab 标题
func (t Tab) Title(value string) Tab {
	return t.set("title", value)
}

// UnmountOnExit 卡片隐藏就销毁卡片节点
func (t Tab) UnmountOnExit(value bool) Tab {
	return t.set("unmountOnExit", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (t Tab) UseMobileUI(value bool) Tab {
	return t.set("useMobileUI", value)
}

// Visible 是否显示
func (t Tab) Visible(value bool) Tab {
	return t.set("visible", value)
}

// VisibleOn 是否显示表达式
func (t Tab) VisibleOn(value string) Tab {
	return t.set("visibleOn", value)
}
