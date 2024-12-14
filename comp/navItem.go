package comp

// navItem 导航项
type navItem schema

// NavItem 创建一个新的 NavItem 实例
func NavItem() navItem {
	return navItem{}
}

// set 方法用于设置属性并返回自身
func (n navItem) set(key string, value any) navItem {
	n[key] = value
	return n
}

// Active 设置是否激活
func (n navItem) Active(value bool) navItem {
	return n.set("active", value)
}

// Children 设置子项
func (n navItem) Children(value string) navItem {
	return n.set("children", value)
}

// ClassName 设置 css 类名
func (n navItem) ClassName(value string) navItem {
	return n.set("className", value)
}

// Defer 设置是否延迟加载
func (n navItem) Defer(value bool) navItem {
	return n.set("defer", value)
}

// DeferApi 设置延迟加载的 API
func (n navItem) DeferApi(value string) navItem {
	return n.set("deferApi", value)
}

// Disabled 设置是否禁用
func (n navItem) Disabled(value bool) navItem {
	return n.set("disabled", value)
}

// DisabledOn 设置禁用表达式
func (n navItem) DisabledOn(value string) navItem {
	return n.set("disabledOn", value)
}

// DisabledTip 设置禁用提示
func (n navItem) DisabledTip(value string) navItem {
	return n.set("disabledTip", value)
}

// EditorSetting 设置编辑器配置
func (n navItem) EditorSetting(value string) navItem {
	return n.set("editorSetting", value)
}

// Hidden 设置是否隐藏
func (n navItem) Hidden(value bool) navItem {
	return n.set("hidden", value)
}

// HiddenOn 设置隐藏表达式
func (n navItem) HiddenOn(value string) navItem {
	return n.set("hiddenOn", value)
}

// Icon 设置图标类名
func (n navItem) Icon(value string) navItem {
	return n.set("icon", value)
}

// Id 设置组件唯一 id
func (n navItem) ID(value string) navItem {
	return n.set("id", value)
}

// Key 设置唯一键
func (n navItem) Key(value string) navItem {
	return n.set("key", value)
}

// Label 设置文字说明
func (n navItem) Label(value string) navItem {
	return n.set("label", value)
}

// Mode 设置模式
func (n navItem) Mode(value string) navItem {
	return n.set("mode", value)
}

// OnEvent 设置事件动作配置
func (n navItem) OnEvent(value any) navItem {
	return n.set("onEvent", value)
}

// Static 设置是否静态展示
func (n navItem) Static(value bool) navItem {
	return n.set("static", value)
}

// StaticClassName 设置静态展示表单项类名
func (n navItem) StaticClassName(value string) navItem {
	return n.set("staticClassName", value)
}

// StaticInputClassName 设置静态展示表单项 Value 类名
func (n navItem) StaticInputClassName(value string) navItem {
	return n.set("staticInputClassName", value)
}

// StaticLabelClassName 设置静态展示表单项 Label 类名
func (n navItem) StaticLabelClassName(value string) navItem {
	return n.set("staticLabelClassName", value)
}

// StaticOn 设置静态展示表达式
func (n navItem) StaticOn(value string) navItem {
	return n.set("staticOn", value)
}

// StaticPlaceholder 设置静态展示空值占位
func (n navItem) StaticPlaceholder(value string) navItem {
	return n.set("staticPlaceholder", value)
}

// StaticSchema 设置静态展示 schema
func (n navItem) StaticSchema(value string) navItem {
	return n.set("staticSchema", value)
}

// Style 设置组件样式
func (n navItem) Style(value any) navItem {
	return n.set("style", value)
}

// Target 设置目标
func (n navItem) Target(value string) navItem {
	return n.set("target", value)
}

// TestIdBuilder 设置测试 ID 构建器
func (n navItem) TestIdBuilder(value string) navItem {
	return n.set("testIdBuilder", value)
}

// Testid 设置测试 ID
func (n navItem) Testid(value string) navItem {
	return n.set("testid", value)
}

// To 设置导航目标
func (n navItem) To(value string) navItem {
	return n.set("to", value)
}

// Unfolded 设置是否展开
func (n navItem) Unfolded(value bool) navItem {
	return n.set("unfolded", value)
}

// UseMobileUI 设置是否关闭移动端样式
func (n navItem) UseMobileUI(value bool) navItem {
	return n.set("useMobileUI", value)
}

// Visible 设置是否显示
func (n navItem) Visible(value bool) navItem {
	return n.set("visible", value)
}

// VisibleOn 设置显示表达式
func (n navItem) VisibleOn(value string) navItem {
	return n.set("visibleOn", value)
}
