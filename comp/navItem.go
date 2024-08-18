package comp

// NavItem 导航项
type NavItem Schema

// NewNavItem 创建一个新的 NavItem 实例
func NewNavItem() NavItem {
	return NavItem{}
}

// set 方法用于设置属性并返回自身
func (n NavItem) set(key string, value interface{}) NavItem {
	n[key] = value
	return n
}

// Active 设置是否激活
func (n NavItem) Active(value bool) NavItem {
	return n.set("active", value)
}

// Children 设置子项
func (n NavItem) Children(value string) NavItem {
	return n.set("children", value)
}

// ClassName 设置 css 类名
func (n NavItem) ClassName(value string) NavItem {
	return n.set("className", value)
}

// Defer 设置是否延迟加载
func (n NavItem) Defer(value bool) NavItem {
	return n.set("defer", value)
}

// DeferApi 设置延迟加载的 API
func (n NavItem) DeferApi(value string) NavItem {
	return n.set("deferApi", value)
}

// Disabled 设置是否禁用
func (n NavItem) Disabled(value bool) NavItem {
	return n.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (n NavItem) DisabledOn(value string) NavItem {
	return n.set("disabledOn", value)
}

// DisabledTip 设置禁用提示
func (n NavItem) DisabledTip(value string) NavItem {
	return n.set("disabledTip", value)
}

// EditorSetting 设置编辑器配置
func (n NavItem) EditorSetting(value string) NavItem {
	return n.set("editorSetting", value)
}

// Hidden 设置是否隐藏
func (n NavItem) Hidden(value bool) NavItem {
	return n.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (n NavItem) HiddenOn(value string) NavItem {
	return n.set("hiddenOn", value)
}

// Icon 设置图标类名
func (n NavItem) Icon(value string) NavItem {
	return n.set("icon", value)
}

// Id 设置组件唯一 id
func (n NavItem) Id(value string) NavItem {
	return n.set("id", value)
}

// Key 设置唯一键
func (n NavItem) Key(value string) NavItem {
	return n.set("key", value)
}

// Label 设置文字说明
func (n NavItem) Label(value string) NavItem {
	return n.set("label", value)
}

// Mode 设置模式
func (n NavItem) Mode(value string) NavItem {
	return n.set("mode", value)
}

// OnEvent 设置事件动作配置
func (n NavItem) OnEvent(value string) NavItem {
	return n.set("onEvent", value)
}

// Static 设置是否静态展示
func (n NavItem) Static(value bool) NavItem {
	return n.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (n NavItem) StaticClassName(value string) NavItem {
	return n.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项 Value 类名
func (n NavItem) StaticInputClassName(value string) NavItem {
	return n.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项 Label 类名
func (n NavItem) StaticLabelClassName(value string) NavItem {
	return n.set("staticLabelClassName", value)
}

// StaticOn 设置静态展示表达式
func (n NavItem) StaticOn(value string) NavItem {
	return n.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (n NavItem) StaticPlaceholder(value string) NavItem {
	return n.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示 schema
func (n NavItem) StaticSchema(value string) NavItem {
	return n.set("staticSchema", value)
}

// Style 设置组件样式
func (n NavItem) Style(value string) NavItem {
	return n.set("style", value)
}

// Target 设置目标
func (n NavItem) Target(value string) NavItem {
	return n.set("target", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (n NavItem) TestIdBuilder(value string) NavItem {
	return n.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (n NavItem) Testid(value string) NavItem {
	return n.set("testid", value)
}

// To 设置导航目标
func (n NavItem) To(value string) NavItem {
	return n.set("to", value)
}

// Unfolded 设置是否展开
func (n NavItem) Unfolded(value bool) NavItem {
	return n.set("unfolded", value)
}

// UseMobileUI 设置是否关闭移动端样式
func (n NavItem) UseMobileUI(value bool) NavItem {
	return n.set("useMobileUI", value)
}

// Visible 设置是否显示
func (n NavItem) Visible(value bool) NavItem {
	return n.set("visible", value)
}

// VisibleOn 设置显示表达式
func (n NavItem) VisibleOn(value string) NavItem {
	return n.set("visibleOn", value)
}
