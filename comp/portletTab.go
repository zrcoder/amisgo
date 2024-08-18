package comp

// PortletTab
//
// @version 6.7.0
type PortletTab Schema

// NewPortletTab 创建一个新的 PortletTab 实例
func NewPortletTab() PortletTab {
	return PortletTab{}
}

func (p PortletTab) set(key string, value interface{}) PortletTab {
	p[key] = value
	return p
}

// Body 内容
func (p PortletTab) Body(value ...interface{}) PortletTab {
	return p.set("body", value)
}

// ClassName 容器 css 类名
func (p PortletTab) ClassName(value string) PortletTab {
	return p.set("className", value)
}

// Disabled 是否禁用
func (p PortletTab) Disabled(value bool) PortletTab {
	return p.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (p PortletTab) DisabledOn(value string) PortletTab {
	return p.set("disabledOn", value)
}

// EditorSetting 编辑器配置
func (p PortletTab) EditorSetting(value string) PortletTab {
	return p.set("editorSetting", value)
}

// Hidden 是否隐藏
func (p PortletTab) Hidden(value bool) PortletTab {
	return p.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (p PortletTab) HiddenOn(value string) PortletTab {
	return p.set("hiddenOn", value)
}

// Icon 按钮图标
func (p PortletTab) Icon(value string) PortletTab {
	return p.set("icon", value)
}

// IconPosition 图标位置
func (p PortletTab) IconPosition(value string) PortletTab {
	return p.set("iconPosition", value)
}

// ID 组件唯一 id
func (p PortletTab) ID(value string) PortletTab {
	return p.set("id", value)
}

// MountOnEnter 点开时才加载卡片内容
func (p PortletTab) MountOnEnter(value bool) PortletTab {
	return p.set("mountOnEnter", value)
}

// OnEvent 事件动作配置
func (p PortletTab) OnEvent(value string) PortletTab {
	return p.set("onEvent", value)
}

// Reload 设置以后内容每次都会重新渲染
func (p PortletTab) Reload(value bool) PortletTab {
	return p.set("reload", value)
}

// Static 是否静态展示
func (p PortletTab) Static(value bool) PortletTab {
	return p.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (p PortletTab) StaticClassName(value string) PortletTab {
	return p.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (p PortletTab) StaticInputClassName(value string) PortletTab {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (p PortletTab) StaticLabelClassName(value string) PortletTab {
	return p.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (p PortletTab) StaticOn(value string) PortletTab {
	return p.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (p PortletTab) StaticPlaceholder(value string) PortletTab {
	return p.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (p PortletTab) StaticSchema(value string) PortletTab {
	return p.set("staticSchema", value)
}

// Style 组件样式
func (p PortletTab) Style(value string) PortletTab {
	return p.set("style", value)
}

// Tab 内容
func (p PortletTab) Tab(value string) PortletTab {
	return p.set("tab", value)
}

// TestIdBuilder 测试 id 构造器
func (p PortletTab) TestIdBuilder(value string) PortletTab {
	return p.set("testIdBuilder", value)
}

// Testid 测试 id
func (p PortletTab) Testid(value string) PortletTab {
	return p.set("testid", value)
}

// Title Tab 标题
func (p PortletTab) Title(value string) PortletTab {
	return p.set("title", value)
}

// Toolbar 可以在右侧配置点其他功能按钮，随着 tab 切换而切换
func (p PortletTab) Toolbar(value string) PortletTab {
	return p.set("toolbar", value)
}

// UnmountOnExit 卡片隐藏就销毁卡片节点
func (p PortletTab) UnmountOnExit(value bool) PortletTab {
	return p.set("unmountOnExit", value)
}

// UseMobileUI 组件级别用来关闭移动端样式
func (p PortletTab) UseMobileUI(value bool) PortletTab {
	return p.set("useMobileUI", value)
}

// Visible 是否显示
func (p PortletTab) Visible(value bool) PortletTab {
	return p.set("visible", value)
}

// VisibleOn 是否显示表达式
func (p PortletTab) VisibleOn(value string) PortletTab {
	return p.set("visibleOn", value)
}
