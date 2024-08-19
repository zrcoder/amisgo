package model

type TabSchema Schema

// Tab 创建一个新的 Tab 实例
func Tab() TabSchema {
	return TabSchema{}
}

func (t TabSchema) set(key string, value any) TabSchema {
	t[key] = value
	return t
}

// Badge 徽标
func (t TabSchema) Badge(value string) TabSchema {
	return t.set("badge", value)
}

// Body 内容 (内容)
func (t TabSchema) Body(value ...any) TabSchema {
	return t.set("body", value)
}

// ClassName 容器 css 类名
func (t TabSchema) ClassName(value string) TabSchema {
	return t.set("className", value)
}

// Closable 是否可关闭，优先级高于 tabs 的 closable
func (t TabSchema) Closable(value bool) TabSchema {
	return t.set("closable", value)
}

// Disabled 是否禁用
func (t TabSchema) Disabled(value bool) TabSchema {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (t TabSchema) DisabledOn(value string) TabSchema {
	return t.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t TabSchema) EditorSetting(value string) TabSchema {
	return t.set("editorSetting", value)
}

// Hash 设置以后将跟url的hash对应
func (t TabSchema) Hash(value string) TabSchema {
	return t.set("hash", value)
}

// Hidden 是否隐藏
func (t TabSchema) Hidden(value bool) TabSchema {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (t TabSchema) HiddenOn(value string) TabSchema {
	return t.set("hiddenOn", value)
}

// Horizontal 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
func (t TabSchema) Horizontal(value string) TabSchema {
	return t.set("horizontal", value)
}

// Icon 按钮图标
func (t TabSchema) Icon(value string) TabSchema {
	return t.set("icon", value)
}

// IconPosition 可选值: left | right
func (t TabSchema) IconPosition(value string) TabSchema {
	return t.set("iconPosition", value)
}

// Id 组件唯一 id，主要用于日志采集
func (t TabSchema) Id(value string) TabSchema {
	return t.set("id", value)
}

// Mode 配置子表单项默认的展示方式。 可选值: normal | inline | horizontal
func (t TabSchema) Mode(value string) TabSchema {
	return t.set("mode", value)
}

// MountOnEnter 点开时才加载卡片内容
func (t TabSchema) MountOnEnter(value bool) TabSchema {
	return t.set("mountOnEnter", value)
}

// OnEvent 事件动作配置
func (t TabSchema) OnEvent(value any) TabSchema {
	return t.set("onEvent", value)
}

// Reload 设置以后内容每次都会重新渲染
func (t TabSchema) Reload(value bool) TabSchema {
	return t.set("reload", value)
}

// Static 是否静态展示
func (t TabSchema) Static(value bool) TabSchema {
	return t.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (t TabSchema) StaticClassName(value string) TabSchema {
	return t.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (t TabSchema) StaticInputClassName(value string) TabSchema {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (t TabSchema) StaticLabelClassName(value string) TabSchema {
	return t.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (t TabSchema) StaticOn(value string) TabSchema {
	return t.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (t TabSchema) StaticPlaceholder(value string) TabSchema {
	return t.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (t TabSchema) StaticSchema(value string) TabSchema {
	return t.set("staticSchema", value)
}

// Style 组件样式
func (t TabSchema) Style(value string) TabSchema {
	return t.set("style", value)
}

// Tab 内容 (内容)
func (t TabSchema) Tab(value ...any) TabSchema {
	return t.set("tab", value)
}

// TestIdBuilder
func (t TabSchema) TestIdBuilder(value string) TabSchema {
	return t.set("testIdBuilder", value)
}

// Testid
func (t TabSchema) Testid(value string) TabSchema {
	return t.set("testid", value)
}

// Title Tab 标题
func (t TabSchema) Title(value any) TabSchema {
	return t.set("title", value)
}

// UnmountOnExit 卡片隐藏就销毁卡片节点
func (t TabSchema) UnmountOnExit(value bool) TabSchema {
	return t.set("unmountOnExit", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (t TabSchema) UseMobileUI(value bool) TabSchema {
	return t.set("useMobileUI", value)
}

// Visible 是否显示
func (t TabSchema) Visible(value bool) TabSchema {
	return t.set("visible", value)
}

// VisibleOn 是否显示表达式
func (t TabSchema) VisibleOn(value string) TabSchema {
	return t.set("visibleOn", value)
}
