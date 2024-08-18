package comp

// Portlet
//
// @version 6.7.0
type Portlet Schema

// NewPortlet 创建一个新的 Portlet 实例
func NewPortlet() Portlet {
	return Portlet{}.set("type", "portlet")
}

func (p Portlet) set(key string, value interface{}) Portlet {
	p[key] = value
	return p
}

// ClassName 容器 css 类名
func (p Portlet) ClassName(value string) Portlet {
	return p.set("className", value)
}

// ContentClassName 内容类名
func (p Portlet) ContentClassName(value string) Portlet {
	return p.set("contentClassName", value)
}

// Description 标题右侧的描述
func (p Portlet) Description(value string) Portlet {
	return p.set("description", value)
}

// Disabled 是否禁用
func (p Portlet) Disabled(value bool) Portlet {
	return p.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (p Portlet) DisabledOn(value string) Portlet {
	return p.set("disabledOn", value)
}

// Divider header和内容是否展示分割线
func (p Portlet) Divider(value bool) Portlet {
	return p.set("divider", value)
}

// EditorSetting 编辑器配置
func (p Portlet) EditorSetting(value string) Portlet {
	return p.set("editorSetting", value)
}

// Hidden 是否隐藏
func (p Portlet) Hidden(value bool) Portlet {
	return p.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (p Portlet) HiddenOn(value string) Portlet {
	return p.set("hiddenOn", value)
}

// HideHeader 隐藏头部
func (p Portlet) HideHeader(value bool) Portlet {
	return p.set("hideHeader", value)
}

// ID 组件唯一 id
func (p Portlet) ID(value string) Portlet {
	return p.set("id", value)
}

// LinksClassName 链接外层类名
func (p Portlet) LinksClassName(value string) Portlet {
	return p.set("linksClassName", value)
}

// MountOnEnter 卡片是否只有在点开的时候加载
func (p Portlet) MountOnEnter(value bool) Portlet {
	return p.set("mountOnEnter", value)
}

// OnEvent 事件动作配置
func (p Portlet) OnEvent(value string) Portlet {
	return p.set("onEvent", value)
}

// Scrollable 是否支持溢出滚动
func (p Portlet) Scrollable(value bool) Portlet {
	return p.set("scrollable", value)
}

// Source 关联已有数据
func (p Portlet) Source(value string) Portlet {
	return p.set("source", value)
}

// Static 是否静态展示
func (p Portlet) Static(value bool) Portlet {
	return p.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (p Portlet) StaticClassName(value string) Portlet {
	return p.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (p Portlet) StaticInputClassName(value string) Portlet {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (p Portlet) StaticLabelClassName(value string) Portlet {
	return p.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (p Portlet) StaticOn(value string) Portlet {
	return p.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (p Portlet) StaticPlaceholder(value string) Portlet {
	return p.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (p Portlet) StaticSchema(value string) Portlet {
	return p.set("staticSchema", value)
}

// Style 自定义样式
func (p Portlet) Style(value string) Portlet {
	return p.set("style", value)
}

// Tabs 标签页
func (p Portlet) Tabs(value string) Portlet {
	return p.set("tabs", value)
}

// TabsClassName 标签页类名
func (p Portlet) TabsClassName(value string) Portlet {
	return p.set("tabsClassName", value)
}

// TabsMode 展示形式
func (p Portlet) TabsMode(value string) Portlet {
	return p.set("tabsMode", value)
}

// TestIdBuilder 测试 id 构造器
func (p Portlet) TestIdBuilder(value string) Portlet {
	return p.set("testIdBuilder", value)
}

// Testid 测试 id
func (p Portlet) Testid(value string) Portlet {
	return p.set("testid", value)
}

// Toolbar 可以在右侧配置点其他功能按钮
func (p Portlet) Toolbar(value string) Portlet {
	return p.set("toolbar", value)
}

// UnmountOnExit 卡片隐藏的时候是否销毁卡片内容
func (p Portlet) UnmountOnExit(value bool) Portlet {
	return p.set("unmountOnExit", value)
}

// UseMobileUI 组件级别关闭移动端样式
func (p Portlet) UseMobileUI(value bool) Portlet {
	return p.set("useMobileUI", value)
}

// Visible 是否显示
func (p Portlet) Visible(value bool) Portlet {
	return p.set("visible", value)
}

// VisibleOn 是否显示表达式
func (p Portlet) VisibleOn(value string) Portlet {
	return p.set("visibleOn", value)
}
