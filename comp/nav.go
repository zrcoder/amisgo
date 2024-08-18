package comp

// Nav 导航渲染器
type Nav Schema

// NewNav 创建一个新的 Nav 实例
func NewNav() Nav {
	return Nav{}.set("type", "nav")
}

// set 方法用于设置属性并返回自身
func (n Nav) set(key string, value interface{}) Nav {
	n[key] = value
	return n
}

// Accordion 设置手风琴展开，仅垂直inline模式支持
func (n Nav) Accordion(value bool) Nav {
	return n.set("accordion", value)
}

// Badge 设置角标
func (n Nav) Badge(value string) Nav {
	return n.set("badge", value)
}

// ClassName 设置容器 css 类名
func (n Nav) ClassName(value string) Nav {
	return n.set("className", value)
}

// Collapsed 设置控制菜单缩起
func (n Nav) Collapsed(value bool) Nav {
	return n.set("collapsed", value)
}

// DefaultOpenLevel 设置默认展开层级
func (n Nav) DefaultOpenLevel(value string) Nav {
	return n.set("defaultOpenLevel", value)
}

// DeferApi 设置懒加载 api
func (n Nav) DeferApi(value string) Nav {
	return n.set("deferApi", value)
}

// Disabled 设置是否禁用
func (n Nav) Disabled(value bool) Nav {
	return n.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (n Nav) DisabledOn(value string) Nav {
	return n.set("disabledOn", value)
}

// DragOnSameLevel 设置仅允许同层级拖拽
func (n Nav) DragOnSameLevel(value bool) Nav {
	return n.set("dragOnSameLevel", value)
}

// Draggable 设置可拖拽
func (n Nav) Draggable(value bool) Nav {
	return n.set("draggable", value)
}

// EditorSetting 设置编辑器配置
func (n Nav) EditorSetting(value string) Nav {
	return n.set("editorSetting", value)
}

// ExpandIcon 设置自定义展开图标
func (n Nav) ExpandIcon(value string) Nav {
	return n.set("expandIcon", value)
}

// ExpandPosition 设置自定义展开图标位置
func (n Nav) ExpandPosition(value string) Nav {
	return n.set("expandPosition", value)
}

// Hidden 设置是否隐藏
func (n Nav) Hidden(value bool) Nav {
	return n.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (n Nav) HiddenOn(value string) Nav {
	return n.set("hiddenOn", value)
}

// Id 设置组件唯一 id
func (n Nav) Id(value string) Nav {
	return n.set("id", value)
}

// IndentSize 设置缩进大小
func (n Nav) IndentSize(value string) Nav {
	return n.set("indentSize", value)
}

// ItemActions 设置更多操作菜单列表
func (n Nav) ItemActions(value string) Nav {
	return n.set("itemActions", value)
}

// ItemBadge 设置角标
func (n Nav) ItemBadge(value string) Nav {
	return n.set("itemBadge", value)
}

// Level 设置最多展示多少层级
func (n Nav) Level(value string) Nav {
	return n.set("level", value)
}

// Links 设置链接地址集合
func (n Nav) Links(value string) Nav {
	return n.set("links", value)
}

// Mode 设置垂直模式下的菜单打开方式
func (n Nav) Mode(value string) Nav {
	return n.set("mode", value)
}

// OnEvent 设置事件动作配置
func (n Nav) OnEvent(value string) Nav {
	return n.set("onEvent", value)
}

// Overflow 设置横向导航时自动收纳配置
func (n Nav) Overflow(value string) Nav {
	return n.set("overflow", value)
}

// PopupClassName 设置子菜单项展开浮层样式
func (n Nav) PopupClassName(value string) Nav {
	return n.set("popupClassName", value)
}

// SaveOrderApi 设置保存排序的 api
func (n Nav) SaveOrderApi(value string) Nav {
	return n.set("saveOrderApi", value)
}

// SearchConfig 设置搜索框相关配置
func (n Nav) SearchConfig(value string) Nav {
	return n.set("searchConfig", value)
}

// Searchable 设置是否开启搜索
func (n Nav) Searchable(value bool) Nav {
	return n.set("searchable", value)
}

// ShowKey 控制仅展示指定key菜单下的子菜单项
func (n Nav) ShowKey(value string) Nav {
	return n.set("showKey", value)
}

// Source 设置可以通过 API 拉取
func (n Nav) Source(value string) Nav {
	return n.set("source", value)
}

// Stacked 设置排列方式，true 为垂直排列，false 为水平排列
func (n Nav) Stacked(value bool) Nav {
	return n.set("stacked", value)
}

// Static 设置是否静态展示
func (n Nav) Static(value bool) Nav {
	return n.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (n Nav) StaticClassName(value string) Nav {
	return n.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项 Value 类名
func (n Nav) StaticInputClassName(value string) Nav {
	return n.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项 Label 类名
func (n Nav) StaticLabelClassName(value string) Nav {
	return n.set("staticLabelClassName", value)
}

// StaticOn 设置静态展示表达式
func (n Nav) StaticOn(value string) Nav {
	return n.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (n Nav) StaticPlaceholder(value string) Nav {
	return n.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示 schema
func (n Nav) StaticSchema(value string) Nav {
	return n.set("staticSchema", value)
}

// Style 设置组件样式
func (n Nav) Style(value string) Nav {
	return n.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (n Nav) TestIdBuilder(value string) Nav {
	return n.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (n Nav) Testid(value string) Nav {
	return n.set("testid", value)
}

// ThemeColor 设置主题配色
func (n Nav) ThemeColor(value string) Nav {
	return n.set("themeColor", value)
}

// UseMobileUI 设置是否关闭移动端样式
func (n Nav) UseMobileUI(value bool) Nav {
	return n.set("useMobileUI", value)
}

// Visible 设置是否显示
func (n Nav) Visible(value bool) Nav {
	return n.set("visible", value)
}

// VisibleOn 设置显示表达式
func (n Nav) VisibleOn(value string) Nav {
	return n.set("visibleOn", value)
}
