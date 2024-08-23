package comp

// nav 导航渲染器
type nav schema

// Nav 创建一个新的 Nav 实例
func Nav() nav {
	return nav{}.set("type", "nav")
}

// set 方法用于设置属性并返回自身
func (n nav) set(key string, value any) nav {
	n[key] = value
	return n
}

// Accordion 设置手风琴展开，仅垂直inline模式支持
func (n nav) Accordion(value bool) nav {
	return n.set("accordion", value)
}

// Badge 设置角标
func (n nav) Badge(value string) nav {
	return n.set("badge", value)
}

// ClassName 设置容器 css 类名
func (n nav) ClassName(value string) nav {
	return n.set("className", value)
}

// Collapsed 设置控制菜单缩起
func (n nav) Collapsed(value bool) nav {
	return n.set("collapsed", value)
}

// DefaultOpenLevel 设置默认展开层级
func (n nav) DefaultOpenLevel(value string) nav {
	return n.set("defaultOpenLevel", value)
}

// DeferApi 设置懒加载 api
func (n nav) DeferApi(value string) nav {
	return n.set("deferApi", value)
}

// Disabled 设置是否禁用
func (n nav) Disabled(value bool) nav {
	return n.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (n nav) DisabledOn(value string) nav {
	return n.set("disabledOn", value)
}

// DragOnSameLevel 设置仅允许同层级拖拽
func (n nav) DragOnSameLevel(value bool) nav {
	return n.set("dragOnSameLevel", value)
}

// Draggable 设置可拖拽
func (n nav) Draggable(value bool) nav {
	return n.set("draggable", value)
}

// EditorSetting 设置编辑器配置
func (n nav) EditorSetting(value string) nav {
	return n.set("editorSetting", value)
}

// ExpandIcon 设置自定义展开图标
func (n nav) ExpandIcon(value string) nav {
	return n.set("expandIcon", value)
}

// ExpandPosition 设置自定义展开图标位置
func (n nav) ExpandPosition(value string) nav {
	return n.set("expandPosition", value)
}

// Hidden 设置是否隐藏
func (n nav) Hidden(value bool) nav {
	return n.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (n nav) HiddenOn(value string) nav {
	return n.set("hiddenOn", value)
}

// Id 设置组件唯一 id
func (n nav) Id(value string) nav {
	return n.set("id", value)
}

// IndentSize 设置缩进大小
func (n nav) IndentSize(value string) nav {
	return n.set("indentSize", value)
}

// ItemActions 设置更多操作菜单列表
func (n nav) ItemActions(value string) nav {
	return n.set("itemActions", value)
}

// ItemBadge 设置角标
func (n nav) ItemBadge(value string) nav {
	return n.set("itemBadge", value)
}

// Level 设置最多展示多少层级
func (n nav) Level(value string) nav {
	return n.set("level", value)
}

// Links 设置链接地址集合 TODO
func (n nav) Links(value string) nav {
	return n.set("links", value)
}

// Mode 设置垂直模式下的菜单打开方式
func (n nav) Mode(value string) nav {
	return n.set("mode", value)
}

// OnEvent 设置事件动作配置
func (n nav) OnEvent(value any) nav {
	return n.set("onEvent", value)
}

// Overflow 设置横向导航时自动收纳配置
func (n nav) Overflow(value string) nav {
	return n.set("overflow", value)
}

// PopupClassName 设置子菜单项展开浮层样式
func (n nav) PopupClassName(value string) nav {
	return n.set("popupClassName", value)
}

// SaveOrderApi 设置保存排序的 api
func (n nav) SaveOrderApi(value string) nav {
	return n.set("saveOrderApi", value)
}

// SearchConfig 设置搜索框相关配置
func (n nav) SearchConfig(value string) nav {
	return n.set("searchConfig", value)
}

// Searchable 设置是否开启搜索
func (n nav) Searchable(value bool) nav {
	return n.set("searchable", value)
}

// ShowKey 控制仅展示指定key菜单下的子菜单项
func (n nav) ShowKey(value string) nav {
	return n.set("showKey", value)
}

// Source 设置可以通过 API 拉取
func (n nav) Source(value string) nav {
	return n.set("source", value)
}

// Stacked 设置排列方式，true 为垂直排列，false 为水平排列
func (n nav) Stacked(value bool) nav {
	return n.set("stacked", value)
}

// Static 设置是否静态展示
func (n nav) Static(value bool) nav {
	return n.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (n nav) StaticClassName(value string) nav {
	return n.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项 Value 类名
func (n nav) StaticInputClassName(value string) nav {
	return n.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项 Label 类名
func (n nav) StaticLabelClassName(value string) nav {
	return n.set("staticLabelClassName", value)
}

// StaticOn 设置静态展示表达式
func (n nav) StaticOn(value string) nav {
	return n.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (n nav) StaticPlaceholder(value string) nav {
	return n.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示 schema
func (n nav) StaticSchema(value string) nav {
	return n.set("staticSchema", value)
}

// Style 设置组件样式
func (n nav) Style(value any) nav {
	return n.set("style", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (n nav) TestIdBuilder(value string) nav {
	return n.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (n nav) Testid(value string) nav {
	return n.set("testid", value)
}

// ThemeColor 设置主题配色
func (n nav) ThemeColor(value string) nav {
	return n.set("themeColor", value)
}

// UseMobileUI 设置是否关闭移动端样式
func (n nav) UseMobileUI(value bool) nav {
	return n.set("useMobileUI", value)
}

// Visible 设置是否显示
func (n nav) Visible(value bool) nav {
	return n.set("visible", value)
}

// VisibleOn 设置显示表达式
func (n nav) VisibleOn(value string) nav {
	return n.set("visibleOn", value)
}
