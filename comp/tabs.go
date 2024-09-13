package comp

// tabs 选项卡控件。 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/tabs

type tabs schema

// Tabs 创建一个新的 Tabs 实例
func Tabs() tabs {
	return tabs{}.set("type", "tabs")
}

func (t tabs) set(key string, value any) tabs {
	t[key] = value
	return t
}

// ActiveKey 激活的选项卡，hash 值或索引值，支持使用表达式，可响应上下文数据变化
func (t tabs) ActiveKey(value any) tabs {
	return t.set("activeKey", value)
}

// AddBtnText 自定义增加按钮文案
func (t tabs) AddBtnText(value string) tabs {
	return t.set("addBtnText", value)
}

// Addable 是否支持新增
func (t tabs) Addable(value bool) tabs {
	return t.set("addable", value)
}

// ClassName 容器 css 类名
func (t tabs) ClassName(value string) tabs {
	return t.set("className", value)
}

// Closable 是否支持删除
func (t tabs) Closable(value bool) tabs {
	return t.set("closable", value)
}

// CollapseBtnLabel 折叠按钮文字
func (t tabs) CollapseBtnLabel(value string) tabs {
	return t.set("collapseBtnLabel", value)
}

// CollapseOnExceed 超过多少个时折叠按钮
func (t tabs) CollapseOnExceed(value string) tabs {
	return t.set("collapseOnExceed", value)
}

// ContentClassName 内容类名
func (t tabs) ContentClassName(value string) tabs {
	return t.set("contentClassName", value)
}

// DefaultKey 组件初始化时激活的选项卡，hash 值或索引值，支持使用表达式
func (t tabs) DefaultKey(value any) tabs {
	return t.set("defaultKey", value)
}

// Disabled 是否禁用
func (t tabs) Disabled(value bool) tabs {
	return t.set("disabled", value)
}

// DisabledOn 是否禁用表达式
func (t tabs) DisabledOn(value string) tabs {
	return t.set("disabledOn", value)
}

// Draggable 是否支持拖拽
func (t tabs) Draggable(value bool) tabs {
	return t.set("draggable", value)
}

// Editable 是否可编辑标签名
func (t tabs) Editable(value bool) tabs {
	return t.set("editable", value)
}

// EditorSetting 编辑器配置
func (t tabs) EditorSetting(value string) tabs {
	return t.set("editorSetting", value)
}

// Hidden 是否隐藏
func (t tabs) Hidden(value bool) tabs {
	return t.set("hidden", value)
}

// HiddenOn 是否隐藏表达式
func (t tabs) HiddenOn(value string) tabs {
	return t.set("hiddenOn", value)
}

// Id 组件唯一 id
func (t tabs) Id(value string) tabs {
	return t.set("id", value)
}

// LinksClassName 链接外层类名
func (t tabs) LinksClassName(value string) tabs {
	return t.set("linksClassName", value)
}

// MountOnEnter 卡片是否只有在点开的时候加载？
func (t tabs) MountOnEnter(value bool) tabs {
	return t.set("mountOnEnter", value)
}

// OnEvent 事件动作配置
func (t tabs) OnEvent(value any) tabs {
	return t.set("onEvent", value)
}

// Scrollable 是否导航支持内容溢出滚动
func (t tabs) Scrollable(value bool) tabs {
	return t.set("scrollable", value)
}

// ShowTip 是否显示提示
func (t tabs) ShowTip(value bool) tabs {
	return t.set("showTip", value)
}

// ShowTipClassName tooltip 提示的类名
func (t tabs) ShowTipClassName(value string) tabs {
	return t.set("showTipClassName", value)
}

// SidePosition 编辑器模式，侧边的位置 left/right
func (t tabs) SidePosition(value string) tabs {
	return t.set("sidePosition", value)
}

// Source 关联已有数据
func (t tabs) Source(value string) tabs {
	return t.set("source", value)
}

// Static 是否静态展示
func (t tabs) Static(value bool) tabs {
	return t.set("static", value)
}

// StaticClassName 静态展示表单项类名
func (t tabs) StaticClassName(value string) tabs {
	return t.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名
func (t tabs) StaticInputClassName(value string) tabs {
	return t.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名
func (t tabs) StaticLabelClassName(value string) tabs {
	return t.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式
func (t tabs) StaticOn(value string) tabs {
	return t.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (t tabs) StaticPlaceholder(value string) tabs {
	return t.set("staticPlaceholder", value)
}

// StaticSchema
func (t tabs) StaticSchema(value string) tabs {
	return t.set("staticSchema", value)
}

// Style 组件样式
func (t tabs) Style(value any) tabs {
	return t.set("style", value)
}

// SubFormHorizontal 如果是水平排版，这个属性可以细化水平排版的左右宽度占比
func (t tabs) SubFormHorizontal(value string) tabs {
	return t.set("subFormHorizontal", value)
}

// SubFormMode 配置子表单项默认的展示方式
func (t tabs) SubFormMode(value string) tabs {
	return t.set("subFormMode", value)
}

// Swipeable 是否滑动切换只在移动端生效
func (t tabs) Swipeable(value bool) tabs {
	return t.set("swipeable", value)
}

// Tabs 选项卡成员
func (t tabs) Tabs(value ...any) tabs {
	return t.set("tabs", value)
}

// TabsMode 展示模式，取值可以是 line、card、radio、vertical、chrome、simple、strong、tiled、sidebar
func (t tabs) TabsMode(value string) tabs {
	return t.set("tabsMode", value)
}

// TestIdBuilder
func (t tabs) TestIdBuilder(value string) tabs {
	return t.set("testIdBuilder", value)
}

// Testid
func (t tabs) Testid(value string) tabs {
	return t.set("testid", value)
}

// Toolbar 可以在右侧配置点其他功能按钮
func (t tabs) Toolbar(value string) tabs {
	return t.set("toolbar", value)
}

// UnmountOnExit 卡片隐藏的时候是否销毁卡片内容
func (t tabs) UnmountOnExit(value bool) tabs {
	return t.set("unmountOnExit", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (t tabs) UseMobileUI(value bool) tabs {
	return t.set("useMobileUI", value)
}

// Visible 是否显示
func (t tabs) Visible(value bool) tabs {
	return t.set("visible", value)
}

// VisibleOn 是否显示表达式
func (t tabs) VisibleOn(value string) tabs {
	return t.set("visibleOn", value)
}
