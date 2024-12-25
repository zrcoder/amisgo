package comp

type MTab Schema

// Tab 创建一个新的 Tab 实例
func Tab() MTab {
	return MTab{}
}

func (t MTab) set(key string, value any) MTab {
	t[key] = value
	return t
}

// Badge 徽标
func (t MTab) Badge(value string) MTab {
	return t.set("badge", value)
}

// Body 内容 (内容)
func (t MTab) Body(value ...any) MTab {
	return t.set("body", value)
}

// ClassName 容器 css 类名
func (t MTab) ClassName(value string) MTab {
	return t.set("className", value)
}

// Closable 是否可关闭，优先级高于 tabs 的 closable
func (t MTab) Closable(value bool) MTab {
	return t.set("closable", value)
}

// Disabled 是否禁用
func (t MTab) Disabled(value bool) MTab {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (t MTab) DisabledOn(value string) MTab {
	return t.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (t MTab) EditorSetting(value string) MTab {
	return t.set("editorSetting", value)
}

// Hash 设置以后将跟url的hash对应
func (t MTab) Hash(value string) MTab {
	return t.set("hash", value)
}

// Hidden 是否隐藏
func (t MTab) Hidden(value bool) MTab {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (t MTab) HiddenOn(value string) MTab {
	return t.set("hiddenOn", value)
}

// Horizontal 如果是水平排版，这个属性可以细化水平排版的左右宽度占比。
func (t MTab) Horizontal(value string) MTab {
	return t.set("horizontal", value)
}

// Icon 按钮图标
func (t MTab) Icon(value string) MTab {
	return t.set("icon", value)
}

// IconPosition 可选值: left | right
func (t MTab) IconPosition(value string) MTab {
	return t.set("iconPosition", value)
}

// Id 组件唯一 id，主要用于日志采集
func (t MTab) ID(value string) MTab {
	return t.set("id", value)
}

// Mode 配置子表单项默认的展示方式。 可选值: normal | inline | horizontal
func (t MTab) Mode(value string) MTab {
	return t.set("mode", value)
}

// MountOnEnter 点开时才加载卡片内容
func (t MTab) MountOnEnter(value bool) MTab {
	return t.set("mountOnEnter", value)
}

// OnEvent 事件动作配置
func (t MTab) OnEvent(value any) MTab {
	return t.set("onEvent", value)
}

// Reload 设置以后内容每次都会重新渲染
func (t MTab) Reload(value bool) MTab {
	return t.set("reload", value)
}

// Static 是否静态展示
func (t MTab) Static(value bool) MTab {
	return t.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (t MTab) StaticClassName(value string) MTab {
	return t.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (t MTab) StaticInputClassName(value string) MTab {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (t MTab) StaticLabelClassName(value string) MTab {
	return t.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (t MTab) StaticOn(value string) MTab {
	return t.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (t MTab) StaticPlaceholder(value string) MTab {
	return t.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (t MTab) StaticSchema(value string) MTab {
	return t.set("staticSchema", value)
}

// Style 组件样式
func (t MTab) Style(value string) MTab {
	return t.set("style", value)
}

// Tab 内容 (内容)
func (t MTab) Tab(value ...any) MTab {
	return t.set("tab", value)
}

// TestIdBuilder
func (t MTab) TestIdBuilder(value string) MTab {
	return t.set("testIdBuilder", value)
}

// Testid
func (t MTab) Testid(value string) MTab {
	return t.set("testid", value)
}

// Title Tab 标题
func (t MTab) Title(value any) MTab {
	return t.set("title", value)
}

// UnmountOnExit 卡片隐藏就销毁卡片节点
func (t MTab) UnmountOnExit(value bool) MTab {
	return t.set("unmountOnExit", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (t MTab) UseMobileUI(value bool) MTab {
	return t.set("useMobileUI", value)
}

// Visible 是否显示
func (t MTab) Visible(value bool) MTab {
	return t.set("visible", value)
}

// VisibleOn 是否显示表达式
func (t MTab) VisibleOn(value string) MTab {
	return t.set("visibleOn", value)
}
