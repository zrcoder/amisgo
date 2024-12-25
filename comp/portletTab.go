package comp

// portletTab

type portletTab Schema

// PortletTab 创建一个新的 PortletTab 实例
func PortletTab() portletTab {
	return portletTab{}
}

func (p portletTab) set(key string, value any) portletTab {
	p[key] = value
	return p
}

// Body 内容
func (p portletTab) Body(value ...any) portletTab {
	return p.set("body", value)
}

// ClassName 容器 css 类名
func (p portletTab) ClassName(value string) portletTab {
	return p.set("className", value)
}

// Disabled 是否禁用
func (p portletTab) Disabled(value bool) portletTab {
	return p.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (p portletTab) DisabledOn(value string) portletTab {
	return p.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (p portletTab) EditorSetting(value string) portletTab {
	return p.set("editorSetting", value)
}

// Hidden 是否隐藏
func (p portletTab) Hidden(value bool) portletTab {
	return p.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (p portletTab) HiddenOn(value string) portletTab {
	return p.set("hiddenOn", value)
}

// Icon 按钮图标
func (p portletTab) Icon(value string) portletTab {
	return p.set("icon", value)
}

// IconPosition 图标位置
func (p portletTab) IconPosition(value string) portletTab {
	return p.set("iconPosition", value)
}

// ID 组件唯一 id
func (p portletTab) ID(value string) portletTab {
	return p.set("id", value)
}

// MountOnEnter 点开时才加载卡片内容
func (p portletTab) MountOnEnter(value bool) portletTab {
	return p.set("mountOnEnter", value)
}

// OnEvent 事件动作配置
func (p portletTab) OnEvent(value any) portletTab {
	return p.set("onEvent", value)
}

// Reload 设置以后内容每次都会重新渲染
func (p portletTab) Reload(value bool) portletTab {
	return p.set("reload", value)
}

// Static 是否静态展示
func (p portletTab) Static(value bool) portletTab {
	return p.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (p portletTab) StaticClassName(value string) portletTab {
	return p.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (p portletTab) StaticInputClassName(value string) portletTab {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (p portletTab) StaticLabelClassName(value string) portletTab {
	return p.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (p portletTab) StaticOn(value string) portletTab {
	return p.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (p portletTab) StaticPlaceholder(value string) portletTab {
	return p.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (p portletTab) StaticSchema(value string) portletTab {
	return p.set("staticSchema", value)
}

// Style 组件样式
func (p portletTab) Style(value any) portletTab {
	return p.set("style", value)
}

// Tab 内容
func (p portletTab) Tab(value string) portletTab {
	return p.set("tab", value)
}

// TestIdBuilder 测试 id 构造器
func (p portletTab) TestIdBuilder(value string) portletTab {
	return p.set("testIdBuilder", value)
}

// Testid 测试 id
func (p portletTab) Testid(value string) portletTab {
	return p.set("testid", value)
}

// Title Tab 标题
func (p portletTab) Title(value any) portletTab {
	return p.set("title", value)
}

// Toolbar 可以在右侧配置点其他功能按钮，随着 tab 切换而切换
func (p portletTab) Toolbar(value string) portletTab {
	return p.set("toolbar", value)
}

// UnmountOnExit 卡片隐藏就销毁卡片节点
func (p portletTab) UnmountOnExit(value bool) portletTab {
	return p.set("unmountOnExit", value)
}

// UseMobileUI 组件级别用来关闭移动端样式
func (p portletTab) UseMobileUI(value bool) portletTab {
	return p.set("useMobileUI", value)
}

// Visible 是否显示
func (p portletTab) Visible(value bool) portletTab {
	return p.set("visible", value)
}

// VisibleOn 是否显示表达式
func (p portletTab) VisibleOn(value string) portletTab {
	return p.set("visibleOn", value)
}
