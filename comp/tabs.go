package comp

// Tabs 选项卡控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/tabs
//
// @version 6.7.0
type Tabs Schema

// NewTabs 创建一个新的 Tabs 实例
func NewTabs() Tabs {
	return Tabs{}.set("type", "tabs")
}

func (t Tabs) set(key string, value interface{}) Tabs {
	t[key] = value
	return t
}

// ActiveKey 激活的选项卡，hash值或索引值，支持使用表达式
func (t Tabs) ActiveKey(value string) Tabs {
	return t.set("activeKey", value)
}

// AddBtnText 自定义增加按钮文案
func (t Tabs) AddBtnText(value string) Tabs {
	return t.set("addBtnText", value)
}

// Addable 是否支持新增
func (t Tabs) Addable(value bool) Tabs {
	return t.set("addable", value)
}

// ClassName 容器 css 类名
func (t Tabs) ClassName(value string) Tabs {
	return t.set("className", value)
}

// Closable 是否支持删除
func (t Tabs) Closable(value bool) Tabs {
	return t.set("closable", value)
}

// CollapseBtnLabel 折叠按钮文字
func (t Tabs) CollapseBtnLabel(value string) Tabs {
	return t.set("collapseBtnLabel", value)
}

// CollapseOnExceed 超过多少个时折叠按钮
func (t Tabs) CollapseOnExceed(value string) Tabs {
	return t.set("collapseOnExceed", value)
}

// ContentClassName 内容类名
func (t Tabs) ContentClassName(value string) Tabs {
	return t.set("contentClassName", value)
}

// DefaultKey 初始化激活的选项卡
func (t Tabs) DefaultKey(value string) Tabs {
	return t.set("defaultKey", value)
}

// Disabled 是否禁用
func (t Tabs) Disabled(value bool) Tabs {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (t Tabs) DisabledOn(value string) Tabs {
	return t.set("disabledOn", value)
}

// Draggable 是否支持拖拽
func (t Tabs) Draggable(value bool) Tabs {
	return t.set("draggable", value)
}

// Editable 是否可编辑标签名
func (t Tabs) Editable(value bool) Tabs {
	return t.set("editable", value)
}

// EditorSetting 编辑器配置
func (t Tabs) EditorSetting(value string) Tabs {
	return t.set("editorSetting", value)
}

// Hidden 是否隐藏
func (t Tabs) Hidden(value bool) Tabs {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (t Tabs) HiddenOn(value string) Tabs {
	return t.set("hiddenOn", value)
}

// Id 组件唯一 id
func (t Tabs) Id(value string) Tabs {
	return t.set("id", value)
}

// LinksClassName 链接外层类名
func (t Tabs) LinksClassName(value string) Tabs {
	return t.set("linksClassName", value)
}

// MountOnEnter 卡片是否只有在点开的时候加载？
func (t Tabs) MountOnEnter(value bool) Tabs {
	return t.set("mountOnEnter", value)
}

// OnEvent 事件动作配置
func (t Tabs) OnEvent(value string) Tabs {
	return t.set("onEvent", value)
}

// Scrollable 是否导航支持内容溢出滚动
func (t Tabs) Scrollable(value bool) Tabs {
	return t.set("scrollable", value)
}

// ShowTip 是否显示提示
func (t Tabs) ShowTip(value bool) Tabs {
	return t.set("showTip", value)
}

// ShowTipClassName tooltip 提示的类名
func (t Tabs) ShowTipClassName(value string) Tabs {
	return t.set("showTipClassName", value)
}

// SidePosition 编辑器模式，侧边的位置
func (t Tabs) SidePosition(value string) Tabs {
	return t.set("sidePosition", value)
}

// Source 关联已有数据
func (t Tabs) Source(value string) Tabs {
	return t.set("source", value)
}

// Static 是否静态展示
func (t Tabs) Static(value bool) Tabs {
	return t.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (t Tabs) StaticClassName(value string) Tabs {
	return t.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (t Tabs) StaticInputClassName(value string) Tabs {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (t Tabs) StaticLabelClassName(value string) Tabs {
	return t.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (t Tabs) StaticOn(value string) Tabs {
	return t.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (t Tabs) StaticPlaceholder(value string) Tabs {
	return t.set("staticPlaceholder", value)
}

// StaticSchema
func (t Tabs) StaticSchema(value string) Tabs {
	return t.set("staticSchema", value)
}

// Style 组件样式
func (t Tabs) Style(value string) Tabs {
	return t.set("style", value)
}

// SubFormHorizontal 如果是水平排版，这个属性可以细化水平排版的左右宽度占比
func (t Tabs) SubFormHorizontal(value string) Tabs {
	return t.set("subFormHorizontal", value)
}

// SubFormMode 配置子表单项默认的展示方式
func (t Tabs) SubFormMode(value string) Tabs {
	return t.set("subFormMode", value)
}

// Swipeable 是否滑动切换只在移动端生效
func (t Tabs) Swipeable(value bool) Tabs {
	return t.set("swipeable", value)
}

// Tabs 选项卡成员
func (t Tabs) Tabs(value string) Tabs {
	return t.set("tabs", value)
}

// TabsMode 展示形式
func (t Tabs) TabsMode(value string) Tabs {
	return t.set("tabsMode", value)
}

// TestIdBuilder
func (t Tabs) TestIdBuilder(value string) Tabs {
	return t.set("testIdBuilder", value)
}

// Testid
func (t Tabs) Testid(value string) Tabs {
	return t.set("testid", value)
}

// Toolbar 可以在右侧配置点其他功能按钮
func (t Tabs) Toolbar(value string) Tabs {
	return t.set("toolbar", value)
}

// UnmountOnExit 卡片隐藏的时候是否销毁卡片内容
func (t Tabs) UnmountOnExit(value bool) Tabs {
	return t.set("unmountOnExit", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (t Tabs) UseMobileUI(value bool) Tabs {
	return t.set("useMobileUI", value)
}

// Visible 是否显示
func (t Tabs) Visible(value bool) Tabs {
	return t.set("visible", value)
}

// VisibleOn 是否显示表达式
func (t Tabs) VisibleOn(value string) Tabs {
	return t.set("visibleOn", value)
}
