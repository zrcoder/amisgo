package comp

// portlet
//
// @version 6.7.0
type portlet schema

// Portlet 创建一个新的 Portlet 实例
func Portlet() portlet {
	return portlet{}.set("type", "portlet")
}

func (p portlet) set(key string, value interface{}) portlet {
	p[key] = value
	return p
}

// ClassName 容器 css 类名
func (p portlet) ClassName(value string) portlet {
	return p.set("className", value)
}

// ContentClassName 内容类名
func (p portlet) ContentClassName(value string) portlet {
	return p.set("contentClassName", value)
}

// Description 标题右侧的描述
func (p portlet) Description(value string) portlet {
	return p.set("description", value)
}

// Disabled 是否禁用
func (p portlet) Disabled(value bool) portlet {
	return p.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (p portlet) DisabledOn(value string) portlet {
	return p.set("disabledOn", value)
}

// Divider header和内容是否展示分割线
func (p portlet) Divider(value bool) portlet {
	return p.set("divider", value)
}

// EditorSetting 编辑器配置
func (p portlet) EditorSetting(value string) portlet {
	return p.set("editorSetting", value)
}

// Hidden 是否隐藏
func (p portlet) Hidden(value bool) portlet {
	return p.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (p portlet) HiddenOn(value string) portlet {
	return p.set("hiddenOn", value)
}

// HideHeader 隐藏头部
func (p portlet) HideHeader(value bool) portlet {
	return p.set("hideHeader", value)
}

// ID 组件唯一 id
func (p portlet) ID(value string) portlet {
	return p.set("id", value)
}

// LinksClassName 链接外层类名
func (p portlet) LinksClassName(value string) portlet {
	return p.set("linksClassName", value)
}

// MountOnEnter 卡片是否只有在点开的时候加载
func (p portlet) MountOnEnter(value bool) portlet {
	return p.set("mountOnEnter", value)
}

// OnEvent 事件动作配置
func (p portlet) OnEvent(value string) portlet {
	return p.set("onEvent", value)
}

// Scrollable 是否支持溢出滚动
func (p portlet) Scrollable(value bool) portlet {
	return p.set("scrollable", value)
}

// Source 关联已有数据
func (p portlet) Source(value string) portlet {
	return p.set("source", value)
}

// Static 是否静态展示
func (p portlet) Static(value bool) portlet {
	return p.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (p portlet) StaticClassName(value string) portlet {
	return p.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (p portlet) StaticInputClassName(value string) portlet {
	return p.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (p portlet) StaticLabelClassName(value string) portlet {
	return p.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (p portlet) StaticOn(value string) portlet {
	return p.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (p portlet) StaticPlaceholder(value string) portlet {
	return p.set("staticPlaceholder", value)
}

// StaticSchema 静态展示 schema
func (p portlet) StaticSchema(value string) portlet {
	return p.set("staticSchema", value)
}

// Style 自定义样式
func (p portlet) Style(value string) portlet {
	return p.set("style", value)
}

// Tabs 标签页
func (p portlet) Tabs(value string) portlet {
	return p.set("tabs", value)
}

// TabsClassName 标签页类名
func (p portlet) TabsClassName(value string) portlet {
	return p.set("tabsClassName", value)
}

// TabsMode 展示形式
func (p portlet) TabsMode(value string) portlet {
	return p.set("tabsMode", value)
}

// TestIdBuilder 测试 id 构造器
func (p portlet) TestIdBuilder(value string) portlet {
	return p.set("testIdBuilder", value)
}

// Testid 测试 id
func (p portlet) Testid(value string) portlet {
	return p.set("testid", value)
}

// Toolbar 可以在右侧配置点其他功能按钮
func (p portlet) Toolbar(value string) portlet {
	return p.set("toolbar", value)
}

// UnmountOnExit 卡片隐藏的时候是否销毁卡片内容
func (p portlet) UnmountOnExit(value bool) portlet {
	return p.set("unmountOnExit", value)
}

// UseMobileUI 组件级别关闭移动端样式
func (p portlet) UseMobileUI(value bool) portlet {
	return p.set("useMobileUI", value)
}

// Visible 是否显示
func (p portlet) Visible(value bool) portlet {
	return p.set("visible", value)
}

// VisibleOn 是否显示表达式
func (p portlet) VisibleOn(value string) portlet {
	return p.set("visibleOn", value)
}
